package product

type Product struct {
	ID                string           `json:"id"`
	ChannelID         string           `json:"channelId"`
	ChannelName       string           `json:"channelName"`
	Marketplace       Marketplace      `json:"marketplace"`
	Language          string           `json:"language"`
	LastFetchTime     string           `json:"lastFetchTime"`
	PublishedContent  PublishedContent `json:"publishedContent"`
	ProductInfo       []ProductInfo    `json:"productInfo"`
	Search            Search           `json:"search"`
	CollectionTermIDS []string         `json:"collectionTermIds"`
	ResourceType      string           `json:"resourceType"`
	Links             ProductLinks     `json:"links"`
	Collectionsv2     Collectionsv2    `json:"collectionsv2"`
}

type Collectionsv2 struct {
	GroupedCollectionTermIDS map[string][]string `json:"groupedCollectionTermIds"`
	CollectionTermIDS        []string            `json:"collectionTermIds"`
}

type ProductLinks struct {
	Self Self `json:"self"`
}

type Self struct {
	Ref string `json:"ref"`
}

type ProductInfo struct {
	MerchProduct   MerchProductClass      `json:"merchProduct"`
	MerchPrice     MerchPrice             `json:"merchPrice"`
	Availability   Availability           `json:"availability"`
	ProductContent ProductContent         `json:"productContent"`
	ImageUrls      ImageUrls              `json:"imageUrls"`
	Skus           []Skus                 `json:"skus"`
	AvailableSkus  []AvailableSkusElement `json:"availableSkus"`
}

type Availability struct {
	ID           string       `json:"id"`
	ProductID    string       `json:"productId"`
	ResourceType string       `json:"resourceType"`
	Links        ProductLinks `json:"links"`
	Available    bool         `json:"available"`
}

type AvailableSkusElement struct {
	ID           string                    `json:"id"`
	ProductID    string                    `json:"productId"`
	ResourceType AvailableSkusResourceType `json:"resourceType"`
	Links        ProductLinks              `json:"links"`
	Available    bool                      `json:"available"`
	Level        Level                     `json:"level"`
	SkuID        string                    `json:"skuId"`
}

type ImageUrls struct {
	ProductImageURL string `json:"productImageUrl"`
}

type MerchPrice struct {
	ID               string        `json:"id"`
	SnapshotID       string        `json:"snapshotId"`
	ProductID        string        `json:"productId"`
	ParentID         string        `json:"parentId"`
	ParentType       Type          `json:"parentType"`
	ModificationDate string        `json:"modificationDate"`
	Country          Marketplace   `json:"country"`
	Msrp             int64         `json:"msrp"`
	FullPrice        int64         `json:"fullPrice"`
	CurrentPrice     int64         `json:"currentPrice"`
	Currency         string        `json:"currency"`
	Discounted       bool          `json:"discounted"`
	PromoInclusions  []interface{} `json:"promoInclusions"`
	PromoExclusions  []interface{} `json:"promoExclusions"`
	ResourceType     string        `json:"resourceType"`
	Links            ProductLinks  `json:"links"`
}

type MerchProductClass struct {
	ID                        string              `json:"id"`
	SnapshotID                string              `json:"snapshotId"`
	ModificationDate          string              `json:"modificationDate"`
	Status                    string              `json:"status"`
	MerchGroup                Marketplace         `json:"merchGroup"`
	StyleCode                 string              `json:"styleCode"`
	ColorCode                 string              `json:"colorCode"`
	StyleColor                string              `json:"styleColor"`
	PID                       string              `json:"pid"`
	CatalogID                 string              `json:"catalogId"`
	ProductGroupID            string              `json:"productGroupId"`
	Brand                     string              `json:"brand"`
	Channels                  []string            `json:"channels"`
	ConsumerChannels          []ConsumerChannel   `json:"consumerChannels"`
	LegacyCatalogIDS          []string            `json:"legacyCatalogIds"`
	Genders                   []string            `json:"genders"`
	SizeConverterID           string              `json:"sizeConverterId"`
	ValueAddedServices        []ValueAddedService `json:"valueAddedServices"`
	SportTags                 []string            `json:"sportTags"`
	ClassificationConcepts    []interface{}       `json:"classificationConcepts"`
	TaxonomyAttributes        []TaxonomyAttribute `json:"taxonomyAttributes"`
	CommerceCountryInclusions []interface{}       `json:"commerceCountryInclusions"`
	CommerceCountryExclusions []interface{}       `json:"commerceCountryExclusions"`
	AbTestValues              []interface{}       `json:"abTestValues"`
	ProductRollup             ProductRollup       `json:"productRollup"`
	QuantityLimit             int64               `json:"quantityLimit"`
	StyleType                 string              `json:"styleType"`
	ProductType               string              `json:"productType"`
	PublishType               string              `json:"publishType"`
	MainColor                 bool                `json:"mainColor"`
	IsImageAvailable          bool                `json:"isImageAvailable"`
	IsCopyAvailable           bool                `json:"isCopyAvailable"`
	IsAttributionApproved     bool                `json:"isAttributionApproved"`
	ExclusiveAccess           bool                `json:"exclusiveAccess"`
	CommercePublishDate       string              `json:"commercePublishDate"`
	CommerceStartDate         string              `json:"commerceStartDate"`
	ResourceType              Type                `json:"resourceType"`
	Links                     ProductLinks        `json:"links"`
}

