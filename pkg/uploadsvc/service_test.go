package uploadsvc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/schepelin/imgupload/pkg/mocks"
	"github.com/schepelin/imgupload/pkg/uploader"
	"github.com/schepelin/imgupload/pkg/uploadsvc"
	"github.com/stretchr/testify/assert"
)

func TestServiceDesign_UploadImage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	stMock := mocks.NewMockImagesStorage(mockCtrl)
	hasherMock := mocks.NewMockHasher(mockCtrl)
	shortenerMock := mocks.NewMockURLShortener(mockCtrl)

	ctx := context.TODO()
	service := &uploadsvc.UploadService{
		Storage:   stMock,
		Hasher:    hasherMock,
		Shortener: shortenerMock,
	}
	expectedImg := uploader.Image{
		ID:      "img-hash",
		Link:    "http://somewhere",
		RawData: []byte{10, 42, 15},
	}

	gomock.InOrder(
		hasherMock.EXPECT().Generate(expectedImg.RawData).Return(expectedImg.ID),
		shortenerMock.EXPECT().MakeShortURL(expectedImg.ID).Return(expectedImg.Link),
		stMock.EXPECT().SaveImage(ctx, expectedImg).Return(uploader.ErrImgExists),
	)

	link, err := service.UploadImage(ctx, expectedImg.RawData)
	assert.NoError(t, err)
	assert.Equal(t, expectedImg.Link, link)
}
