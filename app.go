package cautious_giggle

import (
	"github.com/tinrab/cautious-giggle/config"
	"github.com/tinrab/cautious-giggle/repositories"
)

type App struct {
	cfg config.Config
}

func NewApp(cfg config.Config) App {
	return App{
		cfg,
	}
}

func (a App) Run() error {
	_, err := repositories.NewStoreRepository(a.cfg.StoreRepository)
	if err != nil {
		return err
	}
	return nil
}
