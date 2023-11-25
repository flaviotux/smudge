package todo

const (
	// Label holds the string label denoting the user type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldDone holds the string denoting the done field in the database.
	FieldDone = "done"
	// FieldDone holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the todo in the database.
	Table = "todos"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldDone,
	FieldUserID,
}

var (
	// PartKey and SortKey are the table columns denoting the primary key
	PartKey = []string{}
	SortKey = []string{"id", "user_id"}
)

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
	// DefaultDone holds the default value on creation for the "done" field.
	DefaultDone bool
)
