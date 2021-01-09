package app

import "context"

type FighterService interface {
	UnitByID(ctx context.Context, id string) (Unit, error)
	ListAllUnits(ctx context.Context) ([]Unit, error)
}

type MoveService interface {
	MoveByID(ctx context.Context, id string) (Move, error)
}

type UnitMoveService interface {
	ListUnitMovesForUnit(ctx context.Context, unit_id string) ([]Move, error)
}
