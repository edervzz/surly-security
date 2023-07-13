package commands

import (
	"context"
	"fmt"
	"surly-security/internal/application/extensions"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/application/validators"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"

	"go.uber.org/zap"
)

func UserConfirmEmailHandler(
	ctx context.Context,
	uow ports.IUnitOfWork,
	localizer localizer.ILocalizer,
	logger *zap.Logger,
	request messages.UserConfirmEmailRequest,
) (messages.UserConfirmEmailResponse, *mistake.M) {
	// 1. request validations
	if mt := validators.UserConfirmEmailRequestValidator(request, localizer); mt != nil {
		return messages.UserConfirmEmailResponse{}, mt
	}
	logger.Debug("request validated")
	// 2. business validations
	if mt := validators.UserConfirmEmailRequestBizValidator(ctx, uow, &request, localizer); mt != nil {
		return messages.UserConfirmEmailResponse{}, mt
	}
	logger.Debug("business rules validated")
	// 3. set entity's value
	request.User.IsActive = true
	request.TokenUser.IsActive = false
	// 4. enqueue main entity
	defer uow.Dequeue(ctx)
	if mt := extensions.UserEnqueue(ctx, uow, fmt.Sprintf("%d", request.User.ID), "UserConfirmEmailHandler", localizer); mt != nil {
		return messages.UserConfirmEmailResponse{}, mt
	}
	// 5. save entities
	defer uow.RollbackWork()
	err := *new(error)
	if err = uow.BeginTransaction(ctx); err == nil {
		if err = uow.User().Update(ctx, &request.User); err == nil {
			if err = uow.TokenUser().Update(ctx, &request.TokenUser); err == nil {
				err = uow.CommitWork()
			}
		}
	}
	// 5.1 handle error
	if err != nil {
		return messages.UserConfirmEmailResponse{}, mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)
	}

	return messages.UserConfirmEmailResponse{IsSuccess: true}, nil
}
