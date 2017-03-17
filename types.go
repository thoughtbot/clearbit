package clearbit

import (
	"encoding/json"
	"strconv"
)

// CombinedResponse is the union of a Person and a Company.
//
// https://clearbit.com/docs?javascript#enrichment-api-combined-api
type CombinedResponse struct {
	Company Company `json:"company"`
	Person  Person  `json:"person"`
}

// Company describes company data known to Clearbit.
//
// https://clearbit.com/docs#enrichment-api-company-api-attributes
type Company struct {
	Description   string         `json:"description"`
	Domain        string         `json:"domain"`
	DomainAliases []string       `json:"domainAliases"`
	EmailProvider bool           `json:"emailProvider"`
	FoundedDate   string         `json:"foundedDate"`
	Geo           Geo            `json:"geo"`
	ID            string         `json:"id"`
	LegalName     string         `json:"legalName"`
	Location      string         `json:"location"`
	Logo          string         `json:"logo"`
	Metrics       CompanyMetrics `json:"metrics"`
	Name          string         `json:"name"`
	Phone         string         `json:"phone"`
	Site          CompanySite    `json:"site"`
	Tags          []string       `json:"tags"`
	Tech          []string       `json:"tech"`
	TimeZone      string         `json:"timeZone"`
	Type          string         `json:"type"`
	URL           string         `json:"url"`
	UTCOffset     int            `json:"utcOffset"`

	AngelList  AngelListProfile  `json:"angellist"`
	CrunchBase CrunchBaseProfile `json:"crunchbase"`
	Facebook   FacebookProfile   `json:"facebook"`
	LinkedIn   LinkedInProfile   `json:"linkedin"`
	Twitter    TwitterProfile    `json:"twitter"`
}

// CompanySite describes information about and from a company's site.
type CompanySite struct {
	EmailAddresses  []string `json:"emailAddresses"`
	H1              string   `json:"h1"`
	MetaAuthor      string   `json:"metaAuthor"`
	MetaDescription string   `json:"metaDescription"`
	PhoneNumbers    []string `json:"phoneNumbers"`
	Title           string   `json:"title"`
	URL             string   `json:"url"`
}

// CompanyMetrics describes a company's metrics across different sites,
// services, and other metrics.
type CompanyMetrics struct {
	AlexaGlobalRank int `json:"alexaGlobalRank"`
	AlexaUsRank     int `json:"alexaUsRank"`
	AnnualRevenue   int `json:"annualRevenue"`
	Employees       int `json:"employees"`
	GoogleRank      int `json:"googleRank"`
	MarketCap       int `json:"marketCap"`
	Raised          int `json:"raised"`
}

// Person describes person data known to Clearbit.
//
// https://clearbit.com/docs#enrichment-api-person-api-attributes
type Person struct {
	Avatar     string     `json:"avatar"`
	Bio        string     `json:"bio"`
	Email      string     `json:"email"`
	Employment Employment `json:"employment"`
	Fuzzy      bool       `json:"fuzzy"`
	Gender     string     `json:"gender"`
	Geo        Geo        `json:"geo"`
	ID         string     `json:"id"`
	Location   string     `json:"location"`
	Name       Name       `json:"name"`
	Site       string     `json:"site"`
	TimeZone   string     `json:"timeZone"`
	UTCOffset  int        `json:"utcOffset"`

	AboutMe    AboutMeProfile    `json:"aboutme"`
	AngelList  AngelListProfile  `json:"angellist"`
	Facebook   FacebookProfile   `json:"facebook"`
	GitHub     GitHubProfile     `json:"github"`
	GooglePlus GooglePlusProfile `json:"googleplus"`
	Gravatar   GravatarProfile   `json:"gravatar"`
	Twitter    TwitterProfile    `json:"twitter"`
	LinkedIn   LinkedInProfile   `json:"linkedin"`
}

// Prospect describes the basic contact information for a person
// found through the Prospector API.
//
// https://clearbit.com/docs#prospector-api-person-search-response-body
type Prospect struct {
	ID    string `json:"id"`
	Name  Name   `json:"name"`
	Title string `json:"title"`
	Email string `json:"email"`
}

