package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/email"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	EmailBulksEndpoint                      = "/email/1/bulks"
	EmailBulksStatusEndpoint                = "/email/1/bulks/status"
	EmailDomainsEndpoint                    = "/email/1/domains"
	EmailDomainByIdEndpoint                 = "/email/1/domains/%s"
	EmailDomainReturnPathEndpoint           = "/email/1/domains/%s/return-path"
	EmailDomainTrackingEndpoint             = "/email/1/domains/%s/tracking"
	EmailDomainVerifyEndpoint               = "/email/1/domains/%s/verify"
	EmailIpManagementDomainByIdEndpoint     = "/email/1/ip-management/domains/%d"
	EmailIpManagementDomainPoolsEndpoint    = "/email/1/ip-management/domains/%d/pools"
	EmailIpManagementDomainPoolByIdEndpoint = "/email/1/ip-management/domains/%d/pools/%s"
	EmailIpManagementIpsEndpoint            = "/email/1/ip-management/ips"
	EmailIpManagementIpByIdEndpoint         = "/email/1/ip-management/ips/%s"
	EmailIpManagementPoolsEndpoint          = "/email/1/ip-management/pools"
	EmailIpManagementPoolByIdEndpoint       = "/email/1/ip-management/pools/%s"
	EmailIpManagementPoolIpsEndpoint        = "/email/1/ip-management/pools/%s/ips"
	EmailIpManagementPoolIpByIdEndpoint     = "/email/1/ip-management/pools/%s/ips/%s"
	EmailSendEndpoint                       = "/email/3/send"
	EmailLogsEndpoint                       = "/email/1/logs"
	EmailReportsEndpoint                    = "/email/1/reports"
	EmailSuppressionsEndpoint               = "/email/1/suppressions"
	EmailSuppressionDomainsEndpoint         = "/email/1/suppressions/domains"
	EmailValidationEndpoint                 = "/email/2/validation"
)

func TestShouldSendFullyFeaturedEmail(t *testing.T) {
	givenTo := "john.smith@somedomain.com"
	givenAnotherTo := `{"to": "alice.grey@somecompany.com","placeholders": {"name": "Alice"}}`
	anotherToInResponse := "alice.grey@somecompany.com"
	givenBulkId := "snxemd8u52v7v84iiu69"
	givenGroupId := int32(1)
	givenGroupName := "PENDING"
	givenId := int32(1)
	givenName := "PENDING_ACCEPTED"
	givenDescription := "Message accepted, pending for delivery."

	givenAttachmentText := "Test file text"
	tempFile, err := os.CreateTemp("", "attachment*.txt")
	assert.Nil(t, err)
	defer os.Remove(tempFile.Name())
	_, err = tempFile.WriteString(givenAttachmentText)
	assert.Nil(t, err)

	givenMessageId := "somExternalMessageId0"
	givenAnotherMessageId := "someExternalMessageId1"

	givenFrom := "Jane Smith <jane.smith@somecompany.com>"
	givenReplyTo := "all.replies@somedomain.com"
	givenSubject := "Mail subject text"
	givenText := "Mail body text"
	givenHtml := "<h1>Html body</h1><p>Rich HTML message body.</p>"
	intermediateReport := true
	givenNotifyUrl := "https://www.example.com/email/advanced"
	givenNotifyContentType := "application/json"
	givenCallbackData := "DLR callback data"

	expectedResponse := fmt.Sprintf(`{
		"bulkId": "%s",
		"messages": [
			{
				"to": "%s",
				"messageId": "%s",
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				}
			},
			{
				"to": "%s",
				"messageId": "%s",
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				}
			}
		]
	}`, givenBulkId, givenTo, givenMessageId, givenGroupId, givenGroupName, givenId, givenName, givenDescription, anotherToInResponse, givenAnotherMessageId, givenGroupId, givenGroupName, givenId, givenName, givenDescription)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", EmailSendEndpoint, expectedResponse, 200)

	response, _, err := infobipClient.EmailAPI.SendEmail(context.Background()).
		To([]string{givenTo, givenAnotherTo}).
		From(givenFrom).
		Subject(givenSubject).
		ReplyTo(givenReplyTo).
		Html(givenHtml).
		Text(givenText).
		Attachment([]*os.File{tempFile}).
		IntermediateReport(intermediateReport).
		NotifyUrl(givenNotifyUrl).
		NotifyContentType(givenNotifyContentType).
		CallbackData(givenCallbackData).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.NotNil(t, response.GetMessages())
	messages := response.GetMessages()
	assert.Equal(t, 2, len(messages))

	expectedStatus := email.SingleMessageStatus{
		GroupId:     givenGroupId,
		GroupName:   givenGroupName,
		Id:          givenId,
		Name:        givenName,
		Description: givenDescription,
	}

	firstMessage := messages[0]
	assert.Equal(t, givenTo, firstMessage.GetTo())
	assert.Equal(t, givenMessageId, firstMessage.GetMessageId())
	assert.Equal(t, expectedStatus, firstMessage.GetStatus())

	secondMessage := messages[1]
	assert.Equal(t, anotherToInResponse, secondMessage.GetTo())
	assert.Equal(t, givenAnotherMessageId, secondMessage.GetMessageId())
	assert.Equal(t, expectedStatus, secondMessage.GetStatus())
}

