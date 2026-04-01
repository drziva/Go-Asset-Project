package handler

import (
	"go-project/internal/constants"
	"go-project/internal/dto"
	apiErrors "go-project/internal/handler/errors"
	"go-project/internal/mappers"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	assetService *service.AssetService
}

func NewAssetHandler(assetService *service.AssetService) *AssetHandler {
	return &AssetHandler{
		assetService,
	}
}

func (h *AssetHandler) CreateAsset(c *gin.Context) {
	userID, exists := c.Get(constants.UserIDKey)

	if !exists {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var dto dto.CreateAssetDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	asset, err := h.assetService.CreateAsset(userID.(uint), dto)
	if err != nil {
		apiErrors.HandleError(c, err)

		return
	}

	c.JSON(http.StatusCreated, mappers.ToAssetResponse(*asset))
}
