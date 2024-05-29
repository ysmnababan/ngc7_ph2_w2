package handler

import (
	"net/http"
	"ngc/model"
	"ngc/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	param_id := c.Param("id")
	id, err := strconv.Atoi(param_id)
	if err != nil || id <= 0 {
		handleError(repo.ErrInvalidId, c)
		return
	}

	var p model.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		handleError(repo.ErrBindJSON, c)
		return
	}

	// validate
	if p.Id == 1 {
		handleError(repo.ErrParam, c)
		return
	}

	err = h.Repo.UpdateProduct(id, p)
	if err != nil {
		handleError(err, c)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "data updated",
			"product": p,
		},
	)
}
