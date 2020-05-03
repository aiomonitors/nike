package types

type NewProducts struct {
	Pages   Pages    `json:"pages"`
	Objects []Object `json:"objects"`
}

type Object struct {
	ID                string                 `json:"id"`
	ChannelID         string                 `json:"channelId"`
	ChannelName       Channel                `json:"channelName"`
	Marketplace       Marketplace            `json:"marketplace"`
	Language          Language               `json:"language"`
	LastFetchTime     string                 `json:"lastFetchTime"`
	PublishedContent  ObjectPublishedContent `json:"publishedContent"`
	ProductInfo       []ObjectProductInfo    `json:"productInfo"`
	Search            Search                 `json:"search"`
	CollectionTermIDS []string               `json:"collectionTermIds"`
	ResourceType      ObjectResourceType     `json:"resourceType"`
	Rollup            Rollup                 `json:"rollup"`
	Links             Links                  `json:"links"`
	Collectionsv2     Collectionsv2          `json:"collectionsv2"`
}

type Collectionsv2 struct {
	GroupedCollectionTermIDS GroupedCollectionTermIDS `json:"groupedCollectionTermIds"`
	CollectionTermIDS        []string                 `json:"collectionTermIds"`
}

type GroupedCollectionTermIDS struct {
	D9A5Bc424B9C4976858AF159Cf99C647 []string `json:"d9a5bc42-4b9c-4976-858a-f159cf99c647"`
}

type Links struct {
	Self Self `json:"self"`
}

type Self struct {
	Ref string `json:"ref"`
}

type ObjectProductInfo struct {
	MerchProduct       MerchProductClass   `json:"merchProduct"`
	MerchPrice         MerchPriceClass     `json:"merchPrice"`
	Availability       Availability        `json:"availability"`
	ProductContent     ProductContent      `json:"productContent"`
	ImageUrls          ImageUrls           `json:"imageUrls"`
	LaunchView         *LaunchView         `json:"launchView,omitempty"`
	CustomizedPreBuild *CustomizedPreBuild `json:"customizedPreBuild,omitempty"`
	ContentCopy        *ContentCopy        `json:"contentCopy,omitempty"`
}

type Availability struct {
	ID           string                   `json:"id"`
	ProductID    string                   `json:"productId"`
	ResourceType AvailabilityResourceType `json:"resourceType"`
	Links        Links                    `json:"links"`
	Available    bool                     `json:"available"`
}

type ContentCopy struct {
	DisplayNames DisplayNames `json:"displayNames"`
}

type DisplayNames struct {
	CustomizationTwoLineName TwoLineName `json:"customizationTwoLineName"`
	TwoLineName              TwoLineName `json:"twoLineName"`
}

type TwoLineName struct {
	Line2 Line2 `json:"line2"`
	Line1 Line1 `json:"line1"`
}

type CustomizedPreBuild struct {
	ID                        string                   `json:"id"`
	DesignID                  string                   `json:"designId"`
	ProductID                 string                   `json:"productId"`
	CatalogID                 string                   `json:"catalogId"`
	BuilderID                 string                   `json:"builderId"`
	Status                    Status                   `json:"status"`
	MerchGroup                Marketplace              `json:"merchGroup"`
	CommerceStartDate         string                   `json:"commerceStartDate"`
	ActiveDate                string                   `json:"activeDate"`
	CommerceCountryInclusions []interface{}            `json:"commerceCountryInclusions"`
	CommerceCountryExclusions []interface{}            `json:"commerceCountryExclusions"`
	ModificationDate          string                   `json:"modificationDate"`
	CreationDate              string                   `json:"creationDate"`
	Groups                    []Group                  `json:"groups"`
	TaxonomyConcepts          []Taxonomy               `json:"taxonomyConcepts"`
	ResourceType              string                   `json:"resourceType"`
	Links                     Links                    `json:"links"`
	Legacy                    CustomizedPreBuildLegacy `json:"legacy"`
	Imagery                   []Imagery                `json:"imagery"`
	Selections                []Selection              `json:"selections"`
	EnteredText               []interface{}            `json:"enteredText"`
}

type Group struct {
	DefaultGroup           bool        `json:"defaultGroup"`
	DefaultPrebuildInGroup bool        `json:"defaultPrebuildInGroup"`
	VisibleInSearch        bool        `json:"visibleInSearch"`
	VisibleInCSR           bool        `json:"visibleInCSR"`
	Legacy                 GroupLegacy `json:"legacy"`
}

type GroupLegacy struct {
	Piid string `json:"piid"`
	Slug string `json:"slug"`
}

type Imagery struct {
	ViewCode               ViewCode `json:"viewCode"`
	ViewNumber             string   `json:"viewNumber"`
	ImageSourceURL         string   `json:"imageSourceURL"`
	ImageSourceURLTemplate string   `json:"imageSourceURLTemplate"`
}

