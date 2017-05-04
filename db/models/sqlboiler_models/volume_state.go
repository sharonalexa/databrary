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
	"github.com/vattle/sqlboiler/types"
)

// VolumeState is an object representing the database table.
type VolumeState struct {
	Volume int        `db:"volume" json:"volumeState_volume"`
	Key    string     `db:"key" json:"volumeState_key"`
	Value  types.JSON `db:"value" json:"volumeState_value"`
	Public bool       `db:"public" json:"volumeState_public"`

	R *volumeStateR `db:"-" json:"-"`
	L volumeStateL  `db:"-" json:"-"`
}

// volumeStateR is where relationships are stored.
type volumeStateR struct {
	Volume *Volume
}

// volumeStateL is where Load methods for each relationship are stored.
type volumeStateL struct{}

var (
	volumeStateColumns               = []string{"volume", "key", "value", "public"}
	volumeStateColumnsWithoutDefault = []string{"volume", "key", "value", "public"}
	volumeStateColumnsWithDefault    = []string{}
	volumeStateColumnsWithCustom     = []string{}

	volumeStatePrimaryKeyColumns = []string{"volume", "key"}
)

type (
	// VolumeStateSlice is an alias for a slice of pointers to VolumeState.
	// This should generally be used opposed to []VolumeState.
	VolumeStateSlice []*VolumeState
	// VolumeStateHook is the signature for custom VolumeState hook methods
	VolumeStateHook func(boil.Executor, *VolumeState) error

	volumeStateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	volumeStateType    = reflect.TypeOf(&VolumeState{})
	volumeStateMapping = queries.MakeStructMapping(volumeStateType)

	volumeStatePrimaryKeyMapping, _ = queries.BindMapping(volumeStateType, volumeStateMapping, volumeStatePrimaryKeyColumns)

	volumeStateInsertCacheMut sync.RWMutex
	volumeStateInsertCache    = make(map[string]insertCache)
	volumeStateUpdateCacheMut sync.RWMutex
	volumeStateUpdateCache    = make(map[string]updateCache)
	volumeStateUpsertCacheMut sync.RWMutex
	volumeStateUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var volumeStateBeforeInsertHooks []VolumeStateHook
var volumeStateBeforeUpdateHooks []VolumeStateHook
var volumeStateBeforeDeleteHooks []VolumeStateHook
var volumeStateBeforeUpsertHooks []VolumeStateHook

var volumeStateAfterInsertHooks []VolumeStateHook
var volumeStateAfterSelectHooks []VolumeStateHook
var volumeStateAfterUpdateHooks []VolumeStateHook
var volumeStateAfterDeleteHooks []VolumeStateHook
var volumeStateAfterUpsertHooks []VolumeStateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VolumeState) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VolumeState) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VolumeState) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VolumeState) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VolumeState) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VolumeState) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VolumeState) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VolumeState) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VolumeState) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeStateAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVolumeStateHook registers your hook function for all future operations.
func AddVolumeStateHook(hookPoint boil.HookPoint, volumeStateHook VolumeStateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		volumeStateBeforeInsertHooks = append(volumeStateBeforeInsertHooks, volumeStateHook)
	case boil.BeforeUpdateHook:
		volumeStateBeforeUpdateHooks = append(volumeStateBeforeUpdateHooks, volumeStateHook)
	case boil.BeforeDeleteHook:
		volumeStateBeforeDeleteHooks = append(volumeStateBeforeDeleteHooks, volumeStateHook)
	case boil.BeforeUpsertHook:
		volumeStateBeforeUpsertHooks = append(volumeStateBeforeUpsertHooks, volumeStateHook)
	case boil.AfterInsertHook:
		volumeStateAfterInsertHooks = append(volumeStateAfterInsertHooks, volumeStateHook)
	case boil.AfterSelectHook:
		volumeStateAfterSelectHooks = append(volumeStateAfterSelectHooks, volumeStateHook)
	case boil.AfterUpdateHook:
		volumeStateAfterUpdateHooks = append(volumeStateAfterUpdateHooks, volumeStateHook)
	case boil.AfterDeleteHook:
		volumeStateAfterDeleteHooks = append(volumeStateAfterDeleteHooks, volumeStateHook)
	case boil.AfterUpsertHook:
		volumeStateAfterUpsertHooks = append(volumeStateAfterUpsertHooks, volumeStateHook)
	}
}

