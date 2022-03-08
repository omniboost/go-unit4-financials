package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetServerTimeRequest() GetServerTimeRequest {
	r := GetServerTimeRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetServerTimeRequest struct {
	client      *Client
	queryParams *GetServerTimeRequestQueryParams
	pathParams  *GetServerTimeRequestPathParams
	method      string
	headers     http.Header
	requestBody GetServerTimeRequestBody
}

func (r GetServerTimeRequest) SOAPAction() string {
	return "getServerTime"
}

func (r GetServerTimeRequest) NewQueryParams() *GetServerTimeRequestQueryParams {
	return &GetServerTimeRequestQueryParams{}
}

type GetServerTimeRequestQueryParams struct {
}

func (p GetServerTimeRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetServerTimeRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetServerTimeRequest) NewPathParams() *GetServerTimeRequestPathParams {
	return &GetServerTimeRequestPathParams{}
}

type GetServerTimeRequestPathParams struct {
}

func (p *GetServerTimeRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetServerTimeRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetServerTimeRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetServerTimeRequest) Method() string {
	return r.method
}

func (r GetServerTimeRequest) NewRequestBody() GetServerTimeRequestBody {
	return GetServerTimeRequestBody{}
}

type GetServerTimeRequestBody struct {
	XMLName xml.Name `xml:"GetServerTime"`
}

func (r *GetServerTimeRequest) RequestBody() *GetServerTimeRequestBody {
	return &r.requestBody
}

func (r *GetServerTimeRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetServerTimeRequest) SetRequestBody(body GetServerTimeRequestBody) {
	r.requestBody = body
}

func (r *GetServerTimeRequest) NewResponseBody() *GetServerTimeRequestResponseBody {
	return &GetServerTimeRequestResponseBody{}
}

type GetServerTimeRequestResponseBody struct {
	XMLName xml.Name `xml:"GetServerTimeResponse"`

	GetServerTimeResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		Status       struct {
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		ServerTime string `xml:"serverTime"`
	} `xml:"getServerTimeResult"`
}

func (r *GetServerTimeRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetServerTimeRequest) Do() (GetServerTimeRequestResponseBody, error) {
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