func TestShouldGetEmailSuppressions(t *testing.T) {
	givenDomainName := "example.com"
	givenEmailAddress := "jane.smith@somecompany.com"
	givenType := email.APISUPPRESSIONTYPE_BOUNCE
	givenCreatedDate := "2024-08-14T14:02:17.366"
	givenCreatedDateTime, _ := time.Parse("2006-01-02T15:04:05.000", givenCreatedDate)
	ibTimeCreatedAt := infobip.Time{
		T: givenCreatedDateTime,
	}
	givenReason := "550 5.1.1 <jane.smith@somecompany.com>: user does not exist"
	givenPage := 0
	givenSize := 100

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"domainName": "%s",
				"emailAddress": "%s",
				"type": "%s",
				"createdDate": "%s",
				"reason": "%s"
			}
		],
		"paging": {
			"page": %d,
			"size": %d
		}
	}`, givenDomainName, givenEmailAddress, givenType, givenCreatedDate, givenReason, givenPage, givenSize)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailSuppressionsEndpoint, givenResponse, 200)

	// Execute the request
	response, _, err := infobipClient.
		EmailAPI.
		GetSuppressions(context.Background()).
		DomainName(givenDomainName).
		Type_(givenType).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.GetResults()))

	result := response.GetResults()[0]
	assert.Equal(t, givenDomainName, result.GetDomainName())
	assert.Equal(t, givenEmailAddress, result.GetEmailAddress())
	assert.Equal(t, string(givenType), result.GetType())
	assert.Equal(t, ibTimeCreatedAt, result.GetCreatedDate())
	assert.Equal(t, givenReason, result.GetReason())

	paging := response.GetPaging()
	assert.Equal(t, int32(givenPage), paging.GetPage())
	assert.Equal(t, int32(givenSize), paging.GetSize())
}

func TestShouldAddEmailSuppressions(t *testing.T) {
	givenDomainName1 := "example.com"
	givenEmailAddresses1 := []string{"jane.smith@somecompany.com", "john.doe@somecompany.com"}
	givenType := email.APIADDSUPPRESSIONTYPE_BOUNCE

	givenDomainName2 := "example.com"
	givenEmailAddresses2 := []string{"john.smith@somecompany.com", "john.perry@gmail.com"}

	givenRequest := fmt.Sprintf(`{
        "suppressions": [
            {
                "domainName": "%s",
                "emailAddress": ["%s", "%s"],
                "type": "%s"
            },
            {
                "domainName": "%s",
                "emailAddress": ["%s", "%s"],
                "type": "%s"
            }
        ]
    }`, givenDomainName1, givenEmailAddresses1[0], givenEmailAddresses1[1], givenType, givenDomainName2, givenEmailAddresses2[0], givenEmailAddresses2[1], givenType)

	request := email.NewAddSuppressionRequest([]email.AddSuppression{
		*email.NewAddSuppression(givenDomainName1, givenEmailAddresses1, givenType),
		*email.NewAddSuppression(givenDomainName2, givenEmailAddresses2, givenType),
	})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", EmailSuppressionsEndpoint, "", 204)

	_, err := infobipClient.
		EmailAPI.
		AddSuppressions(context.Background()).
		AddSuppressionRequest(*request).
		Execute()

	assert.Nil(t, err)
}

func TestShouldDeleteEmailSuppressions(t *testing.T) {
	givenDomainName1 := "example.com"
	givenEmailAddresses1 := []string{"jane.smith@somecompany.com", "john.doe@somecompany.com"}
	givenType := email.APISUPPRESSIONTYPE_BOUNCE

	givenDomainName2 := "example.com"
	givenEmailAddresses2 := []string{"john.smith@somecompany.com", "john.perry@gmail.com"}

	givenRequest := fmt.Sprintf(`{
       "suppressions": [
           {
               "domainName": "%s",
               "emailAddress": ["%s", "%s"],
               "type": "%s"
           },
           {
               "domainName": "%s",
               "emailAddress": ["%s", "%s"],
               "type": "%s"
           }
       ]
   }`, givenDomainName1, givenEmailAddresses1[0], givenEmailAddresses1[1], givenType, givenDomainName2, givenEmailAddresses2[0], givenEmailAddresses2[1], givenType)

	request := email.NewDeleteSuppressionRequest([]email.DeleteSuppression{
		*email.NewDeleteSuppression(givenDomainName1, givenEmailAddresses1, givenType),
		*email.NewDeleteSuppression(givenDomainName2, givenEmailAddresses2, givenType),
	})
	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", EmailSuppressionsEndpoint, "", 204)

	_, err := infobipClient.
		EmailAPI.
		DeleteSuppressions(context.Background()).
		DeleteSuppressionRequest(*request).
		Execute()

	assert.Nil(t, err)
}

func TestShouldGetSuppressionDomains(t *testing.T) {
	givenDomainName1 := "example.com"
	givenDataAccess1 := email.APIDOMAINACCESS_OWNER
	givenReadBounces1 := true
	givenCreateBounces1 := true
	givenDeleteBounces1 := true
	givenReadComplaints1 := true
	givenCreateComplaints1 := true
	givenDeleteComplaints1 := true
	givenReadOverquotas1 := true

	givenDomainName2 := "example.com"
	givenDataAccess2 := email.APIDOMAINACCESS_GRANTED
	givenReadBounces2 := true
	givenCreateBounces2 := true
	givenDeleteBounces2 := false
	givenReadComplaints2 := true
	givenCreateComplaints2 := false
	givenDeleteComplaints2 := false
	givenReadOverquotas2 := false

	givenPage := 0
	givenSize := 100

	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "domainName": "%s",
                "dataAccess": "%s",
                "readBounces": %t,
                "createBounces": %t,
                "deleteBounces": %t,
                "readComplaints": %t,
                "createComplaints": %t,
                "deleteComplaints": %t,
                "readOverquotas": %t
            },
            {
                "domainName": "%s",
                "dataAccess": "%s",
                "readBounces": %t,
                "createBounces": %t,
                "deleteBounces": %t,
                "readComplaints": %t,
                "createComplaints": %t,
                "deleteComplaints": %t,
                "readOverquotas": %t
            }
        ],
        "paging": {
            "page": %d,
            "size": %d
        }
    }`, givenDomainName1, givenDataAccess1, givenReadBounces1, givenCreateBounces1, givenDeleteBounces1, givenReadComplaints1, givenCreateComplaints1, givenDeleteComplaints1, givenReadOverquotas1,
		givenDomainName2, givenDataAccess2, givenReadBounces2, givenCreateBounces2, givenDeleteBounces2, givenReadComplaints2, givenCreateComplaints2, givenDeleteComplaints2, givenReadOverquotas2,
		givenPage, givenSize)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailSuppressionDomainsEndpoint, givenResponse, 200)

	// Execute the request
	response, _, err := infobipClient.
		EmailAPI.
		GetDomains(context.Background()).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 2, len(response.GetResults()))

	result1 := response.GetResults()[0]
	assert.Equal(t, givenDomainName1, result1.GetDomainName())
	assert.Equal(t, givenDataAccess1, result1.GetDataAccess())
	assert.Equal(t, givenReadBounces1, result1.GetReadBounces())
	assert.Equal(t, givenCreateBounces1, result1.GetCreateBounces())
	assert.Equal(t, givenDeleteBounces1, result1.GetDeleteBounces())
	assert.Equal(t, givenReadComplaints1, result1.GetReadComplaints())
	assert.Equal(t, givenCreateComplaints1, result1.GetCreateComplaints())
	assert.Equal(t, givenDeleteComplaints1, result1.GetDeleteComplaints())
	assert.Equal(t, givenReadOverquotas1, result1.GetReadOverquotas())

	result2 := response.GetResults()[1]
	assert.Equal(t, givenDomainName2, result2.GetDomainName())
	assert.Equal(t, givenDataAccess2, result2.GetDataAccess())
	assert.Equal(t, givenReadBounces2, result2.GetReadBounces())
	assert.Equal(t, givenCreateBounces2, result2.GetCreateBounces())
	assert.Equal(t, givenDeleteBounces2, result2.GetDeleteBounces())
	assert.Equal(t, givenReadComplaints2, result2.GetReadComplaints())
	assert.Equal(t, givenCreateComplaints2, result2.GetCreateComplaints())
	assert.Equal(t, givenDeleteComplaints2, result2.GetDeleteComplaints())
	assert.Equal(t, givenReadOverquotas2, result2.GetReadOverquotas())

	paging := response.GetPaging()
	assert.Equal(t, int32(givenPage), paging.GetPage())
	assert.Equal(t, int32(givenSize), paging.GetSize())
}

