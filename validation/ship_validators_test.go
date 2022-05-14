package validation

import (
	"testing"

	"github.com/birorichard/WorldOfDelivery/model"
	"github.com/birorichard/WorldOfDelivery/test_data"
)

func TestStartShipTrackingShouldAddNewRouteToTheCache(t *testing.T) {
	cachedRoute := model.ShipRouteCache{
		TableData: model.ShipRouteDTO{
			SourcePortId:      3,
			DestinationPortId: 12,
			Steps: []model.Position{
				{X: 23, Y: 3, StepOrder: 0},
				{X: 23, Y: 4, StepOrder: 1},
				{X: 23, Y: 5, StepOrder: 2},
				{X: 23, Y: 6, StepOrder: 3},
			},
			Commited: false,
		},
		PlannedDestinationPortId: 12,
		Discovered:               true,
		ShipId:                   test_data.TestShipIds[1],
	}

	shipDestinationReachedDto := model.ShipReachedDestinationDTO{
		ShipPositionDTO: model.ShipPositionDTO{
			ShipId: test_data.TestShipIds[0],
			X:      670,
			Y:      512,
		},
		PortId:  99,
		Payment: 10,
	}

	validationResult := IsReachingDestinationValid(&cachedRoute, &shipDestinationReachedDto)

	if validationResult {
		t.Fatal("Destination PortIds wasn't compared properly, validation result should be false, but returned with true!")
	}

	shipDestinationReachedDto.PortId = 12
	validationResult = IsReachingDestinationValid(&cachedRoute, &shipDestinationReachedDto)
	if !validationResult {
		t.Fatal("Destination PortIds wasn't compared properly, validation result should be true, but it returned with false!")
	}

}

func TestIsShipMovementValidShouldReturnsFalseInCaseOfForbiddenStep(t *testing.T) {
	prevPosX := []int{1, 1, 1, 2, 3, 4, 5, 5, 5, 5, 1, 1, 2, 5, 3}
	prevPosY := []int{2, 3, 4, 5, 5, 5, 4, 3, 2, 1, 5, 1, 5, 5, 3}

	posX := 3
	posY := 3

	for index := range prevPosX {
		if IsShipMovementValid(posX, posY, prevPosX[index], prevPosY[index]) {
			t.Fatal("Forbidden ship movement should not be accepted!")
		}
	}
}

func TestIsShipMovementValidShouldReturnsTrueInCaseOfAcceptedStep(t *testing.T) {
	prevPosX := []int{2, 3, 4, 2, 4, 2, 3, 4}
	prevPosY := []int{2, 2, 2, 3, 3, 4, 4, 4}

	posX := 3
	posY := 3

	for index := range prevPosX {
		if !IsShipMovementValid(posX, posY, prevPosX[index], prevPosY[index]) {
			t.Fatal("Accepted ship movement should be accepted!")
		}
	}
}
