package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	SmsSendMessages        string = "/sms/3/messages"
	SmsLogsEndpoint        string = "/sms/3/logs"
	SmsReportsEndpoint     string = "/sms/3/reports"
	SmsSendPreviewEndpoint string = "/sms/1/preview"
	SmsInboxReports        string = "/sms/1/inbox/reports"
	SmsBulks               string = "/sms/1/bulks"
	SmsBulksStatus         string = "/sms/1/bulks/status"
)

const (
	PendingStatusGroupId     int                      = 1
	PendingStatusGroupName   sms.MessageGeneralStatus = sms.MESSAGEGENERALSTATUS_PENDING
	PendingStatusId          int                      = 26
	PendingStatusName        string                   = "PENDING_ACCEPTED"
	PendingStatusDescription string                   = "Message sent to next instance"
)

const (
	StatusGroupId     int                      = 3
	StatusGroupName   sms.MessageGeneralStatus = sms.MESSAGEGENERALSTATUS_DELIVERED
	StatusId          int                      = 5
	StatusName        string                   = "DELIVERED_TO_HANDSET"
	StatusDescription string                   = "Message delivered to handset"
)

const (
	NoErrorStatusGroupId     int                   = 0
	NoErrorStatusGroupName   sms.MessageErrorGroup = sms.MESSAGEERRORGROUP_OK
	NoErrorStatusId          int                   = 0
	NoErrorStatusName        string                = "NO_ERROR"
	NoErrorStatusDescription string                = "No error"
	NoErrorPermanent         bool                  = true
)

func TestSendSms(t *testing.T) {
	givenBulkId := "2034072219640523072"
	givenDestination := "41793026727"
	givenMessageId := "2250be2d4219-3af1-78856-aabe-1362af1edfd2"

	givenSender := "InfoSMS"
	givenText := "This is a sample message"
	givenMessageCount := 1

	// Expected response and request bodies
	givenResponse := fmt.Sprintf(`
        {
            "bulkId": "%s",
            "messages": [
                {
                    "messageId": "%s",
                    "status": {
                        "groupId": %d,
                        "groupName": "%s",
                        "id": %d,
                        "name": "%s",
                        "description": "%s"
                    },
                    "destination": "%s",
                    "details": {
                        "messageCount": %d
                    }
                }
            ]
        }`,
		givenBulkId,
		givenMessageId,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		givenDestination,
		givenMessageCount,
	)

	givenRequest := fmt.Sprintf(`
        {
            "messages": [
                {
                    "sender": "%s",
                    "destinations": [
                        {
                            "to": "%s"
                        }
                    ],
                    "content": {
                        "text": "%s"
                    }
                }
            ]
        }`,
		givenSender,
		givenDestination,
		givenText,
	)

	// Prepare the message
	destinations := []sms.Destination{
		{To: givenDestination},
	}

	content := sms.MessageContent{
		TextMessageContent: sms.NewTextMessageContent(givenText),
	}

	givenMessage := sms.NewMessage(destinations, content)
	givenMessage.SetSender(givenSender)

	request := sms.NewRequestEnvelope([]sms.Message{
		*givenMessage,
	})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SmsSendMessages, givenResponse, 200)

	// Execute the request
	responseBody, r, _ := infobipClient.
		SmsAPI.
		SendSmsMessages(context.Background()).
		RequestEnvelope(*request).Execute()

	// Ensure r and responseBody are not nil before dereferencing
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.True(t, len(responseBody.GetMessages()) == 1)

	msg := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	smsStatus := msg.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), smsStatus.GetGroupId())
	assert.Equal(t, PendingStatusGroupName, smsStatus.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), smsStatus.GetId())
	assert.Equal(t, PendingStatusName, smsStatus.GetName())
	assert.Equal(t, PendingStatusDescription, smsStatus.GetDescription())
}

func TestShouldSendFlashSms(t *testing.T) {
	givenBulkId := "2034072219640523072"
	givenDestination := "41793026727"
	givenMessageId := "2250be2d4219-3af1-78856-aabe-1362af1edfd2"

	givenSender := "InfoSMS"
	givenText := "Toto, I've got a feeling we're not in Kansas anymore."
	isFlash := true

	// Expected response and request bodies
	givenResponse := fmt.Sprintf(`
        {
            "bulkId": "%s",
            "messages": [
                {
                    "messageId": "%s",
                    "status": {
                        "groupId": %d,
                        "groupName": "%s",
                        "id": %d,
                        "name": "%s",
                        "description": "%s"
                    },
                    "destination": "%s",
                    "details": {
                        "messageCount": %d
                    }
                }
            ]
        }`,
		givenBulkId,
		givenMessageId,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		givenDestination,
		1, // Message count is 1
	)

	givenRequest := fmt.Sprintf(`
        {
            "messages": [
                {
                    "sender": "%s",
                    "destinations": [
                        {
                            "to": "%s"
                        }
                    ],
                    "content": {
                        "text": "%s"
                    },
                    "options": {
                        "flash": %t
                    }
                }
            ]
        }`,
		givenSender,
		givenDestination,
		givenText,
		isFlash,
	)

	// Prepare the message
	destinations := []sms.Destination{
		{To: givenDestination},
	}

	content := sms.MessageContent{
		TextMessageContent: sms.NewTextMessageContent(givenText),
	}

	options := sms.MessageOptions{Flash: &isFlash}

	givenMessage := sms.NewMessage(destinations, content)
	givenMessage.SetSender(givenSender)
	givenMessage.SetOptions(options)

	request := sms.NewRequestEnvelope([]sms.Message{
		*givenMessage,
	})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SmsSendMessages, givenResponse, 200)

	// Execute the request
	responseBody, r, _ := infobipClient.
		SmsAPI.
		SendSmsMessages(context.Background()).
		RequestEnvelope(*request).Execute()

	// Ensure r and responseBody are not nil before dereferencing
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.True(t, len(responseBody.GetMessages()) == 1)

	msg := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	smsStatus := msg.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), smsStatus.GetGroupId())
	assert.Equal(t, PendingStatusGroupName, smsStatus.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), smsStatus.GetId())
	assert.Equal(t, PendingStatusName, smsStatus.GetName())
	assert.Equal(t, PendingStatusDescription, smsStatus.GetDescription())
}

