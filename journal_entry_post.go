package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewJournalEntryPostRequest() JournalEntryPostRequest {
	r := JournalEntryPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntryPostRequest struct {
	client      *Client
	queryParams *JournalEntryPostRequestQueryParams
	pathParams  *JournalEntryPostRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntryPostRequestBody
}

func (r JournalEntryPostRequest) NewQueryParams() *JournalEntryPostRequestQueryParams {
	return &JournalEntryPostRequestQueryParams{}
}

type JournalEntryPostRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p JournalEntryPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntryPostRequest) QueryParams() *JournalEntryPostRequestQueryParams {
	return r.queryParams
}

func (r JournalEntryPostRequest) NewPathParams() *JournalEntryPostRequestPathParams {
	return &JournalEntryPostRequestPathParams{}
}

type JournalEntryPostRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *JournalEntryPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *JournalEntryPostRequest) PathParams() *JournalEntryPostRequestPathParams {
	return r.pathParams
}

func (r *JournalEntryPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntryPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntryPostRequest) Method() string {
	return r.method
}

func (r JournalEntryPostRequest) NewRequestBody() JournalEntryPostRequestBody {
	return JournalEntryPostRequestBody{}
}

type JournalEntryPostRequestBody struct {
	JournalEntry
}

func (r *JournalEntryPostRequest) RequestBody() *JournalEntryPostRequestBody {
	return &r.requestBody
}

func (r *JournalEntryPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *JournalEntryPostRequest) SetRequestBody(body JournalEntryPostRequestBody) {
	r.requestBody = body
}

func (r *JournalEntryPostRequest) NewResponseBody() *JournalEntryPostResponseBody {
	return &JournalEntryPostResponseBody{}
}

type JournalEntryPostResponseBody struct {
}

func (r *JournalEntryPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/journalEntry", r.PathParams())
	return &u, err
}

func (r *JournalEntryPostRequest) Do() (JournalEntryPostResponseBody, error) {
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
