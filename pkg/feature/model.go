package feature

import "time"

type Feature struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Active       bool      `db:"active"`
	Responsible  int       `db:"responsible"`
	CreationDate time.Time `db:"creation_date"`
	UpdateDate   time.Time `db:"update_date"`
}
