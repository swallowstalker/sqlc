package compiler

import (
	"fmt"
	"strings"

	"github.com/sqlc-dev/sqlc/internal/sql/ast"
	"github.com/sqlc-dev/sqlc/internal/sql/astutils"
)

// This is mainly copy-pasted from internal/postgresql/parse.go
func stringSlice(list *ast.List) []string {
	items := []string{}
	for _, item := range list.Items {
		fmt.Printf("stringSlice ast.List.Items %v\n", item)
		if n, ok := item.(*ast.String); ok {
			items = append(items, n.Str)
		}
	}
	return items
}

type Relation struct {
	Catalog string
	Schema  string
	Name    string
}

func parseRelation(node ast.Node) (*Relation, error) {
	switch n := node.(type) {
	case *ast.Boolean:
		fmt.Printf("compat parseRelation ast.Boolean\n")
		if n == nil {
			return nil, fmt.Errorf("unexpected nil in %T node", n)
		}
		return &Relation{Name: "bool"}, nil

	case *ast.List:
		fmt.Printf("compat parseRelation ast.List\n")
		if n == nil {
			return nil, fmt.Errorf("unexpected nil in %T node", n)
		}
		parts := stringSlice(n)
		fmt.Printf("stringSlice result: %v\n", parts)
		switch len(parts) {
		case 1:
			return &Relation{
				Name: parts[0],
			}, nil
		case 2:
			return &Relation{
				Schema: parts[0],
				Name:   parts[1],
			}, nil
		case 3:
			return &Relation{
				Catalog: parts[0],
				Schema:  parts[1],
				Name:    parts[2],
			}, nil
		default:
			return nil, fmt.Errorf("invalid name: %s", astutils.Join(n, "."))
		}

	case *ast.RangeVar:
		fmt.Printf("compat parseRelation ast.RangeVar\n")
		if n == nil {
			return nil, fmt.Errorf("unexpected nil in %T node", n)
		}
		name := Relation{}
		if n.Catalogname != nil {
			name.Catalog = *n.Catalogname
		}
		if n.Schemaname != nil {
			name.Schema = *n.Schemaname
		}
		if n.Relname != nil {
			name.Name = *n.Relname
		}
		return &name, nil

	case *ast.TypeName:
		fmt.Printf("compat parseRelation ast.TypeName\n")
		if n == nil {
			return nil, fmt.Errorf("unexpected nil in %T node", n)
		}
		if n.Names != nil {
			return parseRelation(n.Names)
		} else {
			return &Relation{Name: n.Name}, nil
		}

	default:
		return nil, fmt.Errorf("unexpected node type: %T", node)
	}
}

func ParseTableName(node ast.Node) (*ast.TableName, error) {
	rel, err := parseRelation(node)
	if err != nil {
		return nil, fmt.Errorf("parse table name: %w", err)
	}
	return &ast.TableName{
		Catalog: rel.Catalog,
		Schema:  rel.Schema,
		Name:    rel.Name,
	}, nil
}

func ParseTypeName(node ast.Node) (*ast.TypeName, error) {
	rel, err := parseRelation(node)
	if err != nil {
		return nil, fmt.Errorf("parse type name: %w", err)
	}
	return &ast.TypeName{
		Catalog: rel.Catalog,
		Schema:  rel.Schema,
		Name:    rel.Name,
	}, nil
}

func ParseRelationString(name string) (*Relation, error) {
	parts := strings.Split(name, ".")
	switch len(parts) {
	case 1:
		return &Relation{
			Name: parts[0],
		}, nil
	case 2:
		return &Relation{
			Schema: parts[0],
			Name:   parts[1],
		}, nil
	case 3:
		return &Relation{
			Catalog: parts[0],
			Schema:  parts[1],
			Name:    parts[2],
		}, nil
	default:
		return nil, fmt.Errorf("invalid name: %s", name)
	}
}
