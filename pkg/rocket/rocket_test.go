package rocket

import (
	"testing"
)

func TestNewRocket(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	got := rocket.Name
	want := "Saturn-V"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestIgniteRocket(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	//nolint:errcheck
	rocket.Ignite()
	got := rocket.Engines.Ignited
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIgniteRocketEmpty(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)

	err := rocket.Ignite()
	if err == nil {
		t.Errorf("Expected error when attempting to ignite rocket with 0 fuel")
	}
}

func TestThrottleUpEnginesNotIgnited(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	_, err := rocket.ThrottleUp(100)

	if err == nil {
		t.Errorf("Expected error when attempting to throttle up engines which are not ignited first")
	}
}

//nolint:errcheck
func TestThrottleUp(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	rocket.Ignite()

	rocket.ThrottleUp(100)
	rocket.ThrottleUp(100)

	got, _ := rocket.ThrottleUp(100)
	want := 300

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

//nolint:errcheck
func TestThrottleUpExceedMaxSpped(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	rocket.Ignite()

	rocket.ThrottleUp(25000)
	_, err := rocket.ThrottleUp(100)

	if err == nil {
		t.Errorf("Expected error when attempting to throttle up over max speed")
	}
}

func TestThrottleDownEnginesNotIgnited(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	_, err := rocket.ThrottleDown(100)

	if err == nil {
		t.Errorf("Expected error when attempting to throttle down engines which are not ignited first")
	}
}

//nolint:errcheck
func TestThrottleDown(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	rocket.Ignite()

	rocket.ThrottleUp(1000)
	rocket.ThrottleUp(100)
	rocket.ThrottleUp(100)
	rocket.ThrottleDown(10)
	got, _ := rocket.ThrottleDown(10)
	want := 1180

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

//nolint:errcheck
func TestThrottleDownToUnder1000(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	rocket.Ignite()

	rocket.ThrottleUp(900)
	_, err := rocket.ThrottleDown(10)

	if err == nil {
		t.Errorf("Expected error when attempting to throttle down with resulting speed less than 1000")
	}
}

//nolint:errcheck
func TestCurrentSpeed(t *testing.T) {
	rocket := NewRocket("Saturn-V", "Nasa", 5, 25000)
	rocket.AddFuel(100)

	rocket.Ignite()

	rocket.ThrottleUp(1000)
	rocket.ThrottleUp(100)
	rocket.ThrottleDown(10)
	rocket.ThrottleUp(100)
	rocket.ThrottleDown(5)

	got := rocket.CurrentSpeed()
	want := 1185

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