// OneP returns a single volumeState record from the query, and panics on error.
func (q volumeStateQuery) OneP() *VolumeState {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single volumeState record from the query.
func (q volumeStateQuery) One() (*VolumeState, error) {
	o := &VolumeState{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for volume_state")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all VolumeState records from the query, and panics on error.
func (q volumeStateQuery) AllP() VolumeStateSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all VolumeState records from the query.
func (q volumeStateQuery) All() (VolumeStateSlice, error) {
	var o VolumeStateSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VolumeState slice")
	}

	if len(volumeStateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all VolumeState records in the query, and panics on error.
func (q volumeStateQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all VolumeState records in the query.
func (q volumeStateQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count volume_state rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q volumeStateQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q volumeStateQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if volume_state exists")
	}

	return count > 0, nil
}

// VolumeG pointed to by the foreign key.
func (o *VolumeState) VolumeG(mods ...qm.QueryMod) volumeQuery {
	return o.VolumeByFk(boil.GetDB(), mods...)
}

// Volume pointed to by the foreign key.
func (o *VolumeState) VolumeByFk(exec boil.Executor, mods ...qm.QueryMod) volumeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.Volume),
	}

	queryMods = append(queryMods, mods...)

	query := Volumes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"volume\"")

	return query
}

// LoadVolume allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (volumeStateL) LoadVolume(e boil.Executor, singular bool, maybeVolumeState interface{}) error {
	var slice []*VolumeState
	var object *VolumeState

	count := 1
	if singular {
		object = maybeVolumeState.(*VolumeState)
	} else {
		slice = *maybeVolumeState.(*VolumeStateSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &volumeStateR{}
		}
		args[0] = object.Volume
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &volumeStateR{}
			}
			args[i] = obj.Volume
		}
	}

	query := fmt.Sprintf(
		"select * from \"volume\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Volume")
	}
	defer results.Close()

	var resultSlice []*Volume
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Volume")
	}

	if len(volumeStateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Volume = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Volume == foreign.ID {
				local.R.Volume = foreign
				break
			}
		}
	}

	return nil
}

// SetVolumeG of the volume_state to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeStates.
// Uses the global database handle.
func (o *VolumeState) SetVolumeG(insert bool, related *Volume) error {
	return o.SetVolume(boil.GetDB(), insert, related)
}

