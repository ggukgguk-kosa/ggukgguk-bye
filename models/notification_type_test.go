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

func testNotificationTypes(t *testing.T) {
	t.Parallel()

	query := NotificationTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNotificationTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
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

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NotificationTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NotificationTypeExists(ctx, tx, o.NotificationTypeID)
	if err != nil {
		t.Errorf("Unable to check if NotificationType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NotificationTypeExists to return true, but got false.")
	}
}

func testNotificationTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	notificationTypeFound, err := FindNotificationType(ctx, tx, o.NotificationTypeID)
	if err != nil {
		t.Error(err)
	}

	if notificationTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNotificationTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NotificationTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNotificationTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NotificationTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNotificationTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	notificationTypeOne := &NotificationType{}
	notificationTypeTwo := &NotificationType{}
	if err = randomize.Struct(seed, notificationTypeOne, notificationTypeDBTypes, false, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}
	if err = randomize.Struct(seed, notificationTypeTwo, notificationTypeDBTypes, false, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = notificationTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = notificationTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NotificationTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNotificationTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notificationTypeOne := &NotificationType{}
	notificationTypeTwo := &NotificationType{}
	if err = randomize.Struct(seed, notificationTypeOne, notificationTypeDBTypes, false, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}
	if err = randomize.Struct(seed, notificationTypeTwo, notificationTypeDBTypes, false, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = notificationTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = notificationTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func notificationTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func notificationTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NotificationType) error {
	*o = NotificationType{}
	return nil
}

func testNotificationTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &NotificationType{}
	o := &NotificationType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NotificationType object: %s", err)
	}

	AddNotificationTypeHook(boil.BeforeInsertHook, notificationTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	notificationTypeBeforeInsertHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.AfterInsertHook, notificationTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	notificationTypeAfterInsertHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.AfterSelectHook, notificationTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	notificationTypeAfterSelectHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.BeforeUpdateHook, notificationTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	notificationTypeBeforeUpdateHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.AfterUpdateHook, notificationTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	notificationTypeAfterUpdateHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.BeforeDeleteHook, notificationTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	notificationTypeBeforeDeleteHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.AfterDeleteHook, notificationTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	notificationTypeAfterDeleteHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.BeforeUpsertHook, notificationTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	notificationTypeBeforeUpsertHooks = []NotificationTypeHook{}

	AddNotificationTypeHook(boil.AfterUpsertHook, notificationTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	notificationTypeAfterUpsertHooks = []NotificationTypeHook{}
}

func testNotificationTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotificationTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(notificationTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotificationTypeToManyNotifications(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NotificationType
	var b, c Notification

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.NotificationTypeID = a.NotificationTypeID
	c.NotificationTypeID = a.NotificationTypeID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Notifications().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.NotificationTypeID == b.NotificationTypeID {
			bFound = true
		}
		if v.NotificationTypeID == c.NotificationTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := NotificationTypeSlice{&a}
	if err = a.L.LoadNotifications(ctx, tx, false, (*[]*NotificationType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Notifications); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Notifications = nil
	if err = a.L.LoadNotifications(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Notifications); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testNotificationTypeToManyAddOpNotifications(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NotificationType
	var b, c, d, e Notification

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, notificationTypeDBTypes, false, strmangle.SetComplement(notificationTypePrimaryKeyColumns, notificationTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Notification{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, notificationDBTypes, false, strmangle.SetComplement(notificationPrimaryKeyColumns, notificationColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Notification{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNotifications(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.NotificationTypeID != first.NotificationTypeID {
			t.Error("foreign key was wrong value", a.NotificationTypeID, first.NotificationTypeID)
		}
		if a.NotificationTypeID != second.NotificationTypeID {
			t.Error("foreign key was wrong value", a.NotificationTypeID, second.NotificationTypeID)
		}

		if first.R.NotificationType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.NotificationType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Notifications[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Notifications[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Notifications().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testNotificationTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
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

func testNotificationTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNotificationTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NotificationTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	notificationTypeDBTypes = map[string]string{`NotificationTypeID`: `varchar`, `NotificationTypeComment`: `varchar`}
	_                       = bytes.MinRead
)

func testNotificationTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(notificationTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(notificationTypeAllColumns) == len(notificationTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNotificationTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(notificationTypeAllColumns) == len(notificationTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NotificationType{}
	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, notificationTypeDBTypes, true, notificationTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(notificationTypeAllColumns, notificationTypePrimaryKeyColumns) {
		fields = notificationTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			notificationTypeAllColumns,
			notificationTypePrimaryKeyColumns,
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

	slice := NotificationTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNotificationTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(notificationTypeAllColumns) == len(notificationTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLNotificationTypeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NotificationType{}
	if err = randomize.Struct(seed, &o, notificationTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NotificationType: %s", err)
	}

	count, err := NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, notificationTypeDBTypes, false, notificationTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NotificationType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NotificationType: %s", err)
	}

	count, err = NotificationTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
