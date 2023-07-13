package ports

import (
	"surly-security/internal/domain/entities"

	"github.com/edervzz/maya/abstractions"
)

type IUserRepository interface {
	abstractions.ICreateRepository[*entities.User]
	abstractions.IUpdateRepository[*entities.User]
	abstractions.IReadByExternalIDRepository[string, entities.User]
	abstractions.IReadSingleRepository[int64, entities.User]
}
