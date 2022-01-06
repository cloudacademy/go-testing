package rocket

import (
	"errors"
)

type EngineConfig struct {
	Count   int
	Type    string
	Ignited bool
}

type Rocket struct {
	Name         string
	Manufacturer string
	Engines      EngineConfig
	Maxspeed     int
	Fuel         int
	Speed        int
}

func NewRocket(name string, manufacturer string, engines int, maxSpeed int, fuel int) *Rocket {
	engConf := EngineConfig{
		Count:   engines,
		Ignited: false,
	}

	return &Rocket{
		Name:         name,
		Manufacturer: manufacturer,
		Engines:      engConf,
		Maxspeed:     maxSpeed,
		Fuel:         fuel,
		Speed:        0,
	}
}

func (r *Rocket) Ignite() error {
	if r.Fuel == 0 {
		return errors.New("empty tank")
	}
	r.Engines.Ignited = true
	return nil
}

func (r *Rocket) ThrottleUp(amount int) (int, error) {
	if r.Engines.Ignited {
		if r.Speed+amount <= r.Maxspeed {
			r.Speed = r.Speed + amount
			return r.Speed, nil
		} else {
			return r.Speed, errors.New("exceeds max speed")
		}
	} else {
		return r.Speed, errors.New("engines need to be ignited first")
	}
}

func (r *Rocket) ThrottleDown(amount int) (int, error) {
	if r.Engines.Ignited {
		if r.Speed-amount <= 1000 {
			return r.Speed, errors.New("new speed too low")
		} else {
			r.Speed = r.Speed - amount
			return r.Speed, nil
		}
	} else {
		return r.Speed, errors.New("engines need to be ignited first")
	}
}

func (r *Rocket) CurrentSpeed() int {
	return r.Speed
}
