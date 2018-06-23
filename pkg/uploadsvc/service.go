package uploadsvc

import (
	"context"

	"github.com/schepelin/imgupload/pkg/storage"
	"github.com/schepelin/imgupload/pkg/uploader"
)

// UploadService contains all the dependecines to perform use cases
// impelements uploader.UploadService interface
type UploadService struct {
	Storage   storage.ImagesStorage
	Shortener uploader.URLShortener
	Hasher    uploader.Hasher
}

func (us *UploadService) GetImage(ctx context.Context, imgID string) (*uploader.Image, error) {
	return us.Storage.GetImage(ctx, imgID)
}

func (us *UploadService) UploadImage(ctx context.Context, raw []byte) (string, error) {
	id := us.Hasher.Generate(raw)

	img := uploader.Image{
		ID:      id,
		RawData: raw,
		Link:    us.Shortener.MakeShortURL(id),
	}
	err := us.Storage.SaveImage(ctx, img)
	switch {
	case err == uploader.ErrImgExists:
		return img.Link, nil
	case err != nil:
		return "", err
	default:
		return img.Link, nil
	}
}
