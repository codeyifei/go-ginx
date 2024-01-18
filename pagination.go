package ginx

type IPagination interface {
	Page() int
	PageSize() int
}

type PaginationRequest struct {
	Page     int `form:"page" default:"1"`
	PageSize int `form:"page_size" default:"20"`
}

func (r *PaginationRequest) ToMeta(total int) *PaginationMeta {
	return &PaginationMeta{
		Total:    total,
		Page:     r.Page,
		PageSize: r.PageSize,
	}
}

func (r *PaginationRequest) ToSqlPagination() (offset, limit int) {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 {
		r.PageSize = 20
	}

	return (r.Page - 1) * r.PageSize, r.PageSize
}

type PaginationMeta struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PaginationWrapper[T any] struct {
	List []T             `json:"list"`
	Meta *PaginationMeta `json:"meta"`
}