func TestShouldAddDomain(t *testing.T) {
	givenDomainName := "example.com"
	givenDkimKeyLength := int32(1024)
	givenTargetedDailyTraffic := int64(15)

	givenDomainId := int64(1)
	givenActive := false
	givenTracking := true
	givenDnsRecords := "string"
	givenVerified := true
	givenBlocked := false
	givenCreatedAtString := "2021-08-25T16:00:00.000+0000"
	givenCreatedAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAtString)
	ibTimeCreatedAt := infobip.Time{
		T: givenCreatedAt,
	}

	givenResponse := fmt.Sprintf(`{
	   "domainId": %d,
	   "domainName": "%s",
	   "active": %t,
	   "tracking": {
	       "clicks": %t,
	       "opens": %t,
	       "unsubscribe": %t
	   },
	   "dnsRecords": [
	       {
	           "recordType": "%s",
	           "name": "%s",
	           "expectedValue": "%s",
	           "verified": %t
	       }
	   ],
	   "blocked": %t,
	   "createdAt": "%s"
	}`, givenDomainId, givenDomainName, givenActive, givenTracking, givenTracking, givenTracking, givenDnsRecords, givenDnsRecords, givenDnsRecords, givenVerified, givenBlocked, givenCreatedAtString)

	givenRequest := fmt.Sprintf(`{
        "domainName": "%s",
        "dkimKeyLength": %d,
        "targetedDailyTraffic": %d
    }`, givenDomainName, givenDkimKeyLength, givenTargetedDailyTraffic)

	request := email.NewAddDomainRequest(givenDomainName, givenTargetedDailyTraffic)
	dkimKeyLength := int32(1024)
	request.DkimKeyLength = &dkimKeyLength

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", EmailDomainsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		AddDomain(context.Background()).
		AddDomainRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, actualRequest)
	assert.Equal(t, givenDomainId, response.GetDomainId())
	assert.Equal(t, givenDomainName, response.GetDomainName())
	assert.Equal(t, givenActive, response.GetActive())
	assert.Equal(t, &givenTracking, response.GetTracking().Clicks)
	assert.Equal(t, &givenTracking, response.GetTracking().Opens)
	assert.Equal(t, &givenTracking, response.GetTracking().Unsubscribe)
	assert.Equal(t, 1, len(response.GetDnsRecords()))
	record := response.GetDnsRecords()[0]
	assert.Equal(t, givenDnsRecords, record.GetRecordType())
	assert.Equal(t, givenDnsRecords, record.GetName())
	assert.Equal(t, givenDnsRecords, record.GetExpectedValue())
	assert.Equal(t, givenVerified, record.GetVerified())
	assert.Equal(t, givenBlocked, response.GetBlocked())
	assert.Equal(t, ibTimeCreatedAt, response.GetCreatedAt())
}

