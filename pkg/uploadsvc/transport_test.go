package uploadsvc_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/schepelin/imgupload/pkg/uploadsvc"
	"github.com/stretchr/testify/assert"
)

func TestTransportGetImage_Handler(t *testing.T) {
	r := uploadsvc.GetImageResponse{
		Img:    createSampleImage(),
		Format: "png",
	}
	w := httptest.NewRecorder()
	err := uploadsvc.EncodeGetImageResponse(context.TODO(), w, r)

	assert.NoError(t, err)
	resp := w.Result()
	assert.Equal(t, "image/png", resp.Header.Get("Content-Type"))
}
