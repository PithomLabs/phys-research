"""Reproduce the closed-form numeric checks behind C-015/C-017/C-018 (D-006 H2/H3, D-004 parable).

Run: python3 check_toys.py
Stdlib only (math). All checks print observed values and assert tolerances.
"""
import math

print("=== D-004 parable (C-018): equal S_vN, positive S_rel ===")
rho = (0.5, 0.25, 0.25)
s1 = (0.5, 0.25, 0.25)
s2 = (0.25, 0.5, 0.25)
S = lambda p: -sum(x * math.log2(x) for x in p)
Srel = lambda p, q: sum(x * (math.log2(x) - math.log2(y)) for x, y in zip(p, q))
print("S(rho), S(s1), S(s2) =", S(rho), S(s1), S(s2))
print("S(rho||s1) =", Srel(rho, s1), " S(rho||s2) =", Srel(rho, s2))
assert abs(S(rho) - 1.5) < 1e-12 and abs(S(s1) - 1.5) < 1e-12 and abs(S(s2) - 1.5) < 1e-12
assert abs(Srel(rho, s1)) < 1e-12 and abs(Srel(rho, s2) - 0.25) < 1e-12
print("OK: equal entropies (1.5 bits) coexist with S_rel = 0.25 bits > 0\n")

print("=== D-006 H2 (C-017): pure-state family, S=0, QFI=4 ===")
for th in [0.0, 0.3, 1.0, 2.5]:
    c, s = math.cos(th), math.sin(th)
    dpsq = s * s + c * c            # <dpsi|dpsi>
    ov = c * (-s) + s * c            # <psi|dpsi>
    F = 4 * (dpsq - abs(ov) ** 2)
    print(f"theta={th}: S_vN=0 (rank-1 projector), QFI={F}")
    assert abs(F - 4.0) < 1e-12
print("OK: flat (zero) entropy with constant nonzero QFI=4\n")

print("=== D-006 H3 repaired toy (C-015): qutrit, S=1 crossing, smooth Bures ===")


def S3(lam):
    return -(lam * math.log2(lam) + (1 - lam) * math.log2((1 - lam) / 2))


lo, hi = 0.5, 1.0 - 1e-12
for _ in range(200):
    mid = 0.5 * (lo + hi)
    if S3(mid) > 1.0:
        lo = mid
    else:
        hi = mid
lam_star = 0.5 * (lo + hi)
h = 1e-7
dS = (S3(lam_star + h) - S3(lam_star - h)) / (2 * h)
print(f"lambda*={lam_star:.6f} S={S3(lam_star):.12f} dS/dlam={dS:.6f}")
assert abs(S3(lam_star) - 1.0) < 1e-9 and abs(dS) > 1.0  # nonzero slope => S=t solvable


def bures2(t):
    a = lam_star
    b = lam_star + 0.01 * t
    pa = (a, (1 - a) / 2, (1 - a) / 2)
    pb = (b, (1 - b) / 2, (1 - b) / 2)
    fid = sum(math.sqrt(x * y) for x, y in zip(pa, pb)) ** 2  # commuting => classical
    return 2 * (1 - math.sqrt(fid))


e = 1e-4
d1 = (bures2(e) - bures2(-e)) / (2 * e)
d2a = (bures2(e) - 2 * bures2(0.0) + bures2(-e)) / e**2
d2b = (bures2(2 * e) - 2 * bures2(e) + bures2(0.0)) / e**2
print(f"Bures^2: d1={d1:.3e} d2a={d2a:.6f} d2b={d2b:.6f} |d2a-d2b|={abs(d2a - d2b):.2e}")
assert abs(d2a - d2b) < 1e-6  # smooth: finite-difference 2nd derivs agree
print("OK: smooth state path, S=t solvable locally, min(S1,S2) kinks independently\n")
print("ALL CHECKS PASSED")
