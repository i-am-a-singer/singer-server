package database

import "github.com/boltdb/bolt"

// Entity is an interface for dealing with database entity
// Type return type name which is key of correspond bolt database scheme
type Entity interface {
	Type() string
	Encode() ([]byte, error)
	Decode(seq []byte) error
}

// DB is an interface for database CURD operatation with Entity object.
type DB interface {
	OpenDB(loader func(*bolt.Tx) error)
	Close() error
	Query(id int, e *Entity)
}
