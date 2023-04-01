package simu

import (
    "fmt"
    "github.com/metux/freecity/core/base"
)

type NotifySimuNextHour struct {
    Date base.Date
}

func (n NotifySimuNextHour) String() string {
    return fmt.Sprintf("next hour: %s", n.Date)
}

type NotifySimuNextDay struct {
    Date base.Date
}

func (n NotifySimuNextDay) String() string {
    return fmt.Sprintf("next day: %s", n.Date)
}

type NotifySimuNextMonth struct {
    Date base.Date
}

func (n NotifySimuNextMonth) String() string {
    return fmt.Sprintf("next month: %s", n.Date)
}

type NotifySimuNextYear struct {
    Date base.Date
}

func (n NotifySimuNextYear) String() string {
    return fmt.Sprintf("next year: %s", n.Date)
}
