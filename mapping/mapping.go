package mapping

import (
	"github.com/birorichard/WorldOfDelivery/model"
)

func MapFromShipRouteCacheToRouteEntities(dto *model.ShipRouteCache) []model.RouteEntity {
	entities := []model.RouteEntity{}

	sourceId := dto.TableData.SourcePortId
	destinationId := dto.TableData.DestinationPortId
	for index, element := range dto.TableData.Steps {
		entities = append(entities, model.RouteEntity{SourcePortId: sourceId, DestinationPortId: destinationId, PosX: element.X, PosY: element.Y, StepOrder: index})
	}

	return entities

}

// func MapFromShipLeavePortDTOToShipRouteCache(dto *model.ShipLeavePortDTO) model.ShipRouteCache {
// 	return model.ShipRouteCache{
// 		TableData: model.ShipRouteDTO{
// 			SourcePortId:      dto.PortId,
// 			DestinationPortId: dto.DestinationPort,
// 			Steps:             []model.Position{},
// 			Commited:          false,
// 		},
// 		Discovered:               false,
// 		ShipId:                   dto.ShipId,
// 		PlannedDestinationPortId: dto.DestinationPort,
// 	}
// }
