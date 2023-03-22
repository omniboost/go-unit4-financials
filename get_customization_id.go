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

func (c *Client) NewGetCustomizationIDRequest() GetCustomizationIDRequest {
	r := GetCustomizationIDRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCustomizationIDRequest struct {
	client      *Client
	queryParams *GetCustomizationIDRequestQueryParams
	pathParams  *GetCustomizationIDRequestPathParams
	method      string
	headers     http.Header
	requestBody GetCustomizationIDRequestBody
}

func (r GetCustomizationIDRequest) SOAPAction() string {
	return "getCustomizationId"
}

func (r GetCustomizationIDRequest) NewQueryParams() *GetCustomizationIDRequestQueryParams {
	return &GetCustomizationIDRequestQueryParams{}
}

type GetCustomizationIDRequestQueryParams struct {
}

func (p GetCustomizationIDRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomizationIDRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetCustomizationIDRequest) NewPathParams() *GetCustomizationIDRequestPathParams {
	return &GetCustomizationIDRequestPathParams{}
}

type GetCustomizationIDRequestPathParams struct {
}

func (p *GetCustomizationIDRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomizationIDRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetCustomizationIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomizationIDRequest) Method() string {
	return r.method
}

func (r GetCustomizationIDRequest) NewRequestBody() GetCustomizationIDRequestBody {
	return GetCustomizationIDRequestBody{}
}

type GetCustomizationIDRequestBody struct {
	XMLName xml.Name `xml:"getCustomizationId"`

	CustomizationType struct {
		GetCustomizationType string `xml:"getCustomizationType,attr,omitempty"`
	} `xml:"customizationType"`
	IncludeInactives bool `xml:"includeInactives"`
}

func (r *GetCustomizationIDRequest) RequestBody() *GetCustomizationIDRequestBody {
	return &r.requestBody
}

func (r *GetCustomizationIDRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetCustomizationIDRequest) SetRequestBody(body GetCustomizationIDRequestBody) {
	r.requestBody = body
}

func (r *GetCustomizationIDRequest) NewResponseBody() *GetCustomizationIDRequestResponseBody {
	return &GetCustomizationIDRequestResponseBody{}
}

type GetCustomizationIDRequestResponseBody struct {
	XMLName xml.Name `xml:"getCustomizationIdResponse"`

	GetCustomizationIdResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		Status       struct {
			Text      string `xml:",chardata"`
			IsSuccess string `xml:"isSuccess,attr"`
		} `xml:"status"`
		TotalRecords         string `xml:"totalRecords"`
		CustomizationRefList struct {
			CustomizationRef CustomizationRefs `xml:"customizationRef"`
		} `xml:"customizationRefList"`
	} `xml:"getCustomizationIdResult"`
}

func (r *GetCustomizationIDRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *GetCustomizationIDRequest) Do() (GetCustomizationIDRequestResponseBody, error) {
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

type GetCustomizationIDBasic struct {
	ExternalId       SearchMultiSelectField `xml:"externalId,omitempty"`
	ExternalIdString SearchStringField      `xml:"externalIdString,omitempty"`
	InternalId       SearchMultiSelectField `xml:"internalId,omitempty"`
	InternalIdNumber SearchLongField        `xml:"internalIdNumber,omitempty"`
	IsInactive       SearchBooleanField     `xml:"isInactive,omitempty"`
	Name             SearchStringField      `xml:"name,omitempty"`
	NameNoHierarchy  SearchStringField      `xml:"nameNoHierarchy,omitempty"`
	Subsidiary       SearchMultiSelectField `xml:"subsidiary,omitempty"`
	CustomFieldList  SearchCustomFieldList  `xml:"platformCommon:customFieldList,omitempty"`
}

func (c GetCustomizationIDBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c GetCustomizationIDBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
