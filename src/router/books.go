package router

import (
	"dialog-interview/src/errs"
	"dialog-interview/src/models"
	"dialog-interview/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @Summary      Create new Books
// @Description  POST models.Book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        Book   body     models.Book  true  "Book"
// @Success      200  {object}	string
// @Failure      400  {object}  errs.Handling
// @Router       /books/create [post]
func CreateBooks(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_BIND_OBJECT}
		return c.JSON(http.StatusBadRequest, eh)
	}

	err := services.CreateBookService(book)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}
	return c.JSON(http.StatusOK, "Book was inserted with Success!")
}

// @Summary      Get Books
// @Description  GET models.Book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        name   path      string  false  "Book Name"
// @Param        edition   path      string  false  "Book Edition"
// @Param        publication_year   path      string  false  "Book Publication_Year"
// @Param        author   path      string  false  "Book Author ID"
// @Success      200  {object}	string
// @Failure      400  {object}  errs.Handling
// @Router       /books/read [get]
func ReadBooks(c echo.Context) error {
	name := c.QueryParam("name")
	edition := c.QueryParam("edition")
	publicationYear := c.QueryParam("publication_year")
	author := c.QueryParam("author")

	books, err := services.ReadBooksService(name, edition, publicationYear, author)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, books)
}

// @Summary      Update Books
// @Description  PUT models.Book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        Book   body     models.Book  true  "Book"
// @Success      204
// @Failure      400  {object}  errs.Handling
// @Router       /books/update [put]
func UpdateBooks(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_BIND_OBJECT}
		return c.JSON(http.StatusBadRequest, eh)
	}

	if book.ID < 1 {
		return c.JSON(http.StatusBadRequest, "ID is required")
	}

	err := services.UpdateBooksService(book)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}
	return c.NoContent(http.StatusNoContent)
}

// @Summary      Delete Books
// @Description  DELETE models.Book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Success      204
// @Failure      400  {object}  errs.Handling
// @Router       /books/update [delete]
func DeleteBooks(c echo.Context) error {
	ID := c.Param("id")
	if len(ID) <= 0 {
		return c.JSON(http.StatusBadRequest, "ID is required")
	}

	err := services.DeleteBookService(ID)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.NoContent(http.StatusNoContent)
}
