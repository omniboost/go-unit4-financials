package netsuite

import (
	"github.com/cydev/zero"
)

// func (c *Client) NewSearchRequest() SearchRequest {
// 	r := SearchRequest{
// 		client:  c,
// 		method:  http.MethodPost,
// 		headers: http.Header{},
// 	}

// 	r.queryParams = r.NewQueryParams()
// 	r.pathParams = r.NewPathParams()
// 	r.requestBody = r.NewRequestBody()
// 	return r
// }

// type SearchRequest struct {
// 	client      *Client
// 	queryParams *SearchRequestQueryParams
// 	pathParams  *SearchRequestPathParams
// 	method      string
// 	headers     http.Header
// 	requestBody SearchRequestBody
// }

// func (r SearchRequest) SOAPAction() string {
// 	return "search"
// }

// func (r SearchRequest) NewQueryParams() *SearchRequestQueryParams {
// 	return &SearchRequestQueryParams{}
// }

// type SearchRequestQueryParams struct {
// }

// func (p SearchRequestQueryParams) ToURLValues() (url.Values, error) {
// 	encoder := utils.NewSchemaEncoder()
// 	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
// 	params := url.Values{}

// 	err := encoder.Encode(p, params)
// 	if err != nil {
// 		return params, err
// 	}

// 	return params, nil
// }

// func (r *SearchRequest) QueryParams() QueryParams {
// 	return r.queryParams
// }

// func (r SearchRequest) NewPathParams() *SearchRequestPathParams {
// 	return &SearchRequestPathParams{}
// }

// type SearchRequestPathParams struct {
// }

// func (p *SearchRequestPathParams) Params() map[string]string {
// 	return map[string]string{}
// }

// func (r *SearchRequest) PathParams() PathParams {
// 	return r.pathParams
// }

// func (r *SearchRequest) SetMethod(method string) {
// 	r.method = method
// }

// func (r *SearchRequest) Method() string {
// 	return r.method
// }

// func (r SearchRequest) NewRequestBody() SearchRequestBody {
// 	return SearchRequestBody{}
// }

// type SearchRequestBody struct {
// 	XMLName xml.Name `xml:"platformMsgs:search"`

// 	SearchRecord struct { // SearchRecordBasic
// 		Type  string      `xml:"xsi:type,attr"`
// 		Basic interface{} `xml:"listRel:basic"`
// 	} `xml:"platformMsgs:searchRecord"`
// }

// func (r *SearchRequest) RequestBody() *SearchRequestBody {
// 	return &r.requestBody
// }

// func (r *SearchRequest) RequestBodyInterface() interface{} {
// 	return &r.requestBody
// }

// func (r *SearchRequest) SetRequestBody(body SearchRequestBody) {
// 	r.requestBody = body
// }

// func (r *SearchRequest) NewResponseBody() *SearchRequestResponseBody {
// 	return &SearchRequestResponseBody{}
// }

// type SearchRequestResponseBody struct {
// 	XMLName xml.Name `xml:"SearchResponse"`
// }

// func (r *SearchRequest) URL() (*url.URL, error) {
// 	u, err := r.client.GetEndpointURL("", r.PathParams())
// 	return &u, err
// }

// func (r *SearchRequest) Do() (SearchRequestResponseBody, error) {
// 	var err error

// 	// Create http request
// 	req, err := r.client.NewRequest(nil, r)
// 	if err != nil {
// 		return *r.NewResponseBody(), err
// 	}

// 	// Process query parameters
// 	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
// 	if err != nil {
// 		return *r.NewResponseBody(), err
// 	}

// 	responseBody := r.NewResponseBody()
// 	_, err = r.client.Do(req, responseBody)
// 	if err != nil {
// 		return *responseBody, errors.WithStack(err)
// 	}

// 	return *responseBody, nil
// }

type SearchStringField struct {
	Operator    SearchStringFieldOperator `xml:"operator,attr"`
	SearchValue string                    `xml:"platformCore:searchValue"`
}

func (s SearchStringField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SearchStringFieldOperator string

type SearchMultiSelectField struct{}

func (s SearchMultiSelectField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SearchDoubleField struct{}

func (s SearchDoubleField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SearchEnumMultiSelectField struct{}

func (s SearchEnumMultiSelectField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SearchLongField struct{}

func (s SearchLongField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SearchBooleanField struct{}

func (s SearchBooleanField) IsEmpty() bool {
	return zero.IsZero(s)
}

type SearchCustomFieldList struct{}

func (s SearchCustomFieldList) IsEmpty() bool {
	return zero.IsZero(s)
}