func TestSendFullyFeaturedSmsMessage(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenFirstDestinationMessageId := "abc123"
	givenGroupId := 1
	givenGroupName := "PENDING"
	givenStatusId := 26
	givenStatusName := "PENDING_ACCEPTED"
	givenStatusDescription := "Message sent to next instance"
	givenFirstDestination := "385991112222"
	givenDetailsMessageCount := 1
	givenSecondDestinationMessageId := "xyz100"
	givenSecondDestination := "385911231234"
	givenCallbackData := "transactionId=100"
	deliveryTimeWindowDayMonday := sms.DELIVERYDAY_MONDAY
	deliveryTimeWindowDaySunday := sms.DELIVERYDAY_SUNDAY
	deliveryTimeWindowFromHours := 12
	deliveryTimeWindowFromMinutes := 10
	deliveryTimeWindowToHours := 20
	deliveryTimeWindowToMinutes := 15
	givenFrom := "Info"
	givenIntermediateReport := true
	givenLanguageCode := sms.LANGUAGECODE_AUTODETECT
	smsLanguageCode := sms.Language{
		LanguageCode: &givenLanguageCode,
	}
	givenTransliterationCode := sms.TRANSLITERATIONCODE_CENTRAL_EUROPEAN
	givenNotifyContentType := "application/json"
	givenNotifyUrl := "https://www.example.com/sms/advanced"
	givenSendAtString := "2021-08-24T15:00:00.000+0000"
	givenText := "Laku noć"
	givenTransliteration := "CENTRAL_EUROPEAN"
	givenValidityPeriod := 720
	givenValidityPeriodTimeUnit := sms.VALIDITYPERIODTIMEUNIT_HOURS
	givenUrlOptionsShortenUrl := true
	givenUrlOptionsTrackClicks := false
	givenUrlOptionsTrackingUrl := "https://ib.com"
	givenUrlOptionsRemoveProtocol := true
	givenUrlOptionsCustomDomain := "example.com"
	givenTrackingType := "MY_CAMPAIGN"
	givenIncludeSmsCountInResponse := true
	givenUseConversionTracking := true
	givenCampaingReferenceId := "summersale"

	// Expected request body
	expectedRequest := fmt.Sprintf(`{
		"messages": [
			{
				"sender": "%s",
				"destinations": [
					{
						"to": "%s",
						"messageId": "%s"
					},
					{
						"to": "%s"
					}
				],
				"content": {
					"text": "%s",
					"transliteration": "%s",
					"language": {
						"languageCode": "%s"
					}
				},
				"options": {
					"validityPeriod": {
						"amount": %d,
						"timeUnit": "%s"
					},
					"campaignReferenceId": "%s"
				},
				"webhooks": {
					"delivery": {
						"url": "%s",
						"intermediateReport": %t
					},
					"contentType": "%s",
					"callbackData": "%s"
				}
			},
			{
				"sender": "%s",
				"destinations": [
					{
						"to": "%s"
					}
				],
				"content": {
					"text": "%s"
				},
				"options": {
					"deliveryTimeWindow": {
						"days": [
							"%s",
							"%s"
						],
						"from": {
							"hour": %d,
							"minute": %d
						},
						"to": {
							"hour": %d,
							"minute": %d
						}
					}
				}
			}
		],
		"options": {
			"schedule": {
				"bulkId": "%s",
				"sendAt": "%s"
			},
			"tracking": {
				"shortenUrl": %t,
				"trackClicks": %t,
				"trackingUrl": "%s",
				"removeProtocol": %t,
				"customDomain": "%s"
			},
			"includeSmsCountInResponse": %t,
			"conversionTracking": {
				"useConversionTracking": %t,
				"conversionTrackingName": "%s"
			}
		}
	}`, givenFrom, givenFirstDestination, givenFirstDestinationMessageId, givenSecondDestination, givenText, givenTransliteration,
		givenLanguageCode, givenValidityPeriod, givenValidityPeriodTimeUnit, givenCampaingReferenceId, givenNotifyUrl, givenIntermediateReport,
		givenNotifyContentType, givenCallbackData, givenFrom, givenSecondDestination, givenText, deliveryTimeWindowDayMonday, deliveryTimeWindowDaySunday,
		deliveryTimeWindowFromHours, deliveryTimeWindowFromMinutes, deliveryTimeWindowToHours, deliveryTimeWindowToMinutes, givenBulkId, givenSendAtString,
		givenUrlOptionsShortenUrl, givenUrlOptionsTrackClicks, givenUrlOptionsTrackingUrl, givenUrlOptionsRemoveProtocol, givenUrlOptionsCustomDomain,
		givenIncludeSmsCountInResponse, givenUseConversionTracking, givenTrackingType)

	// Expected response
	givenResponse := fmt.Sprintf(`{
		"bulkId": "%s",
		"messages": [
			{
				"messageId": "%s",
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				},
				"destination": "%s",
				"details": {
					"messageCount": %d
				}
			},
			{
				"messageId": "%s",
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				},
				"destination": "%s",
				"details": {
					"messageCount": %d
				}
			}
		]
	}`, givenBulkId, givenFirstDestinationMessageId, givenGroupId, givenGroupName, givenStatusId, givenStatusName, givenStatusDescription,
		givenFirstDestination, givenDetailsMessageCount, givenSecondDestinationMessageId, givenGroupId, givenGroupName, givenStatusId,
		givenStatusName, givenStatusDescription, givenSecondDestination, givenDetailsMessageCount)

	// Prepare the messages
	firstDestinations := []sms.Destination{
		{To: givenFirstDestination, MessageId: &givenFirstDestinationMessageId},
		{To: givenSecondDestination},
	}
	secondDestinations := []sms.Destination{
		{To: givenSecondDestination},
	}

	firstContent := sms.TextMessageContent{
		Text:            givenText,
		Transliteration: &givenTransliterationCode,
		Language:        &smsLanguageCode,
	}

	secondContent := sms.TextMessageContent{
		Text: givenText,
	}

	firstOptions := sms.MessageOptions{
		ValidityPeriod: &sms.ValidityPeriod{
			Amount:   int32(givenValidityPeriod),
			TimeUnit: &givenValidityPeriodTimeUnit,
		},
		CampaignReferenceId: &givenCampaingReferenceId,
	}

	secondOptions := sms.MessageOptions{
		DeliveryTimeWindow: &sms.DeliveryTimeWindow{
			Days: []sms.DeliveryDay{deliveryTimeWindowDayMonday, deliveryTimeWindowDaySunday},
			From: &sms.DeliveryTime{Hour: int32(deliveryTimeWindowFromHours), Minute: int32(deliveryTimeWindowFromMinutes)},
			To:   &sms.DeliveryTime{Hour: int32(deliveryTimeWindowToHours), Minute: int32(deliveryTimeWindowToMinutes)},
		},
	}

	webhooks := sms.Webhooks{
		Delivery: &sms.MessageDeliveryReporting{
			Url:                &givenNotifyUrl,
			IntermediateReport: &givenIntermediateReport,
		},
		ContentType:  &givenNotifyContentType,
		CallbackData: &givenCallbackData,
	}

	firstContentLog := sms.MessageContent{
		TextMessageContent: &firstContent,
	}

	secondContentLog := sms.MessageContent{
		TextMessageContent: &secondContent,
	}

	firstMessage := sms.NewMessage(firstDestinations, firstContentLog)
	firstMessage.SetOptions(firstOptions)
	firstMessage.SetWebhooks(webhooks)
	firstMessage.SetSender(givenFrom)

	secondMessage := sms.NewMessage(secondDestinations, secondContentLog)
	secondMessage.SetOptions(secondOptions)
	secondMessage.SetSender(givenFrom)

	request := sms.NewRequestEnvelope([]sms.Message{*firstMessage, *secondMessage})

	givenSentAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAtString)
	ibTimeGivenSentAt := infobip.Time{
		T: givenSentAt,
	}

	requestOptions := sms.MessageRequestOptions{
		Schedule: &sms.RequestSchedulingSettings{
			BulkId: &givenBulkId,
			SendAt: &ibTimeGivenSentAt,
		},
		Tracking: &sms.UrlOptions{
			ShortenUrl:     &givenUrlOptionsShortenUrl,
			TrackClicks:    &givenUrlOptionsTrackClicks,
			TrackingUrl:    &givenUrlOptionsTrackingUrl,
			RemoveProtocol: &givenUrlOptionsRemoveProtocol,
			CustomDomain:   &givenUrlOptionsCustomDomain,
		},
		IncludeSmsCountInResponse: &givenIncludeSmsCountInResponse,
		ConversionTracking: &sms.Tracking{
			UseConversionTracking:  &givenUseConversionTracking,
			ConversionTrackingName: &givenTrackingType,
		},
	}

	request.SetOptions(requestOptions)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SmsSendMessages, givenResponse, 200)

	responseBody, r, _ := infobipClient.
		SmsAPI.
		SendSmsMessages(context.Background()).
		RequestEnvelope(*request).Execute()

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	assert.True(t, len(responseBody.GetMessages()) == 2)

	message1 := responseBody.GetMessages()[0]
	assert.Equal(t, givenFirstDestinationMessageId, message1.GetMessageId())
	assert.Equal(t, givenFirstDestination, message1.GetDestination())
	smsStatus1 := message1.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), smsStatus1.GetGroupId())
	assert.Equal(t, PendingStatusGroupName, smsStatus1.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), smsStatus1.GetId())
	assert.Equal(t, PendingStatusName, smsStatus1.GetName())
	assert.Equal(t, PendingStatusDescription, smsStatus1.GetDescription())

	message2 := responseBody.GetMessages()[1]
	assert.Equal(t, givenSecondDestinationMessageId, message2.GetMessageId())
	assert.Equal(t, givenSecondDestination, message2.GetDestination())
	smsStatus2 := message2.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), smsStatus2.GetGroupId())
	assert.Equal(t, PendingStatusGroupName, smsStatus2.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), smsStatus2.GetId())
	assert.Equal(t, PendingStatusName, smsStatus2.GetName())
	assert.Equal(t, PendingStatusDescription, smsStatus2.GetDescription())

	messageDetails1 := message2.GetDetails()
	messageDetails2 := message2.GetDetails()
	assert.Equal(t, int32(1), messageDetails1.GetMessageCount())
	assert.Equal(t, int32(1), messageDetails2.GetMessageCount())
}

