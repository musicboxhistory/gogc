package equipmentstatus

import (
	"gogc/src/model"
)

type EquipmentStatus struct {
	Key    string                `json:"key" bson:"key"`
	Status model.EquipmentStatus `json:"status,omitempty" bson:"status,omitempty"`
}
