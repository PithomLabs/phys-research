// Package equivariance is the 1D single-particle Bohmian continuity,
// current, and guidance evidence fixture: executed fragments of the
// doc-02/doc-03 chain (audit/02-continuity-derivation.md,
// audit/03-flow-equivariance.md) with replayed DIFFERENTIATE/SIMPLIFY
// evidence.
//
// Status discipline (see manifest.json): on explicit test wave-function
// pairs in test units (hbar = m = 1, test input only, never claimed
// solutions of the Schrödinger equation) the current assembly,
// continuity residual (solution pair → zero; non-solution pair →
// pinned nonzero form), product-rule expansion identity, and guidance
// quotient (with recorded NEQ(rho,0) conditions) are EXECUTED with
// replayed evidence. Flow ODE existence, pushforward/measure transport,
// the uniqueness lemma, global existence, node behavior, and the
// equivariance theorems (E-loc)/(E-glob) are REPRESENTED relations or
// UNSUPPORTED as recorded — never claimed executed. This is a
// symmetry-reduced 1D fragment, not proof of the general theory.
package equivariance

import (
	"fmt"

	"github.com/PithomLabs/phys-artifact"
	"github.com/PithomLabs/phys-math"
)

// Fixture holds every minted artifact of the equivariance evidence chain.
type Fixture struct {
	Store    physmath.Store
	ChartRef artifact.ArtifactRef // coordinate_system/1 1D line chart
	// Solution pair: R1 = 1, S1 = 2q (satisfies continuity).
	R1FnRef, R1Ref artifact.ArtifactRef
	S1FnRef, S1Ref artifact.ArtifactRef
	Rho1Ref        artifact.ArtifactRef // expression/1 rho1 = R1^2
	DtRho1Ref      artifact.ArtifactRef // expression/1 d_t rho1
	DqS1Ref        artifact.ArtifactRef // expression/1 d_q S1
	J1Ref          artifact.ArtifactRef // expression/1 j1 = rho1 d_q S1
	DqJ1Ref        artifact.ArtifactRef // expression/1 d_q j1
	Res1Ref        artifact.ArtifactRef // expression/1 residual d_t rho1 + d_q j1
	Res1ZeroRef    artifact.ArtifactRef // relation/1 EQ(Res1, 0)
	V1Ref          artifact.ArtifactRef // expression/1 v1 = j1 rho1^-1
	Guide1RelRef   artifact.ArtifactRef // relation/1 EQ(v1, j1 rho1^-1)
	// Non-solution pair: R2 = t^2+q^2+1, S2 = t q.
	R2FnRef, R2Ref artifact.ArtifactRef
	S2FnRef, S2Ref artifact.ArtifactRef
	Rho2Ref        artifact.ArtifactRef
	DtRho2Ref      artifact.ArtifactRef
	DqS2Ref        artifact.ArtifactRef
	DtR2Ref        artifact.ArtifactRef // expression/1 d_t R2
	DqR2Ref        artifact.ArtifactRef // expression/1 d_q R2
	D2qS2Ref       artifact.ArtifactRef // expression/1 d^2_q S2
	J2Ref          artifact.ArtifactRef
	DqJ2Ref        artifact.ArtifactRef
	DqJ2SRef       artifact.ArtifactRef // expression/1 simp-staged d_q j2
	Res2Ref        artifact.ArtifactRef
	PRRelRef       artifact.ArtifactRef // relation/1 EQ(executed d_q j2, manual product-rule assembly)
	V2Ref          artifact.ArtifactRef
	Guide2RelRef   artifact.ArtifactRef
	// Node pair: R3 = q, S3 = q (rho3 vanishes at q = 0).
	R3FnRef, R3Ref artifact.ArtifactRef
	S3FnRef, S3Ref artifact.ArtifactRef
	Rho3Ref        artifact.ArtifactRef
	J3Ref          artifact.ArtifactRef
	V3Ref          artifact.ArtifactRef
	Guide3RelRef   artifact.ArtifactRef
	// Traveling pair: R4 = (q-3t)^2+1, S4 = 3q (nontrivial cancellation).
	R4FnRef, R4Ref artifact.ArtifactRef
	S4FnRef, S4Ref artifact.ArtifactRef
	Rho4Ref        artifact.ArtifactRef
	DtRho4Ref      artifact.ArtifactRef
	DqS4Ref        artifact.ArtifactRef
	J4Ref          artifact.ArtifactRef
	DqJ4Ref        artifact.ArtifactRef
	Res4Ref        artifact.ArtifactRef
	V4Ref          artifact.ArtifactRef
	Guide4RelRef   artifact.ArtifactRef
	// Represented-only claim relations (no derivation targets them).
	FlowRelRef  artifact.ArtifactRef // relation/1 flow ODE (LHS represented)
	PushRelRef  artifact.ArtifactRef // relation/1 pushforward transport
	EquivRelRef artifact.ArtifactRef // relation/1 (E-loc) claim
	Derivs      []artifact.ArtifactRef
	Symbols     map[string]artifact.ArtifactRef
	Conds       map[string]artifact.ArtifactRef
}

