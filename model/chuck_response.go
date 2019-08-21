package model

type ChuckResponseModel struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
	Value   string `json:"value"`
}
