package paginator

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewPaginator(t *testing.T) {
	t.Run("returns error when items is not a slice", func(t *testing.T) {
		_, err := NewPaginator("not a slice")
		if err == nil {
			t.Errorf("expected an error, but got nil")
		}
		if !errors.Is(err, errInvalidItems) {
			t.Errorf("expected %v, but got %v", errInvalidItems, err)
		}
	})

	t.Run("returns a Paginator instance when items is a slice", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		p, err := NewPaginator(items)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		res, _ := p.Paginate(1)
		if !reflect.DeepEqual(res.([]int), items) {
			t.Errorf("expected %v, but got %v", items, res.([]int))
		}
		if p.pageSize != 30 {
			t.Errorf("expected default page size of 30, but got %d", p.pageSize)
		}
	})
}

func TestPaginator_SetPageSize(t *testing.T) {
	p := &Paginator{}

	p.SetPageSize(10)
	if p.pageSize != 10 {
		t.Errorf("expected page size of 10, but got %d", p.pageSize)
	}

	p.SetPageSize(20)
	if p.pageSize != 20 {
		t.Errorf("expected page size of 20, but got %d", p.pageSize)
	}
}

func TestPaginator_Paginate(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	p, _ := NewPaginator(items)

	t.Run("returns error when page is less than 1", func(t *testing.T) {
		_, err := p.Paginate(0)
		if err == nil {
			t.Errorf("expected an error, but got nil")
		}
		if !errors.Is(err, errInvalidPage) {
			t.Errorf("expected %v, but got %v", errInvalidPage, err)
		}
	})

	t.Run("returns error when page is greater than total pages", func(t *testing.T) {
		_, err := p.Paginate(2)
		if err == nil {
			t.Errorf("expected an error, but got nil")
		}
		if !errors.Is(err, errInvalidPage) {
			t.Errorf("expected %v, but got %v", errInvalidPage, err)
		}
	})

	t.Run("returns a slice of items for the given page", func(t *testing.T) {
		p.SetPageSize(3)
		page1, _ := p.Paginate(1)
		if !reflect.DeepEqual(page1, []int{1, 2, 3}) {
			t.Errorf("expected %v, but got %v", []int{1, 2, 3}, page1)
		}

		page2, _ := p.Paginate(2)
		if !reflect.DeepEqual(page2, []int{4, 5}) {
			t.Errorf("expected %v, but got %v", []int{4, 5}, page2)
		}
	})
}

func TestPaginator_PaginateWithDetails(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	p, _ := NewPaginator(items)

	t.Run("returns a PageWithDetails struct for the given page", func(t *testing.T) {
		p.SetPageSize(3)

		tests := []struct {
			pageNum int
			want    *PageWithDetails
		}{
			{
				pageNum: 1,
				want: &PageWithDetails{
					Result:      []int{1, 2, 3},
					CurrentPage: 1,
					TotalPages:  2,
					TotalCount:  5,
					PageSize:    3,
					HasNextPage: true,
					HasPrevPage: false,
				},
			},
			{
				pageNum: 2,
				want: &PageWithDetails{
					Result:      []int{4, 5},
					CurrentPage: 2,
					TotalPages:  2,
					TotalCount:  5,
					PageSize:    3,
					HasNextPage: false,
					HasPrevPage: true,
				},
			},
		}

		for _, tt := range tests {
			t.Run("", func(t *testing.T) {
				got, _ := p.PaginateWithDetails(tt.pageNum)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("expected %v, but got %v", tt.want, got)
				}
			})
		}
	})
}