type CustomizedPreBuildLegacy struct {
	StyleCode   string      `json:"styleCode"`
	ColorCode   string      `json:"colorCode"`
	Pbid        string      `json:"pbid"`
	PathName    string      `json:"pathName"`
	PathVersion int64       `json:"pathVersion"`
	ProductID   string      `json:"productId"`
	Locale      Language    `json:"locale"`
	Country     Marketplace `json:"country"`
	Size        Translate   `json:"size"`
}

type Translate struct {
}

type Selection struct {
	QuestionPath string `json:"questionPath"`
	AnswerID     string `json:"answerId"`
}

type Taxonomy struct {
	ResourceType TaxonomyConceptResourceType `json:"resourceType"`
	IDS          []string                    `json:"ids"`
}

type ImageUrls struct {
	ProductImageURL string `json:"productImageUrl"`
}

type LaunchView struct {
	ID             string `json:"id"`
	ProductID      string `json:"productId"`
	Method         string `json:"method"`
	StartEntryDate string `json:"startEntryDate"`
}

type MerchPriceClass struct {
	Country       Marketplace            `json:"country"`
	FullPrice     int64                  `json:"fullPrice"`
	CurrentPrice  int64                  `json:"currentPrice"`
	EmployeePrice *int64                 `json:"employeePrice,omitempty"`
	Currency      Currency               `json:"currency"`
	Discounted    bool                   `json:"discounted"`
	ResourceType  MerchPriceResourceType `json:"resourceType"`
	Links         Links                  `json:"links"`
}

type MerchProductClass struct {
	ID                        string                   `json:"id"`
	SnapshotID                string                   `json:"snapshotId"`
	ModificationDate          string                   `json:"modificationDate"`
	Status                    Status                   `json:"status"`
	MerchGroup                Marketplace              `json:"merchGroup"`
	StyleCode                 string                   `json:"styleCode"`
	ColorCode                 string                   `json:"colorCode"`
	StyleColor                string                   `json:"styleColor"`
	PID                       string                   `json:"pid"`
	CatalogID                 string                   `json:"catalogId"`
	ProductGroupID            *string                  `json:"productGroupId,omitempty"`
	Brand                     Brand                    `json:"brand"`
	Channels                  []Channel                `json:"channels"`
	ConsumerChannels          []ConsumerChannel        `json:"consumerChannels"`
	LegacyCatalogIDS          []string                 `json:"legacyCatalogIds"`
	Genders                   []Gender                 `json:"genders"`
	SizeConverterID           string                   `json:"sizeConverterId"`
	ValueAddedServices        []ValueAddedService      `json:"valueAddedServices"`
	SportTags                 []SportTag               `json:"sportTags"`
	ClassificationConcepts    []interface{}            `json:"classificationConcepts"`
	TaxonomyAttributes        []Taxonomy               `json:"taxonomyAttributes"`
	CommerceCountryInclusions []interface{}            `json:"commerceCountryInclusions"`
	CommerceCountryExclusions []interface{}            `json:"commerceCountryExclusions"`
	AbTestValues              []interface{}            `json:"abTestValues"`
	ProductRollup             ProductRollup            `json:"productRollup"`
	QuantityLimit             int64                    `json:"quantityLimit"`
	StyleType                 StyleType                `json:"styleType"`
	ProductType               ProductType              `json:"productType"`
	PublishType               *PublishType             `json:"publishType,omitempty"`
	MainColor                 bool                     `json:"mainColor"`
	IsImageAvailable          bool                     `json:"isImageAvailable"`
	IsCopyAvailable           bool                     `json:"isCopyAvailable"`
	IsAttributionApproved     bool                     `json:"isAttributionApproved"`
	ExclusiveAccess           bool                     `json:"exclusiveAccess"`
	HardLaunch                *bool                    `json:"hardLaunch,omitempty"`
	CommercePublishDate       string                   `json:"commercePublishDate"`
	CommerceStartDate         string                   `json:"commerceStartDate"`
	ResourceType              MerchProductResourceType `json:"resourceType"`
	Links                     Links                    `json:"links"`
	InventoryOverride         *bool                    `json:"inventoryOverride,omitempty"`
	HideFromCSR               *bool                    `json:"hideFromCSR,omitempty"`
	SoftLaunchDate            *string                  `json:"softLaunchDate,omitempty"`
	CommerceEndDate           *string                  `json:"commerceEndDate,omitempty"`
	Customization             *Customization           `json:"customization,omitempty"`
	NotifyMeIndicator         *bool                    `json:"notifyMeIndicator,omitempty"`
}

type ConsumerChannel struct {
	ID           string                      `json:"id"`
	ResourceType ConsumerChannelResourceType `json:"resourceType"`
}

type Customization struct {
	NikeIDSlug string `json:"nikeIdSlug"`
}

