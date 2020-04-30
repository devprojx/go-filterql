# filterql
Converts query string into partial sql code. 

# Install
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
