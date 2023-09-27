package main

import (
	"context"
)

type Bats struct{}

func (m *Bats) Bats(ctx context.Context, dir *Directory, args []string) (*Container, error) {
	return run(dir, args), nil
}

func (m *Bats) Foo(ctx context.Context) (*Directory, error) {
	return nil, nil
}

func (dir *Directory) Bats(ctx context.Context, args []string) (*Container, error) {
	return run(dir, args), nil
}

func run(dir *Directory, args []string) *Container {
	return image().
		WithMountedDirectory("/src", dir).
		WithWorkdir("/src").
		WithExec(args)
}

func image() *Container {
	return dag.
		Container().
		From("bats/bats:v1.10.0")
}
