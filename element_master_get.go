package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewElementMasterGetRequest() ElementMasterGetRequest {
	r := ElementMasterGetRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ElementMasterGetRequest struct {
	client      *Client
	queryParams *ElementMasterGetRequestQueryParams
	pathParams  *ElementMasterGetRequestPathParams
	method      string
	headers     http.Header
	requestBody ElementMasterGetRequestBody
}

func (r ElementMasterGetRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/ElementMaster/Get"
}

func (r ElementMasterGetRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r ElementMasterGetRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/elementmaster/elementmaster-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/elementmaster/elementmaster-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:elem"}, Value: "http://www.coda.com/efinance/schemas/elementmaster"},
	}
}

func (r ElementMasterGetRequest) NewQueryParams() *ElementMasterGetRequestQueryParams {
	return &ElementMasterGetRequestQueryParams{}
}

type ElementMasterGetRequestQueryParams struct {
}

func (p ElementMasterGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ElementMasterGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ElementMasterGetRequest) NewPathParams() *ElementMasterGetRequestPathParams {
	return &ElementMasterGetRequestPathParams{}
}

type ElementMasterGetRequestPathParams struct {
}

func (p *ElementMasterGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ElementMasterGetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ElementMasterGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ElementMasterGetRequest) Method() string {
	return r.method
}

func (r ElementMasterGetRequest) NewRequestBody() ElementMasterGetRequestBody {
	return ElementMasterGetRequestBody{}
}

type ElementMasterGetRequestBody struct {
	XMLName xml.Name `xml:"GetRequest"`
	Key     struct {
		CmpCode string `xml:"elem:CmpCode"`
		Level   int    `xml:"elem:Level"`
		Code    string `xml:"elem:Code,omitempty"`
	} `xml:"Key"`
}

func (r *ElementMasterGetRequest) RequestBody() *ElementMasterGetRequestBody {
	return &r.requestBody
}

func (r *ElementMasterGetRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ElementMasterGetRequest) SetRequestBody(body ElementMasterGetRequestBody) {
	r.requestBody = body
}

func (r *ElementMasterGetRequest) NewResponseBody() *ElementMasterGetRequestResponseBody {
	return &ElementMasterGetRequestResponseBody{}
}

type ElementMasterGetRequestResponseBody struct {
	XMLName xml.Name `xml:"GetResponse"`

	CmpCode string  `xml:"CmpCode"`
	Level   string  `xml:"Level"`
	Code    string  `xml:"Code"`
	Element Element `xml:"Element"`
}

func (r *ElementMasterGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/elementmaster/elementmaster-14.0", r.PathParams())
	return &u, err
}

func (r *ElementMasterGetRequest) Do() (ElementMasterGetRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}