func TestShouldGetAllDomains(t *testing.T) {
	givenDomainName := "example.com"
	givenPaging := int32(0)
	givenDomainId := int64(1)
	givenActive := false
	givenTracking := true
	givenDnsRecords := "string"
	givenVerified := true
	givenBlocked := false
	givenCreatedAtString := "2021-08-25T16:00:00.000+0000"
	givenCreatedAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAtString)
	ibTimeCreatedAt := infobip.Time{
		T: givenCreatedAt,
	}

	givenResponse := fmt.Sprintf(`{
       "paging": {
           "page": %d,
           "size": %d,
           "totalPages": %d,
           "totalResults": %d
       },
       "results": [
           {
               "domainId": %d,
               "domainName": "%s",
               "active": %t,
               "tracking": {
                   "clicks": %t,
                   "opens": %t,
                   "unsubscribe": %t
               },
               "dnsRecords": [
                   {
                       "recordType": "%s",
                       "name": "%s",
                       "expectedValue": "%s",
                       "verified": %t
                   }
               ],
               "blocked": %t,
               "createdAt": "%s"
           }
       ]
   }`, givenPaging, givenPaging, givenPaging, givenPaging, givenDomainId, givenDomainName, givenActive, givenTracking, givenTracking, givenTracking, givenDnsRecords, givenDnsRecords, givenDnsRecords, givenVerified, givenBlocked, givenCreatedAtString)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailDomainsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		GetAllDomains(context.Background()).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, &givenPaging, response.GetPaging().Page)
	assert.Equal(t, &givenPaging, response.GetPaging().Size)
	assert.Equal(t, &givenPaging, response.GetPaging().TotalPages)
	assert.Equal(t, &givenPaging, response.GetPaging().TotalResults)

	assert.Equal(t, 1, len(response.GetResults()))
	result := response.GetResults()[0]
	assert.Equal(t, givenDomainId, result.GetDomainId())
	assert.Equal(t, givenDomainName, result.GetDomainName())
	assert.Equal(t, givenActive, result.GetActive())
	assert.Equal(t, &givenTracking, result.GetTracking().Clicks)
	assert.Equal(t, &givenTracking, result.GetTracking().Opens)
	assert.Equal(t, &givenTracking, result.GetTracking().Unsubscribe)
	assert.Equal(t, 1, len(result.GetDnsRecords()))
	record := result.GetDnsRecords()[0]
	assert.Equal(t, givenDnsRecords, record.GetRecordType())
	assert.Equal(t, givenDnsRecords, record.GetName())
	assert.Equal(t, givenDnsRecords, record.GetExpectedValue())
	assert.Equal(t, givenVerified, record.GetVerified())
	assert.Equal(t, givenBlocked, result.GetBlocked())
	assert.Equal(t, ibTimeCreatedAt, result.GetCreatedAt())
}

func TestShouldGetDomainDetails(t *testing.T) {
	givenDomainName := "example.com"
	givenDomainId := int64(1)
	givenActive := false
	givenTracking := true
	givenDnsRecords := "string"
	givenVerified := true
	givenBlocked := false
	givenCreatedAtString := "2021-08-25T16:00:00.000+0000"
	givenCreatedAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAtString)
	ibTimeCreatedAt := infobip.Time{
		T: givenCreatedAt,
	}
	givenResponse := fmt.Sprintf(`{
        "domainId": %d,
        "domainName": "%s",
        "active": %t,
        "tracking": {
            "clicks": %t,
            "opens": %t,
            "unsubscribe": %t
        },
        "dnsRecords": [
            {
                "recordType": "%s",
                "name": "%s",
                "expectedValue": "%s",
                "verified": %t
            }
        ],
        "blocked": %t,
        "createdAt": "%s"
    }`, givenDomainId, givenDomainName, givenActive, givenTracking, givenTracking, givenTracking, givenDnsRecords, givenDnsRecords, givenDnsRecords, givenVerified, givenBlocked, givenCreatedAtString)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(EmailDomainByIdEndpoint, givenDomainName), givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		GetDomainDetails(context.Background(), givenDomainName).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenDomainId, response.GetDomainId())
	assert.Equal(t, givenDomainName, response.GetDomainName())
	assert.Equal(t, givenActive, response.GetActive())
	assert.Equal(t, &givenTracking, response.GetTracking().Clicks)
	assert.Equal(t, &givenTracking, response.GetTracking().Opens)
	assert.Equal(t, &givenTracking, response.GetTracking().Unsubscribe)
	assert.Equal(t, 1, len(response.GetDnsRecords()))
	record := response.GetDnsRecords()[0]
	assert.Equal(t, givenDnsRecords, record.GetRecordType())
	assert.Equal(t, givenDnsRecords, record.GetName())
	assert.Equal(t, givenDnsRecords, record.GetExpectedValue())
	assert.Equal(t, givenVerified, record.GetVerified())
	assert.Equal(t, givenBlocked, response.GetBlocked())
	assert.Equal(t, ibTimeCreatedAt, response.GetCreatedAt())
}

func TestShouldDeleteDomain(t *testing.T) {
	givenDomainName := "example.com"
	givenStatusCode := 204

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf(EmailDomainByIdEndpoint, givenDomainName), "", givenStatusCode)

	_, err := infobipClient.
		EmailAPI.
		DeleteDomain(context.Background(), givenDomainName).
		Execute()

	assert.Nil(t, err)
}