type ConsumerChannel struct {
	ID           string `json:"id"`
	ResourceType string `json:"resourceType"`
}

type ProductRollup struct {
	Type string `json:"type"`
	Key  string `json:"key"`
}

type TaxonomyAttribute struct {
	ResourceType string   `json:"resourceType"`
	IDS          []string `json:"ids"`
}

type ValueAddedService struct {
	ID      string `json:"id"`
	VasType string `json:"vasType"`
}

type ProductContent struct {
	GlobalPID                      string        `json:"globalPid"`
	LangLocale                     string        `json:"langLocale"`
	ColorDescription               string        `json:"colorDescription"`
	Slug                           string        `json:"slug"`
	FullTitle                      string        `json:"fullTitle"`
	Title                          string        `json:"title"`
	Subtitle                       string        `json:"subtitle"`
	DescriptionHeading             string        `json:"descriptionHeading"`
	Description                    string        `json:"description"`
	TechSpec                       string        `json:"techSpec"`
	ManufacturingCountriesOfOrigin []interface{} `json:"manufacturingCountriesOfOrigin"`
	SizeChart                      string        `json:"sizeChart"`
	Colors                         []Color       `json:"colors"`
	BestFor                        []interface{} `json:"bestFor"`
	Athletes                       []interface{} `json:"athletes"`
	Widths                         []Width       `json:"widths"`
}

type Color struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type Width struct {
	Value          string `json:"value"`
	LocalizedValue string `json:"localizedValue"`
}

type Skus struct {
	ID                    string                 `json:"id"`
	SnapshotID            string                 `json:"snapshotId"`
	ProductID             string                 `json:"productId"`
	ParentID              string                 `json:"parentId"`
	ParentType            Type                   `json:"parentType"`
	CatalogSkuID          string                 `json:"catalogSkuId"`
	ModificationDate      string                 `json:"modificationDate"`
	MerchGroup            Marketplace            `json:"merchGroup"`
	StockKeepingUnitID    string                 `json:"stockKeepingUnitId"`
	Gtin                  string                 `json:"gtin"`
	NikeSize              string                 `json:"nikeSize"`
	SizeConversionID      string                 `json:"sizeConversionId"`
	CountrySpecifications []CountrySpecification `json:"countrySpecifications"`
	ResourceType          SkusResourceType       `json:"resourceType"`
	Links                 ProductLinks           `json:"links"`
}

type CountrySpecification struct {
	Country       Marketplace `json:"country"`
	LocalizedSize string      `json:"localizedSize"`
	TaxInfo       TaxInfo     `json:"taxInfo"`
}

type TaxInfo struct {
	CommodityCode string  `json:"commodityCode"`
	Vat           float64 `json:"vat"`
}

type PublishedContent struct {
	Preview            bool                       `json:"preview"`
	Marketplace        Marketplace                `json:"marketplace"`
	CollectionGroupID  string                     `json:"collectionGroupId"`
	CreatedDateTime    string                     `json:"createdDateTime"`
	Language           string                     `json:"language"`
	ViewStartDate      string                     `json:"viewStartDate"`
	Type               string                     `json:"type"`
	Version            string                     `json:"version"`
	Analytics          Analytics                  `json:"analytics"`
	Nodes              []PublishedContentNode     `json:"nodes"`
	PayloadType        string                     `json:"payloadType"`
	PublishStartDate   string                     `json:"publishStartDate"`
	SupportedLanguages []interface{}              `json:"supportedLanguages"`
	PublishEndDate     string                     `json:"publishEndDate"`
	SubType            string                     `json:"subType"`
	Links              PublishedContentLinks      `json:"links"`
	ID                 string                     `json:"id"`
	Properties         PublishedContentProperties `json:"properties"`
	ResourceType       string                     `json:"resourceType"`
}

