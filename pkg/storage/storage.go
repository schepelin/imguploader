package storage

import (
	"context"

	"github.com/schepelin/imgupload/pkg/uploader"
)

//go:generate mockgen -source=storage.go -destination ../mocks/mock_storage.go -package mocks

type ImagesStorage interface {
	SaveImage(ctx context.Context, img uploader.Image) error
	GetImage(ctx context.Context, imgID string) (img *uploader.Image, err error)
}
