// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/compensate"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CompensateUpdate is the builder for updating Compensate entities.
type CompensateUpdate struct {
	config
	hooks     []Hook
	mutation  *CompensateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CompensateUpdate builder.
func (cu *CompensateUpdate) Where(ps ...predicate.Compensate) *CompensateUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CompensateUpdate) SetCreatedAt(u uint32) *CompensateUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CompensateUpdate) SetNillableCreatedAt(u *uint32) *CompensateUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CompensateUpdate) AddCreatedAt(u int32) *CompensateUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CompensateUpdate) SetUpdatedAt(u uint32) *CompensateUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CompensateUpdate) AddUpdatedAt(u int32) *CompensateUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CompensateUpdate) SetDeletedAt(u uint32) *CompensateUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CompensateUpdate) SetNillableDeletedAt(u *uint32) *CompensateUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CompensateUpdate) AddDeletedAt(u int32) *CompensateUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetOrderID sets the "order_id" field.
func (cu *CompensateUpdate) SetOrderID(u uuid.UUID) *CompensateUpdate {
	cu.mutation.SetOrderID(u)
	return cu
}

// SetStart sets the "start" field.
func (cu *CompensateUpdate) SetStart(u uint32) *CompensateUpdate {
	cu.mutation.ResetStart()
	cu.mutation.SetStart(u)
	return cu
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (cu *CompensateUpdate) SetNillableStart(u *uint32) *CompensateUpdate {
	if u != nil {
		cu.SetStart(*u)
	}
	return cu
}

// AddStart adds u to the "start" field.
func (cu *CompensateUpdate) AddStart(u int32) *CompensateUpdate {
	cu.mutation.AddStart(u)
	return cu
}

// ClearStart clears the value of the "start" field.
func (cu *CompensateUpdate) ClearStart() *CompensateUpdate {
	cu.mutation.ClearStart()
	return cu
}

// SetEnd sets the "end" field.
func (cu *CompensateUpdate) SetEnd(u uint32) *CompensateUpdate {
	cu.mutation.ResetEnd()
	cu.mutation.SetEnd(u)
	return cu
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (cu *CompensateUpdate) SetNillableEnd(u *uint32) *CompensateUpdate {
	if u != nil {
		cu.SetEnd(*u)
	}
	return cu
}

// AddEnd adds u to the "end" field.
func (cu *CompensateUpdate) AddEnd(u int32) *CompensateUpdate {
	cu.mutation.AddEnd(u)
	return cu
}

// ClearEnd clears the value of the "end" field.
func (cu *CompensateUpdate) ClearEnd() *CompensateUpdate {
	cu.mutation.ClearEnd()
	return cu
}

// SetMessage sets the "message" field.
func (cu *CompensateUpdate) SetMessage(s string) *CompensateUpdate {
	cu.mutation.SetMessage(s)
	return cu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cu *CompensateUpdate) SetNillableMessage(s *string) *CompensateUpdate {
	if s != nil {
		cu.SetMessage(*s)
	}
	return cu
}

// ClearMessage clears the value of the "message" field.
func (cu *CompensateUpdate) ClearMessage() *CompensateUpdate {
	cu.mutation.ClearMessage()
	return cu
}

// Mutation returns the CompensateMutation object of the builder.
func (cu *CompensateUpdate) Mutation() *CompensateMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompensateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompensateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompensateUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompensateUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompensateUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CompensateUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if compensate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized compensate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := compensate.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CompensateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CompensateUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CompensateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   compensate.Table,
			Columns: compensate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compensate.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: compensate.FieldOrderID,
		})
	}
	if value, ok := cu.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldStart,
		})
	}
	if value, ok := cu.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldStart,
		})
	}
	if cu.mutation.StartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: compensate.FieldStart,
		})
	}
	if value, ok := cu.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldEnd,
		})
	}
	if value, ok := cu.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldEnd,
		})
	}
	if cu.mutation.EndCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: compensate.FieldEnd,
		})
	}
	if value, ok := cu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: compensate.FieldMessage,
		})
	}
	if cu.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: compensate.FieldMessage,
		})
	}
	_spec.Modifiers = cu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{compensate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CompensateUpdateOne is the builder for updating a single Compensate entity.
type CompensateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CompensateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CompensateUpdateOne) SetCreatedAt(u uint32) *CompensateUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CompensateUpdateOne) SetNillableCreatedAt(u *uint32) *CompensateUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CompensateUpdateOne) AddCreatedAt(u int32) *CompensateUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CompensateUpdateOne) SetUpdatedAt(u uint32) *CompensateUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CompensateUpdateOne) AddUpdatedAt(u int32) *CompensateUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CompensateUpdateOne) SetDeletedAt(u uint32) *CompensateUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CompensateUpdateOne) SetNillableDeletedAt(u *uint32) *CompensateUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CompensateUpdateOne) AddDeletedAt(u int32) *CompensateUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetOrderID sets the "order_id" field.
func (cuo *CompensateUpdateOne) SetOrderID(u uuid.UUID) *CompensateUpdateOne {
	cuo.mutation.SetOrderID(u)
	return cuo
}

