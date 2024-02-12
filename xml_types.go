package financials

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/pkg/errors"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

type DateTime struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02T15:04:05"))
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("02/01/2006", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("02.01.2006", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(dt.Time.Format("2006-01-02T15:04:05.999"))
}

func (dt *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	dt.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	dt.Time, err = time.Parse("2006-01-02T15:04:05.999", value)
	return err
}

func (dt DateTime) IsEmpty() bool {
	return dt.Time.IsZero()
}

type Bool bool

func (b *Bool) UnmarshalJSON(text []byte) (err error) {
	var bl bool
	err = json.Unmarshal(text, &bl)
	if err == nil {
		*b = Bool(bl)
		return nil
	}

	var str string
	err = json.Unmarshal(text, &str)
	if err == nil {
		return nil
	}

	if str == "" {
		return nil
	}

	if str == "F" {
		*b = false
	}

	return errors.New("FML")
}

type Decimal struct {
	decimal.Big
}

func (d Decimal) IsEmpty() bool {
	f, _ := d.Big.Float64()
	return f == 0.0
}

func (d Decimal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s, err := d.Big.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(s, start)
}
