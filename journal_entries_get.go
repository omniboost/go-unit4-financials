package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewJournalEntriesGetRequest() JournalEntriesGetRequest {
	r := JournalEntriesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntriesGetRequest struct {
	client      *Client
	queryParams *JournalEntriesGetRequestQueryParams
	pathParams  *JournalEntriesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntriesGetRequestBody
}

func (r JournalEntriesGetRequest) NewQueryParams() *JournalEntriesGetRequestQueryParams {
	return &JournalEntriesGetRequestQueryParams{}
}

type JournalEntriesGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p JournalEntriesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntriesGetRequest) QueryParams() *JournalEntriesGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntriesGetRequest) NewPathParams() *JournalEntriesGetRequestPathParams {
	return &JournalEntriesGetRequestPathParams{}
}

type JournalEntriesGetRequestPathParams struct {
}

func (p *JournalEntriesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalEntriesGetRequest) PathParams() *JournalEntriesGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntriesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntriesGetRequest) Method() string {
	return r.method
}

func (r JournalEntriesGetRequest) NewRequestBody() JournalEntriesGetRequestBody {
	return JournalEntriesGetRequestBody{}
}

type JournalEntriesGetRequestBody struct {
}

func (r *JournalEntriesGetRequest) RequestBody() *JournalEntriesGetRequestBody {
	return nil
}

func (r *JournalEntriesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalEntriesGetRequest) SetRequestBody(body JournalEntriesGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntriesGetRequest) NewResponseBody() *JournalEntriesGetResponseBody {
	return &JournalEntriesGetResponseBody{}
}

type JournalEntriesGetResponseBody struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Count   int  `json:"count"`
	HasMore bool `json:"hasMore"`
	Items   []struct {
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		ID string `json:"id"`
	} `json:"items"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

func (r *JournalEntriesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/journalEntry", r.PathParams())
	return &u, err
}

func (r *JournalEntriesGetRequest) Do() (JournalEntriesGetResponseBody, error) {
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
