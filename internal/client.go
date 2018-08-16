package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobip/infobip-api-go-client/pkg"
)

type Querier interface {
	Query() url.Values
}

type QuerierFunc func() url.Values

func (q QuerierFunc) Query() url.Values {
	return q()
}

var (
	defaultHTTPClient = http.DefaultClient
	defaultBaseURL    = "https://api.infobip.com"
)

type Client struct {
	BaseURL    string
	Authorizer infobip.Authorizer
	HTTPClient *http.Client
}

func (c Client) getClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return defaultHTTPClient
}

func (c Client) getBaseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}
	return defaultBaseURL
}

func CallAPI(client *Client, method string, path string, pathParams map[string]string, query Querier, reqBody interface{}, resBodyPtr interface{}) error {
	url, err := composeURL(client.getBaseURL(), toPath(path, pathParams), toQuery(query))
	if err != nil {
		return err
	}

	body, err := toBodyReader(reqBody)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return err
	}
	fillHeaders(&httpReq.Header, client.Authorizer)

	httpRes, err := client.getClient().Do(httpReq)
	if err != nil {
		return err
	}

	httpResBody, err := handleResponse(httpRes)
	if err != nil {
		return err
	}

	if len(httpResBody) > 0 {
		err = json.Unmarshal(httpResBody, resBodyPtr)
		if err != nil {
			return err
		}
	}

	return nil
}

func composeURL(baseURL string, path string, query string) (*url.URL, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	full, err := base.Parse(path + query)
	return full, err
}

func toPath(path string, pathParams map[string]string) string {
	if pathParams != nil {
		changes := make([]string, 0, len(pathParams)*2)
		for key, val := range pathParams {
			changes = append(changes, key)
			changes = append(changes, val)
		}
		repl := strings.NewReplacer(changes...)
		return repl.Replace(path)
	}
	return path
}

func toQuery(querier Querier) string {
	if querier == nil {
		return ""
	}

	var query string
	for key, val := range querier.Query() {
		joined := strings.Join(val, ",")
		escaped := url.QueryEscape(joined)
		query += fmt.Sprintf("&%s=%s", key, escaped)
	}
	if query != "" {
		query = "?" + query[1:]
	}
	return query
}

func toBodyReader(reqBody interface{}) (io.Reader, error) {
	if reqBody == nil {
		return nil, nil
	}
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(reqBody)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func fillHeaders(header *http.Header, creds infobip.Authorizer) {
	header.Add("Authorization", creds.Authorization())
	header.Add("Content-Type", "application/json")
	header.Add("Accept", "application/json")
	header.Add("User-Agent", "messageFlow:APIv5_GO_CLIENT")
}

func handleResponse(res *http.Response) ([]byte, error) {
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if !success(res.StatusCode) {
		var apiErr infobip.ApiError
		apiErr.StatusCode = res.StatusCode
		_ = json.Unmarshal(resBody, &apiErr)
		return nil, apiErr
	}

	return resBody, nil
}

func success(statusCode int) bool {
	return http.StatusOK <= statusCode && statusCode < http.StatusBadRequest
}
