package ports

import (
	"surly-security/internal/domain/entities"

	"github.com/edervzz/maya/abstractions"
)

type ITokenUserRepository interface {
	abstractions.ICreateRepository[*entities.TokenUser]
	abstractions.IReadSingleRepository[string, entities.TokenUser]
	abstractions.IUpdateRepository[*entities.TokenUser]
}
