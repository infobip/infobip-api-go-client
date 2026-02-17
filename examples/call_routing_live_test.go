package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
)

// Minimal route creation that forwards to a single SIP endpoint.
func TestCreateCallRouteSimple(t *testing.T) {
	client, auth := newClientAndAuth()

	sipEndpoint := voice.NewCallRoutingSipEndpoint("<SIP_TRUNK_ID>")
	sipEndpoint.SetUsername("<SIP_USERNAME>")
	destination := voice.NewCallRoutingEndpointDestination(
		voice.CallRoutingSipEndpointAsCallRoutingEndpoint(sipEndpoint),
	)

	routeReq := voice.NewCallRoutingRouteRequest("Support route", []voice.CallRoutingDestination{
		voice.CallRoutingEndpointDestinationAsCallRoutingDestination(destination),
	})

	apiResponse, httpResponse, err := client.
		CallRoutingAPI.
		CreateCallRoute(auth).
		CallRoutingRouteRequest(*routeReq).
		Execute()

	if err != nil {
		t.Fatalf("Failed to create call route: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Id == nil {
		t.Fatalf("Invalid response: expected route id, got: %+v", apiResponse)
	}
}

// Simulate route selection with criteria, weighted destinations, and recording options.
func TestSimulateRouteSelectionAdvanced(t *testing.T) {
	client, auth := newClientAndAuth()

	primary := voice.NewCallRoutingEndpointDestination(
		voice.CallRoutingPhoneEndpointAsCallRoutingEndpoint(
			voice.NewCallRoutingPhoneEndpoint(),
		),
	)
	primary.SetPriority(1)
	primary.SetConnectTimeout(25)
	recording := voice.NewCallRoutingRecording(voice.CALLROUTINGRECORDINGTYPE_AUDIO)
	recordingComposition := voice.NewCallRoutingRecordingComposition(true)
	recording.SetRecordingComposition(*recordingComposition)
	primary.SetRecording(*recording)

	failover := voice.NewCallRoutingEndpointDestination(
		voice.CallRoutingSipEndpointAsCallRoutingEndpoint(
			voice.NewCallRoutingSipEndpoint("<SIP_TRUNK_FAILOVER>"),
		),
	)
	failover.SetPriority(2)
	failover.SetWeight(50)

	simEndpoint := voice.CallRoutingSimulatorPhoneEndpointAsCallRoutingSimulatorCallEndpoint(
		voice.NewCallRoutingSimulatorPhoneEndpoint(),
	)

	simRequest := voice.NewCallRoutingRouteSimulatorRequest(simEndpoint)
	simRequest.SetTo("<TARGET_NUMBER>")
	simRequest.SetUseDisabledRoutes(false)

	apiResponse, httpResponse, err := client.CallRoutingAPI.
		SimulateRouteSelection(auth).
		CallRoutingRouteSimulatorRequest(*simRequest).
		Execute()

	if err != nil {
		t.Fatalf("Failed to simulate route selection: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Id == nil {
		t.Fatalf("Invalid response: expected simulated route id, got: %+v", apiResponse)
	}
}

func TestCallRoutingFullBundledFlow(t *testing.T) {
	client, auth := newClientAndAuth()

	// --- Create -------------------------------------------------------------
	sipEndpoint := voice.NewCallRoutingSipEndpoint("<SIP_TRUNK_ID>")
	sipEndpoint.SetUsername("<SIP_USERNAME>")
	destination := voice.NewCallRoutingEndpointDestination(
		voice.CallRoutingSipEndpointAsCallRoutingEndpoint(sipEndpoint),
	)
	destination.SetConnectTimeout(20)

	createReq := voice.NewCallRoutingRouteRequest("E2E Call Route", []voice.CallRoutingDestination{
		voice.CallRoutingEndpointDestinationAsCallRoutingDestination(destination),
	})

	createResp, httpResp, err := client.
		CallRoutingAPI.
		CreateCallRoute(auth).
		CallRoutingRouteRequest(*createReq).
		Execute()
	if err != nil {
		t.Fatalf("Create route failed: %v (http=%+v)", err, httpResp)
	}
	routeID := createResp.GetId()

	// --- Get single --------------------------------------------------------
	if _, httpResp, err := client.
		CallRoutingAPI.
		GetCallRoute(auth, routeID).
		Execute(); err != nil {
		t.Fatalf("Get route failed: %v (http=%+v)", err, httpResp)
	}

	// --- List --------------------------------------------------------------
	page, httpResp, err := client.
		CallRoutingAPI.
		GetCallRoutes(auth).
		Execute()
	if err != nil || page == nil || len(page.GetResults()) == 0 {
		t.Fatalf("List routes failed: %v (http=%+v)", err, httpResp)
	}

	// --- Update ------------------------------------------------------------
	updatedSip := voice.NewCallRoutingSipEndpoint("<SIP_TRUNK_ID>")
	updatedSip.SetUsername("<SIP_USERNAME>")
	updatedDestination := voice.NewCallRoutingEndpointDestination(
		voice.CallRoutingSipEndpointAsCallRoutingEndpoint(updatedSip),
	)
	updatedDestination.SetConnectTimeout(30)
	updateReq := voice.NewCallRoutingRouteRequest("E2E Call Route Updated", []voice.CallRoutingDestination{
		voice.CallRoutingEndpointDestinationAsCallRoutingDestination(updatedDestination),
	})

	updateResp, httpResp, err := client.
		CallRoutingAPI.
		UpdateCallRoute(auth, routeID).
		CallRoutingRouteRequest(*updateReq).
		Execute()
	if err != nil {
		t.Fatalf("Update route failed: %v (http=%+v)", err, httpResp)
	}

	// --- Delete ------------------------------------------------------------
	if _, httpResp, err := client.
		CallRoutingAPI.
		DeleteCallRoute(auth, routeID).
		Execute(); err != nil {
		t.Fatalf("Delete route failed: %v (http=%+v)", err, httpResp)
	}

	if routeID == "" || updateResp.GetName() == "" {
		t.Fatalf("Expected non-empty IDs/names (routeID=%s updatedName=%s)", routeID, updateResp.GetName())
	}

	t.Logf("RouteID=%s totalListed=%d updatedName=%s", routeID, len(page.GetResults()), updateResp.GetName())
}
