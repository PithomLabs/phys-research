// Package qsgt001 is the minimal QSG-T001 theorem fixture: it pins the
// audit's c-number results as golden expression artifacts, replays the
// rational-arithmetic skeleton of the conservation argument, and records
// every universal/analytical step as an explicitly unsupported sketch
// step. It proves nothing about QSG-T001; docs 03–04 prove it. See
// AGENTS.md (this directory) for the binding area contract.
package qsgt001

import (
	"bytes"
	"fmt"

	"github.com/PithomLabs/phys-artifact"
	"github.com/PithomLabs/phys-math"
)

// Display names for bound symbols (ASCII; meaning fixed by AGENTS.md).
const (
	DisplayM2     = "m2"
	DisplayR      = "R"
	DisplayG      = "g"
	DisplayRmn    = "Rmn"
	DisplayR2     = "R2"
	DisplayBoxR   = "BoxR"
	DisplayNNR    = "NNR"
	DisplayBoxRmn = "BoxRmn"
	DisplayDNu    = "DNuRamu"
	DisplayDMu    = "DMuRanu"
	DisplayRAbsq  = "Rabsq"
	DisplayRR     = "RR"
	DisplayRiemC  = "RiemC"
	DisplayHddH   = "HddH"
	DisplayHd2    = "Hd2"
	DisplayHdH2   = "HdH2"
)

// StatusSketchUnsupported is the only sketch-step status. There is no
// supported/executed status: analytical steps are never replayed here.
const StatusSketchUnsupported = "unsupported"

// SketchStep is one human proof obligation. Deliberately a bare record:
// there is no Verify or Replay method on QSGSketch anywhere in this
// module (asserted by TestSketchUnsupported).
type SketchStep struct {
	ID        string
	Statement string
	Evidence  string
	Status    string
}

// QSGSketch records the universal/analytical proof obligations. No Verify.
type QSGSketch struct {
	Steps []SketchStep
}

// BuildSketch returns the six unsupported obligations of (T).
func BuildSketch() QSGSketch {
	return QSGSketch{Steps: []SketchStep{
		{ID: "SK-1", Statement: "universal conservation identity (T) over all Hadamard states and backgrounds", Evidence: "doc 08 (N4.7); human proof", Status: StatusSketchUnsupported},
		{ID: "SK-2", Statement: "residual [nabla' J] = (1/3)nabla[J] via corrected TI and Q1-split", Evidence: "doc 08 N4.2/N4.4; human proof", Status: StatusSketchUnsupported},
		{ID: "SK-3", Statement: "Hadamard recursion closure and remainder limit [J] = -3[v1]/(4pi^2)", Evidence: "doc 08 N2/N4.3 (v0/v1-transport); human proof", Status: StatusSketchUnsupported},
		{ID: "SK-4", Statement: "point-splitting subtraction procedure and coincidence limits", Evidence: "doc 08 N1/N2; human procedure", Status: StatusSketchUnsupported},
		{ID: "SK-5", Statement: "Hadamard state semantics, two-point function properties, and standard imports (F1/F3/L1/G2/A9)", Evidence: "doc 08 N1/N2.1/N4.4; human semantics", Status: StatusSketchUnsupported},
		{ID: "SK-6", Statement: "bi-differential operator algebra and parallel transport", Evidence: "doc 03 §0.4, doc 08 N4.2; human proof", Status: StatusSketchUnsupported},
	}}
}

// Fixture holds the store plus every pinned reference and the derivation
// structs (carried so Verify replays without a bespoke codec; stored
// bytes are still checked by re-encode comparison). All gates compare
// resolved payload bytes, never bare references.
type Fixture struct {
	Store physmath.Store

	DCombo6, DCancel, DThird, DTrace, DFlat                          artifact.ArtifactRef
	DCombo6Deriv, DCancelDeriv, DThirdDeriv, DTraceDeriv, DFlatDeriv physmath.Derivation

	ComboResult, ComboExpected   artifact.ArtifactRef
	CancelResult, CancelExpected artifact.ArtifactRef
	ThirdResult, ThirdExpected   artifact.ArtifactRef
	TraceResult, TraceExpected   artifact.ArtifactRef
	FlatResult, FlatExpected     artifact.ArtifactRef
	NonzeroRef, ZeroRef          artifact.ArtifactRef

	V0Shape, V0AntiB artifact.ArtifactRef
	H1Shape, H1AntiB artifact.ArtifactRef
	H2Shape, H2AntiB artifact.ArtifactRef
	FLRW001          artifact.ArtifactRef

	Sketch QSGSketch
}

