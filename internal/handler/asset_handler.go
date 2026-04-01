package handler

import (
	"go-project/internal/dto"
	apiErrors "go-project/internal/handler/errors"
	"go-project/internal/handler/utils"
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
	userID := utils.ExtractUserID(c)

	var dto dto.CreateAssetDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	asset, err := h.assetService.CreateAsset(userID, dto)
	if err != nil {
		apiErrors.HandleError(c, err)

		return
	}

	c.JSON(http.StatusCreated, mappers.ToAssetResponse(*asset))
}

func (h *AssetHandler) GetAllAssets(c *gin.Context) {
	assets, err := h.assetService.GetAllAssets()

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	var assetResponses []dto.AssetResponse

	for _, asset := range assets {
		assetResponses = append(assetResponses, mappers.ToAssetResponse(asset))
	}

	c.JSON(http.StatusOK, assetResponses)
}
