// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Ras96/traq-kinano-cli/ent/alias"
	"github.com/gofrs/uuid"
)

// Alias is the model entity for the Alias schema.
type Alias struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Short holds the value of the "short" field.
	Short string `json:"short,omitempty"`
	// Long holds the value of the "long" field.
	Long string `json:"long,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Alias) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case alias.FieldShort, alias.FieldLong:
			values[i] = new(sql.NullString)
		case alias.FieldID, alias.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Alias", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Alias fields.
func (a *Alias) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case alias.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case alias.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				a.UserID = *value
			}
		case alias.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				a.Short = value.String
			}
		case alias.FieldLong:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field long", values[i])
			} else if value.Valid {
				a.Long = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Alias.
// Note that you need to call Alias.Unwrap() before calling this method if this Alias
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Alias) Update() *AliasUpdateOne {
	return (&AliasClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Alias entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Alias) Unwrap() *Alias {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Alias is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Alias) String() string {
	var builder strings.Builder
	builder.WriteString("Alias(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", short=")
	builder.WriteString(a.Short)
	builder.WriteString(", long=")
	builder.WriteString(a.Long)
	builder.WriteByte(')')
	return builder.String()
}

// AliasSlice is a parsable slice of Alias.
type AliasSlice []*Alias

func (a AliasSlice) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}