func TestShouldUpdateTrackingEvents(t *testing.T) {
	givenDomainName := "example.com"
	givenDomainId := int64(1)
	givenActive := false
	givenTracking := true
	givenUnsubscribe := true
	givenDnsRecords := "string"
	givenVerified := true
	givenBlocked := false
	givenCreatedAtString := "2021-08-25T16:00:00.000+0000"
	givenCreatedAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAtString)
	ibTimeCreatedAt := infobip.Time{
		T: givenCreatedAt,
	}
	givenResponse := fmt.Sprintf(`{
       "domainId": %d,
       "domainName": "%s",
       "active": %t,
       "tracking": {
           "clicks": %t,
           "opens": %t,
           "unsubscribe": %t
       },
       "dnsRecords": [
           {
               "recordType": "%s",
               "name": "%s",
               "expectedValue": "%s",
               "verified": %t
           }
       ],
       "blocked": %t,
       "createdAt": "%s"
   }`, givenDomainId, givenDomainName, givenActive, givenTracking, givenTracking, givenUnsubscribe, givenDnsRecords, givenDnsRecords, givenDnsRecords, givenVerified, givenBlocked, givenCreatedAtString)

	givenRequest := fmt.Sprintf(`{
       "open": %t,
       "clicks": %t,
       "unsubscribe": %t
   }`, givenTracking, givenTracking, givenUnsubscribe)

	request := email.NewTrackingEventRequest()
	request.SetOpen(givenTracking)
	request.SetClicks(givenTracking)
	request.SetUnsubscribe(givenUnsubscribe)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf(EmailDomainTrackingEndpoint, givenDomainName), givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		UpdateTrackingEvents(context.Background(), givenDomainName).
		TrackingEventRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenDomainId, response.GetDomainId())
	assert.Equal(t, givenDomainName, response.GetDomainName())
	assert.Equal(t, givenActive, response.GetActive())
	assert.Equal(t, &givenTracking, response.GetTracking().Clicks)
	assert.Equal(t, &givenTracking, response.GetTracking().Opens)
	assert.Equal(t, &givenTracking, response.GetTracking().Unsubscribe)
	assert.Equal(t, 1, len(response.GetDnsRecords()))
	record := response.GetDnsRecords()[0]
	assert.Equal(t, givenDnsRecords, record.GetRecordType())
	assert.Equal(t, givenDnsRecords, record.GetName())
	assert.Equal(t, givenDnsRecords, record.GetExpectedValue())
	assert.Equal(t, givenVerified, record.GetVerified())
	assert.Equal(t, givenBlocked, response.GetBlocked())
	assert.Equal(t, ibTimeCreatedAt, response.GetCreatedAt())
}

func TestShouldVerifyDomain(t *testing.T) {
	givenDomainName := "example.com"
	givenStatusCode := 202

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf(EmailDomainVerifyEndpoint, givenDomainName), "", givenStatusCode)

	_, err := infobipClient.
		EmailAPI.
		VerifyDomain(context.Background(), givenDomainName).
		Execute()

	assert.Nil(t, err)
}

func TestShouldValidateEmail(t *testing.T) {
	givenTo := "john.smith@abc.com"
	givenValidSyntax := true
	givenDidYouMean := ""

	expectedRequest := fmt.Sprintf(`{
        "to": "%s"
    }`, givenTo)

	givenResponse := fmt.Sprintf(`{
        "to": "%s",
        "validMailbox": "true",
        "validSyntax": %t,
        "catchAll": false,
        "didYouMean": "%s",
        "disposable": false,
        "roleBased": true
    }`, givenTo, givenValidSyntax, givenDidYouMean)

	request := email.NewValidationRequest(givenTo)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", EmailValidationEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		ValidateEmailAddresses(context.Background()).
		ValidationRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenTo, response.GetTo())
	assert.Equal(t, givenValidSyntax, response.GetValidSyntax())
	assert.Equal(t, givenDidYouMean, response.GetDidYouMean())
}

