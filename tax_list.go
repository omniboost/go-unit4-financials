package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewTaxListRequest() TaxListRequest {
	r := TaxListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxListRequest struct {
	client      *Client
	queryParams *TaxListRequestQueryParams
	pathParams  *TaxListRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxListRequestBody
}

func (r TaxListRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/Tax/List"
}

func (r TaxListRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r TaxListRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/tax/tax-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/tax/tax-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:tax"}, Value: "http://www.coda.com/efinance/schemas/tax"},
	}
}

func (r TaxListRequest) NewQueryParams() *TaxListRequestQueryParams {
	return &TaxListRequestQueryParams{}
}

type TaxListRequestQueryParams struct {
}

func (p TaxListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TaxListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r TaxListRequest) NewPathParams() *TaxListRequestPathParams {
	return &TaxListRequestPathParams{}
}

type TaxListRequestPathParams struct {
}

func (p *TaxListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *TaxListRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxListRequest) Method() string {
	return r.method
}

func (r TaxListRequest) NewRequestBody() TaxListRequestBody {
	return TaxListRequestBody{}
}

type TaxListRequestBody struct {
	XMLName xml.Name `xml:"web:ListRequest"`

	Filter ReqKeys `xml:"web:Filter"`
}

func (r *TaxListRequest) RequestBody() *TaxListRequestBody {
	return &r.requestBody
}

func (r *TaxListRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *TaxListRequest) SetRequestBody(body TaxListRequestBody) {
	r.requestBody = body
}

func (r *TaxListRequest) NewResponseBody() *TaxListRequestResponseBody {
	return &TaxListRequestResponseBody{}
}

type TaxListRequestResponseBody struct {
	XMLName xml.Name `xml:"ListResponse"`

	Webservice string `xml:"webservice,attr"`
	Com        string `xml:"com,attr"`
	Filter     struct {
		MaxKeys string `xml:"MaxKeys"`
		Key     struct {
			CmpCode string `xml:"CmpCode"`
			Code    string `xml:"Code"`
		} `xml:"Key"`
	} `xml:"Filter"`
	Keys []struct {
		Code      string `xml:"Code"`
		ShortName string `xml:"ShortName"`
	} `xml:"Keys>Key"`
}

func (r *TaxListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/tax/tax-14.0", r.PathParams())
	return &u, err
}

func (r *TaxListRequest) Do() (TaxListRequestResponseBody, error) {
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
