package lan

import (
	"net/http"
)

var LanApp *App

type App struct {
	Handlers *handler
	Server *http.Server
}

func init(){
	LanApp = NewApp()
}

func NewApp() *App{
	app := &App{Handlers: NewHandler(), Server: &http.Server{}}
	return app
}

func (app *App) Run(addr string){
	app.Server.Handler = app.Handlers
	app.Server.Addr = addr
	app.Server.ListenAndServe()
}