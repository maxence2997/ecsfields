// Package ecsfields is a namespace package. It contains no exported symbols.
//
// Typed, ECS-compliant field constructors live in logger-specific
// sub-packages. v0.1.0 ships a single sub-package:
//
//   - [github.com/maxence2997/ecsfields/zap] — zap.Field constructors that
//     emit Elastic Common Schema (ECS) 8.17 compliant keys and value types.
//
// Future logger support (slog, etc.) will be added as additional sibling
// sub-packages without breaking changes.
package ecsfields
