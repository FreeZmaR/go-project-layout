package build

var (
	mode       = DevelopMode
	commitHash = "none"
	configPath = "./../config.yml"
)

func IsProductionMode() bool {
	return mode == ProductionMode
}

func IsDevelopMode() bool {
	return mode == DevelopMode
}

func CommitHash() string {
	return commitHash
}

func ConfigPath() string {
	return configPath
}
