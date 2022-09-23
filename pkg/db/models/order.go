// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

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

// Order is an object representing the database table.
type Order struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code       string    `boil:"code" json:"code" toml:"code" yaml:"code"`
	AccountID  int64     `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	Items      string    `boil:"items" json:"items,omitempty" toml:"items" yaml:"items,omitempty"`
	TotalPrice int64     `boil:"total_price" json:"total_price" toml:"total_price" yaml:"total_price"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *orderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderColumns = struct {
	ID         string
	Code       string
	AccountID  string
	Items      string
	TotalPrice string
	CreatedAt  string
}{
	ID:         "id",
	Code:       "code",
	AccountID:  "account_id",
	Items:      "items",
	TotalPrice: "total_price",
	CreatedAt:  "created_at",
}

var OrderTableColumns = struct {
	ID         string
	Code       string
	AccountID  string
	Items      string
	TotalPrice string
	CreatedAt  string
}{
	ID:         "order.id",
	Code:       "order.code",
	AccountID:  "order.account_id",
	Items:      "order.items",
	TotalPrice: "order.total_price",
	CreatedAt:  "order.created_at",
}

// Generated where

func (w whereHelperstring) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelperstring) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var OrderWhere = struct {
	ID         whereHelperint64
	Code       whereHelperstring
	AccountID  whereHelperint64
	Items      whereHelperstring
	TotalPrice whereHelperint64
	CreatedAt  whereHelpertime_Time
}{
	ID:         whereHelperint64{field: "\"order\".\"id\""},
	Code:       whereHelperstring{field: "\"order\".\"code\""},
	AccountID:  whereHelperint64{field: "\"order\".\"account_id\""},
	Items:      whereHelperstring{field: "\"order\".\"items\""},
	TotalPrice: whereHelperint64{field: "\"order\".\"total_price\""},
	CreatedAt:  whereHelpertime_Time{field: "\"order\".\"created_at\""},
}

// OrderRels is where relationship names are stored.
var OrderRels = struct {
	IDAccount string
	ItemID string
}{
	IDAccount: "IDAccount",
	ItemID: "ItemID",
}

// orderR is where relationships are stored.
type orderR struct {
	IDAccount *Account `boil:"IDAccount" json:"IDAccount" toml:"IDAccount" yaml:"IDAccount"`
	ItemID *Item `boil:"ItemID" json:"ItemID" toml:"ItemID" yaml:"ItemID"`
}

// NewStruct creates a new relationship struct
func (*orderR) NewStruct() *orderR {
	return &orderR{}
}

func (r *orderR) GetIDAccount() *Account {
	if r == nil {
		return nil
	}
	return r.IDAccount
}

// orderL is where Load methods for each relationship are stored.
type orderL struct{}

var (
	orderAllColumns            = []string{"id", "code", "account_id", "items", "total_price", "created_at"}
	orderColumnsWithoutDefault = []string{"code"}
	orderColumnsWithDefault    = []string{"id", "account_id", "items", "total_price", "created_at"}
	orderPrimaryKeyColumns     = []string{"id"}
	orderGeneratedColumns      = []string{}
)

type (
	// OrderSlice is an alias for a slice of pointers to Order.
	// This should almost always be used instead of []Order.
	OrderSlice []*Order
	// OrderHook is the signature for custom Order hook methods
	OrderHook func(context.Context, boil.ContextExecutor, *Order) error

	orderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderType                 = reflect.TypeOf(&Order{})
	orderMapping              = queries.MakeStructMapping(orderType)
	orderPrimaryKeyMapping, _ = queries.BindMapping(orderType, orderMapping, orderPrimaryKeyColumns)
	orderInsertCacheMut       sync.RWMutex
	orderInsertCache          = make(map[string]insertCache)
	orderUpdateCacheMut       sync.RWMutex
	orderUpdateCache          = make(map[string]updateCache)
	orderUpsertCacheMut       sync.RWMutex
	orderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderAfterSelectHooks []OrderHook

var orderBeforeInsertHooks []OrderHook
var orderAfterInsertHooks []OrderHook

var orderBeforeUpdateHooks []OrderHook
var orderAfterUpdateHooks []OrderHook

var orderBeforeDeleteHooks []OrderHook
var orderAfterDeleteHooks []OrderHook

var orderBeforeUpsertHooks []OrderHook
var orderAfterUpsertHooks []OrderHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Order) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Order) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Order) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Order) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Order) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Order) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Order) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Order) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Order) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderHook registers your hook function for all future operations.
func AddOrderHook(hookPoint boil.HookPoint, orderHook OrderHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderAfterSelectHooks = append(orderAfterSelectHooks, orderHook)
	case boil.BeforeInsertHook:
		orderBeforeInsertHooks = append(orderBeforeInsertHooks, orderHook)
	case boil.AfterInsertHook:
		orderAfterInsertHooks = append(orderAfterInsertHooks, orderHook)
	case boil.BeforeUpdateHook:
		orderBeforeUpdateHooks = append(orderBeforeUpdateHooks, orderHook)
	case boil.AfterUpdateHook:
		orderAfterUpdateHooks = append(orderAfterUpdateHooks, orderHook)
	case boil.BeforeDeleteHook:
		orderBeforeDeleteHooks = append(orderBeforeDeleteHooks, orderHook)
	case boil.AfterDeleteHook:
		orderAfterDeleteHooks = append(orderAfterDeleteHooks, orderHook)
	case boil.BeforeUpsertHook:
		orderBeforeUpsertHooks = append(orderBeforeUpsertHooks, orderHook)
	case boil.AfterUpsertHook:
		orderAfterUpsertHooks = append(orderAfterUpsertHooks, orderHook)
	}
}

