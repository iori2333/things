package things

import (
	"context"
	"os"

	"gopkg.in/yaml.v3"
	core "things/core/service"
	gateway "things/gateway/service"
	persistence "things/persistence/service"
)

const kDefaultConfig = "config.yaml"

var config struct {
	Things      *core.Config        `yaml:"things,omitempty"`
	Gateway     *gateway.Config     `yaml:"gateway,omitempty"`
	Persistence *persistence.Config `yaml:"persistence,omitempty"`
}

func Start(ctx context.Context, path string) error {
	if path == "" {
		path = kDefaultConfig
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	if err = yaml.NewDecoder(f).Decode(&config); err != nil {
		return err
	}

	core.Start(ctx, config.Things)
	gateway.Start(ctx, config.Gateway)
	persistence.Start(ctx, config.Persistence)
	return nil
}
