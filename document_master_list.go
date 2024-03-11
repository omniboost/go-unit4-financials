package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewDocumentMasterListRequest() DocumentMasterListRequest {
	r := DocumentMasterListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DocumentMasterListRequest struct {
	client      *Client
	queryParams *DocumentMasterListRequestQueryParams
	pathParams  *DocumentMasterListRequestPathParams
	method      string
	headers     http.Header
	requestBody DocumentMasterListRequestBody
}

func (r DocumentMasterListRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/DocumentMaster/List"
}

func (r DocumentMasterListRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r DocumentMasterListRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/documentmaster/documentmaster-17.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/documentmaster/documentmaster-17.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:elem"}, Value: "http://www.coda.com/efinance/schemas/documentmaster"},
	}
}

func (r DocumentMasterListRequest) NewQueryParams() *DocumentMasterListRequestQueryParams {
	return &DocumentMasterListRequestQueryParams{}
}

type DocumentMasterListRequestQueryParams struct {
}

func (p DocumentMasterListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DocumentMasterListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r DocumentMasterListRequest) NewPathParams() *DocumentMasterListRequestPathParams {
	return &DocumentMasterListRequestPathParams{}
}

type DocumentMasterListRequestPathParams struct {
}

func (p *DocumentMasterListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentMasterListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *DocumentMasterListRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentMasterListRequest) Method() string {
	return r.method
}

func (r DocumentMasterListRequest) NewRequestBody() DocumentMasterListRequestBody {
	return DocumentMasterListRequestBody{}
}

type DocumentMasterListRequestBody struct {
	XMLName xml.Name `xml:"web:ListRequest"`

	Filter DocSeletKey `xml:"web:Filter"`
}

func (r *DocumentMasterListRequest) RequestBody() *DocumentMasterListRequestBody {
	return &r.requestBody
}

func (r *DocumentMasterListRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *DocumentMasterListRequest) SetRequestBody(body DocumentMasterListRequestBody) {
	r.requestBody = body
}

func (r *DocumentMasterListRequest) NewResponseBody() *DocumentMasterListRequestResponseBody {
	return &DocumentMasterListRequestResponseBody{}
}

type DocumentMasterListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetEnvironmentResponse"`
}

func (r *DocumentMasterListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/documentmaster/documentmaster-17.0", r.PathParams())
	return &u, err
}

func (r *DocumentMasterListRequest) Do() (DocumentMasterListRequestResponseBody, error) {
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

