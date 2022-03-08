package netsuite

type RecordRef struct {
	InternalId string `xml:"internalId,attr"`
	Name       string `xml:"name"`
}

type Customers []Customer

type Customer struct {
	InternalId           string    `xml:"internalId,attr"`
	ExternalId           string    `xml:"externalId,attr"`
	Type                 string    `xml:"type,attr"`
	ListRel              string    `xml:"listRel,attr"`
	EntityId             string    `xml:"entityId"`
	IsPerson             string    `xml:"isPerson"`
	CompanyName          string    `xml:"companyName"`
	EntityStatus         RecordRef `xml:"entityStatus"`
	Phone                string    `xml:"phone"`
	Email                string    `xml:"email"`
	DefaultAddress       string    `xml:"defaultAddress"`
	IsInactive           string    `xml:"isInactive"`
	Language             string    `xml:"language"`
	DateCreated          string    `xml:"dateCreated"`
	Subsidiary           RecordRef `xml:"subsidiary"`
	VatRegNumber         string    `xml:"vatRegNumber"`
	Terms                RecordRef `xml:"terms"`
	CreditLimit          string    `xml:"creditLimit"`
	UnbilledOrders       string    `xml:"unbilledOrders"`
	Currency             RecordRef `xml:"currency"`
	ShipComplete         string    `xml:"shipComplete"`
	AlcoholRecipientType string    `xml:"alcoholRecipientType"`
	ReceivablesAccount   RecordRef `xml:"receivablesAccount"`
	LastModifiedDate     string    `xml:"lastModifiedDate"`
	Stage                string    `xml:"stage"`
	IsBudgetApproved     string    `xml:"isBudgetApproved"`
	CustomFieldList      struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type CustomFields []CustomField

type CustomField struct {
	InternalId string `xml:"internalId,attr"`
	ScriptId   string `xml:"scriptId,attr"`
	Type       string `xml:"type,attr"`
	Value      struct {
		InternalId string `xml:"internalId,attr"`
		TypeId     string `xml:"typeId,attr"`
		Name       string `xml:"name"`
	} `xml:"value"`
}

type Accounts []Account

type Account struct {
	InternalId      string    `xml:"internalId,attr"`
	ExternalId      string    `xml:"externalId,attr"`
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
