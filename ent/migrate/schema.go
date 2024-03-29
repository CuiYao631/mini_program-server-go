// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ResourcesColumns holds the columns for the "resources" table.
	ResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "desc", Type: field.TypeString},
		{Name: "explain", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "is_top", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ResourcesTable holds the schema information for the "resources" table.
	ResourcesTable = &schema.Table{
		Name:       "resources",
		Columns:    ResourcesColumns,
		PrimaryKey: []*schema.Column{ResourcesColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "tag", Type: field.TypeString},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_tag",
				Unique:  true,
				Columns: []*schema.Column{TagsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "user_type", Type: field.TypeEnum, Enums: []string{"publisher", "wechat", "login", "touristt"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TagResourcesColumns holds the columns for the "tag_resources" table.
	TagResourcesColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeString},
		{Name: "resources_id", Type: field.TypeString},
	}
	// TagResourcesTable holds the schema information for the "tag_resources" table.
	TagResourcesTable = &schema.Table{
		Name:       "tag_resources",
		Columns:    TagResourcesColumns,
		PrimaryKey: []*schema.Column{TagResourcesColumns[0], TagResourcesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_resources_tag_id",
				Columns:    []*schema.Column{TagResourcesColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_resources_resources_id",
				Columns:    []*schema.Column{TagResourcesColumns[1]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ResourcesTable,
		TagsTable,
		UsersTable,
		TagResourcesTable,
	}
)

func init() {
	TagResourcesTable.ForeignKeys[0].RefTable = TagsTable
	TagResourcesTable.ForeignKeys[1].RefTable = ResourcesTable
}
