package financials

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/elliotchance/pie/pie"
	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-unit4-financials/" + libraryVersion
	mediaType      = "text/xml"
	charset        = "utf-8"
)

var (
	BaseURL string = "https://webservices.financials.com/services/NetSuitePort_2022_2"
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL string

	// credentials
	user        string
	password    string
	companyCode string
	locale      string
	session     string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// func (c Client) ApplicationID() string {
// 	return c.applicationID
// }

// func (c *Client) SetApplicationID(applicationID string) {
// 	c.applicationID = applicationID
// }

func (c Client) User() string {
	return c.user
}

func (c *Client) SetUser(user string) {
	c.user = user
}

func (c Client) Password() string {
	return c.password
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c Client) CompanyCode() string {
	return c.companyCode
}

func (c *Client) SetCompanyCode(companyCode string) {
	c.companyCode = companyCode
}

func (c Client) Locale() string {
	return c.locale
}

func (c *Client) SetLocale(locale string) {
	c.locale = locale
}

func (c Client) BaseURL() (*url.URL, error) {
	tmpl, err := template.New("host").Parse(c.baseURL)
	if err != nil {
		return &url.URL{}, err
	}
	buf := new(bytes.Buffer)
	// err = tmpl.Execute(buf, map[string]interface{}{"account_id": c.companyID})
	err = tmpl.Execute(buf, map[string]interface{}{})
	if err != nil {
		return &url.URL{}, err
	}
	return url.Parse(buf.String())
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) (url.URL, error) {
	clientURL, err := c.BaseURL()
	if err != nil {
		return url.URL{}, err
	}

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		return url.URL{}, err
	}

	clientURL.Path = buf.String()
	return *clientURL, nil
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to xml
	buf := new(bytes.Buffer)
	if req.SOAPBodyInterface() != nil {
		soapRequest := NewRequestEnvelope()
		if x, ok := req.(interface{ SOAPNS() []xml.Attr }); ok {
			soapRequest.NS = append(soapRequest.NS, x.SOAPNS()...)
		}
		soapRequest.Header = req.SOAPHeader()
		soapRequest.Header.Options.Locale = c.Locale()
		soapRequest.Body.ActionBody = req.SOAPBodyInterface()

		enc := xml.NewEncoder(buf)
		enc.Indent("", "  ")
		err := enc.Encode(soapRequest)
		if err != nil {
			return nil, err
		}

		err = enc.Flush()
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	u, err := req.URL()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(req.Method(), u.String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("SOAPAction", req.SOAPAction())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is xml decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	soapResponse := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body: Body{
			ActionBody: body,
		},
	}

	soapError := SOAPError{Response: httpResp}
	errResp := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body: Body{
			ActionBody: &soapError,
		},
	}

	soapFault := SOAPFault{Response: httpResp}
	faultResp := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body: Body{
			ActionBody: &soapFault,
		},
	}

	statusResponseBody := StatusResponseBody{Response: httpResp}
	statusResp := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body: Body{
			ActionBody: &statusResponseBody,
		},
	}

	err = c.Unmarshal(httpResp.Body, soapResponse, errResp, faultResp, statusResp)
	if err != nil {
		return httpResp, err
	}

	// if statusResponseBody.Node.Status.Error() != "" {
	// 	return httpResp, statusResponseBody.Node.Status
	// }

	if soapError.Error() != "" {
		return httpResp, soapError
	}

	if soapFault.Error() != "" {
		return httpResp, soapFault
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := xml.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			log.Println(e)
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

func (c *Client) Session() (string, error) {
	// fetch a new token if it isn't set already
	if c.session == "" {
		var err error
		c.session, err = c.NewSession()
		if err != nil {
			return "", err
		}
	}

	return c.session, nil
}

func (c *Client) NewSession() (string, error) {
	req := c.NewAuthenticateRequest()
	resp, err := req.Do()
	if err != nil {
		return "", err
	}

	return resp.Session, nil

}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a xml response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert xml to struct
	if len(data) != 0 {
		err = xml.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("%s: Expected Content-Type \"%s\", got \"%s\"", response.Status, mediaType, contentType)
	}

	return nil
}

// {
//   "type": "https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html#sec10.4.4",
//   "title": "Forbidden",
//   "status": 403,
//   "o:errorDetails": [
//     {
//       "detail": "The account record is only available as a beta record. Enable the REST Record Service (Beta) feature in Setup > Company > Enable Features to work with this record.",
//       "o:errorCode": "INSUFFICIENT_PERMISSION"
//     }
//   ]
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
	Err      string
}

func (r *ErrorResponse) Error() string {
	return r.Err
}

type SOAPError struct {
	// HTTP response that caused this error
	Response *http.Response
}

func (e SOAPError) Error() string {
	return ""
}

type SOAPFault struct {
	// HTTP response that caused this error
	Response *http.Response

	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
	Detail      struct {
		Reason struct {
			Text string `xml:"Text"`
			Path string `xml:"Path"`
		} `xml:"Reason"`
	} `xml:"detail"`
}

func (f SOAPFault) Error() string {
	l := []string{f.Faultcode, f.Faultstring, f.Detail.Reason.Text, f.Detail.Reason.Path}
	ll := []string{}
	for _, v := range l {
		if v != "" {
			ll = append(ll, v)
		}
	}

	ll = pie.Strings(ll).Unique()

	return strings.Join(ll, ", ")
}

type StatusResponseBody struct {
	// HTTP response that caused this error
	Response *http.Response

	Node struct {
	} `xml:",any"`
}
