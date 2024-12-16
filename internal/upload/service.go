package upload

import (
	"context"
	"fmt"
	"go-docker/pkg/constants"
	"go-docker/pkg/utils"
	"log"
	"mime/multipart"
	"strings"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

type UploadService struct{}

func (s *UploadService) validateFile(file *multipart.FileHeader) error {
	if file.Size > constants.MaxFileSize {
		return fmt.Errorf("ファイルの上限サイズより大きいです。")
	}

	filename := file.Filename
	ext := strings.ToLower(filename[strings.LastIndex(filename, "."):])
	for _, allowedExt := range constants.AllowedExtensions {
		if ext == allowedExt {
			return nil
		}
	}

	return fmt.Errorf("ファイルの拡張子が不正です。")
}

func (s *UploadService) uploadToImageKit(ik *imagekit.ImageKit, folder, filename string, file multipart.File) (string, error) {
	ctx := context.Background()
	resp, err := ik.Uploader.Upload(ctx, file, uploader.UploadParam{
		FileName: filename,
		Folder: folder,
		UseUniqueFileName: utils.BoolPtr(true),
	})
	if err != nil {
		log.Printf("画像アップロードエラー: %v", err)
		return "", fmt.Errorf("画像のアップロードに失敗しました。")
	}

	return resp.Data.Url, nil
}

func NewUploadService() *UploadService {
	return &UploadService{}
}