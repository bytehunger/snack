package main

import (
	"strings"
	"time"
)

const w3cDateFormat = "2006-01-02"

// W3CDate allows to parse and render dates in a YYYY-MM-DD format.
type W3CDate time.Time

func (d W3CDate) String() string {
	t := time.Time(d)

	if t.IsZero() {
		return ""
	}

	return t.Format(w3cDateFormat)
}

func (d W3CDate) IsDate() bool {
	t := time.Time(d)
	return !t.IsZero()
}

func (d *W3CDate) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *W3CDate) MarshalYAML() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *W3CDate) UnmarshalJSON(b []byte) error {
	return d.Unmarshal(string(b))
}

func (d *W3CDate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)

	if err != nil {
		return err
	}

	return d.Unmarshal(s)
}

func (d *W3CDate) Unmarshal(s string) error {
	if s == "" {
		return nil
	}

	t, err := time.Parse(w3cDateFormat, strings.Trim(s, `"`))

	if err != nil {
		return err
	}

	*d = W3CDate(t)

	return nil
}
