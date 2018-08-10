package infobip

import (
	"fmt"
)

type ApiError struct {
	StatusCode   int
	RequestError RequestError `json:"requestError"`
}

func (err ApiError) Error() string {
	return fmt.Sprintf("Error calling Infobip API\nstatus: %d\nresponse: %+v", err.StatusCode, err.RequestError)
}

type RequestError struct {
	ServiceException ServiceException `json:"serviceException"`
}

type ServiceException struct {
	MessageID string `json:"messageId"`
	Text      string `json:"text"`
}
