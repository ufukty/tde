# astw/copy/clean_copy

Differences of clean_copy from copy:

- `Object` and `Scope` type of fields are replaced with `nil` to remove circular references that cause error on serialization.

- Many o  ast.File 