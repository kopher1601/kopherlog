package repository

import "kopherlog/ent"

type EntRepository[T any] interface {
	Save(t T) error
}

type entRepository struct {
	client *ent.Client
}