func resolvePayload(store physmath.Store, ref artifact.ArtifactRef) ([]byte, error) {
	_, payload, err := store.Lookup(ref)
	return payload, err
}

func applyRule(store physmath.Store, rule string, inputs ...artifact.ArtifactRef) (artifact.ArtifactRef, error) {
	res, err := physmath.Apply(store, physmath.Invocation{
		ProducerNamespace: "physmath",
		RuleID:            rule,
		InputRefs:         inputs,
	})
	if err != nil {
		return artifact.ArtifactRef{}, err
	}
	return res.OutputRef, nil
}

func sym(display string, ordinal uint32) physmath.Expr {
	return physmath.NewBoundSymbol(display, physmath.DeriveParameterID(ordinal))
}

func rat(num, den int64) physmath.Expr  { return physmath.MustRational(num, den) }
func neg(e physmath.Expr) physmath.Expr { return physmath.Neg(e) }

// arithIdentity mints premise input, simplifies it, gates the output
// against expected, and packages the single-step derivation.
func arithIdentity(store physmath.Store, input, expected physmath.Expr, tag string) (in, out, exp, drvRef artifact.ArtifactRef, drv physmath.Derivation, err error) {
	if in, err = physmath.PutExpression(store, input); err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s put input: %w", tag, err)
	}
	out, err = applyRule(store, physmath.RuleExprSimplify, in)
	if err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s simplify: %w", tag, err)
	}
	if exp, err = physmath.PutExpression(store, expected); err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s put expected: %w", tag, err)
	}
	rawOut, err := resolvePayload(store, out)
	if err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s resolve out: %w", tag, err)
	}
	rawExp, err := resolvePayload(store, exp)
	if err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s resolve exp: %w", tag, err)
	}
	if !bytes.Equal(rawOut, rawExp) {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s gate: output bytes differ", tag)
	}
	drv = physmath.Derivation{
		PremiseRefs: []artifact.ArtifactRef{in},
		Steps: []physmath.DerivationStep{
			{Ordinal: 0, ProducerNamespace: "physmath", RuleID: physmath.RuleExprSimplify, InputRefs: []artifact.ArtifactRef{in}, OutputRef: out},
		},
		ResultRefs: []artifact.ArtifactRef{out},
	}
	if err := drv.Validate(); err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s validate: %w", tag, err)
	}
	raw, err := drv.Encode()
	if err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s encode: %w", tag, err)
	}
	if drvRef, err = store.Put(physmath.FamilyDerivation, raw); err != nil {
		return in, out, exp, drvRef, drv, fmt.Errorf("%s store: %w", tag, err)
	}
	return in, out, exp, drvRef, drv, nil
}

