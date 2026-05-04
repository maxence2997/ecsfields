// labels.go — generic ECS escape hatches.
//
// ECS reference (8.17):
//   labels (object)         -> https://www.elastic.co/guide/en/ecs/8.17/ecs-base.html
//   numeric_labels (object) -> https://www.elastic.co/guide/en/ecs/8.17/ecs-base.html
//   tags (keyword[])        -> https://www.elastic.co/guide/en/ecs/8.17/ecs-base.html
//
// labels.* values are stored as keyword (string).
// numeric_labels.* values are stored as scaled_float (float64).
// tags is a flat array of keywords.

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Label emits a custom string label under the ECS labels.<name> slot.
// Use this for free-form keyword metadata that has no dedicated ECS field.
func Label(name, value string) zapcore.Field {
	return zap.String("labels."+name, value)
}

// NumericLabel emits a custom numeric label under the ECS numeric_labels.<name>
// slot. ECS stores these as scaled_float; convert your domain unit to float64
// at the call site.
func NumericLabel(name string, value float64) zapcore.Field {
	return zap.Float64("numeric_labels."+name, value)
}

// Tags emits the ECS top-level tags array (keyword[]). Pass any number of
// string tags; an empty call emits an empty array.
func Tags(values ...string) zapcore.Field {
	return zap.Strings("tags", values)
}
