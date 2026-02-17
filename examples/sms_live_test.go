package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
)

func TestSendSMS(t *testing.T) {
	infobipClient, auth := newClientAndAuth()

	destinations := []sms.Destination{
		{To: "<DESTINATION>"},
	}

	content := sms.MessageContent{
		TextMessageContent: sms.NewTextMessageContent("Congratulations on sending your first message with GO library."),
	}

	givenMessage := sms.NewMessage(destinations, content)
	givenMessage.SetSender("<SENDER>")

	request := sms.NewRequestEnvelope([]sms.Message{
		*givenMessage,
	})

	// Send the SMS message
	apiResponse, httpResponse, err := infobipClient.
		SmsAPI.
		SendSmsMessages(auth).
		RequestEnvelope(*request).
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to send SMS: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestSendAdvancedSMSFeatures(t *testing.T) {
	client, auth := newClientAndAuth()

	destinations := []sms.Destination{
		{To: "<DESTINATION>"},
	}

	content := sms.MessageContent{
		TextMessageContent: sms.NewTextMessageContent("Artık Ulusal Dil Tanımlayıcısı ile Türkçe karakterli smslerinizi rahatlıkla iletebilirsiniz."),
	}
	content.TextMessageContent.SetTransliteration(sms.TRANSLITERATIONCODE_TURKISH)
	languageCode := sms.LANGUAGECODE_TR
	content.TextMessageContent.SetLanguage(sms.Language{LanguageCode: &languageCode})

	message := sms.NewMessage(destinations, content)
	message.SetSender("<SENDER>")
	validityPeriodTimeUnit := sms.VALIDITYPERIODTIMEUNIT_MINUTES
	message.SetOptions(sms.MessageOptions{
		ValidityPeriod:      &sms.ValidityPeriod{Amount: 720, TimeUnit: &validityPeriodTimeUnit},
		CampaignReferenceId: infobip.PtrString("summersale"),
	})
	message.SetWebhooks(sms.Webhooks{
		Delivery:     &sms.MessageDeliveryReporting{Url: infobip.PtrString("https://www.example.com/sms/advanced"), IntermediateReport: infobip.PtrBool(true)},
		ContentType:  infobip.PtrString("application/json"),
		CallbackData: infobip.PtrString("DLR callback data"),
	})

	request := sms.NewRequestEnvelope([]sms.Message{*message})
	sendAtTime := time.Now().Add(10 * time.Second).UTC()

	request.SetOptions(sms.MessageRequestOptions{
		Schedule: &sms.RequestSchedulingSettings{
			BulkId: infobip.PtrString("BULK-ID-123-xyz"),
			SendAt: &infobip.Time{T: sendAtTime},
		},
		Tracking: &sms.UrlOptions{
			ShortenUrl:     infobip.PtrBool(true),
			TrackClicks:    infobip.PtrBool(true),
			TrackingUrl:    infobip.PtrString("https://example.com/click-report"),
			RemoveProtocol: infobip.PtrBool(true),
			CustomDomain:   infobip.PtrString("example.com"),
		},
		IncludeSmsCountInResponse: infobip.PtrBool(true),
		ConversionTracking:        &sms.Tracking{UseConversionTracking: infobip.PtrBool(true), ConversionTrackingName: infobip.PtrString("MY_CAMPAIGN")},
	})

	apiResponse, httpResponse, err := client.SmsAPI.SendSmsMessages(auth).
		RequestEnvelope(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send advanced SMS: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestGetSmsDeliveryReports(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		SmsAPI.
		GetOutboundSmsMessageDeliveryReports(auth).
		Limit(10).
		Execute(); err != nil {
		t.Fatalf("Failed to fetch SMS delivery reports: %v (http=%+v)", err, httpResp)
	}
}

func TestGetSmsMessageLogs(t *testing.T) {
	client, auth := newClientAndAuth()

	since := infobip.Time{T: time.Now().Add(-24 * time.Hour)}
	until := infobip.Time{T: time.Now()}

	if _, httpResp, err := client.
		SmsAPI.
		GetOutboundSmsMessageLogs(auth).
		SentSince(since).
		SentUntil(until).
		Limit(10).
		Execute(); err != nil {
		t.Fatalf("Failed to fetch SMS message logs: %v (http=%+v)", err, httpResp)
	}
}

func TestPreviewSmsMessage(t *testing.T) {
	client, auth := newClientAndAuth()

	previewRequest := sms.NewPreviewRequest("Preview this message before sending.")
	previewRequest.SetLanguageCode("AUTODETECT")
	previewRequest.SetTransliteration("ALL")

	if _, httpResp, err := client.
		SmsAPI.
		PreviewSmsMessage(auth).
		PreviewRequest(*previewRequest).
		Execute(); err != nil {
		t.Fatalf("Failed to preview SMS message: %v (http=%+v)", err, httpResp)
	}
}

func TestGetInboundSmsMessages(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		SmsAPI.
		GetInboundSmsMessages(auth).
		Limit(10).
		Execute(); err != nil {
		t.Fatalf("Failed to fetch inbound SMS messages: %v (http=%+v)", err, httpResp)
	}
}

func TestGetScheduledSmsMessages(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		SmsAPI.
		GetScheduledSmsMessages(auth).
		BulkId("BULK-ID-123-xyz").
		Execute(); err != nil {
		t.Fatalf("Failed to get scheduled SMS messages: %v (http=%+v)", err, httpResp)
	}
}

func TestRescheduleSmsMessages(t *testing.T) {
	client, auth := newClientAndAuth()

	newSendAt := infobip.Time{T: time.Now().Add(1 * time.Minute)}
	request := sms.NewBulkRequest(newSendAt)

	if _, httpResp, err := client.
		SmsAPI.
		RescheduleSmsMessages(auth).
		BulkId("BULK-ID-123-xyz").
		BulkRequest(*request).
		Execute(); err != nil {
		t.Fatalf("Failed to reschedule SMS bulk: %v (http=%+v)", err, httpResp)
	}
}

func TestGetScheduledSmsMessagesStatus(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		SmsAPI.
		GetScheduledSmsMessagesStatus(auth).
		BulkId("BULK-ID-123-xyz").
		Execute(); err != nil {
		t.Fatalf("Failed to get scheduled SMS status: %v (http=%+v)", err, httpResp)
	}
}

func TestUpdateScheduledSmsMessagesStatus(t *testing.T) {
	client, auth := newClientAndAuth()

	updateRequest := sms.NewUpdateStatusRequest(sms.BULKSTATUS_PAUSED)

	if _, httpResp, err := client.
		SmsAPI.
		UpdateScheduledSmsMessagesStatus(auth).
		BulkId("BULK-ID-123-xyz").
		UpdateStatusRequest(*updateRequest).
		Execute(); err != nil {
		t.Fatalf("Failed to update scheduled SMS status: %v (http=%+v)", err, httpResp)
	}
}