// OneG returns a single order record from the query using the global executor.
func (q orderQuery) OneG(ctx context.Context) (*Order, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single order record from the query.
func (q orderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Order, error) {
	o := &Order{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for order")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Order records from the query using the global executor.
func (q orderQuery) AllG(ctx context.Context) (OrderSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Order records from the query.
func (q orderQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderSlice, error) {
	var o []*Order

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to Order slice")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Order records in the query using the global executor
func (q orderQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Order records in the query.
func (q orderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count order rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q orderQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q orderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if order exists")
	}

	return count > 0, nil
}

// IDAccount pointed to by the foreign key.
func (o *Order) IDAccount(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	return Accounts(queryMods...)
}

// LoadIDAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (orderL) LoadIDAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrder interface{}, mods queries.Applicator) error {
	var slice []*Order
	var object *Order

	if singular {
		var ok bool
		object, ok = maybeOrder.(*Order)
		if !ok {
			object = new(Order)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrder)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrder))
			}
		}
	} else {
		s, ok := maybeOrder.(*[]*Order)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrder)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrder))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orderR{}
		}
		args = append(args, object.ID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orderR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`accounts`),
		qm.WhereIn(`accounts.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for accounts")
	}

	if len(orderAfterSelectHooks) != 0 {
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
		object.R.IDAccount = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.IDOrder = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.ID {
				local.R.IDAccount = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.IDOrder = local
				break
			}
		}
	}

	return nil
}

// SetIDAccountG of the order to the related item.
// Sets o.R.IDAccount to related.
// Adds o to related.R.IDOrder.
// Uses the global database handle.
func (o *Order) SetIDAccountG(ctx context.Context, insert bool, related *Account) error {
	return o.SetIDAccount(ctx, boil.GetContextDB(), insert, related)
}

// SetIDAccount of the order to the related item.
// Sets o.R.IDAccount to related.
// Adds o to related.R.IDOrder.
func (o *Order) SetIDAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"id"}),
		strmangle.WhereClause("\"", "\"", 2, orderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ID = related.ID
	if o.R == nil {
		o.R = &orderR{
			IDAccount: related,
		}
	} else {
		o.R.IDAccount = related
	}

	if related.R == nil {
		related.R = &accountR{
			IDOrder: o,
		}
	} else {
		related.R.IDOrder = o
	}

	return nil
}

// Orders retrieves all the records using an executor.
func Orders(mods ...qm.QueryMod) orderQuery {
	mods = append(mods, qm.From("\"order\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"order\".*"})
	}

	return orderQuery{q}
}

// FindOrderG retrieves a single record by ID.
func FindOrderG(ctx context.Context, iD int64, selectCols ...string) (*Order, error) {
	return FindOrder(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrder(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Order, error) {
	orderObj := &Order{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from order")
	}

	if err = orderObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderObj, err
	}

	return orderObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Order) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Order) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no order provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderInsertCacheMut.RLock()
	cache, cached := orderInsertCache[key]
	orderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderType, orderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into order")
	}

	if !cached {
		orderInsertCacheMut.Lock()
		orderInsertCache[key] = cache
		orderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Order record using the global executor.
// See Update for more documentation.
func (o *Order) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Order.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Order) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderUpdateCacheMut.RLock()
	cache, cached := orderUpdateCache[key]
	orderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update order, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, append(wl, orderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update order row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for order")
	}

	if !cached {
		orderUpdateCacheMut.Lock()
		orderUpdateCache[key] = cache
		orderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q orderQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q orderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for order")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrderSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all order")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Order) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Order) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no order provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	orderUpsertCacheMut.RLock()
	cache, cached := orderUpsertCache[key]
	orderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert order, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderPrimaryKeyColumns))
			copy(conflict, orderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"order\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderType, orderMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert order")
	}

	if !cached {
		orderUpsertCacheMut.Lock()
		orderUpsertCache[key] = cache
		orderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Order record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Order) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Order record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Order) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no Order provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderPrimaryKeyMapping)
	sql := "DELETE FROM \"order\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for order")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q orderQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q orderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no orderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for order")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OrderSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for order")
	}

	if len(orderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Order) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: no Order provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Order) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: empty OrderSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order\".* FROM \"order\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in OrderSlice")
	}

	*o = slice

	return nil
}

// OrderExistsG checks if the Order row exists.
func OrderExistsG(ctx context.Context, iD int64) (bool, error) {
	return OrderExists(ctx, boil.GetContextDB(), iD)
}

// OrderExists checks if the Order row exists.
func OrderExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if order exists")
	}

	return exists, nil
}
