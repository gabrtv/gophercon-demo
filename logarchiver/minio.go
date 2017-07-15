package main

import (
	"fmt"
	"io"
	"log"

	"time"

	"github.com/minio/minio-go"
)

func newMinioClient(cfg *minioConfig) (*minio.Client, error) {

	minioClient, err := minio.New(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey, cfg.SSL)
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}

func createBucket(cfg *minioConfig, c *minio.Client) error {

	err := c.MakeBucket(cfg.BucketName, cfg.Location)
	if err != nil {
		exists, err := c.BucketExists(cfg.BucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", cfg.BucketName)
		} else {
			return err
		}
	}
	log.Printf("Successfully created %s\n", cfg.BucketName)
	return nil
}

func uploadFile(cfg *minioConfig, c *minio.Client, r io.Reader, t time.Time) error {

	objectName := fmt.Sprintf("logs-%v.out", t.Format(time.RFC3339))

	n, err := c.PutObject(cfg.BucketName, objectName, r, "application/octet-stream")
	if err != nil {
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	return nil
}
