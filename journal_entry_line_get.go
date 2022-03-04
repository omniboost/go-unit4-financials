package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-soap/utils"
)

func (c *Client) NewJournalEntryLineGetRequest() JournalEntryLineGetRequest {
	r := JournalEntryLineGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntryLineGetRequest struct {
	client      *Client
	queryParams *JournalEntryLineGetRequestQueryParams
	pathParams  *JournalEntryLineGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntryLineGetRequestBody
}

func (r JournalEntryLineGetRequest) NewQueryParams() *JournalEntryLineGetRequestQueryParams {
	return &JournalEntryLineGetRequestQueryParams{}
}

type JournalEntryLineGetRequestQueryParams struct {
	Fields Fields `schema:"fields,omitempty"`
}

func (p JournalEntryLineGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntryLineGetRequest) QueryParams() *JournalEntryLineGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntryLineGetRequest) NewPathParams() *JournalEntryLineGetRequestPathParams {
	return &JournalEntryLineGetRequestPathParams{}
}

type JournalEntryLineGetRequestPathParams struct {
	JournalEntryID     int `schema:"journal_entry_id"`
	JournalEntryLineNo int `schema:"journal_entry_line_no"`
}

func (p *JournalEntryLineGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"journal_entry_id":      strconv.Itoa(p.JournalEntryID),
		"journal_entry_line_no": strconv.Itoa(p.JournalEntryLineNo),
	}
}

func (r *JournalEntryLineGetRequest) PathParams() *JournalEntryLineGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntryLineGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntryLineGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntryLineGetRequest) Method() string {
	return r.method
}

func (r JournalEntryLineGetRequest) NewRequestBody() JournalEntryLineGetRequestBody {
	return JournalEntryLineGetRequestBody{}
}

type JournalEntryLineGetRequestBody struct {
}

func (r *JournalEntryLineGetRequest) RequestBody() *JournalEntryLineGetRequestBody {
	return nil
}

func (r *JournalEntryLineGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalEntryLineGetRequest) SetRequestBody(body JournalEntryLineGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntryLineGetRequest) NewResponseBody() *JournalEntryLineGetResponseBody {
	return &JournalEntryLineGetResponseBody{}
}

type JournalEntryLineGetResponseBody struct {
	Links   Links `json:"links"`
	Account struct {
		Links   []interface{} `json:"links"`
		ID      string        `json:"id"`
		RefName string        `json:"refName"`
	} `json:"account"`
	Cleared             bool    `json:"cleared"`
	Credit              float64 `json:"credit"`
	Debit               float64 `json:"debit"`
	Custcol2663Isperson bool    `json:"custcol_2663_isperson"`
	Eliminate           bool    `json:"eliminate"`
	Line                int     `json:"line"`
}

func (r *JournalEntryLineGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/journalEntry/{{.journal_entry_id}}/line/{{.journal_entry_line_no}}", r.PathParams())
	return &u, err
}

func (r *JournalEntryLineGetRequest) Do() (JournalEntryLineGetResponseBody, error) {
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
