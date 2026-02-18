# Messages API quickstart

This quick guide aims to help you start with [Infobip Messages API](https://www.infobip.com/docs/api/platform/messages-api/sending-message/send-messages-api-message). After reading it, you should know how to use Messages API, send various types of messages, receive incoming messages, and receive delivery reports.

Messages API supports 10 different channels: SMS, MMS, RCS, WhatsApp, Viber Business Messages, Viber Bots, Apple Messages for Business, Instagram Direct Messages, Messenger, and LINE Official Notification.

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

## Activate your test senders

Before sending a message using Messages API, you need to activate your sender(s) and connect to our test domain.

Here you can find the example on how to activate and use **WhatsApp and SMS channels**.

To activate the WhatsApp test sender, add the **447860099299** Infobip sender to your WhatsApp contacts and send a message containing your Infobip account username.

To use the SMS test sender, simply send a message by using **InfoSMS** sender.

You are now ready to send your first message.

**IMPORTANT NOTE:** Keep in mind that for test purposes you can only send messages to a number you registered when you created your Infobip account.

## Send your first message

The easiest way to start with Messages API is to send a text message. First you need to prepare the message you want to send, like on snippet below:

````go
channel := messagesapi.OUTBOUNDMESSAGECHANNEL_SMS
sender := "<SENDER>"
content := messagesapi.MessageContent{
	Body: messagesapi.MessageBody{
		MessageTextBody: &messagesapi.MessageTextBody{
			Text: "Congratulations on sending your first message with GO library.",
		},
	},
}
destinations := []messagesapi.MessageDestination{
	{
		ToDestination: messagesapi.NewToDestination("<DESTINATION>"),
	},
}

givenMessage := messagesapi.NewMessage(
	channel,
	sender,
	destinations,
	content,
)

request := messagesapi.NewRequest([]messagesapi.RequestMessagesInner{
	messagesapi.MessageAsRequestMessagesInner(givenMessage),
})
````

Send the message invoking the appropriate send method and store the results in a new variable.

````go
apiResponse, httpResponse, err := infobipClient.
	MessagesAPI.
	SendMessagesApiMessage(auth).
	Request(*request).
	Execute()
````

Once the invocation finishes, you can inspect the results and print a status description, as shown below.

````go
fmt.Printf("Response: %+v\n", apiResponse)
fmt.Printf("HTTP Response Details: %+v\n", httpResponse)
````

## How to receive messages

To receive messages using Messages API you must set up the webhook.

Basically, that is just an endpoint implemented on your side where you will accept the requests when a new message arrives. That endpoint will be called by the Infobip API whenever we receive an incoming message for your registered sender(s).

```go
    func receiveIncomingMessagesApi(w http.ResponseWriter, req *http.Request) {
    	reqBody, _ := io.ReadAll(req.Body)
        var result messagesapi.IncomingMessage
        err := json.Unmarshal(reqBody, &result)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
    
        msgs := result.GetResults()
        for _, msg := range msgs {
    		if msg.MoEvent != nil {
    			moEvent := msg.MoEvent
            	fmt.Println(fmt.Sprintf("%s - %s - %s", moEvent.GetSender(), moEvent.GetChannel(), moEvent.Event))
    		}
    		if msg.InboundTypingStartedEvent != nil {
    			inboundTypingStartedEvent := msg.InboundTypingStartedEvent
    			fmt.Println(fmt.Sprintf("%s - %s - %s", inboundTypingStartedEvent.GetSender(), inboundTypingStartedEvent.GetChannel(), inboundTypingStartedEvent.Event))
    		}
    		if msg.InboundTypingStoppedEvent != nil {
    			inboundTypingStoppedEvent := msg.InboundTypingStoppedEvent
    			fmt.Println(fmt.Sprintf("%s - %s - %s", inboundTypingStoppedEvent.GetSender(), inboundTypingStoppedEvent.GetChannel(), inboundTypingStoppedEvent.Event))
    		}
        }
    }
    
    func main() {
        http.HandleFunc("/incoming-messages", receiveIncomingMessagesApi)
    }
```

You can find more details about the structure of the message you can expect on your endpoint on [docs page](https://www.infobip.com/docs/api/platform/messages-api/incoming-message/receive-messages-api-incoming-messages).

## How to receive delivery reports

For each message that you send out, you can get a message delivery report in real time. Subscribe for reports by contacting our support team at <support@infobip.com>. e.g. `https://{yourDomain}/delivery-reports`

```go
    func receiveDeliveryReports(w http.ResponseWriter, req *http.Request) {
    	reqBody, _ := io.ReadAll(req.Body)
        var result messagesapi.DeliveryReport
        err := json.Unmarshal(reqBody, &result)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
    
    	dlrs := result.GetResults()
    	for _, dlr := range dlrs {
    		fmt.Println(fmt.Sprintf("%s - %s", dlr.MessageId, dlr.GetStatus().Name))
    	}
    }
    
    func main() {
    	http.HandleFunc("/delivery-reports", receiveDeliveryReports)
    }
```

## Use adaptationMode to automatically modify message types

Enhance your Messages API requests by using the `adaptationMode` parameter. It allows you to send messages even if they are unsupported by the channel.

When you set adaptationMode to `true,` Messages API automatically adjusts the message to remove any unsupported elements, ensuring successful delivery.

For instance, if you'd like to include an image in your WhatsApp and SMS messages, set adaptationMode to 'true'. Messages API will handle the delivery for WhatsApp as a message containing an image, while for SMS will provide a link to the image.

On the other hand, if you set adaptationMode to 'false' and try to send a message with an unsupported element to a channel, an error will occur. Make sure to choose the right setting based on your specific channel and content requirements.