func TestShouldGetScheduledEmails(t *testing.T) {
	givenBulkId := "BULK-ID-123-xyz"
	givenExternalBulkId := "SOME_USER_DEFINE_BULK_123"
	givenResponse := fmt.Sprintf(`{
       "externalBulkId": "%s",
       "bulks": [
           {
               "bulkId": "%s",
               "sendAt": "2021-08-25T16:00:00.000+0000"
           }
       ]
   }`, givenExternalBulkId, givenBulkId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailBulksEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		GetScheduledEmails(context.Background()).
		BulkId(givenExternalBulkId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenExternalBulkId, response.GetExternalBulkId())
	assert.Equal(t, 1, len(response.GetBulks()))
	assert.Equal(t, givenBulkId, response.GetBulks()[0].GetBulkId())
}

func TestShouldGetScheduledEmailsStatuses(t *testing.T) {
	givenBulkId := "BULK-ID-123-xyz"
	givenExternalBulkId := "SOME_USER_DEFINE_BULK_123"
	givenStatus := email.BULKSTATUS_PENDING
	givenResponse := fmt.Sprintf(`{
       "externalBulkId": "%s",
       "bulks": [
           {
               "bulkId": "%s",
               "status": "%s"
           }
       ]
   }`, givenExternalBulkId, givenBulkId, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailBulksStatusEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		GetScheduledEmailStatuses(context.Background()).
		BulkId(givenExternalBulkId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenExternalBulkId, response.GetExternalBulkId())
	assert.Equal(t, 1, len(response.GetBulks()))
	assert.Equal(t, givenBulkId, response.GetBulks()[0].GetBulkId())
	assert.Equal(t, givenStatus, response.GetBulks()[0].GetStatus())
}

func TestShouldRescheduleEmails(t *testing.T) {
	givenBulkId := "BULK-ID-123-xyz"
	givenSendAtString := "2021-06-08T17:42:05.390+0100"
	givenSendAt, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAtString)
	ibTimeGivenSendAt := infobip.Time{
		T: givenSendAt,
	}

	givenRequest := fmt.Sprintf(`{
        "sendAt": "%s"
    }`, givenSendAtString)

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "sendAt": "%s"
    }`, givenBulkId, givenSendAtString)

	request := email.NewBulkRescheduleRequest(ibTimeGivenSendAt)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", EmailBulksEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		RescheduleEmails(context.Background()).
		BulkId(givenBulkId).
		BulkRescheduleRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeGivenSendAt, response.GetSendAt())
}

func TestShouldUpdateScheduledEmailStatuses(t *testing.T) {
	givenBulkId := "string"
	givenStatus := email.BULKSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
       "bulkId": "%s",
       "status": "%s"
   }`, givenBulkId, givenStatus)

	expectedRequest := fmt.Sprintf(`{
       "status": "%s"
   }`, givenStatus)

	request := email.NewBulkUpdateStatusRequest(givenStatus)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", EmailBulksStatusEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		UpdateScheduledEmailStatuses(context.Background()).
		BulkId(givenBulkId).
		BulkUpdateStatusRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldGetEmailDeliveryReports(t *testing.T) {
	givenSentAt := "2021-08-25T16:10:00.000+0500"
	givenDoneAt := "2021-08-25T16:11:00.000+0500"
	expectedSentAt, _ := time.Parse(time.RFC3339, "2021-08-25T16:10:00+05:00")
	expectedDoneAt, _ := time.Parse(time.RFC3339, "2021-08-25T16:11:00+05:00")

	ibTimeSentAt := infobip.Time{
		T: expectedSentAt,
	}

	ibTimeDoneAt := infobip.Time{
		T: expectedDoneAt,
	}
	givenBulkId := "csdstgteet4fath2pclbq"
	givenMessageId := "45653761-3a88-4060-869e-ae372adc7a51"
	givenTo := "john.doe@email.com"

	expectedResponse := fmt.Sprintf(`{
        "results": [
            {
                "bulkId": "%s",
                "messageId": "%s",
                "to": "%s",
                "sentAt": "%s",
                "doneAt": "%s",
                "messageCount": 1,
                "price": {
                    "pricePerMessage": 0,
                    "currency": "UNKNOWN"
                },
                "status": {
                    "groupId": 3,
                    "groupName": "DELIVERED",
                    "id": 5,
                    "name": "DELIVERED_TO_HANDSET",
                    "description": "Message delivered to handset"
                },
                "error": {
                    "groupId": 0,
                    "groupName": "OK",
                    "id": 0,
                    "name": "NO_ERROR",
                    "description": "No Error",
                    "permanent": false
                },
                "channel": "EMAIL"
            }
        ]
    }`, givenBulkId, givenMessageId, givenTo, givenSentAt, givenDoneAt)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailReportsEndpoint, expectedResponse, 200)

	response, _, err := infobipClient.
		EmailAPI.
		GetEmailDeliveryReports(context.Background()).
		BulkId(givenBulkId).
		MessageId(givenMessageId).
		Limit(1).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.GetResults()))

	report := response.GetResults()[0]
	assert.Equal(t, givenBulkId, report.GetBulkId())
	assert.Equal(t, givenMessageId, report.GetMessageId())
	assert.Equal(t, givenTo, report.GetTo())
	assert.Equal(t, ibTimeSentAt, report.GetSentAt())
	assert.Equal(t, ibTimeDoneAt, report.GetDoneAt())

	status := report.GetStatus()
	assert.Equal(t, int32(3), status.GetGroupId())
	assert.Equal(t, "DELIVERED", status.GetGroupName())
	assert.Equal(t, int32(5), status.GetId())
	assert.Equal(t, "DELIVERED_TO_HANDSET", status.GetName())
	assert.Equal(t, "Message delivered to handset", status.GetDescription())

	error := report.GetError()
	assert.Equal(t, int32(0), error.GetGroupId())
	assert.Equal(t, "OK", error.GetGroupName())
	assert.Equal(t, int32(0), error.GetId())
	assert.Equal(t, "NO_ERROR", error.GetName())
	assert.Equal(t, "No Error", error.GetDescription())
	assert.Equal(t, false, error.GetPermanent())
}

func TestShouldGetIpManagementIps(t *testing.T) {
	givenId := "DB3F9D439088BF73F5560443C8054AC4"
	givenIp := "198.51.100.0"

	givenResponse := fmt.Sprintf(`[
  {
    "id": "%s",
    "ip": "%s"
  }
 ]`, givenId, givenIp)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailIpManagementIpsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.EmailAPI.GetAllIps(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, givenId, response[0].GetId())
	assert.Equal(t, givenIp, response[0].GetIp())
}

func TestShouldGetIpManagementIp(t *testing.T) {
	givenId := "DB3F9D439088BF73F5560443C8054AC4"
	givenIp := "198.51.100.0"
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenPoolName := "IP pool name"

	givenResponse := fmt.Sprintf(`{
  "id": "%s",
  "ip": "%s",
  "pools": [
    {
      "id": "%s",
      "name": "%s"
    }
  ]
 }`, givenId, givenIp, givenPoolId, givenPoolName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(EmailIpManagementIpByIdEndpoint, givenId), givenResponse, 200)

	response, _, err := infobipClient.EmailAPI.GetIpDetails(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenIp, response.GetIp())
	assert.Equal(t, 1, len(response.GetPools()))
	assert.Equal(t, givenPoolId, response.GetPools()[0].GetId())
	assert.Equal(t, givenPoolName, response.GetPools()[0].GetName())
}

