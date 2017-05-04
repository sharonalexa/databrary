// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// Funder is an object representing the database table.
type Funder struct {
	FundrefID int64  `db:"fundref_id" json:"funder_fundref_id"`
	Name      string `db:"name" json:"funder_name"`

	R *funderR `db:"-" json:"-"`
	L funderL  `db:"-" json:"-"`
}

// funderR is where relationships are stored.
type funderR struct {
	VolumeFundings VolumeFundingSlice
}

// funderL is where Load methods for each relationship are stored.
type funderL struct{}

var (
	funderColumns               = []string{"fundref_id", "name"}
	funderColumnsWithoutDefault = []string{"fundref_id", "name"}
	funderColumnsWithDefault    = []string{}
	funderColumnsWithCustom     = []string{}

	funderPrimaryKeyColumns = []string{"fundref_id"}
)

type (
	// FunderSlice is an alias for a slice of pointers to Funder.
	// This should generally be used opposed to []Funder.
	FunderSlice []*Funder
	// FunderHook is the signature for custom Funder hook methods
	FunderHook func(boil.Executor, *Funder) error

	funderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	funderType    = reflect.TypeOf(&Funder{})
	funderMapping = queries.MakeStructMapping(funderType)

	funderPrimaryKeyMapping, _ = queries.BindMapping(funderType, funderMapping, funderPrimaryKeyColumns)

	funderInsertCacheMut sync.RWMutex
	funderInsertCache    = make(map[string]insertCache)
	funderUpdateCacheMut sync.RWMutex
	funderUpdateCache    = make(map[string]updateCache)
	funderUpsertCacheMut sync.RWMutex
	funderUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var funderBeforeInsertHooks []FunderHook
var funderBeforeUpdateHooks []FunderHook
var funderBeforeDeleteHooks []FunderHook
var funderBeforeUpsertHooks []FunderHook

var funderAfterInsertHooks []FunderHook
var funderAfterSelectHooks []FunderHook
var funderAfterUpdateHooks []FunderHook
var funderAfterDeleteHooks []FunderHook
var funderAfterUpsertHooks []FunderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Funder) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range funderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Funder) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range funderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Funder) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range funderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Funder) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range funderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Funder) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range funderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Funder) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range funderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Funder) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range funderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Funder) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range funderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Funder) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range funderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFunderHook registers your hook function for all future operations.
func AddFunderHook(hookPoint boil.HookPoint, funderHook FunderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		funderBeforeInsertHooks = append(funderBeforeInsertHooks, funderHook)
	case boil.BeforeUpdateHook:
		funderBeforeUpdateHooks = append(funderBeforeUpdateHooks, funderHook)
	case boil.BeforeDeleteHook:
		funderBeforeDeleteHooks = append(funderBeforeDeleteHooks, funderHook)
	case boil.BeforeUpsertHook:
		funderBeforeUpsertHooks = append(funderBeforeUpsertHooks, funderHook)
	case boil.AfterInsertHook:
		funderAfterInsertHooks = append(funderAfterInsertHooks, funderHook)
	case boil.AfterSelectHook:
		funderAfterSelectHooks = append(funderAfterSelectHooks, funderHook)
	case boil.AfterUpdateHook:
		funderAfterUpdateHooks = append(funderAfterUpdateHooks, funderHook)
	case boil.AfterDeleteHook:
		funderAfterDeleteHooks = append(funderAfterDeleteHooks, funderHook)
	case boil.AfterUpsertHook:
		funderAfterUpsertHooks = append(funderAfterUpsertHooks, funderHook)
	}
}

// OneP returns a single funder record from the query, and panics on error.
func (q funderQuery) OneP() *Funder {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single funder record from the query.
func (q funderQuery) One() (*Funder, error) {
	o := &Funder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for funder")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Funder records from the query, and panics on error.
func (q funderQuery) AllP() FunderSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Funder records from the query.
func (q funderQuery) All() (FunderSlice, error) {
	var o FunderSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Funder slice")
	}

	if len(funderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Funder records in the query, and panics on error.
func (q funderQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Funder records in the query.
func (q funderQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count funder rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q funderQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q funderQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if funder exists")
	}

	return count > 0, nil
}

// VolumeFundingsG retrieves all the volume_funding's volume funding.
func (o *Funder) VolumeFundingsG(mods ...qm.QueryMod) volumeFundingQuery {
	return o.VolumeFundingsByFk(boil.GetDB(), mods...)
}

// VolumeFundings retrieves all the volume_funding's volume funding with an executor.
func (o *Funder) VolumeFundingsByFk(exec boil.Executor, mods ...qm.QueryMod) volumeFundingQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"funder\"=?", o.FundrefID),
	)

	query := VolumeFundings(exec, queryMods...)
	queries.SetFrom(query.Query, "\"volume_funding\" as \"a\"")
	return query
}

