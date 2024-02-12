package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewInputPostRequest() InputPostRequest {
	r := InputPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InputPostRequest struct {
	client      *Client
	queryParams *InputPostRequestQueryParams
	pathParams  *InputPostRequestPathParams
	method      string
	headers     http.Header
	requestBody InputPostRequestBody
}

func (r InputPostRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/Input/Post"
}

func (r InputPostRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r InputPostRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/inputext/input-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/inputext/input-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:inp"}, Value: "http://www.coda.com/efinance/schemas/inputext"},
		{Name: xml.Name{Space: "", Local: "xmlns:tran"}, Value: "http://www.coda.com/efinance/schemas/transaction"},
	}
}

func (r InputPostRequest) NewQueryParams() *InputPostRequestQueryParams {
	return &InputPostRequestQueryParams{}
}

type InputPostRequestQueryParams struct {
}

func (p InputPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InputPostRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r InputPostRequest) NewPathParams() *InputPostRequestPathParams {
	return &InputPostRequestPathParams{}
}

type InputPostRequestPathParams struct {
}

func (p *InputPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InputPostRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *InputPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *InputPostRequest) Method() string {
	return r.method
}

func (r InputPostRequest) NewRequestBody() InputPostRequestBody {
	return InputPostRequestBody{}
}

type InputPostRequestBody struct {
	XMLName xml.Name `xml:"web:PostRequest"`

	PostOptions struct {
		Postto string `xml:"postto,attr"`
	} `xml:"web:PostOptions"`
	Transaction Transaction `xml:"web:Transaction"`
}

func (r *InputPostRequest) RequestBody() *InputPostRequestBody {
	return &r.requestBody
}

func (r *InputPostRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InputPostRequest) SetRequestBody(body InputPostRequestBody) {
	r.requestBody = body
}

func (r *InputPostRequest) NewResponseBody() *InputPostRequestResponseBody {
	return &InputPostRequestResponseBody{}
}

type InputPostRequestResponseBody struct {
	XMLName xml.Name `xml:"GetEnvironmentResponse"`
}

func (r *InputPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/inputext/input-14.0", r.PathParams())
	return &u, err
}

func (r *InputPostRequest) Do() (InputPostRequestResponseBody, error) {
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