func sym(s physmath.Store, name string) artifact.ArtifactRef {
	ref, err := physmath.PutExpression(s, physmath.NewEphemeralSymbol(name))
	if err != nil {
		panic(err)
	}
	return ref
}

func condNEQ(s physmath.Store, e physmath.Expr) artifact.ArtifactRef {
	body, err := physmath.EncodeRelationBody(physmath.OpNEQ, e, physmath.MustRational(0, 1))
	if err != nil {
		panic(err)
	}
	rel, err := s.Put(physmath.FamilyRelation, body)
	if err != nil {
		panic(err)
	}
	cond, err := s.Put(physmath.FamilyCondition, rel.Encode())
	if err != nil {
		panic(err)
	}
	return cond
}

func diffParams(fn artifact.ArtifactRef, pid physmath.ParameterID, order uint32) []byte {
	return append(physmath.BoundParameterRef{DeclaringArtifactRef: fn, ParameterID: pid}.Encode(),
		artifact.EncodeUInt32(order)...)
}

func mintFn(s physmath.Store, ids []physmath.ParameterID, names []string, body physmath.Expr) artifact.ArtifactRef {
	bodyRef, err := physmath.PutExpression(s, body)
	if err != nil {
		panic(err)
	}
	fn := physmath.Function{ParameterIDs: ids, ParamNames: names, BodyRef: bodyRef}
	if err := physmath.ValidateFunctionBody(fn, body); err != nil {
		panic(err)
	}
	payload, err := fn.Encode()
	if err != nil {
		panic(err)
	}
	ref, err := s.Put(physmath.FamilyFunction, payload)
	if err != nil {
		panic(err)
	}
	return ref
}

func refOf(s physmath.Store, fnRef artifact.ArtifactRef) (artifact.ArtifactRef, error) {
	fam, payload, err := s.Lookup(fnRef)
	if err != nil {
		return artifact.ArtifactRef{}, err
	}
	if fam != physmath.FamilyFunction {
		return artifact.ArtifactRef{}, fmt.Errorf("not function/1")
	}
	fn, err := physmath.DecodeFunction(artifact.NewDecoder(payload))
	if err != nil {
		return artifact.ArtifactRef{}, err
	}
	return fn.BodyRef, nil
}

func mustResolve(s physmath.Store, ref artifact.ArtifactRef) physmath.Expr {
	fam, payload, err := s.Lookup(ref)
	if err != nil {
		panic(err)
	}
	if fam != physmath.FamilyExpression {
		panic("not expression/1")
	}
	e, err := physmath.DecodeExpr(artifact.NewDecoder(payload))
	if err != nil {
		panic(err)
	}
	return e
}

func diffReplay(s physmath.Store, in, fn artifact.ArtifactRef, pid physmath.ParameterID, order uint32, conds []artifact.ArtifactRef) (artifact.ArtifactRef, artifact.ArtifactRef) {
	res, err := physmath.Apply(s, physmath.Invocation{
		ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleDifferentiate,
		InputRefs: []artifact.ArtifactRef{in, fn}, ConditionRefs: conds,
		Parameters: diffParams(fn, pid, order),
	})
	if err != nil {
		panic(err)
	}
	d := physmath.Derivation{
		PremiseRefs: []artifact.ArtifactRef{in, fn},
		Steps: []physmath.DerivationStep{{
			Ordinal: 0, ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleDifferentiate,
			InputRefs: []artifact.ArtifactRef{in, fn}, OutputRef: res.OutputRef,
			ConditionRefs: res.ConditionRefs, Parameters: diffParams(fn, pid, order),
		}},
		ResultRefs: []artifact.ArtifactRef{res.OutputRef},
	}
	if err := physmath.Replay(s, d, nil); err != nil {
		panic(err)
	}
	return res.OutputRef, mustMintDerivation(s, d)
}

