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

func testRecordTests(t *testing.T) {
	t.Parallel()

	query := RecordTests()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRecordTestsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
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

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordTestsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RecordTests().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordTestsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecordTestSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordTestsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RecordTestExists(ctx, tx, o.RecordID)
	if err != nil {
		t.Errorf("Unable to check if RecordTest exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RecordTestExists to return true, but got false.")
	}
}

func testRecordTestsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	recordTestFound, err := FindRecordTest(ctx, tx, o.RecordID)
	if err != nil {
		t.Error(err)
	}

	if recordTestFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRecordTestsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RecordTests().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRecordTestsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RecordTests().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRecordTestsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	recordTestOne := &RecordTest{}
	recordTestTwo := &RecordTest{}
	if err = randomize.Struct(seed, recordTestOne, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}
	if err = randomize.Struct(seed, recordTestTwo, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recordTestOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recordTestTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RecordTests().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRecordTestsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordTestOne := &RecordTest{}
	recordTestTwo := &RecordTest{}
	if err = randomize.Struct(seed, recordTestOne, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}
	if err = randomize.Struct(seed, recordTestTwo, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recordTestOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recordTestTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func recordTestBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func recordTestAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RecordTest) error {
	*o = RecordTest{}
	return nil
}

func testRecordTestsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RecordTest{}
	o := &RecordTest{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, recordTestDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RecordTest object: %s", err)
	}

	AddRecordTestHook(boil.BeforeInsertHook, recordTestBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	recordTestBeforeInsertHooks = []RecordTestHook{}

	AddRecordTestHook(boil.AfterInsertHook, recordTestAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	recordTestAfterInsertHooks = []RecordTestHook{}

	AddRecordTestHook(boil.AfterSelectHook, recordTestAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	recordTestAfterSelectHooks = []RecordTestHook{}

	AddRecordTestHook(boil.BeforeUpdateHook, recordTestBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	recordTestBeforeUpdateHooks = []RecordTestHook{}

	AddRecordTestHook(boil.AfterUpdateHook, recordTestAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	recordTestAfterUpdateHooks = []RecordTestHook{}

	AddRecordTestHook(boil.BeforeDeleteHook, recordTestBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	recordTestBeforeDeleteHooks = []RecordTestHook{}

	AddRecordTestHook(boil.AfterDeleteHook, recordTestAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	recordTestAfterDeleteHooks = []RecordTestHook{}

	AddRecordTestHook(boil.BeforeUpsertHook, recordTestBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	recordTestBeforeUpsertHooks = []RecordTestHook{}

	AddRecordTestHook(boil.AfterUpsertHook, recordTestAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	recordTestAfterUpsertHooks = []RecordTestHook{}
}

func testRecordTestsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecordTestsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(recordTestColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecordTestToOneMemberUsingMember(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RecordTest
	var foreign Member

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, recordTestDBTypes, false, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, memberDBTypes, false, memberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Member struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MemberID = foreign.MemberID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Member().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.MemberID != foreign.MemberID {
		t.Errorf("want: %v, got %v", foreign.MemberID, check.MemberID)
	}

	ranAfterSelectHook := false
	AddMemberHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Member) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := RecordTestSlice{&local}
	if err = local.L.LoadMember(ctx, tx, false, (*[]*RecordTest)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Member = nil
	if err = local.L.LoadMember(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testRecordTestToOneMediaFileUsingMediaFile(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RecordTest
	var foreign MediaFile

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mediaFileDBTypes, false, mediaFileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MediaFile struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.MediaFileID, foreign.MediaFileID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.MediaFile().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.MediaFileID, foreign.MediaFileID) {
		t.Errorf("want: %v, got %v", foreign.MediaFileID, check.MediaFileID)
	}

	ranAfterSelectHook := false
	AddMediaFileHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *MediaFile) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := RecordTestSlice{&local}
	if err = local.L.LoadMediaFile(ctx, tx, false, (*[]*RecordTest)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MediaFile == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.MediaFile = nil
	if err = local.L.LoadMediaFile(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MediaFile == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testRecordTestToOneSetOpMemberUsingMember(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecordTest
	var b, c Member

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recordTestDBTypes, false, strmangle.SetComplement(recordTestPrimaryKeyColumns, recordTestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, memberDBTypes, false, strmangle.SetComplement(memberPrimaryKeyColumns, memberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, memberDBTypes, false, strmangle.SetComplement(memberPrimaryKeyColumns, memberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Member{&b, &c} {
		err = a.SetMember(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Member != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RecordTests[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MemberID != x.MemberID {
			t.Error("foreign key was wrong value", a.MemberID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MemberID))
		reflect.Indirect(reflect.ValueOf(&a.MemberID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MemberID != x.MemberID {
			t.Error("foreign key was wrong value", a.MemberID, x.MemberID)
		}
	}
}
func testRecordTestToOneSetOpMediaFileUsingMediaFile(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecordTest
	var b, c MediaFile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recordTestDBTypes, false, strmangle.SetComplement(recordTestPrimaryKeyColumns, recordTestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MediaFile{&b, &c} {
		err = a.SetMediaFile(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.MediaFile != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RecordTests[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.MediaFileID, x.MediaFileID) {
			t.Error("foreign key was wrong value", a.MediaFileID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MediaFileID))
		reflect.Indirect(reflect.ValueOf(&a.MediaFileID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.MediaFileID, x.MediaFileID) {
			t.Error("foreign key was wrong value", a.MediaFileID, x.MediaFileID)
		}
	}
}

func testRecordTestToOneRemoveOpMediaFileUsingMediaFile(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecordTest
	var b MediaFile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recordTestDBTypes, false, strmangle.SetComplement(recordTestPrimaryKeyColumns, recordTestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mediaFileDBTypes, false, strmangle.SetComplement(mediaFilePrimaryKeyColumns, mediaFileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetMediaFile(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveMediaFile(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.MediaFile().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.MediaFile != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.MediaFileID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.RecordTests) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRecordTestsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
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

func testRecordTestsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecordTestSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRecordTestsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RecordTests().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	recordTestDBTypes = map[string]string{`RecordID`: `int`, `MemberID`: `varchar`, `RecordComment`: `varchar`, `RecordCreatedAt`: `datetime`, `MediaFileID`: `char`, `RecordLocationY`: `float`, `RecordLocationX`: `float`, `RecordIsOpen`: `tinyint`, `RecordShareTo`: `varchar`, `RecordShareAccepted`: `tinyint`}
	_                 = bytes.MinRead
)

func testRecordTestsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(recordTestPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(recordTestAllColumns) == len(recordTestPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRecordTestsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(recordTestAllColumns) == len(recordTestPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RecordTest{}
	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recordTestDBTypes, true, recordTestPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(recordTestAllColumns, recordTestPrimaryKeyColumns) {
		fields = recordTestAllColumns
	} else {
		fields = strmangle.SetComplement(
			recordTestAllColumns,
			recordTestPrimaryKeyColumns,
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

	slice := RecordTestSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRecordTestsUpsert(t *testing.T) {
	t.Parallel()

	if len(recordTestAllColumns) == len(recordTestPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRecordTestUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RecordTest{}
	if err = randomize.Struct(seed, &o, recordTestDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RecordTest: %s", err)
	}

	count, err := RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, recordTestDBTypes, false, recordTestPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecordTest struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RecordTest: %s", err)
	}

	count, err = RecordTests().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}