type ProductRollup struct {
	Type ProductRollupType `json:"type"`
	Key  string            `json:"key"`
}

type ValueAddedService struct {
	ID      string  `json:"id"`
	VasType VasType `json:"vasType"`
}

type ProductContent struct {
	ColorDescription string  `json:"colorDescription"`
	Slug             string  `json:"slug"`
	Title            string  `json:"title"`
	Subtitle         string  `json:"subtitle"`
	Colors           []Color `json:"colors"`
}

type Color struct {
	Type ColorType `json:"type"`
	Name string    `json:"name"`
	Hex  string    `json:"hex"`
}

type ObjectPublishedContent struct {
	PublishStartDate string           `json:"publishStartDate"`
	CreatedDateTime  string           `json:"createdDateTime"`
	PublishEndDate   string           `json:"publishEndDate"`
	ViewStartDate    string           `json:"viewStartDate"`
	Properties       PurpleProperties `json:"properties"`
}

type PurpleProperties struct {
	Custom         Custom           `json:"custom"`
	Title          string           `json:"title"`
	Products       []ProductElement `json:"products"`
	ProductCard    ProductCard      `json:"productCard"`
	Publish        Publish          `json:"publish"`
	Subtitle       string           `json:"subtitle"`
	ConsumerLabels []interface{}    `json:"consumerLabels"`
	ThreadType     ThreadType       `json:"threadType"`
	SEO            SEO              `json:"seo"`
	RelatedThreads []interface{}    `json:"relatedThreads"`
}

type Custom struct {
	HideFromUpcoming []interface{} `json:"hideFromUpcoming"`
}

type ProductCard struct {
	Properties       ProductCardProperties   `json:"properties"`
	Nodes            []interface{}           `json:"nodes"`
	ID               string                  `json:"id"`
	Type             ProductCardType         `json:"type"`
	SubType          SubType                 `json:"subType"`
	Version          string                  `json:"version"`
	CreationDate     string                  `json:"creationDate"`
	ModificationDate string                  `json:"modificationDate"`
	Classifications  []interface{}           `json:"classifications"`
	Transforms       []interface{}           `json:"transforms"`
	Translate        Translate               `json:"translate"`
	Language         Language                `json:"language"`
	TargetLanguages  []interface{}           `json:"targetLanguages"`
	Variants         []interface{}           `json:"variants"`
	SoftRef          bool                    `json:"softRef"`
	Analytics        Analytics               `json:"analytics"`
	ResourceType     ProductCardResourceType `json:"resourceType"`
}

type Analytics struct {
	HashKey string `json:"hashKey"`
}

type ProductCardProperties struct {
	SquarishURL string    `json:"squarishURL"`
	PortraitID  *string   `json:"portraitId,omitempty"`
	AltText     string    `json:"altText"`
	PortraitURL *string   `json:"portraitURL,omitempty"`
	Squarish    Portrait  `json:"squarish"`
	Portrait    *Portrait `json:"portrait,omitempty"`
	SquarishID  string    `json:"squarishId"`
}

type Portrait struct {
	ID   string       `json:"id"`
	Type PortraitType `json:"type"`
	URL  string       `json:"url"`
}

type ProductElement struct {
	ProductID  string `json:"productId"`
	StyleColor string `json:"styleColor"`
}

type Publish struct {
	CollectionGroups []string `json:"collectionGroups"`
	Collections      []string `json:"collections"`
	Countries        []string `json:"countries"`
}

type SEO struct {
	Slug string `json:"slug"`
}

type Rollup struct {
	TotalThreads int64           `json:"totalThreads"`
	Threads      []ThreadElement `json:"threads"`
}

type ThreadElement struct {
	ID                string                 `json:"id"`
	ChannelID         string                 `json:"channelId"`
	ChannelName       Channel                `json:"channelName"`
	Marketplace       Marketplace            `json:"marketplace"`
	Language          Language               `json:"language"`
	LastFetchTime     string                 `json:"lastFetchTime"`
	PublishedContent  ThreadPublishedContent `json:"publishedContent"`
	ProductInfo       []ThreadProductInfo    `json:"productInfo"`
	Search            Search                 `json:"search"`
	CollectionTermIDS []string               `json:"collectionTermIds"`
	ResourceType      ObjectResourceType     `json:"resourceType"`
	Links             Links                  `json:"links"`
	Collectionsv2     Collectionsv2          `json:"collectionsv2"`
}

type ThreadProductInfo struct {
	MerchProduct       MerchProductClass   `json:"merchProduct"`
	MerchPrice         MerchPriceClass     `json:"merchPrice"`
	Availability       Availability        `json:"availability"`
	ProductContent     ProductContent      `json:"productContent"`
	ImageUrls          ImageUrls           `json:"imageUrls"`
	CustomizedPreBuild *CustomizedPreBuild `json:"customizedPreBuild,omitempty"`
	ContentCopy        *ContentCopy        `json:"contentCopy,omitempty"`
}