func TestSendBinarySmsMessage(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenDestination := "41793026727"
	givenMessageId := "MESSAGE-ID-123-xyz"
	givenFrom := "InfoSMS"
	givenHex := "54 65 73 74 20 6d 65 73 73 61 67 65 2e"
	givenDataCoding := int32(0)
	givenEsmClass := int32(0)
	givenNotifyUrl := "https://www.example.com/sms/advanced"
	givenNotifyContentType := "application/json"
	givenCallbackData := "DLR callback data"
	givenValidityPeriod := 720
	givenValidityPeriodTimeUnit := sms.VALIDITYPERIODTIMEUNIT_HOURS
	givenSendAtString := "2019-11-09T16:00:00.000+0000"
	givenIntermediateReport := true

	// Expected request body
	expectedRequest := fmt.Sprintf(`{
		"messages": [
			{
				"sender": "%s",
				"destinations": [
					{
						"to": "%s",
						"messageId": "%s"
					}
				],
				"content": {
					"hex": "%s",
					"dataCoding": %d,
					"esmClass": %d
				},
				"options": {
					"validityPeriod": {
						"amount": %d,
						"timeUnit": "%s"
					}
				},
				"webhooks": {
					"delivery": {
						"url": "%s",
						"intermediateReport": true
					},
					"contentType": "%s",
					"callbackData": "%s"
				}
			}
		],
		"options": {
			"schedule": {
				"bulkId": "%s",
				"sendAt": "%s"
			}
		}
	}`, givenFrom, givenDestination, givenMessageId, givenHex, givenDataCoding, givenEsmClass,
		givenValidityPeriod, givenValidityPeriodTimeUnit, givenNotifyUrl, givenNotifyContentType, givenCallbackData,
		givenBulkId, givenSendAtString)

	// Expected response
	expectedResponse := fmt.Sprintf(`{
		"bulkId": "%s",
		"messages": [
			{
				"messageId": "%s",
				"status": {
					"groupId": 1,
					"groupName": "PENDING",
					"id": 26,
					"name": "PENDING_ACCEPTED",
					"description": "Message sent to next instance"
				},
				"destination": "%s",
				"details": {
					"messageCount": 1
				}
			}
		]
	}`, givenBulkId, givenMessageId, givenDestination)

	// Prepare the message
	destinations := []sms.Destination{
		{To: givenDestination, MessageId: &givenMessageId},
	}

	content := sms.BinaryContent{
		Hex:        givenHex,
		DataCoding: &givenDataCoding,
		EsmClass:   &givenEsmClass,
	}

	contentLog := sms.MessageContent{
		BinaryContent: &content,
	}

	options := sms.MessageOptions{
		ValidityPeriod: &sms.ValidityPeriod{
			Amount:   int32(givenValidityPeriod),
			TimeUnit: &givenValidityPeriodTimeUnit,
		},
	}
	webhooks := sms.Webhooks{
		Delivery: &sms.MessageDeliveryReporting{
			Url:                &givenNotifyUrl,
			IntermediateReport: &givenIntermediateReport,
		},
		ContentType:  &givenNotifyContentType,
		CallbackData: &givenCallbackData,
	}

	message := sms.NewMessage(destinations, contentLog)
	message.SetSender(givenFrom)
	message.SetOptions(options)
	message.SetWebhooks(webhooks)

	request := sms.NewRequestEnvelope([]sms.Message{*message})

	givenSentAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAtString)
	ibTimeGivenSentAt := infobip.Time{
		T: givenSentAt,
	}

	requestOptions := sms.MessageRequestOptions{
		Schedule: &sms.RequestSchedulingSettings{
			BulkId: &givenBulkId,
			SendAt: &ibTimeGivenSentAt,
		},
	}

	request.SetOptions(requestOptions)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SmsSendMessages, expectedResponse, 200)

	responseBody, r, _ := infobipClient.
		SmsAPI.
		SendSmsMessages(context.Background()).
		RequestEnvelope(*request).Execute()

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	assert.True(t, len(responseBody.GetMessages()) == 1)

	message1 := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, message1.GetMessageId())
	assert.Equal(t, givenDestination, message1.GetDestination())
	smsStatus := message1.GetStatus()
	assert.Equal(t, int32(1), smsStatus.GetGroupId())
	assert.Equal(t, PendingStatusGroupName, smsStatus.GetGroupName())
	assert.Equal(t, int32(26), smsStatus.GetId())
	assert.Equal(t, "PENDING_ACCEPTED", smsStatus.GetName())
	assert.Equal(t, "Message sent to next instance", smsStatus.GetDescription())
}

