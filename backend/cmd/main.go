package main

import (
	"os"

	_ "blogapp/docs"

	"blogapp/src/core/domain/application/usecases/authusecases"
	"blogapp/src/core/domain/application/usecases/commentusecases"
	"blogapp/src/core/domain/application/usecases/likeusecases"
	"blogapp/src/core/domain/application/usecases/postusecases"
	"blogapp/src/core/domain/application/usecases/profileusecases"

	"blogapp/src/entrypoint/groups"
	"blogapp/src/entrypoint/http/handlers/auth"
	"blogapp/src/entrypoint/http/handlers/comment"
	"blogapp/src/entrypoint/http/handlers/like"
	"blogapp/src/entrypoint/http/handlers/post"
	"blogapp/src/entrypoint/http/handlers/profile"

	// ← Добавь эту строку
	"blogapp/src/entrypoint/http/middleware"   // путь к папке middleware

	"blogapp/src/infrastructure/config"
	"blogapp/src/infrastructure/persistence/repository"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// 1. Загружаем .env
	godotenv.Load("env/.env")

	// 2. Подключаемся к базе
	db := config.ConnectDatabase()
	config.RunMigrations(db)

	// 3. JWT Secret
	jwtSecretStr := os.Getenv("JWT_SECRET")
	if jwtSecretStr == "" {
		jwtSecretStr = "your-super-secret-key-maruf-2026-change-in-production" // fallback
	}
	jwtSecret := []byte(jwtSecretStr)

	// 4. Репозитории
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepositoryImpl(db)
	likeRepo := repository.NewLikeRepositoryImpl(db)
	profileRepo := repository.NewProfileRepositoryImpl(db)

	// 5. UseCases
	registerUC := authusecases.NewRegisterUseCase(userRepo, profileRepo)
	// ← Обнови создание LoginUseCase (передаём jwtSecret!)
	loginUC := authusecases.NewLoginUseCase(userRepo, jwtSecret)

	postCreateUC := postusecases.NewPostUseCase(postRepo)
	postListUC := postusecases.NewPostListUseCase(postRepo)
	postUpdateUC := postusecases.NewPostUpdateUseCase(postRepo)
	postDeleteUC := postusecases.NewPostDeleteUseCase(postRepo)

	commentCreateUC := commentusecases.NewCommentsCreateUseCase(commentRepo)
	commentDeleteUC := commentusecases.NewCommentsDeleteUseCase(commentRepo)

	likeCreateUC := likeusecases.NewLikeUseCase(likeRepo)
	likeDeleteUC := likeusecases.NewLikeDeleteUseCase(likeRepo)

	profileUC := profileusecases.NewProfileUseCase(profileRepo)

	// 6. Handlers
	registerHandler := auth.NewRegisterHandler(registerUC)
	loginHandler := auth.NewLoginHandler(loginUC)

	postCreateHandler := post.NewPostHandler(postCreateUC)
	postListHandler := post.NewPostListHandler(postListUC)
	postUpdateHandler := post.NewPostUpdateHandler(postUpdateUC)
	postDeleteHandler := post.NewPostDeleteHandler(postDeleteUC)

	commentCreateHandler := comment.NewCommentCreateHandler(commentCreateUC)
	commentDeleteHandler := comment.NewCommentDeleteHandler(commentDeleteUC)

	likeCreateHandler := like.NewLikeCreateHandler(likeCreateUC)
	likeDeleteHandler := like.NewLikeDeleteHandler(likeDeleteUC)

	profileUpdateHandler := profile.NewProfileUpdateHandler(profileUC)

	// 7. Groups
	authGroup := groups.NewAuthGroup(registerHandler, loginHandler)
	postGroup := groups.NewPostGroup(postCreateHandler, postListHandler, postUpdateHandler, postDeleteHandler)
	commentGroup := groups.NewCommentGroup(commentCreateHandler, commentDeleteHandler)
	likeGroup := groups.NewLikeGroup(likeCreateHandler, likeDeleteHandler)
	profileGroup := groups.NewProfileGroup(profileUpdateHandler)

	// 8. Echo сервер
	e := echo.New()

	// Публичные роуты (без JWT)
	authGroup.RegisterRoutes(e)

	// Защищённые роуты (с JWT)
	jwtMiddleware := middleware.NewJWT(jwtSecret)


	postGroup.RegisterRoutes(e.Group("/posts", jwtMiddleware))
	commentGroup.RegisterRoutes(e.Group("/comments", jwtMiddleware))
	likeGroup.RegisterRoutes(e.Group("/likes", jwtMiddleware))
	profileGroup.RegisterRoutes(e.Group("/profile", jwtMiddleware))
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}