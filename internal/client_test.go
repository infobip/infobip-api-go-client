package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/infobip/infobip-api-go-client/pkg"
)

type Model struct {
	Foo string
	Baz int
}

type RequestCaptor struct {
	Req  *http.Request
	Body []byte
}

func TestFormatingURL(t *testing.T) {
	var reqCaptor RequestCaptor
	ts := givenOKServer(`""`, &reqCaptor)
	defer ts.Close()
	c := givenClient(ts)
	p := map[string]string{
		"{param2}": "2",
		"{param1}": "1",
		"{unused}": "unused",
	}
	q := QuerierFunc(func() url.Values {
		v := make(url.Values)
		v.Set("foo", "b a r")
		v.Set("baz", "42")
		return v
	})
	var actualRes string

	err := CallAPI(&c, "GET", "/test/{param1}/{param2}/{param1}", p, q, nil, &actualRes)

	if err != nil {
		t.Fatalf("Expected to make a GET request, but was error: %+v", err)
	}
	expected, _ := url.Parse("/test/1/2/1?foo=b+a+r&baz=42")
	actual, _ := url.Parse(reqCaptor.Req.RequestURI)
	if expected.Path != actual.Path || !reflect.DeepEqual(expected.Query(), actual.Query()) {
		t.Errorf("Expected request to url `%s`, but was `%s`", expected, reqCaptor.Req.RequestURI)
	}
}

func TestAddingHeaders(t *testing.T) {
	var reqCaptor RequestCaptor
	ts := givenOKServer(`""`, &reqCaptor)
	defer ts.Close()
	c := givenClient(ts)
	var actualRes string

	err := CallAPI(&c, "GET", "/test", nil, nil, nil, &actualRes)

	if err != nil {
		t.Fatalf("Expected to make a GET request, but was error: %+v", err)
	}
	expected := map[string][]string{
		"Authorization":   []string{"Auth"},
		"Content-Type":    []string{"application/json"},
		"Accept":          []string{"application/json"},
		"User-Agent":      []string{"messageFlow:APIv5_GO_CLIENT"},
		"Accept-Encoding": []string{"gzip"},
	}
	actual := reqCaptor.Req.Header
	if !reflect.DeepEqual(actual, http.Header(expected)) {
		t.Errorf("Expected request headers to be %+v but were %+v", expected, actual)
	}
}

func TestMarshalingRequest(t *testing.T) {
	var reqCaptor RequestCaptor
	ts := givenOKServer(`""`, &reqCaptor)
	defer ts.Close()
	c := givenClient(ts)
	var givenReq = Model{
		Foo: "bar",
		Baz: 42,
	}
	var actualRes string

	err := CallAPI(&c, "POST", "/test", nil, nil, givenReq, &actualRes)

	if err != nil {
		t.Fatalf("Expected to make a POST request, but was error: %+v", err)
	}
	var actual Model
	json.Unmarshal(reqCaptor.Body, &actual)
	if !reflect.DeepEqual(actual, givenReq) {
		t.Errorf("Expected sent body to be %+v, but was %+v", givenReq, actual)
	}
}

func TestUnmarshalingResponse(t *testing.T) {
	givenRes := `{ "foo": "bar", "baz": 42}`
	ts := givenOKServer(givenRes, nil)
	defer ts.Close()
	c := givenClient(ts)
	var actualRes Model

	err := CallAPI(&c, "GET", "/test", nil, nil, nil, &actualRes)

	if err != nil {
		t.Fatalf("Expected to make a GET request, but was error: %s", err.Error())
	}
	if actualRes.Foo != "bar" || actualRes.Baz != 42 {
		t.Errorf("Expected to receive `%s`, but was %+v", givenRes, actualRes)
	}
}

func TestUnmarshalingAPIError(t *testing.T) {
	givenRes := `{
		"requestError": {
			"serviceException": {
				"messageId": "BAD_REQUEST",
				"text": "Missing foo parameter."
			}
		}
	}`
	ts := givenServer(givenRes, http.StatusBadRequest, nil)
	defer ts.Close()
	c := givenClient(ts)

	err := CallAPI(&c, "GET", "/test", nil, nil, nil, nil)

	if err == nil {
		t.Error("Expected error but there was none")
	}
	ibErr, ok := err.(infobip.ApiError)
	if !ok {
		t.Errorf("Expected infobip.ApiError, but was %+v", err)
	}
	if ibErr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusCode to be %d, but was %d", http.StatusBadRequest, ibErr.StatusCode)
	}
	exc := ibErr.RequestError.ServiceException
	if exc.MessageID != "BAD_REQUEST" || exc.Text != "Missing foo parameter." {
		t.Errorf("Expected error to represent `%s`, but was %+v", givenRes, ibErr)
	}
}

func givenOKServer(mockResBody string, reqCaptor *RequestCaptor) *httptest.Server {
	return givenServer(mockResBody, http.StatusOK, reqCaptor)
}

func givenServer(mockResBody string, mockStatusCode int, reqCaptor *RequestCaptor) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if reqCaptor != nil {
			reqCaptor.Req = r
			reqCaptor.Body, _ = ioutil.ReadAll(r.Body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(mockStatusCode)
		fmt.Fprintln(w, mockResBody)
	}))
}

func givenClient(ts *httptest.Server) Client {
	return Client{
		BaseURL: ts.URL,
		Authorizer: infobip.AuthorizerFunc(func() string {
			return "Auth"
		}),
		HTTPClient: http.DefaultClient,
	}
}