func TestShouldGetIpManagementPools(t *testing.T) {
	givenId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenName := "IP pool name"

	givenResponse := fmt.Sprintf(`[
  {
    "id": "%s",
    "name": "%s"
  }
 ]`, givenId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", EmailIpManagementPoolsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.EmailAPI.GetIpPools(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, givenId, response[0].GetId())
	assert.Equal(t, givenName, response[0].GetName())
}

func TestShouldCreateIpManagementPool(t *testing.T) {
	givenId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenName := "IP pool name"

	givenRequest := email.IpPoolCreateRequest{
		Name: givenName,
	}

	expectedRequest := fmt.Sprintf(`{
  "name": "%s"
 }`, givenName)

	givenResponse := fmt.Sprintf(`{
  "id": "%s",
  "name": "%s"
 }`, givenId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", EmailIpManagementPoolsEndpoint, givenResponse, 201)

	actualRequest, _ := json.Marshal(givenRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.EmailAPI.CreateIpPool(context.Background()).IpPoolCreateRequest(givenRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
}

func TestShouldGetIpManagementPool(t *testing.T) {
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenPoolName := "IP pool name"
	givenIpId := "DB3F9D439088BF73F5560443C8054AC4"
	givenIp := "198.51.100.0"

	givenResponse := fmt.Sprintf(`{
  "id": "%s",
  "name": "%s",
  "ips": [
    {
      "id": "%s",
      "ip": "%s"
    }
  ]
 }`, givenPoolId, givenPoolName, givenIpId, givenIp)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(EmailIpManagementPoolByIdEndpoint, givenPoolId), givenResponse, 200)

	response, _, err := infobipClient.EmailAPI.GetIpPool(context.Background(), givenPoolId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenPoolId, response.GetId())
	assert.Equal(t, givenPoolName, response.GetName())
	assert.Equal(t, 1, len(response.GetIps()))
	assert.Equal(t, givenIpId, response.GetIps()[0].GetId())
	assert.Equal(t, givenIp, response.GetIps()[0].GetIp())
}

func TestShouldUpdateIpManagementPool(t *testing.T) {
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenName := "Updated IP pool name"

	givenRequest := email.IpPoolCreateRequest{
		Name: givenName,
	}

	expectedRequest := fmt.Sprintf(`{
  "name": "%s"
 }`, givenName)

	givenResponse := fmt.Sprintf(`{
  "id": "%s",
  "name": "%s"
 }`, givenPoolId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf(EmailIpManagementPoolByIdEndpoint, givenPoolId), givenResponse, 200)

	actualRequest, _ := json.Marshal(givenRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.EmailAPI.UpdateIpPool(context.Background(), givenPoolId).IpPoolCreateRequest(givenRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenPoolId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
}

func TestShouldDeleteIpManagementPool(t *testing.T) {
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenStatusCode := 204

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf(EmailIpManagementPoolByIdEndpoint, givenPoolId), "", givenStatusCode)

	_, err := infobipClient.EmailAPI.DeleteIpPool(context.Background(), givenPoolId).Execute()

	assert.Nil(t, err)
}

func TestShouldAssignIpToPool(t *testing.T) {
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenIpId := "DB3F9D439088BF73F5560443C8054AC4"

	givenRequest := email.IpPoolAssignIpRequest{
		IpId: givenIpId,
	}

	expectedRequest := fmt.Sprintf(`{
  "ipId": "%s"
 }`, givenIpId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf(EmailIpManagementPoolIpsEndpoint, givenPoolId), "", 204)

	actualRequest, _ := json.Marshal(givenRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	_, err := infobipClient.EmailAPI.AssignIpToPool(context.Background(), givenPoolId).IpPoolAssignIpRequest(givenRequest).Execute()

	assert.Nil(t, err)
}

func TestShouldRemoveIpFromPool(t *testing.T) {
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenIpId := "DB3F9D439088BF73F5560443C8054AC4"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf(EmailIpManagementPoolIpByIdEndpoint, givenPoolId, givenIpId), "", 204)

	_, err := infobipClient.EmailAPI.RemoveIpFromPool(context.Background(), givenPoolId, givenIpId).Execute()

	assert.Nil(t, err)
}

func TestShouldGetIpManagementDomain(t *testing.T) {
	givenDomainId := int64(1)
	givenDomainName := "domain.com"
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenPoolName := "IP pool name"
	givenPriority := int32(0)
	givenIpId := "DB3F9D439088BF73F5560443C8054AC4"
	givenIp := "198.51.100.0"

	givenResponse := fmt.Sprintf(`{
  "id": %d,
  "name": "%s",
  "pools": [{
    "id": "%s",
    "name": "%s",
    "priority": %d,
    "ips": [
      {
        "id": "%s",
        "ip": "%s"
      }
    ]
  }]
 }`, givenDomainId, givenDomainName, givenPoolId, givenPoolName, givenPriority, givenIpId, givenIp)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(EmailIpManagementDomainByIdEndpoint, givenDomainId), givenResponse, 200)

	response, _, err := infobipClient.EmailAPI.GetIpDomain(context.Background(), givenDomainId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenDomainId, response.GetId())
	assert.Equal(t, givenDomainName, response.GetName())
	actualPools := response.GetPools()[0]
	assert.Equal(t, givenPoolId, actualPools.GetId())
	assert.Equal(t, givenPoolName, actualPools.GetName())
	assert.Equal(t, givenPriority, actualPools.GetPriority())
	assert.Equal(t, 1, len(actualPools.GetIps()))
	assert.Equal(t, givenIpId, actualPools.GetIps()[0].GetId())
	assert.Equal(t, givenIp, actualPools.GetIps()[0].GetIp())
}

func TestShouldAssignPoolToDomain(t *testing.T) {
	givenDomainId := int64(1)
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenPriority := int32(0)

	givenRequest := email.DomainIpPoolAssignRequest{
		PoolId:   givenPoolId,
		Priority: givenPriority,
	}

	expectedRequest := fmt.Sprintf(`{
  "poolId": "%s",
  "priority": %d
 }`, givenPoolId, givenPriority)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf(EmailIpManagementDomainPoolsEndpoint, givenDomainId), "", 204)

	actualRequest, _ := json.Marshal(givenRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	_, err := infobipClient.EmailAPI.AssignPoolToDomain(context.Background(), givenDomainId).DomainIpPoolAssignRequest(givenRequest).Execute()

	assert.Nil(t, err)
}

func TestShouldUpdatePoolPriorityInDomain(t *testing.T) {
	givenDomainId := int64(1)
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenPriority := int32(0)

	givenRequest := email.DomainIpPoolUpdateRequest{
		Priority: givenPriority,
	}

	expectedRequest := fmt.Sprintf(`{
  "priority": %d
 }`, givenPriority)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf(EmailIpManagementDomainPoolByIdEndpoint, givenDomainId, givenPoolId), "", 204)

	actualRequest, _ := json.Marshal(givenRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	_, err := infobipClient.EmailAPI.UpdateDomainPoolPriority(context.Background(), givenDomainId, givenPoolId).DomainIpPoolUpdateRequest(givenRequest).Execute()

	assert.Nil(t, err)
}

func TestShouldRemovePoolFromDomain(t *testing.T) {
	givenDomainId := int64(1)
	givenPoolId := "08A3A7608750CC6E6080325A6ADF45B6"
	givenStatusCode := 204

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf(EmailIpManagementDomainPoolByIdEndpoint, givenDomainId, givenPoolId), "", givenStatusCode)

	_, err := infobipClient.EmailAPI.RemoveIpPoolFromDomain(context.Background(), givenDomainId, givenPoolId).Execute()

	assert.Nil(t, err)
}

// ==============================================================================
// Webhook deserialization tests
// ==============================================================================

func TestReceiveEmailDeliveryReportsWebhook(t *testing.T) {
	givenBulkId := "bulk-email-123"
	givenMessageId := "msg-email-456"
	givenTo := "user@example.com"
	givenSentAt := "2025-02-06T15:35:12.000+0000"
	givenDoneAt := "2025-02-06T15:35:14.000+0000"

	givenPayload := fmt.Sprintf(`
		{
			"results": [
				{
					"bulkId": "%s",
					"messageId": "%s",
					"to": "%s",
					"sentAt": "%s",
					"doneAt": "%s",
					"smsCount": 0,
					"price": {
						"pricePerMessage": 0.0,
						"currency": "EUR"
					},
					"status": {
						"groupId": 3,
						"groupName": "DELIVERED",
						"id": 5,
						"name": "DELIVERED_TO_HANDSET",
						"description": "Message delivered to handset"
					},
					"error": {
						"groupId": 0,
						"groupName": "OK",
						"id": 0,
						"name": "NO_ERROR",
						"description": "No error",
						"permanent": false
					}
				}
			]
		}`,
		givenBulkId,
		givenMessageId,
		givenTo,
		givenSentAt,
		givenDoneAt,
	)

	var responseBody email.DLRReportResponse
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, responseBody, "Expected non-nil response body")
	assert.Equal(t, 1, len(responseBody.GetResults()))

	result := responseBody.GetResults()[0]
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenTo, result.GetTo())

	status := result.GetStatus()
	assert.Equal(t, int32(3), status.GetGroupId())
	assert.Equal(t, "DELIVERED", status.GetGroupName())
	assert.Equal(t, int32(5), status.GetId())
	assert.Equal(t, "DELIVERED_TO_HANDSET", status.GetName())

	errorInfo := result.GetError()
	assert.Equal(t, int32(0), errorInfo.GetGroupId())
	assert.Equal(t, "NO_ERROR", errorInfo.GetName())
	assert.Equal(t, false, errorInfo.GetPermanent())
}

