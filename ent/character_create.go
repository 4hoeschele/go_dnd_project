// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/4hoeschele/go_dnd_project/ent/character"
	"github.com/google/uuid"
)

// CharacterCreate is the builder for creating a Character entity.
type CharacterCreate struct {
	config
	mutation *CharacterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CharacterCreate) SetName(s string) *CharacterCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetClass sets the "class" field.
func (cc *CharacterCreate) SetClass(s string) *CharacterCreate {
	cc.mutation.SetClass(s)
	return cc
}

// SetRace sets the "race" field.
func (cc *CharacterCreate) SetRace(s string) *CharacterCreate {
	cc.mutation.SetRace(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CharacterCreate) SetID(u uuid.UUID) *CharacterCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the CharacterMutation object of the builder.
func (cc *CharacterCreate) Mutation() *CharacterMutation {
	return cc.mutation
}

// Save creates the Character in the database.
func (cc *CharacterCreate) Save(ctx context.Context) (*Character, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CharacterCreate) SaveX(ctx context.Context) *Character {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CharacterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CharacterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CharacterCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Character.name"`)}
	}
	if _, ok := cc.mutation.Class(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required field "Character.class"`)}
	}
	if _, ok := cc.mutation.Race(); !ok {
		return &ValidationError{Name: "race", err: errors.New(`ent: missing required field "Character.race"`)}
	}
	return nil
}

func (cc *CharacterCreate) sqlSave(ctx context.Context) (*Character, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CharacterCreate) createSpec() (*Character, *sqlgraph.CreateSpec) {
	var (
		_node = &Character{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(character.Table, sqlgraph.NewFieldSpec(character.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Class(); ok {
		_spec.SetField(character.FieldClass, field.TypeString, value)
		_node.Class = value
	}
	if value, ok := cc.mutation.Race(); ok {
		_spec.SetField(character.FieldRace, field.TypeString, value)
		_node.Race = value
	}
	return _node, _spec
}

// CharacterCreateBulk is the builder for creating many Character entities in bulk.
type CharacterCreateBulk struct {
	config
	err      error
	builders []*CharacterCreate
}

// Save creates the Character entities in the database.
func (ccb *CharacterCreateBulk) Save(ctx context.Context) ([]*Character, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Character, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CharacterCreateBulk) SaveX(ctx context.Context) []*Character {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CharacterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CharacterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
