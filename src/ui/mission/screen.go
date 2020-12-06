// Package mission contains the screen that displays the world map to the user. The user can then use the controller to
// select his next mission.
package mission

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/util"
)

const crossHairImagePath = "./images/tiles/other/crosshair.png"

var locationMarkerColor = common.ParseHexColor("#fea400")

const locationMarkerRadius = 10
const worldMapImagePath = "./images/tiles/other/world-map.jpg"
const worldMapWidth = 1280
const worldMapHeight = 783

type Screen struct {
	availableMissions    []*assets.Mission
	crossHairSprite      *pixel.Sprite
	inputController      input.Controller
	missionsInitialized  bool
	screenChangeRequired common.ScreenChangeCallback
	selectedMission      *assets.Mission
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
	worldMapSprite       *pixel.Sprite
}

func (s *Screen) SetInputController(inputCtrl input.Controller) {
	s.inputController = inputCtrl
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	if !assets.MissionRepository.Initialized() {
		assets.MissionRepository.Initialize()
	} else {
		s.initializeMissions()
	}
	s.crossHairSprite = common.LoadSprite(crossHairImagePath)
	s.worldMapSprite = common.LoadSprite(worldMapImagePath)
}

func (s *Screen) initializeMissions() {
	remainingMissions, err := engine.MissionController.RemainingMissions()
	if nil != err {
		logging.Error.Fatalf("Failed to retrieve list of remaining missions: %v", err)
	}
	if 0 == len(remainingMissions) {
		logging.Error.Fatalf("List of remaining missions is empty. Game should have ended!")
	}

	s.availableMissions = remainingMissions
	s.selectedMission = remainingMissions[0]
}

func (s *Screen) Update(_ int64) {
	if !s.missionsInitialized && assets.MissionRepository.Initialized() {
		s.missionsInitialized = true
		s.initializeMissions()
	}

	s.drawWorldMap()

	if s.missionsInitialized {
		s.processUserInput()
		s.drawMissionLocations()
	}

	// TODO: Draw image of client
	// TODO: Draw mission briefing
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.Mission)
}

func (s *Screen) drawWorldMap() {
	var factor = s.getWorldMapScalingFactor()
	var mapCenter = s.getWorldMapCenter()
	s.worldMapSprite.Draw(s.window, pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(mapCenter))
}

func (s *Screen) drawMissionLocations() {
	var factor = s.getWorldMapScalingFactor()
	var mapCenter = s.getWorldMapCenter()
	var mapBottomLeft = pixel.Vec{
		X: mapCenter.X - (worldMapWidth*factor)/2,
		Y: mapCenter.Y - (worldMapHeight*factor)/2,
	}

	for _, city := range s.availableMissions {
		var cityLocation = pixel.Vec{
			X: mapBottomLeft.X + (city.Location.Longitude * factor),
			Y: mapBottomLeft.Y + (worldMapHeight * factor) - (city.Location.Latitude * factor),
		}
		s.drawLocationMarker(cityLocation)
		if city.Name == s.selectedMission.Name {
			s.crossHairSprite.Draw(s.window, pixel.IM.Moved(cityLocation))
		}
	}
}

func (s *Screen) getWorldMapScalingFactor() float64 {
	var factorX = (s.window.Bounds().Max.X - 100) / s.worldMapSprite.Picture().Bounds().Max.X
	var factorY = (s.window.Bounds().Max.Y * 3 / 4) / s.worldMapSprite.Picture().Bounds().Max.Y
	return util.Min(factorX, factorY)
}

func (s *Screen) getWorldMapCenter() (result pixel.Vec) {
	result = s.window.Bounds().Center()
	result.Y -= s.window.Bounds().Max.Y / 8
	return
}

func (s *Screen) drawLocationMarker(location pixel.Vec) {
	imd := imdraw.New(nil)
	imd.Color = locationMarkerColor
	imd.Push(location)
	imd.Circle(locationMarkerRadius, 5)
	imd.Draw(s.window)
}

func (s *Screen) processUserInput() {
	var uiEventState = s.inputController.GetControllerUiEventStateCombined()
	if nil != uiEventState {
		if uiEventState.PressedButton {
			s.screenChangeRequired(common.BuyYourWeapons)
		} else {
			var nextMission = s.selectedMission
			var err error = nil
			if uiEventState.MovedUp {
				nextMission, err = engine.MissionController.NextMissionNorth(&s.selectedMission.Location)
			} else if uiEventState.MovedDown {
				nextMission, err = engine.MissionController.NextMissionSouth(&s.selectedMission.Location)
			} else if uiEventState.MovedLeft {
				nextMission, err = engine.MissionController.NextMissionWest(&s.selectedMission.Location)
			} else if uiEventState.MovedRight {
				nextMission, err = engine.MissionController.NextMissionEast(&s.selectedMission.Location)
			}
			if nil != err {
				logging.Error.Fatalf("Failed to get next mission north: %v", err)
			}
			if nil != nextMission {
				s.selectedMission = nextMission
			}
		}
	}
}
