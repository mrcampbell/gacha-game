package app

import "context"

type UnitService interface {
	UnitByID(ctx context.Context, id string) (Unit, error)
	ListAllUnits(ctx context.Context) ([]Unit, error)
	CreateUnit(ctx context.Context, unit Unit) (Unit, error)
}

type MoveService interface {
	MoveByID(ctx context.Context, id string) (Move, error)
}

type UnitMoveService interface {
	ListUnitMovesForUnit(ctx context.Context, unitID string) ([]Move, error)
	ListUnitMovesAtLevelForUnit(ctx context.Context, unitID string, level int32) ([]Move, error)
}

type FighterService interface {
	CreateFighter(ctx context.Context, userID, unitID string, level int32) (Fighter, error)
	FighterByID(ctx context.Context, userID, fighterID string) (Fighter, error)
	FightersByUserID(ctx context.Context, userID string) ([]Fighter, error)
}
