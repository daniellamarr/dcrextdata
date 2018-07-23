// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testPosdatatables(t *testing.T) {
	t.Parallel()

	query := Posdatatables(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPosdatatablesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = posdatatable.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPosdatatablesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Posdatatables(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPosdatatablesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PosdatatableSlice{posdatatable}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPosdatatablesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PosdatatableExists(tx, posdatatable.ID)
	if err != nil {
		t.Errorf("Unable to check if Posdatatable exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PosdatatableExistsG to return true, but got false.")
	}
}
func testPosdatatablesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	posdatatableFound, err := FindPosdatatable(tx, posdatatable.ID)
	if err != nil {
		t.Error(err)
	}

	if posdatatableFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPosdatatablesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Posdatatables(tx).Bind(posdatatable); err != nil {
		t.Error(err)
	}
}

func testPosdatatablesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Posdatatables(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPosdatatablesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatableOne := &Posdatatable{}
	posdatatableTwo := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatableOne, posdatatableDBTypes, false, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}
	if err = randomize.Struct(seed, posdatatableTwo, posdatatableDBTypes, false, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatableOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = posdatatableTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Posdatatables(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPosdatatablesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	posdatatableOne := &Posdatatable{}
	posdatatableTwo := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatableOne, posdatatableDBTypes, false, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}
	if err = randomize.Struct(seed, posdatatableTwo, posdatatableDBTypes, false, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatableOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = posdatatableTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func posdatatableBeforeInsertHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableAfterInsertHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableAfterSelectHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableBeforeUpdateHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableAfterUpdateHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableBeforeDeleteHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableAfterDeleteHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableBeforeUpsertHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func posdatatableAfterUpsertHook(e boil.Executor, o *Posdatatable) error {
	*o = Posdatatable{}
	return nil
}

func testPosdatatablesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Posdatatable{}
	o := &Posdatatable{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, posdatatableDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Posdatatable object: %s", err)
	}

	AddPosdatatableHook(boil.BeforeInsertHook, posdatatableBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	posdatatableBeforeInsertHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.AfterInsertHook, posdatatableAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	posdatatableAfterInsertHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.AfterSelectHook, posdatatableAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	posdatatableAfterSelectHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.BeforeUpdateHook, posdatatableBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	posdatatableBeforeUpdateHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.AfterUpdateHook, posdatatableAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	posdatatableAfterUpdateHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.BeforeDeleteHook, posdatatableBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	posdatatableBeforeDeleteHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.AfterDeleteHook, posdatatableAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	posdatatableAfterDeleteHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.BeforeUpsertHook, posdatatableBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	posdatatableBeforeUpsertHooks = []PosdatatableHook{}

	AddPosdatatableHook(boil.AfterUpsertHook, posdatatableAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	posdatatableAfterUpsertHooks = []PosdatatableHook{}
}
func testPosdatatablesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPosdatatablesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx, posdatatableColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPosdatatablesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = posdatatable.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPosdatatablesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PosdatatableSlice{posdatatable}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPosdatatablesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Posdatatables(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	posdatatableDBTypes = map[string]string{`Apienabled`: `character varying`, `Apiversionssupported`: `character varying`, `ID`: `integer`, `Immature`: `character varying`, `Lastupdated`: `character varying`, `Launched`: `character varying`, `Live`: `character varying`, `Missed`: `character varying`, `Network`: `character varying`, `Poolfees`: `character varying`, `Posid`: `integer`, `Proportionlive`: `character varying`, `Proportionmissed`: `character varying`, `Timestamp`: `character varying`, `URL`: `character varying`, `Usercount`: `character varying`, `Usercountactive`: `character varying`, `Voted`: `character varying`}
	_                   = bytes.MinRead
)

func testPosdatatablesUpdate(t *testing.T) {
	t.Parallel()

	if len(posdatatableColumns) == len(posdatatablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	if err = posdatatable.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPosdatatablesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(posdatatableColumns) == len(posdatatablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	posdatatable := &Posdatatable{}
	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, posdatatable, posdatatableDBTypes, true, posdatatablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(posdatatableColumns, posdatatablePrimaryKeyColumns) {
		fields = posdatatableColumns
	} else {
		fields = strmangle.SetComplement(
			posdatatableColumns,
			posdatatablePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(posdatatable))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PosdatatableSlice{posdatatable}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPosdatatablesUpsert(t *testing.T) {
	t.Parallel()

	if len(posdatatableColumns) == len(posdatatablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	posdatatable := Posdatatable{}
	if err = randomize.Struct(seed, &posdatatable, posdatatableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = posdatatable.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Posdatatable: %s", err)
	}

	count, err := Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &posdatatable, posdatatableDBTypes, false, posdatatablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Posdatatable struct: %s", err)
	}

	if err = posdatatable.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Posdatatable: %s", err)
	}

	count, err = Posdatatables(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}