package build

const (
	DevelopMode    Mode = 1
	ProductionMode Mode = 2
)

type Mode int

func (m Mode) String() string {
	switch m {
	case DevelopMode:
		return "develop"
	case ProductionMode:
		return "production"
	default:
		return "unknown"
	}
}