// Name describes a person's name.
type Name struct {
	FamilyName string `json:"familyName"`
	FullName   string `json:"fullName"`
	GivenName  string `json:"givenName"`
}

// Employment describes a person's current employment status.
type Employment struct {
	Domain    string `json:"domain"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Seniority string `json:"seniority"`
	Title     string `json:"title"`
}

// Geo represents a person or comapny's location.
//
// Some fields are only available for companies.
type Geo struct {
	City         string  `json:"city"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"countryCode"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	PostalCode   string  `json:"postalCode"`
	State        string  `json:"state"`
	StateCode    string  `json:"stateCode"`
	StreetName   string  `json:"streetName"`
	StreetNumber string  `json:"streetNumber"`
	SubPremise   string  `json:"subPremise"`
}

// AboutMeProfile describes a person's About.me profile.
type AboutMeProfile struct {
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
	Handle string `json:"handle"`
}

// AngelListProfile describes a person or company's AngelList profile.
type AngelListProfile struct {
	Avatar    string `json:"avatar"`
	Bio       string `json:"bio"`
	Blog      string `json:"blog"`
	Followers int    `json:"followers"`
	Handle    string `json:"handle"`
	ID        int    `json:"id"`
	Site      string `json:"site"`
}

// CrunchBaseProfile describes a company's CrunchBase profile.
type CrunchBaseProfile struct {
	Handle string `json:"handle"`
}

// FacebookProfile describes a person or company's Facebook profile.
type FacebookProfile struct {
	Handle string `json:"handle"`
}

// GitHubProfile describes a person's GitHub profile.
type GitHubProfile struct {
	Avatar    string `json:"avatar"`
	Blog      string `json:"blog"`
	Company   string `json:"company"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Handle    string `json:"handle"`
}

// GooglePlusProfile describes a person's Google+ profile.
type GooglePlusProfile struct {
	Handle string `json:"handle"`
}

// GravatarProfile describes a person's public Gravatar profile.
type GravatarProfile struct {
	DefaultAvatarURL string           `json:"avatar"`
	Avatars          []GravatarAvatar `json:"avatars"`
	Handle           string           `json:"handle"`
	URLs             []GravatarURL    `json:"urls"`
}

// GravatarAvatar describes an avatar associated with a person's Gravatar
// profile. For example, the type may be "thumbnail" and the URL to view it.
type GravatarAvatar struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// GravatarURL describes a URL associated with a person's Gravatar profile,
// like a link to their blog.
type GravatarURL struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// LinkedInProfile describes a person or company's public LinkedIn profile.
type LinkedInProfile struct {
	Handle string `json:"handle"`
}

// TwitterProfile describes a person or company's public Twitter profile.
type TwitterProfile struct {
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	Followers int       `json:"followers"`
	Following int       `json:"following"`
	Handle    string    `json:"handle"`
	ID        TwitterID `json:"id"`
	Location  string    `json:"location"`
	Site      string    `json:"site"`
}

// TwitterID is a person or company's Twitter profile ID.
//
// It is stored as a string, though it can be also be an integer
// in Clearbit API responses.
type TwitterID string

// UnmarshalJSON implements json.Unmarshaler.
func (id *TwitterID) UnmarshalJSON(data []byte) error {
	if err := id.unmarshalString(data); err != nil {
		return id.unmarshalInt(data)
	}

	return nil
}

func (id *TwitterID) unmarshalString(data []byte) error {
	return json.Unmarshal(data, (*string)(id))
}

func (id *TwitterID) unmarshalInt(data []byte) error {
	var intID int

	if err := json.Unmarshal(data, &intID); err != nil {
		return err
	}

	*id = TwitterID(strconv.Itoa(intID))
	return nil
}

// ErrorResponse describes the structure of non-successful responses from the
// Clearbit API.
type ErrorResponse struct {
	Error `json:"error"`
}

// Error describes an error returned by Clearbit's API.
//
// For a list of error types, see https://clearbit.com/docs#errors-error-types.
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
