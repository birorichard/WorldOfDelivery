package validation

import (
	"math"

	"github.com/birorichard/WorldOfDelivery/model"
)

func IsShipMovementValid(PosX int, PosY int, PrevPosX int, PrevPosY int) bool {
	if PrevPosX == PosX && math.Abs(float64(PrevPosY-PosY)) == 1 ||
		PrevPosY == PosY && math.Abs(float64(PrevPosX-PosX)) == 1 ||
		math.Abs(float64(PrevPosX-PosX)) == 1 && math.Abs(float64(PrevPosY-PosY)) == 1 {
		return true
	}

	return false
}

func IsReachingDestinationValid(route *model.ShipRouteCache, dto *model.ShipReachedDestinationDTO) bool {

	return route.PlannedDestinationPortId == dto.PortId
}
