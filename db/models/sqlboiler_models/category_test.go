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

func testCategories(t *testing.T) {
	t.Parallel()

	query := Categories(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCategoriesLive(t *testing.T) {
	all, err := Categories(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all Categories err: ", err)
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
	dumpCmd := exec.Command("pg_dump", "--data-only", dbMain.DbName, "-t", "category")
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
	dumpCmd = exec.Command("pg_dump", "--data-only", dbMain.LiveTestDBName, "-t", "category")
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
		t.Fatalf("CategoriesLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testCategoriesDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = category.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Categories(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CategorySlice{category}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCategoriesExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CategoryExists(tx, category.ID)
	if err != nil {
		t.Errorf("Unable to check if Category exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CategoryExistsG to return true, but got false.")
	}
}
func testCategoriesFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	categoryFound, err := FindCategory(tx, category.ID)
	if err != nil {
		t.Error(err)
	}

	if categoryFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCategoriesBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Categories(tx).Bind(category); err != nil {
		t.Error(err)
	}
}

func testCategoriesOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Categories(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCategoriesAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	categoryOne := &Category{}
	categoryTwo := &Category{}
	if err = randomize.Struct(seed, categoryOne, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTwo, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = categoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = categoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Categories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	categoryOne := &Category{}
	categoryTwo := &Category{}
	if err = randomize.Struct(seed, categoryOne, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTwo, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = categoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = categoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func categoryBeforeInsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterInsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterSelectHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeUpdateHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterUpdateHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeDeleteHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterDeleteHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeUpsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterUpsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func testCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	empty := &Category{}

	AddCategoryHook(boil.BeforeInsertHook, categoryBeforeInsertHook)
	if err = category.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", category)
	}
	categoryBeforeInsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterInsertHook, categoryAfterInsertHook)
	if err = category.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", category)
	}
	categoryAfterInsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterSelectHook, categoryAfterSelectHook)
	if err = category.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", category)
	}
	categoryAfterSelectHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeUpdateHook, categoryBeforeUpdateHook)
	if err = category.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", category)
	}
	categoryBeforeUpdateHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterUpdateHook, categoryAfterUpdateHook)
	if err = category.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", category)
	}
	categoryAfterUpdateHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeDeleteHook, categoryBeforeDeleteHook)
	if err = category.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", category)
	}
	categoryBeforeDeleteHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterDeleteHook, categoryAfterDeleteHook)
	if err = category.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", category)
	}
	categoryAfterDeleteHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeUpsertHook, categoryBeforeUpsertHook)
	if err = category.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", category)
	}
	categoryBeforeUpsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterUpsertHook, categoryAfterUpsertHook)
	if err = category.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(category, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", category)
	}
	categoryAfterUpsertHooks = []CategoryHook{}
}
func testCategoriesInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx, categoryColumns...); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoryToManyRecords(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Category
	var b, c Record

	foreignBlacklist := recordColumnsWithDefault
	if err := randomize.Struct(seed, &b, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := categoryColumnsWithDefault
	if err := randomize.Struct(seed, &a, categoryDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	b.Category = a.ID
	c.Category = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	record, err := a.RecordsByFk(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range record {
		if v.Category == b.Category {
			bFound = true
		}
		if v.Category == c.Category {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CategorySlice{&a}
	if err = a.L.LoadRecords(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Records); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Records = nil
	if err = a.L.LoadRecords(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Records); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", record)
	}
}

func testCategoryToManyMetrics(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Category
	var b, c Metric

	foreignBlacklist := metricColumnsWithDefault
	foreignBlacklist = append(foreignBlacklist, metricColumnsWithCustom...)

	if err := randomize.Struct(seed, &b, metricDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Metric struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, metricDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Metric struct: %s", err)
	}
	b.Release = custom_types.NullReleaseRandom()
	c.Release = custom_types.NullReleaseRandom()
	b.Type = custom_types.DataTypeRandom()
	c.Type = custom_types.DataTypeRandom()

	localBlacklist := categoryColumnsWithDefault
	if err := randomize.Struct(seed, &a, categoryDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	b.Category = a.ID
	c.Category = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	metric, err := a.MetricsByFk(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range metric {
		if v.Category == b.Category {
			bFound = true
		}
		if v.Category == c.Category {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CategorySlice{&a}
	if err = a.L.LoadMetrics(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Metrics); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Metrics = nil
	if err = a.L.LoadMetrics(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Metrics); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", metric)
	}
}

func testCategoryToManyAddOpRecords(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Category
	var b, c, d, e Record

	seed := randomize.NewSeed()
	localComplelementList := strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)
	if err = randomize.Struct(seed, &a, categoryDBTypes, false, localComplelementList...); err != nil {
		t.Fatal(err)
	}

	foreignComplementList := strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)

	foreigners := []*Record{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recordDBTypes, false, foreignComplementList...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Record{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRecords(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Category {
			t.Error("foreign key was wrong value", a.ID, first.Category)
		}
		if a.ID != second.Category {
			t.Error("foreign key was wrong value", a.ID, second.Category)
		}

		if first.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Records[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Records[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RecordsByFk(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCategoryToManyAddOpMetrics(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Category
	var b, c, d, e Metric

	seed := randomize.NewSeed()
	localComplelementList := strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)
	if err = randomize.Struct(seed, &a, categoryDBTypes, false, localComplelementList...); err != nil {
		t.Fatal(err)
	}

	foreignComplementList := strmangle.SetComplement(metricPrimaryKeyColumns, metricColumnsWithoutDefault)
	foreignComplementList = append(foreignComplementList, metricColumnsWithCustom...)

	foreigners := []*Metric{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, metricDBTypes, false, foreignComplementList...); err != nil {
			t.Fatal(err)
		}
		x.Release = custom_types.NullReleaseRandom()
		x.Type = custom_types.DataTypeRandom()

	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Metric{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMetrics(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Category {
			t.Error("foreign key was wrong value", a.ID, first.Category)
		}
		if a.ID != second.Category {
			t.Error("foreign key was wrong value", a.ID, second.Category)
		}

		if first.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Metrics[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Metrics[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MetricsByFk(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCategoriesReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = category.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CategorySlice{category}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCategoriesSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Categories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	categoryDBTypes = map[string]string{`Description`: `text`, `ID`: `smallint`, `Name`: `character varying`}
	_               = bytes.MinRead
)

func testCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if len(categoryColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := categoryColumnsWithDefault

	if err = randomize.Struct(seed, category, categoryDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err = category.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(categoryColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := categoryPrimaryKeyColumns

	if err = randomize.Struct(seed, category, categoryDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(categoryColumns, categoryPrimaryKeyColumns) {
		fields = categoryColumns
	} else {
		fields = strmangle.SetComplement(
			categoryColumns,
			categoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(category))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CategorySlice{category}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(categoryColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Category: %s", err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := categoryPrimaryKeyColumns

	if err = randomize.Struct(seed, category, categoryDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err = category.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Category: %s", err)
	}

	count, err = Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
