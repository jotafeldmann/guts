package pagination

import (
	"erroraggregator"
	"fmt"
	"net/http"
	"strconv"
)

const fromParameterRequirements = "Required natural number"
const sizeParameterRequirements = "Required natural number, greater than zero"

// Pagination component
type Pagination struct {
	from    int
	size    int
	results interface{}
	total   int
}

// HTTPResponse of pagination
type HTTPResponse struct {
	Total   int         `json:"total"`
	Results interface{} `json:"results"`
}

// GetHTTPResponse returns pagination HTTPResponse
func (p *Pagination) GetHTTPResponse() HTTPResponse {
	return HTTPResponse{
		Results: p.results,
		Total:   p.total,
	}
}

// New creates a pagination complex struct instance
func New(from, size int) (*Pagination, error) {
	err := validatePaginationData(from, size)

	if err != nil {
		return nil, err
	}

	p := new(Pagination)
	p.from = from
	p.size = size
	return p, nil
}

// ParsePaginationFromRequest parse pagination data from request URL query
func ParsePaginationFromRequest(r *http.Request) (from, size int, err error) {
	errors := erroraggregator.New()

	fromParameter := r.URL.Query().Get("from")
	sizeParameter := r.URL.Query().Get("size")

	if fromParameter == "" {
		errors.Append(fmt.Errorf("Invalid empty 'from' parameter received. %s", fromParameterRequirements))
	} else {
		from, err = strconv.Atoi(fromParameter)

		if err != nil {
			errors.Append(fmt.Errorf("Invalid 'from' = '%v' parameter received. %s", from, fromParameterRequirements))
		}
	}

	if sizeParameter == "" {
		errors.Append(fmt.Errorf("Invalid empty 'size' parameter received. %s", sizeParameterRequirements))
	} else {
		size, err = strconv.Atoi(sizeParameter)

		if err != nil {
			errors.Append(fmt.Errorf("Invalid 'size' = '%v' parameter received. %s", size, sizeParameterRequirements))
		}
	}

	err = validatePaginationData(from, size)

	if err != nil {
		errors.Append(err)
	}

	if errors.GotErrors() {
		return 0, 0, errors.GetErrorMessages()
	}

	return from, size, err
}

// PaginateQuery return a query with paginate string
func (p *Pagination) PaginateQuery(sqlQuery string) string {
	queryLimit := fmt.Sprintf("LIMIT %d", p.size)
	sqlQuery = fmt.Sprintf(sqlQuery+" %s", queryLimit)

	queryOffset := fmt.Sprintf("OFFSET %d", p.from)
	sqlQuery = fmt.Sprintf(sqlQuery+" %s", queryOffset)

	return sqlQuery
}

// SetResults set the page with the items
func (p *Pagination) SetResults(results interface{}) {
	p.results = results
}

// SetTotalItems set the total items of all pages
func (p *Pagination) SetTotalItems(total int) {
	p.total = total
}

func validatePaginationData(from, size int) error {

	errors := erroraggregator.New()

	if from < 0 {
		errors.Append(fmt.Errorf("Invalid 'from' = '%d' parameter received. %s", from, fromParameterRequirements))
	}

	if size < 1 {
		errors.Append(fmt.Errorf("Invalid 'size' = '%d' parameter received. %s", size, sizeParameterRequirements))
	}

	return errors.GetErrorMessages()
}
