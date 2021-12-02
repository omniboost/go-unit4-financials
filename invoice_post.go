package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewInvoicePostRequest() InvoicePostRequest {
	r := InvoicePostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicePostRequest struct {
	client      *Client
	queryParams *InvoicePostRequestQueryParams
	pathParams  *InvoicePostRequestPathParams
	method      string
	headers     http.Header
	requestBody InvoicePostRequestBody
}

func (r InvoicePostRequest) NewQueryParams() *InvoicePostRequestQueryParams {
	return &InvoicePostRequestQueryParams{}
}

type InvoicePostRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p InvoicePostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicePostRequest) QueryParams() *InvoicePostRequestQueryParams {
	return r.queryParams
}

func (r InvoicePostRequest) NewPathParams() *InvoicePostRequestPathParams {
	return &InvoicePostRequestPathParams{}
}

type InvoicePostRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *InvoicePostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *InvoicePostRequest) PathParams() *InvoicePostRequestPathParams {
	return r.pathParams
}

func (r *InvoicePostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicePostRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicePostRequest) Method() string {
	return r.method
}

func (r InvoicePostRequest) NewRequestBody() InvoicePostRequestBody {
	return InvoicePostRequestBody{}
}

type InvoicePostRequestBody struct {
	Invoice
}

func (r *InvoicePostRequest) RequestBody() *InvoicePostRequestBody {
	return &r.requestBody
}

func (r *InvoicePostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InvoicePostRequest) SetRequestBody(body InvoicePostRequestBody) {
	r.requestBody = body
}

func (r *InvoicePostRequest) NewResponseBody() *InvoicePostResponseBody {
	return &InvoicePostResponseBody{}
}

type InvoicePostResponseBody struct {
}

func (r *InvoicePostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/invoice", r.PathParams())
	return &u, err
}

func (r *InvoicePostRequest) Do() (InvoicePostResponseBody, error) {
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
