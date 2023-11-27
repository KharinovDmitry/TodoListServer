package app

import (
	"TodoListServer/internal/config"
	"TodoListServer/internal/server/handlers"
	"TodoListServer/internal/server/middleware"
	"TodoListServer/internal/services"
	"TodoListServer/internal/storage"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	port    int
	router  *mux.Router
	log     *slog.Logger
	storage *storage.Storage

	boardService handlers.IBoardService
	userService  handlers.IUserService
}

func New(cfg *config.Config) (*App, error) {
	log := setupLogger()

	router := mux.NewRouter()

	strg, err := storage.New(cfg.ConnStr)
	if err != nil {
		return nil, err
	}

	boardService := services.NewBoardService(log, strg.Boards)
	userService := services.NewUserService(log, strg.Users, cfg.TokenTTL)

	app := &App{
		port:         cfg.Port,
		router:       router,
		log:          log,
		storage:      strg,
		boardService: boardService,
		userService:  userService,
	}

	app.setupRouter()

	log.Info("App was configured")
	return app, nil
}

func (a *App) MustRun() {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.router); err != nil {
		panic(err.Error())
	}

	a.log.Info("App was started")
	defer a.storage.Close()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	//TODO GRACEFUL SHUTDOWN
	a.log.Info("App was closed")
}

func (a *App) setupRouter() {

	boardSubrouter := a.router.PathPrefix("/board").Subrouter()
	boardSubrouter.HandleFunc("/update", handlers.UpdateBoard(a.log, a.boardService)).Methods("POST")
	boardSubrouter.HandleFunc("/create", handlers.CreateBoard(a.log, a.boardService)).Methods("POST")
	boardSubrouter.HandleFunc("/delete/{id}", handlers.DeleteBoard(a.log, a.boardService)).Methods("POST")
	boardSubrouter.HandleFunc("/{id}", handlers.GetBoard(a.log, a.boardService)).Methods("GET")
	boardSubrouter.Use(middleware.Auth)

	a.router.HandleFunc("/sign-up", handlers.RegisterUser(a.log, a.userService)).Methods("POST")
	a.router.HandleFunc("/sign-in", handlers.LoginUser(a.log, a.userService)).Methods("POST")

	a.router.HandleFunc("/user/{id}", handlers.GetUser(a.log, a.userService)).Methods("GET")

	a.router.Use(middleware.Logging(a.log))
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil)) //TODO
}
