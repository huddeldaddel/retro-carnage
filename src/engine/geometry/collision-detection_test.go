package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//--- Up -------------------------------------------------------------------------------------------------------------//

func TestShouldFindCollisionForRectMovingUpAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 5, Y: 10, Width: 5, Height: 2}
	var stillRect = &Rectangle{X: 2, Y: 5, Width: 25, Height: 2}
	var collision = getCollisionForMovementUp(movingRect, stillRect, &Point{X: 0, Y: -5})

	assert.NotNil(t, collision)
	assert.InDelta(t, 7.0, collision.Y, 0.00001)
}

func TestShouldFindCollisionForRectMovingUpAgainstAnotherSmallRectLeft(t *testing.T) {
	var movingRect = &Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 1, Y: 1, Width: 2, Height: 2}
	var collision = getCollisionForMovementUp(movingRect, stillRect, &Point{X: 0, Y: -4})

	assert.NotNil(t, collision)
	assert.InDelta(t, 3.0, collision.Y, 0.00001)
}

func TestShouldFindCollisionForRectMovingUpAgainstAnotherSmallRectRight(t *testing.T) {
	var movingRect = &Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 3, Y: 1, Width: 2, Height: 2}
	var collision = getCollisionForMovementUp(movingRect, stillRect, &Point{X: 0, Y: -4})

	assert.NotNil(t, collision)
	assert.InDelta(t, 3.0, collision.Y, 0.00001)
}

func TestShouldStopTheUpMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 5, Y: 10, Width: 5, Height: 2}
	var stillRect = &Rectangle{X: 2, Y: 5, Width: 25, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Up, &Point{X: 0, Y: -5})

	assert.NotNil(t, result)
	assert.InDelta(t, 5.0, result.X, 0.00001)
	assert.InDelta(t, 7.0, result.Y, 0.00001)
	assert.InDelta(t, 5.0, result.Width, 0.00001)
	assert.InDelta(t, 2.0, result.Height, 0.00001)
}

func TestShouldStopTheUpMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 5, Y: 10, Width: 5, Height: 2}
	var stillRect = &Rectangle{X: 7, Y: 5, Width: 1, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Up, &Point{X: 0, Y: -6})

	assert.NotNil(t, result)
	assert.InDelta(t, 5.0, result.X, 0.00001)
	assert.InDelta(t, 7.0, result.Y, 0.00001)
	assert.InDelta(t, 5.0, result.Width, 0.00001)
	assert.InDelta(t, 2.0, result.Height, 0.00001)
}

//--- Down -----------------------------------------------------------------------------------------------------------//

func TestShouldFindCollisionForRectMovingDownAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 6, Y: 3, Width: 2, Height: 1}
	var stillRect = &Rectangle{X: 2, Y: 8, Width: 25, Height: 2}
	var collision = getCollisionForMovementDown(movingRect, stillRect, &Point{X: 0, Y: 5})

	assert.NotNil(t, collision)
	assert.InDelta(t, 8.0, collision.Y, 0.00001)
}

func TestShouldFindCollisionForRectMovingDownAgainstAnotherRectLeft(t *testing.T) {
	var movingRect = &Rectangle{X: 3, Y: 1, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var collision = getCollisionForMovementDown(movingRect, stillRect, &Point{X: 0, Y: 4})

	assert.NotNil(t, collision)
	assert.InDelta(t, 4.0, collision.Y, 0.00001)
}

func TestShouldFindCollisionForRectMovingDownAgainstAnotherRectRight(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 1, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var collision = getCollisionForMovementDown(movingRect, stillRect, &Point{X: 0, Y: 4})

	assert.NotNil(t, collision)
	assert.InDelta(t, 4.0, collision.Y, 0.00001)
}

func TestShouldStopTheDownMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 1, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Down, &Point{X: 0, Y: 4})

	assert.NotNil(t, result)
	assert.InDelta(t, 1.0, result.X, 0.00001)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
	assert.InDelta(t, 2.0, result.Width, 0.00001)
	assert.InDelta(t, 2.0, result.Height, 0.00001)
}

func TestShouldStopTheDownMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 1, Width: 4, Height: 2}
	var stillRect = &Rectangle{X: 2, Y: 4, Width: 2, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Down, &Point{X: 0, Y: 5})

	assert.NotNil(t, result)
	assert.InDelta(t, 1.0, result.X, 0.00001)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
	assert.InDelta(t, 4.0, result.Width, 0.00001)
	assert.InDelta(t, 2.0, result.Height, 0.00001)
}

//--- Left -----------------------------------------------------------------------------------------------------------//

func TestShouldFindCollisionForRectMovingLeftAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 5, Y: 3, Width: 3, Height: 3}
	var stillRect = &Rectangle{X: 1, Y: 1, Width: 1, Height: 9}
	var collision = getCollisionForMovementLeft(movingRect, stillRect, &Point{X: -5, Y: 0})

	assert.NotNil(t, collision)
	assert.InDelta(t, 2.0, collision.X, 0.00001)
}

func TestShouldFindCollisionForRectMovingLeftAgainstAnotherRectTop(t *testing.T) {
	var movingRect = &Rectangle{X: 3, Y: 3, Width: 1, Height: 3}
	var stillRect = &Rectangle{X: 1, Y: 1, Width: 1, Height: 3}
	var collision = getCollisionForMovementLeft(movingRect, stillRect, &Point{X: -5, Y: 0})

	assert.NotNil(t, collision)
	assert.InDelta(t, 2.0, collision.X, 0.00001)
}

func TestShouldFindCollisionForRectMovingLeftAgainstAnotherRectBottom(t *testing.T) {
	var movingRect = &Rectangle{X: 3, Y: 3, Width: 1, Height: 3}
	var stillRect = &Rectangle{X: 1, Y: 5, Width: 1, Height: 3}
	var collision = getCollisionForMovementLeft(movingRect, stillRect, &Point{X: -5, Y: 0})

	assert.NotNil(t, collision)
	assert.InDelta(t, 2.0, collision.X, 0.00001)
}

func TestShouldStopTheLeftMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 5, Y: 3, Width: 3, Height: 3}
	var stillRect = &Rectangle{X: 1, Y: 1, Width: 1, Height: 9}
	var result = StopMovementOnCollision(movingRect, stillRect, Left, &Point{X: -5, Y: 0})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.00001)
	assert.InDelta(t, 3.0, result.Y, 0.00001)
	assert.InDelta(t, 3.0, result.Width, 0.00001)
	assert.InDelta(t, 3.0, result.Height, 0.00001)
}

func TestShouldStopTheLeftMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 5, Y: 3, Width: 3, Height: 3}
	var stillRect = &Rectangle{X: 1, Y: 4, Width: 1, Height: 1}
	var result = StopMovementOnCollision(movingRect, stillRect, Left, &Point{X: -6, Y: 0})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.00001)
	assert.InDelta(t, 3.0, result.Y, 0.00001)
	assert.InDelta(t, 3.0, result.Width, 0.00001)
	assert.InDelta(t, 3.0, result.Height, 0.00001)
}

//--- Right ----------------------------------------------------------------------------------------------------------//

func TestShouldFindCollisionForRectMovingRightAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 3, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 4, Y: 2, Width: 1, Height: 4}
	var collision = getCollisionForMovementRight(movingRect, stillRect, &Point{X: 3, Y: 0})

	assert.NotNil(t, collision)
	assert.InDelta(t, 4.0, collision.X, 0.00001)
}

func TestShouldFindCollisionForRectMovingRightAgainstAnotherRectTop(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 3, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 4, Y: 2, Width: 1, Height: 2}
	var collision = getCollisionForMovementRight(movingRect, stillRect, &Point{X: 3, Y: 0})

	assert.NotNil(t, collision)
	assert.InDelta(t, 4.0, collision.X, 0.00001)
}

func TestShouldFindCollisionForRectMovingRightAgainstAnotherRectBottom(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 3, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 4, Y: 4, Width: 1, Height: 2}
	var collision = getCollisionForMovementRight(movingRect, stillRect, &Point{X: 3, Y: 0})

	assert.NotNil(t, collision)
	assert.InDelta(t, 4.0, collision.X, 0.00001)
}

func TestShouldStopTheRightMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 3, Width: 2, Height: 2}
	var stillRect = &Rectangle{X: 4, Y: 4, Width: 1, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Right, &Point{X: 3, Y: 0})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.00001)
	assert.InDelta(t, 3.0, result.Y, 0.00001)
	assert.InDelta(t, 2.0, result.Width, 0.00001)
	assert.InDelta(t, 2.0, result.Height, 0.00001)
}

func TestShouldStopTheRightMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 3, Width: 2, Height: 9}
	var stillRect = &Rectangle{X: 4, Y: 4, Width: 1, Height: 2}
	var result = StopMovementOnCollision(movingRect, stillRect, Right, &Point{X: 6, Y: 0})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.00001)
	assert.InDelta(t, 3.0, result.Y, 0.00001)
	assert.InDelta(t, 2.0, result.Width, 0.00001)
	assert.InDelta(t, 9.0, result.Height, 0.00001)
}

//--- Up Right -------------------------------------------------------------------------------------------------------//

func TestShouldStopTheUpRightMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 6, Width: 1, Height: 1}
	var stillRect = &Rectangle{X: 3, Y: 1, Width: 4, Height: 4}
	var result = stopUpRightMovement(movingRect, stillRect, &Point{X: 3, Y: -3})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.00001)
	assert.InDelta(t, 5.0, result.Y, 0.00001)
	assert.InDelta(t, 1.0, result.Width, 0.00001)
	assert.InDelta(t, 1.0, result.Height, 0.00001)
}

func TestShouldStopTheUpRightMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 4, Width: 6, Height: 6}
	var stillRect = &Rectangle{X: 4, Y: 2, Width: 1, Height: 1}
	var result = stopUpRightMovement(movingRect, stillRect, &Point{X: 2, Y: -3})

	assert.NotNil(t, result)
	assert.InDelta(t, 1.666, result.X, 0.01)
	assert.InDelta(t, 3.0, result.Y, 0.00001)
	assert.InDelta(t, 6.0, result.Width, 0.00001)
	assert.InDelta(t, 6.0, result.Height, 0.00001)
}

//--- Down Right -----------------------------------------------------------------------------------------------------//

func TestShouldStopTheDownRightMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 1, Width: 1, Height: 1}
	var stillRect = &Rectangle{X: 3, Y: 3, Width: 4, Height: 4}
	var result = stopDownRightMovement(movingRect, stillRect, &Point{X: 3, Y: 3})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.01)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
	assert.InDelta(t, 1.0, result.Width, 0.00001)
	assert.InDelta(t, 1.0, result.Height, 0.00001)
}

func TestShouldStopTheDownRightMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 1, Width: 4, Height: 4}
	var stillRect = &Rectangle{X: 4, Y: 6, Width: 1, Height: 1}
	var result = stopDownRightMovement(movingRect, stillRect, &Point{X: 3, Y: 3})

	assert.NotNil(t, result)
	assert.InDelta(t, 2.0, result.X, 0.01)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
	assert.InDelta(t, 4.0, result.Width, 0.00001)
	assert.InDelta(t, 4.0, result.Height, 0.00001)
}

//--- Up Left --------------------------------------------------------------------------------------------------------//

func TestShouldStopTheUpLeftMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 6, Y: 4, Width: 1, Height: 1}
	var stillRect = &Rectangle{X: 1, Y: 1, Width: 4, Height: 4}
	var result = stopDownLeftMovement(movingRect, stillRect, &Point{X: -2, Y: -2})

	assert.NotNil(t, result)
	assert.InDelta(t, 5.0, result.X, 0.01)
	assert.InDelta(t, 3.0, result.Y, 0.00001)
	assert.InDelta(t, 1.0, result.Width, 0.00001)
	assert.InDelta(t, 1.0, result.Height, 0.00001)
}

func TestShouldStopTheUpLeftMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 1, Y: 3, Width: 4, Height: 4}
	var stillRect = &Rectangle{X: 1, Y: 1, Width: 1, Height: 1}
	var result = stopUpLeftMovement(movingRect, stillRect, &Point{X: -3, Y: -3})

	assert.NotNil(t, result)
	assert.InDelta(t, 0.0, result.X, 0.01)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
	assert.InDelta(t, 4.0, result.Width, 0.00001)
	assert.InDelta(t, 4.0, result.Height, 0.00001)
}

//--- Down Left ------------------------------------------------------------------------------------------------------//

func TestShouldStopTheDownLeftMovementOfARectAgainstCenterOfLargerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 4, Y: 3, Width: 1, Height: 1}
	var stillRect = &Rectangle{X: 1, Y: 5, Width: 4, Height: 4}
	var result = stopDownLeftMovement(movingRect, stillRect, &Point{X: -3, Y: 3})

	assert.NotNil(t, result)
	assert.InDelta(t, 3.0, result.X, 0.01)
	assert.InDelta(t, 4.0, result.Y, 0.00001)
	assert.InDelta(t, 1.0, result.Width, 0.00001)
	assert.InDelta(t, 1.0, result.Height, 0.00001)
}

func TestShouldStopTheDownLeftMovementOfARectAgainstCenterOfSmallerRect(t *testing.T) {
	var movingRect = &Rectangle{X: 4, Y: 1, Width: 4, Height: 4}
	var stillRect = &Rectangle{X: 4, Y: 6, Width: 1, Height: 1}
	var result = stopDownLeftMovement(movingRect, stillRect, &Point{X: -2, Y: 2})

	assert.NotNil(t, result)
	assert.InDelta(t, 3.0, result.X, 0.01)
	assert.InDelta(t, 2.0, result.Y, 0.00001)
	assert.InDelta(t, 4.0, result.Width, 0.00001)
	assert.InDelta(t, 4.0, result.Height, 0.00001)
}
