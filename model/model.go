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

type DiveCompleteDTO struct {
	ShipPositionDTO
	Loot int
}

type ExplosionDTO struct {
	ShipPositionDTO
	ShipsDestroyed int
	MovesLeft      int
}
