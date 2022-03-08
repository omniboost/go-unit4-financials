package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewContactSearchRequest() ContactSearchRequest {
	r := ContactSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContactSearchRequest struct {
	client      *Client
	queryParams *ContactSearchRequestQueryParams
	pathParams  *ContactSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody ContactSearchRequestBody
}

func (r ContactSearchRequest) SOAPAction() string {
	return "search"
}

func (r ContactSearchRequest) NewQueryParams() *ContactSearchRequestQueryParams {
	return &ContactSearchRequestQueryParams{}
}

type ContactSearchRequestQueryParams struct {
}

func (p ContactSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ContactSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ContactSearchRequest) NewPathParams() *ContactSearchRequestPathParams {
	return &ContactSearchRequestPathParams{}
}

type ContactSearchRequestPathParams struct {
}

func (p *ContactSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContactSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ContactSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *ContactSearchRequest) Method() string {
	return r.method
}

func (r ContactSearchRequest) NewRequestBody() ContactSearchRequestBody {
	return ContactSearchRequestBody{}
}

type ContactSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct { // SearchRecordBasic
		Type  string             `xml:"xsi:type,attr"`
		Basic ContactSearchBasic `xml:"listRel:basic"`
	} `xml:"platformMsgs:searchRecord"`
}

func (r *ContactSearchRequest) RequestBody() *ContactSearchRequestBody {
	return &r.requestBody
}

func (r *ContactSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ContactSearchRequest) SetRequestBody(body ContactSearchRequestBody) {
	r.requestBody = body
}

func (r *ContactSearchRequest) NewResponseBody() *ContactSearchRequestResponseBody {
	return &ContactSearchRequestResponseBody{}
}

type ContactSearchRequestResponseBody struct {
	XMLName xml.Name `xml:"searchResponse"`

	SearchResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		Status       struct {
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		TotalRecords string `xml:"totalRecords"`
		PageSize     string `xml:"pageSize"`
		TotalPages   string `xml:"totalPages"`
		PageIndex    string `xml:"pageIndex"`
		SearchId     string `xml:"searchId"`
		RecordList   struct {
			Record Customers `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *ContactSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *ContactSearchRequest) Do() (ContactSearchRequestResponseBody, error) {
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

type ContactSearchBasic struct {
	CompanyName SearchStringField `xml:"platformCommon:companyName,omitempty"`
	Email       SearchStringField `xml:"platformCommon:email,omitempty"`
}

func (c ContactSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c ContactSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
