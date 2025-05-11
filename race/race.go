package race

import (
	"github.com/Cameron-MacKinnon/CSP3341-amazing-race/racer"
)

const distance int = 100  // Total race track length (units)
const tickSpeed int = 200 // Tick interval in milliseconds

// Controller manages the race state and racer goroutines
type Controller struct {
	Racers    []racer.Racer         // All racers participating in the race
	Distance  int                   // Length of the race track
	TickSpeed int                   // Delay between updates (ms)
	Updates   chan racer.RaceUpdate // Channel for racer updates
}

// NewController creates and configures a new race controller
func NewController(racers []racer.Racer) Controller {
	return Controller{
		Racers:    racers,
		Distance:  distance,
		TickSpeed: tickSpeed,
		Updates:   make(chan racer.RaceUpdate),
	}
}

// StartRacers launches a goroutine for each racer
func (c *Controller) StartRacers() {
	for _, r := range c.Racers {
		go r.Move(c.Distance, c.TickSpeed, c.Updates)
	}
}
