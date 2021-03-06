// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBloggers(t *testing.T) {
	t.Parallel()

	query := Bloggers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBloggersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBloggersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Bloggers().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBloggersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BloggerSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBloggersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BloggerExists(tx, o.Index)
	if err != nil {
		t.Errorf("Unable to check if Blogger exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BloggerExists to return true, but got false.")
	}
}

func testBloggersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bloggerFound, err := FindBlogger(tx, o.Index)
	if err != nil {
		t.Error(err)
	}

	if bloggerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBloggersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Bloggers().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testBloggersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Bloggers().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBloggersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bloggerOne := &Blogger{}
	bloggerTwo := &Blogger{}
	if err = randomize.Struct(seed, bloggerOne, bloggerDBTypes, false, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}
	if err = randomize.Struct(seed, bloggerTwo, bloggerDBTypes, false, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = bloggerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bloggerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bloggers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBloggersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bloggerOne := &Blogger{}
	bloggerTwo := &Blogger{}
	if err = randomize.Struct(seed, bloggerOne, bloggerDBTypes, false, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}
	if err = randomize.Struct(seed, bloggerTwo, bloggerDBTypes, false, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = bloggerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bloggerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bloggerBeforeInsertHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerAfterInsertHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerAfterSelectHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerBeforeUpdateHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerAfterUpdateHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerBeforeDeleteHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerAfterDeleteHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerBeforeUpsertHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func bloggerAfterUpsertHook(e boil.Executor, o *Blogger) error {
	*o = Blogger{}
	return nil
}

func testBloggersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Blogger{}
	o := &Blogger{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bloggerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Blogger object: %s", err)
	}

	AddBloggerHook(boil.BeforeInsertHook, bloggerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bloggerBeforeInsertHooks = []BloggerHook{}

	AddBloggerHook(boil.AfterInsertHook, bloggerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bloggerAfterInsertHooks = []BloggerHook{}

	AddBloggerHook(boil.AfterSelectHook, bloggerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bloggerAfterSelectHooks = []BloggerHook{}

	AddBloggerHook(boil.BeforeUpdateHook, bloggerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bloggerBeforeUpdateHooks = []BloggerHook{}

	AddBloggerHook(boil.AfterUpdateHook, bloggerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bloggerAfterUpdateHooks = []BloggerHook{}

	AddBloggerHook(boil.BeforeDeleteHook, bloggerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bloggerBeforeDeleteHooks = []BloggerHook{}

	AddBloggerHook(boil.AfterDeleteHook, bloggerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bloggerAfterDeleteHooks = []BloggerHook{}

	AddBloggerHook(boil.BeforeUpsertHook, bloggerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bloggerBeforeUpsertHooks = []BloggerHook{}

	AddBloggerHook(boil.AfterUpsertHook, bloggerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bloggerAfterUpsertHooks = []BloggerHook{}
}

func testBloggersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBloggersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(bloggerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBloggersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBloggersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BloggerSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testBloggersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bloggers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bloggerDBTypes = map[string]string{`Index`: `int`, `ID`: `varchar`, `Nick`: `varchar`, `DOB`: `datetime`, `Blogname`: `varchar`, `Blogskin`: `int`, `Email`: `varchar`, `Visitor`: `int`, `Intro`: `varchar`, `Blogs`: `int`, `IP`: `char`, `TS`: `tinyint`, `LastLog`: `timestamp`, `LastPost`: `datetime`, `Activate`: `tinyint`, `Reveal`: `tinyint`, `Userpic`: `int`, `Lang`: `varchar`, `Widget`: `varchar`}
	_              = bytes.MinRead
)

func testBloggersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bloggerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bloggerAllColumns) == len(bloggerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBloggersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bloggerAllColumns) == len(bloggerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Blogger{}
	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bloggerDBTypes, true, bloggerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bloggerAllColumns, bloggerPrimaryKeyColumns) {
		fields = bloggerAllColumns
	} else {
		fields = strmangle.SetComplement(
			bloggerAllColumns,
			bloggerPrimaryKeyColumns,
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

	slice := BloggerSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBloggersUpsert(t *testing.T) {
	t.Parallel()

	if len(bloggerAllColumns) == len(bloggerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBloggerUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Blogger{}
	if err = randomize.Struct(seed, &o, bloggerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Blogger: %s", err)
	}

	count, err := Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bloggerDBTypes, false, bloggerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blogger struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Blogger: %s", err)
	}

	count, err = Bloggers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
