package servicemodel

type Level struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FetchLevelName struct {
	Name string `json:"name"`
}
