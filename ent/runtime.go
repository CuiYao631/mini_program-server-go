// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/CuiYao631/mini_program-server-go/ent/resources"
	"github.com/CuiYao631/mini_program-server-go/ent/schema"
	"github.com/CuiYao631/mini_program-server-go/ent/tag"
	"github.com/CuiYao631/mini_program-server-go/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	resourcesFields := schema.Resources{}.Fields()
	_ = resourcesFields
	// resourcesDescCreatedAt is the schema descriptor for created_at field.
	resourcesDescCreatedAt := resourcesFields[7].Descriptor()
	// resources.DefaultCreatedAt holds the default value on creation for the created_at field.
	resources.DefaultCreatedAt = resourcesDescCreatedAt.Default.(func() time.Time)
	// resourcesDescUpdatedAt is the schema descriptor for updated_at field.
	resourcesDescUpdatedAt := resourcesFields[8].Descriptor()
	// resources.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	resources.DefaultUpdatedAt = resourcesDescUpdatedAt.Default.(func() time.Time)
	// resources.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	resources.UpdateDefaultUpdatedAt = resourcesDescUpdatedAt.UpdateDefault.(func() time.Time)
	// resourcesDescID is the schema descriptor for id field.
	resourcesDescID := resourcesFields[0].Descriptor()
	// resources.DefaultID holds the default value on creation for the id field.
	resources.DefaultID = resourcesDescID.Default.(func() string)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagFields[0].Descriptor()
	// tag.DefaultID holds the default value on creation for the id field.
	tag.DefaultID = tagDescID.Default.(func() string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
}
