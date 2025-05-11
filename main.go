// The Amazing Race is a terminal-based racing simulation written in Go.
// It uses goroutines and channels to simulate a cast of 'racers' (represented
// by emojis), each with their own speed and crash chance, competing to reach
// the finish line.
//
// Racers move concurrently with randomized behavior, and each update is
// rendered in real-time. The simulation uses visual progress bars, crash
// penalties via stun mechanics, and end-of-race podium results.
//
// Author: Cameron MacKinnon
package main

import (
	"fmt"
	"sort"

	"github.com/Cameron-MacKinnon/CSP3341-amazing-race/race"
	"github.com/Cameron-MacKinnon/CSP3341-amazing-race/racer"
	"github.com/Cameron-MacKinnon/CSP3341-amazing-race/render"
)

func main() {
	// Load racers
	racers := racer.Presets

	// Create a controller and a renderer
	con := race.NewController(racers)
	renderer := render.NewRenderer(con.Distance)

	// Keep track of how many goroutines are complete, and each racer's final
	// race update
	totalRacers := len(con.Racers)
	numFinished := 0
	var finishers []racer.RaceUpdate

	// Holds the most recent update for each racer
	state := make(map[int]racer.RaceUpdate)

	// Clear screen incase anything is left over from a previous run
	fmt.Print("\033[H\033[2J")

	// Start the race (each racer runs in its own goroutine)
	go con.StartRacers()

	// Listen for state updates from racers
	for update := range con.Updates {
		// Save the latest update for this racer
		state[update.RacerID] = update

		// Redraw race with updated state
		renderer.Draw(state, racers)

		// Handle finished racers
		if update.Finished {
			numFinished++
			update.FinishPlace = numFinished
			finishers = append(finishers, update)

			// Once all racers finish, close the update channel to end the loop
			if numFinished == totalRacers {
				close(con.Updates)
			}
		}
	}

	// Sort the finishers list by finish order
	sort.Slice(finishers, func(i, j int) bool {
		return finishers[i].FinishPlace < finishers[j].FinishPlace
	})

	// Announce the first 3 finishers
	fmt.Println("\n🏁 Final Standings:")
	for i, finisher := range finishers {
		if i >= 3 {
			break
		}
		r := racers[finisher.RacerID-1]
		place := []string{"🥇 1st", "🥈 2nd", "🥉 3rd"}[i]
		fmt.Printf("%s — %s\t%s\n", place, r.Emoji, r.Name)
	}
	fmt.Println()
}