func simpReplay(s physmath.Store, in artifact.ArtifactRef, conds []artifact.ArtifactRef) (artifact.ArtifactRef, artifact.ArtifactRef) {
	res, err := physmath.Apply(s, physmath.Invocation{
		ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleExprSimplify,
		InputRefs: []artifact.ArtifactRef{in}, ConditionRefs: conds,
	})
	if err != nil {
		panic(err)
	}
	d := physmath.Derivation{
		PremiseRefs: []artifact.ArtifactRef{in},
		Steps: []physmath.DerivationStep{{
			Ordinal: 0, ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleExprSimplify,
			InputRefs: []artifact.ArtifactRef{in}, OutputRef: res.OutputRef,
			ConditionRefs: res.ConditionRefs,
		}},
		ResultRefs: []artifact.ArtifactRef{res.OutputRef},
	}
	if err := physmath.Replay(s, d, nil); err != nil {
		panic(err)
	}
	return res.OutputRef, mustMintDerivation(s, d)
}

func mustMintDerivation(s physmath.Store, d physmath.Derivation) artifact.ArtifactRef {
	raw, err := d.Encode()
	if err != nil {
		panic(err)
	}
	ref, err := s.Put(physmath.FamilyDerivation, raw)
	if err != nil {
		panic(err)
	}
	return ref
}

func spaceCollection(s physmath.Store, name string) artifact.ArtifactRef {
	desig, err := physmath.PutExpression(s, physmath.NewEphemeralSymbol(name))
	if err != nil {
		panic(err)
	}
	raw, err := physmath.Collection{ElementRefs: []artifact.ArtifactRef{desig}}.Encode()
	if err != nil {
		panic(err)
	}
	ref, err := s.Put(physmath.FamilyCollection, raw)
	if err != nil {
		panic(err)
	}
	return ref
}

func putExpr(s physmath.Store, e physmath.Expr) artifact.ArtifactRef {
	ref, err := physmath.PutExpression(s, e)
	if err != nil {
		panic(err)
	}
	return ref
}

func putRelation(s physmath.Store, op physmath.ComparisonOp, left, right physmath.Expr) artifact.ArtifactRef {
	body, err := physmath.EncodeRelationBody(op, left, right)
	if err != nil {
		panic(err)
	}
	ref, err := s.Put(physmath.FamilyRelation, body)
	if err != nil {
		panic(err)
	}
	return ref
}

func rat(n, d int64) physmath.Expr { return physmath.MustRational(n, d) }

