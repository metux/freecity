package base

import (
    "time"
    "gopkg.in/yaml.v3"
)

type Month = time.Month

const DateTime = "2006-01-02 15:04:05"

// FIXME: very naive implementation, should use time package internally
type Date struct {
    ts time.Time
}

func (d Date) MarshalYAML() (interface{}, error) {
    return d.ts.Format(DateTime), nil
}

func (d *Date) UnmarshalYAML(value *yaml.Node) error {
    var tmpStr string

    if err := value.Decode(&tmpStr); err != nil {
        return err
    }

    tt,err := time.Parse(DateTime, tmpStr)
    if err == nil {
        d.ts = tt
        return nil
    }

    return err
}

func (d *Date) NextHour() {
    d.ts = d.ts.Add(time.Hour)
}

func (d *Date) NextDay() {
    d.ts = d.ts.AddDate(0, 0, 1)
}

func (d Date) StartOfDay() bool {
    return d.ts.Hour() == 0
}

func (d Date) StartOfMonth() bool {
    return d.ts.Day() == 1 && d.StartOfDay()
}

func (d Date) StartOfYear() bool {
    return (d.ts.Month() == time.January) && (d.ts.Day() == 1) && d.StartOfMonth()
}

func (d Date) Date() (year int, month Month, day int) {
    return d.ts.Date()
}

func (d Date) String() string {
    return d.ts.Format(DateTime)
}
