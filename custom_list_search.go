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

func (c *Client) NewCustomListSearchRequest() CustomListSearchRequest {
	r := CustomListSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomListSearchRequest struct {
	client      *Client
	queryParams *CustomListSearchRequestQueryParams
	pathParams  *CustomListSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomListSearchRequestBody
}

func (r CustomListSearchRequest) SOAPAction() string {
	return "search"
}

func (r CustomListSearchRequest) NewQueryParams() *CustomListSearchRequestQueryParams {
	return &CustomListSearchRequestQueryParams{}
}

type CustomListSearchRequestQueryParams struct {
}

func (p CustomListSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomListSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomListSearchRequest) NewPathParams() *CustomListSearchRequestPathParams {
	return &CustomListSearchRequestPathParams{}
}

type CustomListSearchRequestPathParams struct {
}

func (p *CustomListSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomListSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CustomListSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomListSearchRequest) Method() string {
	return r.method
}

func (r CustomListSearchRequest) NewRequestBody() CustomListSearchRequestBody {
	body := CustomListSearchRequestBody{}
	body.SearchRecord.Type = "setupCustom:CustomListSearch"
	return body
}

type CustomListSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct {
		XMLName xml.Name `xml:"platformMsgs:searchRecord"`

		Type  string                `xml:"xsi:type,attr"`
		Basic CustomListSearchBasic `xml:"setupCustom:basic"`
	} `xml:"platformMsgs:searchRecord"`
}

func (r *CustomListSearchRequest) RequestBody() *CustomListSearchRequestBody {
	return &r.requestBody
}

func (r *CustomListSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomListSearchRequest) SetRequestBody(body CustomListSearchRequestBody) {
	r.requestBody = body
}

func (r *CustomListSearchRequest) NewResponseBody() *CustomListSearchRequestResponseBody {
	return &CustomListSearchRequestResponseBody{}
}

type CustomListSearchRequestResponseBody struct {
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
			Record CustomLists `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *CustomListSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *CustomListSearchRequest) Do() (CustomListSearchRequestResponseBody, error) {
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

type CustomListSearchBasic struct {
}

func (c CustomListSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomListSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
