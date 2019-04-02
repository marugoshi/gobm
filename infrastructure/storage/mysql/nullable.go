package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

// https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

func (ni NullInt64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int64, nil
}

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return errors.Wrap(err, "can not scan.")
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{Int64: i.Int64, Valid: false}
	} else {
		*ni = NullInt64{Int64: i.Int64, Valid: true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	result, err := json.Marshal(ni.Int64)
	if err != nil {
		return nil, errors.Wrap(err, "can not marshal.")
	}
	return result, nil
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = err == nil
	return errors.Wrap(err, "can not unmarshal.")
}

// NullBool is an alias for sql.NullBool data type
type NullBool sql.NullBool

func (nb NullBool) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.Bool, nil
}

// Scan implements the Scanner interface for NullBool
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return errors.Wrap(err, "can not scan.")
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NullBool{Bool: b.Bool, Valid: false}
	} else {
		*nb = NullBool{Bool: b.Bool, Valid: true}
	}

	return nil
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	result, err := json.Marshal(nb.Bool)
	if err != nil {
		return nil, errors.Wrap(err, "can not marshal.")
	}
	return result, err
}

// UnmarshalJSON for NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = err == nil
	return errors.Wrap(err, "can not unmarshal.")
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 sql.NullFloat64

func (nf NullFloat64) Value() (driver.Value, error) {
	if !nf.Valid {
		return nil, nil
	}
	return nf.Float64, nil
}

// Scan implements the Scanner interface for NullFloat64
func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return errors.Wrap(err, "can not scan.")
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{Float64: f.Float64, Valid: false}
	} else {
		*nf = NullFloat64{Float64: f.Float64, Valid: true}
	}

	return nil
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	result, err := json.Marshal(nf.Float64)
	if err != nil {
		return nil, errors.Wrap(err, "can not marshal.")
	}
	return result, err
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = err == nil
	return errors.Wrap(err, "can not unmarshal.")
}

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

func (ns NullString) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return nil, nil
}

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return errors.Wrap(err, "can not scan.")
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{String: s.String, Valid: false}
	} else {
		*ns = NullString{String: s.String, Valid: true}
	}

	return nil
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	result, err := json.Marshal(ns.String)
	if err != nil {
		return nil, errors.Wrap(err, "can not marshal.")
	}
	return result, err
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = err == nil
	return errors.Wrap(err, "can not unmarshal.")
}

// NullTime is an alias for mysql.NullTime data type focused on mysql date
type NullDate mysql.NullTime

func (nt NullDate) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// Scan implements the Scanner interface for NullTime
func (nt *NullDate) Scan(value interface{}) error {
	var t mysql.NullTime
	if err := t.Scan(value); err != nil {
		return errors.Wrap(err, "can not scan.")
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullDate{Time: t.Time, Valid: false}
	} else {
		*nt = NullDate{Time: t.Time, Valid: true}
	}

	return nil
}

func NullDateFormat() string {
	return "2006-01-02"
}

// MarshalJSON for NullTime
func (nt *NullDate) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(NullDateTimeFormat()))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullDate) UnmarshalJSON(b []byte) error {
	s := string(b)

	loc, _ := time.LoadLocation("Asia/Tokyo")
	x, err := time.ParseInLocation(`"`+NullDateFormat()+`"`, s, loc)
	if err != nil {
		nt.Valid = false
		return errors.Wrap(err, "can not unmarshal.")
	}

	nt.Time = x
	nt.Valid = true
	return nil
}

// NullTime is an alias for mysql.NullTime data type
type NullDateTime mysql.NullTime

func NewNullDateTime(year int, month time.Month, day, hour, min, sec int) NullDateTime {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return NullDateTime{
		Time:  time.Date(year, month, day, hour, min, sec, 0, loc),
		Valid: true,
	}
}

func (nt NullDateTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// Scan implements the Scanner interface for NullTime
func (nt *NullDateTime) Scan(value interface{}) error {
	var t mysql.NullTime
	if err := t.Scan(value); err != nil {
		return errors.Wrap(err, "can not scan.")
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullDateTime{Time: t.Time, Valid: false}
	} else {
		*nt = NullDateTime{Time: t.Time, Valid: true}
	}

	return nil
}

func NullDateTimeFormat() string {
	return "2006-01-02 15:04:05"
}

// MarshalJSON for NullTime
func (nt *NullDateTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(NullDateTimeFormat()))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullDateTime) UnmarshalJSON(b []byte) error {
	s := string(b)

	loc, _ := time.LoadLocation("Asia/Tokyo")
	x, err := time.ParseInLocation(`"`+NullDateTimeFormat()+`"`, s, loc)
	if err != nil {
		nt.Valid = false
		return errors.Wrap(err, "can not unmarshal.")
	}

	nt.Time = x
	nt.Valid = true
	return nil
}
