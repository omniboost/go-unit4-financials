package financials

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewAuthenticateRequest() AuthenticateRequest {
	r := AuthenticateRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AuthenticateRequest struct {
	client      *Client
	queryParams *AuthenticateRequestQueryParams
	pathParams  *AuthenticateRequestPathParams
	method      string
	headers     http.Header
	requestBody AuthenticateRequestBody
}

func (r AuthenticateRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/wsadapter/authenticate/Authenticate"
}

func (r AuthenticateRequest) NewQueryParams() *AuthenticateRequestQueryParams {
	return &AuthenticateRequestQueryParams{}
}

type AuthenticateRequestQueryParams struct {
}

func (p AuthenticateRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AuthenticateRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r AuthenticateRequest) NewPathParams() *AuthenticateRequestPathParams {
	return &AuthenticateRequestPathParams{}
}

type AuthenticateRequestPathParams struct {
}

func (p *AuthenticateRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AuthenticateRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *AuthenticateRequest) SetMethod(method string) {
	r.method = method
}

func (r *AuthenticateRequest) Method() string {
	return r.method
}

func (r AuthenticateRequest) NewRequestBody() AuthenticateRequestBody {
	return AuthenticateRequestBody{}
}

type AuthenticateRequestBody struct {
	XMLName xml.Name `xml:"AuthenticateRequest"`

	SuppressDefaultUser string `xml:"SuppressDefaultUser"`
}

func (r *AuthenticateRequest) RequestBody() *AuthenticateRequestBody {
	return &r.requestBody
}

func (r *AuthenticateRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AuthenticateRequest) SetRequestBody(body AuthenticateRequestBody) {
	r.requestBody = body
}

func (r *AuthenticateRequest) NewResponseBody() *AuthenticateRequestResponseBody {
	return &AuthenticateRequestResponseBody{}
}

type AuthenticateRequestResponseBody struct {
	XMLName xml.Name `xml:"AuthenticateResponse"`

	User                   string `xml:"User"`
	UserTimestamp          string `xml:"UserTimestamp"`
	Company                string `xml:"Company"`
	Capability             string `xml:"Capability"`
	CapabilityTimestamp    string `xml:"CapabilityTimestamp"`
	PasswordExpiryWarning  string `xml:"PasswordExpiryWarning"`
	RecordLogoff           string `xml:"RecordLogoff"`
	FunctionalSecurityHash string `xml:"FunctionalSecurityHash"`
	Session                string `xml:"Session"`
}

func (r *AuthenticateRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/wsadapter/authenticate/authenticate-2.0", r.PathParams())
	return &u, err
}

func (r *AuthenticateRequest) Do() (AuthenticateRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.SetBasicAuth(fmt.Sprintf("%s\\%s", r.client.User(), r.client.CompanyCode()), r.client.Password())

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
