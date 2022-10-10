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

func (c *Client) NewCustomFieldTypeSearchRequest() CustomFieldTypeSearchRequest {
	r := CustomFieldTypeSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomFieldTypeSearchRequest struct {
	client      *Client
	queryParams *CustomFieldTypeSearchRequestQueryParams
	pathParams  *CustomFieldTypeSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomFieldTypeSearchRequestBody
}

func (r CustomFieldTypeSearchRequest) SOAPAction() string {
	return "search"
}

func (r CustomFieldTypeSearchRequest) NewQueryParams() *CustomFieldTypeSearchRequestQueryParams {
	return &CustomFieldTypeSearchRequestQueryParams{}
}

type CustomFieldTypeSearchRequestQueryParams struct {
}

func (p CustomFieldTypeSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomFieldTypeSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomFieldTypeSearchRequest) NewPathParams() *CustomFieldTypeSearchRequestPathParams {
	return &CustomFieldTypeSearchRequestPathParams{}
}

type CustomFieldTypeSearchRequestPathParams struct {
}

func (p *CustomFieldTypeSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomFieldTypeSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CustomFieldTypeSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomFieldTypeSearchRequest) Method() string {
	return r.method
}

func (r CustomFieldTypeSearchRequest) NewRequestBody() CustomFieldTypeSearchRequestBody {
	body := CustomFieldTypeSearchRequestBody{}
	body.SearchRecord.Type = "setupCustom:CustomFieldTypeSearch"
	return body
}

type CustomFieldTypeSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct {
		XMLName xml.Name `xml:"platformMsgs:searchRecord"`

		Type  string                     `xml:"xsi:type,attr"`
		Basic CustomFieldTypeSearchBasic `xml:"setupCustom:basic"`
	} `xml:"platformMsgs:searchRecord"`
}

func (r *CustomFieldTypeSearchRequest) RequestBody() *CustomFieldTypeSearchRequestBody {
	return &r.requestBody
}

func (r *CustomFieldTypeSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomFieldTypeSearchRequest) SetRequestBody(body CustomFieldTypeSearchRequestBody) {
	r.requestBody = body
}

func (r *CustomFieldTypeSearchRequest) NewResponseBody() *CustomFieldTypeSearchRequestResponseBody {
	return &CustomFieldTypeSearchRequestResponseBody{}
}

type CustomFieldTypeSearchRequestResponseBody struct {
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
		SearchID     string `xml:"searchId"`
		RecordList   struct {
			Record CustomFieldTypes `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *CustomFieldTypeSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *CustomFieldTypeSearchRequest) Do() (CustomFieldTypeSearchRequestResponseBody, error) {
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

type CustomFieldTypeSearchBasic struct {
}

func (c CustomFieldTypeSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomFieldTypeSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
