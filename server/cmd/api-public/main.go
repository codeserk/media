package main

import (
	"fmt"
	"media/internal/cache"
	"media/internal/controller"
	"media/internal/controller/signature"
	"media/internal/email"
	"media/internal/jwt"
	"media/internal/logger"
	"media/internal/media/mediamod"
	"media/internal/mongo"
	"media/module/ai/aimod"
	"media/module/book/bookmod"
	"media/module/user/usermod"
	"net/http"
	"path/filepath"
	"strconv"

	_ "media/cmd/api-public/docs"
	"media/cmd/api-public/module/auth"
	"media/cmd/api-public/module/book"
	"media/cmd/api-public/module/health"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title                       Media API / Public
// @version                     1.0
// @description                 Public API for @codeserk / media
// @BasePath                    /
// @securityDefinitions.apikey  SignatureApp
// @in                          header
// @name                        x-signature-app
// @securityDefinitions.apikey  Signature
// @in                          header
// @name                        x-signature
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	log := zerolog.New(logger.Output()).With().
		Timestamp().
		Logger()

	conf, err := ParseConfig()
	if err != nil {
		log.Fatal().Err(fmt.Errorf("failed to load config: %v", err)).Msg("")
		return
	}

	cacheService, err := cache.New(&conf.Redis)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("load cache: %v", err)).Msg("")
		return
	}

	db, err := mongo.New(&conf.Mongo)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("load mongo: %v", err)).Msg("")
		return
	}

	storage := mediamod.NewStorageService(&conf.Media)
	images := mediamod.NewImageService(&conf.Media, storage)
	signaturesService := signature.NewService(&conf.API.Signature)
	ai := aimod.NewChatService(&conf.AI)
	emails := email.New(&conf.Email)

	jwtService := jwt.New(conf.JWT)
	userRepo := usermod.NewRepository(db)
	userAuth := usermod.NewAuthService(conf.Dashboard, userRepo, cacheService, jwtService, emails)

	bookRepository := bookmod.NewRepository(db)
	bookSources := bookmod.NewSourceService(&conf.BookSource)
	bookProcess := bookmod.NewProcessService(images, ai)
	bookRead := bookmod.NewReadService(bookRepository, bookSources, bookProcess)

	router := mux.NewRouter()
	router.Use(controller.CorsMiddleware)
	health.Handle(router)

	sr := router.PathPrefix("/api/v1/").Subrouter()
	sr.Use(signaturesService.Middleware(signature.SignatureAppDashboard))
	sr.Use(userAuth.Middleware)
	auth.Handle(sr, userAuth)
	book.Handle(sr, bookRead)

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/docs.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("cmd", "api", "docs", "swagger.json"))
	}).Methods(http.MethodGet)

	router.Use(logger.Middleware())
	http.Handle("/", router)

	log.Printf("server started at http://0.0.0.0:%v/docs/", conf.API.Port)

	if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(conf.API.Port), router); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}
