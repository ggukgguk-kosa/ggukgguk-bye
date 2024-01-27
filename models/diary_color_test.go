// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDiaryColors(t *testing.T) {
	t.Parallel()

	query := DiaryColors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDiaryColorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
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

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiaryColorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DiaryColors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiaryColorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DiaryColorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiaryColorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DiaryColorExists(ctx, tx, o.DiaryColorID)
	if err != nil {
		t.Errorf("Unable to check if DiaryColor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiaryColorExists to return true, but got false.")
	}
}

func testDiaryColorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	diaryColorFound, err := FindDiaryColor(ctx, tx, o.DiaryColorID)
	if err != nil {
		t.Error(err)
	}

	if diaryColorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDiaryColorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DiaryColors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDiaryColorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DiaryColors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiaryColorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	diaryColorOne := &DiaryColor{}
	diaryColorTwo := &DiaryColor{}
	if err = randomize.Struct(seed, diaryColorOne, diaryColorDBTypes, false, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}
	if err = randomize.Struct(seed, diaryColorTwo, diaryColorDBTypes, false, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = diaryColorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = diaryColorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DiaryColors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiaryColorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	diaryColorOne := &DiaryColor{}
	diaryColorTwo := &DiaryColor{}
	if err = randomize.Struct(seed, diaryColorOne, diaryColorDBTypes, false, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}
	if err = randomize.Struct(seed, diaryColorTwo, diaryColorDBTypes, false, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = diaryColorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = diaryColorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func diaryColorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func diaryColorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DiaryColor) error {
	*o = DiaryColor{}
	return nil
}

func testDiaryColorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DiaryColor{}
	o := &DiaryColor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, diaryColorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DiaryColor object: %s", err)
	}

	AddDiaryColorHook(boil.BeforeInsertHook, diaryColorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	diaryColorBeforeInsertHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.AfterInsertHook, diaryColorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	diaryColorAfterInsertHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.AfterSelectHook, diaryColorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	diaryColorAfterSelectHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.BeforeUpdateHook, diaryColorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	diaryColorBeforeUpdateHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.AfterUpdateHook, diaryColorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	diaryColorAfterUpdateHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.BeforeDeleteHook, diaryColorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	diaryColorBeforeDeleteHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.AfterDeleteHook, diaryColorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	diaryColorAfterDeleteHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.BeforeUpsertHook, diaryColorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	diaryColorBeforeUpsertHooks = []DiaryColorHook{}

	AddDiaryColorHook(boil.AfterUpsertHook, diaryColorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	diaryColorAfterUpsertHooks = []DiaryColorHook{}
}

func testDiaryColorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiaryColorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(diaryColorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiaryColorToOneDiaryUsingDiary(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DiaryColor
	var foreign Diary

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, diaryColorDBTypes, false, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, diaryDBTypes, false, diaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Diary struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DiaryID = foreign.DiaryID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Diary().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.DiaryID != foreign.DiaryID {
		t.Errorf("want: %v, got %v", foreign.DiaryID, check.DiaryID)
	}

	ranAfterSelectHook := false
	AddDiaryHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Diary) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := DiaryColorSlice{&local}
	if err = local.L.LoadDiary(ctx, tx, false, (*[]*DiaryColor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Diary == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Diary = nil
	if err = local.L.LoadDiary(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Diary == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testDiaryColorToOneSetOpDiaryUsingDiary(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DiaryColor
	var b, c Diary

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, diaryColorDBTypes, false, strmangle.SetComplement(diaryColorPrimaryKeyColumns, diaryColorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, diaryDBTypes, false, strmangle.SetComplement(diaryPrimaryKeyColumns, diaryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, diaryDBTypes, false, strmangle.SetComplement(diaryPrimaryKeyColumns, diaryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Diary{&b, &c} {
		err = a.SetDiary(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Diary != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DiaryColors[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DiaryID != x.DiaryID {
			t.Error("foreign key was wrong value", a.DiaryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DiaryID))
		reflect.Indirect(reflect.ValueOf(&a.DiaryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DiaryID != x.DiaryID {
			t.Error("foreign key was wrong value", a.DiaryID, x.DiaryID)
		}
	}
}

func testDiaryColorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
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

func testDiaryColorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DiaryColorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDiaryColorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DiaryColors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	diaryColorDBTypes = map[string]string{`DiaryColorID`: `int`, `DiaryID`: `int`, `DiaryColor`: `varchar`}
	_                 = bytes.MinRead
)

func testDiaryColorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(diaryColorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(diaryColorAllColumns) == len(diaryColorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDiaryColorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(diaryColorAllColumns) == len(diaryColorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DiaryColor{}
	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, diaryColorDBTypes, true, diaryColorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(diaryColorAllColumns, diaryColorPrimaryKeyColumns) {
		fields = diaryColorAllColumns
	} else {
		fields = strmangle.SetComplement(
			diaryColorAllColumns,
			diaryColorPrimaryKeyColumns,
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

	slice := DiaryColorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDiaryColorsUpsert(t *testing.T) {
	t.Parallel()

	if len(diaryColorAllColumns) == len(diaryColorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDiaryColorUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DiaryColor{}
	if err = randomize.Struct(seed, &o, diaryColorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DiaryColor: %s", err)
	}

	count, err := DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, diaryColorDBTypes, false, diaryColorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiaryColor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DiaryColor: %s", err)
	}

	count, err = DiaryColors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}