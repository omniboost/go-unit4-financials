package netsuite

import "net/url"

type Request interface {
	Method() string
	// QueryParams() QueryParams
	PathParams() PathParams
	RequestBodyInterface() interface{}
	URL() (*url.URL, error)
	SOAPAction() string
}

type QueryParams interface {
	ToURLValues() (url.Values, error)
}

type PathParams interface {
	Params() map[string]string
}
