package handler

import (
	"net/http"
	"ngc/model"
	"ngc/repo"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var p model.Product

	if err := ctx.ShouldBindJSON(&p); err != nil {
		handleError(repo.ErrBindJSON, ctx)
		return
	}

	// validate product
	if p.Id == 1 {
		handleError(repo.ErrParam, ctx)
		return
	}

	newProducts, err := h.Repo.CreateProduct(p)
	if err != nil {
		handleError(err, ctx)
		return
	}
	ctx.JSON(http.StatusCreated, newProducts)
}