func TestReceiveEmailTrackingReportsWebhook(t *testing.T) {
	givenNotificationType := "OPENED"
	givenEventId := "evt-123"
	givenDomain := "example.com"
	givenRecipient := "user@example.com"
	givenBulkId := "bulk-email-123"
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"
	givenCampaignReferenceId := "campaignRef"

	givenPayload := fmt.Sprintf(`
		{
			"notificationType": "%s",
			"eventId": "%s",
			"domain": "%s",
			"recipient": "%s",
			"sendDateTime": 1707234912000,
			"messageId": 12345,
			"bulkId": "%s",
			"entityId": "%s",
			"applicationId": "%s",
			"campaignReferenceId": "%s"
		}`,
		givenNotificationType,
		givenEventId,
		givenDomain,
		givenRecipient,
		givenBulkId,
		givenEntityId,
		givenApplicationId,
		givenCampaignReferenceId,
	)

	var responseBody email.TrackReport
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenNotificationType, responseBody.GetNotificationType())
	assert.Equal(t, givenEventId, responseBody.GetEventId())
	assert.Equal(t, givenDomain, responseBody.GetDomain())
	assert.Equal(t, givenRecipient, responseBody.GetRecipient())
	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.Equal(t, givenEntityId, responseBody.GetEntityId())
	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenCampaignReferenceId, responseBody.GetCampaignReferenceId())
}

func TestReceiveEmailTrackingReportsWebhook_Clicked(t *testing.T) {
	givenPayload := `
		{
			"notificationType": "CLICKED",
			"eventId": "evt-456",
			"domain": "example.com",
			"recipient": "user@example.com",
			"url": "https://example.com/click-link",
			"sendDateTime": 1707234912000,
			"messageId": 12345,
			"bulkId": "bulk-email-123",
			"recipientInfo": {
				"deviceType": "Desktop",
				"os": "Windows",
				"deviceName": "Chrome"
			},
			"geoLocation": {
				"countryName": "Germany",
				"city": "Berlin"
			}
		}`

	var responseBody email.TrackReport
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, "CLICKED", responseBody.GetNotificationType())
	assert.Equal(t, "https://example.com/click-link", responseBody.GetUrl())

	recipientInfo := responseBody.GetRecipientInfo()
	assert.Equal(t, "Desktop", recipientInfo.GetDeviceType())
	assert.Equal(t, "Windows", recipientInfo.GetOs())

	geoLocation := responseBody.GetGeoLocation()
	assert.Equal(t, "Germany", geoLocation.GetCountryName())
	assert.Equal(t, "Berlin", geoLocation.GetCity())
}
