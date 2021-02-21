package error

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

//HiveError represent the server themisError format
type HiveError struct {
	Path   string `json:"path"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

/*NewHiveBadRequestError - create a new HiveError for a 400 themisError code
//path is the endpoint where the themisError happened
//detail is the themisError description
*/
func NewHiveBadRequestError(ctx iris.Context, detail string) *HiveError {
	ctx.StatusCode(iris.StatusBadRequest)
	err := &HiveError{
		Path:   ctx.RouteName(),
		Status: http.StatusBadRequest,
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: detail,
	}
	return err
}

/*NewHiveUnauthorizedError - create a new HiveError for a 401 themisError code
//path is the endpoint where the themisError happened
//detail is the themisError description
*/
func NewHiveUnauthorizedError(ctx iris.Context, detail string) *HiveError {
	ctx.StatusCode(iris.StatusUnauthorized)
	return &HiveError{
		Path:   ctx.RouteName(),
		Status: http.StatusUnauthorized,
		Title:  http.StatusText(http.StatusUnauthorized),
		Detail: detail,
	}
}

/*NewHiveInternalServerError - create a new HiveError for a 500 themisError code
//path is the endpoint where the themisError happened
//detail is the themisError description
*/
func NewHiveInternalServerError(ctx iris.Context, detail string) *HiveError {
	ctx.StatusCode(iris.StatusInternalServerError)
	return &HiveError{
		Path:   ctx.RouteName(),
		Status: iris.StatusInternalServerError,
		Title:  http.StatusText(iris.StatusInternalServerError),
		Detail: detail,
	}
}