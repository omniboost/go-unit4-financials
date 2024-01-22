package financials

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-unit4-financials/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name

	NS     []xml.Attr `xml:"-"`
	Header Header `xml:"env:Header,omitempty"`
	Body   Body   `xml:"env:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		NS: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:env"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		},
		Header: NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header Header     `xml:"Header,omitempty"`
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

type Header struct {
	Options Options `xml:"Options"`
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

type Options struct {
	Locale  string `xml:"locale,attr"`
	Session string `xml:"session,attr,omitempty"`
}