// LoadVolumeFundings allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (funderL) LoadVolumeFundings(e boil.Executor, singular bool, maybeFunder interface{}) error {
	var slice []*Funder
	var object *Funder

	count := 1
	if singular {
		object = maybeFunder.(*Funder)
	} else {
		slice = *maybeFunder.(*FunderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &funderR{}
		}
		args[0] = object.FundrefID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &funderR{}
			}
			args[i] = obj.FundrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"volume_funding\" where \"funder\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load volume_funding")
	}
	defer results.Close()

	var resultSlice []*VolumeFunding
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice volume_funding")
	}

	if len(volumeFundingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.VolumeFundings = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FundrefID == foreign.Funder {
				local.R.VolumeFundings = append(local.R.VolumeFundings, foreign)
				break
			}
		}
	}

	return nil
}

// AddVolumeFundingsG adds the given related objects to the existing relationships
// of the funder, optionally inserting them as new records.
// Appends related to o.R.VolumeFundings.
// Sets related.R.Funder appropriately.
// Uses the global database handle.
func (o *Funder) AddVolumeFundingsG(insert bool, related ...*VolumeFunding) error {
	return o.AddVolumeFundings(boil.GetDB(), insert, related...)
}

// AddVolumeFundingsP adds the given related objects to the existing relationships
// of the funder, optionally inserting them as new records.
// Appends related to o.R.VolumeFundings.
// Sets related.R.Funder appropriately.
// Panics on error.
func (o *Funder) AddVolumeFundingsP(exec boil.Executor, insert bool, related ...*VolumeFunding) {
	if err := o.AddVolumeFundings(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddVolumeFundingsGP adds the given related objects to the existing relationships
// of the funder, optionally inserting them as new records.
// Appends related to o.R.VolumeFundings.
// Sets related.R.Funder appropriately.
// Uses the global database handle and panics on error.
func (o *Funder) AddVolumeFundingsGP(insert bool, related ...*VolumeFunding) {
	if err := o.AddVolumeFundings(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddVolumeFundings adds the given related objects to the existing relationships
// of the funder, optionally inserting them as new records.
// Appends related to o.R.VolumeFundings.
// Sets related.R.Funder appropriately.
func (o *Funder) AddVolumeFundings(exec boil.Executor, insert bool, related ...*VolumeFunding) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.Funder = o.FundrefID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"volume_funding\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"funder"}),
				strmangle.WhereClause("\"", "\"", 2, volumeFundingPrimaryKeyColumns),
			)
			values := []interface{}{o.FundrefID, rel.Volume, rel.Funder}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.Funder = o.FundrefID
		}
	}

	if o.R == nil {
		o.R = &funderR{
			VolumeFundings: related,
		}
	} else {
		o.R.VolumeFundings = append(o.R.VolumeFundings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &volumeFundingR{
				Funder: o,
			}
		} else {
			rel.R.Funder = o
		}
	}
	return nil
}

// FundersG retrieves all records.
func FundersG(mods ...qm.QueryMod) funderQuery {
	return Funders(boil.GetDB(), mods...)
}

// Funders retrieves all the records using an executor.
func Funders(exec boil.Executor, mods ...qm.QueryMod) funderQuery {
	mods = append(mods, qm.From("\"funder\""))
	return funderQuery{NewQuery(exec, mods...)}
}

// FindFunderG retrieves a single record by ID.
func FindFunderG(fundrefID int64, selectCols ...string) (*Funder, error) {
	return FindFunder(boil.GetDB(), fundrefID, selectCols...)
}

