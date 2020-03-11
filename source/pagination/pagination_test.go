package pagination

import (
	"net/http/httptest"
	"testing"
)

func TestHappyPath(t *testing.T) {
	pagination, err := New(90, 30)

	if err != nil {
		t.Errorf("error on pagination construction, theres no error %s", err)
		return
	}

	query := "select * from xablau where id = $1"

	paginatedQuery := pagination.PaginateQuery(query)

	if paginatedQuery != "select * from xablau where id = $1 LIMIT 30 OFFSET 90" {
		t.Errorf("error on generated paginated query %s", paginatedQuery)
		return
	}
}

func TestErrorsPaths(t *testing.T) {
	_, err := New(90, 0)

	if err == nil {
		t.Errorf("error on pagination construction, theres an error %s", err)
		return
	}

	if err.Error() != "Invalid 'size' = '0' parameter received. Required natural number, greater than zero" {
		t.Errorf("error on pagination error message about \"Size\" parameter: %v", err)
		return
	}
}

func TestParseFromRequest(t *testing.T) {
	request := httptest.NewRequest("GET", "/xablau/fadsfads", nil)
	_, _, err := ParsePaginationFromRequest(request)

	if err.Error() != "Invalid empty 'from' parameter received. Required natural number; Invalid empty 'size' parameter received. Required natural number, greater than zero; Invalid 'size' = '0' parameter received. Required natural number, greater than zero" {
		t.Errorf("error on request parser, theres an error on message %s", err)
		return
	}

	request = httptest.NewRequest("GET", "/xablau?from=0&size=1", nil)
	from, size, err := ParsePaginationFromRequest(request)

	if err != nil {
		t.Errorf("error on request parser, theres an error on message %s", err)
		return
	}

	if from != 0 {
		t.Errorf("error on request parser, error on 'from' parameter %d", from)
		return
	}

	if size != 1 {
		t.Errorf("error on request parser, error on 'size' parameter %d", size)
		return
	}

}

func TestPaginationParser(t *testing.T) {
	testCases := []struct {
		name          string
		from          int
		size          int
		expectedError bool
	}{
		{"Valid from number", 1, 1, false},
		{"Valid size number", 0, 1, false},
		{"Invalid from natural number - negative", -42, 1, true},
		{"Invalid size natural positive number - zero", 1, 0, true},
		{"Invalid size natural positive number - negative", 1, -42, true},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		err := validatePaginationData(tc.from, tc.size)

		if (tc.expectedError && err == nil) || (!tc.expectedError && err != nil) {
			t.Errorf("Error on PaginationParser: Behavior not expected in '%s'", tc.name)
			return
		}
	}
}
