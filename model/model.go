package model

type ShipPositionDTO struct {
	ShipId string
	X      int
	Y      int
}

type ShipReachedDestinationDTO struct {
	ShipPositionDTO
	PortId  int
	Payment int
}

type ShipLeavePortDTO struct {
	ShipPositionDTO
	PortId          int
	DestinationPort int
}

// StepOrder-t torolni
type Position struct {
	X         int
	Y         int
	StepOrder int
}

// For API response
type ShipRouteDTO struct {
	SourcePortId      int
	DestinationPortId int
	Steps             []Position
	Commited          bool
}

// Inner use for cache routes
type ShipRouteCache struct {
	TableData  ShipRouteDTO
	Discovered bool
	ShipId     string
}
