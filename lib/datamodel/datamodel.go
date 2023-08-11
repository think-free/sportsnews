package datamodel

// ICNews is our internal data model that can be used by any providers
// NOTE : For this exercice, we could have used the same data model as the upstream provider with, but we want to keep the providers independent
// and this facilitates the implementation of new providers
type ICNews struct {
	ID          string   `bson:"id,omitempty" json:"id"`
	TeamID      string   `bson:"teamId,omitempty" json:"teamId"`
	OptaMatchID string   `bson:"optaMatchId,omitempty" json:"optaMatchId"`
	Title       string   `bson:"title,omitempty" json:"title"`
	Type        []string `bson:"type,omitempty" json:"type"`
	Teaser      string   `bson:"teaser,omitempty" json:"teaser"`
	Content     string   `bson:"content,omitempty" json:"content"`
	URL         string   `bson:"url,omitempty" json:"url"`
	ImageURL    string   `bson:"imageUrl,omitempty" json:"imageUrl"`
	GalleryUrls string   `bson:"galleryUrls,omitempty" json:"galleryUrls"`
	VideoURL    string   `bson:"videoUrl,omitempty" json:"videoUrl"`
	Published   string   `bson:"published,omitempty" json:"published"`
}
