// Code generated by entc, DO NOT EDIT.

package resources

import (
	"time"
)

const (
	// Label holds the string label denoting the resources type in the database.
	Label = "resources"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// Table holds the table name of the resources in the database.
	Table = "resources"
	// TagTable is the table that holds the tag relation/edge. The primary key declared below.
	TagTable = "tag_resources"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
)

// Columns holds all SQL columns for resources fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIcon,
	FieldDesc,
	FieldURL,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// TagPrimaryKey and TagColumn2 are the table columns denoting the
	// primary key for the tag relation (M2M).
	TagPrimaryKey = []string{"tag_id", "resources_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
