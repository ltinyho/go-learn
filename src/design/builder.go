package design

type Storage struct {
	options clientOptions
}

func NewStorage(options ...StorageOption) Storage {
	o := clientOptions{}
	for _, f := range options {
		f(&o)
	}
	return Storage{
		options: o,
	}
}

type clientOptions struct {
	name string
	ok   bool
	Options
}

type Options struct {
	namespace string
}

type StorageOption func(options *clientOptions)

func WithName(name string) StorageOption {
	return func(options *clientOptions) {
		options.name = name
	}
}
func WithOk(ok bool) StorageOption {
	return func(options *clientOptions) {
		options.ok = ok
	}
}
