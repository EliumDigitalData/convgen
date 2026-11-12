package typeinfo

import "go/types"

// Substitution records solved bindings for a generic function's type
// parameters, discovered by structurally unifying its declared parameter and
// result pattern types against a concrete type pair.
type Substitution map[*types.TypeParam]types.Type

// unify structurally matches pattern (which may contain bare *types.TypeParam
// leaves belonging to some generic function declaration) against concrete (a
// fully resolved type with no type parameters of its own), extending subst
// with any newly discovered bindings. It reports whether pattern and concrete
// are structurally compatible.
//
// Already-bound type parameters are checked for consistency via
// types.Identical, so the same type parameter occurring in multiple positions
// (e.g. T in both the parameter and result type) must resolve to the same
// concrete type everywhere it appears.
//
// Unification only understands type parameters, pointers, slices, arrays,
// maps, and named types with type arguments; anything else (basic types,
// plain named types, structs, interfaces, signatures) is compared with
// types.Identical. Struct-field-level and interface-method-level unification
// is intentionally not supported.
func unify(pattern, concrete types.Type, subst Substitution) bool {
	pattern = types.Unalias(pattern)
	concrete = types.Unalias(concrete)

	if tp, ok := pattern.(*types.TypeParam); ok {
		if bound, ok := subst[tp]; ok {
			return types.Identical(bound, concrete)
		}
		subst[tp] = concrete
		return true
	}

	switch p := pattern.(type) {
	case *types.Pointer:
		c, ok := concrete.(*types.Pointer)
		return ok && unify(p.Elem(), c.Elem(), subst)

	case *types.Slice:
		c, ok := concrete.(*types.Slice)
		return ok && unify(p.Elem(), c.Elem(), subst)

	case *types.Array:
		c, ok := concrete.(*types.Array)
		return ok && p.Len() == c.Len() && unify(p.Elem(), c.Elem(), subst)

	case *types.Map:
		c, ok := concrete.(*types.Map)
		return ok && unify(p.Key(), c.Key(), subst) && unify(p.Elem(), c.Elem(), subst)

	case *types.Named:
		c, ok := concrete.(*types.Named)
		if !ok {
			return false
		}
		if p.TypeArgs().Len() > 0 {
			if p.TypeArgs().Len() != c.TypeArgs().Len() || p.Origin() != c.Origin() {
				return false
			}
			for i := 0; i < p.TypeArgs().Len(); i++ {
				if !unify(p.TypeArgs().At(i), c.TypeArgs().At(i), subst) {
					return false
				}
			}
			return true
		}
		return types.Identical(pattern, concrete)

	default:
		return types.Identical(pattern, concrete)
	}
}

// UnifyFuncTypeArgs attempts to solve fn's type parameters by structurally
// unifying its declared X/Y pattern types (which may reference the type
// parameters of fn.Object()'s generic signature) against the concrete xType/
// yType pair. On success it returns the solved type arguments in the
// function's type-parameter declaration order, ready for types.Instantiate.
//
// It reports ok == false if fn is not generic, if the shapes of the patterns
// and the concrete types are incompatible, if the two unifications disagree
// on a shared type parameter, or if some declared type parameter is left
// unbound by both unifications (e.g. a type parameter that doesn't appear in
// the function's parameter or result type).
func UnifyFuncTypeArgs(fn Func, xType, yType types.Type) (targs []types.Type, ok bool) {
	sig, ok := fn.Object().Type().(*types.Signature)
	if !ok || sig.TypeParams().Len() == 0 {
		return nil, false
	}

	subst := make(Substitution)
	if !unify(fn.X().Type(), xType, subst) {
		return nil, false
	}
	if !unify(fn.Y().Type(), yType, subst) {
		return nil, false
	}

	n := sig.TypeParams().Len()
	targs = make([]types.Type, n)
	for i := 0; i < n; i++ {
		solved, bound := subst[sig.TypeParams().At(i)]
		if !bound {
			return nil, false
		}
		targs[i] = solved
	}
	return targs, true
}
