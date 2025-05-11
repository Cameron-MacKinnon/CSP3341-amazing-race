# 🏁 The Amazing Race
A terminal-based racing simulator written in Go! Emoji-powered racers with 
unique stats compete to reach the finish line. Each turn carries the possibility
of crashes and stuns to randomise the outcome each time the program is run.

---

## 💡 Features

- ⚙️ Built to highlight goroutines, channels, and idiomatic formatting
- 🔧 Standard lib only, no external dependencies
- 🎭 Each racer has distinct speed and crash stats
- 🛩️ Includes fast, crash-prone racers
- 🐢 Includes slow and steady racers who can still win (albeit extremely rarely...)
- 🐖 Includes middle-of-the pack racers with balanced stats.
- 💥 Crash mechanic includes random stun duration (1–3 turns)
- 📊 Real-time rendering with progress bars and live updates as messages are consumed from the updates channel.
- 🥇 Podium display for the top 3 finishers


---

## 📦 Project Structure
```
.
├── main.go                 // Entry point: sets up and runs the race
├── race/
│   └── race.go             // Race controller: manages racer goroutines and update channel
├── racer/
│   ├── racer.go            // Racer logic: movement, crash/stun handling, message sending
│   └── presets.go          // Predefined racer roster with unique stats and emojis
├── render/
│   └── render.go           // Renderer: draws progress bars and race state in terminal
```

---

## 🚀 Running the Project

```bash
go run main.go
```
---

## 🧑‍💻 Unit & Author Details

**Unit:** CSP3341 — Programming Languages and Paradigms  
**Project:** Assignment 2 - Technical Report and Application
**Author:** Cameron MacKinnon  
**Institution:** Edith Cowan University, Joondalup  
**Semester:** Semester 2, 2025  
**Unit Coordinator:** Dr Ali Hur