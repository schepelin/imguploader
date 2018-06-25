package uploadsvc

import (
	"context"
	"encoding/json"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/schepelin/imgupload/pkg/uploader"
)

func DecodeGetImageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("bad roure")
	}
	return GetImageRequest{
		ImgID: id,
	}, nil
}

func EncodeGetImageResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetImageResponse)
	switch resp.Format {
	case "png":
		return png.Encode(w, resp.Img)
	case "jpeg":
		return jpeg.Encode(w, resp.Img, nil)
	case "gif":
		return gif.Encode(w, resp.Img, nil)
	default:
		w.WriteHeader(http.StatusTeapot)
	}
	return nil
}

func MakeGetImageHandler(svc uploader.UploadService) http.Handler {
	return kithttp.NewServer(
		MakeGetImageEndpoint(svc),
		DecodeGetImageRequest,
		EncodeGetImageResponse,
	)
}

func DecodeUploadImageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func MakeUploadImageHandler(svc uploader.UploadService) http.Handler {
	return kithttp.NewServer(
		MakeUploadEndpoint(svc),
		DecodeUploadImageRequest,
		kithttp.EncodeJSONResponse,
	)
}

func MakeRouter(svc uploader.UploadService) http.Handler {
	m := mux.NewRouter()
	m.Handle("/images", MakeUploadImageHandler(svc)).Methods("POST")
	m.Handle("/images/{id}", MakeGetImageHandler(svc)).Methods("GET")

	return m
}
