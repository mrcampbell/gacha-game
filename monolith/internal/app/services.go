package app

import "context"

type FighterService interface {
	UnitByID(ctx context.Context, id string) (Unit, error)
	ListAllUnits(ctx context.Context) ([]Unit, error)
}
