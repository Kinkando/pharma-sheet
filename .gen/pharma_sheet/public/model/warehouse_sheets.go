//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type WarehouseSheets struct {
	WarehouseID    uuid.UUID
	SpreadsheetID  string
	SheetID        int32
	LatestSyncedAt time.Time
	CreatedAt      time.Time
}
