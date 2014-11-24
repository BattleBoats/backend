package models

import (
	"encoding/json"
	"fmt"

	"backend/errors"

	"github.com/coopernurse/gorp"
)

//TypeConverter Interface for gorp
type TypeConverter struct{}

// ToDb converts val to another type. Called before INSERT/UPDATE operations
func (c TypeConverter) ToDb(val interface{}) (interface{}, error) {
	switch v := val.(type) {
	case *GameBoard:
		json, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		return json, nil
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
	switch target.(type) {
	default:
		fmt.Printf("unexpected type %T\n", target)
	case *GameBoard:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New(nil, "FromDb: Unable to convert GameBoard entry to *string", 500)
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}

		return gorp.CustomScanner{new(string), target, binder}, true
	}

	return gorp.CustomScanner{}, false
}
