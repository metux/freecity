package simu

import (
    "fmt"
)

type NotifySimuNextHour struct {
    Date date
}

func (n NotifySimuNextHour) String() string {
    return fmt.Sprintf("next hour: %s", n.Date)
}

type NotifySimuNextDay struct {
    Date date
}

func (n NotifySimuNextDay) String() string {
    return fmt.Sprintf("next day: %s", n.Date)
}

type NotifySimuNextMonth struct {
    Date date
}

func (n NotifySimuNextMonth) String() string {
    return fmt.Sprintf("next month: %s", n.Date)
}

type NotifySimuNextYear struct {
    Date date
}

func (n NotifySimuNextYear) String() string {
    return fmt.Sprintf("next year: %s", n.Date)
}
