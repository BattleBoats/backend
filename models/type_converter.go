package models

import (
	"github.com/coopernurse/gorp"
)

//TypeConverter Interface for gorp
type TypeConverter struct{}

// ToDb converts val to another type. Called before INSERT/UPDATE operations
func (c TypeConverter) ToDb(val interface{}) (interface{}, error) {
	switch v := val.(type) {
	case NullInt64:
		if v.Valid {
			return v.Int64, nil
		}
	case NullString:
		if v.Valid {
			return v.String, nil
		}
	case NullTime:
		if v.Valid {
			return v.Time, nil
		}
	}
	return val, nil
}

// FromDb returns a CustomScanner appropriate for this type. This will be used
// to hold values returned from SELECT queries.
//
// In particular the CustomScanner returned should implement a Binder
// function appropriate for the Go type you wish to convert the db value to
//
// If bool==false, then no custom scanner will be used for this field.
func (c TypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	return gorp.CustomScanner{}, false
}
