# Config layer

This layer provides configuration for the system. This layer responsible for preparing and describing the system configuration.
Layer has three sublayers: `build`, `load` and `types`.

In my example I use `yaml` file as configuration file and this file located in `config.config.yaml` as `config.example.yaml`, path to the file define in `build.configPath`.

## Build

This sublayer responsible for building the system configuration. It contains `build` function which is called by the system during the build process.
In my example I use `-ldflags` for globals var, but you can use any other way to build your configuration.

## Load

This sublayer responsible for loading the system configuration. It contains `load` function which is called by the system during the start process.
In my example I use `gopkg.in/yaml.v3` package for loading configuration `yaml file`, but you can use any other way to load your configuration.

## Types

This sublayer responsible for describing the system configuration. It contains `types` which is used by the system for describing the configuration.
In my example I use `struct` with `yaml tag` for describing the configuration, but you can use any other way to describe your configuration.