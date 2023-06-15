package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Upload struct {
	ArchiveId uuid.UUID `psql:"ARCHIVE_ID"`
	UserId    uuid.UUID `psql:"USER_ID"`
	BasedOn   uuid.UUID `psql:"BASED_ON"`
	CreatedAt time.Time `psql:"CREATED_AT"`
}

var (
	ErrNoResults = errors.New("No results found with given key.")
)

func (u Upload) Get() error {
	var q = `
		SELECT "USER_ID", "BASED_ON", "CREATED_AT"
		FROM "UPLOAD"
		WHERE "ARCHIVE_ID" = '$1';
	`
	var rows, err = db.Query(context.Background(), q, u.ArchiveId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNoResults
		}
		return errors.Wrap(err, "db.Query is failed")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&u.UserId,
			&u.BasedOn,
			&u.CreatedAt,
		)
		if err != nil {
			return errors.Wrap(err, "rows.Scan is failed")
		}
	}
	return nil
}

func (u Upload) Insert() error {
	var q = `
		INSERT INTO "UPLOAD" 
		("USER_ID, "BASED_ON, "CREATED_AT, "ARCHIVE_ID) 
		VALUES ('$1', '$2', '$3', '$4');
	`
	var rows, err = db.Query(context.Background(), q, u.UserId.String(), u.BasedOn, u.CreatedAt, u.ArchiveId)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer rows.Close()
	for rows.Next() {

	}
	return nil
}

func (u Upload) Update() {

}
