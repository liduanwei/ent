// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
	"github.com/facebookincubator/ent/schema/field"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	mutation *FieldTypeMutation
	hooks    []Hook
}

// SetInt sets the int field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.mutation.SetInt(i)
	return ftc
}

// SetInt8 sets the int8 field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetInt8(i)
	return ftc
}

// SetInt16 sets the int16 field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetInt16(i)
	return ftc
}

// SetInt32 sets the int32 field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetInt32(i)
	return ftc
}

// SetInt64 sets the int64 field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetInt64(i)
	return ftc
}

// SetOptionalInt sets the optional_int field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt(i)
	return ftc
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt8(i)
	return ftc
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt16(i)
	return ftc
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt32(i)
	return ftc
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt64(i)
	return ftc
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNillableInt sets the nillable_int field.
func (ftc *FieldTypeCreate) SetNillableInt(i int) *FieldTypeCreate {
	ftc.mutation.SetNillableInt(i)
	return ftc
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt(*i)
	}
	return ftc
}

// SetNillableInt8 sets the nillable_int8 field.
func (ftc *FieldTypeCreate) SetNillableInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetNillableInt8(i)
	return ftc
}

// SetNillableNillableInt8 sets the nillable_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt8(*i)
	}
	return ftc
}

// SetNillableInt16 sets the nillable_int16 field.
func (ftc *FieldTypeCreate) SetNillableInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetNillableInt16(i)
	return ftc
}

// SetNillableNillableInt16 sets the nillable_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt16(*i)
	}
	return ftc
}

// SetNillableInt32 sets the nillable_int32 field.
func (ftc *FieldTypeCreate) SetNillableInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetNillableInt32(i)
	return ftc
}

// SetNillableNillableInt32 sets the nillable_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt32(*i)
	}
	return ftc
}

// SetNillableInt64 sets the nillable_int64 field.
func (ftc *FieldTypeCreate) SetNillableInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetNillableInt64(i)
	return ftc
}

// SetNillableNillableInt64 sets the nillable_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt64(*i)
	}
	return ftc
}

// SetValidateOptionalInt32 sets the validate_optional_int32 field.
func (ftc *FieldTypeCreate) SetValidateOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetValidateOptionalInt32(i)
	return ftc
}

// SetNillableValidateOptionalInt32 sets the validate_optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetValidateOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalUint sets the optional_uint field.
func (ftc *FieldTypeCreate) SetOptionalUint(u uint) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint(u)
	return ftc
}

// SetNillableOptionalUint sets the optional_uint field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint(u *uint) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint(*u)
	}
	return ftc
}

// SetOptionalUint8 sets the optional_uint8 field.
func (ftc *FieldTypeCreate) SetOptionalUint8(u uint8) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint8(u)
	return ftc
}

// SetNillableOptionalUint8 sets the optional_uint8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint8(u *uint8) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint8(*u)
	}
	return ftc
}

// SetOptionalUint16 sets the optional_uint16 field.
func (ftc *FieldTypeCreate) SetOptionalUint16(u uint16) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint16(u)
	return ftc
}

// SetNillableOptionalUint16 sets the optional_uint16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint16(u *uint16) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint16(*u)
	}
	return ftc
}

// SetOptionalUint32 sets the optional_uint32 field.
func (ftc *FieldTypeCreate) SetOptionalUint32(u uint32) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint32(u)
	return ftc
}

// SetNillableOptionalUint32 sets the optional_uint32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint32(u *uint32) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint32(*u)
	}
	return ftc
}

// SetOptionalUint64 sets the optional_uint64 field.
func (ftc *FieldTypeCreate) SetOptionalUint64(u uint64) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint64(u)
	return ftc
}

// SetNillableOptionalUint64 sets the optional_uint64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint64(u *uint64) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint64(*u)
	}
	return ftc
}

// SetState sets the state field.
func (ftc *FieldTypeCreate) SetState(f fieldtype.State) *FieldTypeCreate {
	ftc.mutation.SetState(f)
	return ftc
}

