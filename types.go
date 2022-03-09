package netsuite

type RecordRef struct {
	InternalID string `xml:"internalId,attr"`
	Name       string `xml:"name"`
}

type Customers []Customer

type Customer struct {
	InternalID           string    `xml:"internalId,attr"`
	ExternalID           string    `xml:"externalId,attr"`
	Type                 string    `xml:"type,attr"`
	ListRel              string    `xml:"listRel,attr"`
	EntityID             string    `xml:"entityId"`
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
	InternalID string `xml:"internalId,attr"`
	ScriptID   string `xml:"scriptId,attr"`
	Type       string `xml:"type,attr"`
	Value      string `xml:"value"`
	// Value      struct {
	// 	InternalID string `xml:"internalId,attr"`
	// 	TypeID     string `xml:"typeId,attr"`
	// 	Name       string `xml:"name"`
	// } `xml:"value"`
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
