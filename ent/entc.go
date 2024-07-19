//go:build ignore
// +build ignore

package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/edge"
	"log"
)

func main() {
	err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(
			&EncodeExtension{},
		),
	)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// EncodeExtension is an implementation of entc.Extension that adds a MarshalJSON
// method to each generated type <T> and inlines the Edges field to the top level JSON.
type EncodeExtension struct {
	entc.DefaultExtension
}

// Templates of the extension.
func (e *EncodeExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("model/additional/jsonencode").
			Parse(`
{{ if $.Edges }}
	// MarshalJSON implements the json.Marshaler interface.
	func ({{ $.Receiver }} *{{ $.Name }}) MarshalJSON() ([]byte, error) {
		type Alias {{ $.Name }}
		return json.Marshal(&struct {
			*Alias
			{{ $.Name }}Edges
		}{
			Alias: (*Alias)({{ $.Receiver }}),
			{{ $.Name }}Edges: {{ $.Receiver }}.Edges,
		})
	}
{{ end }}
`)),
	}
}

// Hooks of the extension.
func (e *EncodeExtension) Hooks() []gen.Hook {
	return []gen.Hook{
		func(next gen.Generator) gen.Generator {
			return gen.GenerateFunc(func(g *gen.Graph) error {
				tag := edge.Annotation{StructTag: `json:"-"`}
				for _, n := range g.Nodes {
					n.Annotations.Set(tag.Name(), tag)
				}
				return next.Generate(g)
			})
		},
	}
}

var _ entc.Extension = (*EncodeExtension)(nil)