// Build mints every golden, derivation, and the sketch. Hand-derivation
// sources (doc 08 repair record): (1.7) for COMBO-6, (N4.5)+(N4.6) for
// CANCEL, (N4.4) Q1-split for THIRD, (N4.8) for TRACE, (N2.5) for
// FLAT-CHAIN and NONZERO, (V0)/(3.1)/(3.2) for shapes, doc 05 G4 for
// FLRW-001.
func Build() (*Fixture, error) {
	store := physmath.NewMemStore()
	fx := &Fixture{Store: store}
	put := func(e physmath.Expr) (artifact.ArtifactRef, error) {
		return physmath.PutExpression(store, e)
	}

	// COMBO-6: 2+4=6 — the (1.7) decomposition (2V1/sigma part + 4V1
	// transport part). Rational skeleton only; the analysis is SK-3.
	_, comboOut, comboExp, comboDrv, comboDeriv, err := arithIdentity(store,
		physmath.Add(rat(2, 1), rat(4, 1)), rat(6, 1), "combo-6")
	if err != nil {
		return nil, err
	}
	fx.ComboResult, fx.ComboExpected, fx.DCombo6, fx.DCombo6Deriv = comboOut, comboExp, comboDrv, comboDeriv

	// CANCEL: -2+2=0 — residual + correction numerators (N4.5)+(N4.6)
	// share the symbolic denominator 8pi^2; the machine checks the
	// rational cancellation, not the analysis (SK-2).
	_, cancelOut, cancelExp, cancelDrv, cancelDeriv, err := arithIdentity(store,
		physmath.Add(neg(rat(2, 1)), rat(2, 1)), rat(0, 1), "cancel")
	if err != nil {
		return nil, err
	}
	fx.CancelResult, fx.CancelExpected, fx.DCancel, fx.DCancelDeriv = cancelOut, cancelExp, cancelDrv, cancelDeriv

	// THIRD: (1/3)*6=2 — the (N4.4) Q1-split one-third applied to
	// the (1.7) six, yielding the residual numerator.
	_, thirdOut, thirdExp, thirdDrv, thirdDeriv, err := arithIdentity(store,
		physmath.Mul(rat(1, 3), rat(6, 1)), rat(2, 1), "third")
	if err != nil {
		return nil, err
	}
	fx.ThirdResult, fx.ThirdExpected, fx.DThird, fx.DThirdDeriv = thirdOut, thirdExp, thirdDrv, thirdDeriv

	// TRACE: -3/4+1=+1/4 — net geometric trace coefficient (N4.8) in
	// units of [v1]/pi^2: remainder trace plus correction trace.
	_, traceOut, traceExp, traceDrv, traceDeriv, err := arithIdentity(store,
		physmath.Add(neg(rat(3, 4)), rat(1, 1)), rat(1, 4), "trace")
	if err != nil {
		return nil, err
	}
	fx.TraceResult, fx.TraceExpected, fx.DTrace, fx.DTraceDeriv = traceOut, traceExp, traceDrv, traceDeriv

	// FLAT-CHAIN: (1/8)*(-3/4)=-3/32 — (N2.5) rational core: flat
	// [v1] rational part times the (N2.2) rational part.
	_, flatOut, flatExp, flatDrv, flatDeriv, err := arithIdentity(store,
		physmath.Mul(rat(1, 8), neg(rat(3, 4))), rat(-3, 32), "flat-chain")
	if err != nil {
		return nil, err
	}
	fx.FlatResult, fx.FlatExpected, fx.DFlat, fx.DFlatDeriv = flatOut, flatExp, flatDrv, flatDeriv

	// NONZERO: -3/32 != 0 — arithmetic core of the homogeneous-W
	// falsification (N2.5) (symbolic m^4/pi^2 documented in AGENTS.md).
	nz, err := put(neg(rat(3, 32)))
	if err != nil {
		return nil, err
	}
	zero, err := put(rat(0, 1))
	if err != nil {
		return nil, err
	}
	nzCanon, err := applyRule(store, physmath.RuleExprCanon, nz)
	if err != nil {
		return nil, err
	}
	zeroCanon, err := applyRule(store, physmath.RuleExprCanon, zero)
	if err != nil {
		return nil, err
	}
	rawNZ, _ := resolvePayload(store, nzCanon)
	rawZero, _ := resolvePayload(store, zeroCanon)
	if bytes.Equal(rawNZ, rawZero) {
		return nil, fmt.Errorf("nonzero gate: -3/32 canon-collided with 0")
	}
	fx.NonzeroRef, fx.ZeroRef = nzCanon, zeroCanon

	m2 := sym(DisplayM2, 0)
	R := sym(DisplayR, 1)
	// V0-SHAPE: m^2/2 - R/12 (blind (V0)); anti: Agent B's R/12 (C5).
	if fx.V0Shape, err = put(physmath.Add(
		physmath.Mul(rat(1, 2), m2),
		physmath.Mul(neg(rat(1, 12)), R))); err != nil {
		return nil, err
	}
	if fx.V0AntiB, err = put(physmath.Mul(rat(1, 12), R)); err != nil {
		return nil, err
	}
	// H1-SHAPE: 2RR - (1/2)gR^2 + 2(gBox-NN)R (blind (3.1)); anti:
	// Agent B's sign-flipped derivative terms (D2).
	Rmn := sym(DisplayRmn, 2)
	g := sym(DisplayG, 3)
	R2 := sym(DisplayR2, 4)
	BoxR := sym(DisplayBoxR, 5)
	NNR := sym(DisplayNNR, 6)
	if fx.H1Shape, err = put(physmath.Add(
		physmath.Mul(rat(2, 1), R, Rmn),
		physmath.Mul(neg(rat(1, 2)), g, R2),
		physmath.Mul(rat(2, 1), g, BoxR),
		physmath.Mul(neg(rat(2, 1)), NNR))); err != nil {
		return nil, err
	}
	if fx.H1AntiB, err = put(physmath.Add(
		physmath.Mul(rat(2, 1), R, Rmn),
		physmath.Mul(neg(rat(1, 2)), g, R2),
		physmath.Mul(neg(rat(2, 1)), g, BoxR),
		physmath.Mul(rat(2, 1), NNR))); err != nil {
		return nil, err
	}
	// H2-SHAPE: blind (3.2) unsimplified form; anti: Agent B's
	// derivative structure (D3).
	DNu := sym(DisplayDNu, 7)
	DMu := sym(DisplayDMu, 8)
	BoxRmn := sym(DisplayBoxRmn, 9)
	RAbsq := sym(DisplayRAbsq, 10)
	RR := sym(DisplayRR, 11)
	if fx.H2Shape, err = put(physmath.Add(
		neg(DNu), neg(DMu), BoxRmn,
		physmath.Mul(rat(1, 2), g, BoxR),
		physmath.Mul(neg(rat(1, 2)), g, RAbsq),
		physmath.Mul(rat(2, 1), RR))); err != nil {
		return nil, err
	}
	RiemC := sym(DisplayRiemC, 12)
	if fx.H2AntiB, err = put(physmath.Add(
		NNR,
		physmath.Mul(neg(rat(1, 1)), g, BoxR),
		neg(BoxRmn),
		physmath.Mul(rat(2, 1), RiemC),
		physmath.Mul(neg(rat(1, 2)), g, RAbsq))); err != nil {
		return nil, err
	}
	// FLRW-001: H^(1)_00 = 18(2HddH - Hd^2 + 6HdH^2), hand-derived in
	// doc 05 G4 from (3.1). Pinned golden (regression, not proof).
	HddH := sym(DisplayHddH, 13)
	Hd2 := sym(DisplayHd2, 14)
	HdH2 := sym(DisplayHdH2, 15)
	if fx.FLRW001, err = put(physmath.Mul(rat(18, 1), physmath.Add(
		physmath.Mul(rat(2, 1), HddH),
		neg(Hd2),
		physmath.Mul(rat(6, 1), HdH2)))); err != nil {
		return nil, err
	}

	if err := checkSeparation(store, fx.V0Shape, fx.V0AntiB, "v0"); err != nil {
		return nil, err
	}
	if err := checkSeparation(store, fx.H1Shape, fx.H1AntiB, "h1"); err != nil {
		return nil, err
	}
	if err := checkSeparation(store, fx.H2Shape, fx.H2AntiB, "h2"); err != nil {
		return nil, err
	}

	fx.Sketch = BuildSketch()
	return fx, nil
}

