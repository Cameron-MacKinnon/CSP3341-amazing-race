package render

import (
	"fmt"
	"strings"

	"github.com/Cameron-MacKinnon/CSP3341-amazing-race/racer"
)

const barWidth int = 100 // Width of the progress bars

// Renderer handles visual display of the race
type Renderer struct {
	TrackLength int
}

// NewRenderer creates a new Renderer with the given track length
func NewRenderer(trackLength int) *Renderer {
	return &Renderer{
		TrackLength: trackLength,
	}
}

// trackEdge visually marks the top and bottom of the track
// (e.g., "=========")
var trackEdge string = fmt.Sprint(strings.Repeat("=", barWidth+26) + "\n")

// buildBar returns a visual progress bar based on racer position
// (e.g., "▓▓▓▓▓░░░░░░░░░░")
func buildBar(pos, total int) string {
	progress := int(float64(pos) / float64(total) * float64(barWidth))
	return strings.Repeat("▓", progress) + strings.Repeat("░", barWidth-progress)
}

// Draw renders the current race state to the terminal. Each racer's emoji,
// name, progress bar, and completion percentage are displayed.
func (r *Renderer) Draw(state map[int]racer.RaceUpdate, racers []racer.Racer) {
	// Clear the screen using ANSI escape sequences
	fmt.Print("\033[H\033[2J")

	// Top border
	fmt.Println(trackEdge)

	// For each racer, get their current state (as per their most recent update)
	// and use it to build and display their progress bar.
	for _, rc := range racers {
		update := state[rc.ID]
		bar := buildBar(update.Position, r.TrackLength)
		percent := int(float64(update.Position) / float64(r.TrackLength) * 100)
		fmt.Printf(
			"%-2s\t%-8s: %-20s  (%3d%%)\n\n",
			rc.Emoji, rc.Name, bar, percent,
		)
	}

	// Bottom border
	fmt.Println(trackEdge)
}