// Build constructs the equivariance evidence chain. It panics on programmer
// error; replay/validation failures of crafted artifacts surface through
// Verify in protocol tests.
func Build() *Fixture {
	s := physmath.NewMemStore()
	f := &Fixture{Store: s, Symbols: map[string]artifact.ArtifactRef{}, Conds: map[string]artifact.ArtifactRef{}}
	pt := physmath.DeriveParameterID(0)
	pq := physmath.DeriveParameterID(1)
	t := physmath.NewBoundSymbol("t", pt)
	q := physmath.NewBoundSymbol("q", pq)
	pids := []physmath.ParameterID{pt, pq}
	names := []string{"t", "q"}

	// 1D line chart for coordinates ["q"].
	qExprRef := putExpr(s, q)
	chart := physmath.CoordinateSystem{Dimension: 1, CoordinateRefs: []artifact.ArtifactRef{qExprRef}, SpaceRef: spaceCollection(s, "Q1")}
	if err := chart.Validate(); err != nil {
		panic(err)
	}
	chartRef, err := s.Put(physmath.FamilyCoordinateSystem, chart.Encode())
	if err != nil {
		panic(err)
	}
	f.ChartRef = chartRef

	// Solution pair: R1 = 1, S1 = 2q.
	r1fn := mintFn(s, pids, names, rat(1, 1))
	f.R1FnRef = r1fn
	f.R1Ref, err = refOf(s, r1fn)
	if err != nil {
		panic(err)
	}
	s1fn := mintFn(s, pids, names, physmath.Mul(rat(2, 1), q))
	f.S1FnRef = s1fn
	f.S1Ref, err = refOf(s, s1fn)
	if err != nil {
		panic(err)
	}
	rho1, rho1Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.R1Ref), mustResolve(s, f.R1Ref))), nil)
	f.Rho1Ref = rho1
	f.Derivs = append(f.Derivs, rho1Deriv)
	dtRho1, dtRho1Deriv := diffReplay(s, f.Rho1Ref, r1fn, pt, 1, nil)
	f.DtRho1Ref = dtRho1
	dqS1, dqS1Deriv := diffReplay(s, f.S1Ref, s1fn, pq, 1, nil)
	f.DqS1Ref = dqS1
	j1, j1Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.Rho1Ref), mustResolve(s, f.DqS1Ref))), nil)
	f.J1Ref = j1
	dqJ1, dqJ1Deriv := diffReplay(s, f.J1Ref, r1fn, pq, 1, nil)
	f.DqJ1Ref = dqJ1
	res1, res1Deriv := simpReplay(s, putExpr(s, physmath.Add(mustResolve(s, f.DtRho1Ref), mustResolve(s, f.DqJ1Ref))), nil)
	f.Res1Ref = res1
	f.Derivs = append(f.Derivs, dtRho1Deriv, dqS1Deriv, j1Deriv, dqJ1Deriv, res1Deriv)
	f.Res1ZeroRef = putRelation(s, physmath.OpEQ, mustResolve(s, f.Res1Ref), rat(0, 1))
	v1, v1Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.J1Ref), physmath.Pow(mustResolve(s, f.Rho1Ref), rat(-1, 1)))), nil)
	f.V1Ref = v1
	f.Derivs = append(f.Derivs, v1Deriv)
	f.Guide1RelRef = putRelation(s, physmath.OpEQ, mustResolve(s, f.V1Ref),
		physmath.Mul(mustResolve(s, f.J1Ref), physmath.Pow(mustResolve(s, f.Rho1Ref), rat(-1, 1))))
	f.Conds["rho1Nonzero"] = condNEQ(s, mustResolve(s, f.Rho1Ref))

	// Non-solution pair: R2 = t^2+q^2+1, S2 = t q.
	r2body := physmath.Add(physmath.Pow(t, rat(2, 1)), physmath.Pow(q, rat(2, 1)), rat(1, 1))
	r2fn := mintFn(s, pids, names, r2body)
	f.R2FnRef = r2fn
	f.R2Ref, err = refOf(s, r2fn)
	if err != nil {
		panic(err)
	}
	s2fn := mintFn(s, pids, names, physmath.Mul(t, q))
	f.S2FnRef = s2fn
	f.S2Ref, err = refOf(s, s2fn)
	if err != nil {
		panic(err)
	}
	rho2, rho2Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.R2Ref), mustResolve(s, f.R2Ref))), nil)
	f.Rho2Ref = rho2
	f.Symbols["rho2"] = rho2
	dtRho2, dtRho2Deriv := diffReplay(s, f.Rho2Ref, r2fn, pt, 1, nil)
	f.DtRho2Ref = dtRho2
	dqS2, dqS2Deriv := diffReplay(s, f.S2Ref, s2fn, pq, 1, nil)
	f.DqS2Ref = dqS2
	j2, j2Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.Rho2Ref), mustResolve(s, f.DqS2Ref))), nil)
	f.J2Ref = j2
	f.Symbols["j2"] = j2
	dqJ2, dqJ2Deriv := diffReplay(s, f.J2Ref, r2fn, pq, 1, nil)
	f.DqJ2Ref = dqJ2
	dqJ2S, dqJ2SVeriv := simpReplay(s, f.DqJ2Ref, nil)
	f.DqJ2SRef = dqJ2S
	f.Derivs = append(f.Derivs, dqJ2SVeriv)
	res2, res2Deriv := simpReplay(s, putExpr(s, physmath.Add(mustResolve(s, f.DtRho2Ref), mustResolve(s, f.DqJ2Ref))), nil)
	f.Res2Ref = res2
	f.Derivs = append(f.Derivs, rho2Deriv, dtRho2Deriv, dqS2Deriv, j2Deriv, dqJ2Deriv, res2Deriv)
	// Manual product-rule assembly of d_q j2 = d_q(rho2 d_q S2).
	dqR2, dqR2Deriv := diffReplay(s, f.R2Ref, r2fn, pq, 1, nil)
	f.Derivs = append(f.Derivs, dqR2Deriv)
	f.DqR2Ref = dqR2
	dtR2, dtR2Deriv := diffReplay(s, f.R2Ref, r2fn, pt, 1, nil)
	f.DtR2Ref = dtR2
	f.Derivs = append(f.Derivs, dtR2Deriv)
	d2qS2, d2qS2Deriv := diffReplay(s, f.DqS2Ref, s2fn, pq, 1, nil)
	f.Derivs = append(f.Derivs, d2qS2Deriv)
	f.D2qS2Ref = d2qS2
	manual := physmath.Add(
		physmath.Mul(rat(2, 1), mustResolve(s, f.R2Ref), mustResolve(s, dqR2), mustResolve(s, f.DqS2Ref)),
		physmath.Mul(mustResolve(s, f.Rho2Ref), mustResolve(s, d2qS2)))
	manualRef, manualDeriv := simpReplay(s, putExpr(s, manual), nil)
	f.Derivs = append(f.Derivs, manualDeriv)
	f.PRRelRef = putRelation(s, physmath.OpEQ, mustResolve(s, f.DqJ2SRef), mustResolve(s, manualRef))
	v2, v2Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.J2Ref), physmath.Pow(mustResolve(s, f.Rho2Ref), rat(-1, 1)))), nil)
	f.V2Ref = v2
	f.Derivs = append(f.Derivs, v2Deriv)
	f.Guide2RelRef = putRelation(s, physmath.OpEQ, mustResolve(s, f.V2Ref),
		physmath.Mul(mustResolve(s, f.J2Ref), physmath.Pow(mustResolve(s, f.Rho2Ref), rat(-1, 1))))
	f.Conds["rho2Nonzero"] = condNEQ(s, mustResolve(s, f.Rho2Ref))

	// Node pair: R3 = q, S3 = q (rho3 = q^2 vanishes at q = 0).
	r3fn := mintFn(s, pids, names, q)
	f.R3FnRef = r3fn
	f.R3Ref, err = refOf(s, r3fn)
	if err != nil {
		panic(err)
	}
	s3fn := mintFn(s, pids, names, q)
	f.S3FnRef = s3fn
	f.S3Ref, err = refOf(s, s3fn)
	if err != nil {
		panic(err)
	}
	rho3, rho3Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.R3Ref), mustResolve(s, f.R3Ref))), nil)
	f.Rho3Ref = rho3
	dqS3, dqS3Deriv := diffReplay(s, f.S3Ref, s3fn, pq, 1, nil)
	j3, j3Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.Rho3Ref), mustResolve(s, dqS3))), nil)
	f.J3Ref = j3
	v3, v3Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.J3Ref), physmath.Pow(mustResolve(s, f.Rho3Ref), rat(-1, 1)))), nil)
	f.V3Ref = v3
	f.Derivs = append(f.Derivs, rho3Deriv, dqS3Deriv, j3Deriv, v3Deriv)
	f.Guide3RelRef = putRelation(s, physmath.OpEQ, mustResolve(s, f.V3Ref),
		physmath.Mul(mustResolve(s, f.J3Ref), physmath.Pow(mustResolve(s, f.Rho3Ref), rat(-1, 1))))
	f.Conds["rho3Nonzero"] = condNEQ(s, mustResolve(s, f.Rho3Ref))

	// Traveling pair: R4 = (q-3t)^2+1, S4 = 3q. Hand algebra: with
	// u = q-3t, R = u^2+1: d_t rho = -12Ru, d_q j = +12Ru, residual 0
	// with both terms generically nonzero (nontrivial cancellation).
	r4body := physmath.Add(physmath.Pow(physmath.Add(q, physmath.Mul(rat(-3, 1), t)), rat(2, 1)), rat(1, 1))
	r4fn := mintFn(s, pids, names, r4body)
	f.R4FnRef = r4fn
	f.R4Ref, err = refOf(s, r4fn)
	if err != nil {
		panic(err)
	}
	s4fn := mintFn(s, pids, names, physmath.Mul(rat(3, 1), q))
	f.S4FnRef = s4fn
	f.S4Ref, err = refOf(s, s4fn)
	if err != nil {
		panic(err)
	}
	rho4, rho4Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.R4Ref), mustResolve(s, f.R4Ref))), nil)
	f.Rho4Ref = rho4
	dtRho4, dtRho4Deriv := diffReplay(s, f.Rho4Ref, r4fn, pt, 1, nil)
	f.DtRho4Ref = dtRho4
	dqS4, dqS4Deriv := diffReplay(s, f.S4Ref, s4fn, pq, 1, nil)
	f.DqS4Ref = dqS4
	j4, j4Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.Rho4Ref), mustResolve(s, f.DqS4Ref))), nil)
	f.J4Ref = j4
	dqJ4, dqJ4Deriv := diffReplay(s, f.J4Ref, r4fn, pq, 1, nil)
	f.DqJ4Ref = dqJ4
	res4, res4Deriv := simpReplay(s, putExpr(s, physmath.Add(mustResolve(s, f.DtRho4Ref), mustResolve(s, f.DqJ4Ref))), nil)
	f.Res4Ref = res4
	f.Derivs = append(f.Derivs, rho4Deriv, dtRho4Deriv, dqS4Deriv, j4Deriv, dqJ4Deriv, res4Deriv)
	v4, v4Deriv := simpReplay(s, putExpr(s, physmath.Mul(mustResolve(s, f.J4Ref), physmath.Pow(mustResolve(s, f.Rho4Ref), rat(-1, 1)))), nil)
	f.V4Ref = v4
	f.Derivs = append(f.Derivs, v4Deriv)
	f.Guide4RelRef = putRelation(s, physmath.OpEQ, mustResolve(s, f.V4Ref),
		physmath.Mul(mustResolve(s, f.J4Ref), physmath.Pow(mustResolve(s, f.Rho4Ref), rat(-1, 1))))
	f.Conds["rho4Nonzero"] = condNEQ(s, mustResolve(s, f.Rho4Ref))

	// Represented-only claim relations over opaque symbols: flow ODE,
	// pushforward transport, and the (E-loc) equivariance claim. No
	// derivation targets them (enforced by protocol tests); they record
	// what the fixture does not execute.
	dQdt := physmath.NewEphemeralSymbol("dQdt")
	vFlow := physmath.NewEphemeralSymbol("vFlow")
	f.FlowRelRef = putRelation(s, physmath.OpEQ, dQdt, vFlow)
	pushLHS := physmath.NewEphemeralSymbol("pushRho0")
	rhoT := physmath.NewEphemeralSymbol("rhoT")
	f.PushRelRef = putRelation(s, physmath.OpEQ, pushLHS, rhoT)
	f.EquivRelRef = putRelation(s, physmath.OpEQ, pushLHS, rhoT)
	return f
}

