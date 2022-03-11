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

func (c *Client) NewCustomerSearchRequest() CustomerSearchRequest {
	r := CustomerSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerSearchRequest struct {
	client      *Client
	queryParams *CustomerSearchRequestQueryParams
	pathParams  *CustomerSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomerSearchRequestBody
}

func (r CustomerSearchRequest) SOAPAction() string {
	return "search"
}

func (r CustomerSearchRequest) NewQueryParams() *CustomerSearchRequestQueryParams {
	return &CustomerSearchRequestQueryParams{}
}

type CustomerSearchRequestQueryParams struct {
}

func (p CustomerSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r CustomerSearchRequest) NewPathParams() *CustomerSearchRequestPathParams {
	return &CustomerSearchRequestPathParams{}
}

type CustomerSearchRequestPathParams struct {
}

func (p *CustomerSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *CustomerSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerSearchRequest) Method() string {
	return r.method
}

func (r CustomerSearchRequest) NewRequestBody() CustomerSearchRequestBody {
	body := CustomerSearchRequestBody{}
	body.SearchRecord.Type = "listRel:CustomerSearch"
	return body
}

type CustomerSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord struct {
		XMLName xml.Name `xml:"platformMsgs:searchRecord"`

		Type  string              `xml:"xsi:type,attr"`
		Basic CustomerSearchBasic `xml:"listRel:basic"`
	} `xml:"platformMsgs:searchRecord"`
}

func (r *CustomerSearchRequest) RequestBody() *CustomerSearchRequestBody {
	return &r.requestBody
}

func (r *CustomerSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerSearchRequest) SetRequestBody(body CustomerSearchRequestBody) {
	r.requestBody = body
}

func (r *CustomerSearchRequest) NewResponseBody() *CustomerSearchRequestResponseBody {
	return &CustomerSearchRequestResponseBody{}
}

