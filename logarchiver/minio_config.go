package main

import (
	"github.com/kelseyhightower/envconfig"
)

type minioConfig struct {
	Endpoint   string `envconfig:"MINIO_ENDPOINT"`
	AccessKey  string `envconfig:"MINIO_ACCESS_KEY"`
	SecretKey  string `envconfig:"MINIO_SECRET_KEY"`
	SSL        bool   `envconfig:"MINIO_SSL"`
	BucketName string `envconfig:"MINIO_BUCKET_NAME"`
	Location   string `envconfig:"MINIO_LOCATION"`
}

func parseMinioConfig() (*minioConfig, error) {
	ret := new(minioConfig)
	if err := envconfig.Process(appName, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
