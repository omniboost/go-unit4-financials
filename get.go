package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetRequest() GetRequest {
	r := GetRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetRequest struct {
	client      *Client
	queryParams *GetRequestQueryParams
	pathParams  *GetRequestPathParams
	method      string
	headers     http.Header
	requestBody GetRequestBody
}

func (r GetRequest) SOAPAction() string {
	return "get"
}

func (r GetRequest) NewQueryParams() *GetRequestQueryParams {
	return &GetRequestQueryParams{}
}

type GetRequestQueryParams struct {
}

func (p GetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetRequest) NewPathParams() *GetRequestPathParams {
	return &GetRequestPathParams{}
}

type GetRequestPathParams struct {
}

func (p *GetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetRequest) Method() string {
	return r.method
}

func (r GetRequest) NewRequestBody() GetRequestBody {
	return GetRequestBody{}
}

type GetRequestBody struct {
	XMLName xml.Name `xml:"get"`

	BaseRef BaseRef `xml:"baseRef"`
}

func (r *GetRequest) RequestBody() *GetRequestBody {
	return &r.requestBody
}

func (r *GetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetRequest) SetRequestBody(body GetRequestBody) {
	r.requestBody = body
}

func (r *GetRequest) NewResponseBody() *GetRequestResponseBody {
	return &GetRequestResponseBody{}
}

type GetRequestResponseBody struct {
	XMLName xml.Name `xml:"getResponse"`

	Text         string `xml:",chardata"`
	Xmlns        string `xml:"xmlns,attr"`
	ReadResponse struct {
		Text         string `xml:",chardata"`
		PlatformMsgs string `xml:"platformMsgs,attr"`
		Status       struct {
			Text         string `xml:",chardata"`
			IsSuccess    string `xml:"isSuccess,attr"`
			PlatformCore string `xml:"platformCore,attr"`
		} `xml:"status"`
		// Record TransactionBodyCustomField `xml:"record"`
		Record struct {
			XML []byte `xml:",innerxml"`
		} `xml:"record"`
	} `xml:"readResponse"`
}

func (r *GetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetRequest) Do() (GetRequestResponseBody, error) {
	var err error

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
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}
