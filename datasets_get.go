package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewDataSetsGetRequest() DataSetsGetRequest {
	r := DataSetsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DataSetsGetRequest struct {
	client      *Client
	queryParams *DataSetsGetRequestQueryParams
	pathParams  *DataSetsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody DataSetsGetRequestBody
}

func (r DataSetsGetRequest) NewQueryParams() *DataSetsGetRequestQueryParams {
	return &DataSetsGetRequestQueryParams{}
}

type DataSetsGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p DataSetsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DataSetsGetRequest) QueryParams() *DataSetsGetRequestQueryParams {
	return r.queryParams
}

func (r DataSetsGetRequest) NewPathParams() *DataSetsGetRequestPathParams {
	return &DataSetsGetRequestPathParams{}
}

type DataSetsGetRequestPathParams struct {
}

func (p *DataSetsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DataSetsGetRequest) PathParams() *DataSetsGetRequestPathParams {
	return r.pathParams
}

func (r *DataSetsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DataSetsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DataSetsGetRequest) Method() string {
	return r.method
}

func (r DataSetsGetRequest) NewRequestBody() DataSetsGetRequestBody {
	return DataSetsGetRequestBody{}
}

type DataSetsGetRequestBody struct {
}

func (r *DataSetsGetRequest) RequestBody() *DataSetsGetRequestBody {
	return nil
}

func (r *DataSetsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *DataSetsGetRequest) SetRequestBody(body DataSetsGetRequestBody) {
	r.requestBody = body
}

func (r *DataSetsGetRequest) NewResponseBody() *DataSetsGetResponseBody {
	return &DataSetsGetResponseBody{}
}

type DataSetsGetResponseBody struct {
}

func (r *DataSetsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/query/v1/dataset", r.PathParams())
	return &u, err
}

func (r *DataSetsGetRequest) Do() (DataSetsGetResponseBody, error) {
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

	req.Header.Del("Accept-Language")

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
