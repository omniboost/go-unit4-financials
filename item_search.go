package financials

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-unit4-financials/omitempty"
	"github.com/omniboost/go-unit4-financials/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewItemSearchRequest() ItemSearchRequest {
	r := ItemSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ItemSearchRequest struct {
	client      *Client
	queryParams *ItemSearchRequestQueryParams
	pathParams  *ItemSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody ItemSearchRequestBody
}

func (r ItemSearchRequest) SOAPAction() string {
	return "search"
}

func (r ItemSearchRequest) NewQueryParams() *ItemSearchRequestQueryParams {
	return &ItemSearchRequestQueryParams{}
}

type ItemSearchRequestQueryParams struct {
}

func (p ItemSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ItemSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ItemSearchRequest) NewPathParams() *ItemSearchRequestPathParams {
	return &ItemSearchRequestPathParams{}
}

type ItemSearchRequestPathParams struct {
}

func (p *ItemSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ItemSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *ItemSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *ItemSearchRequest) Method() string {
	return r.method
}

func (r ItemSearchRequest) NewRequestBody() ItemSearchRequestBody {
	return ItemSearchRequestBody{
		SearchRecord: ItemSearchRecordBasic{
			Type: "listAcct:ItemSearch",
		},
	}
}

type ItemSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord ItemSearchRecordBasic `xml:"platformMsgs:searchRecord"`
}

func (r *ItemSearchRequest) RequestBody() *ItemSearchRequestBody {
	return &r.requestBody
}

func (r *ItemSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ItemSearchRequest) SetRequestBody(body ItemSearchRequestBody) {
	r.requestBody = body
}

func (r *ItemSearchRequest) NewResponseBody() *ItemSearchRequestResponseBody {
	return &ItemSearchRequestResponseBody{}
}

