package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetAllRequest() GetAllRequest {
	r := GetAllRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAllRequest struct {
	client      *Client
	queryParams *GetAllRequestQueryParams
	pathParams  *GetAllRequestPathParams
	method      string
	headers     http.Header
	requestBody GetAllRequestBody
}

func (r GetAllRequest) SOAPAction() string {
	return "getAll"
}

func (r GetAllRequest) NewQueryParams() *GetAllRequestQueryParams {
	return &GetAllRequestQueryParams{}
}

type GetAllRequestQueryParams struct {
}

func (p GetAllRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAllRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetAllRequest) NewPathParams() *GetAllRequestPathParams {
	return &GetAllRequestPathParams{}
}

type GetAllRequestPathParams struct {
}

func (p *GetAllRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAllRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetAllRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAllRequest) Method() string {
	return r.method
}

func (r GetAllRequest) NewRequestBody() GetAllRequestBody {
	return GetAllRequestBody{}
}

type GetAllRequestBody struct {
	XMLName xml.Name `xml:"GetAll"`
}

func (r *GetAllRequest) RequestBody() *GetAllRequestBody {
	return &r.requestBody
}

func (r *GetAllRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetAllRequest) SetRequestBody(body GetAllRequestBody) {
	r.requestBody = body
}

func (r *GetAllRequest) NewResponseBody() *GetAllRequestResponseBody {
	return &GetAllRequestResponseBody{}
}

type GetAllRequestResponseBody struct {
	XMLName xml.Name `xml:"GetAllResponse"`
}

func (r *GetAllRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetAllRequest) Do() (GetAllRequestResponseBody, error) {
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
