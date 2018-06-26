package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/schepelin/imgupload/pkg/uploader"
)

type Storage struct {
	DB *sql.DB
}

func (s *Storage) SaveImage(ctx context.Context, img uploader.Image) error {
	_, err := s.DB.Exec(
		`insert into images(id, raw, link) values($1, $2, $3)`,
		img.ID, string(img.RawData), img.Link,
	)
	return err
}

func (s *Storage) GetImage(ctx context.Context, imgID string) (*uploader.Image, error) {
	var raw, link string
	err := s.DB.QueryRow("select raw, link from images where id=$1", imgID).Scan(
		&raw, &link,
	)
	if err != nil {
		return nil, err
	}
	return &uploader.Image{
		ID:      imgID,
		RawData: []byte(raw),
		Link:    link,
	}, nil
}
