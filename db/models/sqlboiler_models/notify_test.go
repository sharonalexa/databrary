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

	"github.com/databrary/databrary/db/models/custom_types"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testNotifies(t *testing.T) {
	t.Parallel()

	query := Notifies(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNotifiesLive(t *testing.T) {
	all, err := Notifies(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all Notifies err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM notify) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM notify) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("NotifiesLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testNotifiesDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = notify.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotifiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Notifies(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotifiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := NotifySlice{notify}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotifiesExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := NotifyExists(tx, notify.Target, notify.Notice)
	if err != nil {
		t.Errorf("Unable to check if Notify exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NotifyExistsG to return true, but got false.")
	}
}

func testNotifiesFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	notifyFound, err := FindNotify(tx, notify.Target, notify.Notice)
	if err != nil {
		t.Error(err)
	}

	if notifyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNotifiesBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Notifies(tx).Bind(notify); err != nil {
		t.Error(err)
	}
}

func testNotifiesOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Notifies(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNotifiesAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notifyOne := &Notify{}
	notifyTwo := &Notify{}
	if err = randomize.Struct(seed, notifyOne, notifyDBTypes, false, notifyColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize Notify struct: %s", err)
	}
	if err = randomize.Struct(seed, notifyTwo, notifyDBTypes, false, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notifyOne.Delivery = custom_types.NoticeDeliveryRandom()
	notifyTwo.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notifyOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = notifyTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Notifies(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNotifiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notifyOne := &Notify{}
	notifyTwo := &Notify{}
	if err = randomize.Struct(seed, notifyOne, notifyDBTypes, false, notifyColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize Notify struct: %s", err)
	}
	if err = randomize.Struct(seed, notifyTwo, notifyDBTypes, false, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notifyOne.Delivery = custom_types.NoticeDeliveryRandom()
	notifyTwo.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notifyOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = notifyTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func notifyBeforeInsertHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyAfterInsertHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyAfterSelectHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyBeforeUpdateHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyAfterUpdateHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyBeforeDeleteHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyAfterDeleteHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyBeforeUpsertHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func notifyAfterUpsertHook(e boil.Executor, o *Notify) error {
	*o = Notify{}
	return nil
}

func testNotifiesHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	empty := &Notify{}

	AddNotifyHook(boil.BeforeInsertHook, notifyBeforeInsertHook)
	if err = notify.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", notify)
	}
	notifyBeforeInsertHooks = []NotifyHook{}

	AddNotifyHook(boil.AfterInsertHook, notifyAfterInsertHook)
	if err = notify.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", notify)
	}
	notifyAfterInsertHooks = []NotifyHook{}

	AddNotifyHook(boil.AfterSelectHook, notifyAfterSelectHook)
	if err = notify.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", notify)
	}
	notifyAfterSelectHooks = []NotifyHook{}

	AddNotifyHook(boil.BeforeUpdateHook, notifyBeforeUpdateHook)
	if err = notify.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", notify)
	}
	notifyBeforeUpdateHooks = []NotifyHook{}

	AddNotifyHook(boil.AfterUpdateHook, notifyAfterUpdateHook)
	if err = notify.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", notify)
	}
	notifyAfterUpdateHooks = []NotifyHook{}

	AddNotifyHook(boil.BeforeDeleteHook, notifyBeforeDeleteHook)
	if err = notify.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", notify)
	}
	notifyBeforeDeleteHooks = []NotifyHook{}

	AddNotifyHook(boil.AfterDeleteHook, notifyAfterDeleteHook)
	if err = notify.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", notify)
	}
	notifyAfterDeleteHooks = []NotifyHook{}

	AddNotifyHook(boil.BeforeUpsertHook, notifyBeforeUpsertHook)
	if err = notify.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", notify)
	}
	notifyBeforeUpsertHooks = []NotifyHook{}

	AddNotifyHook(boil.AfterUpsertHook, notifyAfterUpsertHook)
	if err = notify.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(notify, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", notify)
	}
	notifyAfterUpsertHooks = []NotifyHook{}
}
func testNotifiesInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotifiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx, notifyColumns...); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotifyToOneNoticeUsingNotice(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Notice
	var local Notify

	foreignBlacklist := noticeColumnsWithDefault
	foreignBlacklist = append(foreignBlacklist, noticeColumnsWithCustom...)

	if err := randomize.Struct(seed, &foreign, noticeDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notice struct: %s", err)
	}
	foreign.Delivery = custom_types.NoticeDeliveryRandom()

	localBlacklist := notifyColumnsWithDefault
	localBlacklist = append(localBlacklist, notifyColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, notifyDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}
	local.Delivery = custom_types.NoticeDeliveryRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Notice = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.NoticeByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NotifySlice{&local}
	if err = local.L.LoadNotice(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Notice == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Notice = nil
	if err = local.L.LoadNotice(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Notice == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNotifyToOneAccountUsingTarget(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Account
	var local Notify

	foreignBlacklist := accountColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, accountDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	localBlacklist := notifyColumnsWithDefault
	localBlacklist = append(localBlacklist, notifyColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, notifyDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}
	local.Delivery = custom_types.NoticeDeliveryRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Target = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TargetByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NotifySlice{&local}
	if err = local.L.LoadTarget(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Target == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Target = nil
	if err = local.L.LoadTarget(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Target == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNotifyToOneSetOpNoticeUsingNotice(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Notify
	var b, c Notice

	foreignBlacklist := strmangle.SetComplement(noticePrimaryKeyColumns, noticeColumnsWithoutDefault)
	foreignBlacklist = append(foreignBlacklist, noticeColumnsWithCustom...)

	if err := randomize.Struct(seed, &b, noticeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notice struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, noticeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notice struct: %s", err)
	}
	b.Delivery = custom_types.NoticeDeliveryRandom()
	c.Delivery = custom_types.NoticeDeliveryRandom()

	localBlacklist := strmangle.SetComplement(notifyPrimaryKeyColumns, notifyColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, notifyColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, notifyDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}
	a.Delivery = custom_types.NoticeDeliveryRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Notice{&b, &c} {
		err = a.SetNotice(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Notice != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Notifies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Notice != x.ID {
			t.Error("foreign key was wrong value", a.Notice)
		}

		if exists, err := NotifyExists(tx, a.Target, a.Notice); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testNotifyToOneSetOpAccountUsingTarget(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Notify
	var b, c Account

	foreignBlacklist := strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, accountDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, accountDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(notifyPrimaryKeyColumns, notifyColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, notifyColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, notifyDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}
	a.Delivery = custom_types.NoticeDeliveryRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Account{&b, &c} {
		err = a.SetTarget(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Target != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TargetNotifies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Target != x.ID {
			t.Error("foreign key was wrong value", a.Target)
		}

		if exists, err := NotifyExists(tx, a.Target, a.Notice); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testNotifiesReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = notify.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testNotifiesReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := NotifySlice{notify}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testNotifiesSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Notifies(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	notifyDBTypes = map[string]string{`Delivery`: `enum.notice_delivery('none','site','weekly','daily','async')`, `Notice`: `smallint`, `Target`: `integer`}
	_             = bytes.MinRead
)

func testNotifiesUpdate(t *testing.T) {
	t.Parallel()

	if len(notifyColumns) == len(notifyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := notifyColumnsWithDefault
	blacklist = append(blacklist, notifyColumnsWithCustom...)

	if err = randomize.Struct(seed, notify, notifyDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	if err = notify.Update(tx); err != nil {
		t.Error(err)
	}
}

func testNotifiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(notifyColumns) == len(notifyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := notifyPrimaryKeyColumns
	blacklist = append(blacklist, notifyColumnsWithCustom...)

	if err = randomize.Struct(seed, notify, notifyDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(notifyColumns, notifyPrimaryKeyColumns) {
		fields = notifyColumns
	} else {
		fields = strmangle.SetComplement(
			notifyColumns,
			notifyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(notify))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := NotifySlice{notify}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testNotifiesUpsert(t *testing.T) {
	t.Parallel()

	if len(notifyColumns) == len(notifyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	notify := &Notify{}
	if err = randomize.Struct(seed, notify, notifyDBTypes, true, notifyColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = notify.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Notify: %s", err)
	}

	count, err := Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := notifyPrimaryKeyColumns

	blacklist = append(blacklist, notifyColumnsWithCustom...)

	if err = randomize.Struct(seed, notify, notifyDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize Notify struct: %s", err)
	}

	notify.Delivery = custom_types.NoticeDeliveryRandom()

	if err = notify.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Notify: %s", err)
	}

	count, err = Notifies(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
