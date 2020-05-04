
#go-filterql
Converts query string into partial sql code using a RHS colon syntax.
Given a query string "``` ?first_name=John:eq:and&last_name:eq:Doe```"
it will generate "```first_name = ? AND last_name = ?```".

### RHS Colon Sytanx

Provided the query string "```?name=john doe:eq:and```", the following components can be extracted:


- **Field/Column name:** ```name```
- **Filter/Search value:** ```john doe```
- **Comparison Operator:** ```:eq```
- **Logical Operator:** ```:and```

**Note** that each component of the RHS syntax following the field name is separated by a colon. The **Logical Operator** is optional, however this will be default to an "AND" operator if more than one filter is provided.

### Supported SQL Operators

SQL Comparision Operators   | RHS COLON SYNTAX 
----------------------------|------------------
=                           | eq               
<=                          | lte              
&gt;=                       | gte              
<                           | lt               
&gt;                        | gt               
<&gt;                       | nt           
LIKE                        | lk          


SQL Logical Operators | RHS COLON SYNTAX 
----------------------|------------------
AND                   | and               
OR                    | or     


## Installation
```
go get github.com/devprojx/go-filterql
```
         
   
## Usage

Example using the [Echo Framework](https://echo.labstack.com/guide) framework

```go
import (
  "net/http"
  "github.com/labstack/echo/v4"

  "github.com/devprojx/go-filterql"
  
  ...
)

func (h *Handler) FindAllCompanies(ctx echo.Context) error {
  //Get query string
  queryString := ctx.QueryString()
  
  //Map defining the column names and types that are 
  //filterable on your endpoint
  possibleFilters := map[string]string{
    "first_name": "string",
    "last_name": "string",
    "age": "int",
  }
  
  //Converts query string to partial parameterized SQL string along with parameters
  filterQuery, params := filterql.ConvertQueryStrToSql(queryString, possibleFilters)

  users := []*models.User{}
  h.DB.Where(filterQuery, params).Find(&users)

  return ctx.JSON(http.StatusOK, users)
}
```
