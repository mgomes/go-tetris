# Go Tetris - Outrun Edition

A Tetris clone written in Go with an Outrun/synthwave aesthetic, featuring neon colors, 3D-styled blocks, and a retro-futuristic visual design.

## Features

- Classic Tetris gameplay with all 7 tetromino pieces
- Outrun-themed visuals with neon colors and synthwave grid background
- 3D-styled blocks with depth and glow effects
- Hold piece functionality
- Ghost piece preview
- Progressive speed increase using Tetris Worlds speed curve
- Perfect clear bonuses and back-to-back Tetris scoring
- Pause functionality
- Score and level tracking with 7-segment style displays

## Prerequisites

- Go 1.18 or higher
- C compiler (for CGO dependencies)
- OpenGL 2.1 support

### Platform-specific requirements

**macOS:**
- Xcode Command Line Tools

**Linux:**
- OpenGL development headers
- X11 development headers
```bash
# Ubuntu/Debian
sudo apt-get install libgl1-mesa-dev xorg-dev

# Fedora
sudo dnf install mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel
```

**Windows:**
- MinGW-w64 or Microsoft C++ Build Tools

## Building

1. Clone the repository:
```bash
git clone https://github.com/mgomes/go-tetris.git
cd go-tetris
```

2. Download dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build
```

## Running

Run the compiled binary:
```bash
./go-tetris
```

Or run directly with Go:
```bash
go run .
```

## Controls

- **Left/Right/Down Arrow** - Move piece left/right/down
- **Up Arrow** - Rotate piece clockwise
- **Shift** - Rotate piece counter-clockwise
- **Space** - Drop piece immediately
- **Left Ctrl** - Hold piece
- **P** - Pause/unpause game
- **R** - Start new game
- **Escape** - Quit

## Scoring System

### Normal Line Clears
- Single: 100 × level
- Double: 300 × level
- Triple: 500 × level
- Tetris: 800 × level

### Perfect Clear Bonuses
- Single-line perfect clear: 800 × level
- Double-line perfect clear: 1200 × level
- Triple-line perfect clear: 1800 × level
- Tetris perfect clear: 2000 × level
- Back-to-back Tetris perfect clear: 3200 × level

## Game Mechanics

- **Levels**: Increase every 10 lines cleared
- **Speed**: Follows Tetris Worlds speed curve (levels 1-20)
- **Wall Kicks**: SRS-inspired rotation system allows pieces to rotate near walls
- **Hold**: Can hold one piece at a time, swaps with current piece

## Technical Details

- Written in Go
- Uses GLFW for window management and input
- OpenGL 2.1 for rendering
- No external game engine dependencies

## License

MIT License

## Acknowledgments

- Inspired by classic Tetris gameplay
- Visual design inspired by Outrun aesthetics and synthwave culture
- Speed curve based on Tetris Worlds standards