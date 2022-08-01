package db

import (
	"context"
	"fmt"
)

type patrimonio struct {
	id     string `json:"id"`
	marca  string `json:"marca"`
	modelo string `json:"modelo"`
	local  string `json:"local"`
}

func (p *patrimonio) GetAll(ctx context.Context) {
	return
}

func (p *patrimonio) Show(ctx context.Context, id string) ([]patrimonio.Patrimonio, error) {

	return nil, nil
}

func (p *patrimonio) Update(ctx context.Context, id string) (bool, error) {

	return false, nil
}

func (p *patrimonio) Delete(ctx context.Context, id string) (bool, error) {

	fmt.Print("Deleting")
	return false, nil
}

func (p *patrimonio) Add(ctx context.Context) (bool, error) {
	fmt.Print("Adding")
	return false, nil
}
