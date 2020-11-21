// Package start contains the second screen shown to the user. The screen displays a copyright notice and the dedication
// lines. Once the screen has been loaded, the theme song gets buffered. The next screen gets displayed when the theme
// song has been fully buffered.
package start

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"image/color"
	"retro-carnage.net/assets"
	"retro-carnage.net/engine/geometry"
	"retro-carnage.net/engine/input"
	"retro-carnage.net/ui/common"
)

const txtFirstLine = "RETRO CARNAGE"
const txtSecondLine = "(C) 2020 THOMAS WERNER"
const txtThirdLine = "Dedicated to Jonathan Werner"
const txtFourthLine = "Inspired by 'DOGS OF WAR'"
const txtFifthLine = "(C) 1989 by Elite Systems Ltd."

type Screen struct {
	screenChangeRequired common.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *common.Stereo
	textDimensions       map[string]*geometry.Point
	themeLoaded          bool
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(_ *input.Controller) {}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.screenChangeTimeout = 0
	s.textDimensions = common.GetTextDimensions(text.New(pixel.V(0, 0), common.DefaultAtlas),
		txtFirstLine, txtSecondLine, txtThirdLine, txtFourthLine, txtFifthLine)

	s.stereo = common.NewStereo()
	s.themeLoaded = false
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs
	if s.themeLoaded {
		s.screenChangeRequired(common.Title)
	}
	s.renderScreen()
	if !s.themeLoaded && (s.screenChangeTimeout > 100) {
		// TODO: Buffer theme song asynchronously
		// This next call will buffer the song if it's not buffered already. That blocks the main thread for a couple of
		// seconds. It would be much cooler to start the buffer process asynchronously. Access to the map used to store
		// the songs could be protected by a Mutex. Then the buffering would be triggered once. Then we'd check if the
		// buffering has finished by accessing the map to see if the song it there.
		s.stereo.PlaySong(assets.ThemeSong)
		s.themeLoaded = true
	}
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.Start)
}

func (s *Screen) renderScreen() {
	s.drawLineToScreen(txtFirstLine, 2.5, common.Red)
	s.drawLineToScreen(txtSecondLine, 1, common.Yellow)
	s.drawLineToScreen(txtThirdLine, -1, common.Green)
	s.drawLineToScreen(txtFourthLine, -2.5, common.Green)
	s.drawLineToScreen(txtFifthLine, -4, common.Green)
}

func (s *Screen) drawLineToScreen(line string, offsetMultiplier float64, color color.Color) {
	var vertCenter = s.window.Bounds().Max.Y / 2

	var lineDimensions = s.textDimensions[line]
	var lineX = (s.window.Bounds().Max.X - lineDimensions.X) / 2
	var lineY = vertCenter + offsetMultiplier*lineDimensions.Y

	var txt = text.New(pixel.V(lineX, lineY), common.DefaultAtlas)
	txt.Color = color
	_, _ = fmt.Fprint(txt, line)
	txt.Draw(s.window, pixel.IM)
}