type ItemSearchRequestResponseBody struct {
	XMLName xml.Name `xml:"searchResponse"`

	SearchResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		TotalRecords string `xml:"totalRecords"`
		PageSize     string `xml:"pageSize"`
		TotalPages   string `xml:"totalPages"`
		PageIndex    string `xml:"pageIndex"`
		SearchID     string `xml:"searchId"`
		RecordList   struct {
			Record Items `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *ItemSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *ItemSearchRequest) Do() (ItemSearchRequestResponseBody, error) {
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

type ItemSearchBasic struct {
	AccBookRevRecForecastRule SearchMultiSelectField `xml:"accBookRevRecForecastRule,omitempty"`
	Account SearchMultiSelectField `xml:"account,omitempty"`
	AccountingBook SearchMultiSelectField `xml:"accountingBook,omitempty"`
	AccountingBookAmortization SearchMultiSelectField `xml:"accountingBookAmortization,omitempty"`
	AccountingBookCreatePlansOn SearchMultiSelectField `xml:"accountingBookCreatePlansOn,omitempty"`
	AccountingBookRevRecRule SearchMultiSelectField `xml:"accountingBookRevRecRule,omitempty"`
	AccountingBookRevRecSchedule SearchMultiSelectField `xml:"accountingBookRevRecSchedule,omitempty"`
	AllowedShippingMethod SearchMultiSelectField `xml:"allowedShippingMethod,omitempty"`
	AlternateDemandSourceItem SearchMultiSelectField `xml:"alternateDemandSourceItem,omitempty"`
	AtpLeadTime SearchDoubleField `xml:"atpLeadTime,omitempty"`
	AtpMethod SearchEnumMultiSelectField `xml:"atpMethod,omitempty"`
	AutoLeadTime SearchBooleanField `xml:"autoLeadTime,omitempty"`
	AutoPreferredStockLevel SearchBooleanField `xml:"autoPreferredStockLevel,omitempty"`
	AutoReorderPoint SearchBooleanField `xml:"autoReorderPoint,omitempty"`
	AvailableToPartners SearchBooleanField `xml:"availableToPartners,omitempty"`
	AverageCost SearchDoubleField `xml:"averageCost,omitempty"`
	BackwardConsumptionDays SearchLongField `xml:"backwardConsumptionDays,omitempty"`
	BinNumber SearchStringField `xml:"binNumber,omitempty"`
	BinOnHandAvail SearchDoubleField `xml:"binOnHandAvail,omitempty"`
	BinOnHandCount SearchDoubleField `xml:"binOnHandCount,omitempty"`
	BomQuantity SearchDoubleField `xml:"bomQuantity,omitempty"`
	BuildEntireAssembly SearchBooleanField `xml:"buildEntireAssembly,omitempty"`
	BuildTime SearchDoubleField `xml:"buildTime,omitempty"`
	BuildTimeLotSize SearchDoubleField `xml:"buildTimeLotSize,omitempty"`
	BuyItNowPrice SearchDoubleField `xml:"buyItNowPrice,omitempty"`
	Caption SearchStringField `xml:"caption,omitempty"`
	Category SearchMultiSelectField `xml:"category,omitempty"`
	Class SearchMultiSelectField `xml:"class,omitempty"`
	Component SearchMultiSelectField `xml:"component,omitempty"`
	ComponentOf SearchMultiSelectField `xml:"componentOf,omitempty"`
	ComponentYield SearchDoubleField `xml:"componentYield,omitempty"`
	ConsumptionUnit SearchMultiSelectField `xml:"consumptionUnit,omitempty"`
	ContingentRevenueHandling SearchBooleanField `xml:"contingentRevenueHandling,omitempty"`
	ConversionRate SearchDoubleField `xml:"conversionRate,omitempty"`
	CopyDescription SearchBooleanField `xml:"copyDescription,omitempty"`
	CorrelatedItem SearchMultiSelectField `xml:"correlatedItem,omitempty"`
	CorrelatedItemCorrelation SearchDoubleField `xml:"correlatedItemCorrelation,omitempty"`
	CorrelatedItemCount SearchLongField `xml:"correlatedItemCount,omitempty"`
	CorrelatedItemLift SearchDoubleField `xml:"correlatedItemLift,omitempty"`
	CorrelatedItemPurchaseRate SearchDoubleField `xml:"correlatedItemPurchaseRate,omitempty"`
	Cost SearchDoubleField `xml:"cost,omitempty"`
	CostAccountingStatus SearchEnumMultiSelectField `xml:"costAccountingStatus,omitempty"`
	CostCategory SearchMultiSelectField `xml:"costCategory,omitempty"`
	CostEstimate SearchDoubleField `xml:"costEstimate,omitempty"`
	CostEstimateType SearchEnumMultiSelectField `xml:"costEstimateType,omitempty"`
	CostingMethod SearchEnumMultiSelectField `xml:"costingMethod,omitempty"`
	CountryOfManufacture SearchEnumMultiSelectField `xml:"countryOfManufacture,omitempty"`
	Created SearchDateField `xml:"created,omitempty"`
	CreateJob SearchBooleanField `xml:"createJob,omitempty"`
	CreateRevenuePlansOn SearchMultiSelectField `xml:"createRevenuePlansOn,omitempty"`
	DateViewed SearchDateField `xml:"dateViewed,omitempty"`
	DaysBeforeExpiration SearchDoubleField `xml:"daysBeforeExpiration,omitempty"`
	DefaultReturnCost SearchDoubleField `xml:"defaultReturnCost,omitempty"`
	DefaultShippingMethod SearchMultiSelectField `xml:"defaultShippingMethod,omitempty"`
	DeferRevRec SearchBooleanField `xml:"deferRevRec,omitempty"`
	DemandModifier SearchDoubleField `xml:"demandModifier,omitempty"`
	DemandSource SearchEnumMultiSelectField `xml:"demandSource,omitempty"`
	DemandTimeFence SearchLongField `xml:"demandTimeFence,omitempty"`
	Department SearchMultiSelectField `xml:"department,omitempty"`
	DirectRevenuePosting SearchBooleanField `xml:"directRevenuePosting,omitempty"`
	DisplayIneBayStore SearchBooleanField `xml:"displayIneBayStore,omitempty"`
	DisplayName SearchStringField `xml:"displayName,omitempty"`
	DistributionCategory SearchMultiSelectField `xml:"distributionCategory,omitempty"`
	DistributionNetwork SearchMultiSelectField `xml:"distributionNetwork,omitempty"`
	DontShowPrice SearchBooleanField `xml:"dontShowPrice,omitempty"`
	EBayItemDescription SearchStringField `xml:"eBayItemDescription,omitempty"`
	EBayItemSubtitle SearchStringField `xml:"eBayItemSubtitle,omitempty"`
	EBayItemTitle SearchStringField `xml:"eBayItemTitle,omitempty"`
	EbayRelistingOption SearchEnumMultiSelectField `xml:"ebayRelistingOption,omitempty"`
	EffectiveBomControl SearchEnumMultiSelectField `xml:"effectiveBomControl,omitempty"`
	EffectiveDate SearchDateField `xml:"effectiveDate,omitempty"`
	EffectiveRevision SearchMultiSelectField `xml:"effectiveRevision,omitempty"`
	EnableCatchWeight SearchBooleanField `xml:"enableCatchWeight,omitempty"`
	EndAuctionsWhenOutOfStock SearchBooleanField `xml:"endAuctionsWhenOutOfStock,omitempty"`
	ExcludeFromSitemap SearchBooleanField `xml:"excludeFromSitemap,omitempty"`
	ExternalID SearchMultiSelectField `xml:"externalId,omitempty"`
	ExternalIDString SearchStringField `xml:"externalIdString,omitempty"`
	FeaturedDescription SearchStringField `xml:"featuredDescription,omitempty"`
	FeedDescription SearchStringField `xml:"feedDescription,omitempty"`
	FeedName SearchStringField `xml:"feedName,omitempty"`
	FixedBuildTime SearchDoubleField `xml:"fixedBuildTime,omitempty"`
	FixedLotSize SearchDoubleField `xml:"fixedLotSize,omitempty"`
	ForwardConsumptionDays SearchLongField `xml:"forwardConsumptionDays,omitempty"`
	FraudRisk SearchEnumMultiSelectField `xml:"fraudRisk,omitempty"`
	FroogleProductFeed SearchBooleanField `xml:"froogleProductFeed,omitempty"`
	FutureHorizon SearchLongField `xml:"futureHorizon,omitempty"`
	FxCost SearchDoubleField `xml:"fxCost,omitempty"`
	GenerateAccruals SearchBooleanField `xml:"generateAccruals,omitempty"`
	GiftCertAuthCode SearchStringField `xml:"giftCertAuthCode,omitempty"`
	GiftCertEmail SearchStringField `xml:"giftCertEmail,omitempty"`
	GiftCertExpDate SearchDateField `xml:"giftCertExpDate,omitempty"`
	GiftCertFrom SearchStringField `xml:"giftCertFrom,omitempty"`
	GiftCertMsg SearchStringField `xml:"giftCertMsg,omitempty"`
	GiftCertOrigAmt SearchStringField `xml:"giftCertOrigAmt,omitempty"`
	GiftCertRecipient SearchStringField `xml:"giftCertRecipient,omitempty"`
	HierarchyNode SearchMultiSelectField `xml:"hierarchyNode,omitempty"`
	HierarchyVersion SearchMultiSelectField `xml:"hierarchyVersion,omitempty"`
	ImageUrl SearchStringField `xml:"imageUrl,omitempty"`
	IncludeChildren SearchBooleanField `xml:"includeChildren,omitempty"`
	InternalID SearchMultiSelectField `xml:"internalId,omitempty"`
	InternalIDNumber SearchLongField `xml:"internalIdNumber,omitempty"`
	InventoryLocation SearchMultiSelectField `xml:"inventoryLocation,omitempty"`
	InvtClassification SearchEnumMultiSelectField `xml:"invtClassification,omitempty"`
	InvtCountInterval SearchLongField `xml:"invtCountInterval,omitempty"`
	IsAvailable SearchBooleanField `xml:"isAvailable,omitempty"`
	IsDropShipItem SearchBooleanField `xml:"isDropShipItem,omitempty"`
	IsFulfillable SearchBooleanField `xml:"isFulfillable,omitempty"`
	IsGcoCompliant SearchBooleanField `xml:"isGcoCompliant,omitempty"`
	IsInactive SearchBooleanField `xml:"isInactive,omitempty"`
	IsLotItem SearchBooleanField `xml:"isLotItem,omitempty"`
	IsOnline SearchBooleanField `xml:"isOnline,omitempty"`
	IsPreferredVendor SearchBooleanField `xml:"isPreferredVendor,omitempty"`
	IsSerialItem SearchBooleanField `xml:"isSerialItem,omitempty"`
	IsSpecialOrderItem SearchBooleanField `xml:"isSpecialOrderItem,omitempty"`
	IsSpecialWorkOrderItem SearchBooleanField `xml:"isSpecialWorkOrderItem,omitempty"`
	IsStorePickupAllowed SearchBooleanField `xml:"isStorePickupAllowed,omitempty"`
	IssueProduct SearchMultiSelectField `xml:"issueProduct,omitempty"`
	IsTaxable SearchBooleanField `xml:"isTaxable,omitempty"`
	IsVsoeBundle SearchBooleanField `xml:"isVsoeBundle,omitempty"`
	IsWip SearchBooleanField `xml:"isWip,omitempty"`
	ItemID SearchStringField `xml:"itemId,omitempty"`
	ItemRevenueCategory SearchMultiSelectField `xml:"itemRevenueCategory,omitempty"`
	ItemUrl SearchStringField `xml:"itemUrl,omitempty"`
	LastInvtCountDate SearchDateField `xml:"lastInvtCountDate,omitempty"`
	LastModifiedDate SearchDateField `xml:"lastModifiedDate,omitempty"`
	LastPurchasePrice SearchDoubleField `xml:"lastPurchasePrice,omitempty"`
	LastQuantityAvailableChange SearchDateField `xml:"lastQuantityAvailableChange,omitempty"`
	LeadTime SearchLongField `xml:"leadTime,omitempty"`
	ListingDuration SearchEnumMultiSelectField `xml:"listingDuration,omitempty"`
	Location SearchMultiSelectField `xml:"location,omitempty"`
	LocationAllowStorePickup SearchBooleanField `xml:"locationAllowStorePickup,omitempty"`
	LocationAtpLeadTime SearchDoubleField `xml:"locationAtpLeadTime,omitempty"`
	LocationAverageCost SearchDoubleField `xml:"locationAverageCost,omitempty"`
	LocationBuildTime SearchDoubleField `xml:"locationBuildTime,omitempty"`
	LocationCost SearchDoubleField `xml:"locationCost,omitempty"`
	LocationCostAccountingStatus SearchEnumMultiSelectField `xml:"locationCostAccountingStatus,omitempty"`
	LocationDefaultReturnCost SearchDoubleField `xml:"locationDefaultReturnCost,omitempty"`
	LocationDemandSource SearchEnumMultiSelectField `xml:"locationDemandSource,omitempty"`
	LocationDemandTimeFence SearchLongField `xml:"locationDemandTimeFence,omitempty"`
	LocationFixedLotSize SearchDoubleField `xml:"locationFixedLotSize,omitempty"`
	LocationInventoryCostTemplate SearchMultiSelectField `xml:"locationInventoryCostTemplate,omitempty"`
	LocationInvtClassification SearchEnumMultiSelectField `xml:"locationInvtClassification,omitempty"`
	LocationInvtCountInterval SearchLongField `xml:"locationInvtCountInterval,omitempty"`
	LocationLastInvtCountDate SearchDateField `xml:"locationLastInvtCountDate,omitempty"`
	LocationLeadTime SearchLongField `xml:"locationLeadTime,omitempty"`
	LocationNextInvtCountDate SearchDateField `xml:"locationNextInvtCountDate,omitempty"`
	LocationPeriodicLotSizeDays SearchLongField `xml:"locationPeriodicLotSizeDays,omitempty"`
	LocationPeriodicLotSizeType SearchEnumMultiSelectField `xml:"locationPeriodicLotSizeType,omitempty"`
	LocationPreferredStockLevel SearchDoubleField `xml:"locationPreferredStockLevel,omitempty"`
	LocationQtyAvailForStorePickup SearchDoubleField `xml:"locationQtyAvailForStorePickup,omitempty"`
	LocationQuantityAvailable SearchDoubleField `xml:"locationQuantityAvailable,omitempty"`
	LocationQuantityBackOrdered SearchDoubleField `xml:"locationQuantityBackOrdered,omitempty"`
	LocationQuantityCommitted SearchDoubleField `xml:"locationQuantityCommitted,omitempty"`
	LocationQuantityInTransit SearchDoubleField `xml:"locationQuantityInTransit,omitempty"`
	LocationQuantityOnHand SearchDoubleField `xml:"locationQuantityOnHand,omitempty"`
	LocationQuantityOnOrder SearchDoubleField `xml:"locationQuantityOnOrder,omitempty"`
	LocationReorderPoint SearchDoubleField `xml:"locationReorderPoint,omitempty"`
	LocationRescheduleInDays SearchLongField `xml:"locationRescheduleInDays,omitempty"`
	LocationRescheduleOutDays SearchLongField `xml:"locationRescheduleOutDays,omitempty"`
	LocationSafetyStockLevel SearchDoubleField `xml:"locationSafetyStockLevel,omitempty"`
	LocationStorePickupBufferStock SearchDoubleField `xml:"locationStorePickupBufferStock,omitempty"`
	LocationSupplyLotSizingMethod SearchEnumMultiSelectField `xml:"locationSupplyLotSizingMethod,omitempty"`
	LocationSupplyTimeFence SearchLongField `xml:"locationSupplyTimeFence,omitempty"`
	LocationSupplyType SearchEnumMultiSelectField `xml:"locationSupplyType,omitempty"`
	LocationTotalValue SearchDoubleField `xml:"locationTotalValue,omitempty"`
	LocBackwardConsumptionDays SearchLongField `xml:"locBackwardConsumptionDays,omitempty"`
	LocForwardConsumptionDays SearchLongField `xml:"locForwardConsumptionDays,omitempty"`
	LowerWarningLimit SearchDoubleField `xml:"lowerWarningLimit,omitempty"`
	Manufacturer SearchStringField `xml:"manufacturer,omitempty"`
	Manufactureraddr1 SearchStringField `xml:"manufactureraddr1,omitempty"`
	ManufacturerCity SearchStringField `xml:"manufacturerCity,omitempty"`
	ManufacturerState SearchStringField `xml:"manufacturerState,omitempty"`
	ManufacturerTariff SearchStringField `xml:"manufacturerTariff,omitempty"`
	ManufacturerTaxID SearchStringField `xml:"manufacturerTaxId,omitempty"`
	ManufacturerZip SearchStringField `xml:"manufacturerZip,omitempty"`
	ManufacturingChargeItem SearchBooleanField `xml:"manufacturingChargeItem,omitempty"`
	MatchBillToReceipt SearchBooleanField `xml:"matchBillToReceipt,omitempty"`
	Matrix SearchBooleanField `xml:"matrix,omitempty"`
	MatrixChild SearchBooleanField `xml:"matrixChild,omitempty"`
	MaximumQuantity SearchLongField `xml:"maximumQuantity,omitempty"`
	MetaTagHtml SearchStringField `xml:"metaTagHtml,omitempty"`
	MinimumQuantity SearchLongField `xml:"minimumQuantity,omitempty"`
	MossApplies SearchBooleanField `xml:"mossApplies,omitempty"`
	Mpn SearchStringField `xml:"mpn,omitempty"`
	MultManufactureAddr SearchBooleanField `xml:"multManufactureAddr,omitempty"`
	NexTagCategory SearchStringField `xml:"nexTagCategory,omitempty"`
	NexTagProductFeed SearchBooleanField `xml:"nexTagProductFeed,omitempty"`
	NextInvtCountDate SearchDateField `xml:"nextInvtCountDate,omitempty"`
	NumActiveListings SearchLongField `xml:"numActiveListings,omitempty"`
	NumberAllowedDownloads SearchDoubleField `xml:"numberAllowedDownloads,omitempty"`
	NumCurrentlyListed SearchLongField `xml:"numCurrentlyListed,omitempty"`
	ObsoleteDate SearchDateField `xml:"obsoleteDate,omitempty"`
	ObsoleteRevision SearchMultiSelectField `xml:"obsoleteRevision,omitempty"`
	OfferSupport SearchBooleanField `xml:"offerSupport,omitempty"`
	OnlineCustomerPrice SearchDoubleField `xml:"onlineCustomerPrice,omitempty"`
	OnSpecial SearchBooleanField `xml:"onSpecial,omitempty"`
	OtherVendor SearchMultiSelectField `xml:"otherVendor,omitempty"`
	OutOfStockBehavior SearchMultiSelectField `xml:"outOfStockBehavior,omitempty"`
	OverallQuantityPricingType SearchEnumMultiSelectField `xml:"overallQuantityPricingType,omitempty"`
	OverheadType SearchEnumMultiSelectField `xml:"overheadType,omitempty"`
	PageTitle SearchStringField `xml:"pageTitle,omitempty"`
	Parent SearchMultiSelectField `xml:"parent,omitempty"`
	PeriodicLotSizeDays SearchLongField `xml:"periodicLotSizeDays,omitempty"`
	PeriodicLotSizeType SearchEnumMultiSelectField `xml:"periodicLotSizeType,omitempty"`
	PlanningItemCategory SearchMultiSelectField `xml:"planningItemCategory,omitempty"`
	PreferenceCriterion SearchStringField `xml:"preferenceCriterion,omitempty"`
	PreferredBin SearchBooleanField `xml:"preferredBin,omitempty"`
	PreferredLocation SearchMultiSelectField `xml:"preferredLocation,omitempty"`
	PreferredStockLevel SearchDoubleField `xml:"preferredStockLevel,omitempty"`
	PreferredStockLevelDays SearchLongField `xml:"preferredStockLevelDays,omitempty"`
	Price SearchDoubleField `xml:"price,omitempty"`
	PricesIncludeTax SearchBooleanField `xml:"pricesIncludeTax,omitempty"`
	PricingGroup SearchMultiSelectField `xml:"pricingGroup,omitempty"`
	PrimaryCategory SearchLongField `xml:"primaryCategory,omitempty"`
	PurchaseOrderAmount SearchDoubleField `xml:"purchaseOrderAmount,omitempty"`
	PurchaseOrderQuantity SearchDoubleField `xml:"purchaseOrderQuantity,omitempty"`
	PurchaseOrderQuantityDiff SearchDoubleField `xml:"purchaseOrderQuantityDiff,omitempty"`
	PurchaseUnit SearchMultiSelectField `xml:"purchaseUnit,omitempty"`
	QuantityAvailable SearchDoubleField `xml:"quantityAvailable,omitempty"`
	QuantityBackOrdered SearchDoubleField `xml:"quantityBackOrdered,omitempty"`
	QuantityCommitted SearchDoubleField `xml:"quantityCommitted,omitempty"`
	QuantityOnHand SearchDoubleField `xml:"quantityOnHand,omitempty"`
	QuantityOnOrder SearchDoubleField `xml:"quantityOnOrder,omitempty"`
	QuantityPricingSchedule SearchMultiSelectField `xml:"quantityPricingSchedule,omitempty"`
	ReceiptAmount SearchDoubleField `xml:"receiptAmount,omitempty"`
	ReceiptQuantity SearchDoubleField `xml:"receiptQuantity,omitempty"`
	ReceiptQuantityDiff SearchDoubleField `xml:"receiptQuantityDiff,omitempty"`
	ReorderMultiple SearchLongField `xml:"reorderMultiple,omitempty"`
	ReorderPoint SearchDoubleField `xml:"reorderPoint,omitempty"`
	RescheduleInDays SearchLongField `xml:"rescheduleInDays,omitempty"`
	RescheduleOutDays SearchLongField `xml:"rescheduleOutDays,omitempty"`
	ReservePrice SearchDoubleField `xml:"reservePrice,omitempty"`
	RevenueAllocationGroup SearchMultiSelectField `xml:"revenueAllocationGroup,omitempty"`
	RevenueRecognitionRule SearchMultiSelectField `xml:"revenueRecognitionRule,omitempty"`
	RevRecForecastRule SearchMultiSelectField `xml:"revRecForecastRule,omitempty"`
	RevRecSchedule SearchMultiSelectField `xml:"revRecSchedule,omitempty"`
	RoundUpAsComponent SearchBooleanField `xml:"roundUpAsComponent,omitempty"`
	SafetyStockLevel SearchDoubleField `xml:"safetyStockLevel,omitempty"`
	SafetyStockLevelDays SearchLongField `xml:"safetyStockLevelDays,omitempty"`
	SalesDescription SearchStringField `xml:"salesDescription,omitempty"`
	SaleUnit SearchMultiSelectField `xml:"saleUnit,omitempty"`
	SameAsPrimaryBookAmortization SearchBooleanField `xml:"sameAsPrimaryBookAmortization,omitempty"`
	SameAsPrimaryBookRevRec SearchBooleanField `xml:"sameAsPrimaryBookRevRec,omitempty"`
	ScheduleBCode SearchEnumMultiSelectField `xml:"scheduleBCode,omitempty"`
	ScheduleBNumber SearchStringField `xml:"scheduleBNumber,omitempty"`
	ScheduleBQuantity SearchStringField `xml:"scheduleBQuantity,omitempty"`
	SearchKeywords SearchStringField `xml:"searchKeywords,omitempty"`
	SeasonalDemand SearchBooleanField `xml:"seasonalDemand,omitempty"`
	SecondaryConsumptionUnit SearchMultiSelectField `xml:"secondaryConsumptionUnit,omitempty"`
	SecondaryPurchaseUnit SearchMultiSelectField `xml:"secondaryPurchaseUnit,omitempty"`
	SecondarySaleUnit SearchMultiSelectField `xml:"secondarySaleUnit,omitempty"`
	SecondaryStockUnit SearchMultiSelectField `xml:"secondaryStockUnit,omitempty"`
	SecondaryUnitsType SearchMultiSelectField `xml:"secondaryUnitsType,omitempty"`
	SellOnEBay SearchBooleanField `xml:"sellOnEBay,omitempty"`
	SerialNumber SearchStringField `xml:"serialNumber,omitempty"`
	SerialNumberLocation SearchMultiSelectField `xml:"serialNumberLocation,omitempty"`
	ShipIndividually SearchBooleanField `xml:"shipIndividually,omitempty"`
	ShipPackage SearchMultiSelectField `xml:"shipPackage,omitempty"`
	ShippingCarrier SearchEnumMultiSelectField `xml:"shippingCarrier,omitempty"`
	ShippingRate SearchDoubleField `xml:"shippingRate,omitempty"`
	ShoppingDotComCategory SearchStringField `xml:"shoppingDotComCategory,omitempty"`
	ShoppingProductFeed SearchBooleanField `xml:"shoppingProductFeed,omitempty"`
	ShopzillaCategoryID SearchLongField `xml:"shopzillaCategoryId,omitempty"`
	ShopzillaProductFeed SearchBooleanField `xml:"shopzillaProductFeed,omitempty"`
	SitemapPriority SearchEnumMultiSelectField `xml:"sitemapPriority,omitempty"`
	SoftDescriptor SearchMultiSelectField `xml:"softDescriptor,omitempty"`
	StartingPrice SearchDoubleField `xml:"startingPrice,omitempty"`
	StockDescription SearchStringField `xml:"stockDescription,omitempty"`
	StockUnit SearchMultiSelectField `xml:"stockUnit,omitempty"`
	StoreDescription SearchStringField `xml:"storeDescription,omitempty"`
	Subsidiary SearchMultiSelectField `xml:"subsidiary,omitempty"`
	SubType SearchEnumMultiSelectField `xml:"subType,omitempty"`
	SupplyLotSizingMethod SearchEnumMultiSelectField `xml:"supplyLotSizingMethod,omitempty"`
	SupplyReplenishmentMethod SearchEnumMultiSelectField `xml:"supplyReplenishmentMethod,omitempty"`
	SupplyTimeFence SearchLongField `xml:"supplyTimeFence,omitempty"`
	SupplyType SearchEnumMultiSelectField `xml:"supplyType,omitempty"`
	TaxCode SearchMultiSelectField `xml:"taxCode,omitempty"`
	TaxSchedule SearchMultiSelectField `xml:"taxSchedule,omitempty"`
	ThumbnailUrl SearchStringField `xml:"thumbnailUrl,omitempty"`
	TotalValue SearchDoubleField `xml:"totalValue,omitempty"`
	TrackLandedCost SearchBooleanField `xml:"trackLandedCost,omitempty"`
	TransferPrice SearchDoubleField `xml:"transferPrice,omitempty"`
	Type SearchEnumMultiSelectField `xml:"type,omitempty"`
	UnitsType SearchMultiSelectField `xml:"unitsType,omitempty"`
	UpcCode SearchStringField `xml:"upcCode,omitempty"`
	UpperWarningLimit SearchDoubleField `xml:"upperWarningLimit,omitempty"`
	UrlComponent SearchStringField `xml:"urlComponent,omitempty"`
	UseBins SearchBooleanField `xml:"useBins,omitempty"`
	UseComponentYield SearchBooleanField `xml:"useComponentYield,omitempty"`
	UseMarginalRates SearchBooleanField `xml:"useMarginalRates,omitempty"`
	Vendor SearchMultiSelectField `xml:"vendor,omitempty"`
	VendorCode SearchStringField `xml:"vendorCode,omitempty"`
	VendorCost SearchDoubleField `xml:"vendorCost,omitempty"`
	VendorCostEntered SearchDoubleField `xml:"vendorCostEntered,omitempty"`
	VendorName SearchStringField `xml:"vendorName,omitempty"`
	VendorPriceCurrency SearchMultiSelectField `xml:"vendorPriceCurrency,omitempty"`
	VsoeDeferral SearchEnumMultiSelectField `xml:"vsoeDeferral,omitempty"`
	VsoeDelivered SearchBooleanField `xml:"vsoeDelivered,omitempty"`
	VsoePermitDiscount SearchEnumMultiSelectField `xml:"vsoePermitDiscount,omitempty"`
	VsoePrice SearchDoubleField `xml:"vsoePrice,omitempty"`
	VsoeSopGroup SearchEnumMultiSelectField `xml:"vsoeSopGroup,omitempty"`
	WebSite SearchMultiSelectField `xml:"webSite,omitempty"`
	Weight SearchDoubleField `xml:"weight,omitempty"`
	YahooProductFeed SearchBooleanField `xml:"yahooProductFeed,omitempty"`
	CustomFieldList SearchCustomFieldList `xml:"customFieldList,omitempty"`
}

func (c ItemSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c ItemSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}

