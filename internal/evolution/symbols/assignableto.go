package symbols

import (
	"fmt"
	"go/types"
)

// var (t)ype <-- (v)alue type
func containsSomethingAssignableTo(v, t types.Type) bool {
	switch v := v.(type) {
	case *types.Array:
		if types.AssignableTo(v, t) {
			return true
		}
		return containsSomethingAssignableTo(v.Elem(), t)

	case *types.Basic:
		return types.AssignableTo(v, t)

	case *types.Chan:
		if types.AssignableTo(v, t) {
			return true
		}
		return containsSomethingAssignableTo(v.Elem(), t)

	case *types.Interface:
		return types.AssignableTo(v, t)

	case *types.Map:
		if types.AssignableTo(v, t) {
			return true
		} else if containsSomethingAssignableTo(v.Key(), t) {
			return true
		}
		return containsSomethingAssignableTo(v.Elem(), t)

	case *types.Named:
		return types.AssignableTo(v, t)

	case *types.Pointer:
		if types.AssignableTo(v, t) {
			return true
		}
		return containsSomethingAssignableTo(v.Elem(), t)

	case *types.Signature:
		if types.AssignableTo(v, t) {
			return true
		}
		return containsSomethingAssignableTo(v.Results(), t)

	case *types.Slice:
		if types.AssignableTo(v, t) {
			return true
		}
		return containsSomethingAssignableTo(v.Elem(), t)

	case *types.Struct:
		if types.AssignableTo(v, t) {
			return true
		}
		for i := 0; i < v.NumFields(); i++ {
			if e := v.Field(i); e != nil {
				if containsSomethingAssignableTo(e.Type(), t) {
					return true
				}
			}
		}
		return false

	case *types.Tuple:
		if types.AssignableTo(v, t) {
			return true
		}
		for i := 0; i < v.Len(); i++ {
			if e := v.At(i); e != nil {
				if containsSomethingAssignableTo(e.Type(), t) {
					return true
				}
			}
		}
		return false

	case *types.TypeParam:
		return types.AssignableTo(v, t)

	case *types.Union:
		return types.AssignableTo(v, t)

	default:
		panic(fmt.Errorf("unhandled case for the type %T", v))
	}
}
