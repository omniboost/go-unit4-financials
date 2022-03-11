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

func (c *Client) NewSalesTaxItemSearchRequest() SalesTaxItemSearchRequest {
	r := SalesTaxItemSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesTaxItemSearchRequest struct {
	client      *Client
	queryParams *SalesTaxItemSearchRequestQueryParams
	pathParams  *SalesTaxItemSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesTaxItemSearchRequestBody
}

func (r SalesTaxItemSearchRequest) SOAPAction() string {
	return "search"
}

func (r SalesTaxItemSearchRequest) NewQueryParams() *SalesTaxItemSearchRequestQueryParams {
	return &SalesTaxItemSearchRequestQueryParams{}
}

type SalesTaxItemSearchRequestQueryParams struct {
}

func (p SalesTaxItemSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SalesTaxItemSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SalesTaxItemSearchRequest) NewPathParams() *SalesTaxItemSearchRequestPathParams {
	return &SalesTaxItemSearchRequestPathParams{}
}

type SalesTaxItemSearchRequestPathParams struct {
}

func (p *SalesTaxItemSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesTaxItemSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SalesTaxItemSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesTaxItemSearchRequest) Method() string {
	return r.method
}

func (r SalesTaxItemSearchRequest) NewRequestBody() SalesTaxItemSearchRequestBody {
	return SalesTaxItemSearchRequestBody{
		SearchRecord: SalesTaxItemSearchRecordBasic{
			Type: "listAcct:SalesTaxItemSearch",
		},
	}
}

type SalesTaxItemSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord SalesTaxItemSearchRecordBasic `xml:"platformMsgs:searchRecord"`
}

func (r *SalesTaxItemSearchRequest) RequestBody() *SalesTaxItemSearchRequestBody {
	return &r.requestBody
}

func (r *SalesTaxItemSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SalesTaxItemSearchRequest) SetRequestBody(body SalesTaxItemSearchRequestBody) {
	r.requestBody = body
}

func (r *SalesTaxItemSearchRequest) NewResponseBody() *SalesTaxItemSearchRequestResponseBody {
	return &SalesTaxItemSearchRequestResponseBody{}
}

type SalesTaxItemSearchRequestResponseBody struct {
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
			Record SalesTaxItems `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *SalesTaxItemSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *SalesTaxItemSearchRequest) Do() (SalesTaxItemSearchRequestResponseBody, error) {
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

type SalesTaxItemSearchBasic struct {
	Description      SearchStringField      `xml:"platformCommon:description,omitempty"`
	ExternalID       SearchMultiSelectField `xml:"platformCommon:externalId,omitempty"`
	ExternalIDString SearchStringField      `xml:"platformCommon:externalIdString,omitempty"`
	InternalID       SearchMultiSelectField `xml:"platformCommon:internalId,omitempty"`
	InternalIDNumber SearchLongField        `xml:"platformCommon:internalIdNumber,omitempty"`
	IsInactive       SearchBooleanField     `xml:"platformCommon:isInactive,omitempty"`
	ItemID           SearchStringField      `xml:"platformCommon:itemId,omitempty"`
	Name             SearchStringField      `xml:"platformCommon:name,omitempty"`
	TaxType          SearchMultiSelectField `xml:"platformCommon:taxType,omitempty"`
}

func (c SalesTaxItemSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c SalesTaxItemSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
