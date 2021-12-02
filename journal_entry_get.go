package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewJournalEntryGetRequest() JournalEntryGetRequest {
	r := JournalEntryGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntryGetRequest struct {
	client      *Client
	queryParams *JournalEntryGetRequestQueryParams
	pathParams  *JournalEntryGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntryGetRequestBody
}

func (r JournalEntryGetRequest) NewQueryParams() *JournalEntryGetRequestQueryParams {
	return &JournalEntryGetRequestQueryParams{}
}

type JournalEntryGetRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p JournalEntryGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntryGetRequest) QueryParams() *JournalEntryGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntryGetRequest) NewPathParams() *JournalEntryGetRequestPathParams {
	return &JournalEntryGetRequestPathParams{}
}

type JournalEntryGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *JournalEntryGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *JournalEntryGetRequest) PathParams() *JournalEntryGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntryGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntryGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntryGetRequest) Method() string {
	return r.method
}

func (r JournalEntryGetRequest) NewRequestBody() JournalEntryGetRequestBody {
	return JournalEntryGetRequestBody{}
}

type JournalEntryGetRequestBody struct {
}

func (r *JournalEntryGetRequest) RequestBody() *JournalEntryGetRequestBody {
	return nil
}

func (r *JournalEntryGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalEntryGetRequest) SetRequestBody(body JournalEntryGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntryGetRequest) NewResponseBody() *JournalEntryGetResponseBody {
	return &JournalEntryGetResponseBody{}
}

type JournalEntryGetResponseBody struct {
	Links Links `json:"links"`
	JournalEntry
}

func (r *JournalEntryGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/journalEntry/{{.id}}", r.PathParams())
	return &u, err
}

func (r *JournalEntryGetRequest) Do() (JournalEntryGetResponseBody, error) {
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
