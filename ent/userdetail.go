// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"orse/ent/user"
	"orse/ent/userdetail"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// UserDetail is the model entity for the UserDetail schema.
type UserDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	// 名称
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank float64 `json:"rank,omitempty"`
	// Pic holds the value of the "pic" field.
	// 头像
	Pic string `json:"pic,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserDetailQuery when eager-loading is set.
	Edges UserDetailEdges `json:"edges"`
}

// UserDetailEdges holds the relations/edges for other nodes in the graph.
type UserDetailEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserDetailEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserDetail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userdetail.FieldRank:
			values[i] = new(sql.NullFloat64)
		case userdetail.FieldID, userdetail.FieldUserID, userdetail.FieldAge:
			values[i] = new(sql.NullInt64)
		case userdetail.FieldName, userdetail.FieldPic:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserDetail fields.
func (ud *UserDetail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ud.ID = int(value.Int64)
		case userdetail.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ud.UserID = int(value.Int64)
			}
		case userdetail.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ud.Name = value.String
			}
		case userdetail.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				ud.Age = int(value.Int64)
			}
		case userdetail.FieldRank:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				ud.Rank = value.Float64
			}
		case userdetail.FieldPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic", values[i])
			} else if value.Valid {
				ud.Pic = value.String
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserDetail entity.
func (ud *UserDetail) QueryUser() *UserQuery {
	return (&UserDetailClient{config: ud.config}).QueryUser(ud)
}

// Update returns a builder for updating this UserDetail.
// Note that you need to call UserDetail.Unwrap() before calling this method if this UserDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (ud *UserDetail) Update() *UserDetailUpdateOne {
	return (&UserDetailClient{config: ud.config}).UpdateOne(ud)
}

// Unwrap unwraps the UserDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ud *UserDetail) Unwrap() *UserDetail {
	tx, ok := ud.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserDetail is not a transactional entity")
	}
	ud.config.driver = tx.drv
	return ud
}

// String implements the fmt.Stringer.
func (ud *UserDetail) String() string {
	var builder strings.Builder
	builder.WriteString("UserDetail(")
	builder.WriteString(fmt.Sprintf("id=%v", ud.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ud.UserID))
	builder.WriteString(", name=")
	builder.WriteString(ud.Name)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", ud.Age))
	builder.WriteString(", rank=")
	builder.WriteString(fmt.Sprintf("%v", ud.Rank))
	builder.WriteString(", pic=")
	builder.WriteString(ud.Pic)
	builder.WriteByte(')')
	return builder.String()
}

// UserDetails is a parsable slice of UserDetail.
type UserDetails []*UserDetail

func (ud UserDetails) config(cfg config) {
	for _i := range ud {
		ud[_i].config = cfg
	}
}
