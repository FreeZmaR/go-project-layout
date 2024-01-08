package build

import "fmt"

var (
	mode       = DevelopMode
	commitHash = "none"
	configPath = "./config/config.yaml"
)

func IsProductionMode() bool {
	return mode == ProductionMode
}

func IsDevelopMode() bool {
	return mode == DevelopMode
}

func GetModeString() string {
	return mode.String()
}

func CommitHash() string {
	return commitHash
}

func ConfigPath() string {
	return configPath
}

func Info() string {
	return fmt.Sprintf("mode: %s, commit: %s, config: %s", mode, commitHash, configPath)
}