// checkSeparation gates that a golden shape and its anti-golden have
// distinct canonical payloads (regression encoding of a falsification).
func checkSeparation(store physmath.Store, shape, anti artifact.ArtifactRef, tag string) error {
	cs, err := applyRule(store, physmath.RuleExprCanon, shape)
	if err != nil {
		return fmt.Errorf("%s canon shape: %w", tag, err)
	}
	ca, err := applyRule(store, physmath.RuleExprCanon, anti)
	if err != nil {
		return fmt.Errorf("%s canon anti: %w", tag, err)
	}
	rawS, err := resolvePayload(store, cs)
	if err != nil {
		return err
	}
	rawA, err := resolvePayload(store, ca)
	if err != nil {
		return err
	}
	if bytes.Equal(rawS, rawA) {
		return fmt.Errorf("%s separation: anti-golden canon-collided with golden", tag)
	}
	return nil
}

func gateEqual(store physmath.Store, got, want artifact.ArtifactRef, tag string) error {
	rawGot, err := resolvePayload(store, got)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	rawWant, err := resolvePayload(store, want)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if !bytes.Equal(rawGot, rawWant) {
		return fmt.Errorf("%s: payload bytes differ", tag)
	}
	return nil
}

// Verify replays every derivation (checking stored bytes by re-encode)
// and re-checks every gate from the stored references. It verifies the
// fixture's own regression content — replay here is not proof of
// QSG-T001 (see SK-1).
func Verify(fx *Fixture) error {
	for _, item := range []struct {
		tag string
		ref artifact.ArtifactRef
		drv physmath.Derivation
	}{
		{"combo-6", fx.DCombo6, fx.DCombo6Deriv},
		{"cancel", fx.DCancel, fx.DCancelDeriv},
		{"third", fx.DThird, fx.DThirdDeriv},
		{"trace", fx.DTrace, fx.DTraceDeriv},
		{"flat-chain", fx.DFlat, fx.DFlatDeriv},
	} {
		stored, err := resolvePayload(fx.Store, item.ref)
		if err != nil {
			return fmt.Errorf("%s lookup: %w", item.tag, err)
		}
		enc, err := item.drv.Encode()
		if err != nil {
			return fmt.Errorf("%s encode: %w", item.tag, err)
		}
		if !bytes.Equal(stored, enc) {
			return fmt.Errorf("%s: stored derivation bytes differ from re-encode", item.tag)
		}
		if err := physmath.Replay(fx.Store, item.drv, map[artifact.ArtifactRef]bool{}); err != nil {
			return fmt.Errorf("%s replay: %w", item.tag, err)
		}
	}
	if err := gateEqual(fx.Store, fx.ComboResult, fx.ComboExpected, "combo-6"); err != nil {
		return err
	}
	if err := gateEqual(fx.Store, fx.CancelResult, fx.CancelExpected, "cancel"); err != nil {
		return err
	}
	if err := gateEqual(fx.Store, fx.ThirdResult, fx.ThirdExpected, "third"); err != nil {
		return err
	}
	if err := gateEqual(fx.Store, fx.TraceResult, fx.TraceExpected, "trace"); err != nil {
		return err
	}
	if err := gateEqual(fx.Store, fx.FlatResult, fx.FlatExpected, "flat-chain"); err != nil {
		return err
	}
	rawNZ, err := resolvePayload(fx.Store, fx.NonzeroRef)
	if err != nil {
		return err
	}
	rawZero, err := resolvePayload(fx.Store, fx.ZeroRef)
	if err != nil {
		return err
	}
	if bytes.Equal(rawNZ, rawZero) {
		return fmt.Errorf("nonzero gate: collision")
	}
	if err := checkSeparation(fx.Store, fx.V0Shape, fx.V0AntiB, "v0"); err != nil {
		return err
	}
	if err := checkSeparation(fx.Store, fx.H1Shape, fx.H1AntiB, "h1"); err != nil {
		return err
	}
	if err := checkSeparation(fx.Store, fx.H2Shape, fx.H2AntiB, "h2"); err != nil {
		return err
	}
	for _, s := range fx.Sketch.Steps {
		if s.Status != StatusSketchUnsupported {
			return fmt.Errorf("sketch %s: status %q is not %q", s.ID, s.Status, StatusSketchUnsupported)
		}
	}
	return nil
}
