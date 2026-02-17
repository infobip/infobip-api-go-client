package examples

import (
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
)

func TestTfaFullBundledFlow(t *testing.T) {
	client, auth := newClientAndAuth()

	// --- Application lifecycle ------------------------------------------------
	appReq := sms.NewApplicationRequest("E2E TFA App")
	appReq.SetEnabled(true)
	cfg := sms.NewApplicationConfiguration()
	cfg.SetPinAttempts(5)
	cfg.SetAllowMultiplePinVerifications(true)
	cfg.SetPinTimeToLive("10m")
	appReq.SetConfiguration(*cfg)

	appResp, httpResp, err := client.
		TfaAPI.
		CreateTfaApplication(auth).
		ApplicationRequest(*appReq).
		Execute()
	if err != nil {
		t.Fatalf("Create app failed: %v (http=%+v)", err, httpResp)
	}
	appId := appResp.GetApplicationId()

	// Get single app
	if _, httpResp, err := client.TfaAPI.GetTfaApplication(auth, appId).Execute(); err != nil {
		t.Fatalf("Get app failed: %v (http=%+v)", err, httpResp)
	}

	// List apps
	apps, httpResp, err := client.TfaAPI.GetTfaApplications(auth).Execute()
	if err != nil || len(apps) == 0 {
		t.Fatalf("List apps failed: %v (http=%+v)", err, httpResp)
	}

	// Update app
	appUpdate := sms.NewApplicationRequest("E2E TFA App Updated")
	appUpdate.SetEnabled(true)
	appUpdate.SetConfiguration(*cfg)
	if _, httpResp, err := client.
		TfaAPI.
		UpdateTfaApplication(auth, appId).
		ApplicationRequest(*appUpdate).
		Execute(); err != nil {
		t.Fatalf("Update app failed: %v (http=%+v)", err, httpResp)
	}

	// --- Template creation (SMS, Voice, Email) -------------------------------
	smsTplReq := sms.NewCreateMessageRequest("Your code is {{pin}}", sms.PINTYPE_NUMERIC)
	smsTplReq.SetPinLength(4)
	smsTplReq.SetSenderId("Infobip 2FA")
	smsTplReq.SetLanguage(sms.TFALANGUAGE_EN)
	smsTplResp, httpResp, err := client.
		TfaAPI.
		CreateTfaMessageTemplate(auth, appId).
		CreateMessageRequest(*smsTplReq).
		Execute()
	if err != nil {
		t.Fatalf("Create SMS template failed: %v (http=%+v)", err, httpResp)
	}
	smsMsgId := smsTplResp.GetMessageId()

	voiceTplReq := sms.NewCreateMessageRequest("Voice code {{pin}}", sms.PINTYPE_NUMERIC)
	voiceTplReq.SetPinLength(4)
	voiceTplReq.SetLanguage(sms.TFALANGUAGE_EN)
	voiceTplReq.SetVoiceName("Joanna")
	voiceTplReq.SetRepeatDTMF("1#")
	voiceTplResp, httpResp, err := client.
		TfaAPI.
		CreateTfaMessageTemplate(auth, appId).
		CreateMessageRequest(*voiceTplReq).
		Execute()
	if err != nil {
		t.Fatalf("Create Voice template failed: %v (http=%+v)", err, httpResp)
	}
	voiceMsgId := voiceTplResp.GetMessageId()

	emailTplReq := sms.NewCreateEmailMessageRequest(1234) // replace with real broadcast email template ID
	emailTplReq.SetFrom("company@example.com")
	emailTplReq.SetPinType(sms.PINTYPE_NUMERIC)
	emailTplReq.SetPinLength(4)
	emailTplResp, httpResp, err := client.
		TfaAPI.
		CreateTfaEmailMessageTemplate(auth, appId).
		CreateEmailMessageRequest(*emailTplReq).
		Execute()
	if err != nil {
		t.Fatalf("Create Email template failed: %v (http=%+v)", err, httpResp)
	}
	emailMsgId := emailTplResp.GetMessageId()

	// --- List and get templates ---------------------------------------------
	msgs, httpResp, err := client.TfaAPI.GetTfaMessageTemplates(auth, appId).Execute()
	if err != nil || len(msgs) == 0 {
		t.Fatalf("List templates failed: %v (http=%+v)", err, httpResp)
	}
	if _, httpResp, err := client.TfaAPI.GetTfaMessageTemplate(auth, appId, smsMsgId).Execute(); err != nil {
		t.Fatalf("Get SMS template failed: %v (http=%+v)", err, httpResp)
	}
	if _, httpResp, err := client.TfaAPI.GetTfaMessageTemplate(auth, appId, voiceMsgId).Execute(); err != nil {
		t.Fatalf("Get Voice template failed: %v (http=%+v)", err, httpResp)
	}
	if _, httpResp, err := client.TfaAPI.GetTfaMessageTemplate(auth, appId, emailMsgId).Execute(); err != nil {
		t.Fatalf("Get Email template failed: %v (http=%+v)", err, httpResp)
	}

	// --- Update templates ---------------------------------------------------
	smsUpdate := sms.NewUpdateMessageRequest()
	smsUpdate.SetPinLength(6)
	smsUpdate.SetLanguage(sms.TFALANGUAGE_EN)
	smsUpdate.SetSenderId("Infobip 2FA")
	if _, httpResp, err := client.
		TfaAPI.
		UpdateTfaMessageTemplate(auth, appId, smsMsgId).
		UpdateMessageRequest(*smsUpdate).
		Execute(); err != nil {
		t.Fatalf("Update SMS template failed: %v (http=%+v)", err, httpResp)
	}

	voiceUpdate := sms.NewUpdateMessageRequest()
	voiceUpdate.SetPinLength(6)
	voiceUpdate.SetLanguage(sms.TFALANGUAGE_EN)
	voiceUpdate.SetVoiceName("Joanna")
	voiceUpdate.SetRepeatDTMF("2#")
	if _, httpResp, err := client.
		TfaAPI.
		UpdateTfaMessageTemplate(auth, appId, voiceMsgId).
		UpdateMessageRequest(*voiceUpdate).
		Execute(); err != nil {
		t.Fatalf("Update Voice template failed: %v (http=%+v)", err, httpResp)
	}

	emailUpdate := sms.NewUpdateEmailMessageRequest()
	emailUpdate.SetEmailTemplateId(1234)
	emailUpdate.SetFrom("company@example.com")
	emailUpdate.SetPinType(sms.PINTYPE_ALPHANUMERIC)
	emailUpdate.SetPinLength(6)
	if _, httpResp, err := client.
		TfaAPI.
		UpdateTfaEmailMessageTemplate(auth, appId, emailMsgId).
		UpdateEmailMessageRequest(*emailUpdate).
		Execute(); err != nil {
		t.Fatalf("Update Email template failed: %v (http=%+v)", err, httpResp)
	}

	// --- Send & resend PINs --------------------------------------------------
	smsStart := sms.NewStartAuthenticationRequest(appId, smsMsgId, "<DESTINATION>")
	smsSend, httpResp, err := client.
		TfaAPI.
		SendTfaPinCodeOverSms(auth).
		StartAuthenticationRequest(*smsStart).
		Execute()
	if err != nil {
		t.Fatalf("Send SMS PIN failed: %v (http=%+v)", err, httpResp)
	}
	if _, httpResp, err := client.
		TfaAPI.
		ResendTfaPinCodeOverSms(auth, smsSend.GetPinId()).
		ResendPinRequest(*sms.NewResendPinRequest()).
		Execute(); err != nil {
		t.Fatalf("Resend SMS PIN failed: %v (http=%+v)", err, httpResp)
	}

	voiceStart := sms.NewStartAuthenticationRequest(appId, voiceMsgId, "<DESTINATION>")
	voiceSend, httpResp, err := client.
		TfaAPI.
		SendTfaPinCodeOverVoice(auth).
		StartAuthenticationRequest(*voiceStart).
		Execute()
	if err != nil {
		t.Fatalf("Send Voice PIN failed: %v (http=%+v)", err, httpResp)
	}
	if _, httpResp, err := client.
		TfaAPI.
		ResendTfaPinCodeOverVoice(auth, voiceSend.GetPinId()).
		ResendPinRequest(*sms.NewResendPinRequest()).
		Execute(); err != nil {
		t.Fatalf("Resend Voice PIN failed: %v (http=%+v)", err, httpResp)
	}

	emailStart := sms.NewStartEmailAuthenticationRequest(appId, emailMsgId, "<DESTINATION_EMAIL>")
	emailSend, httpResp, err := client.
		TfaAPI.
		Send2faPinCodeOverEmail(auth).
		StartEmailAuthenticationRequest(*emailStart).
		Execute()
	if err != nil {
		t.Fatalf("Send Email PIN failed: %v (http=%+v)", err, httpResp)
	}
	if _, httpResp, err := client.
		TfaAPI.
		Resend2faPinCodeOverEmail(auth, emailSend.GetPinId()).
		ResendPinRequestViaEmail(*sms.NewResendPinRequestViaEmail()).
		Execute(); err != nil {
		t.Fatalf("Resend Email PIN failed: %v (http=%+v)", err, httpResp)
	}

	// Optional: verification status check
	if _, httpResp, err := client.
		TfaAPI.
		GetTfaVerificationStatus(auth, appId).
		Msisdn("<PHONE_NUMBER>").
		Execute(); err != nil {
		t.Fatalf("Verification status failed: %v (http=%+v)", err, httpResp)
	}

	// Minimal assertions to ensure IDs exist
	if appId == "" || smsMsgId == "" || voiceMsgId == "" || emailMsgId == "" || smsSend.GetPinId() == "" {
		t.Fatalf("Expected non-empty IDs in flow (app=%s smsMsg=%s voiceMsg=%s emailMsg=%s pin=%s)", appId, smsMsgId, voiceMsgId, emailMsgId, smsSend.GetPinId())
	}

	// Log key artifacts for manual inspection
	t.Logf("AppID=%s SMSMsgID=%s VoiceMsgID=%s EmailMsgID=%s SmsPinID=%s EmailPinID=%s VoicePinID=%s",
		appId, smsMsgId, voiceMsgId, emailMsgId, smsSend.GetPinId(), emailSend.GetPinId(), voiceSend.GetPinId())
	t.Logf("Total templates listed: %d, total apps listed: %d", len(msgs), len(apps))
}
