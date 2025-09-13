package dto

type Paginate struct {
	Page     int `json:"page"`
	PerPage int `json:"perpage"`
	Total    int `json:"total"`
	TotalPage int `json:"totalpage"`
	
}