package upload

import (
	"context"
	"go-docker/pkg/constants"
	"go-docker/pkg/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"go-docker/internal/db"
	"go-docker/models"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

type UploadService struct{}

func (s *UploadService) validateFile(file *multipart.FileHeader) error {
	if file.Size > constants.MaxFileSize {
		return utils.NewCustomError(http.StatusBadRequest, "ファイルの上限サイズより大きいです。")
	}

	filename := file.Filename
	ext := strings.ToLower(filename[strings.LastIndex(filename, "."):])
	for _, allowedExt := range constants.AllowedExtensions {
		if ext == allowedExt {
			return nil
		}
	}

	return utils.NewCustomError(http.StatusBadRequest, "ファイルの拡張子が不正です。")
}

func (s *UploadService) uploadToImageKit(ik *imagekit.ImageKit, folder *string, filename *string, file *multipart.File) (*UploadToImageKitResponse, error) {
	ctx := context.Background()
	resp, err := ik.Uploader.Upload(ctx, *file, uploader.UploadParam{
		FileName:          *filename,
		Folder:            *folder,
		UseUniqueFileName: utils.BoolPtr(true),
	})
	if err != nil {
		log.Printf("画像アップロードエラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "画像のアップロードに失敗しました。")
	}

	return &UploadToImageKitResponse{Url: resp.Data.Url, FileId: resp.Data.FileId}, nil
}

func (s *UploadService) validateUploadImages(c *gin.Context) ([]*multipart.FileHeader, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "multipart formのパースに失敗")
	}

	files := form.File["images"]
	if len(files) == 0 {
		return nil, utils.NewCustomError(http.StatusBadRequest, "ファイルが選択されていません。")
	} else if len(files) > 10 {
		return nil, utils.NewCustomError(http.StatusBadRequest, "ファイルの上限選択数を超えています。")
	}

	return files, nil
}

func (s *UploadService) UploadImagesService(ik *imagekit.ImageKit, folder *string, files []*multipart.FileHeader) (*[]string, error) {
	var imageUrls []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "ファイルの開封に失敗")
		}
		defer src.Close()

		if err := s.validateFile(file); err != nil {
			return nil, err
		}
		image, err := uploadService.uploadToImageKit(ik, folder, &file.Filename, &src)
		if err != nil {
			return nil, err
		}
		tempImage := models.TempImage{
			FileId:    image.FileId,
			Image:     image.Url,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}
		if err := db.DB.Create(&tempImage).Error; err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "画像情報の保存に失敗しました")
		}

		imageUrls = append(imageUrls, image.Url)
	}

	return &imageUrls, nil
}

func NewUploadService() *UploadService {
	return &UploadService{}
}