// FindFunderGP retrieves a single record by ID, and panics on error.
func FindFunderGP(fundrefID int64, selectCols ...string) *Funder {
	retobj, err := FindFunder(boil.GetDB(), fundrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFunder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFunder(exec boil.Executor, fundrefID int64, selectCols ...string) (*Funder, error) {
	funderObj := &Funder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"funder\" where \"fundref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, fundrefID)

	err := q.Bind(funderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from funder")
	}

	return funderObj, nil
}

// FindFunderP retrieves a single record by ID with an executor, and panics on error.
func FindFunderP(exec boil.Executor, fundrefID int64, selectCols ...string) *Funder {
	retobj, err := FindFunder(exec, fundrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Funder) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Funder) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Funder) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Funder) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no funder provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(funderColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	funderInsertCacheMut.RLock()
	cache, cached := funderInsertCache[key]
	funderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			funderColumns,
			funderColumnsWithDefault,
			funderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(funderType, funderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(funderType, funderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"funder\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"funder\" DEFAULT VALUES"
		}

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into funder")
	}

	if !cached {
		funderInsertCacheMut.Lock()
		funderInsertCache[key] = cache
		funderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Funder record. See Update for
// whitelist behavior description.
func (o *Funder) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Funder record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Funder) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Funder, and panics on error.
// See Update for whitelist behavior description.
func (o *Funder) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Funder.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Funder) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	funderUpdateCacheMut.RLock()
	cache, cached := funderUpdateCache[key]
	funderUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(funderColumns, funderPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update funder, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"funder\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, funderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(funderType, funderMapping, append(wl, funderPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update funder row")
	}

	if !cached {
		funderUpdateCacheMut.Lock()
		funderUpdateCache[key] = cache
		funderUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q funderQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q funderQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for funder")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FunderSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FunderSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FunderSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FunderSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), funderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"UPDATE \"funder\" SET %s WHERE (\"fundref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(funderPrimaryKeyColumns), len(colNames)+1, len(funderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in funder slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Funder) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Funder) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Funder) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Funder) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no funder provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(funderColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	funderUpsertCacheMut.RLock()
	cache, cached := funderUpsertCache[key]
	funderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			funderColumns,
			funderColumnsWithDefault,
			funderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			funderColumns,
			funderPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert funder, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(funderPrimaryKeyColumns))
			copy(conflict, funderPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"funder\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(funderType, funderMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(funderType, funderMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert funder")
	}

	if !cached {
		funderUpsertCacheMut.Lock()
		funderUpsertCache[key] = cache
		funderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Funder record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Funder) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Funder record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Funder) DeleteG() error {
	if o == nil {
		return errors.New("models: no Funder provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Funder record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Funder) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Funder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Funder) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Funder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), funderPrimaryKeyMapping)
	query := "DELETE FROM \"funder\" WHERE \"fundref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from funder")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q funderQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q funderQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no funderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from funder")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FunderSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FunderSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Funder slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FunderSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FunderSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Funder slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(funderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), funderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"DELETE FROM \"funder\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, funderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(funderPrimaryKeyColumns), 1, len(funderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from funder slice")
	}

	if len(funderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Funder) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Funder) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Funder) ReloadG() error {
	if o == nil {
		return errors.New("models: no Funder provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Funder) Reload(exec boil.Executor) error {
	ret, err := FindFunder(exec, o.FundrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FunderSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FunderSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FunderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FunderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FunderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	funders := FunderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), funderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"SELECT \"funder\".* FROM \"funder\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, funderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(funderPrimaryKeyColumns), 1, len(funderPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, query, args...)

	err := q.Bind(&funders)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FunderSlice")
	}

	*o = funders

	return nil
}

// FunderExists checks if the Funder row exists.
func FunderExists(exec boil.Executor, fundrefID int64) (bool, error) {
	var exists bool

	query := "select exists(select 1 from \"funder\" where \"fundref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, fundrefID)
	}

	row := exec.QueryRow(query, fundrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if funder exists")
	}

	return exists, nil
}

// FunderExistsG checks if the Funder row exists.
func FunderExistsG(fundrefID int64) (bool, error) {
	return FunderExists(boil.GetDB(), fundrefID)
}

// FunderExistsGP checks if the Funder row exists. Panics on error.
func FunderExistsGP(fundrefID int64) bool {
	e, err := FunderExists(boil.GetDB(), fundrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FunderExistsP checks if the Funder row exists. Panics on error.
func FunderExistsP(exec boil.Executor, fundrefID int64) bool {
	e, err := FunderExists(exec, fundrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
