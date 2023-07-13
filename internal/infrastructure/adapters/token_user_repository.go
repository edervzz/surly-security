package adapters

import (
	"context"
	"surly-security/internal/domain/entities"

	"github.com/edervzz/maya"
)

type tokenUserRepository struct {
	dbContext maya.IDbContext
}

func NewTokenUserRepository(dbContext maya.IDbContext) tokenUserRepository {
	return tokenUserRepository{
		dbContext: dbContext,
	}
}

func (r tokenUserRepository) Create(ctx context.Context, entity *entities.TokenUser) error {
	_, err := r.dbContext.Add(ctx, entity)
	return err
}

func (r tokenUserRepository) Update(ctx context.Context, entity *entities.TokenUser) error {
	_, err := r.dbContext.Update(ctx, entity)
	return err
}

func (r tokenUserRepository) Read(key string, ctx context.Context) (entities.TokenUser, error) {
	entity := entities.TokenUser{}

	rows, err := r.dbContext.Read(
		ctx,
		&entities.TokenUser{},
		map[string]any{
			"token": key,
		},
	)

	if err == nil && rows != nil && rows.Next() {
		err = rows.Scan(
			&entity.Token,
			&entity.Type,
			&entity.UserID,
			&entity.ExpiresAt,
			&entity.IsActive,
		)
	}
	return entity, err
}
