// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package models

import (
	"bytes"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testVolumeLinks(t *testing.T) {
	t.Parallel()

	query := VolumeLinks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testVolumeLinksLive(t *testing.T) {
	all, err := VolumeLinks(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all VolumeLinks err: ", err)
	}
	tx, err := dbMain.liveTestDbConn.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: ", err)
	}
	for _, v := range all {
		err := v.Insert(tx)
		if err != nil {
			t.Fatalf("failed to failed to insert %s because of %s", v, err)
		}

	}
	err = tx.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: ", err)
	}
	bf := &bytes.Buffer{}
	dumpCmd := exec.Command("pg_dump", "--data-only", dbMain.DbName, "-t", "volume_link")
	dumpCmd.Env = append(os.Environ(), dbMain.pgEnv()...)
	dumpCmd.Stdout = bf
	err = dumpCmd.Start()
	if err != nil {
		t.Fatalf("failed to start dump from live db because of %s", err)
	}
	dumpCmd.Wait()
	if err != nil {
		t.Fatalf("failed to wait dump from live db because of %s", err)
	}
	bg := &bytes.Buffer{}
	dumpCmd = exec.Command("pg_dump", "--data-only", dbMain.LiveTestDBName, "-t", "volume_link")
	dumpCmd.Env = append(os.Environ(), dbMain.pgEnv()...)
	dumpCmd.Stdout = bg
	err = dumpCmd.Start()
	if err != nil {
		t.Fatalf("failed to start dump from test db because of %s", err)
	}
	dumpCmd.Wait()
	if err != nil {
		t.Fatalf("failed to wait dump from test db because of %s", err)
	}
	bfslice := sort.StringSlice(difflib.SplitLines(bf.String()))
	gfslice := sort.StringSlice(difflib.SplitLines(bg.String()))
	bfslice.Sort()
	gfslice.Sort()
	diff := difflib.ContextDiff{
		A:        bfslice,
		B:        gfslice,
		FromFile: "databrary",
		ToFile:   "test",
		Context:  1,
	}
	result, _ := difflib.GetContextDiffString(diff)
	if len(result) > 0 {
		t.Fatalf("VolumeLinksLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testVolumeLinksDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = volumeLink.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeLinksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = VolumeLinks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeLinksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := VolumeLinkSlice{volumeLink}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testVolumeLinksExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := VolumeLinkExists(tx, volumeLink.Volume, volumeLink.URL)
	if err != nil {
		t.Errorf("Unable to check if VolumeLink exists: %s", err)
	}
	if !e {
		t.Errorf("Expected VolumeLinkExistsG to return true, but got false.")
	}
}
func testVolumeLinksFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	volumeLinkFound, err := FindVolumeLink(tx, volumeLink.Volume, volumeLink.URL)
	if err != nil {
		t.Error(err)
	}

	if volumeLinkFound == nil {
		t.Error("want a record, got nil")
	}
}
func testVolumeLinksBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = VolumeLinks(tx).Bind(volumeLink); err != nil {
		t.Error(err)
	}
}

func testVolumeLinksOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := VolumeLinks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testVolumeLinksAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLinkOne := &VolumeLink{}
	volumeLinkTwo := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLinkOne, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}
	if err = randomize.Struct(seed, volumeLinkTwo, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLinkOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = volumeLinkTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := VolumeLinks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testVolumeLinksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLinkOne := &VolumeLink{}
	volumeLinkTwo := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLinkOne, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}
	if err = randomize.Struct(seed, volumeLinkTwo, volumeLinkDBTypes, false, volumeLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLinkOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = volumeLinkTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func volumeLinkBeforeInsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterInsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterSelectHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkBeforeUpdateHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterUpdateHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkBeforeDeleteHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterDeleteHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkBeforeUpsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func volumeLinkAfterUpsertHook(e boil.Executor, o *VolumeLink) error {
	*o = VolumeLink{}
	return nil
}

func testVolumeLinksHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	empty := &VolumeLink{}

	AddVolumeLinkHook(boil.BeforeInsertHook, volumeLinkBeforeInsertHook)
	if err = volumeLink.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkBeforeInsertHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterInsertHook, volumeLinkAfterInsertHook)
	if err = volumeLink.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkAfterInsertHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterSelectHook, volumeLinkAfterSelectHook)
	if err = volumeLink.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkAfterSelectHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.BeforeUpdateHook, volumeLinkBeforeUpdateHook)
	if err = volumeLink.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkBeforeUpdateHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterUpdateHook, volumeLinkAfterUpdateHook)
	if err = volumeLink.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkAfterUpdateHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.BeforeDeleteHook, volumeLinkBeforeDeleteHook)
	if err = volumeLink.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkBeforeDeleteHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterDeleteHook, volumeLinkAfterDeleteHook)
	if err = volumeLink.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkAfterDeleteHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.BeforeUpsertHook, volumeLinkBeforeUpsertHook)
	if err = volumeLink.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkBeforeUpsertHooks = []VolumeLinkHook{}

	AddVolumeLinkHook(boil.AfterUpsertHook, volumeLinkAfterUpsertHook)
	if err = volumeLink.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeLink, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", volumeLink)
	}
	volumeLinkAfterUpsertHooks = []VolumeLinkHook{}
}
func testVolumeLinksInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testVolumeLinksInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx, volumeLinkColumns...); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testVolumeLinkToOneVolumeUsingVolume(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Volume
	var local VolumeLink

	foreignBlacklist := volumeColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, volumeDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	localBlacklist := volumeLinkColumnsWithDefault
	if err := randomize.Struct(seed, &local, volumeLinkDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Volume = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.VolumeByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := VolumeLinkSlice{&local}
	if err = local.L.LoadVolume(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Volume = nil
	if err = local.L.LoadVolume(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testVolumeLinkToOneSetOpVolumeUsingVolume(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a VolumeLink
	var b, c Volume

	foreignBlacklist := strmangle.SetComplement(volumePrimaryKeyColumns, volumeColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, volumeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, volumeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(volumeLinkPrimaryKeyColumns, volumeLinkColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, volumeLinkDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Volume{&b, &c} {
		err = a.SetVolume(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Volume != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.VolumeLinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Volume != x.ID {
			t.Error("foreign key was wrong value", a.Volume)
		}

		if exists, err := VolumeLinkExists(tx, a.Volume, a.URL); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testVolumeLinksReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = volumeLink.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeLinksReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := VolumeLinkSlice{volumeLink}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testVolumeLinksSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := VolumeLinks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	volumeLinkDBTypes = map[string]string{`Head`: `text`, `URL`: `text`, `Volume`: `integer`}
	_                 = bytes.MinRead
)

func testVolumeLinksUpdate(t *testing.T) {
	t.Parallel()

	if len(volumeLinkColumns) == len(volumeLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := volumeLinkColumnsWithDefault

	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	if err = volumeLink.Update(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeLinksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(volumeLinkColumns) == len(volumeLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := volumeLinkPrimaryKeyColumns

	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(volumeLinkColumns, volumeLinkPrimaryKeyColumns) {
		fields = volumeLinkColumns
	} else {
		fields = strmangle.SetComplement(
			volumeLinkColumns,
			volumeLinkPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(volumeLink))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := VolumeLinkSlice{volumeLink}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testVolumeLinksUpsert(t *testing.T) {
	t.Parallel()

	if len(volumeLinkColumns) == len(volumeLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	volumeLink := &VolumeLink{}
	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeLink.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert VolumeLink: %s", err)
	}

	count, err := VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := volumeLinkPrimaryKeyColumns

	if err = randomize.Struct(seed, volumeLink, volumeLinkDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeLink struct: %s", err)
	}

	if err = volumeLink.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert VolumeLink: %s", err)
	}

	count, err = VolumeLinks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
