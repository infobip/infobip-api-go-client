## Two-Factor Authentication (2FA)

#### Application setup

The application represents your service. It’s a good practice to have separate applications for separate services.

```go
    request := sms.NewApplicationRequest("2FA application")

	apiResponse, httpResponse, err := infobipClient.
        TfaAPI.
		CreateTfaApplication(auth).
		ApplicationRequest(*request).
		Execute()

    appId := apiResponse.ApplicationId
```

#### Message template setup

Message template is the message body with the PIN placeholder that is sent to end users.
```go
	pinLength := int32(4)
	
    request := sms.NewCreateMessageRequest("Your pin is {{pin}}", sms.PINTYPE_ALPHANUMERIC)
    request.PinLength = &pinLength
    
    apiResponse, httpResponse, err := infobipClient.
        TfaAPI.
        CreateTfaMessageTemplate(auth, appId).
        CreateMessageRequest(*request).
        Execute()
	
	msgId := apiResponse.MessageId
```

#### Send 2FA code via SMS

After setting up the application and message template, you can start generating and sending PIN codes via SMS to the provided destination address.

```go
    from := "InfoSMS"
    to := "41793026727"
    
    request := sms.NewStartAuthenticationRequest(*appId, *msgId, to)
    request.From = &from

	apiResponse, httpResponse, err := infobipClient.
        TfaAPI.
		SendTfaPinCodeOverSms(auth).
		StartAuthenticationRequest(*request).
		Execute()
		
    pinId := apiResponse.PinId
```

#### Verify phone number

Verify a phone number to confirm successful 2FA authentication.

```go
	request := sms.NewVerifyPinRequest("1598")

	apiResponse, httpResponse, err := infobipClient.
        TfaAPI.
		VerifyTfaPhoneNumber(auth, *pinId).
		VerifyPinRequest(*request).
		Execute()
```
