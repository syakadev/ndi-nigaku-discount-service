package resmodel

type PaginationData struct {
	TotalData int `json:"total_data"`
	Page      int `json:"page"`
	Size      int `json:"size"`
}

type NoDataResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DatasResponse struct {
	Success    bool          `json:"success"`
	Message    string        `json:"message"`
	Datas      []interface{} `json:"datas"`
	Pagination interface{}   `json:"pagination"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