// SetNillableState sets the state field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableState(f *fieldtype.State) *FieldTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// SetOptionalFloat sets the optional_float field.
func (ftc *FieldTypeCreate) SetOptionalFloat(f float64) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat(f)
	return ftc
}

// SetNillableOptionalFloat sets the optional_float field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat(*f)
	}
	return ftc
}

// SetOptionalFloat32 sets the optional_float32 field.
func (ftc *FieldTypeCreate) SetOptionalFloat32(f float32) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat32(f)
	return ftc
}

// SetNillableOptionalFloat32 sets the optional_float32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat32(f *float32) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat32(*f)
	}
	return ftc
}

// SetDatetime sets the datetime field.
func (ftc *FieldTypeCreate) SetDatetime(t time.Time) *FieldTypeCreate {
	ftc.mutation.SetDatetime(t)
	return ftc
}

// SetNillableDatetime sets the datetime field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDatetime(t *time.Time) *FieldTypeCreate {
	if t != nil {
		ftc.SetDatetime(*t)
	}
	return ftc
}

// SetDecimal sets the decimal field.
func (ftc *FieldTypeCreate) SetDecimal(f float64) *FieldTypeCreate {
	ftc.mutation.SetDecimal(f)
	return ftc
}

// SetNillableDecimal sets the decimal field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDecimal(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetDecimal(*f)
	}
	return ftc
}

// SetDir sets the dir field.
func (ftc *FieldTypeCreate) SetDir(h http.Dir) *FieldTypeCreate {
	ftc.mutation.SetDir(h)
	return ftc
}

// SetNillableDir sets the dir field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDir(h *http.Dir) *FieldTypeCreate {
	if h != nil {
		ftc.SetDir(*h)
	}
	return ftc
}

// SetNdir sets the ndir field.
func (ftc *FieldTypeCreate) SetNdir(h http.Dir) *FieldTypeCreate {
	ftc.mutation.SetNdir(h)
	return ftc
}

// SetNillableNdir sets the ndir field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNdir(h *http.Dir) *FieldTypeCreate {
	if h != nil {
		ftc.SetNdir(*h)
	}
	return ftc
}

// SetStr sets the str field.
func (ftc *FieldTypeCreate) SetStr(ss sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetStr(ss)
	return ftc
}

// SetNullStr sets the null_str field.
func (ftc *FieldTypeCreate) SetNullStr(ss sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetNullStr(ss)
	return ftc
}

// SetLink sets the link field.
func (ftc *FieldTypeCreate) SetLink(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLink(s)
	return ftc
}