func TestSendFlashBinarySms(t *testing.T) {
	// Given values
	givenBulkId := "2034072219640523072"
	givenTo := "41793026727"
	givenMessageId := "2250be2d4219-3af1-78856-aabe-1362af1edfd2"
	givenFrom := "InfoSMS"
	givenHex := "0048 0065 006c 006c 006f 0020 0077 006f 0072 006c 0064 002c 0020 039a 03b1 03bb 03b7 03bc 03ad 03c1 03b1 0020 03ba 03cc 03c3 03bc 03b5 002c 0020 30b3 30f3 30cb 30c1 30cf"
	givenDataCoding := int32(0)
	givenEsmClass := int32(0)
	isFlash := true

	//Expected request body
	expectedRequest := fmt.Sprintf(`{
	   "messages": [
	       {
	           "sender": "%s",
	           "destinations": [
	               {
	                   "to": "%s"
	               }
	           ],
	           "content": {
	               "hex": "%s",
	               "dataCoding": %d,
	               "esmClass": %d
	           },
	           "options": {
	               "flash": %t
	           }
	       }
	   ]
	}`, givenFrom, givenTo, givenHex, givenDataCoding, givenEsmClass, isFlash)

	// Expected response
	givenResponse := fmt.Sprintf(`
        {
            "bulkId": "%s",
            "messages": [
                {
                    "messageId": "%s",
                    "status": {
                        "groupId": %d,
                        "groupName": "%s",
                        "id": %d,
                        "name": "%s",
                        "description": "%s"
                    },
                    "destination": "%s",
                    "details": {
                        "messageCount": %d
                    }
                }
            ]
        }`,
		givenBulkId,
		givenMessageId,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		givenTo,
		1, // Message count is 1
	)

	// Prepare the message
	destination := sms.Destination{
		To: givenTo,
	}

	content := sms.BinaryContent{
		Hex:        givenHex,
		DataCoding: &givenDataCoding,
		EsmClass:   &givenEsmClass,
	}

	options := sms.MessageOptions{
		Flash: &isFlash,
	}

	destinations := []sms.Destination{
		destination,
	}

	message := sms.Message{
		Destinations: destinations,
		Sender:       &givenFrom,
		Content: sms.MessageContent{
			BinaryContent: &content,
		},
		Options: &options,
	}

	request := sms.RequestEnvelope{
		Messages: []sms.Message{message},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", SmsSendMessages, givenResponse, 200)

	responseBody, r, _ := infobipClient.
		SmsAPI.
		SendSmsMessages(context.Background()).
		RequestEnvelope(request).Execute()

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)
	// Assertions
	assert.NotNil(t, responseBody)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	assert.True(t, len(responseBody.GetMessages()) == 1)

	assert.True(t, len(responseBody.GetMessages()) == 1)

	message1 := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, message1.GetMessageId())
	assert.Equal(t, givenTo, message1.GetDestination())
	smsStatus := message1.GetStatus()
	assert.Equal(t, int32(1), smsStatus.GetGroupId())
	assert.Equal(t, PendingStatusGroupName, smsStatus.GetGroupName())
	assert.Equal(t, int32(26), smsStatus.GetId())
	assert.Equal(t, "PENDING_ACCEPTED", smsStatus.GetName())
	assert.Equal(t, "Message sent to next instance", smsStatus.GetDescription())
}

