# GO-PAGINATE
    
[![GoDoc](https://godoc.org/github.com/deestarks/go-paginate?status.svg)](https://godoc.org/github.com/deestarks/go-paginate)
[![Go Report Card](https://goreportcard.com/badge/github.com/deestarks/go-paginate)](https://goreportcard.com/report/github.com/deestarks/go-paginate)

This package provides a simple implementation of pagination for any slice of data.

## Usage
Import package:
```go
import "github.com/deestarks/go-paginate"
```

Then, create a new `Paginator` for your slice of data:
```go
p, err := paginator.NewPaginator(yourSliceOfData)
if err != nil {
    // handle error
}
```

You can then set the page size for the paginator (default is 30):
```go
p.SetPageSize(50)
```

To paginate your data, call the `Paginate` method with the desired page number, which returns a slice of the data for the given page:
```go
page := 1
result, err := p.Paginate(page)
if err != nil {
    // handle error
}
// use paginated result
```

Alternatively, you can use `PaginateWithDetails` to get a `PageWithDetails` struct containing details of the pagination, along with the paginated result:
```go
page := 1
pageWithDetails, err := p.PaginateWithDetails(page)
if err != nil {
    // handle error
}
// use paginated result and details
```

## Example
```go
package main

import (
    "fmt"
    "github.com/deestarks/go-paginate"
)

type User struct {
    ID   int
    Name string
}

func main() {
    users := []User{
        {ID: 1, Name: "John"},
        {ID: 2, Name: "Jane"},
        {ID: 3, Name: "Bob"},
        {ID: 4, Name: "Alice"},
        {ID: 5, Name: "Tom"},
        {ID: 6, Name: "Jerry"},
        {ID: 7, Name: "Mickey"},
        {ID: 8, Name: "Minnie"},
        {ID: 9, Name: "Donald"},
        {ID: 10, Name: "Daisy"},
    }

    // create a new paginator for the users slice
    p, err := paginator.NewPaginator(users)
    if err != nil {
        panic(err)
    }

    // set the page size to 5 (default is 30)
    p.SetPageSize(5)

    // get the first page
    page := 1

    // paginate the users
    pageWithoutDetails, err := p.Paginate(page)
    if err != nil {
        panic(err)
    }

    pageWithDetails, err := p.PaginateWithDetails(page)
    if err != nil {
        panic(err)
    }

    // print the paginated results
    fmt.Println(pageWithoutDetails.([]User))
    // Output:
    // [{1 John} {2 Jane} {3 Bob} {4 Alice} {5 Tom}]

    fmt.Println("Current Page:", pageWithDetails.CurrentPage)
    fmt.Println("Total pages:", pageWithDetails.TotalPages)
    fmt.Println("Total items:", pageWithDetails.TotalCount)
    fmt.Println("Items per page:", pageWithDetails.PageSize)
    fmt.Println("Has next page:", pageWithDetails.HasNextPage)
    fmt.Println("Has previous page:", pageWithDetails.HasPrevPage)
    fmt.Println("Results:", pageWithDetails.Results.([]User))
    // Output:
    // Current Page: 1
    // Total pages: 2
    // Total items: 10
    // Items per page: 5
    // Has next page: true
    // Has previous page: false
    // Results: [{1 John} {2 Jane} {3 Bob} {4 Alice} {5 Tom}]

    // JSON output of pageWithDetails:
    // {
    //     "current_page": 1,
    //     "total_pages": 2,
    //     "total_count": 10,
    //     "page_size": 5,
    //     "has_next_page": true,
    //     "has_prev_page": false,
    //     "results": [
    //         {
    //             "ID": 1,
    //             "Name": "John"
    //         },
    //         ...
    //     ]
    // }
}
```
