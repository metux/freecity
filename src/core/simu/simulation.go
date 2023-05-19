package simu

import (
    "log"
    "github.com/metux/freecity/core/items"
)

type Simulator struct {
    terrain * items.TerrainMap
    notify items.NotifyHandler
}

func (sim * Simulator) Init(tm * items.TerrainMap, n items.NotifyHandler) {
    log.Println("init simu")
    sim.terrain = tm
    sim.notify = n
}

func (sim * Simulator) SetNotify(nh items.NotifyHandler) items.NotifyHandler {
    old := sim.notify
    sim.notify = nh
    return old
}

func (sim * Simulator) Tick() {
    sim.terrain.Date.NextHour()
    sim.simNextHour()

    if (sim.terrain.Date.StartOfYear()) {
        sim.simNextYear()
    }

    if (sim.terrain.Date.StartOfMonth()) {
        sim.simNextMonth()
    }

    if (sim.terrain.Date.StartOfDay()) {
        sim.simNextDay()
    }
}

func (sim * Simulator) simNextHour() {
    sim.notify.NotifyEmit(act, NotifySimuNextHour{sim.terrain.Date})
}

func (sim * Simulator) simNextDay() {
    sim.notify.NotifyEmit(act, NotifySimuNextDay{sim.terrain.Date})
}

func (sim * Simulator) simNextMonth() {
    sim.notify.NotifyEmit(act, NotifySimuNextMonth{sim.terrain.Date})
}

func (sim * Simulator) simNextYear() {
    sim.notify.NotifyEmit(act, NotifySimuNextYear{sim.terrain.Date})
}
