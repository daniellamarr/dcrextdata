// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMempools(t *testing.T) {
	t.Parallel()

	query := Mempools()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMempoolsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMempoolsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Mempools().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMempoolsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MempoolSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMempoolsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MempoolExists(ctx, tx, o.FirstSeenTime)
	if err != nil {
		t.Errorf("Unable to check if Mempool exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MempoolExists to return true, but got false.")
	}
}

func testMempoolsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mempoolFound, err := FindMempool(ctx, tx, o.FirstSeenTime)
	if err != nil {
		t.Error(err)
	}

	if mempoolFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMempoolsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Mempools().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMempoolsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Mempools().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMempoolsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mempoolOne := &Mempool{}
	mempoolTwo := &Mempool{}
	if err = randomize.Struct(seed, mempoolOne, mempoolDBTypes, false, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}
	if err = randomize.Struct(seed, mempoolTwo, mempoolDBTypes, false, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mempoolOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mempoolTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Mempools().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMempoolsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mempoolOne := &Mempool{}
	mempoolTwo := &Mempool{}
	if err = randomize.Struct(seed, mempoolOne, mempoolDBTypes, false, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}
	if err = randomize.Struct(seed, mempoolTwo, mempoolDBTypes, false, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mempoolOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mempoolTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testMempoolsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMempoolsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mempoolColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMempoolsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMempoolsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MempoolSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMempoolsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Mempools().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mempoolDBTypes = map[string]string{`FirstSeenTime`: `integer`, `BlockReceiveTime`: `integer`, `TotalSent`: `double precision`, `LastBlockHeight`: `double precision`, `Size`: `integer`, `RegularTransactionCount`: `integer`, `TicketCount`: `integer`, `VoteCount`: `integer`, `RevocationCount`: `integer`}
	_              = bytes.MinRead
)

func testMempoolsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mempoolPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mempoolAllColumns) == len(mempoolPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMempoolsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mempoolAllColumns) == len(mempoolPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Mempool{}
	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mempoolDBTypes, true, mempoolPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mempoolAllColumns, mempoolPrimaryKeyColumns) {
		fields = mempoolAllColumns
	} else {
		fields = strmangle.SetComplement(
			mempoolAllColumns,
			mempoolPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MempoolSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMempoolsUpsert(t *testing.T) {
	t.Parallel()

	if len(mempoolAllColumns) == len(mempoolPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Mempool{}
	if err = randomize.Struct(seed, &o, mempoolDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Mempool: %s", err)
	}

	count, err := Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mempoolDBTypes, false, mempoolPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mempool struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Mempool: %s", err)
	}

	count, err = Mempools().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
