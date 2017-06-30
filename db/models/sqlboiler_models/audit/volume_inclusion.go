// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package audit

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/databrary/databrary/db/models/custom_types"
	"github.com/databrary/sqlboiler/boil"
	"github.com/databrary/sqlboiler/queries"
	"github.com/databrary/sqlboiler/queries/qm"
	"github.com/databrary/sqlboiler/strmangle"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"sync"
	"time"
)

// VolumeInclusion is an object representing the database table.
type VolumeInclusion struct {
	AuditTime   time.Time            `db:"audit_time" json:"volumeInclusion_audit_time"`
	AuditUser   int                  `db:"audit_user" json:"volumeInclusion_audit_user"`
	AuditIP     custom_types.Inet    `db:"audit_ip" json:"volumeInclusion_audit_ip"`
	AuditAction custom_types.Action  `db:"audit_action" json:"volumeInclusion_audit_action"`
	Container   int                  `db:"container" json:"volumeInclusion_container"`
	Segment     custom_types.Segment `db:"segment" json:"volumeInclusion_segment"`
	Volume      int                  `db:"volume" json:"volumeInclusion_volume"`

	R *volumeInclusionR `db:"-" json:"-"`
	L volumeInclusionL  `db:"-" json:"-"`
}

// volumeInclusionR is where relationships are stored.
type volumeInclusionR struct {
}

// volumeInclusionL is where Load methods for each relationship are stored.
type volumeInclusionL struct{}

var (
	volumeInclusionColumns               = []string{"audit_time", "audit_user", "audit_ip", "audit_action", "container", "segment", "volume"}
	volumeInclusionColumnsWithoutDefault = []string{"audit_user", "audit_ip", "audit_action", "container", "segment", "volume"}
	volumeInclusionColumnsWithDefault    = []string{"audit_time"}
	volumeInclusionColumnsWithCustom     = []string{"audit_ip", "audit_action", "segment"}
)

type (
	// VolumeInclusionSlice is an alias for a slice of pointers to VolumeInclusion.
	// This should generally be used opposed to []VolumeInclusion.
	VolumeInclusionSlice []*VolumeInclusion
	// VolumeInclusionHook is the signature for custom VolumeInclusion hook methods
	VolumeInclusionHook func(boil.Executor, *VolumeInclusion) error

	volumeInclusionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	volumeInclusionType    = reflect.TypeOf(&VolumeInclusion{})
	volumeInclusionMapping = queries.MakeStructMapping(volumeInclusionType)

	volumeInclusionInsertCacheMut sync.RWMutex
	volumeInclusionInsertCache    = make(map[string]insertCache)
	volumeInclusionUpdateCacheMut sync.RWMutex
	volumeInclusionUpdateCache    = make(map[string]updateCache)
	volumeInclusionUpsertCacheMut sync.RWMutex
	volumeInclusionUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var volumeInclusionBeforeInsertHooks []VolumeInclusionHook
var volumeInclusionBeforeUpdateHooks []VolumeInclusionHook
var volumeInclusionBeforeDeleteHooks []VolumeInclusionHook
var volumeInclusionBeforeUpsertHooks []VolumeInclusionHook

var volumeInclusionAfterInsertHooks []VolumeInclusionHook
var volumeInclusionAfterSelectHooks []VolumeInclusionHook
var volumeInclusionAfterUpdateHooks []VolumeInclusionHook
var volumeInclusionAfterDeleteHooks []VolumeInclusionHook
var volumeInclusionAfterUpsertHooks []VolumeInclusionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VolumeInclusion) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VolumeInclusion) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VolumeInclusion) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VolumeInclusion) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VolumeInclusion) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VolumeInclusion) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VolumeInclusion) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VolumeInclusion) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VolumeInclusion) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeInclusionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVolumeInclusionHook registers your hook function for all future operations.
func AddVolumeInclusionHook(hookPoint boil.HookPoint, volumeInclusionHook VolumeInclusionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		volumeInclusionBeforeInsertHooks = append(volumeInclusionBeforeInsertHooks, volumeInclusionHook)
	case boil.BeforeUpdateHook:
		volumeInclusionBeforeUpdateHooks = append(volumeInclusionBeforeUpdateHooks, volumeInclusionHook)
	case boil.BeforeDeleteHook:
		volumeInclusionBeforeDeleteHooks = append(volumeInclusionBeforeDeleteHooks, volumeInclusionHook)
	case boil.BeforeUpsertHook:
		volumeInclusionBeforeUpsertHooks = append(volumeInclusionBeforeUpsertHooks, volumeInclusionHook)
	case boil.AfterInsertHook:
		volumeInclusionAfterInsertHooks = append(volumeInclusionAfterInsertHooks, volumeInclusionHook)
	case boil.AfterSelectHook:
		volumeInclusionAfterSelectHooks = append(volumeInclusionAfterSelectHooks, volumeInclusionHook)
	case boil.AfterUpdateHook:
		volumeInclusionAfterUpdateHooks = append(volumeInclusionAfterUpdateHooks, volumeInclusionHook)
	case boil.AfterDeleteHook:
		volumeInclusionAfterDeleteHooks = append(volumeInclusionAfterDeleteHooks, volumeInclusionHook)
	case boil.AfterUpsertHook:
		volumeInclusionAfterUpsertHooks = append(volumeInclusionAfterUpsertHooks, volumeInclusionHook)
	}
}

// OneP returns a single volumeInclusion record from the query, and panics on error.
func (q volumeInclusionQuery) OneP() *VolumeInclusion {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single volumeInclusion record from the query.
func (q volumeInclusionQuery) One() (*VolumeInclusion, error) {
	o := &VolumeInclusion{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for volume_inclusion")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all VolumeInclusion records from the query, and panics on error.
func (q volumeInclusionQuery) AllP() VolumeInclusionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all VolumeInclusion records from the query.
func (q volumeInclusionQuery) All() (VolumeInclusionSlice, error) {
	var o VolumeInclusionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VolumeInclusion slice")
	}

	if len(volumeInclusionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all VolumeInclusion records in the query, and panics on error.
func (q volumeInclusionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all VolumeInclusion records in the query.
func (q volumeInclusionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count volume_inclusion rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q volumeInclusionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q volumeInclusionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if volume_inclusion exists")
	}

	return count > 0, nil
}

// VolumeInclusionsG retrieves all records.
func VolumeInclusionsG(mods ...qm.QueryMod) volumeInclusionQuery {
	return VolumeInclusions(boil.GetDB(), mods...)
}

// VolumeInclusions retrieves all the records using an executor.
func VolumeInclusions(exec boil.Executor, mods ...qm.QueryMod) volumeInclusionQuery {
	mods = append(mods, qm.From("\"audit\".\"volume_inclusion\""))
	return volumeInclusionQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VolumeInclusion) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *VolumeInclusion) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *VolumeInclusion) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *VolumeInclusion) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no volume_inclusion provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeInclusionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	volumeInclusionInsertCacheMut.RLock()
	cache, cached := volumeInclusionInsertCache[key]
	volumeInclusionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			volumeInclusionColumns,
			volumeInclusionColumnsWithDefault,
			volumeInclusionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(volumeInclusionType, volumeInclusionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(volumeInclusionType, volumeInclusionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit\".\"volume_inclusion\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit\".\"volume_inclusion\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into volume_inclusion")
	}

	if !cached {
		volumeInclusionInsertCacheMut.Lock()
		volumeInclusionInsertCache[key] = cache
		volumeInclusionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}