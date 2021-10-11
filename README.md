# Infobip API Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/infobip/infobip-api-go-client.svg)](https://pkg.go.dev/github.com/infobip/infobip-api-go-client)
[![MIT License](https://badgen.net/github/license/infobip/infobip-api-go-client)](https://opensource.org/licenses/MIT)

This is a Go module for Infobip API and you can use it as a dependency to add [Infobip APIs][apidocs] to your application.
To use this, you'll need an Infobip account. If not already having one, you can create a [free trial][freetrial] account [here][signup].

Built on top of [OpenAPI Specification](https://swagger.io/specification/), powered by [OpenAPI Generator](https://openapi-generator.tech/).

<img src="https://udesigncss.com/wp-content/uploads/2020/01/Infobip-logo-transparent.png" height="124px" alt="Infobip" />

#### Table of contents:
* [Documentation](#documentation)
* [General Info](#general-info)
* [Installation](#installation)
* [Quickstart](#quickstart)
* [Ask for help](#ask-for-help)

## Documentation

Infobip API Documentation can be found [here][apidocs].

## General Info
For `infobip-api-go-client` versioning we use [Semantic Versioning][semver] scheme.

Published under [MIT License][license].

Minimum Go version supported by this library is 1.13.

## Installation
Pull the library by using the following command:
```shell
go get github.com/infobip/infobip-api-go-client/v2
```

To use our library you have to import it in your project files:
```shell
import "github.com/infobip/infobip-api-go-client/v2"
```

Afterwards, to update your `go.mod` and `go.sum` files, simply run the following:
```shell
go mod tidy
```

## Quickstart

#### Initialize the Client

In order to initialize the client first you need to set some basics like configuration and authentication. 

We support multiple authentication methods, e.g. you can use [API Key](https://www.infobip.com/docs/essentials/api-authentication#api-key-header) or any of the other supported approaches.

To see your base URL, log in to the [Infobip API Resource][apidocs] hub with your Infobip credentials.

Let's first set the configuration field.
```go
	configuration := infobip.NewConfiguration()
	configuration.Host = "<put your base URL here>"
```

Now you can initialize the API client.
```go
	infobipClient := infobip.NewAPIClient(configuration)
```

#### Authentication
After that is done, we should set the authentication method.
There are several ways how that can be done.

If you prefer to use Basic Auth:
```go
	auth := context.WithValue(context.Background(), infobip.ContextBasicAuth, infobip.BasicAuth{
		UserName: "<put your username here>",
		Password: "<put your password here>",
	})
```

If you prefer to use Api Key:
```go
	auth := context.WithValue(context.Background(), infobip.ContextAPIKey, "<put your Infobip Api key here>")
```

If you prefer to use IBSSO token:
```go
	auth := context.WithValue(context.Background(), infobip.ContextIBSSO, "<put your IBSSO token here>")
```

Or you might want to use Infobip generated OAuth2 token:
```go
	auth := context.WithValue(context.Background(), infobip.ContextAccessToken, "<put your Infobip generated OAuth2 token here>")
```
To understand how to generate above mentioned tokens, check [this](https://www.infobip.com/docs/essentials/api-authentication) page.

In order to use the `auth` you need to provide it as a param to the endpoint you are triggering, eg.
```go
	apiResponse, httpResponse, err := infobipClient.
            SendSmsApi.
            GetOutboundSmsMessageLogs(auth).
            From("<put your From field filter here>").
            Execute()
```

That is it, now you are set to use the API.

#### Send an SMS
Simple example for sending an SMS message.

```go
	request := infobip.NewSmsAdvancedTextualRequest()

	destination := infobip.NewSmsDestination("41793026727")

	from := "InfoSMS"
	text := "This is a dummy SMS message"
	message := infobip.NewSmsTextualMessage()
	message.From = &from
	message.Destinations = &[]infobip.SmsDestination{*destination}
	message.Text = &text
```

Send the message and inspect the `err` field for more information in case of failure.
You can get the HTTP status code from `httpResponse` property, and more details about error from `ErrorContent` property.

```go
	apiResponse, httpResponse, err := infobipClient.
		SendSmsApi.
		SendSmsMessage(auth).
		SmsAdvancedTextualRequest(*request).
		Execute()

    // If, for any reason, you don't want to use some of the response fields, feel free to replace them with an underscore wildcard, eg.
    //  apiResponse, _, _ := ...
```

In order to be able to access our generated `Exception` models you will have to do a little bit of type casting.
Methods provided within `ServiceException` object are `GetMessageId()` and `GetText()` referring to the nature of the exception. 
Withing the `_nethttp.Response` response field (in example labeled as `httpResponse`) you can find pretty much everything that is referring to the HTTP status, e.g. `Status`, `StatusCode`, `Body`, `Header`, etc.

```go
	if err != nil {
        apiErr, isApiErr := err.(infobip.GenericOpenAPIError)
        if isApiErr {
            ibErr, isIbErr := apiErr.Model().(infobip.SmsApiException)
            if isIbErr {
                ibErr.RequestError.ServiceException.GetMessageId()
                ibErr.RequestError.ServiceException.GetText()
            }
        }
        return
    }
```

Additionally, from the successful response (`SmsResponse` object) you can pull out the `bulkId` and `messageId`(s) and use them to fetch a delivery report for given message or bulk.
Bulk ID will be received only when you send a message to more than one destination address or multiple messages in a single request.

```go
    bulkId := *apiResponse.BulkId
    messageId := *apiResponse.Messages[0].MessageId
```

#### Receive sent SMS report
For each SMS that you send out, we can send you a message delivery report in real time. All you need to do is specify your endpoint when sending SMS in `notifyUrl` field of `SmsTextualMessage`, or subscribe for reports by contacting our support team.
e.g. `https://{yourDomain}/delivery-reports`

Example of webhook implementation:

```go
    func DeliveryReports(w http.ResponseWriter, req *http.Request) {
        reqBody, _ := ioutil.ReadAll(req.Body)
        var report SmsDeliveryResult
        json.Unmarshal(reqBody, &report)
        fmt.Fprintf("Delivery report: %s", report)
    }

    func main() {
        http.HandleFunc("/delivery-reports", deliveryReports)
    }
```
If you prefer to use your own serializer, please pay attention to the supported [date format](https://www.infobip.com/docs/essentials/integration-best-practices#date-formats).
In this library, we are wrapping around the `time.Time` model with our own `infobip.Time` in order to enforce the time format to be serialized properly.

#### Fetching delivery reports
If you are for any reason unable to receive real time delivery reports on your endpoint, you can use `messageId` or `bulkId` to fetch them.
Each request will return a batch of delivery reports - only once.

```go
    limit := int32(10)
    bulkId := "bulk-1234"
    messageId := "msg-123"

    apiResponse, httpResponse, err := infobipClient.
		SendSmsApi.
		GetOutboundSmsMessageDeliveryReports(auth).
		BulkId(bulkId).
		MessageId(messageId).
        Limit(limit).
		Execute()
```

#### Unicode & SMS preview
Infobip API supports Unicode characters and automatically detects encoding. Unicode and non-standard GSM characters use additional space, avoid unpleasant surprises and check how different message configurations will affect your message text, number of characters and message parts.

```go
    text := "Let's see how many characters will remain unused in this message."
    request := *infobip.NewSmsPreviewRequest(text)

	apiResponse, httpResponse, err := infobipClient.
		SendSmsApi.
		PreviewSmsMessage(auth).
		SmsPreviewRequest(request).
		Execute()
```

#### Receive incoming SMS
If you want to receive SMS messages from your subscribers we can have them delivered to you in real time. When you buy and configure a number capable of receiving SMS, specify your endpoint as explained [here](https://www.infobip.com/docs/api#channels/sms/receive-inbound-sms-messages).
e.g. `https://{yourDomain}/incoming-sms`.

Example of webhook implementation:

```go
    func IncomingSms(w http.ResponseWriter, req *http.Request) {
        reqBody, _ := ioutil.ReadAll(req.Body)
        var msg SmsInboundMessage
        json.Unmarshal(reqBody, &msg)
        fmt.Fprintf("Received message: %s", msg)
    }
    func main() {
        http.HandleFunc("/incoming-sms", incomingSms)
    }
```
#### Two-Factor Authentication (2FA)
For 2FA quick start guide please check [these examples](two-factor-authentication.md).

## Ask for help

Feel free to open issues on the repository for any issue or feature request. As per pull requests, for details check the `CONTRIBUTING` [file][contributing] related to it - in short, we will not merge any pull requests, this code is auto-generated.

If it is, however, something that requires our imminent attention feel free to contact us @ [support@infobip.com](mailto:support@infobip.com).

[apidocs]: https://www.infobip.com/docs/api
[freetrial]: https://www.infobip.com/docs/freetrial
[signup]: https://www.infobip.com/signup
[semver]: https://semver.org
[license]: LICENSE
[contributing]: CONTRIBUTING.md
