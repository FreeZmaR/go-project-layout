package config

import (
	"errors"
	"github.com/FreeZmaR/go-project-layout/config/build"
	"github.com/FreeZmaR/go-project-layout/config/load"
	"github.com/FreeZmaR/go-project-layout/config/types"
)

func LoadInbox() (*types.Inbox, error) {
	type inbox struct {
		Data *types.Inbox `yaml:"inbox"`
	}

	wrapper := &inbox{}

	err := loadByMode(wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Data, nil
}

func LoadOutbox() (*types.Outbox, error) {
	type outbox struct {
		Data *types.Outbox `yaml:"outbox"`
	}

	wrapper := &outbox{}

	err := loadByMode(wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Data, nil
}

func loadByMode(dst any) error {
	if build.IsDevelopMode() {
		return load.Develop(build.ConfigPath(), dst)
	}

	if build.IsProductionMode() {
		return load.Production(build.ConfigPath(), dst)
	}

	return errors.New("unknown mode")
}
