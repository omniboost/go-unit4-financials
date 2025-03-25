package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewElementMasterAddRequest() ElementMasterAddRequest {
	r := ElementMasterAddRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ElementMasterAddRequest struct {
	client      *Client
	queryParams *ElementMasterAddRequestQueryParams
	pathParams  *ElementMasterAddRequestPathParams
	method      string
	headers     http.Header
	requestBody ElementMasterAddRequestBody
}

func (r ElementMasterAddRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/ElementMaster/Add"
}

func (r ElementMasterAddRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r ElementMasterAddRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/elementmaster/elementmaster-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/elementmaster/elementmaster-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:elem"}, Value: "http://www.coda.com/efinance/schemas/elementmaster"},
	}
}

func (r ElementMasterAddRequest) NewQueryParams() *ElementMasterAddRequestQueryParams {
	return &ElementMasterAddRequestQueryParams{}
}

type ElementMasterAddRequestQueryParams struct {
}

func (p ElementMasterAddRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ElementMasterAddRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ElementMasterAddRequest) NewPathParams() *ElementMasterAddRequestPathParams {
	return &ElementMasterAddRequestPathParams{}
}

type ElementMasterAddRequestPathParams struct {
}

func (p *ElementMasterAddRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ElementMasterAddRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ElementMasterAddRequest) SetMethod(method string) {
	r.method = method
}

func (r *ElementMasterAddRequest) Method() string {
	return r.method
}

func (r ElementMasterAddRequest) NewRequestBody() ElementMasterAddRequestBody {
	return ElementMasterAddRequestBody{}
}

type ElementMasterAddRequestBody struct {
	XMLName        xml.Name          `xml:"AddRequest"`
	Element        Element           `xml:"Element"`
	FlexiFieldData *FlexiFieldMaster `xml:"FlexiFieldData"`
}

func (r *ElementMasterAddRequest) RequestBody() *ElementMasterAddRequestBody {
	return &r.requestBody
}

func (r *ElementMasterAddRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ElementMasterAddRequest) SetRequestBody(body ElementMasterAddRequestBody) {
	r.requestBody = body
}

func (r *ElementMasterAddRequest) NewResponseBody() *ElementMasterAddRequestResponseBody {
	return &ElementMasterAddRequestResponseBody{}
}

type ElementMasterAddRequestResponseBody struct {
	XMLName xml.Name `xml:"GetEnvironmentResponse"`
}

func (r *ElementMasterAddRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/elementmaster/elementmaster-14.0", r.PathParams())
	return &u, err
}

func (r *ElementMasterAddRequest) Do() (ElementMasterAddRequestResponseBody, error) {
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
