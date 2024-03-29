// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// RecordKeyword is an object representing the database table.
type RecordKeyword struct {
	RecordKeywordID int    `boil:"record_keyword_id" json:"record_keyword_id" toml:"record_keyword_id" yaml:"record_keyword_id"`
	RecordID        int    `boil:"record_id" json:"record_id" toml:"record_id" yaml:"record_id"`
	RecordKeyword   string `boil:"record_keyword" json:"record_keyword" toml:"record_keyword" yaml:"record_keyword"`

	R *recordKeywordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recordKeywordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecordKeywordColumns = struct {
	RecordKeywordID string
	RecordID        string
	RecordKeyword   string
}{
	RecordKeywordID: "record_keyword_id",
	RecordID:        "record_id",
	RecordKeyword:   "record_keyword",
}

var RecordKeywordTableColumns = struct {
	RecordKeywordID string
	RecordID        string
	RecordKeyword   string
}{
	RecordKeywordID: "record_keyword.record_keyword_id",
	RecordID:        "record_keyword.record_id",
	RecordKeyword:   "record_keyword.record_keyword",
}

// Generated where

var RecordKeywordWhere = struct {
	RecordKeywordID whereHelperint
	RecordID        whereHelperint
	RecordKeyword   whereHelperstring
}{
	RecordKeywordID: whereHelperint{field: "`record_keyword`.`record_keyword_id`"},
	RecordID:        whereHelperint{field: "`record_keyword`.`record_id`"},
	RecordKeyword:   whereHelperstring{field: "`record_keyword`.`record_keyword`"},
}

// RecordKeywordRels is where relationship names are stored.
var RecordKeywordRels = struct {
	Record string
}{
	Record: "Record",
}

// recordKeywordR is where relationships are stored.
type recordKeywordR struct {
	Record *Record `boil:"Record" json:"Record" toml:"Record" yaml:"Record"`
}

// NewStruct creates a new relationship struct
func (*recordKeywordR) NewStruct() *recordKeywordR {
	return &recordKeywordR{}
}

func (r *recordKeywordR) GetRecord() *Record {
	if r == nil {
		return nil
	}
	return r.Record
}

// recordKeywordL is where Load methods for each relationship are stored.
type recordKeywordL struct{}

var (
	recordKeywordAllColumns            = []string{"record_keyword_id", "record_id", "record_keyword"}
	recordKeywordColumnsWithoutDefault = []string{"record_id", "record_keyword"}
	recordKeywordColumnsWithDefault    = []string{"record_keyword_id"}
	recordKeywordPrimaryKeyColumns     = []string{"record_keyword_id"}
	recordKeywordGeneratedColumns      = []string{}
)

type (
	// RecordKeywordSlice is an alias for a slice of pointers to RecordKeyword.
	// This should almost always be used instead of []RecordKeyword.
	RecordKeywordSlice []*RecordKeyword
	// RecordKeywordHook is the signature for custom RecordKeyword hook methods
	RecordKeywordHook func(context.Context, boil.ContextExecutor, *RecordKeyword) error

	recordKeywordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recordKeywordType                 = reflect.TypeOf(&RecordKeyword{})
	recordKeywordMapping              = queries.MakeStructMapping(recordKeywordType)
	recordKeywordPrimaryKeyMapping, _ = queries.BindMapping(recordKeywordType, recordKeywordMapping, recordKeywordPrimaryKeyColumns)
	recordKeywordInsertCacheMut       sync.RWMutex
	recordKeywordInsertCache          = make(map[string]insertCache)
	recordKeywordUpdateCacheMut       sync.RWMutex
	recordKeywordUpdateCache          = make(map[string]updateCache)
	recordKeywordUpsertCacheMut       sync.RWMutex
	recordKeywordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recordKeywordAfterSelectMu sync.Mutex
var recordKeywordAfterSelectHooks []RecordKeywordHook

var recordKeywordBeforeInsertMu sync.Mutex
var recordKeywordBeforeInsertHooks []RecordKeywordHook
var recordKeywordAfterInsertMu sync.Mutex
var recordKeywordAfterInsertHooks []RecordKeywordHook

var recordKeywordBeforeUpdateMu sync.Mutex
var recordKeywordBeforeUpdateHooks []RecordKeywordHook
var recordKeywordAfterUpdateMu sync.Mutex
var recordKeywordAfterUpdateHooks []RecordKeywordHook

var recordKeywordBeforeDeleteMu sync.Mutex
var recordKeywordBeforeDeleteHooks []RecordKeywordHook
var recordKeywordAfterDeleteMu sync.Mutex
var recordKeywordAfterDeleteHooks []RecordKeywordHook

var recordKeywordBeforeUpsertMu sync.Mutex
var recordKeywordBeforeUpsertHooks []RecordKeywordHook
var recordKeywordAfterUpsertMu sync.Mutex
var recordKeywordAfterUpsertHooks []RecordKeywordHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecordKeyword) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecordKeyword) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecordKeyword) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecordKeyword) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecordKeyword) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecordKeyword) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecordKeyword) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecordKeyword) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecordKeyword) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recordKeywordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecordKeywordHook registers your hook function for all future operations.
func AddRecordKeywordHook(hookPoint boil.HookPoint, recordKeywordHook RecordKeywordHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		recordKeywordAfterSelectMu.Lock()
		recordKeywordAfterSelectHooks = append(recordKeywordAfterSelectHooks, recordKeywordHook)
		recordKeywordAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		recordKeywordBeforeInsertMu.Lock()
		recordKeywordBeforeInsertHooks = append(recordKeywordBeforeInsertHooks, recordKeywordHook)
		recordKeywordBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		recordKeywordAfterInsertMu.Lock()
		recordKeywordAfterInsertHooks = append(recordKeywordAfterInsertHooks, recordKeywordHook)
		recordKeywordAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		recordKeywordBeforeUpdateMu.Lock()
		recordKeywordBeforeUpdateHooks = append(recordKeywordBeforeUpdateHooks, recordKeywordHook)
		recordKeywordBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		recordKeywordAfterUpdateMu.Lock()
		recordKeywordAfterUpdateHooks = append(recordKeywordAfterUpdateHooks, recordKeywordHook)
		recordKeywordAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		recordKeywordBeforeDeleteMu.Lock()
		recordKeywordBeforeDeleteHooks = append(recordKeywordBeforeDeleteHooks, recordKeywordHook)
		recordKeywordBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		recordKeywordAfterDeleteMu.Lock()
		recordKeywordAfterDeleteHooks = append(recordKeywordAfterDeleteHooks, recordKeywordHook)
		recordKeywordAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		recordKeywordBeforeUpsertMu.Lock()
		recordKeywordBeforeUpsertHooks = append(recordKeywordBeforeUpsertHooks, recordKeywordHook)
		recordKeywordBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		recordKeywordAfterUpsertMu.Lock()
		recordKeywordAfterUpsertHooks = append(recordKeywordAfterUpsertHooks, recordKeywordHook)
		recordKeywordAfterUpsertMu.Unlock()
	}
}

