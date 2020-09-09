package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/infobip/infobip-api-go-client/pkg"
	"github.com/infobip/infobip-api-go-client/pkg/sms"
)

func main() {
	username := flag.String("username", "username", "an Infobip username used to authenticate API requests")
	password := flag.String("password", "password", "an Infobip password used to authenticate API requests")
	phone := flag.String("phone", "", "a GSM formatted phone number used as a destination of SMS messages")
	flag.Parse()

	client := sms.Client{
		BaseURL:    "https://api.infobip.com", //optional parameter
		Authorizer: infobip.NewBasicCredentials(*username, *password),
		HTTPClient: http.DefaultClient, //optional
	}

	sendSms(client, newMultiMessage(*phone))
	sendSms(client, newMessageWithScheduling(*phone))
	sendSms(client, newMessageWithURLTracking(*phone))
	sendSms(client, newMessageWithNotifications(*phone))
	sendSms(client, newMessageWithIndiaDltParameters(*phone))

	response := sendSms(client, newSimpleMessage(*phone))

	time.Sleep(time.Minute * 5)

	if response != nil {
		messageID := response.Messages[0].MessageID
		query := sms.GetSentSMSDeliveryReportsQuery{
			MessageID: messageID,
			Limit:     15,
		}
		retrieveStatus(client, query)
	}

}

func sendSms(client sms.Client, req sms.SMSAdvancedTextualRequest) *sms.SMSResponse {
	res, err := client.SendMultipleTextualSMSAdvanced(req)

	if err != nil {
		fmt.Println("SMS sending error:")
		fmt.Println(err)
		return nil
	}

	fmt.Println("SMS sending success:")
	fmt.Printf("%+v\n", *res)
	return res
}

func retrieveStatus(client sms.Client, query sms.GetSentSMSDeliveryReportsQuery) *sms.SMSReportResponse {
	res, err := client.GetSentSMSDeliveryReports(query)

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
						sms.DD_MONDAY,
						sms.DD_TUESDAY,
						sms.DD_WEDNESDAY,
						sms.DD_THURSDAY,
						sms.DD_FRIDAY,
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

func newMessageWithIndiaDltParameters(destination string) sms.SMSAdvancedTextualRequest {
	return sms.SMSAdvancedTextualRequest{
		Messages: []sms.Message{
			sms.Message{
				Destinations: []sms.Destination{sms.Destination{To: destination}},
				From:         "InfoSMS",
				Text:         "SMS with India DLT parameters",
				Regional: &sms.RegionalOptions {
					IndiaDLT: &sms.IndiaDLTOptions {
						ContentTemplateID: "content template id",
						PrincipalEntityID: "principal entity id",
					},
				},
			},
		},
	}
}