func TestSendSmsPreview(t *testing.T) {
	givenOriginalText := "Let's see how many characters will remain unused in this message."
	givenTextPreview := "Let's see how many characters will remain unused in this message."
	givenMessageCount := int32(1)
	givenCharactersRemaining := int32(95)

	givenResponse := fmt.Sprintf(`
		{
			"originalText": "%s",
			"previews": [{
				"textPreview": "%s",
				"messageCount": %d,
				"charactersRemaining": %d,
				"configuration": {
				}
			}]
       }
		`,
		givenOriginalText,
		givenTextPreview,
		givenMessageCount,
		givenCharactersRemaining,
	)

	givenRequest := fmt.Sprintf(`
		{
           "text": "%s"
       }
		`,
		givenOriginalText,
	)

	smsPreviewRequest := sms.NewPreviewRequest(givenOriginalText)

	actualRequest, _ := json.Marshal(smsPreviewRequest)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SmsSendPreviewEndpoint, givenResponse, 200)

	responseBody, r, _ := infobipClient.
		SmsAPI.
		PreviewSmsMessage(context.Background()).
		PreviewRequest(*smsPreviewRequest).
		Execute()

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenOriginalText, responseBody.GetOriginalText())
	assert.NotNil(t, responseBody.GetPreviews())
	assert.True(t, len(responseBody.GetPreviews()) == 1)
	assert.Equal(t, givenTextPreview, responseBody.GetPreviews()[0].GetTextPreview())
	assert.Equal(t, givenMessageCount, responseBody.GetPreviews()[0].GetMessageCount())
	assert.Equal(t, givenCharactersRemaining, responseBody.GetPreviews()[0].GetCharactersRemaining())
}

