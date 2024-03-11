package financials

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-unit4-financials/omitempty"
)

type ElementFilterListFilter struct {
	MaxKeys   int    `xml:"com:MaxKeys"`   // The maximum number of items to be returned. The value zero has the special meaning of 'no limit'.
	CmpCode   string `xml:"com:CmpCode"`   // The company code.
	ShortName string `xml:"com:ShortName"` // The short name.
	Level     int    `xml:"elem:Level"`    // The element level to which the element filter master applies.
}

type ElmReqFullKeys struct {
	MaxKeys     int `xml:"MaxKeys"`
	FullWildKey struct {
		CmpCode string `xml:"CmpCode"`
		Level   int    `xml:"Level"`
		Code    string `xml:"Code"`
	} `xml:"FullWildKey"`
}

type Transaction struct {
	Header Header `xml:"Header"`
	Lines  Lines  `xml:"Lines>Line"`
}

type Header struct {
	Key struct {
		CmpCode string `xml:"CmpCode"`
		Code    string `xml:"Code"`
	} `xml:"Key"`
	Period          string   `xml:"Period"`
	CurCode         string   `xml:"CurCode"`
	Date            DateTime `xml:"Date"`
	Description     string   `xml:"Description"`
	OriginalCompany string   `xml:"OriginalCompany"`
}

func (h Header) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(h, e, start)
}

type Lines []Line

type Line struct {
	Number           int      `xml:"Number,omitempty"`
	AccountCode      string   `xml:"AccountCode"`
	BalancingAccount string   `xml:"BalancingAccount,omitempty"`
	DocValue         Decimal  `xml:"DocValue"`
	LineType         string   `xml:"LineType"`
	LineSense        string   `xml:"LineSense"`
	Description      string   `xml:"Description"`
	DueDate          DateTime `xml:"DueDate,omitempty"`
	ExtRef1          string   `xml:"ExtRef1,omitempty"`
	ExtRef2          string   `xml:"ExtRef2,omitempty"`
	ExtRef3          string   `xml:"ExtRef3,omitempty"`
	ExtRef4          string   `xml:"ExtRef4,omitempty"`
	ExtRef5          string   `xml:"ExtRef5,omitempty"`
	ExtRef6          string   `xml:"ExtRef6,omitempty"`
	Taxes            Taxes    `xml:"Taxes>Tax,omitempty"`
	TaxLineCode      string   `xml:"TaxLineCode,omitempty"`
	DocTaxTurnover   Decimal  `xml:"DocTaxTurnover,omitempty"`
	DocSumTax        Decimal  `xml:"DocSumTax,omitempty"`
}

func (l Line) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(l, e, start)
}

type Taxes []Tax

func (t Taxes) IsEmpty() bool {
	return len(t) == 0
}

type Tax struct {
	Code  string  `xml:"Code"`
	Value Decimal `xml:"Value"`
}

func (t Tax) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(t, e, start)
}

func (t Tax) IsEmpty() bool {
	return zero.IsZero(t)
}

type Element struct {
	XMLName          xml.Name  `xml:"Element"`
	CmpCode          string    `xml:"CmpCode"`
	Level            int       `xml:"Level"`
	Code             string    `xml:"Code"`
	Name             string    `xml:"Name"`
	ShortName        string    `xml:"ShortName,omitempty"`
	Matchable        bool      `xml:"Matchable"`
	AccountType      string    `xml:"AccountType"`
	CustomerSupplier bool      `xml:"CustomerSupplier"`
	IsCustomer       bool      `xml:"IsCustomer"`
	IsSupplier       bool      `xml:"IsSupplier"`
	Terms            string    `xml:"Terms"`
	PayStatus        string    `xml:"PayStatus"`
	Addresses        Addresses `xml:"Addresses>Address"`
}

type Addresses []Address

type Address struct {
	DefaultAddress bool   `xml:"DefaultAddress"`
	Name           string `xml:"Name"`
	Address1       string `xml:"Address1,omitempty"`
	Address2       string `xml:"Address2,omitempty"`
	Address3       string `xml:"Address3,omitempty"`
	Address4       string `xml:"Address4,omitempty"`
	Address5       string `xml:"Address5,omitempty"`
	Address6       string `xml:"Address6,omitempty"`
	PostCode       string `xml:"PostCode"`
	Category       string `xml:"Category"`
	Country        string `xml:"Country"`
}

type DocSeletKey struct {
	MaxKeys   int    `xml:"com:MaxKeys"`   // The maximum number of items to be returned. The value zero has the special meaning of 'no limit'.
	ShortName string `xml:"com:ShortName"` // The short name.
}

type ReqKeys struct {
	MaxKeys int `xml:"com:MaxKeys"` // The maximum number of items to be returned. The value zero has the special meaning of 'no limit'.
	Key     struct {
		CmpCode string `xml:"com:CmpCode,omitepty"`
		Code    string `xml:"com:Code,omitempty"`
	} `xml:"com:Key"`
}

type TaxMaster struct {
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
}