// Verify replays every stored derivation and checks the binding of each
// guidance relation to its quotient construction and NEQ condition, plus
// the residual-zero and product-rule relation contents.
func (f *Fixture) Verify() error {
	for i, dref := range f.Derivs {
		if err := replayDerivation(f.Store, i, dref); err != nil {
			return err
		}
	}
	if err := checkGuideBinding(f.Store, f.Guide1RelRef, f.V1Ref, f.J1Ref, f.Rho1Ref, f.Conds["rho1Nonzero"], "pair1"); err != nil {
		return err
	}
	if err := checkGuideBinding(f.Store, f.Guide2RelRef, f.V2Ref, f.J2Ref, f.Rho2Ref, f.Conds["rho2Nonzero"], "pair2"); err != nil {
		return err
	}
	if err := checkGuideBinding(f.Store, f.Guide3RelRef, f.V3Ref, f.J3Ref, f.Rho3Ref, f.Conds["rho3Nonzero"], "pair3"); err != nil {
		return err
	}
	if err := checkGuideBinding(f.Store, f.Guide4RelRef, f.V4Ref, f.J4Ref, f.Rho4Ref, f.Conds["rho4Nonzero"], "pair4"); err != nil {
		return err
	}
	if err := checkResidualZero(f.Store, f.Res1ZeroRef, f.Res1Ref); err != nil {
		return err
	}
	return checkPRRelation(f.Store, f.PRRelRef, f.DqJ2SRef)
}

