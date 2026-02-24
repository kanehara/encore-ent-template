package ent

// Re-generate the ent client from the schema definitions.
// Run from the module root with: go generate ./...
//
//go:generate go run entgo.io/ent/cmd/ent generate --feature sql/upsert,intercept,schema/snapshot ./schema
