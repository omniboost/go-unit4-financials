package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewEnvironmentGetRequest() EnvironmentGetRequest {
	r := EnvironmentGetRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EnvironmentGetRequest struct {
	client      *Client
	queryParams *EnvironmentGetRequestQueryParams
	pathParams  *EnvironmentGetRequestPathParams
	method      string
	headers     http.Header
	requestBody EnvironmentGetRequestBody
}

func (r EnvironmentGetRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/AppServer/GetEnvironment"
}

func (r EnvironmentGetRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r EnvironmentGetRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/appserver/appserver-12.0/webservice"},
	}
}

func (r EnvironmentGetRequest) NewQueryParams() *EnvironmentGetRequestQueryParams {
	return &EnvironmentGetRequestQueryParams{}
}

type EnvironmentGetRequestQueryParams struct {
}

func (p EnvironmentGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EnvironmentGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r EnvironmentGetRequest) NewPathParams() *EnvironmentGetRequestPathParams {
	return &EnvironmentGetRequestPathParams{}
}

type EnvironmentGetRequestPathParams struct {
}

func (p *EnvironmentGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EnvironmentGetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *EnvironmentGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EnvironmentGetRequest) Method() string {
	return r.method
}

func (r EnvironmentGetRequest) NewRequestBody() EnvironmentGetRequestBody {
	return EnvironmentGetRequestBody{}
}

type EnvironmentGetRequestBody struct {
	XMLName xml.Name `xml:"GetEnvironmentRequest"`
}

func (r *EnvironmentGetRequest) RequestBody() *EnvironmentGetRequestBody {
	return &r.requestBody
}

func (r *EnvironmentGetRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *EnvironmentGetRequest) SetRequestBody(body EnvironmentGetRequestBody) {
	r.requestBody = body
}

func (r *EnvironmentGetRequest) NewResponseBody() *EnvironmentGetRequestResponseBody {
	return &EnvironmentGetRequestResponseBody{}
}

type EnvironmentGetRequestResponseBody struct {
	XMLName xml.Name `xml:"GetEnvironmentResponse"`

	Text        string `xml:",chardata"`
	Webservice  string `xml:"webservice,attr"`
	Xmlns       string `xml:"xmlns,attr"`
	Com         string `xml:"com,attr"`
	Environment struct {
		CmpCode           string `xml:"CmpCode"`
		CapCode           string `xml:"CapCode"`
		HomeCurr          string `xml:"HomeCurr"`
		DateOrder         string `xml:"DateOrder"`
		DateDisplay       string `xml:"DateDisplay"`
		DateSep           string `xml:"DateSep"`
		HomeCurrDps       string `xml:"HomeCurrDps"`
		HomeCurrSymbol    string `xml:"HomeCurrSymbol"`
		HomeCurrSymbolPos string `xml:"HomeCurrSymbolPos"`
		HomeCurrLinkType  string `xml:"HomeCurrLinkType"`
		HomeCurrParent    string `xml:"HomeCurrParent"`
		CurrentPeriod     string `xml:"CurrentPeriod"`
		CmpTimeStamp      string `xml:"CmpTimeStamp"`
		CapTimeStamp      string `xml:"CapTimeStamp"`
		UsrTimeStamp      string `xml:"UsrTimeStamp"`
		ServerVersion     string `xml:"ServerVersion"`
		SubsChar          string `xml:"SubsChar"`
	} `xml:"Environment"`
}

func (r *EnvironmentGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/appserver/appserver-12.0", r.PathParams())
	return &u, err
}

func (r *EnvironmentGetRequest) Do() (EnvironmentGetRequestResponseBody, error) {
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
