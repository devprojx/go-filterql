package filterql

import (
	"net/url"
	"strconv"
	"strings"
)

var (
	comparisonOperators = map[string]string{
		"eq":  " = ",
		"lte": " =< ",
		"gte": " >= ",
		"lt":  " < ",
		"gt":  " > ",
		"nt":  " <> ",
		"lk":  " LIKE ",
	}

	logicalOperators = map[string]string{
		"and": " AND ",
		"or":  " OR ",
	}
)

//FilteredResult defines the final output of a parse querystring
type FilteredResult struct {
	Field     string
	Type      string
	Value     interface{}
	Operator  string
	Condition string
}

//QueryStringParser converts URL querystring into a slice of `FilteredResult.`
//Given the querystring `?first_name=john:eq:and&last_name=doe:eq`
func QueryStringParser(queryStr string, filters map[string]string) []FilteredResult {
	//define custom map type to allowduplicate keys
	type Map struct {
		Key   string
		Value string
	}

	params := []Map{}
	searchFilters := []FilteredResult{}

	parts := strings.Split(queryStr, "&")

	//build a key/value map of the querystring by
	//storing the query as key and the fragment as the value
	for _, part := range parts {
		split := strings.Split(part, "=")

		if len(split) > 1 && split[1] != "" {
			params = append(params, Map{
				Key:   split[0],
				Value: split[1],
			})
		} else {
			params = append(params, Map{
				Key:   split[0],
				Value: "",
			})
		}
	}

	//
	for _, param := range params {
		for name, varType := range filters {
			if param.Key == name {
				esc, _ := url.QueryUnescape(param.Value)
				parseValue, operator, condition := RHSParser(esc, varType)

				searchFilters = append(searchFilters, FilteredResult{
					Field:     param.Key,
					Type:      varType,
					Value:     parseValue,
					Operator:  operator,
					Condition: condition,
				})
				break
			}
		}
	}

	return searchFilters
}

//RHSParser separates the fragment part of the query string into three parts
//value, comparison operator (=, >, <, <>, <=, >=, LIKE) and logical operator (AND/OR).
func RHSParser(queryStrValue string, valueType string) (value interface{}, comparisonOperator string, logicOperator string) {
	var val interface{}
	var cOperator string = " = "
	var lOperator string = " AND "

	parts := strings.Split(queryStrValue, ":")
	len := len(parts)

	if valueType == "int" {
		var number int64
		number, _ = strconv.ParseInt(parts[0], 10, 64)
		val = number
	} else if valueType == "float" {
		number := 0.0
		number, _ = strconv.ParseFloat(parts[0], 64)
		val = number
	} else {
		val = parts[0]
	}

	if len == 1 {
		cOperator = comparisonOperators["eq"]
		lOperator = " AND "
		return val, cOperator, lOperator
	}

	if comparisonOperators[parts[1]] != "" {
		cOperator = comparisonOperators[parts[1]]
	}

	if len == 3 {
		if logicalOperators[parts[2]] != "" {
			lOperator = logicalOperators[parts[2]]
		}
	}

	return val, cOperator, lOperator
}
