// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/CuiYao631/mini_program-server-go/ent/resources"
)

// Resources is the model entity for the Resources schema.
type Resources struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Explain holds the value of the "explain" field.
	Explain string `json:"explain,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// IsTop holds the value of the "is_top" field.
	IsTop bool `json:"is_top,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourcesQuery when eager-loading is set.
	Edges ResourcesEdges `json:"edges"`
}

// ResourcesEdges holds the relations/edges for other nodes in the graph.
type ResourcesEdges struct {
	// Tag holds the value of the tag edge.
	Tag []*Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e ResourcesEdges) TagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resources) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case resources.FieldIsTop:
			values[i] = new(sql.NullBool)
		case resources.FieldID, resources.FieldName, resources.FieldIcon, resources.FieldDesc, resources.FieldExplain, resources.FieldURL:
			values[i] = new(sql.NullString)
		case resources.FieldCreatedAt, resources.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Resources", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Resources fields.
func (r *Resources) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resources.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case resources.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case resources.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				r.Icon = value.String
			}
		case resources.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				r.Desc = value.String
			}
		case resources.FieldExplain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field explain", values[i])
			} else if value.Valid {
				r.Explain = value.String
			}
		case resources.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				r.URL = value.String
			}
		case resources.FieldIsTop:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_top", values[i])
			} else if value.Valid {
				r.IsTop = value.Bool
			}
		case resources.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case resources.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTag queries the "tag" edge of the Resources entity.
func (r *Resources) QueryTag() *TagQuery {
	return (&ResourcesClient{config: r.config}).QueryTag(r)
}

// Update returns a builder for updating this Resources.
// Note that you need to call Resources.Unwrap() before calling this method if this Resources
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Resources) Update() *ResourcesUpdateOne {
	return (&ResourcesClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Resources entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Resources) Unwrap() *Resources {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Resources is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Resources) String() string {
	var builder strings.Builder
	builder.WriteString("Resources(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", icon=")
	builder.WriteString(r.Icon)
	builder.WriteString(", desc=")
	builder.WriteString(r.Desc)
	builder.WriteString(", explain=")
	builder.WriteString(r.Explain)
	builder.WriteString(", url=")
	builder.WriteString(r.URL)
	builder.WriteString(", is_top=")
	builder.WriteString(fmt.Sprintf("%v", r.IsTop))
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ResourcesSlice is a parsable slice of Resources.
type ResourcesSlice []*Resources

func (r ResourcesSlice) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