// OneG returns a single recordKeyword record from the query using the global executor.
func (q recordKeywordQuery) OneG(ctx context.Context) (*RecordKeyword, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single recordKeyword record from the query.
func (q recordKeywordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecordKeyword, error) {
	o := &RecordKeyword{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for record_keyword")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all RecordKeyword records from the query using the global executor.
func (q recordKeywordQuery) AllG(ctx context.Context) (RecordKeywordSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all RecordKeyword records from the query.
func (q recordKeywordQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecordKeywordSlice, error) {
	var o []*RecordKeyword

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecordKeyword slice")
	}

	if len(recordKeywordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all RecordKeyword records in the query using the global executor
func (q recordKeywordQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all RecordKeyword records in the query.
func (q recordKeywordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count record_keyword rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q recordKeywordQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q recordKeywordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if record_keyword exists")
	}

	return count > 0, nil
}

// Record pointed to by the foreign key.
func (o *RecordKeyword) Record(mods ...qm.QueryMod) recordQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`record_id` = ?", o.RecordID),
	}

	queryMods = append(queryMods, mods...)

	return Records(queryMods...)
}

// LoadRecord allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recordKeywordL) LoadRecord(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecordKeyword interface{}, mods queries.Applicator) error {
	var slice []*RecordKeyword
	var object *RecordKeyword

	if singular {
		var ok bool
		object, ok = maybeRecordKeyword.(*RecordKeyword)
		if !ok {
			object = new(RecordKeyword)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRecordKeyword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRecordKeyword))
			}
		}
	} else {
		s, ok := maybeRecordKeyword.(*[]*RecordKeyword)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRecordKeyword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRecordKeyword))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &recordKeywordR{}
		}
		args[object.RecordID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recordKeywordR{}
			}

			args[obj.RecordID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`record`),
		qm.WhereIn(`record.record_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Record")
	}

	var resultSlice []*Record
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Record")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for record")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for record")
	}

	if len(recordAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Record = foreign
		if foreign.R == nil {
			foreign.R = &recordR{}
		}
		foreign.R.RecordKeywords = append(foreign.R.RecordKeywords, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RecordID == foreign.RecordID {
				local.R.Record = foreign
				if foreign.R == nil {
					foreign.R = &recordR{}
				}
				foreign.R.RecordKeywords = append(foreign.R.RecordKeywords, local)
				break
			}
		}
	}

	return nil
}

// SetRecordG of the recordKeyword to the related item.
// Sets o.R.Record to related.
// Adds o to related.R.RecordKeywords.
// Uses the global database handle.
func (o *RecordKeyword) SetRecordG(ctx context.Context, insert bool, related *Record) error {
	return o.SetRecord(ctx, boil.GetContextDB(), insert, related)
}

// SetRecord of the recordKeyword to the related item.
// Sets o.R.Record to related.
// Adds o to related.R.RecordKeywords.
func (o *RecordKeyword) SetRecord(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Record) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `record_keyword` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"record_id"}),
		strmangle.WhereClause("`", "`", 0, recordKeywordPrimaryKeyColumns),
	)
	values := []interface{}{related.RecordID, o.RecordKeywordID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RecordID = related.RecordID
	if o.R == nil {
		o.R = &recordKeywordR{
			Record: related,
		}
	} else {
		o.R.Record = related
	}

	if related.R == nil {
		related.R = &recordR{
			RecordKeywords: RecordKeywordSlice{o},
		}
	} else {
		related.R.RecordKeywords = append(related.R.RecordKeywords, o)
	}

	return nil
}

// RecordKeywords retrieves all the records using an executor.
func RecordKeywords(mods ...qm.QueryMod) recordKeywordQuery {
	mods = append(mods, qm.From("`record_keyword`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`record_keyword`.*"})
	}

	return recordKeywordQuery{q}
}

// FindRecordKeywordG retrieves a single record by ID.
func FindRecordKeywordG(ctx context.Context, recordKeywordID int, selectCols ...string) (*RecordKeyword, error) {
	return FindRecordKeyword(ctx, boil.GetContextDB(), recordKeywordID, selectCols...)
}

// FindRecordKeyword retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecordKeyword(ctx context.Context, exec boil.ContextExecutor, recordKeywordID int, selectCols ...string) (*RecordKeyword, error) {
	recordKeywordObj := &RecordKeyword{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `record_keyword` where `record_keyword_id`=?", sel,
	)

	q := queries.Raw(query, recordKeywordID)

	err := q.Bind(ctx, exec, recordKeywordObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from record_keyword")
	}

	if err = recordKeywordObj.doAfterSelectHooks(ctx, exec); err != nil {
		return recordKeywordObj, err
	}

	return recordKeywordObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *RecordKeyword) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecordKeyword) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no record_keyword provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recordKeywordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recordKeywordInsertCacheMut.RLock()
	cache, cached := recordKeywordInsertCache[key]
	recordKeywordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recordKeywordAllColumns,
			recordKeywordColumnsWithDefault,
			recordKeywordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recordKeywordType, recordKeywordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recordKeywordType, recordKeywordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `record_keyword` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `record_keyword` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `record_keyword` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, recordKeywordPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into record_keyword")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.RecordKeywordID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == recordKeywordMapping["record_keyword_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.RecordKeywordID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for record_keyword")
	}

CacheNoHooks:
	if !cached {
		recordKeywordInsertCacheMut.Lock()
		recordKeywordInsertCache[key] = cache
		recordKeywordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single RecordKeyword record using the global executor.
// See Update for more documentation.
func (o *RecordKeyword) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the RecordKeyword.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecordKeyword) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recordKeywordUpdateCacheMut.RLock()
	cache, cached := recordKeywordUpdateCache[key]
	recordKeywordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recordKeywordAllColumns,
			recordKeywordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update record_keyword, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `record_keyword` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, recordKeywordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recordKeywordType, recordKeywordMapping, append(wl, recordKeywordPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update record_keyword row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for record_keyword")
	}

	if !cached {
		recordKeywordUpdateCacheMut.Lock()
		recordKeywordUpdateCache[key] = cache
		recordKeywordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q recordKeywordQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q recordKeywordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for record_keyword")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for record_keyword")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RecordKeywordSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecordKeywordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recordKeywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `record_keyword` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, recordKeywordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recordKeyword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recordKeyword")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *RecordKeyword) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLRecordKeywordUniqueColumns = []string{
	"record_keyword_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecordKeyword) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no record_keyword provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recordKeywordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRecordKeywordUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	recordKeywordUpsertCacheMut.RLock()
	cache, cached := recordKeywordUpsertCache[key]
	recordKeywordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			recordKeywordAllColumns,
			recordKeywordColumnsWithDefault,
			recordKeywordColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			recordKeywordAllColumns,
			recordKeywordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert record_keyword, could not build update column list")
		}

		ret := strmangle.SetComplement(recordKeywordAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`record_keyword`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `record_keyword` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(recordKeywordType, recordKeywordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recordKeywordType, recordKeywordMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for record_keyword")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.RecordKeywordID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == recordKeywordMapping["record_keyword_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(recordKeywordType, recordKeywordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for record_keyword")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for record_keyword")
	}

CacheNoHooks:
	if !cached {
		recordKeywordUpsertCacheMut.Lock()
		recordKeywordUpsertCache[key] = cache
		recordKeywordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single RecordKeyword record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *RecordKeyword) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single RecordKeyword record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecordKeyword) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecordKeyword provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recordKeywordPrimaryKeyMapping)
	sql := "DELETE FROM `record_keyword` WHERE `record_keyword_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from record_keyword")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for record_keyword")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q recordKeywordQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q recordKeywordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recordKeywordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from record_keyword")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for record_keyword")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o RecordKeywordSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecordKeywordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recordKeywordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recordKeywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `record_keyword` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, recordKeywordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recordKeyword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for record_keyword")
	}

	if len(recordKeywordAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *RecordKeyword) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no RecordKeyword provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RecordKeyword) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecordKeyword(ctx, exec, o.RecordKeywordID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecordKeywordSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty RecordKeywordSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecordKeywordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecordKeywordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recordKeywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `record_keyword`.* FROM `record_keyword` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, recordKeywordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecordKeywordSlice")
	}

	*o = slice

	return nil
}

// RecordKeywordExistsG checks if the RecordKeyword row exists.
func RecordKeywordExistsG(ctx context.Context, recordKeywordID int) (bool, error) {
	return RecordKeywordExists(ctx, boil.GetContextDB(), recordKeywordID)
}

// RecordKeywordExists checks if the RecordKeyword row exists.
func RecordKeywordExists(ctx context.Context, exec boil.ContextExecutor, recordKeywordID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `record_keyword` where `record_keyword_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, recordKeywordID)
	}
	row := exec.QueryRowContext(ctx, sql, recordKeywordID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if record_keyword exists")
	}

	return exists, nil
}

// Exists checks if the RecordKeyword row exists.
func (o *RecordKeyword) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RecordKeywordExists(ctx, exec, o.RecordKeywordID)
}
