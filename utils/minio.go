package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var MinioBucket string

func InitMinio() {
	endpoint := os.Getenv("MINIO_BASE_URL")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucket := os.Getenv("MINIO_BUCKET")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))

	if endpoint == "" || accessKey == "" || secretKey == "" || bucket == "" {
		log.Fatal("Gagal konfigurasi Minio")
	}

	MinioBucket = bucket

	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Gagal konek ke MinIO: %v", err)
	}
}

// utils/minio.go
func UploadToMinio(file *multipart.FileHeader, prefix string) (string, error) {
	bucket := os.Getenv("MINIO_BUCKET")
	tanggal := time.Now().Format("02-01-2006")
	objectName := fmt.Sprintf("%s/%s_%d", tanggal, prefix, time.Now().UnixNano())

	// buka file multipart
	fileReader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileReader.Close()

	// upload langsung stream ke MinIO
	_, err = MinioClient.PutObject(context.Background(), bucket, objectName, fileReader, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	// kembalikan nama objek atau URL file di MinIO
	return objectName, nil
}

func GetPublicURL(objectName string, expiry time.Duration) (string, error) {
	if MinioClient == nil {
		return "", fmt.Errorf("Minio client belum diinisialisasi")
	}
	presignedURL, err := MinioClient.PresignedGetObject(
		context.Background(),
		MinioBucket,
		objectName,
		expiry,
		nil,
	)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
