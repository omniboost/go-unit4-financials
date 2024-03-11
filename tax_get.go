package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewTaxGetRequest() TaxGetRequest {
	r := TaxGetRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxGetRequest struct {
	client      *Client
	queryParams *TaxGetRequestQueryParams
	pathParams  *TaxGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxGetRequestBody
}

func (r TaxGetRequest) SOAPAction() string {
	return "uri-coda-webservice/{version}/finance/Tax/Get"
}

func (r TaxGetRequest) SOAPHeader() SOAPHeader {
	session, _ := r.client.Session()
	return SOAPHeader{
		Options: Options{
			Session: session,
		},
	}
}

func (r TaxGetRequest) SOAPNS() []xml.Attr {
	return []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://www.coda.com/efinance/schemas/tax/tax-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://www.coda.com/efinance/schemas/tax/tax-14.0/webservice"},
		{Name: xml.Name{Space: "", Local: "xmlns:elem"}, Value: "http://www.coda.com/efinance/schemas/tax"},
	}
}

func (r TaxGetRequest) NewQueryParams() *TaxGetRequestQueryParams {
	return &TaxGetRequestQueryParams{}
}

type TaxGetRequestQueryParams struct {
}

func (p TaxGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TaxGetRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r TaxGetRequest) NewPathParams() *TaxGetRequestPathParams {
	return &TaxGetRequestPathParams{}
}

type TaxGetRequestPathParams struct {
}

func (p *TaxGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxGetRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *TaxGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxGetRequest) Method() string {
	return r.method
}

func (r TaxGetRequest) NewRequestBody() TaxGetRequestBody {
	return TaxGetRequestBody{}
}

type TaxGetRequestBody struct {
	XMLName xml.Name `xml:"GetRequest"`
	Key     struct {
		CmpCode string `xml:"elem:CmpCode"`
		Code    string `xml:"elem:Code,omitempty"`
	} `xml:"Key"`
}

func (r *TaxGetRequest) RequestBody() *TaxGetRequestBody {
	return &r.requestBody
}

func (r *TaxGetRequest) SOAPBodyInterface() interface{} {
	return &r.requestBody
}

func (r *TaxGetRequest) SetRequestBody(body TaxGetRequestBody) {
	r.requestBody = body
}

func (r *TaxGetRequest) NewResponseBody() *TaxGetRequestResponseBody {
	return &TaxGetRequestResponseBody{}
}

type TaxGetRequestResponseBody struct {
	XMLName xml.Name `xml:"GetResponse"`

	Text       string `xml:",chardata"`
	Webservice string `xml:"webservice,attr"`
	Xmlns      string `xml:"xmlns,attr"`
	Com        string `xml:"com,attr"`
	CmpCode    string `xml:"CmpCode"`
	Code       string `xml:"Code"`
	Tax        struct {
		TimeStamp           string `xml:"TimeStamp"`
		CmpCode             string `xml:"CmpCode"`
		Code                string `xml:"Code"`
		Name                string `xml:"Name"`
		ShortName           string `xml:"ShortName"`
		RecoveryScope       string `xml:"RecoveryScope"`
		Rev                 string `xml:"Rev"`
		RecoverAcc          string `xml:"RecoverAcc"`
		RecoverRevAcc       string `xml:"RecoverRevAcc"`
		IrrecoverToGoods    string `xml:"IrrecoverToGoods"`
		IrrecoverAcc        string `xml:"IrrecoverAcc"`
		IrrecoverRevToGoods string `xml:"IrrecoverRevToGoods"`
		IrrecoverRevAcc     string `xml:"IrrecoverRevAcc"`
		DestCode            string `xml:"DestCode"`
		Intercompany        string `xml:"Intercompany"`
		IsDeferredVAT       string `xml:"IsDeferredVAT"`
		DeferredAccount     string `xml:"DeferredAccount"`
		CollectionAccount   string `xml:"CollectionAccount"`
		RateInfoList        struct {
			RateInfo struct {
				EffectiveDate      string `xml:"EffectiveDate"`
				Rate               string `xml:"Rate"`
				RecoveryPercentage string `xml:"RecoveryPercentage"`
			} `xml:"RateInfo"`
		} `xml:"RateInfoList"`
	} `xml:"Tax"`
}

func (r *TaxGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/services/finance/tax/tax-14.0", r.PathParams())
	return &u, err
}

func (r *TaxGetRequest) Do() (TaxGetRequestResponseBody, error) {
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
