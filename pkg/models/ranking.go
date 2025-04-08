package models

type RankingListResponse struct {
	Total   int           `json:"total"`
	Ranking []RankingItem `json:"ranking"`
}

type RankingItem struct {
	Index  uint
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Points int    `json:"points"`
}
