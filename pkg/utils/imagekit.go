package utils

import (
	"log"
	"os"

	"github.com/imagekit-developer/imagekit-go"
)

func NewImageKit() *imagekit.ImageKit {
	publicKey := os.Getenv("IMAGE_KIT_PUBLIC_KEY")
	privateKey := os.Getenv("IMAGE_KIT_PRIVATE_KEY")
	urlEndPoint := os.Getenv("IMAGE_KIT_URL_ENDPOINT")
	log.Printf(publicKey)
	log.Printf(privateKey)
	log.Printf(urlEndPoint)
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PublicKey: publicKey,
		PrivateKey: privateKey,
		UrlEndpoint: urlEndPoint,
	})

	if ik == nil {
		log.Printf("imageKitの初期化に失敗")
	}

	return ik
}