func replayDerivation(s physmath.Store, i int, dref artifact.ArtifactRef) error {
	fam, payload, err := s.Lookup(dref)
	if err != nil {
		return fmt.Errorf("derivation %d unresolvable: %w", i, err)
	}
	if fam != physmath.FamilyDerivation {
		return fmt.Errorf("derivation %d is not derivation/1", i)
	}
	d, err := physmath.DecodeDerivation(payload)
	if err != nil {
		return fmt.Errorf("derivation %d undecodable: %w", i, err)
	}
	if err := physmath.Replay(s, d, nil); err != nil {
		return fmt.Errorf("derivation %d does not replay: %w", i, err)
	}
	return nil
}

// checkGuideBinding verifies a guidance relation is EQ(v, j rho^-1) with
// v the stored guidance artifact and quotient sides matching the stored
// current/density, and that the binding NEQ condition exists with body
// NEQ(rho, 0). A node-blind assembly (same quotient, no condition)
// fails this check.
func checkGuideBinding(s physmath.Store, relRef, vRef, jRef, rhoRef, condRef artifact.ArtifactRef, tag string) error {
	op, left, right, err := decodeRelationPayload(s, relRef)
	if err != nil {
		return fmt.Errorf("%s guidance relation: %w", tag, err)
	}
	if op != string(physmath.OpEQ) {
		return fmt.Errorf("%s guidance relation is not an equality", tag)
	}
	vb, err := exprBytes(s, vRef)
	if err != nil {
		return fmt.Errorf("%s guidance v: %w", tag, err)
	}
	lb, err := left.Encode()
	if err != nil {
		return fmt.Errorf("%s guidance left unencodable: %w", tag, err)
	}
	if string(lb) != string(vb) {
		return fmt.Errorf("%s guidance left side is not the stored v", tag)
	}
	want, err := physmath.Canonicalize(physmath.Mul(mustResolve(s, jRef),
		physmath.Pow(mustResolve(s, rhoRef), rat(-1, 1))))
	if err != nil {
		return fmt.Errorf("%s guidance expectation unbuildable: %w", tag, err)
	}
	got, err := physmath.Canonicalize(right)
	if err != nil {
		return fmt.Errorf("%s guidance right uncanonicalizable: %w", tag, err)
	}
	if !got.Equal(want) {
		return fmt.Errorf("%s guidance right side is not j rho^-1", tag)
	}
	cop, cleft, cright, err := decodeCondition(s, condRef)
	if err != nil {
		return fmt.Errorf("%s guidance condition: %w", tag, err)
	}
	if cop != string(physmath.OpNEQ) {
		return fmt.Errorf("%s guidance condition is not NEQ", tag)
	}
	rhob, err := exprBytes(s, rhoRef)
	if err != nil {
		return fmt.Errorf("%s guidance rho: %w", tag, err)
	}
	clb, err := cleft.Encode()
	if err != nil {
		return fmt.Errorf("%s condition left unencodable: %w", tag, err)
	}
	if string(clb) != string(rhob) {
		return fmt.Errorf("%s condition does not mention stored rho", tag)
	}
	zero, err := rat(0, 1).Encode()
	if err != nil {
		return fmt.Errorf("%s zero unencodable: %w", tag, err)
	}
	crb, err := cright.Encode()
	if err != nil {
		return fmt.Errorf("%s condition right unencodable: %w", tag, err)
	}
	if string(crb) != string(zero) {
		return fmt.Errorf("%s condition is not NEQ(rho, 0)", tag)
	}
	return nil
}

