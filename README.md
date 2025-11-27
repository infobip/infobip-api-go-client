# Infobip API Go Client

<img src="https://cdn-web.infobip.com/uploads/2023/01/Infobip-logo.svg" height="93px" alt="Infobip" />

[![Go Reference](https://pkg.go.dev/badge/github.com/infobip/infobip-api-go-client.svg)](https://pkg.go.dev/github.com/infobip/infobip-api-go-client)
[![MIT License](https://badgen.net/github/license/infobip/infobip-api-go-client)](https://opensource.org/licenses/MIT)

This is a Go module for Infobip API and you can use it as a dependency when you want to consume [Infobip APIs][apidocs] in your application.
To use this, you'll need an Infobip account. You can create a free trial account [here][signup].

`infobip-api-go-client` is built on top of [OpenAPI Specification](https://swagger.io/specification/), powered by [OpenAPI Generator](https://openapi-generator.tech/).

#### Table of contents:
* [Documentation](#documentation)
* [General Info](#general-info)
* [Installation](#installation)
* [Quickstart](#quickstart)
* [Ask for help](#ask-for-help)

## Documentation

Detailed documentation about Infobip API can be found [here][apidocs].
The current version of this library includes this subset of Infobip products:
* [SMS](https://www.infobip.com/docs/api/channels/sms)
* [Messages API](https://www.infobip.com/docs/api/platform/messages-api)
* [Email](https://www.infobip.com/docs/api/channels/email)
* [Voice](https://www.infobip.com/docs/api/channels/voice)
* [Moments](https://www.infobip.com/docs/api/customer-engagement/moments)

## General Info
For `infobip-api-go-client` versioning we use [Semantic Versioning][semver] scheme.

Published under [MIT License][license].

Minimum `Go` version supported by this library is `1.18`.

## Installation
Pull the library by using the following command:
```shell
go get github.com/infobip/infobip-api-go-client/v3
```

To use our library you have to import it in your project files:
```shell
import "github.com/infobip/infobip-api-go-client/v3"
```

Afterwards, to update your `go.mod` and `go.sum` files, simply run the following:
```shell
go mod tidy
```

## Quickstart

#### Initialize the Client

In order to initialize the client first you need to set some basics like configuration and authentication.

We support multiple authentication methods, e.g. you can use [API Key](https://www.infobip.com/docs/essentials/api-authentication#api-key-header) or any of the other supported approaches.
Once you have an [Infobip account](https://www.infobip.com/signup), you can manage your API keys through the Infobip [API key management](https://portal.infobip.com/settings/accounts/api-keys) page.

To see your base URL, log in to the [Infobip API Resource][apidocs] hub with your Infobip credentials or visit your [Infobip account](https://portal.infobip.com/homepage/).

Let's first set the configuration field.
```go
configuration := infobip.NewConfiguration()
configuration.Host = "<YOUR_BASE_URL>"
```

Now you can initialize the API client.
```go
infobipClient := api.NewAPIClient(configuration)
```

For details, check the [client](https://github.com/infobip/infobip-api-go-client/blob/master/v3/pkg/infobip/client.go) source code.

#### Authentication
After that is done, we should set the authentication method.

```go
auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<YOUR_API_PREFIX>"},
		},
	)
```

To understand how to generate above mentioned tokens, check [this](https://www.infobip.com/docs/essentials/api-authentication) page.

That is it, now you are set to use the API.

#### Send an SMS
Here's a basic example of sending the SMS message.

```go
    import (
    	"context"
    	"fmt"

    	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
    	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
    	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
    )

    ...

    destinations := []sms.Destination{
		{To: "421907372599"},
	}

	content := sms.MessageContent{
		TextMessageContent: sms.NewTextMessageContent("Congratulations on sending your first message with GO library."),
	}

	givenMessage := sms.NewMessage(destinations, content)
	givenMessage.SetSender("421902028966")

	request := sms.NewRequestEnvelope([]sms.Message{
		*givenMessage,
	})

	// Send the SMS message
	apiResponse, httpResponse, err := infobipClient.
		SmsAPI.
		SendSmsMessages(auth).
		RequestEnvelope(*request).
		Execute()
```
```
    // If, for any reason, you don't want to use some of the response fields, feel free to replace them with an underscore wildcard, eg.
    // apiResponse, _, _ := ...
```

In order to be able to access our generated `Exception` models, you will have to do a little bit of type casting.
Methods provided within `ServiceException` object are `GetMessageId()` and `GetText()` referring to the nature of the exception.
Withing the `http.Response` response field (in example labeled as `httpResponse`) you can find pretty much everything that is referring to the HTTP status, e.g. `Status`, `StatusCode`, `Body`, `Header`, etc.

```go
if err != nil {
    apiErr, isApiErr := err.(*api.GenericOpenAPIError)
    if isApiErr {
        ibErr, isIbErr := apiErr.Model().(sms.ApiException)
        if isIbErr {
            errMessageId := ibErr.RequestError.ServiceException.GetMessageId()
            errText := ibErr.RequestError.ServiceException.GetText()
            // HANDLE THE IB ERROR EXCEPTION
        } else {
            // HANDLE THE API ERROR EXCEPTION
        }
    } else {
        // HANDLE THE EXCEPTION
    }
    return
}
```

Additionally, you can retrieve a the `bulkId` and `messageId` from the successful response (`SmsResponse` object) to use for fetching a delivery report for a given message or a bulk.
Bulk ID is received only when you send a message to more than one destination address or send multiple messages in a single request.

```go
bulkId := *apiResponse.BulkId
messageId := *apiResponse.Messages[0].MessageId
```

#### Receive sent SMS report

For each SMS that you send out, we can send you a message delivery report in real time. All you need to do is specify your endpoint when sending SMS in `notifyUrl` field of `SmsTextualMessage`, or subscribe for reports by contacting our support team.
e.g. `https://{yourDomain}/delivery-reports`

Example of webhook implementation:

```go
func deliveryReports(w http.ResponseWriter, req *http.Request) {
    reqBody, _ := io.ReadAll(req.Body)
    var report sms.DeliveryResult
    err := json.Unmarshal(reqBody, &report)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if report.HasResults() {
        results := report.GetResults()
        for _, result := range results {
            status := result.GetStatus()
            fmt.Println(fmt.Sprintf("%s - %s", result.GetMessageId(), status.GetName()))
        }
    }
}

func main() {
    http.HandleFunc("/delivery-reports", deliveryReports)
}
```

If you prefer to use your own serializer, please pay attention to the supported [date format](https://www.infobip.com/docs/essentials/integration-best-practices#date-formats).
In this library, we are wrapping around the `time.Time` model with our own `infobip.Time` in order to enforce the time format to be serialized properly.

#### Unicode & SMS preview
Infobip API supports Unicode characters and automatically detects encoding. Unicode and non-standard GSM characters use additional space, avoid unpleasant surprises and check how different message configurations will affect your message text, number of characters and message parts.

```go
text := "Let's see how many characters will remain unused in this message."
request := sms.NewPreviewRequest(text)

apiResponse, httpResponse, err := infobipClient.
    SmsAPI.
    PreviewSmsMessage(auth).
    PreviewRequest(*request).
    Execute()
```

#### Receive incoming SMS
If you want to receive SMS messages from your subscribers we can have them delivered to you in real time. When you buy and configure a number capable of receiving SMS, specify your endpoint as explained [here](https://www.infobip.com/docs/api#channels/sms/receive-inbound-sms-messages).
e.g. `https://{yourDomain}/incoming-sms`.

Example of webhook implementation:

```go
func incomingSms(w http.ResponseWriter, req *http.Request) {
    reqBody, _ := io.ReadAll(req.Body)
    var result sms.InboundMessageResult
    err := json.Unmarshal(reqBody, &result)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if result.HasResults() {
        msgs := result.GetResults()
        for _, msg := range msgs {
            fmt.Println(fmt.Sprintf("%s - %s", msg.GetFrom(), msg.GetCleanText()))
        }
    }
}

func main() {
    http.HandleFunc("/incoming-sms", incomingSms)
}
```

#### Two-Factor Authentication (2FA)
For the 2FA quick start guide please check [these examples](two-factor-authentication.md).

#### Messages API
For the Messages API quick start guide, view [these examples](messages-api.md).

#### Email
For the Email quick start guide, view [these examples](email.md).

#### Moments
For the Moments quick start guide, view [these examples](moments.md).

## Versioning

This project follows a pragmatic Semantic Versioning approach.  
For full details on how versions are managed, please see our [Versioning guide][versioning].

## Ask for help

Feel free to open an issue on the repository if you see any problem or want to request a feature. As per pull requests, for details check the `CONTRIBUTING` [file][contributing] related to it - in short, we will not merge any pull requests, this code is auto generated.

If it is, however, something that requires our immediate attention, feel free to contact us @ [support@infobip.com](mailto:support@infobip.com).

[apidocs]: https://www.infobip.com/docs/api
[signup]: https://www.infobip.com/signup
[semver]: https://semver.org
[license]: LICENSE
[contributing]: CONTRIBUTING.md
[versioning]: VERSIONING.md