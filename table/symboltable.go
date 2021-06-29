package table

import (
	"github.com/iFREEGROUP/dxf/format"
	"github.com/iFREEGROUP/dxf/handle"
)

// SymbolTable is interface for AcDbSymbolTableRecord.
type SymbolTable interface {
	IsSymbolTable() bool
	Format(format.Formatter)
	Handle() int
	SetHandle(*int)
	SetOwner(handle.Handler)
	Name() string
}
