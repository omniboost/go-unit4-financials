package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewInvoiceGetRequest() InvoiceGetRequest {
	r := InvoiceGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoiceGetRequest struct {
	client      *Client
	queryParams *InvoiceGetRequestQueryParams
	pathParams  *InvoiceGetRequestPathParams
	method      string
	headers     http.Header
	requestBody InvoiceGetRequestBody
}

func (r InvoiceGetRequest) NewQueryParams() *InvoiceGetRequestQueryParams {
	return &InvoiceGetRequestQueryParams{}
}

type InvoiceGetRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p InvoiceGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoiceGetRequest) QueryParams() *InvoiceGetRequestQueryParams {
	return r.queryParams
}

func (r InvoiceGetRequest) NewPathParams() *InvoiceGetRequestPathParams {
	return &InvoiceGetRequestPathParams{}
}

type InvoiceGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *InvoiceGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *InvoiceGetRequest) PathParams() *InvoiceGetRequestPathParams {
	return r.pathParams
}

func (r *InvoiceGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoiceGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoiceGetRequest) Method() string {
	return r.method
}

func (r InvoiceGetRequest) NewRequestBody() InvoiceGetRequestBody {
	return InvoiceGetRequestBody{}
}

type InvoiceGetRequestBody struct {
}

func (r *InvoiceGetRequest) RequestBody() *InvoiceGetRequestBody {
	return nil
}

func (r *InvoiceGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *InvoiceGetRequest) SetRequestBody(body InvoiceGetRequestBody) {
	r.requestBody = body
}

func (r *InvoiceGetRequest) NewResponseBody() *InvoiceGetResponseBody {
	return &InvoiceGetResponseBody{}
}

type InvoiceGetResponseBody struct {
	Links Links `json:"links"`
	Invoice
}

func (r *InvoiceGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/invoice/{{.id}}", r.PathParams())
	return &u, err
}

func (r *InvoiceGetRequest) Do() (InvoiceGetResponseBody, error) {
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
