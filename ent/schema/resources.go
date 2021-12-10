package schema

import (
	"time"

	"entgo.io/ent"
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
		field.String("title"),                      //标题
		field.String("icon"),                       //图标
		field.String("tag"),                        //标签
		field.String("desc"),                       //描述
		field.String("url"),                        //链接
		field.String("created_user_name"),          //创建人员
		field.String("updated_user_name"),          //更新人员
		field.Time("created_at").Default(time.Now), //创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), //更新时间
	}
}

// Edges of the Resources.
func (Resources) Edges() []ent.Edge {
	return nil
}
