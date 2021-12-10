package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/CuiYao631/mini_program-server-go/entity"
	"github.com/segmentio/ksuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(
			func() string {
				return ksuid.New().String()
			}), //id
		field.String("name"),
		field.String("password"),
		field.Enum("user_type").Values(
			entity.UserTypePublisher.String(),
			entity.UserTypeWechat.String(),
			entity.UserTypeLogin.String(),
			entity.UserTypeTourist.String(),
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
