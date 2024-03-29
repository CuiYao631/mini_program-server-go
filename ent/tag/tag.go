// Code generated by entc, DO NOT EDIT.

package tag

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ResourcesTable is the table that holds the resources relation/edge. The primary key declared below.
	ResourcesTable = "tag_resources"
	// ResourcesInverseTable is the table name for the Resources entity.
	// It exists in this package in order to avoid circular dependency with the "resources" package.
	ResourcesInverseTable = "resources"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldTag,
}

var (
	// ResourcesPrimaryKey and ResourcesColumn2 are the table columns denoting the
	// primary key for the resources relation (M2M).
	ResourcesPrimaryKey = []string{"tag_id", "resources_id"}
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
