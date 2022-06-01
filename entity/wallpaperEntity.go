package entity

type AutoGenerated []struct {
	ID                     string        `json:"id"`
	CreatedAt              string        `json:"created_at"`
	UpdatedAt              string        `json:"updated_at"`
	PromotedAt             string        `json:"promoted_at"`
	Width                  int           `json:"width"`
	Height                 int           `json:"height"`
	Color                  string        `json:"color"`
	BlurHash               string        `json:"blur_hash"`
	Description            string        `json:"description"`
	AltDescription         string        `json:"alt_description"`
	Urls                   Urls          `json:"urls"`
	Links                  Links         `json:"links"`
	Categories             []interface{} `json:"categories"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	User                   User          `json:"user"`
}
type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}
type Links struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}