package base

import "time"

type Base struct {
	Id         uint      `json:"id" db:"Id"`
	IsActive   bool      `json:"isactive" db:"IsActive"`
	IsDeleted  bool      `json:"isdeleted" db:"IsDeleted"`
	CreateDate time.Time `json:"createdate" db:"CreateDate"`
	LastUpdate time.Time `json:"lastupdate" db:"LastUpdate"`
}
