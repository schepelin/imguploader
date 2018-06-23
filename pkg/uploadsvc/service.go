package uploadsvc

import (
	"context"

	"github.com/schepelin/imgupload/pkg/storage"
	"github.com/schepelin/imgupload/pkg/uploader"
)

// UploadService contains all the dependecines to perform use cases
// impelement uploader.ImageUploadService interface
type UploadService struct {
	Storage   storage.ImagesStorage
	Shortener uploader.URLShortener
	Hasher    uploader.Hasher
}

func (us *UploadService) GetImage(
	ctx context.Context, imgID string,
) (*uploader.Image, error) {
	return us.Storage.GetImage(ctx, imgID)
}
