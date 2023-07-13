package ports

import "context"

type IUnitOfWork interface {
	BeginTransaction(ctx context.Context) error
	User() IUserRepository
	TokenUser() ITokenUserRepository
	CommitWork() error
	RollbackWork() error
	Enqueue(ctx context.Context, id string, entity any, info string) error
	Dequeue(ctx context.Context) error
}
