package database

type DB interface {
	InstantiateClient(uri string, readTimeout uint, writeTimeout uint) error
}
