package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"time"
)

type Minio struct {
	storage *minio.Client
}

func InitMinio() (Minio, error) {
	endpoint := "localhost:9000"
	accessKeyId := "minio_admin"
	secretAccessKey := "minioAdmin"

	useSSL := false

	//Initialize minIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println("Init:", err.Error())
		return Minio{}, err
	}
	ctx := context.Background()
	status, err := minioClient.BucketExists(ctx, "videos")
	if err != nil {
		return Minio{}, fmt.Errorf("failed to check minio bucket. Error %v", err)
	}
	if status != true {
		err = minioClient.MakeBucket(ctx, "videos", minio.MakeBucketOptions{})
	}
	return Minio{
		storage: minioClient,
	}, nil
}

func (s *Minio) UploadImage(path string, response *fasthttp.Response) error {
	_, err := s.storage.PutObject(context.Background(), "videos", path, bytes.NewReader(response.Body()),
		int64(response.Header.ContentLength()), minio.PutObjectOptions{
			ContentType: string(response.Header.ContentType()),
		})
	if err != nil {
		log.Println("error while uploading img to minio storage. Err: " + err.Error())
		return err
	}
	return nil
}

func (s *Minio) UploadConvertedImage(path string, reader io.Reader) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	_, err := s.storage.PutObject(context.Background(), "videos", path, reader, int64(buf.Len()), minio.PutObjectOptions{
		ContentType: "image/webp",
	})
	if err != nil {
		log.Println("error while uploading converted image", err.Error())
		return err
	}

	return nil
}

func (s *Minio) GenerateUploadUrl(bucket, filePath string) (string, error) {
	expire := time.Second * 24 * 60 * 60
	preSignedUrl, err := s.storage.PresignedPutObject(context.Background(), bucket, filePath, expire)
	if err != nil {
		log.Println("error while generating preSigned url. Error: ", err.Error())
	}
	return preSignedUrl.String(), nil
}
