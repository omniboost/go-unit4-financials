package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-soap/utils"
)

func (c *Client) NewCustomerGetRequest() CustomerGetRequest {
	r := CustomerGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerGetRequest struct {
	client      *Client
	queryParams *CustomerGetRequestQueryParams
	pathParams  *CustomerGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomerGetRequestBody
}

func (r CustomerGetRequest) NewQueryParams() *CustomerGetRequestQueryParams {
	return &CustomerGetRequestQueryParams{}
}

type CustomerGetRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p CustomerGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(Fields{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerGetRequest) QueryParams() *CustomerGetRequestQueryParams {
	return r.queryParams
}

func (r CustomerGetRequest) NewPathParams() *CustomerGetRequestPathParams {
	return &CustomerGetRequestPathParams{}
}

type CustomerGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *CustomerGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *CustomerGetRequest) PathParams() *CustomerGetRequestPathParams {
	return r.pathParams
}

func (r *CustomerGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerGetRequest) Method() string {
	return r.method
}

func (r CustomerGetRequest) NewRequestBody() CustomerGetRequestBody {
	return CustomerGetRequestBody{}
}

type CustomerGetRequestBody struct {
}

func (r *CustomerGetRequest) RequestBody() *CustomerGetRequestBody {
	return nil
}

func (r *CustomerGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerGetRequest) SetRequestBody(body CustomerGetRequestBody) {
	r.requestBody = body
}

func (r *CustomerGetRequest) NewResponseBody() *CustomerGetResponseBody {
	return &CustomerGetResponseBody{}
}

type CustomerGetResponseBody struct {
	Links Links `json:"links"`
	Customer
}

func (r *CustomerGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/customer/{{.id}}", r.PathParams())
	return &u, err
}

func (r *CustomerGetRequest) Do() (CustomerGetResponseBody, error) {
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