type Analytics struct {
	HashKey string `json:"hashKey"`
}

type PublishedContentLinks struct {
	Self string `json:"self"`
}

type PublishedContentNode struct {
	Analytics  Analytics        `json:"analytics"`
	Nodes      []NodeNode       `json:"nodes"`
	SubType    string           `json:"subType"`
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Version    string           `json:"version"`
	Properties FluffyProperties `json:"properties"`
}

type NodeNode struct {
	Analytics  Analytics        `json:"analytics"`
	SubType    string           `json:"subType"`
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Version    string           `json:"version"`
	Properties PurpleProperties `json:"properties"`
}

type PurpleProperties struct {
	PortraitID   string   `json:"portraitId"`
	SquarishURL  string   `json:"squarishURL"`
	LandscapeID  string   `json:"landscapeId"`
	AltText      string   `json:"altText"`
	PortraitURL  string   `json:"portraitURL"`
	LandscapeURL string   `json:"landscapeURL"`
	Title        string   `json:"title"`
	Portrait     Portrait `json:"portrait"`
	Squarish     Portrait `json:"squarish"`
	SEOName      string   `json:"seoName"`
	SquarishID   string   `json:"squarishId"`
	Subtitle     string   `json:"subtitle"`
	ColorTheme   string   `json:"colorTheme"`
}

type Portrait struct {
	ID   string   `json:"id"`
	Type TypeEnum `json:"type"`
	URL  string   `json:"url"`
}

type FluffyProperties struct {
	ContainerType string `json:"containerType"`
	Loop          bool   `json:"loop"`
	Subtitle      string `json:"subtitle"`
	ColorTheme    string `json:"colorTheme"`
	AutoPlay      bool   `json:"autoPlay"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Speed         int64  `json:"speed"`
}

type PublishedContentProperties struct {
	ProductCard    ProductCard      `json:"productCard"`
	Custom         Custom           `json:"custom"`
	Publish        Publish          `json:"publish"`
	Subtitle       string           `json:"subtitle"`
	ConsumerLabels []interface{}    `json:"consumerLabels"`
	ThreadType     string           `json:"threadType"`
	Title          string           `json:"title"`
	SEO            SEO              `json:"seo"`
	Products       []ProductElement `json:"products"`
}

type Custom struct {
}

type ProductCard struct {
	SoftRef          bool                  `json:"softRef"`
	Transforms       []interface{}         `json:"transforms"`
	Language         string                `json:"language"`
	Variants         []interface{}         `json:"variants"`
	Type             string                `json:"type"`
	CreationDate     string                `json:"creationDate"`
	Version          string                `json:"version"`
	Translate        Custom                `json:"translate"`
	Analytics        Analytics             `json:"analytics"`
	Classifications  []interface{}         `json:"classifications"`
	TargetLanguages  []interface{}         `json:"targetLanguages"`
	ModificationDate string                `json:"modificationDate"`
	Nodes            []interface{}         `json:"nodes"`
	SubType          string                `json:"subType"`
	ID               string                `json:"id"`
	Properties       ProductCardProperties `json:"properties"`
	ResourceType     string                `json:"resourceType"`
}

type ProductCardProperties struct {
	PortraitID  string   `json:"portraitId"`
	SquarishURL string   `json:"squarishURL"`
	AltText     string   `json:"altText"`
	PortraitURL string   `json:"portraitURL"`
	Portrait    Portrait `json:"portrait"`
	Squarish    Portrait `json:"squarish"`
	SquarishID  string   `json:"squarishId"`
}

type ProductElement struct {
	ProductID  string `json:"productId"`
	StyleColor string `json:"styleColor"`
}

type Publish struct {
	CollectionGroups []string      `json:"collectionGroups"`
	Collections      []string      `json:"collections"`
	Countries        []Marketplace `json:"countries"`
}

type SEO struct {
	Slug string `json:"slug"`
}

type Search struct {
	ConceptIDS []string `json:"conceptIds"`
}

type Marketplace string

const (
	Us Marketplace = "US"
)

type Level string

const (
	High Level = "HIGH"
	Low  Level = "LOW"
	Oos  Level = "OOS"
)

type AvailableSkusResourceType string

const (
	AvailableSkus AvailableSkusResourceType = "availableSkus"
)

type Type string

const (
	MerchProduct Type = "merchProduct"
)

type SkusResourceType string

const (
	MerchSku SkusResourceType = "merchSku"
)

type TypeEnum string

const (
	TypeProduct TypeEnum = "product"
)
