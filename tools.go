//go:build tools

// This file pins tool dependencies so they are tracked in go.mod/go.sum.
// This ensures reproducible code generation: `go run entgo.io/ent/cmd/ent`
// will always use the version pinned here rather than fetching @latest.
//
// See: https://www.alexedwards.net/blog/using-go-modules-with-build-tags-for-tools
package tools

import (
	// ent code generator – pinned to the same version as entgo.io/ent in go.mod.
	// Used by: go generate ./... (via api/ent/generate.go)
	//      and: go run entgo.io/ent/cmd/ent new --target api/ent/schema <Name>
	_ "entgo.io/ent/cmd/ent"
)
