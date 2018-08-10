package sms

//This is a generated file and is not intended for modification!

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/infobip/infobip-api-go-client/pkg"
)

var (
	//reference importted packages just in case auto-generated code doesn't
	_ = errors.New
	_ = url.Parse
	_ = strconv.Itoa
	_ = strings.Replace
	_ = infobip.NewApiKeyCredentials
)

type SMSTextualRequest struct {
	From string   `json:"from,omitempty"`
	To   []string `json:"to,omitempty"`
	Text string   `json:"text,omitempty"`
}

type SMSAdvancedTextualRequest struct {
	Tracking *Tracking `json:"tracking,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	BulkID   string    `json:"bulkId,omitempty"`
}

type Tracking struct {
	Track      string `json:"track,omitempty"`
	ProcessKey string `json:"processKey,omitempty"`
	Type       string `json:"type,omitempty"`
	BaseURL    string `json:"baseUrl,omitempty"`
}

type Message struct {
	From               string              `json:"from,omitempty"`
	To                 []string            `json:"to,omitempty"`
	Destinations       []Destination       `json:"destinations,omitempty"`
	Text               string              `json:"text,omitempty"`
	Binary             *BinaryContent      `json:"binary,omitempty"`
	Flash              bool                `json:"flash,omitempty"`
	Language           *Language           `json:"language,omitempty"`
	Transliteration    string              `json:"transliteration,omitempty"`
	Notify             bool                `json:"notify,omitempty"`
	IntermediateReport bool                `json:"intermediateReport,omitempty"`
	NotifyURL          string              `json:"notifyUrl,omitempty"`
	NotifyContentType  string              `json:"notifyContentType,omitempty"`
	CallbackData       string              `json:"callbackData,omitempty"`
	ValidityPeriod     int64               `json:"validityPeriod,omitempty"`
	SendAt             *infobip.Time       `json:"sendAt,omitempty"`
	DeliveryTimeWindow *DeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	CampaignID         string              `json:"campaignId,omitempty"`
	OperatorClientID   string              `json:"operatorClientId,omitempty"`
}

type Destination struct {
	To        string `json:"to,omitempty"`
	MessageID string `json:"messageId,omitempty"`
}

type BinaryContent struct {
	Hex        string `json:"hex,omitempty"`
	DataCoding int    `json:"dataCoding,omitempty"`
	EsmClass   int    `json:"esmClass,omitempty"`
}

type Language struct {
	SingleShift  bool   `json:"singleShift,omitempty"`
	LockingShift bool   `json:"lockingShift,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
}

type DeliveryTimeWindow struct {
	From *DeliveryTime `json:"from,omitempty"`
	To   *DeliveryTime `json:"to,omitempty"`
	Days []DeliveryDay `json:"days,omitempty"`
}

type DeliveryTime struct {
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
}

type DeliveryDay int

const (
	MONDAY DeliveryDay = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

var DeliveryDayValues = []DeliveryDay{
	MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY,
}

func (day DeliveryDay) String() string {
	names := []string{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
	if day < MONDAY || day > SUNDAY {
		return "Unknown"
	}
	return names[day]
}

func (dd DeliveryDay) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dd.String() + `"`), nil
}

func (dd *DeliveryDay) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	switch s {
	case "MONDAY":
		*dd = MONDAY
	case "TUESDAY":
		*dd = TUESDAY
	case "WEDNESDAY":
		*dd = WEDNESDAY
	case "THURSDAY":
		*dd = THURSDAY
	case "FRIDAY":
		*dd = FRIDAY
	case "SATURDAY":
		*dd = SATURDAY
	case "SUNDAY":
		*dd = SUNDAY
	default:
		return errors.New("Unknown DeliveryDay value " + s)
	}
	return nil
}

type SMSResponse struct {
	BulkID             string               `json:"bulkId,omitempty"`
	TrackingProcessKey string               `json:"trackingProcessKey,omitempty"`
	Messages           []SMSResponseDetails `json:"messages,omitempty"`
}

type SMSResponseDetails struct {
	To        string  `json:"to,omitempty"`
	SmsCount  int     `json:"smsCount,omitempty"`
	MessageID string  `json:"messageId,omitempty"`
	Status    *Status `json:"status,omitempty"`
}

type Status struct {
	GroupID     int    `json:"groupId,omitempty"`
	GroupName   string `json:"groupName,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Action      string `json:"action,omitempty"`
}

type SMSReportResponse struct {
	Results []SMSReport `json:"results,omitempty"`
}

type SMSReport struct {
	BulkID       string       `json:"bulkId,omitempty"`
	MessageID    string       `json:"messageId,omitempty"`
	To           string       `json:"to,omitempty"`
	From         string       `json:"from,omitempty"`
	Text         string       `json:"text,omitempty"`
	SentAt       infobip.Time `json:"sentAt,omitempty"`
	DoneAt       infobip.Time `json:"doneAt,omitempty"`
	SmsCount     int          `json:"smsCount,omitempty"`
	MccMnc       string       `json:"mccMnc,omitempty"`
	Price        *Price       `json:"price,omitempty"`
	Status       *Status      `json:"status,omitempty"`
	Error        *Error       `json:"error,omitempty"`
	CallbackData string       `json:"callbackData,omitempty"`
}

type Error struct {
	GroupID     int    `json:"groupId,omitempty"`
	GroupName   string `json:"groupName,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Permanent   bool   `json:"permanent,omitempty"`
}

type Price struct {
	PricePerMessage float64 `json:"pricePerMessage,omitempty"`
	PricePerLookup  float64 `json:"pricePerLookup,omitempty"`
	Currency        string  `json:"currency,omitempty"`
}

type GetSentSmsDeliveryReportsQuery struct {
	BulkID    string
	MessageID string
	Limit     int
}

func (query GetSentSmsDeliveryReportsQuery) Query() url.Values {
	q := make(url.Values)
	if query.BulkID != "" {
		q.Set("bulkId", query.BulkID)
	}
	if query.MessageID != "" {
		q.Set("messageId", query.MessageID)
	}
	if query.Limit != 0 {
		q.Set("limit", strconv.Itoa(query.Limit))
	}
	return q
}
