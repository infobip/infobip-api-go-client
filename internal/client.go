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

type Client struct {
	BaseURL    string
	Authorizer infobip.Authorizer
	HTTPCLient *http.Client
}

func CallAPI(client *Client, method string, path string, query Querier, reqBody interface{}, resBodyPtr interface{}) error {
	url, err := composeURL(client.BaseURL, path, toQuery(query))
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

	httpRes, err := client.HTTPCLient.Do(httpReq)
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