// SetVolumeP of the volume_state to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeStates.
// Panics on error.
func (o *VolumeState) SetVolumeP(exec boil.Executor, insert bool, related *Volume) {
	if err := o.SetVolume(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVolumeGP of the volume_state to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeStates.
// Uses the global database handle and panics on error.
func (o *VolumeState) SetVolumeGP(insert bool, related *Volume) {
	if err := o.SetVolume(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVolume of the volume_state to the related item.
// Sets o.R.Volume to related.
// Adds o to related.R.VolumeStates.
func (o *VolumeState) SetVolume(exec boil.Executor, insert bool, related *Volume) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"volume_state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"volume"}),
		strmangle.WhereClause("\"", "\"", 2, volumeStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Volume, o.Key}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Volume = related.ID

	if o.R == nil {
		o.R = &volumeStateR{
			Volume: related,
		}
	} else {
		o.R.Volume = related
	}

	if related.R == nil {
		related.R = &volumeR{
			VolumeStates: VolumeStateSlice{o},
		}
	} else {
		related.R.VolumeStates = append(related.R.VolumeStates, o)
	}

	return nil
}

// VolumeStatesG retrieves all records.
func VolumeStatesG(mods ...qm.QueryMod) volumeStateQuery {
	return VolumeStates(boil.GetDB(), mods...)
}

// VolumeStates retrieves all the records using an executor.
func VolumeStates(exec boil.Executor, mods ...qm.QueryMod) volumeStateQuery {
	mods = append(mods, qm.From("\"volume_state\""))
	return volumeStateQuery{NewQuery(exec, mods...)}
}

// FindVolumeStateG retrieves a single record by ID.
func FindVolumeStateG(volume int, key string, selectCols ...string) (*VolumeState, error) {
	return FindVolumeState(boil.GetDB(), volume, key, selectCols...)
}

// FindVolumeStateGP retrieves a single record by ID, and panics on error.
func FindVolumeStateGP(volume int, key string, selectCols ...string) *VolumeState {
	retobj, err := FindVolumeState(boil.GetDB(), volume, key, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindVolumeState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVolumeState(exec boil.Executor, volume int, key string, selectCols ...string) (*VolumeState, error) {
	volumeStateObj := &VolumeState{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"volume_state\" where \"volume\"=$1 AND \"key\"=$2", sel,
	)

	q := queries.Raw(exec, query, volume, key)

	err := q.Bind(volumeStateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from volume_state")
	}

	return volumeStateObj, nil
}

// FindVolumeStateP retrieves a single record by ID with an executor, and panics on error.
func FindVolumeStateP(exec boil.Executor, volume int, key string, selectCols ...string) *VolumeState {
	retobj, err := FindVolumeState(exec, volume, key, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VolumeState) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *VolumeState) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *VolumeState) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *VolumeState) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no volume_state provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeStateColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	volumeStateInsertCacheMut.RLock()
	cache, cached := volumeStateInsertCache[key]
	volumeStateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			volumeStateColumns,
			volumeStateColumnsWithDefault,
			volumeStateColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(volumeStateType, volumeStateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(volumeStateType, volumeStateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"volume_state\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"volume_state\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into volume_state")
	}

	if !cached {
		volumeStateInsertCacheMut.Lock()
		volumeStateInsertCache[key] = cache
		volumeStateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single VolumeState record. See Update for
// whitelist behavior description.
func (o *VolumeState) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single VolumeState record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *VolumeState) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the VolumeState, and panics on error.
// See Update for whitelist behavior description.
func (o *VolumeState) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the VolumeState.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *VolumeState) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	volumeStateUpdateCacheMut.RLock()
	cache, cached := volumeStateUpdateCache[key]
	volumeStateUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(volumeStateColumns, volumeStatePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update volume_state, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"volume_state\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, volumeStatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(volumeStateType, volumeStateMapping, append(wl, volumeStatePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update volume_state row")
	}

	if !cached {
		volumeStateUpdateCacheMut.Lock()
		volumeStateUpdateCache[key] = cache
		volumeStateUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q volumeStateQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q volumeStateQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for volume_state")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VolumeStateSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o VolumeStateSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o VolumeStateSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VolumeStateSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), volumeStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"UPDATE \"volume_state\" SET %s WHERE (\"volume\",\"key\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(volumeStatePrimaryKeyColumns), len(colNames)+1, len(volumeStatePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in volumeState slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *VolumeState) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *VolumeState) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *VolumeState) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *VolumeState) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no volume_state provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeStateColumnsWithDefault, o)

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

	volumeStateUpsertCacheMut.RLock()
	cache, cached := volumeStateUpsertCache[key]
	volumeStateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			volumeStateColumns,
			volumeStateColumnsWithDefault,
			volumeStateColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			volumeStateColumns,
			volumeStatePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert volume_state, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(volumeStatePrimaryKeyColumns))
			copy(conflict, volumeStatePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"volume_state\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(volumeStateType, volumeStateMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(volumeStateType, volumeStateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert volume_state")
	}

	if !cached {
		volumeStateUpsertCacheMut.Lock()
		volumeStateUpsertCache[key] = cache
		volumeStateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single VolumeState record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *VolumeState) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single VolumeState record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *VolumeState) DeleteG() error {
	if o == nil {
		return errors.New("models: no VolumeState provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single VolumeState record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *VolumeState) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single VolumeState record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VolumeState) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no VolumeState provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), volumeStatePrimaryKeyMapping)
	query := "DELETE FROM \"volume_state\" WHERE \"volume\"=$1 AND \"key\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from volume_state")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q volumeStateQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q volumeStateQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no volumeStateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from volume_state")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o VolumeStateSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o VolumeStateSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no VolumeState slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o VolumeStateSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VolumeStateSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no VolumeState slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(volumeStateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), volumeStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"DELETE FROM \"volume_state\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, volumeStatePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(volumeStatePrimaryKeyColumns), 1, len(volumeStatePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from volumeState slice")
	}

	if len(volumeStateAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *VolumeState) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *VolumeState) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *VolumeState) ReloadG() error {
	if o == nil {
		return errors.New("models: no VolumeState provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VolumeState) Reload(exec boil.Executor) error {
	ret, err := FindVolumeState(exec, o.Volume, o.Key)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VolumeStateSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VolumeStateSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VolumeStateSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty VolumeStateSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VolumeStateSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	volumeStates := VolumeStateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), volumeStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	query := fmt.Sprintf(
		"SELECT \"volume_state\".* FROM \"volume_state\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, volumeStatePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(volumeStatePrimaryKeyColumns), 1, len(volumeStatePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, query, args...)

	err := q.Bind(&volumeStates)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VolumeStateSlice")
	}

	*o = volumeStates

	return nil
}

// VolumeStateExists checks if the VolumeState row exists.
func VolumeStateExists(exec boil.Executor, volume int, key string) (bool, error) {
	var exists bool

	query := "select exists(select 1 from \"volume_state\" where \"volume\"=$1 AND \"key\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, volume, key)
	}

	row := exec.QueryRow(query, volume, key)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if volume_state exists")
	}

	return exists, nil
}

// VolumeStateExistsG checks if the VolumeState row exists.
func VolumeStateExistsG(volume int, key string) (bool, error) {
	return VolumeStateExists(boil.GetDB(), volume, key)
}

// VolumeStateExistsGP checks if the VolumeState row exists. Panics on error.
func VolumeStateExistsGP(volume int, key string) bool {
	e, err := VolumeStateExists(boil.GetDB(), volume, key)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// VolumeStateExistsP checks if the VolumeState row exists. Panics on error.
func VolumeStateExistsP(exec boil.Executor, volume int, key string) bool {
	e, err := VolumeStateExists(exec, volume, key)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
