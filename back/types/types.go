package types

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type GoogleSearchApiResponse struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		PreviousPage []struct {
			Title                  string `json:"title"`
			TotalResults           string `json:"totalResults"`
			SearchTerms            string `json:"searchTerms"`
			Count                  int    `json:"count"`
			StartIndex             int    `json:"startIndex"`
			StartPage              int    `json:"startPage"`
			Language               string `json:"language"`
			InputEncoding          string `json:"inputEncoding"`
			OutputEncoding         string `json:"outputEncoding"`
			Safe                   string `json:"safe"`
			Cx                     string `json:"cx"`
			Sort                   string `json:"sort"`
			Filter                 string `json:"filter"`
			Gl                     string `json:"gl"`
			Cr                     string `json:"cr"`
			GoogleHost             string `json:"googleHost"`
			DisableCnTwTranslation string `json:"disableCnTwTranslation"`
			Hq                     string `json:"hq"`
			Hl                     string `json:"hl"`
			SiteSearch             string `json:"siteSearch"`
			SiteSearchFilter       string `json:"siteSearchFilter"`
			ExactTerms             string `json:"exactTerms"`
			ExcludeTerms           string `json:"excludeTerms"`
			LinkSite               string `json:"linkSite"`
			OrTerms                string `json:"orTerms"`
			RelatedSite            string `json:"relatedSite"`
			DateRestrict           string `json:"dateRestrict"`
			LowRange               string `json:"lowRange"`
			HighRange              string `json:"highRange"`
			FileType               string `json:"fileType"`
			Rights                 string `json:"rights"`
			SearchType             string `json:"searchType"`
			ImgSize                string `json:"imgSize"`
			ImgType                string `json:"imgType"`
			ImgColorType           string `json:"imgColorType"`
			ImgDominantColor       string `json:"imgDominantColor"`
		} `json:"previousPage"`
		Request []struct {
			Title                  string `json:"title"`
			TotalResults           string `json:"totalResults"`
			SearchTerms            string `json:"searchTerms"`
			Count                  int    `json:"count"`
			StartIndex             int    `json:"startIndex"`
			StartPage              int    `json:"startPage"`
			Language               string `json:"language"`
			InputEncoding          string `json:"inputEncoding"`
			OutputEncoding         string `json:"outputEncoding"`
			Safe                   string `json:"safe"`
			Cx                     string `json:"cx"`
			Sort                   string `json:"sort"`
			Filter                 string `json:"filter"`
			Gl                     string `json:"gl"`
			Cr                     string `json:"cr"`
			GoogleHost             string `json:"googleHost"`
			DisableCnTwTranslation string `json:"disableCnTwTranslation"`
			Hq                     string `json:"hq"`
			Hl                     string `json:"hl"`
			SiteSearch             string `json:"siteSearch"`
			SiteSearchFilter       string `json:"siteSearchFilter"`
			ExactTerms             string `json:"exactTerms"`
			ExcludeTerms           string `json:"excludeTerms"`
			LinkSite               string `json:"linkSite"`
			OrTerms                string `json:"orTerms"`
			RelatedSite            string `json:"relatedSite"`
			DateRestrict           string `json:"dateRestrict"`
			LowRange               string `json:"lowRange"`
			HighRange              string `json:"highRange"`
			FileType               string `json:"fileType"`
			Rights                 string `json:"rights"`
			SearchType             string `json:"searchType"`
			ImgSize                string `json:"imgSize"`
			ImgType                string `json:"imgType"`
			ImgColorType           string `json:"imgColorType"`
			ImgDominantColor       string `json:"imgDominantColor"`
		} `json:"request"`
		NextPage []struct {
			Title                  string `json:"title"`
			TotalResults           string `json:"totalResults"`
			SearchTerms            string `json:"searchTerms"`
			Count                  int    `json:"count"`
			StartIndex             int    `json:"startIndex"`
			StartPage              int    `json:"startPage"`
			Language               string `json:"language"`
			InputEncoding          string `json:"inputEncoding"`
			OutputEncoding         string `json:"outputEncoding"`
			Safe                   string `json:"safe"`
			Cx                     string `json:"cx"`
			Sort                   string `json:"sort"`
			Filter                 string `json:"filter"`
			Gl                     string `json:"gl"`
			Cr                     string `json:"cr"`
			GoogleHost             string `json:"googleHost"`
			DisableCnTwTranslation string `json:"disableCnTwTranslation"`
			Hq                     string `json:"hq"`
			Hl                     string `json:"hl"`
			SiteSearch             string `json:"siteSearch"`
			SiteSearchFilter       string `json:"siteSearchFilter"`
			ExactTerms             string `json:"exactTerms"`
			ExcludeTerms           string `json:"excludeTerms"`
			LinkSite               string `json:"linkSite"`
			OrTerms                string `json:"orTerms"`
			RelatedSite            string `json:"relatedSite"`
			DateRestrict           string `json:"dateRestrict"`
			LowRange               string `json:"lowRange"`
			HighRange              string `json:"highRange"`
			FileType               string `json:"fileType"`
			Rights                 string `json:"rights"`
			SearchType             string `json:"searchType"`
			ImgSize                string `json:"imgSize"`
			ImgType                string `json:"imgType"`
			ImgColorType           string `json:"imgColorType"`
			ImgDominantColor       string `json:"imgDominantColor"`
		} `json:"nextPage"`
	} `json:"queries"`
	Promotions []interface{} `json:"promotions"`
	Context    struct {
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Spelling struct {
		CorrectedQuery     string `json:"correctedQuery"`
		HTMLCorrectedQuery string `json:"htmlCorrectedQuery"`
	} `json:"spelling"`
	Items []Items `json:"items"`
}

type Items struct {
	DisplayLink      string `json:"displayLink"`
	FormattedURL     string `json:"formattedUrl"`
	HTMLFormattedURL string `json:"htmlFormattedUrl"`
	HTMLSnippet      string `json:"htmlSnippet"`
	HTMLTitle        string `json:"htmlTitle"`
	Kind             string `json:"kind"`
	Link             string `json:"link"`
	Pagemap          struct {
		CseImage []struct {
			Src string `json:"src"`
		} `json:"cse_image"`
		CseThumbnail []struct {
			Height string `json:"height"`
			Src    string `json:"src"`
			Width  string `json:"width"`
		} `json:"cse_thumbnail"`
		Metatags []interface{} `json:"metatags"`
	} `json:"pagemap"`
	Snippet string `json:"snippet"`
	Title   string `json:"title"`
}

type Search struct {
	// User        primitive.ObjectID `bson:"user_id"`
	ID          primitive.ObjectID `bson:"_id"         json:"id"`
	Term        string             `bson:"term"        json:"term"`
	Competitors []string           `bson:"competitors" json:"competitors"`
	CreatedAt   time.Time          `bson:"created_at"  json:"created_at"`
}

type SearchStore interface {
	CreateSearch(context.Context, *CreateSearchPayload) error
	GetSearches(context.Context) ([]*Search, error)
	GetSearchByID(context.Context, primitive.ObjectID) (*Search, error)
}

type TestAddData struct {
	ID        primitive.ObjectID `bson:"_id"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type TestStore interface {
	AddData(context.Context, *TestAddData) error
}

type CreateSearchPayload struct {
	Term        string   `json:"term"`
	Competitors []string `json:"competitors"`
}