func TestShouldGetDeliveryReports(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenMessageId := "MESSAGE-ID-123-xyz"
	givenTo := "41793026727"
	givenSentAt := "2019-11-09T16:00:00.000+0000"
	givenDoneAt := "2019-11-09T16:00:00.000+0000"
	givenSmsCount := 1
	givenPricePerMessage := 0.01
	givenCurrency := "EUR"
	givenGroupId := 3
	givenGroupName := sms.MESSAGEGENERALSTATUS_DELIVERED // Assuming you have these constants defined in the Infobip Go SDK
	givenId := 5
	givenName := "DELIVERED_TO_HANDSET"
	givenDescription := "Message delivered to handset"
	givenErrorGroupId := 0
	givenErrorGroupName := sms.MESSAGEERRORGROUP_OK // Assuming you have these constants defined
	givenErrorId := 0
	givenErrorName := "NO_ERROR"
	givenErrorDescription := "No Error"
	givenErrorPermanent := false
	givenMessageId2 := "12db39c3-7822-4e72-a3ec-c87442c0ffc5"
	givenTo2 := "41793026834"
	givenSender := "InfoSMS"
	givenEntityId := "promotional-traffic-entity"
	givenApplicationId := "marketing-automation-application"

	// Expected response
	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"bulkId": "%s",
				"price": {
					"pricePerMessage": %f,
					"currency": "%s"
				},
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				},
				"error": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s",
					"permanent": %t
				},
				"messageId": "%s",
				"to": "%s",
				"sender": "%s",
				"sentAt": "%s",
				"doneAt": "%s",
				"messageCount": %d,
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				}
			},
			{
				"bulkId": "%s",
				"price": {
					"pricePerMessage": %f,
					"currency": "%s"
				},
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				},
				"error": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s",
					"permanent": %t
				},
				"messageId": "%s",
				"to": "%s",
				"sender": "%s",
				"sentAt": "%s",
				"doneAt": "%s",
				"messageCount": %d,
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				}
			}
		]
	}`,
		givenBulkId, givenPricePerMessage, givenCurrency, givenGroupId, givenGroupName, givenId, givenName, givenDescription,
		givenErrorGroupId, givenErrorGroupName, givenErrorId, givenErrorName, givenErrorDescription, givenErrorPermanent,
		givenMessageId, givenTo, givenSender, givenSentAt, givenDoneAt, givenSmsCount, givenEntityId, givenApplicationId,
		givenBulkId, givenPricePerMessage, givenCurrency, givenGroupId, givenGroupName, givenId, givenName, givenDescription,
		givenErrorGroupId, givenErrorGroupName, givenErrorId, givenErrorName, givenErrorDescription, givenErrorPermanent,
		givenMessageId2, givenTo2, givenSender, givenSentAt, givenDoneAt, givenSmsCount, givenEntityId, givenApplicationId)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", SmsReportsEndpoint, givenResponse, 200) // Assuming you have SetUpSuccessGetRequest defined

	// Get delivery reports
	smsAPI := infobipClient.SmsAPI
	deliveryReportsResponse, _, err := smsAPI.GetOutboundSmsMessageDeliveryReports(context.Background()).Execute()

	if err != nil {
		t.Errorf("Error getting delivery reports: %v", err)
	}

	// Assertions
	assert.NotNil(t, deliveryReportsResponse)
	assert.NotNil(t, deliveryReportsResponse.GetResults())

	results := deliveryReportsResponse.GetResults()
	assert.Len(t, results, 2)

	result1 := results[0]
	assert.Equal(t, givenBulkId, result1.GetBulkId())
	assert.Equal(t, givenMessageId, result1.GetMessageId())
	assert.Equal(t, givenTo, result1.GetTo())
	assert.Equal(t, givenSentAt, result1.GetSentAt().String())
	assert.Equal(t, givenDoneAt, result1.GetDoneAt().String())
	assert.Equal(t, int32(givenSmsCount), result1.GetMessageCount())

	price1 := result1.GetPrice()
	assert.NotNil(t, price1)
	assert.Equal(t, float32(givenPricePerMessage), price1.GetPricePerMessage())
	assert.Equal(t, givenCurrency, price1.GetCurrency())

	status1 := result1.GetStatus()
	assert.Equal(t, int32(givenGroupId), status1.GetGroupId())
	assert.Equal(t, givenGroupName, status1.GetGroupName())
	assert.Equal(t, int32(givenId), status1.GetId())
	assert.Equal(t, givenName, status1.GetName())
	assert.Equal(t, givenDescription, status1.GetDescription())

	error1 := result1.GetError()
	assert.Equal(t, int32(givenErrorGroupId), error1.GetGroupId())
	assert.Equal(t, givenErrorGroupName, error1.GetGroupName())
	assert.Equal(t, int32(givenErrorId), error1.GetId())
	assert.Equal(t, givenErrorName, error1.GetName())
	assert.Equal(t, givenErrorDescription, error1.GetDescription())
	assert.Equal(t, givenErrorPermanent, error1.GetPermanent())

	result2 := results[1]
	assert.Equal(t, givenBulkId, result2.GetBulkId())
	assert.Equal(t, givenMessageId2, result2.GetMessageId())
	assert.Equal(t, givenTo2, result2.GetTo())
	assert.Equal(t, givenSentAt, result2.GetSentAt().String())
	assert.Equal(t, givenDoneAt, result2.GetDoneAt().String())
	assert.Equal(t, int32(givenSmsCount), result2.GetMessageCount())

	price2 := result2.GetPrice()
	assert.NotNil(t, price2)
	assert.Equal(t, float32(givenPricePerMessage), price2.GetPricePerMessage())
	assert.Equal(t, givenCurrency, price2.GetCurrency())

	status2 := result2.GetStatus()
	assert.Equal(t, int32(givenGroupId), status2.GetGroupId())
	assert.Equal(t, givenGroupName, status2.GetGroupName())
	assert.Equal(t, int32(givenId), status2.GetId())
	assert.Equal(t, givenName, status2.GetName())
	assert.Equal(t, givenDescription, status2.GetDescription())

	error2 := result2.GetError()
	assert.Equal(t, int32(givenErrorGroupId), error2.GetGroupId())
	assert.Equal(t, givenErrorGroupName, error2.GetGroupName())
	assert.Equal(t, int32(givenErrorId), error2.GetId())
	assert.Equal(t, givenErrorName, error2.GetName())
	assert.Equal(t, givenErrorDescription, error2.GetDescription())
	assert.Equal(t, givenErrorPermanent, error2.GetPermanent())

}

func TestShouldGetSmsLogs(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenSentSinceString := "2015-02-22T17:42:05.390+0100"
	givenSentSince, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSentSinceString)
	ibTimeGivenSentAt := infobip.Time{
		T: givenSentSince,
	}

	givenMessageIdMessage1 := "MESSAGE-ID-123-xyz"
	givenToMessage1 := "41793026727"
	givenSendAtMessage1 := "2019-11-09T16:00:00.000+0000"
	givenDoneAtMessage1 := "2019-11-09T16:00:00.000+0000"
	givenSmsCountMessage1 := 1
	givenPricePerMessageMessage1 := 0.01
	givenCurrencyMessage1 := "EUR"

	givenMessageIdMessage2 := "12db39c3-7822-4e72-a3ec-c87442c0ffc5"
	givenToMessage2 := "41793026834"
	givenSendAtMessage2 := "2019-11-09T17:00:00.000+0000"
	givenDoneAtMessage2 := "2019-11-09T17:00:00.000+0000"
	givenSmsCountMessage2 := 5
	givenPricePerMessageMessage2 := 0.05
	givenCurrencyMessage2 := "HRK"
	givenApplicationId := "applicationId"
	givenEntityId := "entityId"
	// Expected response
	givenResponse := fmt.Sprintf(`{
       "results": [
           {
               "bulkId": "%s",
               "platform": {
                   "applicationId": "%s",
                   "entityId": "%s"
               },
               "messageId": "%s",
               "destination": "%s",
               "sentAt": "%s",
               "doneAt": "%s",
               "smsCount": %d,
               "price": {
                   "pricePerMessage": %g,
                   "currency": "%s"
               },
               "status": {
                   "groupId": %d,
                   "groupName": "%s",
                   "id": %d,
                   "name": "%s",
                   "description": "%s"
               },
               "error": {
                   "groupId": %d,
                   "groupName": "%s",
                   "id": %d,
                   "name": "%s",
                   "description": "%s",
                   "permanent": %t
               }
           },
           {
               "bulkId": "%s",
               "messageId": "%s",
               "destination": "%s",
               "sentAt": "%s",
               "doneAt": "%s",
               "smsCount": %d,
               "price": {
                   "pricePerMessage": %g,
                   "currency": "%s"
               },
               "status": {
                   "groupId": %d,
                   "groupName": "%s",
                   "id": %d,
                   "name": "%s",
                   "description": "%s"
               },
               "error": {
                   "groupId": %d,
                   "groupName": "%s",
                   "id": %d,
                   "name": "%s",
                   "description": "%s",
                   "permanent": %t
               }
           }
       ]
   }`,
		givenBulkId, givenApplicationId, givenEntityId, givenMessageIdMessage1, givenToMessage1,
		givenSendAtMessage1, givenDoneAtMessage1, givenSmsCountMessage1, givenPricePerMessageMessage1, givenCurrencyMessage1,
		StatusGroupId, StatusGroupName, StatusId, StatusName, StatusDescription,
		NoErrorStatusGroupId, NoErrorStatusGroupName, NoErrorStatusId, NoErrorStatusName, NoErrorStatusDescription, NoErrorPermanent,
		givenBulkId, givenMessageIdMessage2, givenToMessage2,
		givenSendAtMessage2, givenDoneAtMessage2, givenSmsCountMessage2, givenPricePerMessageMessage2, givenCurrencyMessage2,
		StatusGroupId, StatusGroupName, StatusId, StatusName, StatusDescription,
		NoErrorStatusGroupId, NoErrorStatusGroupName, NoErrorStatusId, NoErrorStatusName, NoErrorStatusDescription, NoErrorPermanent)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", SmsLogsEndpoint, givenResponse, 200)

	// Get SMS logs
	smsAPI := infobipClient.SmsAPI
	smsLogsResponse, _, err := smsAPI.GetOutboundSmsMessageLogs(context.Background()).
		BulkId([]string{givenBulkId}).
		SentSince(ibTimeGivenSentAt).
		Execute()

	if err != nil {
		t.Errorf("Error getting SMS logs: %v", err)
	}

	// Assertions
	assert.NotNil(t, smsLogsResponse)
	results := smsLogsResponse.GetResults()
	assert.NotNil(t, results)
	assert.Len(t, results, 2)

	// Assert first log entry
	log := results[0]
	assert.NotNil(t, log)
	assert.Equal(t, givenBulkId, log.GetBulkId())
	assert.Equal(t, givenMessageIdMessage1, log.GetMessageId())
	assert.Equal(t, givenToMessage1, log.GetDestination())
	assert.Equal(t, givenSendAtMessage1, log.GetSentAt().String())
	assert.Equal(t, givenDoneAtMessage1, log.GetDoneAt().String())

	platform := log.GetPlatform()
	assert.NotNil(t, platform)
	assert.Equal(t, givenApplicationId, platform.GetApplicationId())
	assert.Equal(t, givenEntityId, platform.GetEntityId())

	price := log.GetPrice()
	assert.NotNil(t, price)
	assert.Equal(t, float32(givenPricePerMessageMessage1), price.GetPricePerMessage())
	assert.Equal(t, givenCurrencyMessage1, price.GetCurrency())

	status := log.GetStatus()
	assert.Equal(t, int32(StatusGroupId), status.GetGroupId())
	assert.Equal(t, StatusGroupName, status.GetGroupName())
	assert.Equal(t, int32(StatusId), status.GetId())
	assert.Equal(t, StatusName, status.GetName())
	assert.Equal(t, StatusDescription, status.GetDescription())

	resultError := log.GetError()
	assert.Equal(t, int32(NoErrorStatusGroupId), resultError.GetGroupId())
	assert.Equal(t, NoErrorStatusGroupName, resultError.GetGroupName())
	assert.Equal(t, int32(NoErrorStatusId), resultError.GetId())
	assert.Equal(t, NoErrorStatusName, resultError.GetName())
	assert.Equal(t, NoErrorStatusDescription, resultError.GetDescription())
	assert.Equal(t, NoErrorPermanent, resultError.GetPermanent())

	// Assert second log entry
	anotherLog := results[1]
	assert.NotNil(t, anotherLog)
	assert.Equal(t, givenBulkId, anotherLog.GetBulkId())
	assert.Equal(t, givenMessageIdMessage2, anotherLog.GetMessageId())
	assert.Equal(t, givenToMessage2, anotherLog.GetDestination())
	assert.Equal(t, givenSendAtMessage2, anotherLog.GetSentAt().String())
	assert.Equal(t, givenDoneAtMessage2, anotherLog.GetDoneAt().String())

	price2 := anotherLog.GetPrice()
	assert.NotNil(t, price2)
	assert.Equal(t, float32(givenPricePerMessageMessage2), price2.GetPricePerMessage())
	assert.Equal(t, givenCurrencyMessage2, price2.GetCurrency())

	status2 := anotherLog.GetStatus()
	assert.Equal(t, int32(StatusGroupId), status2.GetGroupId())
	assert.Equal(t, StatusGroupName, status2.GetGroupName())
	assert.Equal(t, int32(StatusId), status2.GetId())
	assert.Equal(t, StatusName, status2.GetName())
	assert.Equal(t, StatusDescription, status2.GetDescription())

	resultError2 := anotherLog.GetError()
	assert.Equal(t, int32(NoErrorStatusGroupId), resultError2.GetGroupId())
	assert.Equal(t, NoErrorStatusGroupName, resultError2.GetGroupName())
	assert.Equal(t, int32(NoErrorStatusId), resultError2.GetId())
	assert.Equal(t, NoErrorStatusName, resultError2.GetName())
	assert.Equal(t, NoErrorStatusDescription, resultError2.GetDescription())
	assert.Equal(t, NoErrorPermanent, resultError2.GetPermanent())
}

func TestShouldGetScheduledSmsMessages(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenSendAtString := "2021-08-25T16:00:00.000+0000"

	// Expected response
	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "sendAt": "%s"
    }`, givenBulkId, givenSendAtString)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", SmsBulks, givenResponse, 200)

	// Get scheduled SMS messages
	smsAPI := infobipClient.SmsAPI
	bulkResponse, _, err := smsAPI.GetScheduledSmsMessages(context.Background()).BulkId(givenBulkId).Execute()

	if err != nil {
		t.Errorf("Error getting scheduled SMS messages: %v", err)
	}

	// Assertions
	assert.NotNil(t, bulkResponse)
	assert.Equal(t, givenBulkId, bulkResponse.GetBulkId())
	assert.Equal(t, givenSendAtString, bulkResponse.GetSendAt().String())
}

