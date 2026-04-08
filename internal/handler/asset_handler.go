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

func (h *AssetHandler) GetAssetsForUser(c *gin.Context) {
	userID := utils.ExtractUserID(c)
	assets, err := h.assetService.GetAssetsForUser(userID)

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

func (h *AssetHandler) GetAssetById(c *gin.Context) {
	userID := utils.ExtractUserID(c)
	ID, err := utils.ExtractIDParam(c)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	asset, err := h.assetService.GetAssetById(userID, ID)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, mappers.ToAssetResponse(*asset))
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

func (h *AssetHandler) UpdateAsset(c *gin.Context) {
	var dto *dto.UpdateAssetDTO
	userID := utils.ExtractUserID(c)
	ID, err := utils.ExtractIDParam(c)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	updatedAsset, err := h.assetService.UpdateAsset(userID, ID, *dto)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, mappers.ToAssetResponse(*updatedAsset))
}

func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	userID := utils.ExtractUserID(c)
	ID, err := utils.ExtractIDParam(c)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	err = h.assetService.DeleteAsset(userID, ID)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ADMIN FUNCTIONS
func (h *AssetHandler) GetAnyAssetById(c *gin.Context) {
	ID, err := utils.ExtractIDParam(c)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	asset, err := h.assetService.GetAnyAssetById(ID)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, mappers.ToAssetResponse(*asset))
}

func (h *AssetHandler) UpdateAnyAsset(c *gin.Context) {
	var dto *dto.UpdateAssetDTO
	ID, err := utils.ExtractIDParam(c)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}
	err = c.ShouldBindJSON(&dto)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}
	updatedAsset, err := h.assetService.UpdateAnyAsset(ID, *dto)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, mappers.ToAssetResponse(*updatedAsset))
}
