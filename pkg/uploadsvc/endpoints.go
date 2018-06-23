package uploadsvc

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"

	"github.com/go-kit/kit/endpoint"
	"github.com/schepelin/imgupload/pkg/uploader"
)

type UploadRequest struct {
	Raw string `json:"image"`
}

type UploadResponse struct {
	Link string `json:"link"`
	Err  error  `json:"string,omitempty"`
}

func MakeUploadEndpoint(svc uploader.UploadService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UploadRequest)
		raw, err := base64.StdEncoding.DecodeString(req.Raw)
		if err != nil {
			return nil, err
		}
		b := bytes.NewBuffer(raw)
		_, _, err = image.Decode(b)
		if err != nil {
			return nil, err
		}
		link, err := svc.UploadImage(ctx, raw)
		return UploadResponse{
			Link: link,
			Err:  err,
		}, nil
	}
}

type GetImageRequest struct {
	ImgID string
}

type GetImageResponse struct {
	Img image.Image
}

func MakeGetImageEndpoint(svc uploader.UploadService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetImageRequest)
		imgObj, err := svc.GetImage(ctx, req.ImgID)
		b := bytes.NewBuffer(imgObj.RawData)
		img, _, err := image.Decode(b)
		if err != nil {
			return nil, err
		}
		return GetImageResponse{
			Img: img,
		}, nil
	}
}