func TestShouldRescheduleSmsMessages(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenSendAtString := "2021-02-22T17:42:05.390+0100"

	// Expected request body
	expectedRequest := fmt.Sprintf(`{
        "sendAt": "%s"
    }`, givenSendAtString)

	// Prepare the request
	givenSendAtTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAtString)
	ibTimeGivenSentAt := infobip.Time{
		T: givenSendAtTime,
	}
	bulkRequest := sms.BulkRequest{
		SendAt: ibTimeGivenSentAt,
	}

	// Expected response
	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "sendAt": "%s"
    }`, givenBulkId, givenSendAtString)

	actualRequest, _ := json.Marshal(bulkRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", SmsBulks, givenResponse, 200)

	// Reschedule SMS messages
	smsAPI := infobipClient.SmsAPI
	bulkResponse, _, err := smsAPI.RescheduleSmsMessages(context.Background()).BulkId(givenBulkId).BulkRequest(bulkRequest).Execute()

	if err != nil {
		t.Errorf("Error rescheduling SMS messages: %v", err)
	}

	// Assertions
	assert.NotNil(t, bulkResponse)
	assert.Equal(t, givenBulkId, bulkResponse.GetBulkId())
	assert.Equal(t, givenSendAtString, bulkResponse.GetSendAt().String())
}

func TestShouldGetScheduledSmsMessagesStatus(t *testing.T) {
	// Given values
	givenBulkId := "BULK-ID-123-xyz"
	givenBulkStatusString := "PAUSED"
	givenBulkStatus := sms.BULKSTATUS_PAUSED

	// Expected response
	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "status": "%s"
    }`, givenBulkId, givenBulkStatusString)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", SmsBulksStatus, givenResponse, 200)

	// Get scheduled SMS messages status
	smsAPI := infobipClient.SmsAPI
	bulkStatusResponse, _, err := smsAPI.GetScheduledSmsMessagesStatus(context.Background()).BulkId(givenBulkId).Execute()

	if err != nil {
		t.Errorf("Error getting scheduled SMS messages status: %v", err)
	}

	// Assertions
	assert.NotNil(t, bulkStatusResponse)
	assert.Equal(t, givenBulkId, bulkStatusResponse.GetBulkId())
	assert.Equal(t, givenBulkStatus, bulkStatusResponse.GetStatus())
}

