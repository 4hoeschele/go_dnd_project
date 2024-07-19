package schema


import (
    "github.com/google/uuid"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    // "entgo.io/ent/schema/edge"
)

// as soon as I change something in a schema I need to run 'go generate ./...'

// Character holds the schema definition for the Character entity.
type Character struct {
    // is my model
    // configures the schema for the character
    ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
    // defines the columns of the table
    // {} is struct
    return []ent.Field{
        field.UUID("id", uuid.UUID{}),
        field.String("name"),
        field.String("class"),
        field.String("race"),
    }
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
    // defines foreign key relationship
    return nil
}
