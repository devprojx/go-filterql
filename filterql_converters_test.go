package filterql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertQueryStrToSql(t *testing.T) {
	testTable := []struct {
		testName  string
		queryStr  string
		filters   map[string]string
		expQuery  string
		expParams []interface{}
	}{
		{
			testName: "Only one filter with no logical or comparison operator",
			queryStr: "name=yanik%20blake",
			filters: map[string]string{
				"name": "string",
			},
			expQuery: "name = ?",
			expParams: []interface{}{
				"yanik blake",
			},
		},
		{
			testName: "Only one filter with no logical or comparison operator",
			queryStr: "first_name=yanik",
			filters: map[string]string{
				"first_name": "string",
			},
			expQuery: "first_name = ?",
			expParams: []interface{}{
				"yanik",
			},
		},
		{
			testName: "Multiple Filters",
			queryStr: "first_name=yanik:eq:and&last_name=black:eq&height=172.5",
			filters: map[string]string{
				"first_name": "string",
				"last_name":  "string",
				"height":     "float",
			},
			expQuery: "first_name = ? AND last_name = ? AND height = ?",
			expParams: []interface{}{
				"yanik",
				"black",
				float64(172.5),
			},
		},
		{
			testName: "Duplicate filters",
			queryStr: "age=18:gte:and&age=28:lt",
			filters: map[string]string{
				"age": "int",
			},
			expQuery: "age >= ? AND age < ?",
			expParams: []interface{}{
				int64(18),
				int64(28),
			},
		},
		{
			testName: "Without Value",
			queryStr: "age",
			filters: map[string]string{
				"age": "int",
			},
			expQuery: "age = ?",
			expParams: []interface{}{
				int64(0),
			},
		},
	}
	//run tests
	for _, data := range testTable {
		result, params := ConvertQueryStrToSql(data.queryStr, data.filters)

		if !cmp.Equal(data.expQuery, result) || !cmp.Equal(data.expParams, params) {
			t.Errorf("\n[Query String] input -> %s \ngot -> %v \nexp -> %v", data.queryStr, result, data.expQuery)
			t.Errorf("\n[Parameter ] \ngot -> %v \nexp -> %v", params, data.expParams)
		}
	}

}