// SetNullLink sets the null_link field.
func (ftc *FieldTypeCreate) SetNullLink(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetNullLink(s)
	return ftc
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	if _, ok := ftc.mutation.Int(); !ok {
		return nil, errors.New("ent: missing required field \"int\"")
	}
	if _, ok := ftc.mutation.Int8(); !ok {
		return nil, errors.New("ent: missing required field \"int8\"")
	}
	if _, ok := ftc.mutation.Int16(); !ok {
		return nil, errors.New("ent: missing required field \"int16\"")
	}
	if _, ok := ftc.mutation.Int32(); !ok {
		return nil, errors.New("ent: missing required field \"int32\"")
	}
	if _, ok := ftc.mutation.Int64(); !ok {
		return nil, errors.New("ent: missing required field \"int64\"")
	}
	if v, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		if err := fieldtype.ValidateOptionalInt32Validator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %w", err)
		}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := fieldtype.StateValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"state\": %w", err)
		}
	}
	var (
		err  error
		node *FieldType
	)
	if len(ftc.hooks) == 0 {
		node, err = ftc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FieldTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftc.mutation = mutation
			node, err = ftc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftc.hooks) - 1; i >= 0; i-- {
			mut = ftc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FieldTypeCreate) sqlSave(ctx context.Context) (*FieldType, error) {
	var (
		ft    = &FieldType{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fieldtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fieldtype.FieldID,
			},
		}
	)
	if value, ok := ftc.mutation.Int(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldInt,
		})
		ft.Int = value
	}
	if value, ok := ftc.mutation.Int8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldInt8,
		})
		ft.Int8 = value
	}
	if value, ok := ftc.mutation.Int16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fieldtype.FieldInt16,
		})
		ft.Int16 = value
	}
	if value, ok := ftc.mutation.Int32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldInt32,
		})
		ft.Int32 = value
	}
	if value, ok := ftc.mutation.Int64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldInt64,
		})
		ft.Int64 = value
	}
	if value, ok := ftc.mutation.OptionalInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldOptionalInt,
		})
		ft.OptionalInt = value
	}
	if value, ok := ftc.mutation.OptionalInt8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldOptionalInt8,
		})
		ft.OptionalInt8 = value
	}
	if value, ok := ftc.mutation.OptionalInt16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fieldtype.FieldOptionalInt16,
		})
		ft.OptionalInt16 = value
	}
	if value, ok := ftc.mutation.OptionalInt32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldOptionalInt32,
		})
		ft.OptionalInt32 = value
	}
	if value, ok := ftc.mutation.OptionalInt64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldOptionalInt64,
		})
		ft.OptionalInt64 = value
	}
	if value, ok := ftc.mutation.NillableInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldNillableInt,
		})
		ft.NillableInt = &value
	}
	if value, ok := ftc.mutation.NillableInt8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldNillableInt8,
		})
		ft.NillableInt8 = &value
	}
	if value, ok := ftc.mutation.NillableInt16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fieldtype.FieldNillableInt16,
		})
		ft.NillableInt16 = &value
	}
	if value, ok := ftc.mutation.NillableInt32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldNillableInt32,
		})
		ft.NillableInt32 = &value
	}
	if value, ok := ftc.mutation.NillableInt64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldNillableInt64,
		})
		ft.NillableInt64 = &value
	}
	if value, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldValidateOptionalInt32,
		})
		ft.ValidateOptionalInt32 = value
	}
	if value, ok := ftc.mutation.OptionalUint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: fieldtype.FieldOptionalUint,
		})
		ft.OptionalUint = value
	}
	if value, ok := ftc.mutation.OptionalUint8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: fieldtype.FieldOptionalUint8,
		})
		ft.OptionalUint8 = value
	}
	if value, ok := ftc.mutation.OptionalUint16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: fieldtype.FieldOptionalUint16,
		})
		ft.OptionalUint16 = value
	}
	if value, ok := ftc.mutation.OptionalUint32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fieldtype.FieldOptionalUint32,
		})
		ft.OptionalUint32 = value
	}
	if value, ok := ftc.mutation.OptionalUint64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: fieldtype.FieldOptionalUint64,
		})
		ft.OptionalUint64 = value
	}
	if value, ok := ftc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: fieldtype.FieldState,
		})
		ft.State = value
	}
	if value, ok := ftc.mutation.OptionalFloat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fieldtype.FieldOptionalFloat,
		})
		ft.OptionalFloat = value
	}
	if value, ok := ftc.mutation.OptionalFloat32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: fieldtype.FieldOptionalFloat32,
		})
		ft.OptionalFloat32 = value
	}
	if value, ok := ftc.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fieldtype.FieldDatetime,
		})
		ft.Datetime = value
	}
	if value, ok := ftc.mutation.Decimal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fieldtype.FieldDecimal,
		})
		ft.Decimal = value
	}
	if value, ok := ftc.mutation.Dir(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldDir,
		})
		ft.Dir = value
	}
	if value, ok := ftc.mutation.Ndir(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldNdir,
		})
		ft.Ndir = &value
	}
	if value, ok := ftc.mutation.Str(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldStr,
		})
		ft.Str = value
	}
	if value, ok := ftc.mutation.NullStr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldNullStr,
		})
		ft.NullStr = &value
	}
	if value, ok := ftc.mutation.Link(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldLink,
		})
		ft.Link = value
	}
	if value, ok := ftc.mutation.NullLink(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldNullLink,
		})
		ft.NullLink = &value
	}
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ft.ID = int(id)
	return ft, nil
}
