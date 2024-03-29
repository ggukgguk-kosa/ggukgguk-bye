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

func testMediaTypes(t *testing.T) {
	t.Parallel()

	query := MediaTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMediaTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
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

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MediaTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MediaTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MediaTypeExists(ctx, tx, o.MediaTypeID)
	if err != nil {
		t.Errorf("Unable to check if MediaType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MediaTypeExists to return true, but got false.")
	}
}

func testMediaTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mediaTypeFound, err := FindMediaType(ctx, tx, o.MediaTypeID)
	if err != nil {
		t.Error(err)
	}

	if mediaTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMediaTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MediaTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMediaTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MediaTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMediaTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mediaTypeOne := &MediaType{}
	mediaTypeTwo := &MediaType{}
	if err = randomize.Struct(seed, mediaTypeOne, mediaTypeDBTypes, false, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}
	if err = randomize.Struct(seed, mediaTypeTwo, mediaTypeDBTypes, false, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mediaTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mediaTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MediaTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMediaTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mediaTypeOne := &MediaType{}
	mediaTypeTwo := &MediaType{}
	if err = randomize.Struct(seed, mediaTypeOne, mediaTypeDBTypes, false, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}
	if err = randomize.Struct(seed, mediaTypeTwo, mediaTypeDBTypes, false, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mediaTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mediaTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mediaTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func mediaTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
	*o = MediaType{}
	return nil
}

func testMediaTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MediaType{}
	o := &MediaType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MediaType object: %s", err)
	}

	AddMediaTypeHook(boil.BeforeInsertHook, mediaTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mediaTypeBeforeInsertHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.AfterInsertHook, mediaTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mediaTypeAfterInsertHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.AfterSelectHook, mediaTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mediaTypeAfterSelectHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.BeforeUpdateHook, mediaTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mediaTypeBeforeUpdateHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.AfterUpdateHook, mediaTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mediaTypeAfterUpdateHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.BeforeDeleteHook, mediaTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mediaTypeBeforeDeleteHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.AfterDeleteHook, mediaTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mediaTypeAfterDeleteHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.BeforeUpsertHook, mediaTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mediaTypeBeforeUpsertHooks = []MediaTypeHook{}

	AddMediaTypeHook(boil.AfterUpsertHook, mediaTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mediaTypeAfterUpsertHooks = []MediaTypeHook{}
}

func testMediaTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMediaTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mediaTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMediaTypeToManyMediaFiles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaType
	var b, c MediaFile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.MediaTypeID = a.MediaTypeID
	c.MediaTypeID = a.MediaTypeID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.MediaFiles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.MediaTypeID == b.MediaTypeID {
			bFound = true
		}
		if v.MediaTypeID == c.MediaTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MediaTypeSlice{&a}
	if err = a.L.LoadMediaFiles(ctx, tx, false, (*[]*MediaType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MediaFiles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MediaFiles = nil
	if err = a.L.LoadMediaFiles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MediaFiles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMediaTypeToManyAddOpMediaFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaType
	var b, c, d, e MediaFile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaTypeDBTypes, false, strmangle.SetComplement(mediaTypePrimaryKeyColumns, mediaTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MediaFile{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*MediaFile{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMediaFiles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.MediaTypeID != first.MediaTypeID {
			t.Error("foreign key was wrong value", a.MediaTypeID, first.MediaTypeID)
		}
		if a.MediaTypeID != second.MediaTypeID {
			t.Error("foreign key was wrong value", a.MediaTypeID, second.MediaTypeID)
		}

		if first.R.MediaType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.MediaType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MediaFiles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MediaFiles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MediaFiles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMediaTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
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

func testMediaTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MediaTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMediaTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MediaTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mediaTypeDBTypes = map[string]string{`MediaTypeID`: `varchar`, `MediaTypeExtention`: `varchar`}
	_                = bytes.MinRead
)

func testMediaTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mediaTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mediaTypeAllColumns) == len(mediaTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMediaTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mediaTypeAllColumns) == len(mediaTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MediaType{}
	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mediaTypeDBTypes, true, mediaTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mediaTypeAllColumns, mediaTypePrimaryKeyColumns) {
		fields = mediaTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			mediaTypeAllColumns,
			mediaTypePrimaryKeyColumns,
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

	slice := MediaTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMediaTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(mediaTypeAllColumns) == len(mediaTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMediaTypeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MediaType{}
	if err = randomize.Struct(seed, &o, mediaTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MediaType: %s", err)
	}

	count, err := MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mediaTypeDBTypes, false, mediaTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MediaType: %s", err)
	}

	count, err = MediaTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
