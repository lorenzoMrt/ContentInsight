package mysql

import "time"

const (
	sqlContentTable = "contents"
)

type sqlContent struct {
	Uuid            string    `db:"uuid"`
	Title           string    `db:"title"`
	Description     string    `db:"description"`
	ContentType     string    `db:"contentType"`
	Categories      []byte    `db:"categories"`
	Tags            []byte    `db:"tags"`
	Author          string    `db:"author"`
	PublicationDate time.Time `db:"publicationDate"`
	ContentURL      string    `db:"contentUrl"`
	Duration        *int      `db:"duration"`
	Language        string    `db:"language"`
	CoverImage      string    `db:"coverImage"`
	Metadata        []byte    `db:"metadata"`
	Status          string    `db:"status"`
	Source          string    `db:"source"`
	Visibility      string    `db:"visibility"`
}
