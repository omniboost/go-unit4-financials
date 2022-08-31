package netsuite

import (
	"encoding/xml"
	"strings"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
)

type RecordRef struct {
	InternalID string `xml:"internalId,attr,omitempty"`
	ExternalID string `xml:"externalId,attr,omitempty"`
	Name       string `xml:"name,omitempty"`
}

func (r RecordRef) IsEmpty() bool {
	return zero.IsZero(r)
}

type Customers []Customer

type Customer struct {
	InternalID           string    `xml:"internalId,attr,omitempty"`
	ExternalID           string    `xml:"externalId,attr,omitempty"`
	Type                 string    `xml:"type,attr,omitempty"`
	ListRel              string    `xml:"listRel,attr,omitempty"`
	EntityID             string    `xml:"entityId,omitempty"`
	IsPerson             bool      `xml:"isPerson,omitempty"`
	CompanyName          string    `xml:"companyName,omitempty"`
	EntityStatus         RecordRef `xml:"entityStatus,omitempty"`
	Phone                string    `xml:"phone,omitempty"`
	Email                string    `xml:"email,omitempty"`
	DefaultAddress       string    `xml:"defaultAddress,omitempty"`
	IsInactive           string    `xml:"isInactive,omitempty"`
	Language             string    `xml:"language,omitempty"`
	DateCreated          Date      `xml:"dateCreated,omitempty"`
	Subsidiary           RecordRef `xml:"subsidiary,omitempty"`
	VatRegNumber         string    `xml:"vatRegNumber,omitempty"`
	Terms                RecordRef `xml:"terms,omitempty"`
	CreditLimit          string    `xml:"creditLimit,omitempty"`
	UnbilledOrders       string    `xml:"unbilledOrders,omitempty"`
	Currency             RecordRef `xml:"currency,omitempty"`
	ShipComplete         string    `xml:"shipComplete,omitempty"`
	AlcoholRecipientType string    `xml:"alcoholRecipientType,omitempty"`
	ReceivablesAccount   RecordRef `xml:"receivablesAccount,omitempty"`
	LastModifiedDate     DateTime  `xml:"lastModifiedDate,omitempty"`
	Stage                string    `xml:"stage,omitempty"`
	IsBudgetApproved     string    `xml:"isBudgetApproved,omitempty"`
	FirstName            string    `xml:"firstName,omitempty"`
	LastName             string    `xml:"lastName,omitempty"`
	CustomFieldList      struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

func (c Customer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c Customer) IsEmpty() bool {
	return zero.IsZero(c)
}

type CustomFields []CustomField

type CustomField struct {
	InternalID string           `xml:"internalId,attr,omitempty"`
	ScriptID   string           `xml:"scriptId,attr,omitempty"`
	Type       string           `xml:"xsi:type,attr,omitempty"`
	Value      CustomFieldValue `xml:"value,omitempty"`
}

type Accounts []Account

type Account struct {
	InternalID      string    `xml:"internalId,attr"`
	ExternalID      string    `xml:"externalId,attr"`
	Type            string    `xml:"type,attr"`
	ListAcct        string    `xml:"listAcct,attr"`
	AcctType        string    `xml:"acctType"`
	AcctNumber      string    `xml:"acctNumber"`
	AcctName        string    `xml:"acctName"`
	IncludeChildren string    `xml:"includeChildren"`
	GeneralRate     string    `xml:"generalRate"`
	Parent          RecordRef `xml:"parent"`
	CashFlowRate    string    `xml:"cashFlowRate"`
	IsInactive      string    `xml:"isInactive"`
	Inventory       string    `xml:"inventory"`
	Eliminate       string    `xml:"eliminate"`
	Revalue         string    `xml:"revalue"`
	CustomFieldList struct {
		CustomFields CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
	Currency RecordRef `xml:"currency"`
}

type Departments []Department

type Department struct {
	InternalID      string    `xml:"internalId,attr"`
	Type            string    `xml:"type,attr"`
	ListAcct        string    `xml:"listAcct,attr"`
	Name            string    `xml:"name"`
	IncludeChildren string    `xml:"includeChildren"`
	IsInactive      string    `xml:"isInactive"`
	Parent          RecordRef `xml:"parent"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type Record struct {
	Type   string `xml:"xsi:type,attr"`
	Record interface{}
}

func (r Record) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: r.Type})
	return omitempty.MarshalXML(r.Record, e, start)
}

type Classifications []Classification

type Classification struct {
	InternalID      string `xml:"internalId,attr"`
	Type            string `xml:"type,attr"`
	ListAcct        string `xml:"listAcct,attr"`
	Name            string `xml:"name"`
	IncludeChildren string `xml:"includeChildren"`
	IsInactive      string `xml:"isInactive"`
	CustomFieldList struct {
		CustomFields CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type JournalEntry struct {
	TranGeneral      string            `xml:"tranGeneral,attr,omitempty"`
	InternalId       string            `xml:"internalId,attr,omitempty"`
	Type             string            `xml:"type,attr,omitempty"`
	TranDate         Date              `xml:"tranDate,omitempty"`
	Currency         RecordRef         `xml:"currency,omitempty"`
	ExchangeRate     string            `xml:"exchangeRate,omitempty"`
	TranId           string            `xml:"tranId,omitempty"`
	Subsidiary       RecordRef         `xml:"subsidiary,omitempty"`
	CreatedDate      Date              `xml:"createdDate,omitempty"`
	LastModifiedDate Date              `xml:"lastModifiedDate,omitempty"`
	LineList         JournalEntryLines `xml:"lineList>line"`
	Memo             string            `xml:"memo"`
	CustomFieldList  struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
}

func (j JournalEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j JournalEntry) IsEmpty() bool {
	return zero.IsZero(j)
}

// type JournalEntryLineList struct {
// 	Line JournalEntryLines
// }

// func (j JournalEntryLineList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	return omitempty.MarshalXML(j, e, start)
// }

// func (j JournalEntryLineList) IsEmpty() bool {
// 	return zero.IsZero(j)
// }

type JournalEntryLines []JournalEntryLine

type JournalEntryLine struct {
	XMLName xml.Name `xml:"line"`

	Account         RecordRef `xml:"account"`
	Line            int       `xml:"line,omitempty"`
	Debit           float64   `xml:"debit,omitempty"`
	Memo            string    `xml:"memo"`
	Eliminate       string    `xml:"eliminate,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	Credit     float64   `xml:"credit,omitempty"`
	Department RecordRef `xml:"department,omitempty"`
	TaxCode    RecordRef `xml:"taxCode,omitempty"`
	Tax1Amt    float64   `xml:"tax1Amt,omitempty"`
	Tax1Acct   RecordRef `xml:"tax1Acct,omitempty"`
}

func (j JournalEntryLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j JournalEntryLine) IsEmpty() bool {
	return zero.IsZero(j)
}

type SalesTaxItems []SalesTaxItem

type SalesTaxItem struct {
	ListAcct       string    `xml:"listAcct,attr"`
	InternalID     string    `xml:"internalId,attr"`
	Type           string    `xml:"type,attr"`
	ItemID         string    `xml:"itemId"`
	Description    string    `xml:"description"`
	Rate           string    `xml:"rate"`
	TaxType        RecordRef `xml:"taxType"`
	TaxAgency      RecordRef `xml:"taxAgency"`
	IsInactive     string    `xml:"isInactive"`
	SubsidiaryList struct {
		RecordRef RecordRef `xml:"recordRef"`
	} `xml:"subsidiaryList"`
	IncludeChildren       string `xml:"includeChildren"`
	Eccode                string `xml:"eccode"`
	ReverseCharge         string `xml:"reverseCharge"`
	Service               string `xml:"service"`
	Exempt                string `xml:"exempt"`
	IsDefault             string `xml:"isDefault"`
	ExcludeFromTaxReports string `xml:"excludeFromTaxReports"`
	Available             string `xml:"available"`
	Export                string `xml:"export"`
	CustomFieldList       struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	Parent RecordRef `xml:"parent"`
}

type Invoice struct {
	TranSales             string          `xml:"tranSales,attr,omitempty"`
	InternalId            string          `xml:"internalId,attr,omitempty"`
	Type                  string          `xml:"type,attr,omitempty"`
	ExternalId            string          `xml:"externalId,attr,omitempty"`
	CreatedDate           Date            `xml:"createdDate,omitempty"`
	LastModifiedDate      Date            `xml:"lastModifiedDate,omitempty"`
	Entity                RecordRef       `xml:"entity,omitempty"`
	TranDate              Date            `xml:"tranDate,omitempty"`
	TranId                string          `xml:"tranId,omitempty"`
	Department            RecordRef       `xml:"department,omitempty"`
	Terms                 RecordRef       `xml:"terms,omitempty"`
	Subsidiary            RecordRef       `xml:"subsidiary,omitempty"`
	Currency              RecordRef       `xml:"currency,omitempty"`
	DueDate               Date            `xml:"dueDate,omitempty"`
	OtherRefNum           string          `xml:"otherRefNum,omitempty"`
	Memo                  string          `xml:"memo,omitempty"`
	TotalCostEstimate     string          `xml:"totalCostEstimate,omitempty"`
	EstGrossProfit        string          `xml:"estGrossProfit,omitempty"`
	EstGrossProfitPercent string          `xml:"estGrossProfitPercent,omitempty"`
	Account               RecordRef       `xml:"account,omitempty"`
	ExchangeRate          string          `xml:"exchangeRate,omitempty"`
	CurrencyName          string          `xml:"currencyName,omitempty"`
	BillingAddress        Address         `xml:"billingAddress,omitempty"`
	ShippingAddress       Address         `xml:"shippingAddress,omitempty"`
	ShipIsResidential     string          `xml:"shipIsResidential,omitempty"`
	SubTotal              string          `xml:"subTotal,omitempty"`
	CanHaveStackable      string          `xml:"canHaveStackable,omitempty"`
	TaxTotal              string          `xml:"taxTotal,omitempty"`
	Total                 string          `xml:"total,omitempty"`
	Status                string          `xml:"status,omitempty"`
	Email                 string          `xml:"email,omitempty"`
	VatRegNum             string          `xml:"vatRegNum,omitempty"`
	ItemCostDiscPrint     string          `xml:"itemCostDiscPrint"`
	ExpCostDiscPrint      string          `xml:"expCostDiscPrint"`
	ItemList              InvoiceItemList `xml:"itemList>item"`
	OverrideInstallments  string          `xml:"overrideInstallments"`
	CustomFieldList       struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type Transactions []Transaction

type Transaction struct{}

type Address struct {
	PlatformCommon  string `xml:"platformCommon,attr"`
	InternalId      string `xml:"internalId"`
	Country         string `xml:"country"`
	Attention       string `xml:"attention"`
	Addressee       string `xml:"addressee"`
	Addr1           string `xml:"addr1"`
	City            string `xml:"city"`
	Zip             string `xml:"zip"`
	AddrText        string `xml:"addrText"`
	Override        string `xml:"override"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
	Addr2 string `xml:"addr2"`
}

type InvoiceItemList []InvoiceItem

type InvoiceItem struct {
	Item            RecordRef `xml:"item,omitempty"`
	Line            string    `xml:"line,omitempty"`
	Description     string    `xml:"description,omitempty"`
	Amount          float64   `xml:"amount,omitempty"`
	Quantity        string    `xml:"quantity,omitempty"`
	Price           RecordRef `xml:"price,omitempty"`
	Rate            float64   `xml:"rate,omitempty"`
	Department      RecordRef `xml:"department,omitempty"`
	GrossAmt        float64   `xml:"grossAmt,omitempty"`
	Tax1Amt         float64   `xml:"tax1Amt,omitempty"`
	Tax1Acct        RecordRef `xml:"tax1Acct,omitempty"`
	TaxCode         RecordRef `xml:"taxCode,omitempty"`
	TaxRate1        float64   `xml:"taxRate1,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
}

func (i InvoiceItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

func (i InvoiceItem) IsEmpty() bool {
	return zero.IsZero(i)
}

type Status struct {
	IsSuccess    bool `xml:"isSuccess,attr"`
	StatusDetail struct {
		Type    string `xml:"type,attr"`
		Code    string `xml:"code"`
		Message string `xml:"message"`
	} `xml:"statusDetail"`
}

func (s Status) Error() string {
	if s.IsSuccess == false && s.StatusDetail.Message != "" {
		s := []string{s.StatusDetail.Type, s.StatusDetail.Code, s.StatusDetail.Message}
		return strings.Join(s, ", ")
	}

	return ""
}

type CustomFieldValue struct {
	InternalID string `xml:"internalId,attr,omitempty"`
	TypeID     string `xml:"typeId,attr,omitempty"`
	Name       string `xml:"name,omitempty"`
	Text       string `xml:",chardata"`
}

func (c CustomFieldValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomFieldValue) IsEmpty() bool {
	return zero.IsZero(c)
}

type CreditMemo struct {
	TranSales             string    `xml:"tranSales,attr,omitempty"`
	InternalId            string    `xml:"internalId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"`
	ExternalId            string    `xml:"externalId,attr,omitempty"`
	CreatedDate           Date      `xml:"createdDate,omitempty"`
	LastModifiedDate      Date      `xml:"lastModifiedDate,omitempty"`
	Entity                RecordRef `xml:"entity,omitempty"`
	TranDate              Date      `xml:"tranDate,omitempty"`
	TranId                string    `xml:"tranId,omitempty"`
	Department            RecordRef `xml:"department,omitempty"`
	Terms                 RecordRef `xml:"terms,omitempty"`
	Subsidiary            RecordRef `xml:"subsidiary,omitempty"`
	Currency              RecordRef `xml:"currency,omitempty"`
	DueDate               Date      `xml:"dueDate,omitempty"`
	OtherRefNum           string    `xml:"otherRefNum,omitempty"`
	Memo                  string    `xml:"memo,omitempty"`
	TotalCostEstimate     string    `xml:"totalCostEstimate,omitempty"`
	EstGrossProfit        string    `xml:"estGrossProfit,omitempty"`
	EstGrossProfitPercent string    `xml:"estGrossProfitPercent,omitempty"`
	Account               RecordRef `xml:"account,omitempty"`
	ExchangeRate          string    `xml:"exchangeRate,omitempty"`
	CurrencyName          string    `xml:"currencyName,omitempty"`
	BillingAddress        Address   `xml:"billingAddress,omitempty"`
	ShippingAddress       Address   `xml:"shippingAddress,omitempty"`
	ShipIsResidential     string    `xml:"shipIsResidential,omitempty"`
	SubTotal              string    `xml:"subTotal,omitempty"`
	CanHaveStackable      string    `xml:"canHaveStackable,omitempty"`
	TaxTotal              string    `xml:"taxTotal,omitempty"`
	Total                 string    `xml:"total,omitempty"`
	Status                string    `xml:"status,omitempty"`
	Email                 string    `xml:"email,omitempty"`
	VatRegNum             string    `xml:"vatRegNum,omitempty"`
	// ItemCostDiscPrint     string         `xml:"itemCostDiscPrint"`
	// ExpCostDiscPrint     string             `xml:"expCostDiscPrint"`
	ItemList CreditMemoItemList `xml:"itemList>item"`
	// OverrideInstallments string             `xml:"overrideInstallments"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type CreditMemoItemList []CreditMemoItem

type CreditMemoItem struct {
	Item            RecordRef `xml:"item,omitempty"`
	Line            string    `xml:"line,omitempty"`
	Description     string    `xml:"description,omitempty"`
	Amount          float64   `xml:"amount,omitempty"`
	Quantity        string    `xml:"quantity,omitempty"`
	Price           RecordRef `xml:"price,omitempty"`
	Rate            float64   `xml:"rate,omitempty"`
	Department      RecordRef `xml:"department,omitempty"`
	GrossAmt        float64   `xml:"grossAmt,omitempty"`
	Tax1Amt         float64   `xml:"tax1Amt,omitempty"`
	Tax1Acct        RecordRef `xml:"tax1Acct,omitempty"`
	TaxCode         RecordRef `xml:"taxCode,omitempty"`
	TaxRate1        float64   `xml:"taxRate1,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
}

func (i CreditMemoItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

func (i CreditMemoItem) IsEmpty() bool {
	return zero.IsZero(i)
}

// https://www.netsuite.com/help/helpcenter/en_US/srbrowser/Browser2017_1/schema/record/statisticaljournalentry.html
type StatisticalJournalEntry struct {
	// Approved        bool `xml:"approved"`
	CreatedDate     Date `xml:"createdDate,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	// Select the custom journal entry record, if required.
	CustomForm         RecordRef                    `xml:"customForm,omitempty"`
	LastModifiedDate   Date                         `xml:"lastModifiedDate,omitempty"`
	LineList           StatisticalJournalEntryLines `xml:"lineList>statisticalJournalEntryLine"`
	Memo               string                       `xml:"memo,omitempty"`
	ParentExpenseAlloc RecordRef                    `xml:"parentExpenseAlloc,omitempty"`
	PostingPeriod      RecordRef                    `xml:"postingPeriod,omitempty"`
	ReversalDate       Date                         `xml:"reversalDate,omitempty"`
	ReversalDefer      bool                         `xml:"reversalDefer,omitempty"`
	Subsidiary         RecordRef                    `xml:"subsidiary,omitempty"`
	TranDate           Date                         `xml:"tranDate,omitempty"`
	TranID             string                       `xml:"tranId,omitempty"`
	// This field displays the unit type to associate with this statistical journal entry.
	UnitsType RecordRef `xml:"unitsType,omitempty"`
}

func (j StatisticalJournalEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j StatisticalJournalEntry) IsEmpty() bool {
	return zero.IsZero(j)
}

type StatisticalJournalEntryLines []StatisticalJournalEntryLine

type StatisticalJournalEntryLine struct {
	XMLName xml.Name `xml:"line"`

	Account         RecordRef `xml:"account"`
	Class           RecordRef `xml:"class,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	Debit        float64   `xml:"debit,omitempty"`
	Department   RecordRef `xml:"department,omitempty"`
	Entity       RecordRef `xml:"entity,omitempty"`
	Line         int       `xml:"line,omitempty"`
	LineUnit     RecordRef `xml:"lineUnit,omitempty"`
	Location     RecordRef `xml:"location,omitempty"`
	Memo         string    `xml:"memo,omitempty"`
	PreviewDebit string    `xml:"previewDebit,omitempty"`
	ScheduleNum  RecordRef `xml:"scheduleNum,omitempty"`
}

func (j StatisticalJournalEntryLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j StatisticalJournalEntryLine) IsEmpty() bool {
	return zero.IsZero(j)
}
