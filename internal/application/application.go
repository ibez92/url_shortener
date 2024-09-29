package application

type Server interface {
	Start() error
	Stop() error
}

type Application struct {
	server Server
}

func NewApplication(s Server) *Application {
	return &Application{
		server: s,
	}
}

func (a *Application) Run() error {
	sErr := a.server.Start()
	if sErr != nil {
		return sErr
	}

	return nil
}

func (a *Application) Stop() error {
	sErr := a.server.Stop()
	if sErr != nil {
		return sErr
	}

	return nil
}
