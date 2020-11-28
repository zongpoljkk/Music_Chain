package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

var (
	SERVICE_ACCOUNT = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
)

const (
	BUCKET_NAME = "chula-blockchain-music"
)

func UploadBytesToGCS(filename string, content []byte) (*storage.ObjectAttrs, error) {

	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	newObject := client.Bucket(BUCKET_NAME).Object(filename)
	wc := newObject.NewWriter(context.Background())
	if _, err = wc.Write(content); err != nil {
		return nil, fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("Writer.Close: %v", err)
	}
	return newObject.Attrs(context.Background())
}

func GetObjectSignedURL(filename string, from time.Time, duration time.Duration) (string, error) {
	jsonKey, err := ioutil.ReadFile(SERVICE_ACCOUNT)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        from.Add(duration),
	}
	u, err := storage.SignedURL(BUCKET_NAME, filename, opts)
	if err != nil {
		return "", fmt.Errorf("storage.SignedURL: %v", err)
	}
	return u, nil
}