// checkResidualZero verifies the solution-residual relation is EQ with the
// stored residual on the left and canonical zero on the right.
func checkResidualZero(s physmath.Store, relRef, storedRef artifact.ArtifactRef) error {
	op, left, right, err := decodeRelationPayload(s, relRef)
	if err != nil {
		return fmt.Errorf("solution residual relation: %w", err)
	}
	if op != string(physmath.OpEQ) {
		return fmt.Errorf("solution residual relation is not an equality")
	}
	sb, err := exprBytes(s, storedRef)
	if err != nil {
		return fmt.Errorf("solution residual stored: %w", err)
	}
	lb, err := left.Encode()
	if err != nil {
		return fmt.Errorf("solution residual left unencodable: %w", err)
	}
	if string(lb) != string(sb) {
		return fmt.Errorf("solution residual left side is not the stored artifact")
	}
	got, err := physmath.Canonicalize(right)
	if err != nil {
		return fmt.Errorf("solution residual right uncanonicalizable: %w", err)
	}
	want, err := physmath.Canonicalize(rat(0, 1))
	if err != nil {
		return fmt.Errorf("solution residual zero uncanonicalizable: %w", err)
	}
	if !got.Equal(want) {
		return fmt.Errorf("solution residual right side is not zero")
	}
	gotL, err := physmath.Canonicalize(left)
	if err != nil {
		return fmt.Errorf("solution residual left uncanonicalizable: %w", err)
	}
	if !gotL.Equal(want) {
		return fmt.Errorf("solution residual left side is not zero")
	}
	return nil
}

