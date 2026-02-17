# Moments quickstart

This quick guide aims to help you start with [Infobip Moments API](https://www.infobip.com/docs/api/customer-engagement/moments). After reading it, you should know how to use Moments.

The first step is to add your configuration, initialize the api client and set your authentication:

````go
    configuration := infobip.NewConfiguration()
    configuration.Host = "<YOUR_BASE_URL>"
    
    infobipClient := api.NewAPIClient(configuration)
    
    auth := context.WithValue(
        context.Background(),
        infobip.ContextAPIKeys,
        map[string]infobip.APIKey{
            // Prefix is deprecated; the SDK now prepends "App " automatically.
            "APIKeyHeader": {Key: "<YOUR_API_KEY>"},
        },
    )
````

For details, check the [client](https://github.com/infobip/infobip-api-go-client/blob/master/v3/pkg/infobip/client.go) source code.


## Flow API

### Add participants to flow

To add participants to a flow, you can use the following code:

````go
    campaignId := int64(200000000000001)
    identifier := "test@example.com"
	participantType := moments.FLOWPERSONUNIQUEFIELDTYPE_EMAIL
	notifyUrl := "https://example.com"

    participantUniqueField := *moments.NewFlowPersonUniqueField(identifier, participantType)
    participant := moments.NewFlowParticipant(participantUniqueField)

    request := moments.NewFlowAddFlowParticipantsRequest([]moments.FlowParticipant{*participant})
    request.SetNotifyUrl(notifyUrl)

    response, _, err := infobipClient.
        FlowAPI.
        AddFlowParticipants(auth, campaignId).
        FlowAddFlowParticipantsRequest(*request).
        Execute()
````

### Get a report on participants added to flow

To fetch a report to confirm that all persons have been successfully added to the flow, you can use the following code:

````go
    givenOperationId := "03f2d474-0508-46bf-9f3d-d8e2c28adaea"

    response, _, err := infobipClient.
        FlowAPI.
        GetFlowParticipantsAddedReport(auth, campaignId).
        OperationId(givenOperationId).
        Execute()
````

### Remove person from flow

To remove a person from a flow, you can use the following code:

````go
    externalId := "8edb24b5-0319-48cd-a1d9-1e8bc5d577ab"
    _, err := infobipClient.
        FlowAPI.
        RemovePeopleFromFlow(auth, campaignId).
        ExternalId(externalId).
        Execute()
````


## Forms API

### Get forms

To get all forms, you can use the following code:

````go
    response, _, err := infobipClient.
        FormsAPI.
        GetForms(auth).
        Execute()
````

### Get form by ID

To get a specific form by ID, you can use the following code:

````go
    formId := "cec5dfd2-4238-48e0-933b-9acbdb2e6f5f"

    response, _, err := infobipClient.
        FormsAPI.
        GetForm(auth, formId).
        Execute()
````

### Increment form view count

To increase the view counter of a specific form, you can use the following code:

````go
    response, _, err := infobipClient.
        FormsAPI.
        IncrementViewCount(auth, formId).
        Execute()
````

### Submit form data

To submit data to a specific form, you can use the following code:

````go
    formDataRequest := map[string]interface{}{
        "first_name": "John",
        "last_name": "Doe",
        "company": "Infobip",
        "email": "info@example.com",
    }

    response, _, err := infobipClient.
        FormsAPI.
        SubmitFormData(auth, formId).
        RequestBody(formDataRequest).
        Execute()
````