type ThreadPublishedContent struct {
	PublishStartDate string           `json:"publishStartDate"`
	CreatedDateTime  string           `json:"createdDateTime"`
	PublishEndDate   string           `json:"publishEndDate"`
	ViewStartDate    string           `json:"viewStartDate"`
	Properties       FluffyProperties `json:"properties"`
}

type FluffyProperties struct {
	ThreadType     ThreadType       `json:"threadType"`
	Publish        Publish          `json:"publish"`
	Custom         Translate        `json:"custom"`
	Title          string           `json:"title"`
	Products       []ProductElement `json:"products"`
	ProductCard    ProductCard      `json:"productCard"`
	Subtitle       string           `json:"subtitle"`
	ConsumerLabels []interface{}    `json:"consumerLabels"`
	SEO            SEO              `json:"seo"`
}

type Search struct {
	ConceptIDS []string `json:"conceptIds"`
}

type Pages struct {
	Prev           string `json:"prev"`
	Next           string `json:"next"`
	TotalPages     int64  `json:"totalPages"`
	TotalResources int64  `json:"totalResources"`
}

type Channel string

const (
	NikeApp              Channel = "NikeApp"
	NikeCOM              Channel = "Nike.com"
	NikeStoreExperiences Channel = "Nike Store Experiences"
	Snkrs                Channel = "SNKRS"
)

type Language string

const (
	En Language = "en"
)

type Marketplace string

const (
	Us Marketplace = "US"
)

type AvailabilityResourceType string

const (
	AvailableProducts AvailabilityResourceType = "availableProducts"
)

type Line1 string

const (
	NikeBlazerMidByYou Line1 = "Nike Blazer Mid By You"
)

type Line2 string

const (
	CustomMenSShoe Line2 = "Custom Men's Shoe"
)

type ViewCode string

const (
	A ViewCode = "A"
	B ViewCode = "B"
	C ViewCode = "C"
	D ViewCode = "D"
	E ViewCode = "E"
	F ViewCode = "F"
)

type Status string

const (
	Active Status = "ACTIVE"
)

type TaxonomyConceptResourceType string

const (
	MerchTaxonomyAttributes TaxonomyConceptResourceType = "merch/taxonomy_attributes"
)

type Currency string

const (
	Usd Currency = "USD"
)

type MerchPriceResourceType string

const (
	MerchPrice MerchPriceResourceType = "merchPrice"
)

type Brand string

const (
	Jordan         Brand = "Jordan"
	Nike           Brand = "Nike"
	NikeSportswear Brand = "Nike Sportswear"
)

type ConsumerChannelResourceType string

const (
	GlobalizationConsumerChannels ConsumerChannelResourceType = "globalization/consumer_channels"
)

type Gender string

const (
	Men   Gender = "MEN"
	Women Gender = "WOMEN"
)

type ProductRollupType string

const (
	Standard   ProductRollupType = "Standard"
	WidthGroup ProductRollupType = "WidthGroup"
)

type ProductType string

const (
	Footwear ProductType = "FOOTWEAR"
)

type PublishType string

const (
	Flow   PublishType = "FLOW"
	Launch PublishType = "LAUNCH"
)

type MerchProductResourceType string

const (
	MerchProduct MerchProductResourceType = "merchProduct"
)

type SportTag string

const (
	Basketball SportTag = "Basketball"
	Lifestyle  SportTag = "Lifestyle"
	Running    SportTag = "Running"
	Training   SportTag = "Training"
)

type StyleType string

const (
	Inline StyleType = "INLINE"
	Nikeid StyleType = "NIKEID"
)

type VasType string

const (
	FullCustomization VasType = "FULL_CUSTOMIZATION"
	GiftMessage       VasType = "GIFT_MESSAGE"
	GiftWrap          VasType = "GIFT_WRAP"
)

type ColorType string

const (
	Logo      ColorType = "LOGO"
	Primary   ColorType = "PRIMARY"
	Secondary ColorType = "SECONDARY"
	Simple    ColorType = "SIMPLE"
	Tertiary  ColorType = "TERTIARY"
)

type PortraitType string

const (
	Product PortraitType = "product"
)

type ProductCardResourceType string

const (
	ContentNode ProductCardResourceType = "content/node"
)

type SubType string

const (
	Image SubType = "image"
)

type ProductCardType string

const (
	Card ProductCardType = "card"
)

type ThreadType string

const (
	NikeidSoldier ThreadType = "nikeid_soldier"
	Officer       ThreadType = "officer"
	Soldier       ThreadType = "soldier"
)

type ObjectResourceType string

const (
	Thread ObjectResourceType = "thread"
)
