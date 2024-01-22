package financials

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-unit4-financials/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"Header,omitempty"`
	Body   Body   `xml:"env:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		Header: NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"Header,omitempty"`
	Body   Body   `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "env:Envelope"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/wsadapter/schemas/authenticate/authenticate-2.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:env"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
	}
	for _, ns := range namespaces {
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
