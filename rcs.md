# RCS API quickstart

This example demonstrates how to use the [Infobip RCS API](https://www.infobip.com/docs/api/channels/rcs). You'll learn how to initialize an RCS client, send a message, and pull delivery reports and logs.

The first step is to add your configuration, initialize the api client and set your authentication:

````go
    configuration := infobip.NewConfiguration()
    configuration.Host = "<YOUR_BASE_URL>"

    infobipClient := api.NewAPIClient(configuration)

    auth := context.WithValue(
        context.Background(),
        infobip.ContextAPIKeys,
        map[string]infobip.APIKey{
            "APIKeyHeader": {Key: "<YOUR_API_KEY>"},
        },
    )
````

For details, check the [client](https://github.com/infobip/infobip-api-go-client/blob/master/v3/pkg/infobip/client.go) source code.

## Send an RCS text message

````go
    sender := "<RCS_SENDER_ID>"

    destinations := []rcs.ToDestination{
        *rcs.NewToDestination("<DESTINATION>"),
    }

    textContent := rcs.NewOutboundTextContent("Hello from Infobip RCS!")
    content := rcs.OutboundTextContentAsOutboundMessageContent(textContent)

    message := rcs.NewMessage(sender, destinations, content)

    request := rcs.NewMessageRequest([]rcs.Message{*message})

    apiResponse, httpResponse, err := infobipClient.
        RcsAPI.
        SendRcsMessages(auth).
        MessageRequest(*request).
        Execute()
````

`apiResponse` includes `BulkId` and per-message IDs you can use for reports and logs.

## Get delivery reports (pull)

If you can’t receive webhooks, fetch batches of delivery reports:

````go
    bulkId := apiResponse.GetBulkId()
    limit := int32(20)

    reports, httpResponse, err := infobipClient.
        RcsAPI.
        GetOutboundRcsMessageDeliveryReports(auth).
        BulkId(bulkId).
        Limit(limit).
        Execute()
````

Reports are returned once; store or process them before the next call.

## Get message logs

Retrieve recent logs (last 48h) for troubleshooting or UI display:

````go
    logs, httpResponse, err := infobipClient.
        RcsAPI.
        GetOutboundRcsMessageLogs(auth).
        Sender(sender).
        Destination("<DESTINATION>").
        Limit(20).
        Execute()
````

Use `UseCursor(true)` and the returned cursor for paging through larger result sets.
