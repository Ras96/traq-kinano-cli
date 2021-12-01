// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/Ras96/traq-kinano-cli/ent/alias"
	"github.com/Ras96/traq-kinano-cli/ent/predicate"
	"github.com/gofrs/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAlias = "Alias"
)

// AliasMutation represents an operation that mutates the Alias nodes in the graph.
type AliasMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	user_id       *uuid.UUID
	short         *string
	long          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Alias, error)
	predicates    []predicate.Alias
}

var _ ent.Mutation = (*AliasMutation)(nil)

// aliasOption allows management of the mutation configuration using functional options.
type aliasOption func(*AliasMutation)

// newAliasMutation creates new mutation for the Alias entity.
func newAliasMutation(c config, op Op, opts ...aliasOption) *AliasMutation {
	m := &AliasMutation{
		config:        c,
		op:            op,
		typ:           TypeAlias,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAliasID sets the ID field of the mutation.
func withAliasID(id uuid.UUID) aliasOption {
	return func(m *AliasMutation) {
		var (
			err   error
			once  sync.Once
			value *Alias
		)
		m.oldValue = func(ctx context.Context) (*Alias, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Alias.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAlias sets the old Alias of the mutation.
func withAlias(node *Alias) aliasOption {
	return func(m *AliasMutation) {
		m.oldValue = func(context.Context) (*Alias, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AliasMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AliasMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Alias entities.
func (m *AliasMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AliasMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUserID sets the "user_id" field.
func (m *AliasMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *AliasMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Alias entity.
// If the Alias object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AliasMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *AliasMutation) ResetUserID() {
	m.user_id = nil
}

// SetShort sets the "short" field.
func (m *AliasMutation) SetShort(s string) {
	m.short = &s
}

// Short returns the value of the "short" field in the mutation.
func (m *AliasMutation) Short() (r string, exists bool) {
	v := m.short
	if v == nil {
		return
	}
	return *v, true
}

// OldShort returns the old "short" field's value of the Alias entity.
// If the Alias object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AliasMutation) OldShort(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldShort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldShort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShort: %w", err)
	}
	return oldValue.Short, nil
}

// ResetShort resets all changes to the "short" field.
func (m *AliasMutation) ResetShort() {
	m.short = nil
}

// SetLong sets the "long" field.
func (m *AliasMutation) SetLong(s string) {
	m.long = &s
}

// Long returns the value of the "long" field in the mutation.
func (m *AliasMutation) Long() (r string, exists bool) {
	v := m.long
	if v == nil {
		return
	}
	return *v, true
}

// OldLong returns the old "long" field's value of the Alias entity.
// If the Alias object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AliasMutation) OldLong(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLong is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLong requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLong: %w", err)
	}
	return oldValue.Long, nil
}

// ResetLong resets all changes to the "long" field.
func (m *AliasMutation) ResetLong() {
	m.long = nil
}

// Where appends a list predicates to the AliasMutation builder.
func (m *AliasMutation) Where(ps ...predicate.Alias) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AliasMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Alias).
func (m *AliasMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AliasMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.user_id != nil {
		fields = append(fields, alias.FieldUserID)
	}
	if m.short != nil {
		fields = append(fields, alias.FieldShort)
	}
	if m.long != nil {
		fields = append(fields, alias.FieldLong)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AliasMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case alias.FieldUserID:
		return m.UserID()
	case alias.FieldShort:
		return m.Short()
	case alias.FieldLong:
		return m.Long()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AliasMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case alias.FieldUserID:
		return m.OldUserID(ctx)
	case alias.FieldShort:
		return m.OldShort(ctx)
	case alias.FieldLong:
		return m.OldLong(ctx)
	}
	return nil, fmt.Errorf("unknown Alias field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AliasMutation) SetField(name string, value ent.Value) error {
	switch name {
	case alias.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case alias.FieldShort:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShort(v)
		return nil
	case alias.FieldLong:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLong(v)
		return nil
	}
	return fmt.Errorf("unknown Alias field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AliasMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AliasMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AliasMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Alias numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AliasMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AliasMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AliasMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Alias nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AliasMutation) ResetField(name string) error {
	switch name {
	case alias.FieldUserID:
		m.ResetUserID()
		return nil
	case alias.FieldShort:
		m.ResetShort()
		return nil
	case alias.FieldLong:
		m.ResetLong()
		return nil
	}
	return fmt.Errorf("unknown Alias field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AliasMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AliasMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AliasMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AliasMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AliasMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AliasMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AliasMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Alias unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AliasMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Alias edge %s", name)
}