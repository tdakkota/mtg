package config2

import (
	"fmt"
	"strconv"
)

type TypeConcurrency struct {
	Value uint
}

func (t *TypeConcurrency) Set(value string) error {
	concurrencyValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fmt.Errorf("Value is not uint (%s): %w", value, err)
	}

	if concurrencyValue == 0 {
		return fmt.Errorf("Value should be >0 (%s)", value)
	}

	t.Value = uint(concurrencyValue)

	return nil
}

func (t TypeConcurrency) Get(defaultValue uint) uint {
	if t.Value == 0 {
		return defaultValue
	}

	return t.Value
}

func (t *TypeConcurrency) UnmarshalJSON(data []byte) error {
	return t.Set(string(data))
}

func (t TypeConcurrency) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t TypeConcurrency) String() string {
	return strconv.FormatUint(uint64(t.Value), 10)
}
