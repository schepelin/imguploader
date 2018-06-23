package storage

import (
	"context"

	"github.com/schepelin/imgupload/pkg/uploader"
)

type ImagesStorage interface {
	SaveImage(ctx context.Context, img uploader.Image) error
	GetImage(ctx context.Context, imgID string) (img *uploader.Image, err error)
}
