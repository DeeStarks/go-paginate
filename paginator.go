package paginator

import (
	"errors"
	"reflect"
)

var (
	errInvalidPage  = errors.New("invalid page number")
	errInvalidItems = errors.New("items must be a slice")
)

type PageWithDetails struct {
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	TotalCount  int         `json:"total_count"`
	PageSize    int         `json:"page_size"`
	HasNextPage bool        `json:"has_next_page"`
	HasPrevPage bool        `json:"has_prev_page"`
	Result      interface{} `json:"result"`
}

type Paginator struct {
	data     reflect.Value
	pageSize int
	current  int
}

// NewPaginator returns a new Paginator for the given items
// The items must be a slice of any type
func NewPaginator(items interface{}) (*Paginator, error) {
	// Check if the items are a slice
	ref := reflect.ValueOf(items)
	if ref.Kind() != reflect.Slice {
		return nil, errInvalidItems
	}

	return &Paginator{
		data:     ref,
		pageSize: 30,
	}, nil
}

// SetPageSize sets the page size for the paginator
func (p *Paginator) SetPageSize(size int) {
	p.pageSize = size
}

// Paginate returns a slice of items for the given page number
func (p *Paginator) Paginate(page int) (interface{}, error) {
	if page < 1 {
		return nil, errInvalidPage
	}

	p.current = page

	start := (p.current - 1) * p.pageSize
	end := p.current * p.pageSize

	if start > p.data.Len() {
		return nil, errInvalidPage
	}

	if end > p.data.Len() {
		end = p.data.Len()
	}

	res := reflect.MakeSlice(p.data.Type(), end-start, end-start)
	reflect.Copy(res, p.data.Slice(start, end))
	return res.Interface(), nil
}

// PaginateWithDetails returns a PageWithDetails struct for the given page number
// This struct contains details of the pagination such as the current page, total pages, total count, etc.
// along with the paginated result
func (p *Paginator) PaginateWithDetails(page int) (*PageWithDetails, error) {
	res, err := p.Paginate(page)
	if err != nil {
		return nil, err
	}

	return &PageWithDetails{
		CurrentPage: p.current,
		TotalPages:  p.getPageCount(),
		TotalCount:  p.getTotalCount(),
		PageSize:    p.pageSize,
		HasNextPage: p.hasNextPage(),
		HasPrevPage: p.hasPreviousPage(),
		Result:      res,
	}, nil
}

func (p *Paginator) getPageCount() int {
	return (p.data.Len() + p.pageSize - 1) / p.pageSize
}

func (p *Paginator) getTotalCount() int {
	return p.data.Len()
}

func (p *Paginator) hasNextPage() bool {
	return p.current < p.getPageCount()
}

func (p *Paginator) hasPreviousPage() bool {
	return p.current > 1
}
