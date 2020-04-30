package filterql

//ConvertQueryStrToSql converts URL querystring into a partial sql conditional statement
func ConvertQueryStrToSql(queryStr string, filters map[string]string) (partialSql string, parameters []interface{}) {
	var query string
	var prevlogicalOperator string

	filteredResults := QueryStringParser(queryStr, filters)

	params := make([]interface{}, 0)
	for _, filter := range filteredResults {
		query += prevlogicalOperator + filter.Field + filter.Operator + "?"
		params = append(params, filter.Value)
		prevlogicalOperator = filter.Condition
	}

	return query, params
}
