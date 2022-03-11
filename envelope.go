package netsuite

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"env:Header"`
	Body   Body   `xml:"env:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		Header: NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"Header"`
	Body   Body   `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "env:Envelope"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
		{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
		{Name: xml.Name{Space: "", Local: "xmlns:platformMsgs"}, Value: "urn:messages_2021_2.platform.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:env"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		{Name: xml.Name{Space: "", Local: "xmlns:platformCore"}, Value: "urn:core_2021_2.platform.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:platformCommon"}, Value: "urn:common_2021_2.platform.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:listRel"}, Value: "urn:relationships_2021_2.lists.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranSales"}, Value: "urn:sales_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranPurch"}, Value: "urn:purchases_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:actSched"}, Value: "urn:scheduling_2021_2.activities.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:setupCustom"}, Value: "urn:customization_2021_2.setup.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:listAcct"}, Value: "urn:accounting_2021_2.lists.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranBank"}, Value: "urn:bank_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranCust"}, Value: "urn:customers_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranEmp"}, Value: "urn:employees_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranInvt"}, Value: "urn:inventory_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:listSupport"}, Value: "urn:support_2021_2.lists.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:tranGeneral"}, Value: "urn:general_2021_2.transactions.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:commGeneral"}, Value: "urn:communication_2021_2.general.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:listMkt"}, Value: "urn:marketing_2021_2.lists.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:listWebsite"}, Value: "urn:website_2021_2.lists.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:fileCabinet"}, Value: "urn:filecabinet_2021_2.documents.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:listEmp"}, Value: "urn:employees_2021_2.lists.webservices.netsuite.com"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type Body struct {
	ActionBody interface{} `xml:",any"`
}

type Header struct {
	TokenPassport TokenPassport `xml:"platformMsgs:tokenPassport"`
	// DocumentInfo  struct {
	// 	NSInfo string `xml:"platformMsgs:nsId,omitempty"`
	// } `xml:"platformMsgs:documentInfo,omitempty"`
	Preferences       Preferences       `xml:"platformMsgs:preferences"`
	SearchPreferences SearchPreferences `xml:"platformMsgs:searchPreferences"`
}

func (h Header) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(h, e, start)
}

func (h Header) IsEmpty() bool {
	return zero.IsZero(h)
}

func NewHeader() Header {
	return Header{}
}

type ActionBody interface{}

type TokenPassport struct {
	XMLName     xml.Name `xml:"platformMsgs:tokenPassport"`
	Account     string   `xml:"platformMsgs:account"`
	ConsumerKey string   `xml:"platformMsgs:consumerKey"`
	Token       string   `xml:"platformMsgs:token"`
	Nonce       string   `xml:"platformMsgs:nonce"`
	Timestamp   int64    `xml:"platformMsgs:timestamp"`
	Signature   struct {
		Algorithm string `xml:"algorithm,attr"`
		Text      string `xml:",chardata"`
	} `xml:"platformMsgs:signature"`
}

type Preferences struct {
}

type SearchPreferences struct {
	PageSize            int  `xml:"platformMsgs:pageSize,omitempty"`
	BodyFieldsOnly      bool `xml:"platformMsgs:bodyFieldsOnly"`
	ReturnSearchColumns bool `xml:"platformMsgs:returnSearchColumns"`
}

func (s SearchPreferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(s, e, start)
}

func (s SearchPreferences) IsEmpty() bool {
	return zero.IsZero(s)
}
