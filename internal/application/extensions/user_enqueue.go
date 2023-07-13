package extensions

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"surly-security/internal/application/resources"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"
	"time"

	"github.com/edervzz/mistake"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func UserEnqueue(ctx context.Context, uow ports.IUnitOfWork, userID string, info string, localizer localizer.ILocalizer) *mistake.M {
	if err := uow.Enqueue(ctx, userID, &entities.User{}, info); err != nil {
		mistake.AppendValue(userID)
		mistake.AppendValue("users")
		return mistake.New(
			resources.VALIDATION,
			resources.DB_ENQUEUE,
			mistake.Formatter(localizer.Localize(resources.DB_ENQUEUE)),
		)
	}
	return nil
}

func CreateJwt(subject string, durationSeconds int, iss string, aud string, key string) (string, error) {
	expireTime := time.Now().Add(time.Duration(durationSeconds) * time.Minute)
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ID:        uuid.NewString(),
		ExpiresAt: jwt.NewNumericDate(expireTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    iss,
		Audience:  []string{aud},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

// Generate a pair of tokens. Access token and Refresh token.
func GenerateTokens(
	ctx context.Context,
	user entities.User,
	subject string,
	userID int64,
	expiresAt int,
	refreshExpiresAt int64,
	logger *zap.Logger,
	localizer localizer.ILocalizer,
	uow ports.IUnitOfWork) (string, string, *mistake.M) {
	// 1. create access token
	accessToken, err := CreateJwt(subject, expiresAt, os.Getenv("Jwt:Issuer"), os.Getenv("Jwt:Audience"), os.Getenv("Jwt:Key"))
	if err != nil {
		logger.Debug(err.Error())
		return "", "", mistake.New(
			resources.SERVER_ERROR,
			resources.CREATE_TOKEN,
			localizer.Localize(resources.CREATE_TOKEN),
		)
	}
	// 2. prepare refresh token
	hashData, err := bcrypt.GenerateFromPassword([]byte(uuid.NewString()), bcrypt.DefaultCost)
	if err != nil {
		logger.Debug(err.Error())
		return "", "", mistake.New(
			resources.SERVER_ERROR,
			resources.CREATE_REFRESH_TOKEN,
			localizer.Localize(resources.CREATE_REFRESH_TOKEN),
		)
	}
	refreshToken := fmt.Sprintf("%x", md5.Sum(hashData))
	tokenUserRefresh := entities.TokenUser{
		Token:     refreshToken,
		Type:      entities.TokenTypeRefresh,
		UserID:    userID,
		ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(refreshExpiresAt)).Format(time.DateTime),
		IsActive:  true,
	}
	// 3. enqueue user
	defer uow.Dequeue(ctx)
	if mt := UserEnqueue(ctx, uow, fmt.Sprintf("%d", userID), "LoginHandler", localizer); mt != nil {
		return "", "", mt
	}
	logger.Debug("entity enqueued")
	// 4. save entities
	if err = uow.BeginTransaction(ctx); err == nil {
		if err = uow.TokenUser().Create(ctx, &tokenUserRefresh); err == nil {
			if user.Intents > 0 {
				user.Intents = 0
				err = uow.User().Update(ctx, &user)
			}
			if err == nil {
				err = uow.CommitWork()
			}
		}
	}

	if err != nil {
		defer uow.RollbackWork()
		return "", "", mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)
	}

	return accessToken, refreshToken, nil

}
