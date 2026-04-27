package handler

import (
	"go-project/internal/dto"
	apiErrors "go-project/internal/handler/errors"
	"go-project/internal/handler/utils"
	"go-project/internal/mappers"
	"go-project/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AssetHandler struct {
	assetService *service.AssetService
}

func NewAssetHandler(assetService *service.AssetService) *AssetHandler {
	return &AssetHandler{
		assetService,
	}
}

func (h *AssetHandler) UploadFile(c *gin.Context) (*dto.ServiceCreateAssetDTO, error) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not fetch file",
		})
		return nil, err
	}

	description := c.PostForm("description")

	extension := filepath.Ext(file.Filename)

	filename := c.PostForm("name") + extension
	if filename == "" {
		filename = filepath.Base(file.Filename)
	}

	storedName := uuid.New().String() + extension

	uploadFolderDst := "../../uploads"

	err = os.MkdirAll(uploadFolderDst, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create uploads folder",
		})

		return nil, err
	}

	destination := uploadFolderDst + "/" + storedName

	err = c.SaveUploadedFile(file, destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error persisting file, please try again",
		})

		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buffer := make([]byte, 512)

	_, err = src.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &dto.ServiceCreateAssetDTO{
		FileName:    filename,
		FilePath:    destination,
		FileSize:    file.Size,
		StoredName:  storedName,
		MimeType:    http.DetectContentType(buffer),
		Description: description,
	}, nil
}

func (h *AssetHandler) CreateAsset(c *gin.Context) {
	userID := utils.ExtractUserID(c)

	serviceDTO, err := h.UploadFile(c)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	asset, err := h.assetService.CreateAsset(userID, *serviceDTO)
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

func (h *AssetHandler) DownloadAssetById(c *gin.Context) {
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

	c.Header("Content-Disposition", "attachment; filename="+asset.FileName)
	c.Header("Content-Type", asset.MimeType)

	c.File(asset.FilePath)
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

func (h *AssetHandler) DownloadAnyAssetById(c *gin.Context) {
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

	c.Header("Content-Disposition", "attachment; filename="+asset.FileName)
	c.Header("Content-Type", asset.MimeType)

	c.File(asset.FilePath)
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