// checkPRRelation verifies the product-rule relation structurally: EQ with
// the stored executed divergence on the left and a resolvable expression
// on the right. The identity itself (executed vs manual assembly) is
// asserted by protocol tests with hand-verified commentary, not here.
func checkPRRelation(s physmath.Store, relRef, storedRef artifact.ArtifactRef) error {
	op, left, _, err := decodeRelationPayload(s, relRef)
	if err != nil {
		return fmt.Errorf("product-rule relation: %w", err)
	}
	if op != string(physmath.OpEQ) {
		return fmt.Errorf("product-rule relation is not an equality")
	}
	sb, err := exprBytes(s, storedRef)
	if err != nil {
		return fmt.Errorf("product-rule stored: %w", err)
	}
	lb, err := left.Encode()
	if err != nil {
		return fmt.Errorf("product-rule left unencodable: %w", err)
	}
	if string(lb) != string(sb) {
		return fmt.Errorf("product-rule left side is not the stored divergence")
	}
	return nil
}

func decodeRelationPayload(s physmath.Store, relRef artifact.ArtifactRef) (string, physmath.Expr, physmath.Expr, error) {
	fam, payload, err := s.Lookup(relRef)
	if err != nil {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("unresolvable: %w", err)
	}
	if fam != physmath.FamilyRelation {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("not relation/1")
	}
	return decodeRelationBody(payload)
}

func decodeRelationBody(payload []byte) (string, physmath.Expr, physmath.Expr, error) {
	rd := artifact.NewDecoder(payload)
	opStr, err := rd.String()
	if err != nil {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("missing operator: %w", err)
	}
	left, err := physmath.DecodeExpr(rd)
	if err != nil {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("left undecodable: %w", err)
	}
	right, err := physmath.DecodeExpr(rd)
	if err != nil {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("right undecodable: %w", err)
	}
	if !rd.Exhausted() {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("trailing bytes")
	}
	return opStr, left, right, nil
}

func decodeCondition(s physmath.Store, condRef artifact.ArtifactRef) (string, physmath.Expr, physmath.Expr, error) {
	fam, payload, err := s.Lookup(condRef)
	if err != nil {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("unresolvable: %w", err)
	}
	if fam != physmath.FamilyCondition {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("not condition/1")
	}
	relRef, err := artifact.DecodeArtifactRef(artifact.NewDecoder(payload))
	if err != nil {
		return "", physmath.Expr{}, physmath.Expr{}, fmt.Errorf("condition does not wrap a ref: %w", err)
	}
	return decodeRelationPayload(s, relRef)
}

func exprBytes(s physmath.Store, ref artifact.ArtifactRef) ([]byte, error) {
	fam, payload, err := s.Lookup(ref)
	if err != nil {
		return nil, err
	}
	if fam != physmath.FamilyExpression {
		return nil, fmt.Errorf("not expression/1")
	}
	return payload, nil
}
