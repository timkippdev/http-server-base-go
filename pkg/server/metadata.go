package server

type Metadata struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func NewMetadata(total int, params *PaginationParams) *Metadata {
	return &Metadata{
		Limit:  params.Limit,
		Offset: params.Offset,
		Total:  total,
	}
}
