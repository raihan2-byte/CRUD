package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/data"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type inputHandler struct {
	inputService data.Service
}

func NewInputHandler(inputService data.Service) *inputHandler {
	return &inputHandler{inputService}
}

func (h *inputHandler) GetNames(c *gin.Context) {
	names, err := h.inputService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var response []data.Response

	for _, r := range names {
		namesResponse := convertToNamaResponse(r) 
		response = append(response, namesResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}

func (h *inputHandler) GetName(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	nama, err := h.inputService.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	responseNama := convertToNamaResponse(nama) 
	c.JSON(http.StatusOK, gin.H{
		"data": responseNama,
	})

}



func (h *inputHandler) PostNameHandler(c *gin.Context) {
	var namaku data.User

	err := c.ShouldBindJSON(&namaku)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("harus mengisi %s, variabel: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := h.inputService.Create(namaku)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToNamaResponse(user),
	})
}

func (h *inputHandler) UpdateName(c *gin.Context) {
	var namaku data.User

	err := c.ShouldBindJSON(&namaku)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("harus mengisi %s, variabel: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, err := h.inputService.Update(id, namaku)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToNamaResponse(user),
	})
}

func convertToNamaResponse(nama data.Nama) data.Response {
	return data.Response{
		ID:     nama.ID,
		Name:   nama.Name,
		Email:  nama.Email,
		Age:    nama.Age,
		Rating: nama.Rating,
	}
}

func (h *inputHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	nama, err := h.inputService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToNamaResponse(nama),
	})
}
