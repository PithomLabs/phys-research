// Package protocol_test holds the executable derivation tests for the Bohmian
// equivariance fixture. PROTOCOL: every assertion below is either replayed
// engine evidence (a stored derivation re-executed by Verify/Replay), a
// canonical-form comparison both sides of which are shown, or a hand
// derivation recorded in comments adjacent to the assertion it authorizes.
// Nothing in this file imports the conclusions it checks: hand algebra in
// comments is the authority for mathematical facts the engine cannot reduce.
package protocol_test

import (
	"bytes"
	"testing"

	ms "github.com/PithomLabs/bohm-equiv"
	artifact "github.com/PithomLabs/phys-artifact"
	physmath "github.com/PithomLabs/phys-math"
)

func mustResolve(t *testing.T, s physmath.Store, ref artifact.ArtifactRef) physmath.Expr {
	t.Helper()
	fam, payload, err := s.Lookup(ref)
	if err != nil {
		t.Fatal(err)
	}
	if fam != physmath.FamilyExpression {
		t.Fatalf("ref is %s, want expression/1", fam)
	}
	e, err := physmath.DecodeExpr(artifact.NewDecoder(payload))
	if err != nil {
		t.Fatal(err)
	}
	return e
}

func canon(t *testing.T, e physmath.Expr) physmath.Expr {
	t.Helper()
	c, err := physmath.Canonicalize(e)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func assertCanonEq(t *testing.T, got, want physmath.Expr, msg string) {
	t.Helper()
	if !canon(t, got).Equal(canon(t, want)) {
		t.Fatalf("%s: canonical forms differ", msg)
	}
}

func assertCanonNe(t *testing.T, a, b physmath.Expr, msg string) {
	t.Helper()
	if canon(t, a).Equal(canon(t, b)) {
		t.Fatalf("%s: canonical forms unexpectedly equal", msg)
	}
}

func rat(n, d int64) physmath.Expr { return physmath.MustRational(n, d) }

// boundSyms mirrors the fixture's (name, pid) scheme: t=pid0, q=pid1.
func boundSyms() (t, q physmath.Expr) {
	pt := physmath.DeriveParameterID(0)
	pq := physmath.DeriveParameterID(1)
	return physmath.NewBoundSymbol("t", pt), physmath.NewBoundSymbol("q", pq)
}

func diffParams(fn artifact.ArtifactRef, pid physmath.ParameterID, order uint32) []byte {
	return append(physmath.BoundParameterRef{DeclaringArtifactRef: fn, ParameterID: pid}.Encode(),
		artifact.EncodeUInt32(order)...)
}

// decodeRelationBody mirrors the fixture's relation decoder:
// Enum(operator) || left || right.
func decodeRelationBody(t *testing.T, payload []byte) (string, physmath.Expr, physmath.Expr) {
	t.Helper()
	rd := artifact.NewDecoder(payload)
	opStr, err := rd.String()
	if err != nil {
		t.Fatal(err)
	}
	left, err := physmath.DecodeExpr(rd)
	if err != nil {
		t.Fatal(err)
	}
	right, err := physmath.DecodeExpr(rd)
	if err != nil {
		t.Fatal(err)
	}
	if !rd.Exhausted() {
		t.Fatal("relation body has trailing bytes")
	}
	return opStr, left, right
}

func applySimplify(t *testing.T, s physmath.Store, target artifact.ArtifactRef) artifact.ArtifactRef {
	t.Helper()
	out, err := physmath.Apply(s, physmath.Invocation{
		ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleExprSimplify,
		InputRefs: []artifact.ArtifactRef{target},
	})
	if err != nil {
		t.Fatal(err)
	}
	return out.OutputRef
}

func putExpr(t *testing.T, s physmath.Store, e physmath.Expr) artifact.ArtifactRef {
	t.Helper()
	ref, err := physmath.PutExpression(s, e)
	if err != nil {
		t.Fatal(err)
	}
	return ref
}

// TestVerifyAcceptsFresh pins the end-to-end gate: Build then Verify.
func TestVerifyAcceptsFresh(t *testing.T) {
	if err := ms.Build().Verify(); err != nil {
		t.Fatalf("fresh fixture rejected: %v", err)
	}
}

// TestDerivationInventory pins the full evidence count (30 single-step
// derivations) and replays every one.
func TestDerivationInventory(t *testing.T) {
	f := ms.Build()
	if len(f.Derivs) != 30 {
		t.Fatalf("derivation count = %d, want 30", len(f.Derivs))
	}
	for i, dref := range f.Derivs {
		fam, payload, err := f.Store.Lookup(dref)
		if err != nil {
			t.Fatalf("deriv %d unresolvable: %v", i, err)
		}
		if fam != physmath.FamilyDerivation {
			t.Fatalf("deriv %d family = %s", i, fam)
		}
		d, err := physmath.DecodeDerivation(payload)
		if err != nil {
			t.Fatalf("deriv %d undecodable: %v", i, err)
		}
		if err := d.Validate(); err != nil {
			t.Fatalf("deriv %d invalid: %v", i, err)
		}
		if len(d.Steps) != 1 {
			t.Fatalf("deriv %d has %d steps, want 1 atomic rule", i, len(d.Steps))
		}
		if err := physmath.Replay(f.Store, d, nil); err != nil {
			t.Fatalf("deriv %d replay failed: %v", i, err)
		}
	}
}

// TestPartialsGoldens compares executed partials against hand-built trees.
// Hand authority: raw DIFFERENTIATE-001/SIMPLIFY-001 normal forms observed
// from the engine (factored, no distributivity, zero-annihilation folded).
// The power rule decrements exponents in raw output: d_t(q^2) is the raw
// Mul(0,2,q^1), i.e. 2*q^1*d_t(q) with the zero derivative first.
func TestPartialsGoldens(t *testing.T) {
	tm, q := boundSyms()
	r0, r1, r2 := rat(0, 1), rat(1, 1), rat(2, 1)
	f := ms.Build()
	cases := []struct {
		name string
		ref  artifact.ArtifactRef
		want physmath.Expr
	}{
		{"DtRho1", f.DtRho1Ref, r0},
		{"DqJ1", f.DqJ1Ref, r0},
		{"Res1", f.Res1Ref, r0},
		{"J1", f.J1Ref, r2},
		{"V1", f.V1Ref, physmath.Mul(r2, physmath.Pow(r1, rat(-1, 1)))},
		{"DqS1", f.DqS1Ref, physmath.Add(physmath.Mul(q, r0), physmath.Mul(r2, r1))},
		{"DqS2", f.DqS2Ref, physmath.Add(physmath.Mul(q, r0), physmath.Mul(tm, r1))},
		{"DtR2", f.DtR2Ref, physmath.Add(r0,
			physmath.Mul(r0, r2, physmath.Pow(q, r1)),
			physmath.Mul(r1, r2, physmath.Pow(tm, r1)))},
		{"DqR2", f.DqR2Ref, physmath.Add(r0,
			physmath.Mul(r0, r2, physmath.Pow(tm, r1)),
			physmath.Mul(r1, r2, physmath.Pow(q, r1)))},
		{"D2qS2", f.D2qS2Ref, physmath.Add(
			physmath.Mul(q, r0), physmath.Mul(tm, r0),
			physmath.Mul(r0, r1), physmath.Mul(r0, r1))},
		{"Rho3", f.Rho3Ref, physmath.Mul(q, q)},
		{"J3", f.J3Ref, physmath.Mul(q, q)},
		{"V3", f.V3Ref, physmath.Mul(physmath.Mul(q, q),
			physmath.Pow(physmath.Mul(q, q), rat(-1, 1)))},
	}
	for _, c := range cases {
		assertCanonEq(t, mustResolve(t, f.Store, c.ref), c.want, c.name)
	}
}

// TestCurrentAssemblyGoldens checks the composite current/density builds.
// Hand authority: R2 = t^2+q^2+1, so rho2 = R2^2 and j2 = rho2*d_qS2 with
// d_qS2 = t (raw Add(q*0, t*1)); v2 = j2*rho2^-1 by construction.
func TestCurrentAssemblyGoldens(t *testing.T) {
	tm, q := boundSyms()
	r1, r2 := rat(1, 1), rat(2, 1)
	r2body := physmath.Add(physmath.Pow(tm, r2), physmath.Pow(q, r2), r1)
	f := ms.Build()
	assertCanonEq(t, mustResolve(t, f.Store, f.Rho2Ref),
		physmath.Mul(r2body, r2body), "Rho2")
	assertCanonEq(t, mustResolve(t, f.Store, f.J2Ref),
		physmath.Mul(tm, r2body, r2body), "J2")
	assertCanonEq(t, mustResolve(t, f.Store, f.V2Ref),
		physmath.Mul(mustResolve(t, f.Store, f.J2Ref),
			physmath.Pow(mustResolve(t, f.Store, f.Rho2Ref), rat(-1, 1))), "V2")
	assertCanonEq(t, mustResolve(t, f.Store, f.V4Ref),
		physmath.Mul(mustResolve(t, f.Store, f.J4Ref),
			physmath.Pow(mustResolve(t, f.Store, f.Rho4Ref), rat(-1, 1))), "V4")
}

// TestSolutionResidualGolden decodes the EQ(Res1, 0) verdict relation.
func TestSolutionResidualGolden(t *testing.T) {
	f := ms.Build()
	_, payload, err := f.Store.Lookup(f.Res1ZeroRef)
	if err != nil {
		t.Fatal(err)
	}
	op, left, right := decodeRelationBody(t, payload)
	if op != string(physmath.OpEQ) {
		t.Fatalf("residual relation op = %q", op)
	}
	assertCanonEq(t, left, mustResolve(t, f.Store, f.Res1Ref), "Res1Zero left")
	assertCanonEq(t, right, rat(0, 1), "Res1Zero right")
}

// TestResidualDiscrimination pins the non-solution verdict by separation:
// Res2 is neither zero nor the solution residual, and both of its terms
// are nonzero. Hand authority: with R = t^2+q^2+1,
// d_t rho2 = 2R(2t) = 4tR, d_q j2 = d_q(tR^2) = t(2R)(2q) = 4tqR,
// so Res2 = 4tR(1+q), generically nonzero.
func TestResidualDiscrimination(t *testing.T) {
	f := ms.Build()
	res2 := mustResolve(t, f.Store, f.Res2Ref)
	assertCanonNe(t, res2, rat(0, 1), "Res2 vs zero")
	assertCanonNe(t, res2, mustResolve(t, f.Store, f.Res1Ref), "Res2 vs Res1")
	assertCanonNe(t, mustResolve(t, f.Store, f.DtRho2Ref), rat(0, 1), "DtRho2 vs zero")
	assertCanonNe(t, mustResolve(t, f.Store, f.DqJ2Ref), rat(0, 1), "DqJ2 vs zero")
}

// TestProductRuleIdentity pins the product-rule relation as linkage plus
// hand-authorized equation. PROTOCOL NOTE (negative result, kept honest):
// the simp-staged executed and manual forms are DIFFERENT factored normal
// forms (Mul(t,...) vs Mul(q,t,...)); SIMPLIFY-001 has no distributivity,
// like-term collection, or const-folding, so the engine cannot reduce them
// to a common form — not even after (t,q)=(1,1) substitution. The machine
// contribution: (1) the relation links the stored executed divergence
// (left bytes == DqJ2S payload); (2) both sides substitute totally at
// (1,1) — idempotent re-substitution proves no symbols remain — and stay
// nonzero. The equation itself (both sides denote d_q j2) is hand-proved.
// Hand authority: executed is Diff(Mul(t,R2,R2)) = (d_q t)R2^2 +
// t(2R2)(d_qR2) = 0 + t(2R2)(2q) = 4tqR2; manual is
// 2R2(d_qR2)(d_qS2) + rho2(d^2_qS2) = 2R2(2q)t + R2^2(0) = 4tqR2.
func TestProductRuleIdentity(t *testing.T) {
	f := ms.Build()
	_, payload, err := f.Store.Lookup(f.PRRelRef)
	if err != nil {
		t.Fatal(err)
	}
	op, left, right := decodeRelationBody(t, payload)
	if op != string(physmath.OpEQ) {
		t.Fatalf("product-rule relation op = %q", op)
	}
	_, dqJ2S, err := f.Store.Lookup(f.DqJ2SRef)
	if err != nil {
		t.Fatal(err)
	}
	lb, err := left.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(lb, dqJ2S) {
		t.Fatal("product-rule left is not the stored executed divergence")
	}
	pt := physmath.DeriveParameterID(0)
	pq := physmath.DeriveParameterID(1)
	sub := func(target artifact.ArtifactRef) artifact.ArtifactRef {
		t.Helper()
		return sub11(t, f.Store, target, f.R2FnRef, pt, pq)
	}
	sL := sub(putExpr(t, f.Store, left))
	sR := sub(putExpr(t, f.Store, right))
	if again := sub(putExpr(t, f.Store, left)); again != sL {
		t.Fatal("left substitution nondeterministic")
	}
	if closed := sub(sL); closed != sL {
		t.Fatal("left substitution not total")
	}
	if closed := sub(sR); closed != sR {
		t.Fatal("right substitution not total")
	}
	assertCanonNe(t, mustResolve(t, f.Store, sL), rat(0, 1), "substituted executed divergence vs zero")
	assertCanonNe(t, mustResolve(t, f.Store, sR), rat(0, 1), "substituted manual assembly vs zero")
}

// sub11 substitutes (t,q) = (1,1) via EXPR-SUBSTITUTE-001. Matching is
// by ParameterID, so one entry per pid suffices whatever functions
// declared the symbols; decl must declare both pids.
func sub11(t *testing.T, s physmath.Store, target, decl artifact.ArtifactRef, pt, pq physmath.ParameterID) artifact.ArtifactRef {
	t.Helper()
	one := putExpr(t, s, rat(1, 1))
	mk := func(pid physmath.ParameterID) []byte {
		return append(physmath.BoundParameterRef{DeclaringArtifactRef: decl, ParameterID: pid}.Encode(), one.Encode()...)
	}
	e1, e2 := mk(pt), mk(pq)
	if string(e1) > string(e2) {
		e1, e2 = e2, e1
	}
	params := append(artifact.EncodeUInt32(2), append(e1, e2...)...)
	out, err := physmath.Apply(s, physmath.Invocation{
		ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleExprSubstitute,
		InputRefs:  []artifact.ArtifactRef{target, decl},
		Parameters: params,
	})
	if err != nil {
		t.Fatal(err)
	}
	return out.OutputRef
}

// TestTravelingPair evaluates the pair-4 residual at (t,q) = (1,1).
// Hand authority: with u = q-3t and R = u^2+1, u(1,1) = -2, R = 5,
// d_t rho4 = 2R(2u)(-3) = -12Ru = 120,
// d_q j4 = d_q(3R^2) = 3(2R)(2u) = 12Ru = -120, residual 0.
// The machine contribution: substitution consistency, determinism, the
// nonzero/nontrivial pattern, and wrong-path discrimination. Only the
// final collapse 120 + (-120) = 0 rests on hand arithmetic.
func TestTravelingPair(t *testing.T) {
	f := ms.Build()
	pt := physmath.DeriveParameterID(0)
	pq := physmath.DeriveParameterID(1)
	sub := func(target artifact.ArtifactRef) artifact.ArtifactRef {
		t.Helper()
		return sub11(t, f.Store, target, f.R4FnRef, pt, pq)
	}
	sDt := mustResolve(t, f.Store, sub(f.DtRho4Ref))
	sDq := mustResolve(t, f.Store, sub(f.DqJ4Ref))
	sRes := mustResolve(t, f.Store, sub(f.Res4Ref))
	if r1, r2 := sub(f.Res4Ref), sub(f.Res4Ref); r1 != r2 {
		t.Fatal("substitution nondeterministic")
	}
	assertCanonNe(t, sDt, rat(0, 1), "substituted d_t rho4 vs zero")
	assertCanonNe(t, sDq, rat(0, 1), "substituted d_q j4 vs zero")
	assertCanonNe(t, sDt, sDq, "substituted terms vs each other")
	assertCanonNe(t, sRes, sDt, "substituted residual vs d_t term")
	assertCanonNe(t, sRes, sDq, "substituted residual vs d_q term")
	// RECORDED NEGATIVE RESULT: simp(sub(Res4)) does NOT fold to 0.
	// SIMPLIFY-001 leaves substituted trees unfolded (Pow(1,n) remnants,
	// no rational-tree arithmetic), so the final 120 + (-120) = 0 step
	// rests on the hand algebra above — the same boundary as the
	// minisuperspace substitution precedent. No assertion is made here
	// that the engine cannot prove.
}

// TestAntiGoldens builds two wrong paths with the engine and separates
// each from the corrected residual and from zero. Hand authority:
// wrong-sign: ResW = 4tR - 4tqR = 4tR(1-q); dropped-factor:
// ResD = 4tR + d_qS2... with d_q(t) = 0, ResD = d_t rho2 = 4tR.
func TestAntiGoldens(t *testing.T) {
	f := ms.Build()
	pq := physmath.DeriveParameterID(1)
	res2 := mustResolve(t, f.Store, f.Res2Ref)
	// Wrong-sign current: jW = -j2, differentiated and summed by the engine.
	jW := putExpr(t, f.Store, physmath.Mul(rat(-1, 1), mustResolve(t, f.Store, f.J2Ref)))
	dqJW, err := physmath.Apply(f.Store, physmath.Invocation{
		ProducerNamespace: physmath.NamespacePhysMath, RuleID: physmath.RuleDifferentiate,
		InputRefs:  []artifact.ArtifactRef{jW, f.R2FnRef},
		Parameters: diffParams(f.R2FnRef, pq, 1),
	})
	if err != nil {
		t.Fatal(err)
	}
	resW := mustResolve(t, f.Store, applySimplify(t, f.Store,
		putExpr(t, f.Store, physmath.Add(mustResolve(t, f.Store, f.DtRho2Ref),
			mustResolve(t, f.Store, dqJW.OutputRef)))))
	assertCanonNe(t, resW, res2, "wrong-sign residual vs Res2")
	assertCanonNe(t, resW, rat(0, 1), "wrong-sign residual vs zero")
	// Dropped-factor divergence: d_t rho2 + d_qS2 (missing the j2 assembly).
	resD := mustResolve(t, f.Store, applySimplify(t, f.Store,
		putExpr(t, f.Store, physmath.Add(mustResolve(t, f.Store, f.DtRho2Ref),
			mustResolve(t, f.Store, f.DqS2Ref)))))
	assertCanonNe(t, resD, res2, "dropped-factor residual vs Res2")
	assertCanonNe(t, resD, rat(0, 1), "dropped-factor residual vs zero")
	assertCanonNe(t, resW, resD, "wrong-sign vs dropped-factor")
}

// TestNodeBlindRejected deletes the node-pair condition and requires
// Verify to fail: the velocity verdict is not evidence without it.
func TestNodeBlindRejected(t *testing.T) {
	f := ms.Build()
	delete(f.Conds, "rho3Nonzero")
	if err := f.Verify(); err == nil {
		t.Fatal("Verify accepted fixture with node condition deleted")
	}
}

// TestRepresentedOnly scans every derivation's full reference footprint
// and requires the three claim relations to be absent from all of them.
func TestRepresentedOnly(t *testing.T) {
	f := ms.Build()
	banned := map[artifact.ArtifactRef]string{
		f.FlowRelRef: "FlowRel", f.PushRelRef: "PushRel", f.EquivRelRef: "EquivRel",
	}
	for i, dref := range f.Derivs {
		_, payload, err := f.Store.Lookup(dref)
		if err != nil {
			t.Fatal(err)
		}
		d, err := physmath.DecodeDerivation(payload)
		if err != nil {
			t.Fatal(err)
		}
		var refs []artifact.ArtifactRef
		refs = append(refs, d.PremiseRefs...)
		refs = append(refs, d.ResultRefs...)
		for _, st := range d.Steps {
			refs = append(refs, st.InputRefs...)
			refs = append(refs, st.OutputRef)
			refs = append(refs, st.ConditionRefs...)
		}
		for _, r := range refs {
			if name, ok := banned[r]; ok {
				t.Fatalf("deriv %d references represented-only %s", i, name)
			}
		}
	}
}

// TestLimitationRecord pins the rule allowlist: only SIMPLIFY-001 and
// DIFFERENTIATE-001 appear in any derivation step.
func TestLimitationRecord(t *testing.T) {
	f := ms.Build()
	for i, dref := range f.Derivs {
		_, payload, err := f.Store.Lookup(dref)
		if err != nil {
			t.Fatal(err)
		}
		d, err := physmath.DecodeDerivation(payload)
		if err != nil {
			t.Fatal(err)
		}
		for _, st := range d.Steps {
			if st.ProducerNamespace != physmath.NamespacePhysMath {
				t.Fatalf("deriv %d step uses namespace %q", i, st.ProducerNamespace)
			}
			if st.RuleID != physmath.RuleExprSimplify && st.RuleID != physmath.RuleDifferentiate {
				t.Fatalf("deriv %d step uses rule %q", i, st.RuleID)
			}
		}
	}
}

// TestVerifyRejectsTamperedDerivation flips one payload byte and requires
// Verify to fail (decode or replay must catch it).
func TestVerifyRejectsTamperedDerivation(t *testing.T) {
	f := ms.Build()
	_, payload, err := f.Store.Lookup(f.Derivs[0])
	if err != nil {
		t.Fatal(err)
	}
	bad := bytes.Clone(payload)
	bad[len(bad)/2] ^= 0x01
	badRef, err := f.Store.Put(physmath.FamilyDerivation, bad)
	if err != nil {
		t.Fatal(err)
	}
	f.Derivs[0] = badRef
	if err := f.Verify(); err == nil {
		t.Fatal("Verify accepted tampered derivation")
	}
}

// TestChartStructural decodes the 1D line chart and binds its coordinate.
func TestChartStructural(t *testing.T) {
	_, q := boundSyms()
	f := ms.Build()
	_, payload, err := f.Store.Lookup(f.ChartRef)
	if err != nil {
		t.Fatal(err)
	}
	chart, err := physmath.DecodeCoordinateSystem(payload)
	if err != nil {
		t.Fatal(err)
	}
	if chart.Dimension != 1 {
		t.Fatalf("chart dimension = %d, want 1", chart.Dimension)
	}
	if len(chart.CoordinateRefs) != 1 {
		t.Fatalf("chart has %d coordinates, want 1", len(chart.CoordinateRefs))
	}
	assertCanonEq(t, mustResolve(t, f.Store, chart.CoordinateRefs[0]), q, "chart coordinate")
}

// TestRebuildDeterminism builds the fixture twice and byte-compares every
// exported payload: the artifact is a pure function of the code.
func TestRebuildDeterminism(t *testing.T) {
	a := ms.Build()
	b := ms.Build()
	refs := func(f *ms.Fixture) []artifact.ArtifactRef {
		out := []artifact.ArtifactRef{f.ChartRef,
			f.R1FnRef, f.R1Ref, f.S1FnRef, f.S1Ref, f.Rho1Ref, f.DtRho1Ref,
			f.DqS1Ref, f.J1Ref, f.DqJ1Ref, f.Res1Ref, f.Res1ZeroRef,
			f.V1Ref, f.Guide1RelRef,
			f.R2FnRef, f.R2Ref, f.S2FnRef, f.S2Ref, f.Rho2Ref, f.DtRho2Ref,
			f.DqS2Ref, f.DtR2Ref, f.DqR2Ref, f.D2qS2Ref, f.J2Ref, f.DqJ2Ref,
			f.DqJ2SRef, f.Res2Ref, f.PRRelRef, f.V2Ref, f.Guide2RelRef,
			f.R3FnRef, f.R3Ref, f.S3FnRef, f.S3Ref, f.Rho3Ref, f.J3Ref,
			f.V3Ref, f.Guide3RelRef,
			f.R4FnRef, f.R4Ref, f.S4FnRef, f.S4Ref, f.Rho4Ref, f.DtRho4Ref,
			f.DqS4Ref, f.J4Ref, f.DqJ4Ref, f.Res4Ref, f.V4Ref, f.Guide4RelRef,
			f.FlowRelRef, f.PushRelRef, f.EquivRelRef}
		out = append(out, f.Derivs...)
		for _, k := range []string{"rho1Nonzero", "rho2Nonzero", "rho3Nonzero", "rho4Nonzero"} {
			out = append(out, f.Conds[k])
		}
		for _, k := range []string{"rho2", "j2"} {
			out = append(out, f.Symbols[k])
		}
		return out
	}
	ra, rb := refs(a), refs(b)
	if len(ra) != len(rb) {
		t.Fatal("ref list length mismatch")
	}
	for i := range ra {
		fa, pa, err := a.Store.Lookup(ra[i])
		if err != nil {
			t.Fatal(err)
		}
		fb, pb, err := b.Store.Lookup(rb[i])
		if err != nil {
			t.Fatal(err)
		}
		if fa != fb || !bytes.Equal(pa, pb) {
			t.Fatalf("ref %d differs across rebuilds", i)
		}
	}
}
