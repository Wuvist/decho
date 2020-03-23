// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Comment is an object representing the database table.
type Comment struct {
	Index    int         `boil:"index" json:"index" toml:"index" yaml:"index"`
	Blogger  null.Int    `boil:"blogger" json:"blogger,omitempty" toml:"blogger" yaml:"blogger,omitempty"`
	Article  int         `boil:"article" json:"article" toml:"article" yaml:"article"`
	Title    string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Content  string      `boil:"content" json:"content" toml:"content" yaml:"content"`
	Author   null.String `boil:"author" json:"author,omitempty" toml:"author" yaml:"author,omitempty"`
	Homepage null.String `boil:"homepage" json:"homepage,omitempty" toml:"homepage" yaml:"homepage,omitempty"`
	AddDate  time.Time   `boil:"add_date" json:"add_date" toml:"add_date" yaml:"add_date"`
	IP       null.String `boil:"ip" json:"ip,omitempty" toml:"ip" yaml:"ip,omitempty"`
	Utype    uint8       `boil:"utype" json:"utype" toml:"utype" yaml:"utype"`
	UID      null.String `boil:"uid" json:"uid,omitempty" toml:"uid" yaml:"uid,omitempty"`

	R *commentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommentColumns = struct {
	Index    string
	Blogger  string
	Article  string
	Title    string
	Content  string
	Author   string
	Homepage string
	AddDate  string
	IP       string
	Utype    string
	UID      string
}{
	Index:    "index",
	Blogger:  "blogger",
	Article:  "article",
	Title:    "title",
	Content:  "content",
	Author:   "author",
	Homepage: "homepage",
	AddDate:  "add_date",
	IP:       "ip",
	Utype:    "utype",
	UID:      "uid",
}

// Generated where

type whereHelperuint8 struct{ field string }

func (w whereHelperuint8) EQ(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint8) NEQ(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint8) LT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint8) LTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint8) GT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint8) GTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var CommentWhere = struct {
	Index    whereHelperint
	Blogger  whereHelpernull_Int
	Article  whereHelperint
	Title    whereHelperstring
	Content  whereHelperstring
	Author   whereHelpernull_String
	Homepage whereHelpernull_String
	AddDate  whereHelpertime_Time
	IP       whereHelpernull_String
	Utype    whereHelperuint8
	UID      whereHelpernull_String
}{
	Index:    whereHelperint{field: "`comment`.`index`"},
	Blogger:  whereHelpernull_Int{field: "`comment`.`blogger`"},
	Article:  whereHelperint{field: "`comment`.`article`"},
	Title:    whereHelperstring{field: "`comment`.`title`"},
	Content:  whereHelperstring{field: "`comment`.`content`"},
	Author:   whereHelpernull_String{field: "`comment`.`author`"},
	Homepage: whereHelpernull_String{field: "`comment`.`homepage`"},
	AddDate:  whereHelpertime_Time{field: "`comment`.`add_date`"},
	IP:       whereHelpernull_String{field: "`comment`.`ip`"},
	Utype:    whereHelperuint8{field: "`comment`.`utype`"},
	UID:      whereHelpernull_String{field: "`comment`.`uid`"},
}

// CommentRels is where relationship names are stored.
var CommentRels = struct {
}{}

// commentR is where relationships are stored.
type commentR struct {
}

// NewStruct creates a new relationship struct
func (*commentR) NewStruct() *commentR {
	return &commentR{}
}

// commentL is where Load methods for each relationship are stored.
type commentL struct{}

var (
	commentAllColumns            = []string{"index", "blogger", "article", "title", "content", "author", "homepage", "add_date", "ip", "utype", "uid"}
	commentColumnsWithoutDefault = []string{"blogger", "article", "title", "content", "author", "homepage", "add_date", "ip", "uid"}
	commentColumnsWithDefault    = []string{"index", "utype"}
	commentPrimaryKeyColumns     = []string{"index"}
)

type (
	// CommentSlice is an alias for a slice of pointers to Comment.
	// This should generally be used opposed to []Comment.
	CommentSlice []*Comment
	// CommentHook is the signature for custom Comment hook methods
	CommentHook func(boil.Executor, *Comment) error

	commentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commentType                 = reflect.TypeOf(&Comment{})
	commentMapping              = queries.MakeStructMapping(commentType)
	commentPrimaryKeyMapping, _ = queries.BindMapping(commentType, commentMapping, commentPrimaryKeyColumns)
	commentInsertCacheMut       sync.RWMutex
	commentInsertCache          = make(map[string]insertCache)
	commentUpdateCacheMut       sync.RWMutex
	commentUpdateCache          = make(map[string]updateCache)
	commentUpsertCacheMut       sync.RWMutex
	commentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var commentBeforeInsertHooks []CommentHook
var commentBeforeUpdateHooks []CommentHook
var commentBeforeDeleteHooks []CommentHook
var commentBeforeUpsertHooks []CommentHook

var commentAfterInsertHooks []CommentHook
var commentAfterSelectHooks []CommentHook
var commentAfterUpdateHooks []CommentHook
var commentAfterDeleteHooks []CommentHook
var commentAfterUpsertHooks []CommentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Comment) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range commentBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Comment) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range commentBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Comment) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range commentBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Comment) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range commentBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Comment) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range commentAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Comment) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range commentAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Comment) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range commentAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Comment) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range commentAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Comment) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range commentAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCommentHook registers your hook function for all future operations.
func AddCommentHook(hookPoint boil.HookPoint, commentHook CommentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		commentBeforeInsertHooks = append(commentBeforeInsertHooks, commentHook)
	case boil.BeforeUpdateHook:
		commentBeforeUpdateHooks = append(commentBeforeUpdateHooks, commentHook)
	case boil.BeforeDeleteHook:
		commentBeforeDeleteHooks = append(commentBeforeDeleteHooks, commentHook)
	case boil.BeforeUpsertHook:
		commentBeforeUpsertHooks = append(commentBeforeUpsertHooks, commentHook)
	case boil.AfterInsertHook:
		commentAfterInsertHooks = append(commentAfterInsertHooks, commentHook)
	case boil.AfterSelectHook:
		commentAfterSelectHooks = append(commentAfterSelectHooks, commentHook)
	case boil.AfterUpdateHook:
		commentAfterUpdateHooks = append(commentAfterUpdateHooks, commentHook)
	case boil.AfterDeleteHook:
		commentAfterDeleteHooks = append(commentAfterDeleteHooks, commentHook)
	case boil.AfterUpsertHook:
		commentAfterUpsertHooks = append(commentAfterUpsertHooks, commentHook)
	}
}

