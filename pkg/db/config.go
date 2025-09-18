package db

type Config struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
}
