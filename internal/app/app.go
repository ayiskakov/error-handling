package app

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Exec() error {
	return nil
}
