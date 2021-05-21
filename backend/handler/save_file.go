package handler

import (
	"context"
	"fmt"
	"io"
	"os"
	"log"
	"mime/multipart"
	"net/http"
	"time"
	"github.com/Oasixer/ShopifyChallenge2021/config"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO this is entirely duplicate code w/ get_file, need to add auth to this endpoint and refactor out the duplicate code

const (
	projectID  = "triple-skein-312919" // TODO move these to config
	bucketName = "shopify_imghub"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var uploader *ClientUploader

func init() {
	log.Printf("init handler %s", config.CONFIG.Mode)
	if config.CONFIG.Mode == "test" || config.CONFIG.Mode == "" {
		return
	}
	if config.CONFIG.Mode == "dev" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/k/Downloads/triple-skein-312919-cee9b170f894.json") // temp
	}
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "imageFiles/",
	}
}

func SaveFile(c *gin.Context) {
	f, err := c.FormFile("file_input")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	blobFile, err := f.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileUUID, _ := uuid.NewUUID()
	err = uploader.UploadFile(blobFile, fileUUID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Print("uploaded file succesfully")
	c.JSON(201, gin.H{
		"message": "success",
		"uuid":    fileUUID.String(),
	})
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
