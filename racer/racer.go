package racer

import (
	"math/rand/v2"
	"time"
)

// Racer represents a participant in the race
type Racer struct {
	ID          int     // Unique racer ID
	Name        string  // Racer's name
	Emoji       string  // Visual representation of the racer
	Speed       int     // Distance moved per tick
	CrashChance float64 // Probability of crashing on a tick
	Position    int     // Current position on the track
	Stun        int     // Remaining turns to skip due to crashing
}

// RaceUpdate is a message sent to the controller via the Updates channel
type RaceUpdate struct {
	RacerID     int  // ID of the racer
	Position    int  // Current position on the track
	Crashed     bool // Whether the racer crashed this tick
	Finished    bool // Whether the racer has finished the race
	FinishPlace int  // Final placement (1st, 2nd, etc.)
}

// Move runs in a loop, determining a racer's behavior on each tick. The racer
// either crashes, is stunned, finishes, or moves forward.
func (r *Racer) Move(trackLength int, tickSpeed int, updates chan RaceUpdate) {
	for {

		// Wait for the next tick
		time.Sleep(time.Duration(tickSpeed) * time.Millisecond)

		// If stunned, skip this turn and decrement stun counter
		if r.Stun > 0 {
			r.Stun--
			updates <- RaceUpdate{
				RacerID:  r.ID,
				Position: r.Position,
				Crashed:  true,
				Finished: false,
			}
			continue
		}

		// Roll for crash; apply stun penalty if it occurs
		if rand.Float64() < r.CrashChance {
			r.Stun = rand.IntN(3) + 1 // Stun for 1–3 turns
			updates <- RaceUpdate{
				RacerID:  r.ID,
				Position: r.Position,
				Crashed:  true,
				Finished: false,
			}
			continue
		}

		// Advance position
		r.Position += r.Speed

		// Check for finish
		if r.Position >= trackLength {
			r.Position = trackLength
			updates <- RaceUpdate{
				RacerID:  r.ID,
				Position: r.Position,
				Crashed:  false,
				Finished: true,
			}
			return
		}

		// Normal progress update
		updates <- RaceUpdate{
			RacerID:  r.ID,
			Position: r.Position,
			Crashed:  false,
			Finished: false,
		}
	}
}
