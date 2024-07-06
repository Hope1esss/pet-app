package service

type GigaChat interface {
	Giga(inputMessage string) string
}

type Authorization interface {
}

type Service struct {
	Authorization
	GigaChat
}
