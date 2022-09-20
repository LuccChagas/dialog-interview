package router

import (
	"dialog-interview/src/errs"
	"dialog-interview/src/services"
	"net/http"

	_ "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

// @Summary      Show all authors or filter by Page
// @Description  get models.AuthorsResultSet
// @Tags         Authors
// @Accept       json
// @Produce      json
// @Param        page   path      int  true  "Pagination item"
// @Success      200  {object}  models.AuthorsResultSet
// @Failure      400  {object}  errs.Handling
// @Router       /authors/:page [get]
func GetAuthors(c echo.Context) error {
	page := c.Param("page")

	authors, err := services.GetAuthorsService(page)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, authors)
}

// @Summary      Get authors by Name
// @Description  get []models.Authors
// @Tags         Authors
// @Accept       json
// @Produce      json
// @Param        name   path      string  false  "Author Name"
// @Success      200  {object}  models.Authors
// @Failure      400  {object}  errs.Handling
// @Router       /authors/:name [get]
func GetAuthorsByName(c echo.Context) error {
	name := c.QueryParam("name")
	if len(name) <= 0 {
		return c.String(http.StatusBadRequest, "Params - `name` is required")
	}

	authors, err := services.GetAuthorsByNameService(name)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, authors)
}
