package model

type NewsArticles struct {
	Header
	Items struct {
		Articles []New `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}

type NewsArticle struct {
	Header
	Article New `xml:"NewsArticle"`
}

type Header struct {
	ClubName       string `xml:"ClubName"`
	ClubWebsiteURL string `xml:"ClubWebsiteURL"`
}

type New struct {
	ArticleURL        string `xml:"ArticleURL"`
	ID                string `xml:"NewsArticleID"`
	PublishDate       string `xml:"PublishDate"`
	Taxonomies        string `xml:"Taxonomies"`
	TeaserText        string `xml:"TeaserText"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL"`
	Title             string `xml:"Title"`
	BodyText          string `xml:"BodyText"`
	GalleryImageURLs  string `xml:"GalleryImageURLs"`
	VideoURL          string `xml:"VideoURL"`
	OptaMatchID       string `xml:"OptaMatchId"`
	LastUpdateDate    string `xml:"LastUpdateDate"`
	IsPublished       string `xml:"IsPublished"`
}
