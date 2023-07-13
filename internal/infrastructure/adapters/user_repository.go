package adapters

import (
	"context"
	"surly-security/internal/domain/entities"

	"github.com/edervzz/maya"
)

type userRepository struct {
	dbContext maya.IDbContext
}

func NewUserRepository(dbContext maya.IDbContext) userRepository {
	return userRepository{
		dbContext: dbContext,
	}
}

func (r userRepository) Create(ctx context.Context, entity *entities.User) error {
	sqlresult, err := r.dbContext.Add(ctx, entity)
	entity.ID, err = sqlresult.LastInsertId()
	return err
}

func (r userRepository) Update(ctx context.Context, entity *entities.User) error {
	_, err := r.dbContext.Update(ctx, entity)
	return err
}

func (r userRepository) ReadByExternalID(ctx context.Context, email string) (entities.User, error) {
	user := entities.User{}

	rows, err := r.dbContext.Read(
		ctx,
		&entities.User{},
		map[string]any{
			"email": email,
		},
	)

	if err == nil && rows != nil && rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Fullname,
			&user.PasswordHash,
			&user.EmailConfirmed,
			&user.PhoneConfirmed,
			&user.IsLocked,
			&user.IsActive,
			&user.Intents,
		)
	}

	return user, err
}

func (r userRepository) Read(id int64, ctx context.Context) (entities.User, error) {
	user := entities.User{}

	rows, err := r.dbContext.Read(
		ctx,
		&entities.User{},
		map[string]any{
			"id": id,
		},
	)

	if err == nil && rows != nil && rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Fullname,
			&user.PasswordHash,
			&user.EmailConfirmed,
			&user.PhoneConfirmed,
			&user.IsLocked,
			&user.IsActive,
			&user.Intents,
		)
	}

	return user, err
}
