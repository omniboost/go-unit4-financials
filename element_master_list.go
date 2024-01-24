package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewElementMasterListRequest() ElementMasterListRequest {
	r := ElementMasterListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ElementMasterListRequest struct {
	client      *Client
	queryParams *ElementMasterListRequestQueryParams
	pathParams  *ElementMasterListRequestPathParams
	method      string
	headers     http.Header
	requestBody ElementMasterListRequestBody
}

func (r ElementMasterListRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/ElementMaster/List"
}

func (r ElementMasterListRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r ElementMasterListRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/elementmaster/elementmaster-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/elementmaster/elementmaster-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:elem"}, Value: "http://www.coda.com/efinance/schemas/elementmaster"},
	}
}

func (r ElementMasterListRequest) NewQueryParams() *ElementMasterListRequestQueryParams {
	return &ElementMasterListRequestQueryParams{}
}

type ElementMasterListRequestQueryParams struct {
}

func (p ElementMasterListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ElementMasterListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ElementMasterListRequest) NewPathParams() *ElementMasterListRequestPathParams {
	return &ElementMasterListRequestPathParams{}
}

type ElementMasterListRequestPathParams struct {
}

func (p *ElementMasterListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ElementMasterListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ElementMasterListRequest) SetMethod(method string) {
	r.method = method
}

func (r *ElementMasterListRequest) Method() string {
	return r.method
}

func (r ElementMasterListRequest) NewRequestBody() ElementMasterListRequestBody {
	return ElementMasterListRequestBody{}
}

type ElementMasterListRequestBody struct {
	XMLName xml.Name `xml:"web:ListRequest"`

	Filter ElmReqFullKeys `xml:"web:Filter"`
}

func (r *ElementMasterListRequest) RequestBody() *ElementMasterListRequestBody {
	return &r.requestBody
}

func (r *ElementMasterListRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ElementMasterListRequest) SetRequestBody(body ElementMasterListRequestBody) {
	r.requestBody = body
}

func (r *ElementMasterListRequest) NewResponseBody() *ElementMasterListRequestResponseBody {
	return &ElementMasterListRequestResponseBody{}
}

type ElementMasterListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetEnvironmentResponse"`
}

func (r *ElementMasterListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/elementmaster/elementmaster-14.0", r.PathParams())
	return &u, err
}

func (r *ElementMasterListRequest) Do() (ElementMasterListRequestResponseBody, error) {
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
