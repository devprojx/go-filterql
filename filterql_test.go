package filterql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuerStringParser(t *testing.T) {

	testTable := []struct {
		testName string
		queryStr string
		filters  map[string]string
		exp      []FilteredResult
	}{
		{
			testName: "Only one filter with no logical or comparison operator",
			queryStr: "name=yanik%20blake",
			filters: map[string]string{
				"name": "string",
			},
			exp: []FilteredResult{
				{
					Field:     "name",
					Type:      "string",
					Value:     "yanik blake",
					Operator:  " = ",
					Condition: " AND ",
				},
			},
		},
		{
			testName: "Only one filter with no logical or comparison operator",
			queryStr: "first_name=yanik",
			filters: map[string]string{
				"first_name": "string",
				"last_name":  "string",
				"age":        "int",
			},
			exp: []FilteredResult{
				{
					Field:     "first_name",
					Type:      "string",
					Value:     "yanik",
					Operator:  " = ",
					Condition: " AND ",
				},
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
			exp: []FilteredResult{
				{
					Field:     "first_name",
					Type:      "string",
					Value:     "yanik",
					Operator:  " = ",
					Condition: " AND ",
				},
				{
					Field:     "last_name",
					Type:      "string",
					Value:     "black",
					Operator:  " = ",
					Condition: " AND ",
				},
				{
					Field:     "height",
					Type:      "float",
					Value:     float64(172.5),
					Operator:  " = ",
					Condition: " AND ",
				},
			},
		},
		{
			testName: "Duplicate filters",
			queryStr: "age=18:gte:and&age=28:lt",
			filters: map[string]string{
				"first_name": "string",
				"last_name":  "string",
				"age":        "int",
			},
			exp: []FilteredResult{
				{
					Field:     "age",
					Type:      "int",
					Value:     int64(18),
					Operator:  " >= ",
					Condition: " AND ",
				},
				{
					Field:     "age",
					Type:      "int",
					Value:     int64(28),
					Operator:  " < ",
					Condition: " AND ",
				},
			},
		},
		{
			testName: "Without Value",
			queryStr: "age",
			filters: map[string]string{
				"age": "int",
			},
			exp: []FilteredResult{
				{
					Field:     "age",
					Type:      "int",
					Value:     int64(0),
					Operator:  " = ",
					Condition: " AND ",
				},
			},
		},
	}

	//run tests
	for _, data := range testTable {
		result := QueryStringParser(data.queryStr, data.filters)

		if !cmp.Equal(data.exp, result) {
			t.Errorf("\ninput -> %s \ngot -> %v \nexp -> %v", data.queryStr, result, data.exp)
		}
	}

}
