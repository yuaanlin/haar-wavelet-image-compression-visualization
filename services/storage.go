package services

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/ken20001207/image-compressor/redis"
	"golang.org/x/image/bmp"
	"image"
	"time"
)

// SaveImage : Save image as bytes into redis can return an uid.
// image will expire after 10 minutes
func SaveImage(ctx context.Context, img *image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := bmp.Encode(buf, *img)
	if err != nil {
		return "", err
	}

	uid := uuid.New().String()[:8]
	send := buf.Bytes()
	redis.GetClient().Set(ctx, uid, send, 10*time.Minute)
	return uid, nil
}

// GetImage : Get image from redis by uid
func GetImage(ctx context.Context, uid string) (*image.Image, error) {
	by, err := redis.GetClient().Get(ctx, uid).Bytes()
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(by)
	img, err := bmp.Decode(r)
	if err != nil {
		return nil, err
	}

	return &img, nil
}
