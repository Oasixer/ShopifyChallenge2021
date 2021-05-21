package handler

import (
	"context"
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	// "mime/multipart"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Oasixer/ShopifyChallenge2021/config"
	"github.com/gin-gonic/gin"
)

// TODO this is entirely duplicate code w/ save_file, need to add auth to this endpoint and refactor out the duplicate code

type ClientDownloader struct {
	cl           *storage.Client
	projectID    string
	bucketName   string
	downloadPath string
	w io.Writer
}

var downloader *ClientDownloader

func init() {
	// TODO update this logic to be more robust
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

	buf := &bytes.Buffer{}

	downloader = &ClientDownloader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		downloadPath: "imageFiles/",
		w: buf,
	}
}

func GetFile(c *gin.Context) {
	log.Print("Downloading file")
	uuid := c.Query("uuid")
	fileExt := c.Query("fileExt")

	bytes_, err := downloader.GetFile(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Print("downloaded file succesfully")

	log.Printf("type: %#v", fileExt)
	// http.ServeContent(w, r, fmt.SPrintf("%s.%s", uuid, contentType), modtime, bytes_.NewReader(pdfData))

	c.Data(http.StatusOK, fmt.Sprintf("application/%s",fileExt), bytes_)
}

// GetFile uploads an object
func (c *ClientDownloader) GetFile(uuid string) ([]byte, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := c.cl.Bucket(c.bucketName).Object(c.downloadPath + uuid).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("readFile: unable to open file from bucket %q", bucketName)
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("readFile: unable to read data from bucket %q, file %q: %v", bucketName, uuid, err)
	}

	fmt.Fprintf(c.w, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
	if len(slurp) > 1024 {
		fmt.Fprintf(c.w, "...%s\n", slurp[len(slurp)-1024:])
	} else {
		fmt.Fprintf(c.w, "%s\n", slurp)
	}

	return slurp, nil
}
