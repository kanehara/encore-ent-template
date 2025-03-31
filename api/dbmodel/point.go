package dbmodel

import (
	"database/sql/driver"
	"errors"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

type Point struct {
	*geom.Point
}

func (p *Point) Value() (driver.Value, error) {
	encodedGeometry, err := ewkbhex.Encode(p.Point, ewkbhex.NDR)
	if err != nil {
			return nil, err
	}
	return encodedGeometry, nil
}

func (p *Point) Scan(value interface{}) error {
	stringValue, ok := value.(string)
	if !ok {
			return errors.New("value is no string")
	}
	t, err := ewkbhex.Decode(stringValue)
	if err != nil {
			return err
	}
	p.Point, ok = t.(*geom.Point)
	if !ok {
			return errors.New("value is no point")
	}
	return nil
}
