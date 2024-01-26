package controllers

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IAuthorController interface {
	CreateAuthor(e echo.Context) error
	GetAuthor(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
}

// to access the methods of service and repo
type AuthorController struct {
	authorSvc domain.IAuthorService
}

func NewAuthorController(authorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		authorSvc: authorSvc,
	}
}

func (ac *AuthorController) CreateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	author := &models.AuthorDetail{
		AuthorName:    reqAuthor.AuthorName,
		AuthorAddress: reqAuthor.AuthorAddress,
		Phone:         reqAuthor.Phone,
	}
	fmt.Print(author)
	if err := ac.authorSvc.CreateAuthor(author); err != nil {

		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "AuthorDetail was created successfully")
}

func (ac *AuthorController) GetAuthor(e echo.Context) error {
	tempAuthorID := e.QueryParam("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	author, err := ac.authorSvc.GetAuthors(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, author)
}

func (ac *AuthorController) UpdateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	existingAuthor, err := ac.authorSvc.GetAuthors(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedAuthor := &models.AuthorDetail{
		ID:            uint(authorID),
		AuthorName:    reqAuthor.AuthorName,
		AuthorAddress: reqAuthor.AuthorAddress,
		Phone:         reqAuthor.Phone,
	}
	if updatedAuthor.AuthorName == "" {
		updatedAuthor.AuthorName = existingAuthor[0].AuthorName
	}
	if updatedAuthor.AuthorAddress == "" {
		updatedAuthor.AuthorAddress = existingAuthor[0].AuthorAddress
	}
	if updatedAuthor.Phone == "" {
		updatedAuthor.Phone = existingAuthor[0].Phone
	}
	if err := ac.authorSvc.UpdateAuthor(updatedAuthor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "AuthorDetail was updated successfully")
}

func (ac *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = ac.authorSvc.GetAuthors(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ac.authorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "AuthorDetail was deleted successfully")
}
