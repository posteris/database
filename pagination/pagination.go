package pagination

//Pagination data structure to helps GORM pagination implementation
type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

//GetOffset function to return the pagination offset
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

//GetLimit get the pagination limit, by default it is 10
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

//GetPage function to return the page number
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}

	return p.Page
}
