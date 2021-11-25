package orm

import (
	"gorm.io/gorm"
)

type Basic struct {
	gorm.Model

	guid string `gorm:"uniqueIndex;size:256"` // FIXXME: https://gorm.io/docs/data_types.html
	name string `gorm:"uniqueIndex;size:256"`
	desc string `gorm:"size:1048576"`
	slug string `gorm:"uniqueIndex;size:256"`
}

// ObjType is a type of object
type ObjType struct {
	Basic Basic `gorm:"embedded"`
}

type Object struct {
	Basic   Basic  `gorm:"embedded"`
	ObjType string ``
	//parents
	//func children
}
