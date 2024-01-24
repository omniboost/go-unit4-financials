package financials

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-unit4-financials/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name

	NS     []xml.Attr `xml:"-"`
	Header SOAPHeader `xml:"env:Header,omitempty"`
	Body   Body   `xml:"env:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		NS: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:env"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
			{Name: xml.Name{Space: "", Local: "xmlns:com"}, Value: "http://www.coda.com/efinance/schemas/common"},
		},
		Header: NewSOAPHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header SOAPHeader     `xml:"Header,omitempty"`
	Body   Body       `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "env:Envelope"}

	for _, ns := range env.NS {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)

	return omitempty.MarshalXML(a, e, start)
}

type Body struct {
	ActionBody interface{} `xml:",any"`
}

type SOAPHeader struct {
	Options Options `xml:"Options"`
}

func (h SOAPHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(h, e, start)
}

func (h SOAPHeader) IsEmpty() bool {
	return zero.IsZero(h)
}

func NewSOAPHeader() SOAPHeader {
	return SOAPHeader{}
}

type Options struct {
	Locale  string `xml:"locale,attr"`
	Session string `xml:"session,attr,omitempty"`
}
