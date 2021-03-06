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

// For response
type ShipRouteDTO struct {
	SourcePortId      int
	DestinationPortId int
	Steps             []Position
	Committed         bool
}

// Inner use for cache routes
type ShipRouteCache struct {
	TableData                ShipRouteDTO
	PlannedDestinationPortId int
	Discovered               bool
	ShipId                   string
}

type TableSize struct {
	X int
	Y int
}

type GetAllRoutesResponseDTO struct {
	TableSize TableSize
	Routes    []ShipRouteDTO
}

type RouteEntity struct {
	ID                int
	SourcePortId      int
	DestinationPortId int
	PosX              int
	PosY              int
	StepOrder         int
}
