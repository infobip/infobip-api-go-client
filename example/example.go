package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/infobip/infobip-api-go-client/pkg"
	"github.com/infobip/infobip-api-go-client/pkg/sms"
)

func main() {
	username, password, phone := "username", "password", "phoneNumber"

	client := sms.Client{
		BaseURL:    "https://api.infobip.com",
		Authorizer: infobip.NewBasicCredentials(username, password),
		HTTPCLient: http.DefaultClient,
	}

	sendSms(client, newMultiMessage(phone))
	sendSms(client, newMessageWithScheduling(phone))
	sendSms(client, newMessageWithURLTracking(phone))
	sendSms(client, newMessageWithNotifications(phone))

	response := sendSms(client, newSimpleMessage(phone))

	time.Sleep(time.Minute * 5)

	if response != nil {
		messageID := response.Messages[0].MessageID
		query := sms.GetSentSmsDeliveryReportsQuery{
			MessageID: messageID,
			Limit:     15,
		}
		retrieveStatus(client, query)
	}

}

func sendSms(client sms.Client, req sms.SMSAdvancedTextualRequest) *sms.SMSResponse {
	res, err := client.SendMultipleTextualSmsAdvanced(req)

	if err != nil {
		fmt.Println("SMS sending error:")
		fmt.Println(err)
		return nil
	}

	fmt.Println("SMS sending success:")
	fmt.Printf("%+v\n", *res)
	return res
}

func retrieveStatus(client sms.Client, query sms.GetSentSmsDeliveryReportsQuery) *sms.SMSReportResponse {
	res, err := client.GetSentSmsDeliveryReports(query)

	if err != nil {
		fmt.Println("Status retrieval error:")
		fmt.Println(err)
		return nil
	}

	fmt.Println("Status retrieval success:")
	fmt.Printf("%+v\n", *res)
	return res
}

func newSimpleMessage(destination string) sms.SMSAdvancedTextualRequest {
	return sms.SMSAdvancedTextualRequest{
		Messages: []sms.Message{
			sms.Message{
				Destinations: []sms.Destination{sms.Destination{To: destination}},
				From:         "InfoSMS",
				Text:         "SMS from go library",
			},
		},
	}
}

func newMessageWithScheduling(destination string) sms.SMSAdvancedTextualRequest {
	future := time.Now().Add(time.Minute * 10)
	return sms.SMSAdvancedTextualRequest{
		Messages: []sms.Message{
			sms.Message{
				Destinations:   []sms.Destination{sms.Destination{To: destination}},
				From:           "InfoSMS",
				Text:           fmt.Sprintf("SMS scheduled for %v", future),
				SendAt:         &infobip.Time{T: future},
				ValidityPeriod: 10000,
				DeliveryTimeWindow: &sms.DeliveryTimeWindow{
					Days: []sms.DeliveryDay{
						sms.MONDAY,
						sms.TUESDAY,
						sms.WEDNESDAY,
						sms.THURSDAY,
						sms.FRIDAY,
					},
				},
			},
		},
	}
}

func newMessageWithNotifications(destination string) sms.SMSAdvancedTextualRequest {
	notifyURL := fmt.Sprintf("http://echo-http-requests.appspot.com/push/sfonsv%d", time.Now().UTC().UnixNano())
	return sms.SMSAdvancedTextualRequest{
		Messages: []sms.Message{
			sms.Message{
				Destinations:       []sms.Destination{sms.Destination{To: destination}},
				From:               "InfoSMS",
				Text:               fmt.Sprintf("SMS with notifyURL %s", notifyURL),
				Notify:             true,
				IntermediateReport: true,
				NotifyURL:          notifyURL,
				CallbackData:       "Go lib callback data",
			},
		},
	}
}

func newMessageWithURLTracking(destination string) sms.SMSAdvancedTextualRequest {
	return sms.SMSAdvancedTextualRequest{
		Messages: []sms.Message{
			sms.Message{
				Destinations: []sms.Destination{sms.Destination{To: destination}},
				From:         "InfoSMS",
				Text:         "SMS with URL tracking: https://dev.infobip.com/docs/url-shortening <- find out more!",
			},
		},
		Tracking: &sms.Tracking{Track: "URL"},
	}
}

func newMultiMessage(destination string) sms.SMSAdvancedTextualRequest {
	ts := time.Now().UTC().UnixNano()
	return sms.SMSAdvancedTextualRequest{
		BulkID: fmt.Sprintf("go-lib-bulk-%d", ts),
		Messages: []sms.Message{
			sms.Message{
				Destinations: []sms.Destination{
					sms.Destination{To: destination, MessageID: fmt.Sprintf("go-lib-msg-1-to-1-%d", ts)},
					sms.Destination{To: destination, MessageID: fmt.Sprintf("go-lib-msg-1-to-2-%d", ts)},
				},
				From: "InfoSMS",
				Text: "SMS bulk #1 text",
			},
			sms.Message{
				Destinations: []sms.Destination{
					sms.Destination{To: destination, MessageID: fmt.Sprintf("go-lib-msg-2-to-1-%d", ts)},
					sms.Destination{To: destination, MessageID: fmt.Sprintf("go-lib-msg-2-to-2-%d", ts)},
				},
				From: "InfoSMS",
				Text: "SMS bulk #2 text",
			},
		},
	}
}
