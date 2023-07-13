package main

import (
	"fmt"
	"net/http"
	"os"
	"surly-security/docs"
	"surly-security/internal/application/resources"
	"surly-security/internal/domain/ports"
	"surly-security/internal/infrastructure/adapters"
	"surly-security/internal/infrastructure/migrations"

	"surly-security/internal/webapi/endpoints"
	"surly-security/internal/webapi/middlewares"
	"surly-security/toolkit/localizer"

	"surly-security/toolkit/services"

	"github.com/edervzz/maya"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	_ "github.com/go-sql-driver/mysql"

	midd "github.com/go-chi/chi/middleware"
)

// @title           Security API
// @version         1.0
// @description     Security API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Eder Velázquez
// @contact.url    https://www.linkedin.com/in/oscar-eder-vel%C3%A1zquez-pineda/
// @contact.email  edervzz.work@gmail.com

// @host      localhost:6001
// @BasePath  /security/api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// envars
	basePath := os.Getenv("BASE_PATH")
	appPort := os.Getenv("APP_PORT")
	server := os.Getenv("DB_SERVER")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USR")
	pass := os.Getenv("DB_PWD")

	/*
	 * localizer
	 */
	services.AddSingleton[localizer.ILocalizer](localizer.NewLocalizer(localizer.EN, resources.SecurityMessages))
	/*
	 * logger
	 */
	if logger, err := zap.NewDevelopment(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		services.AddSingleton[*zap.Logger](logger)
	}
	/*
	 * unit of work
	 */
	dbContext := maya.NewDbContext(
		maya.ConnectionString{
			Server: server,
			DbName: dbName,
			Port:   port,
			User:   user,
			Pass:   pass,
		},
		services.Get[*zap.Logger](),
		[]maya.Migration{
			migrations.SecurityDBTables,
		},
	)
	services.AddSingleton[ports.IUnitOfWork](adapters.NewUnitOfWork(dbContext))
	/*
	 * routing
	 */
	r := chi.NewRouter()
	r.Use(midd.AllowContentType("application/json"))
	r.Use(midd.CleanPath)
	r.Use(midd.Logger)
	r.Use(middlewares.UseJsonResponse)
	r.Use(middlewares.UseUserInfo)
	r.Use(middlewares.UseLocalizerLanguage)
	/*5
	 * public endpoints
	 */
	r.Post(basePath+"/users/sign-up", endpoints.SignUp)
	r.Post(basePath+"/users/confirm/email/{token}", endpoints.ConfirmUserByEmail)
	r.Post(basePath+"/users/login", endpoints.Login)
	r.Post(basePath+"/users/login/refresh/{token}", endpoints.RefreshLogin)
	r.Post(basePath+"/users/password/forgot", endpoints.ForgotPassword)
	r.Post(basePath+"/users/password/reset", endpoints.ResetForgotPassword)
	/*
	 * private endpoints
	 */
	r.Group(func(r chi.Router) {
		tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("Jwt:Key")), nil)
		r.Use(jwtauth.Verifier(tokenAuth), jwtauth.Authenticator)
		r.Put(basePath+"/users/me/password", endpoints.ChangePassword)
	})
	/*
	 * Swagger
	 */
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Host = "localhost:" + appPort
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("/swagger/doc.json"))))

	fmt.Println(fmt.Sprintf("listening on port %s", appPort))
	http.ListenAndServe(fmt.Sprintf(":%s", appPort), r)
}