// SetStart sets the "start" field.
func (cuo *CompensateUpdateOne) SetStart(u uint32) *CompensateUpdateOne {
	cuo.mutation.ResetStart()
	cuo.mutation.SetStart(u)
	return cuo
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (cuo *CompensateUpdateOne) SetNillableStart(u *uint32) *CompensateUpdateOne {
	if u != nil {
		cuo.SetStart(*u)
	}
	return cuo
}

// AddStart adds u to the "start" field.
func (cuo *CompensateUpdateOne) AddStart(u int32) *CompensateUpdateOne {
	cuo.mutation.AddStart(u)
	return cuo
}

// ClearStart clears the value of the "start" field.
func (cuo *CompensateUpdateOne) ClearStart() *CompensateUpdateOne {
	cuo.mutation.ClearStart()
	return cuo
}

// SetEnd sets the "end" field.
func (cuo *CompensateUpdateOne) SetEnd(u uint32) *CompensateUpdateOne {
	cuo.mutation.ResetEnd()
	cuo.mutation.SetEnd(u)
	return cuo
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (cuo *CompensateUpdateOne) SetNillableEnd(u *uint32) *CompensateUpdateOne {
	if u != nil {
		cuo.SetEnd(*u)
	}
	return cuo
}

// AddEnd adds u to the "end" field.
func (cuo *CompensateUpdateOne) AddEnd(u int32) *CompensateUpdateOne {
	cuo.mutation.AddEnd(u)
	return cuo
}

// ClearEnd clears the value of the "end" field.
func (cuo *CompensateUpdateOne) ClearEnd() *CompensateUpdateOne {
	cuo.mutation.ClearEnd()
	return cuo
}

// SetMessage sets the "message" field.
func (cuo *CompensateUpdateOne) SetMessage(s string) *CompensateUpdateOne {
	cuo.mutation.SetMessage(s)
	return cuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cuo *CompensateUpdateOne) SetNillableMessage(s *string) *CompensateUpdateOne {
	if s != nil {
		cuo.SetMessage(*s)
	}
	return cuo
}

// ClearMessage clears the value of the "message" field.
func (cuo *CompensateUpdateOne) ClearMessage() *CompensateUpdateOne {
	cuo.mutation.ClearMessage()
	return cuo
}

// Mutation returns the CompensateMutation object of the builder.
func (cuo *CompensateUpdateOne) Mutation() *CompensateMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CompensateUpdateOne) Select(field string, fields ...string) *CompensateUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Compensate entity.
func (cuo *CompensateUpdateOne) Save(ctx context.Context) (*Compensate, error) {
	var (
		err  error
		node *Compensate
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompensateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Compensate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CompensateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompensateUpdateOne) SaveX(ctx context.Context) *Compensate {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompensateUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompensateUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CompensateUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if compensate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized compensate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := compensate.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CompensateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CompensateUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CompensateUpdateOne) sqlSave(ctx context.Context) (_node *Compensate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   compensate.Table,
			Columns: compensate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compensate.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Compensate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, compensate.FieldID)
		for _, f := range fields {
			if !compensate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != compensate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: compensate.FieldOrderID,
		})
	}
	if value, ok := cuo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldStart,
		})
	}
	if value, ok := cuo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldStart,
		})
	}
	if cuo.mutation.StartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: compensate.FieldStart,
		})
	}
	if value, ok := cuo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldEnd,
		})
	}
	if value, ok := cuo.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldEnd,
		})
	}
	if cuo.mutation.EndCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: compensate.FieldEnd,
		})
	}
	if value, ok := cuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: compensate.FieldMessage,
		})
	}
	if cuo.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: compensate.FieldMessage,
		})
	}
	_spec.Modifiers = cuo.modifiers
	_node = &Compensate{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{compensate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}