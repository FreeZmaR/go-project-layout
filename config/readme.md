# Config layer

Этот уровень обеспечивает настройку системы. Этот уровень отвечает за подготовку и описание конфигурации системы. Слой имеет три подслоя: `build`, `load` и `types`.

В моем примере я использую файл `yaml` в качестве файла конфигурации, а этот файл расположен в `config.config.yaml` как `config.example.yaml`, путь к файлу определяется в `build.configPath`.

---

This layer provides system configuration, The layer responsible for preparing and describing the system configuration.
Layer has three sublayers: `build`, `load` and `types`.

In my example I use `yaml` file as configuration file and this file located in `config.config.yaml` as `config.example.yaml`, path to the file define in `build.configPath`.

## Build

Этот подуровень отвечает за построение конфигурации системы для команды `go build`.
В моем примере я использую `-ldflags` для глобальных переменных, но вы можете использовать любой другой способ создания конфигурации.

---

This sublayer responsible for building system configuration for `go build` command.
In my example I use `-ldflags` for globals var, but you can use any other way to build your configuration.

## Load

Этот подуровень отвечает за загрузку конфигурации системы. Он содержит функцию загрузки, которая вызывается системой во время процесса запуска.
В моем примере я использую пакет `gopkg.in/yaml.v3` для загрузки файла конфигурации `yaml`, но вы можете использовать любой другой способ загрузки конфигурации.

---

This sublayer responsible for loading system configuration. It contains `load` function which is called by the system during the start process.
In my example I use `gopkg.in/yaml.v3` package for loading configuration `yaml` file, but you can use any other way to load your configuration.

## Types

Этот подуровень отвечает за описание конфигурации системы. Он содержит `types`(типы), которые используются системой для описания конфигурации.
В моем примере я использую структуру с тегом `yaml` для описания конфигурации, но вы можете использовать любой другой способ описания конфигурации.

---

This sublayer responsible for describing system configuration. It contains `types` which is used by the system for describing the configuration.
In my example I use `struct` with `yaml tag` for describing the configuration, but you can use any other way to describe your configuration.