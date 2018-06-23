package uploadsvc_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/schepelin/imgupload/pkg/mocks"
	"github.com/schepelin/imgupload/pkg/uploadsvc"
	"github.com/stretchr/testify/assert"
)

func createSampleImage() image.Image {
	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})
	return sampleImg
}

func imageToByte(img image.Image) []byte {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func ImageToBase64(img image.Image) string {
	raw := imageToByte(img)
	return base64.StdEncoding.EncodeToString(raw)
}

func TestServiceEndpoint_UploadImage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSvc := mocks.NewMockUploadService(mockCtrl)
	ctx := context.TODO()
	imgStub := createSampleImage()
	expectedRaw := imageToByte(imgStub)
	expectedLink := "http://somewhere"

	endpoint := uploadsvc.MakeUploadEndpoint(mockSvc)
	mockSvc.EXPECT().UploadImage(ctx, expectedRaw).Return(expectedLink, nil)
	req := uploadsvc.UploadRequest{
		Raw: ImageToBase64(imgStub),
	}
	r, err := endpoint(ctx, req)
	assert.NoError(t, err)
	resp := r.(uploadsvc.UploadResponse)
	assert.Equal(t, expectedLink, resp.Link)
}
