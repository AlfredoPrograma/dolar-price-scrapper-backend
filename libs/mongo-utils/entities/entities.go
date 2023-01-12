package entities

import "context"

type Entity interface {
	ToString() string
	Save(ctx context.Context) error
}
