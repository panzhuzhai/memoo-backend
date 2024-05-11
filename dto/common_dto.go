package dto

type PageDto struct {
	PageNumber int64 `form:"pageNumber" json:"pageNumber"`
	PageSize   int   `form:"pageSize" json:"pageSize"`
}
