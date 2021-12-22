package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/segmentio/ksuid"
)

// Resources holds the schema definition for the Resources entity.
type Resources struct {
	ent.Schema
}

// Fields of the Resources.
func (Resources) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(
			func() string {
				return ksuid.New().String()
			}), //id
		field.String("name"),                       //名称
		field.String("icon").Optional(),            //图标
		field.String("desc"),                       //描述
		field.String("url"),                        //url
		field.Time("created_at").Default(time.Now), //创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), //更新时间

	}
}

// Edges of the Resources.
func (Resources) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tag", Tag.Type).Ref("resources"),
	}
}
