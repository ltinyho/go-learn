package errors

type CacheError interface {
	Error() string
}

func New(err CacheError) error {
	return err
}
