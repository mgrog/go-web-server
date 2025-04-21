// go:build tools
// go:generate go run -mod=mod github.com/99designs/gqlgen

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/rubenv/sql-migrate"
)
