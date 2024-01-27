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

func testMediaFiles(t *testing.T) {
	t.Parallel()

	query := MediaFiles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMediaFilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
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

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaFilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MediaFiles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaFilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MediaFileSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMediaFilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MediaFileExists(ctx, tx, o.MediaFileID)
	if err != nil {
		t.Errorf("Unable to check if MediaFile exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MediaFileExists to return true, but got false.")
	}
}

func testMediaFilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mediaFileFound, err := FindMediaFile(ctx, tx, o.MediaFileID)
	if err != nil {
		t.Error(err)
	}

	if mediaFileFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMediaFilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MediaFiles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMediaFilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MediaFiles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMediaFilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mediaFileOne := &MediaFile{}
	mediaFileTwo := &MediaFile{}
	if err = randomize.Struct(seed, mediaFileOne, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}
	if err = randomize.Struct(seed, mediaFileTwo, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mediaFileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mediaFileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MediaFiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMediaFilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mediaFileOne := &MediaFile{}
	mediaFileTwo := &MediaFile{}
	if err = randomize.Struct(seed, mediaFileOne, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}
	if err = randomize.Struct(seed, mediaFileTwo, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mediaFileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mediaFileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mediaFileBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func mediaFileAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
	*o = MediaFile{}
	return nil
}

func testMediaFilesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MediaFile{}
	o := &MediaFile{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mediaFileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MediaFile object: %s", err)
	}

	AddMediaFileHook(boil.BeforeInsertHook, mediaFileBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mediaFileBeforeInsertHooks = []MediaFileHook{}

	AddMediaFileHook(boil.AfterInsertHook, mediaFileAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mediaFileAfterInsertHooks = []MediaFileHook{}

	AddMediaFileHook(boil.AfterSelectHook, mediaFileAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mediaFileAfterSelectHooks = []MediaFileHook{}

	AddMediaFileHook(boil.BeforeUpdateHook, mediaFileBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mediaFileBeforeUpdateHooks = []MediaFileHook{}

	AddMediaFileHook(boil.AfterUpdateHook, mediaFileAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mediaFileAfterUpdateHooks = []MediaFileHook{}

	AddMediaFileHook(boil.BeforeDeleteHook, mediaFileBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mediaFileBeforeDeleteHooks = []MediaFileHook{}

	AddMediaFileHook(boil.AfterDeleteHook, mediaFileAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mediaFileAfterDeleteHooks = []MediaFileHook{}

	AddMediaFileHook(boil.BeforeUpsertHook, mediaFileBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mediaFileBeforeUpsertHooks = []MediaFileHook{}

	AddMediaFileHook(boil.AfterUpsertHook, mediaFileAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mediaFileAfterUpsertHooks = []MediaFileHook{}
}

func testMediaFilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMediaFilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mediaFileColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMediaFileToManyRecords(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c Record

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, recordDBTypes, false, recordColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, recordDBTypes, false, recordColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.MediaFileID, a.MediaFileID)
	queries.Assign(&c.MediaFileID, a.MediaFileID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Records().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.MediaFileID, b.MediaFileID) {
			bFound = true
		}
		if queries.Equal(v.MediaFileID, c.MediaFileID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MediaFileSlice{&a}
	if err = a.L.LoadRecords(ctx, tx, false, (*[]*MediaFile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Records); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Records = nil
	if err = a.L.LoadRecords(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Records); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMediaFileToManyRecordTests(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c RecordTest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.MediaFileID, a.MediaFileID)
	queries.Assign(&c.MediaFileID, a.MediaFileID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RecordTests().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.MediaFileID, b.MediaFileID) {
			bFound = true
		}
		if queries.Equal(v.MediaFileID, c.MediaFileID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MediaFileSlice{&a}
	if err = a.L.LoadRecordTests(ctx, tx, false, (*[]*MediaFile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RecordTests); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RecordTests = nil
	if err = a.L.LoadRecordTests(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RecordTests); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMediaFileToManyAddOpRecords(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c, d, e Record

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Record{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordDBTypes, false, strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Record{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRecords(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.MediaFileID, first.MediaFileID) {
			t.Error("foreign key was wrong value", a.MediaFileID, first.MediaFileID)
		}
		if !queries.Equal(a.MediaFileID, second.MediaFileID) {
			t.Error("foreign key was wrong value", a.MediaFileID, second.MediaFileID)
		}

		if first.R.MediaFile != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.MediaFile != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Records[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Records[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Records().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMediaFileToManySetOpRecords(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c, d, e Record

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Record{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordDBTypes, false, strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetRecords(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Records().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRecords(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Records().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MediaFileID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MediaFileID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.MediaFileID, d.MediaFileID) {
		t.Error("foreign key was wrong value", a.MediaFileID, d.MediaFileID)
	}
	if !queries.Equal(a.MediaFileID, e.MediaFileID) {
		t.Error("foreign key was wrong value", a.MediaFileID, e.MediaFileID)
	}

	if b.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.MediaFile != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.MediaFile != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Records[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Records[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMediaFileToManyRemoveOpRecords(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c, d, e Record

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Record{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordDBTypes, false, strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddRecords(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Records().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRecords(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Records().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MediaFileID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MediaFileID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.MediaFile != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.MediaFile != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Records) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Records[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Records[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMediaFileToManyAddOpRecordTests(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c, d, e RecordTest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RecordTest{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordTestDBTypes, false, strmangle.SetComplement(recordTestPrimaryKeyColumns, recordTestColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*RecordTest{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRecordTests(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.MediaFileID, first.MediaFileID) {
			t.Error("foreign key was wrong value", a.MediaFileID, first.MediaFileID)
		}
		if !queries.Equal(a.MediaFileID, second.MediaFileID) {
			t.Error("foreign key was wrong value", a.MediaFileID, second.MediaFileID)
		}

		if first.R.MediaFile != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.MediaFile != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RecordTests[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RecordTests[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RecordTests().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMediaFileToManySetOpRecordTests(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c, d, e RecordTest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RecordTest{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordTestDBTypes, false, strmangle.SetComplement(recordTestPrimaryKeyColumns, recordTestColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetRecordTests(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RecordTests().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRecordTests(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RecordTests().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MediaFileID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MediaFileID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.MediaFileID, d.MediaFileID) {
		t.Error("foreign key was wrong value", a.MediaFileID, d.MediaFileID)
	}
	if !queries.Equal(a.MediaFileID, e.MediaFileID) {
		t.Error("foreign key was wrong value", a.MediaFileID, e.MediaFileID)
	}

	if b.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.MediaFile != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.MediaFile != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.RecordTests[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.RecordTests[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMediaFileToManyRemoveOpRecordTests(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c, d, e RecordTest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RecordTest{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordTestDBTypes, false, strmangle.SetComplement(recordTestPrimaryKeyColumns, recordTestColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddRecordTests(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RecordTests().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRecordTests(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RecordTests().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MediaFileID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MediaFileID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.MediaFile != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.MediaFile != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.MediaFile != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.RecordTests) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.RecordTests[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.RecordTests[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMediaFileToOneMediaTypeUsingMediaType(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MediaFile
	var foreign MediaType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mediaTypeDBTypes, false, mediaTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MediaTypeID = foreign.MediaTypeID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.MediaType().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.MediaTypeID != foreign.MediaTypeID {
		t.Errorf("want: %v, got %v", foreign.MediaTypeID, check.MediaTypeID)
	}

	ranAfterSelectHook := false
	AddMediaTypeHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *MediaType) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := MediaFileSlice{&local}
	if err = local.L.LoadMediaType(ctx, tx, false, (*[]*MediaFile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MediaType == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.MediaType = nil
	if err = local.L.LoadMediaType(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MediaType == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testMediaFileToOneSetOpMediaTypeUsingMediaType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MediaFile
	var b, c MediaType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mediaTypeDBTypes, false, strmangle.SetComplement(mediaTypePrimaryKeyColumns, mediaTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mediaTypeDBTypes, false, strmangle.SetComplement(mediaTypePrimaryKeyColumns, mediaTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MediaType{&b, &c} {
		err = a.SetMediaType(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.MediaType != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MediaFiles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MediaTypeID != x.MediaTypeID {
			t.Error("foreign key was wrong value", a.MediaTypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MediaTypeID))
		reflect.Indirect(reflect.ValueOf(&a.MediaTypeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MediaTypeID != x.MediaTypeID {
			t.Error("foreign key was wrong value", a.MediaTypeID, x.MediaTypeID)
		}
	}
}

func testMediaFilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
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

func testMediaFilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MediaFileSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMediaFilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MediaFiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mediaFileDBTypes = map[string]string{`MediaFileID`: `char`, `MediaTypeID`: `varchar`, `MediaFileBlocked`: `tinyint`, `MediaFileChecked`: `tinyint`}
	_                = bytes.MinRead
)

func testMediaFilesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mediaFilePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mediaFileAllColumns) == len(mediaFilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMediaFilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mediaFileAllColumns) == len(mediaFilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MediaFile{}
	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mediaFileDBTypes, true, mediaFilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mediaFileAllColumns, mediaFilePrimaryKeyColumns) {
		fields = mediaFileAllColumns
	} else {
		fields = strmangle.SetComplement(
			mediaFileAllColumns,
			mediaFilePrimaryKeyColumns,
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

	slice := MediaFileSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMediaFilesUpsert(t *testing.T) {
	t.Parallel()

	if len(mediaFileAllColumns) == len(mediaFilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMediaFileUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MediaFile{}
	if err = randomize.Struct(seed, &o, mediaFileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MediaFile: %s", err)
	}

	count, err := MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mediaFileDBTypes, false, mediaFilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MediaFile: %s", err)
	}

	count, err = MediaFiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}