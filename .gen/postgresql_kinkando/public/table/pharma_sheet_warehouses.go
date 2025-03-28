//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PharmaSheetWarehouses = newPharmaSheetWarehousesTable("public", "pharma_sheet_warehouses", "")

type pharmaSheetWarehousesTable struct {
	postgres.Table

	// Columns
	WarehouseID postgres.ColumnString
	Name        postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz
	UpdatedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PharmaSheetWarehousesTable struct {
	pharmaSheetWarehousesTable

	EXCLUDED pharmaSheetWarehousesTable
}

// AS creates new PharmaSheetWarehousesTable with assigned alias
func (a PharmaSheetWarehousesTable) AS(alias string) *PharmaSheetWarehousesTable {
	return newPharmaSheetWarehousesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PharmaSheetWarehousesTable with assigned schema name
func (a PharmaSheetWarehousesTable) FromSchema(schemaName string) *PharmaSheetWarehousesTable {
	return newPharmaSheetWarehousesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PharmaSheetWarehousesTable with assigned table prefix
func (a PharmaSheetWarehousesTable) WithPrefix(prefix string) *PharmaSheetWarehousesTable {
	return newPharmaSheetWarehousesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PharmaSheetWarehousesTable with assigned table suffix
func (a PharmaSheetWarehousesTable) WithSuffix(suffix string) *PharmaSheetWarehousesTable {
	return newPharmaSheetWarehousesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPharmaSheetWarehousesTable(schemaName, tableName, alias string) *PharmaSheetWarehousesTable {
	return &PharmaSheetWarehousesTable{
		pharmaSheetWarehousesTable: newPharmaSheetWarehousesTableImpl(schemaName, tableName, alias),
		EXCLUDED:                   newPharmaSheetWarehousesTableImpl("", "excluded", ""),
	}
}

func newPharmaSheetWarehousesTableImpl(schemaName, tableName, alias string) pharmaSheetWarehousesTable {
	var (
		WarehouseIDColumn = postgres.StringColumn("warehouse_id")
		NameColumn        = postgres.StringColumn("name")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		allColumns        = postgres.ColumnList{WarehouseIDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return pharmaSheetWarehousesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		WarehouseID: WarehouseIDColumn,
		Name:        NameColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