func TestShouldGetReceivedSmsMessages(t *testing.T) {
	// Given values
	givenMessageId := "817790313235066447"
	givenFrom := "385916242493"
	givenTo := "385921004026"
	givenText := "QUIZ Correct answer is Paris"
	givenCleanText := "Correct answer is Paris"
	givenKeyword := "QUIZ"
	givenReceivedAtString := "2021-08-25T16:10:00.000+0500"
	givenSmsCount := 1
	givenPricePerMessage := 0.0
	givenCurrency := "EUR"
	givenCallbackData := "callbackData"
	givenMessageCount := 1
	givenPendingMessageCount := 0
	givenApplicationId := "applicationId"
	givenEntityId := "entityId"
	givenLimit := int32(2)

	// Expected response
	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "messageId": "%s",
                "from": "%s",
                "to": "%s",
                "text": "%s",
                "cleanText": "%s",
                "keyword": "%s",
                "receivedAt": "%s",
                "smsCount": %d,
                "price": {
                    "pricePerMessage": %g,
                    "currency": "%s"
                },
                "callbackData": "%s",
                "applicationId": "%s",
                "entityId": "%s"
            }
        ],
        "messageCount": %d,
        "pendingMessageCount": %d
    }`,
		givenMessageId, givenFrom, givenTo, givenText, givenCleanText, givenKeyword, givenReceivedAtString,
		givenSmsCount, givenPricePerMessage, givenCurrency, givenCallbackData, givenApplicationId, givenEntityId,
		givenMessageCount, givenPendingMessageCount)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", SmsInboxReports, givenResponse, 200)

	// Get received SMS messages
	smsAPI := infobipClient.SmsAPI
	inboundMessageResult, _, err := smsAPI.GetInboundSmsMessages(context.Background()).Limit(givenLimit).Execute()

	if err != nil {
		t.Errorf("Error getting received SMS messages: %v", err)
	}

	// Assertions
	assert.NotNil(t, inboundMessageResult)
	assert.Equal(t, int32(givenMessageCount), inboundMessageResult.GetMessageCount())
	assert.Equal(t, int32(givenPendingMessageCount), inboundMessageResult.GetPendingMessageCount())

	messages := inboundMessageResult.Results
	assert.NotNil(t, messages)
	assert.Len(t, messages, 1)

	message := messages[0]
	assert.NotNil(t, message)
	assert.Equal(t, givenMessageId, message.GetMessageId())
	assert.Equal(t, givenFrom, message.GetFrom())
	assert.Equal(t, givenTo, message.GetTo())
	assert.Equal(t, givenText, message.GetText())
	assert.Equal(t, givenCleanText, message.GetCleanText())
	assert.Equal(t, givenKeyword, message.GetKeyword())
	assert.Equal(t, givenReceivedAtString, message.GetReceivedAt().String())
	assert.Equal(t, int32(givenSmsCount), message.GetSmsCount())
	assert.Equal(t, givenCallbackData, message.GetCallbackData())
	assert.Equal(t, givenApplicationId, message.GetApplicationId())
	assert.Equal(t, givenEntityId, message.GetEntityId())

	price := message.Price
	assert.NotNil(t, price)
	assert.Equal(t, givenPricePerMessage, price.GetPricePerMessage())
	assert.Equal(t, givenCurrency, price.GetCurrency())
}