type CustomerSearchRequestResponseBody struct {
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
			Record Customers `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *CustomerSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *CustomerSearchRequest) Do() (CustomerSearchRequestResponseBody, error) {
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

type CustomerSearchBasic struct {
	AccountNumber             SearchStringField          `xml:"platformCommon:accountNumber,omitempty"`
	Address                   SearchStringField          `xml:"platformCommon:address,omitempty"`
	Addressee                 SearchStringField          `xml:"platformCommon:addressee,omitempty"`
	AddressLabel              SearchStringField          `xml:"platformCommon:addressLabel,omitempty"`
	AddressPhone              SearchStringField          `xml:"platformCommon:addressPhone,omitempty"`
	AssignedSite              SearchMultiSelectField     `xml:"platformCommon:assignedSite,omitempty"`
	AssignedSiteID            SearchMultiSelectField     `xml:"platformCommon:assignedSiteId,omitempty"`
	Attention                 SearchStringField          `xml:"platformCommon:attention,omitempty"`
	AvailableOffline          SearchBooleanField         `xml:"platformCommon:availableOffline,omitempty"`
	Balance                   SearchDoubleField          `xml:"platformCommon:balance,omitempty"`
	BillAddress               SearchStringField          `xml:"platformCommon:billAddress,omitempty"`
	BoughtAmount              SearchDoubleField          `xml:"platformCommon:boughtAmount,omitempty"`
	BoughtDate                SearchDateField            `xml:"platformCommon:boughtDate,omitempty"`
	BuyingReason              SearchMultiSelectField     `xml:"platformCommon:buyingReason,omitempty"`
	BuyingTimeFrame           SearchMultiSelectField     `xml:"platformCommon:buyingTimeFrame,omitempty"`
	Category                  SearchMultiSelectField     `xml:"platformCommon:category,omitempty"`
	CcCustomerCode            SearchStringField          `xml:"platformCommon:ccCustomerCode,omitempty"`
	CcDefault                 SearchBooleanField         `xml:"platformCommon:ccDefault,omitempty"`
	CcExpDate                 SearchDateField            `xml:"platformCommon:ccExpDate,omitempty"`
	CcHolderName              SearchStringField          `xml:"platformCommon:ccHolderName,omitempty"`
	CcNumber                  SearchStringField          `xml:"platformCommon:ccNumber,omitempty"`
	CcState                   SearchMultiSelectField     `xml:"platformCommon:ccState,omitempty"`
	CcStateFrom               SearchDateField            `xml:"platformCommon:ccStateFrom,omitempty"`
	CcType                    SearchMultiSelectField     `xml:"platformCommon:ccType,omitempty"`
	City                      SearchStringField          `xml:"platformCommon:city,omitempty"`
	ClassBought               SearchMultiSelectField     `xml:"platformCommon:classBought,omitempty"`
	Comments                  SearchStringField          `xml:"platformCommon:comments,omitempty"`
	CompanyName               SearchStringField          `xml:"platformCommon:companyName,omitempty"`
	ConsolBalance             SearchDoubleField          `xml:"platformCommon:consolBalance,omitempty"`
	ConsolDaysOverdue         SearchLongField            `xml:"platformCommon:consolDaysOverdue,omitempty"`
	ConsolDepositBalance      SearchDoubleField          `xml:"platformCommon:consolDepositBalance,omitempty"`
	ConsolOverdueBalance      SearchDoubleField          `xml:"platformCommon:consolOverdueBalance,omitempty"`
	ConsolUnbilledOrders      SearchDoubleField          `xml:"platformCommon:consolUnbilledOrders,omitempty"`
	Contact                   SearchStringField          `xml:"platformCommon:contact,omitempty"`
	Contribution              SearchLongField            `xml:"platformCommon:contribution,omitempty"`
	ConversionDate            SearchDateField            `xml:"platformCommon:conversionDate,omitempty"`
	Country                   SearchEnumMultiSelectField `xml:"platformCommon:country,omitempty"`
	County                    SearchStringField          `xml:"platformCommon:county,omitempty"`
	CreditHold                SearchEnumMultiSelectField `xml:"platformCommon:creditHold,omitempty"`
	CreditHoldOverride        SearchBooleanField         `xml:"platformCommon:creditHoldOverride,omitempty"`
	CreditLimit               SearchDoubleField          `xml:"platformCommon:creditLimit,omitempty"`
	Currency                  SearchMultiSelectField     `xml:"platformCommon:currency,omitempty"`
	CustStage                 SearchMultiSelectField     `xml:"platformCommon:custStage,omitempty"`
	CustStatus                SearchMultiSelectField     `xml:"platformCommon:custStatus,omitempty"`
	DateClosed                SearchDateField            `xml:"platformCommon:dateClosed,omitempty"`
	DateCreated               SearchDateField            `xml:"platformCommon:dateCreated,omitempty"`
	DaysOverdue               SearchLongField            `xml:"platformCommon:daysOverdue,omitempty"`
	DefaultAllocationStrategy SearchMultiSelectField     `xml:"platformCommon:defaultAllocationStrategy,omitempty"`
	DefaultOrderPriority      SearchDoubleField          `xml:"platformCommon:defaultOrderPriority,omitempty"`
	DefaultTaxReg             SearchMultiSelectField     `xml:"platformCommon:defaultTaxReg,omitempty"`
	DefaultTaxRegText         SearchStringField          `xml:"platformCommon:defaultTaxRegText,omitempty"`
	DepositBalance            SearchDoubleField          `xml:"platformCommon:depositBalance,omitempty"`
	DeptBought                SearchMultiSelectField     `xml:"platformCommon:deptBought,omitempty"`
	DrAccount                 SearchMultiSelectField     `xml:"platformCommon:drAccount,omitempty"`
	Email                     SearchStringField          `xml:"platformCommon:email,omitempty"`
	EmailPreference           SearchEnumMultiSelectField `xml:"platformCommon:emailPreference,omitempty"`
	EmailTransactions         SearchBooleanField         `xml:"platformCommon:emailTransactions,omitempty"`
	EndDate                   SearchDateField            `xml:"platformCommon:endDate,omitempty"`
	EntityID                  SearchStringField          `xml:"platformCommon:entityId,omitempty"`
	EntityStatus              SearchMultiSelectField     `xml:"platformCommon:entityStatus,omitempty"`
	EstimatedBudget           SearchDoubleField          `xml:"platformCommon:estimatedBudget,omitempty"`
	ExplicitConversion        SearchBooleanField         `xml:"platformCommon:explicitConversion,omitempty"`
	ExternalID                SearchMultiSelectField     `xml:"platformCommon:externalId,omitempty"`
	ExternalIDString          SearchStringField          `xml:"platformCommon:externalIdString,omitempty"`
	Fax                       SearchStringField          `xml:"platformCommon:fax,omitempty"`
	FaxTransactions           SearchBooleanField         `xml:"platformCommon:faxTransactions,omitempty"`
	FirstName                 SearchStringField          `xml:"platformCommon:firstName,omitempty"`
	FirstOrderDate            SearchDateField            `xml:"platformCommon:firstOrderDate,omitempty"`
	FirstSaleDate             SearchDateField            `xml:"platformCommon:firstSaleDate,omitempty"`
	FxAccount                 SearchMultiSelectField     `xml:"platformCommon:fxAccount,omitempty"`
	FxBalance                 SearchDoubleField          `xml:"platformCommon:fxBalance,omitempty"`
	FxConsolBalance           SearchDoubleField          `xml:"platformCommon:fxConsolBalance,omitempty"`
	FxConsolUnbilledOrders    SearchDoubleField          `xml:"platformCommon:fxConsolUnbilledOrders,omitempty"`
	FxUnbilledOrders          SearchDoubleField          `xml:"platformCommon:fxUnbilledOrders,omitempty"`
	GiveAccess                SearchBooleanField         `xml:"platformCommon:giveAccess,omitempty"`
	GlobalSubscriptionStatus  SearchEnumMultiSelectField `xml:"platformCommon:globalSubscriptionStatus,omitempty"`
	Group                     SearchMultiSelectField     `xml:"platformCommon:group,omitempty"`
	GroupPricingLevel         SearchMultiSelectField     `xml:"platformCommon:groupPricingLevel,omitempty"`
	HasDuplicates             SearchBooleanField         `xml:"platformCommon:hasDuplicates,omitempty"`
	Image                     SearchStringField          `xml:"platformCommon:image,omitempty"`
	InternalID                SearchMultiSelectField     `xml:"platformCommon:internalId,omitempty"`
	InternalIDNumber          SearchLongField            `xml:"platformCommon:internalIdNumber,omitempty"`
	IsBudgetApproved          SearchBooleanField         `xml:"platformCommon:isBudgetApproved,omitempty"`
	IsDefaultBilling          SearchBooleanField         `xml:"platformCommon:isDefaultBilling,omitempty"`
	IsDefaultShipping         SearchBooleanField         `xml:"platformCommon:isDefaultShipping,omitempty"`
	IsInactive                SearchBooleanField         `xml:"platformCommon:isInactive,omitempty"`
	IsPerson                  SearchBooleanField         `xml:"platformCommon:isPerson,omitempty"`
	IsReportedLead            SearchBooleanField         `xml:"platformCommon:isReportedLead,omitempty"`
	IsShipAddress             SearchBooleanField         `xml:"platformCommon:isShipAddress,omitempty"`
	ItemPricingLevel          SearchMultiSelectField     `xml:"platformCommon:itemPricingLevel,omitempty"`
	ItemPricingUnitPrice      SearchDoubleField          `xml:"platformCommon:itemPricingUnitPrice,omitempty"`
	ItemsBought               SearchMultiSelectField     `xml:"platformCommon:itemsBought,omitempty"`
	ItemsOrdered              SearchMultiSelectField     `xml:"platformCommon:itemsOrdered,omitempty"`
	Language                  SearchEnumMultiSelectField `xml:"platformCommon:language,omitempty"`
	LastModifiedDate          SearchDateField            `xml:"platformCommon:lastModifiedDate,omitempty"`
	LastName                  SearchStringField          `xml:"platformCommon:lastName,omitempty"`
	LastOrderDate             SearchDateField            `xml:"platformCommon:lastOrderDate,omitempty"`
	LastSaleDate              SearchDateField            `xml:"platformCommon:lastSaleDate,omitempty"`
	LeadDate                  SearchDateField            `xml:"platformCommon:leadDate,omitempty"`
	LeadSource                SearchMultiSelectField     `xml:"platformCommon:leadSource,omitempty"`
	Level                     SearchEnumMultiSelectField `xml:"platformCommon:level,omitempty"`
	LocationBought            SearchMultiSelectField     `xml:"platformCommon:locationBought,omitempty"`
	ManualCreditHold          SearchBooleanField         `xml:"platformCommon:manualCreditHold,omitempty"`
	MerchantAccount           SearchMultiSelectField     `xml:"platformCommon:merchantAccount,omitempty"`
	MiddleName                SearchStringField          `xml:"platformCommon:middleName,omitempty"`
	MonthlyClosing            SearchEnumMultiSelectField `xml:"platformCommon:monthlyClosing,omitempty"`
	OnCreditHold              SearchBooleanField         `xml:"platformCommon:onCreditHold,omitempty"`
	OrderedAmount             SearchDoubleField          `xml:"platformCommon:orderedAmount,omitempty"`
	OrderedDate               SearchDateField            `xml:"platformCommon:orderedDate,omitempty"`
	OtherRelationships        SearchEnumMultiSelectField `xml:"platformCommon:otherRelationships,omitempty"`
	OverdueBalance            SearchDoubleField          `xml:"platformCommon:overdueBalance,omitempty"`
	Parent                    SearchMultiSelectField     `xml:"platformCommon:parent,omitempty"`
	ParentItemsBought         SearchMultiSelectField     `xml:"platformCommon:parentItemsBought,omitempty"`
	ParentItemsOrdered        SearchMultiSelectField     `xml:"platformCommon:parentItemsOrdered,omitempty"`
	Partner                   SearchMultiSelectField     `xml:"platformCommon:partner,omitempty"`
	PartnerContribution       SearchLongField            `xml:"platformCommon:partnerContribution,omitempty"`
	PartnerRole               SearchMultiSelectField     `xml:"platformCommon:partnerRole,omitempty"`
	PartnerTeamMember         SearchMultiSelectField     `xml:"platformCommon:partnerTeamMember,omitempty"`
	Pec                       SearchStringField          `xml:"platformCommon:pec,omitempty"`
	Permission                SearchEnumMultiSelectField `xml:"platformCommon:permission,omitempty"`
	Phone                     SearchStringField          `xml:"platformCommon:phone,omitempty"`
	PhoneticName              SearchStringField          `xml:"platformCommon:phoneticName,omitempty"`
	PriceLevel                SearchMultiSelectField     `xml:"platformCommon:priceLevel,omitempty"`
	PricingGroup              SearchMultiSelectField     `xml:"platformCommon:pricingGroup,omitempty"`
	PricingItem               SearchMultiSelectField     `xml:"platformCommon:pricingItem,omitempty"`
	PrintTransactions         SearchBooleanField         `xml:"platformCommon:printTransactions,omitempty"`
	ProspectDate              SearchDateField            `xml:"platformCommon:prospectDate,omitempty"`
	PstExempt                 SearchBooleanField         `xml:"platformCommon:pstExempt,omitempty"`
	ReceivablesAccount        SearchMultiSelectField     `xml:"platformCommon:receivablesAccount,omitempty"`
	ReminderDate              SearchDateField            `xml:"platformCommon:reminderDate,omitempty"`
	ResaleNumber              SearchStringField          `xml:"platformCommon:resaleNumber,omitempty"`
	Role                      SearchMultiSelectField     `xml:"platformCommon:role,omitempty"`
	SalesReadiness            SearchMultiSelectField     `xml:"platformCommon:salesReadiness,omitempty"`
	SalesRep                  SearchMultiSelectField     `xml:"platformCommon:salesRep,omitempty"`
	SalesTeamMember           SearchMultiSelectField     `xml:"platformCommon:salesTeamMember,omitempty"`
	SalesTeamRole             SearchMultiSelectField     `xml:"platformCommon:salesTeamRole,omitempty"`
	Salutation                SearchStringField          `xml:"platformCommon:salutation,omitempty"`
	ShipAddress               SearchStringField          `xml:"platformCommon:shipAddress,omitempty"`
	ShipComplete              SearchBooleanField         `xml:"platformCommon:shipComplete,omitempty"`
	ShippingItem              SearchMultiSelectField     `xml:"platformCommon:shippingItem,omitempty"`
	SourceSite                SearchMultiSelectField     `xml:"platformCommon:sourceSite,omitempty"`
	SourceSiteID              SearchMultiSelectField     `xml:"platformCommon:sourceSiteId,omitempty"`
	Stage                     SearchEnumMultiSelectField `xml:"platformCommon:stage,omitempty"`
	StartDate                 SearchDateField            `xml:"platformCommon:startDate,omitempty"`
	State                     SearchStringField          `xml:"platformCommon:state,omitempty"`
	SubsidBought              SearchMultiSelectField     `xml:"platformCommon:subsidBought,omitempty"`
	Subsidiary                SearchMultiSelectField     `xml:"platformCommon:subsidiary,omitempty"`
	Taxable                   SearchBooleanField         `xml:"platformCommon:taxable,omitempty"`
	Terms                     SearchMultiSelectField     `xml:"platformCommon:terms,omitempty"`
	Territory                 SearchMultiSelectField     `xml:"platformCommon:territory,omitempty"`
	Title                     SearchStringField          `xml:"platformCommon:title,omitempty"`
	UnbilledOrders            SearchDoubleField          `xml:"platformCommon:unbilledOrders,omitempty"`
	Url                       SearchStringField          `xml:"platformCommon:url,omitempty"`
	VatRegNumber              SearchStringField          `xml:"platformCommon:vatRegNumber,omitempty"`
	WebLead                   SearchBooleanField         `xml:"platformCommon:webLead,omitempty"`
	ZipCode                   SearchStringField          `xml:"platformCommon:zipCode,omitempty"`
	CustomFieldList           SearchCustomFieldList      `xml:"platformCommon:customFieldList,omitempty"`
}

func (c CustomerSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomerSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
