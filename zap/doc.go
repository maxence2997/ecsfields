// Package zap provides typed, ECS-compliant zap.Field constructors.
//
// Every exported function returns a zap.Field whose key is an Elastic Common
// Schema (ECS) 8.17 field name, and whose zap value type matches the ECS type
// system (keyword -> zap.String, long -> zap.Int64, scaled_float ->
// zap.Float64, date -> zap.Time, etc.).
//
// The package has no implicit encoder requirement: it works with any zap
// encoder. Constructors are organised by top-level ECS fieldset (one Go file
// per fieldset), and each file references the upstream ECS YAML schema in a
// top-of-file comment.
//
// Naming convention:
//
//   - PascalCase maps to a dotted ECS key. ServiceName -> service.name.
//   - Nested keys join with PascalCase. ServiceNodeName -> service.node.name.
//   - Numeric variants are suffixed with Int. EventCodeInt(7) -> event.code = "7".
//
// Generic escape hatches Label, NumericLabel, and Tags target the ECS
// labels.*, numeric_labels.*, and tags slots respectively.
package zap
