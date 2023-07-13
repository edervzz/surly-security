package adapters

import (
	"context"
	"surly-security/internal/domain/ports"

	"github.com/edervzz/maya"
)

type unitOfWork struct {
	dbContext            maya.IDbContext
	userRepository       ports.IUserRepository
	userTokensRepository ports.ITokenUserRepository
}

func NewUnitOfWork(dbContext maya.IDbContext) unitOfWork {
	return unitOfWork{
		dbContext:            dbContext,
		userRepository:       NewUserRepository(dbContext),
		userTokensRepository: NewTokenUserRepository(dbContext),
	}
}

func (uow unitOfWork) BeginTransaction(ctx context.Context) error {
	return uow.dbContext.BeginTransaction(ctx)
}

func (uow unitOfWork) User() ports.IUserRepository {
	return uow.userRepository
}

func (uow unitOfWork) TokenUser() ports.ITokenUserRepository {
	return uow.userTokensRepository
}

func (uow unitOfWork) CommitWork() error {
	return uow.dbContext.Commit()
}

func (uow unitOfWork) RollbackWork() error {
	return uow.dbContext.Rollback()
}

func (uow unitOfWork) Enqueue(ctx context.Context, id string, entityPtr any, info string) error {
	return uow.dbContext.Enqueue(ctx, id, entityPtr, info)
}

func (uow unitOfWork) Dequeue(ctx context.Context) error {
	return uow.dbContext.Dequeue(ctx)
}
