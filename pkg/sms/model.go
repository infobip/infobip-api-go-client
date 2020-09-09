package sms

//This is a generated file and is not intended for modification!

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/infobip/infobip-api-go-client/pkg"
)

var (
	//reference imported packages just in case auto-generated code doesn't
	_ = errors.New
	_ = url.Parse
	_ = strings.Replace
	_ = fmt.Sprint
	_ = infobip.NewApiKeyCredentials
)

type Status struct {
	GroupID int32 `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	ID int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Action string `json:"action,omitempty"`
}

type SMSResponseDetails struct {
	To string `json:"to,omitempty"`
	Status *Status `json:"status,omitempty"`
	SMSCount int32 `json:"smsCount,omitempty"`
	MessageID string `json:"messageId,omitempty"`
}

type SMSAdvancedTextualRequest struct {
	Tracking *Tracking `json:"tracking,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	BulkID string `json:"bulkId,omitempty"`
}

type DeliveryTimeWindow struct {
	From *DeliveryTime `json:"from,omitempty"`
	To *DeliveryTime `json:"to,omitempty"`
	Days []DeliveryDay `json:"days,omitempty"`
}

type Destination struct {
	To string `json:"to,omitempty"`
	MessageID string `json:"messageId,omitempty"`
}

type Message struct {
	From string `json:"from,omitempty"`
	To []string `json:"to,omitempty"`
	Destinations []Destination `json:"destinations,omitempty"`
	Text string `json:"text,omitempty"`
	Binary *BinaryContent `json:"binary,omitempty"`
	Flash bool `json:"flash,omitempty"`
	Language *Language `json:"language,omitempty"`
	Transliteration string `json:"transliteration,omitempty"`
	Notify bool `json:"notify,omitempty"`
	IntermediateReport bool `json:"intermediateReport,omitempty"`
	NotifyURL string `json:"notifyUrl,omitempty"`
	NotifyContentType string `json:"notifyContentType,omitempty"`
	CallbackData string `json:"callbackData,omitempty"`
	ValidityPeriod int64 `json:"validityPeriod,omitempty"`
	SendAt *infobip.Time `json:"sendAt,omitempty"`
	DeliveryTimeWindow *DeliveryTimeWindow `json:"deliveryTimeWindow,omitempty"`
	CampaignID string `json:"campaignId,omitempty"`
	OperatorClientID string `json:"operatorClientId,omitempty"`
	Regional *RegionalOptions `json:"regional,omitempty"`
}

type Language struct {
	SingleShift bool `json:"singleShift,omitempty"`
	LockingShift bool `json:"lockingShift,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
}

type RegionalOptions struct {
	IndiaDLT *IndiaDLTOptions `json:"indiaDlt,omitempty"`
}

type IndiaDLTOptions struct {
	ContentTemplateID string `json:"contentTemplateId,omitempty"`
	PrincipalEntityID string `json:"principalEntityId,omitempty"`
}

type Error struct {
	GroupID int32 `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	ID int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Permanent bool `json:"permanent,omitempty"`
}

type GetSentSMSDeliveryReportsQuery struct {
	BulkID string
	MessageID string
	Limit int32
}

func (query GetSentSMSDeliveryReportsQuery) Query() url.Values {
	q := make(url.Values)
	if query.BulkID != "" {
		q.Set("bulkId", query.BulkID)
	}
	if query.MessageID != "" {
		q.Set("messageId", query.MessageID)
	}
	if query.Limit != 0 {
		q.Set("limit", fmt.Sprint(query.Limit))
	}
	return q
}

type DeliveryDay int

const (
	DD_MONDAY DeliveryDay = iota
	DD_TUESDAY
	DD_WEDNESDAY
	DD_THURSDAY
	DD_FRIDAY
	DD_SATURDAY
	DD_SUNDAY
)

var DeliveryDayValues = []DeliveryDay{
	DD_MONDAY, DD_TUESDAY, DD_WEDNESDAY, DD_THURSDAY, DD_FRIDAY, DD_SATURDAY, DD_SUNDAY,
}

func (val DeliveryDay) String() string {
	names := []string{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
	if val < DD_MONDAY || val > DD_SUNDAY {
		return "Unknown"
	}
	return names[val]
}

func (val DeliveryDay) MarshalJSON() ([]byte, error) {
	return []byte(`"` + val.String() + `"`), nil
}

func (val *DeliveryDay) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	switch s {
	case "MONDAY":
		*val = DD_MONDAY
	case "TUESDAY":
		*val = DD_TUESDAY
	case "WEDNESDAY":
		*val = DD_WEDNESDAY
	case "THURSDAY":
		*val = DD_THURSDAY
	case "FRIDAY":
		*val = DD_FRIDAY
	case "SATURDAY":
		*val = DD_SATURDAY
	case "SUNDAY":
		*val = DD_SUNDAY
	default:
		return errors.New("Unknown DeliveryDay value " + s)
	}
	return nil
}

type BinaryContent struct {
	Hex string `json:"hex,omitempty"`
	DataCoding int32 `json:"dataCoding,omitempty"`
	EsmClass int32 `json:"esmClass,omitempty"`
}

type SMSReportResponse struct {
	Results []SMSReport `json:"results,omitempty"`
}

type Tracking struct {
	Track string `json:"track,omitempty"`
	ProcessKey string `json:"processKey,omitempty"`
	Type string `json:"type,omitempty"`
	BaseURL string `json:"baseUrl,omitempty"`
}

type SMSResponse struct {
	BulkID string `json:"bulkId,omitempty"`
	TrackingProcessKey string `json:"trackingProcessKey,omitempty"`
	Messages []SMSResponseDetails `json:"messages,omitempty"`
}

type SMSReport struct {
	BulkID string `json:"bulkId,omitempty"`
	MessageID string `json:"messageId,omitempty"`
	To string `json:"to,omitempty"`
	From string `json:"from,omitempty"`
	Text string `json:"text,omitempty"`
	SentAt *infobip.Time `json:"sentAt,omitempty"`
	DoneAt *infobip.Time `json:"doneAt,omitempty"`
	SMSCount int32 `json:"smsCount,omitempty"`
	MccMnc string `json:"mccMnc,omitempty"`
	Price *Price `json:"price,omitempty"`
	Status *Status `json:"status,omitempty"`
	Error *Error `json:"error,omitempty"`
	CallbackData string `json:"callbackData,omitempty"`
}

type DeliveryTime struct {
	Hour int32 `json:"hour,omitempty"`
	Minute int32 `json:"minute,omitempty"`
}

type Price struct {
	PricePerMessage float64 `json:"pricePerMessage,omitempty"`
	PricePerLookup float64 `json:"pricePerLookup,omitempty"`
	Currency string `json:"currency,omitempty"`
}