// OneG returns a single comment record from the query using the global executor.
func (q commentQuery) OneG() (*Comment, error) {
	return q.One(boil.GetDB())
}

// One returns a single comment record from the query.
func (q commentQuery) One(exec boil.Executor) (*Comment, error) {
	o := &Comment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for comment")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Comment records from the query using the global executor.
func (q commentQuery) AllG() (CommentSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Comment records from the query.
func (q commentQuery) All(exec boil.Executor) (CommentSlice, error) {
	var o []*Comment

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Comment slice")
	}

	if len(commentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Comment records in the query, and panics on error.
func (q commentQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Comment records in the query.
func (q commentQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count comment rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q commentQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q commentQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if comment exists")
	}

	return count > 0, nil
}

// Comments retrieves all the records using an executor.
func Comments(mods ...qm.QueryMod) commentQuery {
	mods = append(mods, qm.From("`comment`"))
	return commentQuery{NewQuery(mods...)}
}

// FindCommentG retrieves a single record by ID.
func FindCommentG(index int, selectCols ...string) (*Comment, error) {
	return FindComment(boil.GetDB(), index, selectCols...)
}

// FindComment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindComment(exec boil.Executor, index int, selectCols ...string) (*Comment, error) {
	commentObj := &Comment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `comment` where `index`=?", sel,
	)

	q := queries.Raw(query, index)

	err := q.Bind(nil, exec, commentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from comment")
	}

	return commentObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Comment) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Comment) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comment provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(commentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commentInsertCacheMut.RLock()
	cache, cached := commentInsertCache[key]
	commentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commentAllColumns,
			commentColumnsWithDefault,
			commentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commentType, commentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `comment` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `comment` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `comment` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, commentPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into comment")
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

	o.Index = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == commentMapping["index"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Index,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for comment")
	}

CacheNoHooks:
	if !cached {
		commentInsertCacheMut.Lock()
		commentInsertCache[key] = cache
		commentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Comment record using the global executor.
// See Update for more documentation.
func (o *Comment) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Comment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Comment) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	commentUpdateCacheMut.RLock()
	cache, cached := commentUpdateCache[key]
	commentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update comment, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `comment` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, commentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, append(wl, commentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update comment row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for comment")
	}

	if !cached {
		commentUpdateCacheMut.Lock()
		commentUpdateCache[key] = cache
		commentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q commentQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q commentQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for comment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for comment")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CommentSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommentSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `comment` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in comment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all comment")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Comment) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

var mySQLCommentUniqueColumns = []string{
	"index",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Comment) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comment provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(commentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCommentUniqueColumns, o)

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

	commentUpsertCacheMut.RLock()
	cache, cached := commentUpsertCache[key]
	commentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commentAllColumns,
			commentColumnsWithDefault,
			commentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert comment, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "comment", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `comment` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commentType, commentMapping, ret)
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
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for comment")
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

	o.Index = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == commentMapping["index"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(commentType, commentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for comment")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for comment")
	}

CacheNoHooks:
	if !cached {
		commentUpsertCacheMut.Lock()
		commentUpsertCache[key] = cache
		commentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Comment record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Comment) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Comment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Comment) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Comment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commentPrimaryKeyMapping)
	sql := "DELETE FROM `comment` WHERE `index`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from comment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for comment")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q commentQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comment")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CommentSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommentSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(commentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `comment` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comment")
	}

	if len(commentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Comment) ReloadG() error {
	if o == nil {
		return errors.New("models: no Comment provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Comment) Reload(exec boil.Executor) error {
	ret, err := FindComment(exec, o.Index)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommentSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CommentSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommentSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `comment`.* FROM `comment` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommentSlice")
	}

	*o = slice

	return nil
}

// CommentExistsG checks if the Comment row exists.
func CommentExistsG(index int) (bool, error) {
	return CommentExists(boil.GetDB(), index)
}

// CommentExists checks if the Comment row exists.
func CommentExists(exec boil.Executor, index int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `comment` where `index`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, index)
	}
	row := exec.QueryRow(sql, index)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if comment exists")
	}

	return exists, nil
}
