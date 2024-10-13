package main

import (
	"encoding/json"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// SDL UTIL FUNCTIONS
//

func startSDL() (*sdl.Window, *sdl.Renderer) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, _ := sdl.CreateWindow(
		"Window Title",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		400, 300,
		sdl.WINDOW_SHOWN,
	)

	renderer, _ := sdl.CreateRenderer(
		window, -1, sdl.RENDERER_ACCELERATED,
	)

	return window, renderer
}

func stopSDL(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

// GAME STUFF
//

type Player struct {
	PosX float32
	PosY float32
}

// In game actions
const (
	ActionPlayerMoveUp = iota
	ActionPlayerMoveDown
	ActionPlayerMoveLeft
	ActionPlayerMoveRight
	ActionUIMenu
	ActionLast
)

type Keymap [ActionLast]sdl.Scancode // Map of actions to sdl scancodes

// Map to convert json config file values to scancodes
var KeyConversion map[string]sdl.Scancode = map[string]sdl.Scancode{
	"esc":       sdl.SCANCODE_ESCAPE,
	"tab":       sdl.SCANCODE_TAB,
	"caps":      sdl.SCANCODE_CAPSLOCK,
	"shift":     sdl.SCANCODE_LSHIFT,
	"r_shift":   sdl.SCANCODE_RSHIFT,
	"ctrl":      sdl.SCANCODE_LCTRL,
	"r_ctrl":    sdl.SCANCODE_RCTRL,
	"del":       sdl.SCANCODE_DELETE,
	"home":      sdl.SCANCODE_HOME,
	"page_up":   sdl.SCANCODE_PAGEUP,
	"page_down": sdl.SCANCODE_PAGEDOWN,
	"space":     sdl.SCANCODE_SPACE,
	"up":        sdl.SCANCODE_UP,
	"down":      sdl.SCANCODE_DOWN,
	"left":      sdl.SCANCODE_LEFT,
	"right":     sdl.SCANCODE_RIGHT,
	"a":         sdl.SCANCODE_A,
	"b":         sdl.SCANCODE_B,
	"c":         sdl.SCANCODE_C,
	"d":         sdl.SCANCODE_D,
	"e":         sdl.SCANCODE_E,
	"f":         sdl.SCANCODE_F,
	"g":         sdl.SCANCODE_G,
	"h":         sdl.SCANCODE_H,
	"i":         sdl.SCANCODE_I,
	"j":         sdl.SCANCODE_J,
	"k":         sdl.SCANCODE_K,
	"l":         sdl.SCANCODE_L,
	"m":         sdl.SCANCODE_M,
	"n":         sdl.SCANCODE_N,
	"o":         sdl.SCANCODE_O,
	"p":         sdl.SCANCODE_P,
	"q":         sdl.SCANCODE_Q,
	"r":         sdl.SCANCODE_R,
	"s":         sdl.SCANCODE_S,
	"t":         sdl.SCANCODE_T,
	"u":         sdl.SCANCODE_U,
	"v":         sdl.SCANCODE_V,
	"w":         sdl.SCANCODE_W,
	"x":         sdl.SCANCODE_X,
	"y":         sdl.SCANCODE_Y,
	"z":         sdl.SCANCODE_Z,
	"0":         sdl.SCANCODE_0,
	"1":         sdl.SCANCODE_1,
	"2":         sdl.SCANCODE_2,
	"3":         sdl.SCANCODE_3,
	"4":         sdl.SCANCODE_4,
	"5":         sdl.SCANCODE_5,
	"6":         sdl.SCANCODE_6,
	"7":         sdl.SCANCODE_7,
	"8":         sdl.SCANCODE_8,
	"9":         sdl.SCANCODE_9,
	"-":         sdl.SCANCODE_MINUS,
	"=":         sdl.SCANCODE_EQUALS,
	"[":         sdl.SCANCODE_LEFTBRACKET,
	"]":         sdl.SCANCODE_RIGHTBRACKET,
	"\\":        sdl.SCANCODE_BACKSLASH,
	";":         sdl.SCANCODE_SEMICOLON,
	"'":         sdl.SCANCODE_APOSTROPHE,
	",":         sdl.SCANCODE_COMMA,
	".":         sdl.SCANCODE_PERIOD,
	"/":         sdl.SCANCODE_SLASH,
}

// Function for loading controls from config.json
func loadControls() Keymap {
	var keymap Keymap
	configFile, _ := os.ReadFile("config.json")
	jsonData := make(map[string]string)

	json.Unmarshal(configFile, &jsonData)

	keymap[ActionPlayerMoveUp] = KeyConversion[jsonData["PlayerMoveUp"]]
	keymap[ActionPlayerMoveDown] = KeyConversion[jsonData["PlayerMoveDown"]]
	keymap[ActionPlayerMoveLeft] = KeyConversion[jsonData["PlayerMoveLeft"]]
	keymap[ActionPlayerMoveRight] = KeyConversion[jsonData["PlayerMoveRight"]]
	keymap[ActionUIMenu] = KeyConversion[jsonData["UIMenu"]]

	return keymap
}

// MAIN FUNC
//

func main() {
	window, renderer := startSDL()
	defer stopSDL(window, renderer)

	player := Player{PosX: 200, PosY: 150}
	keymap := loadControls()

	// Main loop
	for active := true; active; {
		//active = handleInputA(&player)
		active = handleInputB(keymap, &player)

		// Clear screen
		renderer.SetDrawColor(0x00, 0x00, 0x1a, 0xff)
		renderer.Clear()

		// Draw Player
		renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
		renderer.FillRectF(
			&sdl.FRect{
				X: player.PosX - 4,
				Y: player.PosY - 4,
				W: 8,
				H: 8,
			},
		)

		// Present
		renderer.Present()
	}
}

// Default input handling function
func handleInputA(player *Player) bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			switch e.Keysym.Scancode {
			case sdl.SCANCODE_W:
				player.PosY -= 4
			case sdl.SCANCODE_S:
				player.PosY += 4
			case sdl.SCANCODE_A:
				player.PosX -= 4
			case sdl.SCANCODE_D:
				player.PosX += 4
			case sdl.SCANCODE_ESCAPE:
				return false
			}
		}
	}

	// Keep main loop running
	return true
}

// Keymap based input handling function
func handleInputB(keymap Keymap, player *Player) bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			switch e.Keysym.Scancode {
			case keymap[ActionPlayerMoveUp]:
				player.PosY -= 4
			case keymap[ActionPlayerMoveDown]:
				player.PosY += 4
			case keymap[ActionPlayerMoveLeft]:
				player.PosX -= 4
			case keymap[ActionPlayerMoveRight]:
				player.PosX += 4
			case keymap[ActionUIMenu]:
				return false
			}
		}
	}

	// Keep main loop running
	return true
}
