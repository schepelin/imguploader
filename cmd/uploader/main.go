package main

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/schepelin/imgupload/pkg/postgres"
	"github.com/schepelin/imgupload/pkg/uploadsvc"
)

type noopShortener struct{}

func (ns noopShortener) MakeShortURL(id string) string {
	return id
}

type HasherMD5 struct{}

func (h HasherMD5) Generate(raw []byte) string {
	hash := md5.New()
	hash.Write(raw)

	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := Run(ctx, logger); err != nil {
		logger.Fatal(err)
	}
}

func Run(ctx context.Context, logger *log.Logger) error {
	db, err := sql.Open("postgres", os.Getenv("UPLOADER_DB"))
	defer db.Close()
	if err != nil {
		return err
	}
	pgSt := postgres.Storage{DB: db}
	uplSvc := uploadsvc.UploadService{
		Storage:   &pgSt,
		Shortener: noopShortener{},
		Hasher:    HasherMD5{},
	}

	srv := &http.Server{
		Addr:    os.Getenv("UPLOADER_LISTEN"),
		Handler: uploadsvc.MakeRouter(&uplSvc),
	}
	errs := make(chan error, 1)
	go func() {
		logger.Printf("Server is starting %v", srv.Addr)
		errs <- srv.ListenAndServe()
	}()

	return <-errs
}
