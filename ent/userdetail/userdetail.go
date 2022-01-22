// Code generated by entc, DO NOT EDIT.

package userdetail

const (
	// Label holds the string label denoting the userdetail type in the database.
	Label = "user_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "rank"
	// FieldPic holds the string denoting the pic field in the database.
	FieldPic = "pic"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userdetail in the database.
	Table = "user_detail"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_detail"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userdetail fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldName,
	FieldAge,
	FieldRank,
	FieldPic,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)
