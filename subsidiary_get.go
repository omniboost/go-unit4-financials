package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewSubsidiaryGetRequest() SubsidiaryGetRequest {
	r := SubsidiaryGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SubsidiaryGetRequest struct {
	client      *Client
	queryParams *SubsidiaryGetRequestQueryParams
	pathParams  *SubsidiaryGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SubsidiaryGetRequestBody
}

func (r SubsidiaryGetRequest) NewQueryParams() *SubsidiaryGetRequestQueryParams {
	return &SubsidiaryGetRequestQueryParams{}
}

type SubsidiaryGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p SubsidiaryGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SubsidiaryGetRequest) QueryParams() *SubsidiaryGetRequestQueryParams {
	return r.queryParams
}

func (r SubsidiaryGetRequest) NewPathParams() *SubsidiaryGetRequestPathParams {
	return &SubsidiaryGetRequestPathParams{}
}

type SubsidiaryGetRequestPathParams struct {
}

func (p *SubsidiaryGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SubsidiaryGetRequest) PathParams() *SubsidiaryGetRequestPathParams {
	return r.pathParams
}

func (r *SubsidiaryGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SubsidiaryGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SubsidiaryGetRequest) Method() string {
	return r.method
}

func (r SubsidiaryGetRequest) NewRequestBody() SubsidiaryGetRequestBody {
	return SubsidiaryGetRequestBody{}
}

type SubsidiaryGetRequestBody struct {
}

func (r *SubsidiaryGetRequest) RequestBody() *SubsidiaryGetRequestBody {
	return nil
}

func (r *SubsidiaryGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SubsidiaryGetRequest) SetRequestBody(body SubsidiaryGetRequestBody) {
	r.requestBody = body
}

func (r *SubsidiaryGetRequest) NewResponseBody() *SubsidiaryGetResponseBody {
	return &SubsidiaryGetResponseBody{}
}

type SubsidiaryGetResponseBody struct {
}

func (r *SubsidiaryGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/subsidiary", r.PathParams())
	return &u, err
}

func (r *SubsidiaryGetRequest) Do() (SubsidiaryGetResponseBody, error) {
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
	return *responseBody, err
}
