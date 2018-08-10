# Infobip API Go client

## Prerequisites

[Go](https://golang.org/)

## Installation

[`go get github.com/infobip/infobip-api-go-client`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies)

## Usage

```go
package main

import (
    "fmt"
	"net/http"

	"github.com/infobip/infobip-api-go-client/pkg"
    "github.com/infobip/infobip-api-go-client/pkg/sms"
)

func main() {
	client := sms.Client{
		BaseURL:    "https://api.infobip.com",
		Authorizer: infobip.NewBasicCredentials("username", "password"),
		HTTPCLient: http.DefaultClient,
	}

	sendRes, err := client.SendMultipleTextualSmsAdvanced(sms.SMSAdvancedTextualRequest{
		Messages: []sms.Message{
			sms.Message{
				Destinations: []sms.Destination{sms.Destination{To: "phoneNumber"}},
				Text:         "SMS from Infobip API go client.",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Minute * 5)

	statusRes, err := client.GetSentSmsDeliveryReports(sms.GetSentSmsDeliveryReportsQuery{
		MessageID: sendRes.Messages[0].MessageID,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Staus of the send SMS is %+v", statusRes.Results[0].Status)
}
```

Find more examples in [`./example`](./example).