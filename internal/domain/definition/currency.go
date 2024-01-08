package definition

import "github.com/FreeZmaR/go-project-layout/internal/lib/utils"

// Currency - The purpose of currency in the field of international transfers is to facilitate the exchange of goods
// and services between countries with different monetary systems.
// Currency serves as a medium of exchange, a unit of account, and a store of value.
// It allows businesses and individuals to conduct transactions across borders and to manage their
// financial affairs in a global economy. Currency exchange rates fluctuate based on supply and demand,
// economic conditions, and geopolitical events, and can have a significant impact on trade and investment flows.

const (
	// RUB Russian ruble
	RUB Currency = 1
	// USD United States dollar
	USD Currency = 2
	// EUR Euro is the official currency of 19 of the 27 member states of the European Union
	EUR Currency = 3
)

type Currency int

func (c Currency) String() string {
	switch c {
	case RUB:
		return "RUB"
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	default:
		return "UNKNOWN"
	}
}

func IsCurrency(currency int) bool {
	return utils.OneOf(
		Currency(currency),
		RUB,
		USD,
		EUR,
	)
}
