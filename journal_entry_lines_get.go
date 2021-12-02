package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewJournalEntryLinesGetRequest() JournalEntryLinesGetRequest {
	r := JournalEntryLinesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntryLinesGetRequest struct {
	client      *Client
	queryParams *JournalEntryLinesGetRequestQueryParams
	pathParams  *JournalEntryLinesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntryLinesGetRequestBody
}

func (r JournalEntryLinesGetRequest) NewQueryParams() *JournalEntryLinesGetRequestQueryParams {
	return &JournalEntryLinesGetRequestQueryParams{}
}

type JournalEntryLinesGetRequestQueryParams struct {
	Fields Fields `schema:"fields,omitempty"`
}

func (p JournalEntryLinesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntryLinesGetRequest) QueryParams() *JournalEntryLinesGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntryLinesGetRequest) NewPathParams() *JournalEntryLinesGetRequestPathParams {
	return &JournalEntryLinesGetRequestPathParams{}
}

type JournalEntryLinesGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *JournalEntryLinesGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *JournalEntryLinesGetRequest) PathParams() *JournalEntryLinesGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntryLinesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntryLinesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntryLinesGetRequest) Method() string {
	return r.method
}

func (r JournalEntryLinesGetRequest) NewRequestBody() JournalEntryLinesGetRequestBody {
	return JournalEntryLinesGetRequestBody{}
}

type JournalEntryLinesGetRequestBody struct {
}

func (r *JournalEntryLinesGetRequest) RequestBody() *JournalEntryLinesGetRequestBody {
	return nil
}

func (r *JournalEntryLinesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalEntryLinesGetRequest) SetRequestBody(body JournalEntryLinesGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntryLinesGetRequest) NewResponseBody() *JournalEntryLinesGetResponseBody {
	return &JournalEntryLinesGetResponseBody{}
}

type JournalEntryLinesGetResponseBody struct {
	Links Links `json:"links"`
	Items []struct {
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
	} `json:"items"`
	TotalResults int `json:"totalResults"`
}

func (r *JournalEntryLinesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/journalEntry/{{.id}}/line", r.PathParams())
	return &u, err
}

func (r *JournalEntryLinesGetRequest) Do() (JournalEntryLinesGetResponseBody, error) {
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
