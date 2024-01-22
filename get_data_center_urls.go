package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetDataCenterURLsRequest() GetDataCenterURLsRequest {
	r := GetDataCenterURLsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetDataCenterURLsRequest struct {
	client      *Client
	queryParams *GetDataCenterURLsRequestQueryParams
	pathParams  *GetDataCenterURLsRequestPathParams
	method      string
	headers     http.Header
	requestBody GetDataCenterURLsRequestBody
}

func (r GetDataCenterURLsRequest) SOAPAction() string {
	return "getDataCenterUrls"
}

func (r GetDataCenterURLsRequest) NewQueryParams() *GetDataCenterURLsRequestQueryParams {
	return &GetDataCenterURLsRequestQueryParams{}
}

type GetDataCenterURLsRequestQueryParams struct {
}

func (p GetDataCenterURLsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDataCenterURLsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetDataCenterURLsRequest) NewPathParams() *GetDataCenterURLsRequestPathParams {
	return &GetDataCenterURLsRequestPathParams{}
}

type GetDataCenterURLsRequestPathParams struct {
}

func (p *GetDataCenterURLsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDataCenterURLsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetDataCenterURLsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDataCenterURLsRequest) Method() string {
	return r.method
}

func (r GetDataCenterURLsRequest) NewRequestBody() GetDataCenterURLsRequestBody {
	return GetDataCenterURLsRequestBody{}
}

type GetDataCenterURLsRequestBody struct {
	XMLName xml.Name `xml:"getDataCenterUrls"`

	Account string `xml:"urn:Account"`
}

func (r *GetDataCenterURLsRequest) RequestBody() *GetDataCenterURLsRequestBody {
	return &r.requestBody
}

func (r *GetDataCenterURLsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetDataCenterURLsRequest) SetRequestBody(body GetDataCenterURLsRequestBody) {
	r.requestBody = body
}

func (r *GetDataCenterURLsRequest) NewResponseBody() *GetDataCenterURLsRequestResponseBody {
	return &GetDataCenterURLsRequestResponseBody{}
}

type GetDataCenterURLsRequestResponseBody struct {
	XMLName xml.Name `xml:"getDataCenterUrlsResponse"`

	GetDataCenterUrlsResult struct {
		Text         string `xml:",chardata"`
		PlatformCore string `xml:"platformCore,attr"`
		Status       struct {
			Text      string `xml:",chardata"`
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		DataCenterUrls struct {
			Text              string `xml:",chardata"`
			RestDomain        string `xml:"restDomain"`
			WebservicesDomain string `xml:"webservicesDomain"`
			SystemDomain      string `xml:"systemDomain"`
		} `xml:"dataCenterUrls"`
	} `xml:"getDataCenterUrlsResult"`
}

func (r *GetDataCenterURLsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetDataCenterURLsRequest) Do() (GetDataCenterURLsRequestResponseBody, error) {
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
