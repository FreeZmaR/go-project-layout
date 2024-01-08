package load

import (
	"fmt"
)

func newLoadFile(details string) error {
	return fmt.Errorf("error load file: %s", details)
}

func newUnmarshalError(details string) error {
	return fmt.Errorf("error unmarshal: %s", details)
}

func newDecodeError(details string) error {
	return fmt.Errorf("error decode security content: %s", details)
}
