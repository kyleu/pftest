package bar

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	BarFirstValue  = Bar{Key: "first_value", Name: "First Value", Start: "2001-01-01", End: "2001-12-31", Version: "0.0.1"}
	BarSecondValue = Bar{Key: "second_value", Name: "Second Value", Start: "2002-01-01", End: "2002-12-31", Version: "0.0.2"}
	BarUnknown     = Bar{Key: "unknown", Name: "Unknown"}

	AllBars = Bars{BarFirstValue, BarSecondValue, BarUnknown}
)

type Bar struct {
	Key         string
	Name        string
	Description string
	Icon        string

	Start   string
	End     string
	Version string
}

func (b Bar) String() string {
	return b.Key
}

func (b Bar) NameSafe() string {
	if b.Name != "" {
		return b.Name
	}
	return b.String()
}

func (b Bar) Matches(xx Bar) bool {
	return b.Key == xx.Key
}

func (b Bar) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(b.Key, false), nil
}

func (b *Bar) UnmarshalJSON(data []byte) error {
	key, err := util.FromJSONString(data)
	if err != nil {
		return err
	}
	*b = AllBars.Get(key, nil)
	return nil
}

func (b Bar) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(b.Key, start)
}

func (b *Bar) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*b = AllBars.Get(key, nil)
	return nil
}

func (b Bar) Value() (driver.Value, error) {
	return b.Key, nil
}

func (b *Bar) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, ok := converted.(string); ok {
			*b = AllBars.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan Bar enum from value [%v]", value)
}

func BarParse(logger util.Logger, keys ...string) Bars {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) Bar {
		return AllBars.Get(x, logger)
	})
}

type Bars []Bar

func (b Bars) Keys() []string {
	return lo.Map(b, func(x Bar, _ int) string {
		return x.Key
	})
}

func (b Bars) Strings() []string {
	return lo.Map(b, func(x Bar, _ int) string {
		return x.String()
	})
}

func (b Bars) Help() string {
	return "Available bar options: [" + strings.Join(b.Strings(), ", ") + "]"
}

func (b Bars) Get(key string, logger util.Logger) Bar {
	for _, value := range b {
		if strings.EqualFold(value.Key, key) {
			return value
		}
	}
	if key == "" {
		return BarUnknown
	}
	msg := fmt.Sprintf("unable to find [Bar] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return BarUnknown
}

func (b Bars) GetByName(name string, logger util.Logger) Bar {
	for _, value := range b {
		if strings.EqualFold(value.Name, name) {
			return value
		}
	}
	if name == "" {
		return BarUnknown
	}
	msg := fmt.Sprintf("unable to find [Bar] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return BarUnknown
}

func (b Bars) GetByStart(input string, logger util.Logger) Bar {
	for _, value := range b {
		if value.Start == input {
			return value
		}
	}
	if input == "" {
		return BarUnknown
	}
	if logger != nil {
		msg := fmt.Sprintf("unable to find [Bar] with Start [%s]", input)
		logger.Warn(msg)
	}
	return BarUnknown
}

func (b Bars) GetByEnd(input string, logger util.Logger) Bar {
	for _, value := range b {
		if value.End == input {
			return value
		}
	}
	if input == "" {
		return BarUnknown
	}
	if logger != nil {
		msg := fmt.Sprintf("unable to find [Bar] with End [%s]", input)
		logger.Warn(msg)
	}
	return BarUnknown
}

func (b Bars) GetByVersion(input string, logger util.Logger) Bar {
	for _, value := range b {
		if value.Version == input {
			return value
		}
	}
	if input == "" {
		return BarUnknown
	}
	if logger != nil {
		msg := fmt.Sprintf("unable to find [Bar] with Version [%s]", input)
		logger.Warn(msg)
	}
	return BarUnknown
}

func (b Bars) Random() Bar {
	return util.RandomElement(b)
}
