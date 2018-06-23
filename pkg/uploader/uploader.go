package uploader

import (
	"context"
)

// Image entity
type Image struct {
	ID      string
	RawData []byte
	Link    string
}

// ImageUploadService represents use cases with Image entity
type ImageUploadService interface {
	UploadImage(ctx context.Context, raw []byte) (link string, err error)
	GetImage(ctx context.Context, imgID string) (*Image, error)
}

// Hasher generate unique ID for image based on its content
type Hasher interface {
	Generate([]byte) string
}

// URLShortener produce short url for image ID
type URLShortener interface {
	MakeShortURL(imgID string) string
}
