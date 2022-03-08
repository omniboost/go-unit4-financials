package netsuite

import (
	"encoding/xml"
)

type RequestEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"s:Header"`
	Body   Body   `xml:"s:Body"`
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
	start.Name = xml.Name{Local: "s:Envelope"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:s"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		{Name: xml.Name{Space: "", Local: "xmlns:urn"}, Value: "urn:messages_2021_2.platform.webservices.netsuite.com"},
		{Name: xml.Name{Space: "", Local: "xmlns:urn1"}, Value: "urn:core_2021_2.platform.webservices.netsuite.com"},
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
	TokenPassport TokenPassport `xml:"urn:tokenPassport"`
	// DocumentInfo  struct {
	// 	NSInfo string `xml:"platformMsgs:nsId,omitempty"`
	// } `xml:"platformMsgs:documentInfo,omitempty"`
}

func NewHeader() Header {
	return Header{}
}

type ActionBody interface{}

type TokenPassport struct {
	XMLName     xml.Name `xml:"urn:tokenPassport"`
	Account     string   `xml:"urn:account"`
	ConsumerKey string   `xml:"urn:consumerKey"`
	Token       string   `xml:"urn:token"`
	Nonce       string   `xml:"urn:nonce"`
	Timestamp   int64    `xml:"urn:timestamp"`
	Signature   struct {
		Algorithm string `xml:"algorithm,attr"`
		Text      string `xml:",chardata"`
	} `xml:"urn:signature"`
}
