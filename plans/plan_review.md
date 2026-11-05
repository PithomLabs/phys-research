## claude

# Adversarial review of the plan (OpenCode + Go + Git)

I used the finding format from `MRC_REVIEW_PROTOCOL.md` §8 (Where / What / Why / Minimum fix) and issue no overall verdict. The protocol is written for physics reasoning, so I applied it by analogy to a tooling plan.

## Material findings

**1. Where:** §7 safeguard, §6 "defer PRs", §8 step 9. **What:** unenforced gate; the policy contradicts itself.
- The plan correctly notes a human-only command isn't human-only if agents hold filesystem access. It then removes `approve`, defers mandatory PRs, and relies on "review through Git." The permissions table is policy, not control.
- Agents can also edit the verifier (`cmd/physctl`), `.opencode/agents/*` and `AGENTS.md`.
- **Fix:**
  - Deny edits to those paths in each agent's own permissions. Per the OpenCode docs I found, subagents use their own permissions, not a restricted copy of the parent's.
  - Have agents work on branches with no merge rights on main.
  - Make `check` fail if protected files change on an agent branch.

**2. Where:** §3 `lint` vs. "no invented syntax." **What:** scope contradiction.
- The phys-pseudo README says no parser, checker or linter is required or implied, and PSEUDO.md says it is not a grammar. A "structure validator" is a de facto grammar that can drift from the spec.
- **Fix:** a closed rule list, each rule citing a PSEUDO.md section (e.g., names take no argument lists, `derive` states an explicit result, `?` sits at column 0). Report warnings only, and pin to a spec commit.

**3. Where:** §3 `hash`. **What:** unsupported premise ("applicable PhysArtifact conventions").
- `ComputeRef` takes `(namespace, schemaVersion, artifactType, payload)`, and identity is that tuple plus the digest. No artifact type exists for `.pseudo` or research records.
- Hashing mutable working files changes the ID on a whitespace edit.
- **Fix:** record one namespace/type/schema decision in `AGENTS.md`, and mint refs only at freeze points (agreed hypotheses, accepted derivations). Use readable IDs (C-001) elsewhere.

**4. Where:** §3 `refs` and `status --open`. **What:** undefined mechanism.
- "Typed relationships" and "open" are never defined. Without a reference syntax, an ID allocator, or a status field, implementers will invent conventions or regex the prose. Two agents can also allocate the same next ID.
- **Fix:** about eight lines of front-matter spec (`id`, `type`, `status: open|blocked|closed`, `refs: [{rel,to}]` with a closed set of relations), plus a `new` command that allocates IDs.

**5. Where:** §5 Formalizer, §8 steps 6–8. **What:** circularity and an omitted spec stage.
- GUIDELINES §10 says the test oracle must not be generated solely by the implementation under test, and that a "hand-derived" comment doesn't establish independence. The plan has the Formalizer prepare artifacts and then lets Go "validate," which shows consistency, not correctness.
- It also skips the §9 capability audit and the specification step. The infra lacks torsion (phys-math: "torsion is never assumed"), density operators and spatial integration.
- **Fix:** add a spec step with a capability audit. Goldens and anti-goldens come from you or an agent other than the implementer, committed before the implementation.

**6. Where:** §5 Adversary. **What:** correlated error; the role is overloaded.
- Nothing requires the Adversary to use a different model or context. The Researcher also extracts, files issues and proposes hypotheses.
- **Fix:** a different `model` in the adversary's frontmatter, a fresh-context subagent, no write access to hypotheses, and reviews that are append-only.

**7. Where:** §6 "remove hash-named files." **What:** append-only is not enforced.
- Git keeps history but doesn't stop a later commit from editing a finding. Deleting a `?` line makes `status` report the point resolved. The plan treats `?` as legitimate but not its removal.
- **Fix:** review files are new-file-only. `check` flags any edited review or removed `?` relative to base unless it cites an evidence ref.

**8. Where:** §4 C-001 example and the "first implementation" demo. **What:** unsupported claims and an infeasible target.
- "Actual — Independent derivation: known divergence problem" rests on no derivation I can see. I only confirmed the title of the 2024 Diósi paper.
- "Expected — Experiment" states a goal, not a prediction under stated assumptions.
- No field carries assumptions, regime or source per line, which is exactly what the Aziz–Howl dispute turned on.
- C-001 needs Riemann–Cartan plus relativistic classical–quantum dynamics plus quantum spin, none of which the infra supports.
- The "non-Markovian" next question is introduced without a source.
- **Fix:** each Expected/Actual line carries `{claim, assumptions, regime, source locator, established_by: source|derivation|experiment}`. Demo with Candidate A on a finite-dimensional toy instead.

**9. Where:** §2 sources, §5 Researcher. **What:** unverified extraction and untrusted input.
- Extracted claims have no locators, so they can't be checked. EBP Sprint 0 intake found substantive gaps in three of five sources.
- Papers are untrusted text read by an agent with tool access, so embedded instructions are a risk.
- **Fix:** every claim has a page/section/equation locator. The extractor can write only under `research/sources/` and run only `physctl`, and never holds merge capability.

**10. Where:** §3 "no `next`," §8. **What:** unbounded autonomy.
- The orchestrating LLM chooses the work and when to stop, with no priority field or per-run bounds.
- **Fix:** a `priority` field you set, sorted by `status --open --json`. One issue and one transition per run.

## Cosmetic

- I can't see the claude/deepseek/gemini/qwen files, so the attributions and the "filepath shadows the package" bug in Qwen's validator are unverifiable from the paste.
- "Five commands" omits the ID allocator from finding 4.

## What holds up

- The plan correctly rejects Qwen's four-keyword restriction and its ban on prose. That is consistent with PSEUDO.md, whose seven constructs are complete and whose attribution lives in prose (§8).
- Keeping `?` legitimate matches §6.
- Separating theory-expected from experiment-expected is right.
- Making SQLite a disposable cache is right.

## Evidential boundary

- **Checked against the repos:** the `ComputeRef` signature, the no-linter statement, GUIDELINES §9 and §10, and the phys-math torsion comment.
- **Not checked:**
  - OpenCode permission semantics for your build, and whether `physctl check` can run fixtures from a separate module (workspace or `replace` behavior).
  - The other proposals' texts.
  - The physics in the C-001 text beyond titles.
  - I ran no Go code.
- **Blocking before a first autonomous run:** findings 1, 4 and 5. Findings 2, 3, 6 and 7 should follow immediately. Findings 8–10 can follow the first dry run.



## gemini

An adversarial review of the proposed Git-native, database-free architecture (`plan.md`) reveals a design that is brilliantly disciplined in its minimalism, yet harbors critical structural vulnerabilities under agentic load.

Stripping away enterprise bloat is the right instinct, but eliminating state management entirely introduces failure modes that will choke autonomous research loops.

---

## 1. The Adversarial Attack Vectors

### Vector 1: The Git-as-Database Scalability & Concurrency Trap

* **The Claim:** No database initially; Git and flat files serve as the single source of truth.


* **The Vulnerability:** OpenCode agents operate asynchronously or in rapid sequence. When multiple agents (Researcher, Adversary, Formalizer) update issues, hypotheses, and references simultaneously, relying on Git flat files (`research/issues/`, `research/hypotheses/`) guarantees **merge conflict storms** and file-locking bottlenecks.
* **The Breaking Point:** Scanning flat markdown files on every invocation of `physctl status` or `physctl refs` to resolve graph relationships ($N$ files, each requiring parsing frontmatter or internal links) will cause CLI latency to explode as the repository grows past a few dozen papers. The plan acknowledges SQLite can be added later as a search index, but waiting for performance to degrade before introducing it is an avoidable trap.



### Vector 2: The Semantic Illusion of `physctl lint`

* **The Claim:** `physctl lint` validates Phys-Pseudo structure and artifact references.


* **The Vulnerability:** Without a robust Abstract Syntax Tree (AST) parser in Go, a lightweight regex-based linter will be easily fooled by autonomous agents. Agents are notorious for generating syntactically *plausible* text that violates deep semantic constraints of physical theories (e.g., index-mismatched tensor contractions or illegal signature transitions). If `physctl lint` only checks structural tags, it passes garbage data into the derivation pipeline.

### Vector 3: The Hallucination Rubber-Stamp (The Human Bottleneck)

* **The Claim:** Human acceptance happens through review and Git integration; agents cannot declare theorems established.


* **The Vulnerability:** The plan correctly identifies that an agent-accessible CLI cannot be trusted with final scientific approval. However, it pushes the entire burden onto **manual human code review of Git diffs**. When the Formalizer agent generates complex mathematical boilerplate via `phys`, `phys-lib`, and `phys-math`, reviewing raw diffs becomes cognitively exhausting. Humans suffering from review fatigue will eventually "rubber-stamp" agent PRs/merges, neutralizing the human-in-the-loop safeguard.



### Vector 4: Diff Noise and State Decay in Flat Files

* **The Claim:** The 4-question expected-versus-actual structure tracks research issues in text files.


* **The Vulnerability:** As an issue evolves through 20 iterations of adversarial challenge, rewriting a flat markdown file leads to messy Git history. The "Actual — Independent derivation" field will be constantly overwritten, making it difficult to trace *when* a specific mathematical obstruction was discovered or *which* adversarial finding broke the previous hypothesis without manually grepping commit logs.

---

## 2. Hardening the Plan (The Patch Set)

To preserve the plan's core virtue—**zero bloat, Git-native canonical record, Go determinism, and OpenCode orchestration**—the following surgical adjustments are required:

| Vulnerability | Hardening Fix |
| --- | --- |
| **Git Concurrency / CLI Latency** | **Embrace the ephemeral SQLite cache from Day 1.** Do not treat SQLite as a competing source of truth; treat it as a *local, read-optimized build artifact* generated automatically by `physctl sync` from Git files. OpenCode queries the local SQLite DB for instant graph traversals; Git remains the immutable ledger.

 |
| **Linter Blindness** | **Enforce structural typing via Go-native AST parsing.** The `physctl lint` command must parse Phys-Pseudo blocks into strict Go structs, validating type signatures and variable scopes before allowing a hypothesis file to be marked as "ready for formalization." |
| **Human Review Fatigue** | **Automate the adversarial audit report.** Before a human ever looks at a Git diff, `physctl check` must output a machine-readable verification report proving that all symbolic constraints in `phys-lib` and `phys-math` closed without algebraic contradiction. The human reviews the *proof summary*, not the raw boilerplate. |

---

## 3. Final Verdict on the Plan

The plan is **viable, but fragile in its execution details.**

By dropping SQLite entirely at the start, you trade infrastructural simplicity for runtime performance agony as soon as your agent ecosystem starts generating interconnected research nodes.

**The Winning Adjustment:** Keep Git as the *storage format* (source of truth), but let Go auto-generate an embedded, local SQLite search/graph index on startup (`.phys.db`, ignored by Git). This gives you the speed of a graph database, the auditability of Git, and zero external infrastructure bloat.



## deepseek

# ADVERSARIAL REVIEW: The Simplified Plan
## Verdict: 7/10. Adopt with six fixes. Three objections are blocking.

---

## Part 1: What the Plan Gets Right

| Claim | Verdict | Justification |
|---|---|---|
| Git-native storage, no DB initially | ✅ Correct | Files + Git = inspectable, diffable, diffable by humans and agents alike. DB can be a rebuildable index later. |
| One Go binary, five commands | ✅ Correct | The minimum viable surface. No daemons, no scheduler. |
| Three agents (Researcher, Adversary, Formalizer) | ✅ Correct | Merging SCOUT+HYPOTHESIZER is right; keeping Adversary is right; renaming DERIVER→Formalizer is right if the Formalizer transcribes rather than derives. |
| "Validation ≠ correctness" | ✅ Correct and important | The plan states this, but (see O-8) does not enforce it. |
| Rejecting invented Phys-Pseudo syntax | ✅ Correct | The four-keyword restriction in the prior proposal was wrong. The seven constructs stay. |
| Rejecting hash-named files when PhysArtifact owns identity | ✅ Correct | Two identity systems is worse than one. |
| The `?` mark is not invalid | ✅ Correct | Unresolved obligations are the *content* of research. The linter must distinguish `MALFORMED` from `OBLIGATION`. |
| Expected/Actual split between theory and experiment | ✅ Correct | Conflating them is the single most common error in the survey's alternatives. |
| "Automate clerical work, make reasoning inspectable, keep proof rigorous" | ✅ Correct principle | But the plan violates its own second clause (see O-16). |

---

## Part 2: Blocking Objections (must fix before first commit)

### **O-1 — The human-authority boundary is a fiction** 🔴 BLOCKING

**What the plan says:**
> "A human-only command in a Go executable is not genuinely human-only if an agent has unrestricted access to that executable and the filesystem. A real acceptance boundary needs a human-controlled Git merge or permissions boundary."

**What the plan then says:**
> "No `approve` command is needed initially: human acceptance happens through review and Git integration."
> "Mandatory PR for every change — **Defer**."

**The contradiction:** The plan correctly identifies that a CLI flag is not a permission boundary, then removes the only mechanism (PR/signed merge) that *is* a permission boundary, and defers it.

**Why it matters:** If OpenCode has shell access (it does), it can run `physctl seal`, `git commit`, and `git push` — all in one session. The "human accepts through Git" step is not enforced. The first dishonest theorem will be indistinguishable from the first honest one.

**Concrete fix (pick one):**
| Option | Cost | Enforces authority? |
|---|---|---|
| **A. GPG-signed merge commits** — `physctl seal` requires a commit signed by a key OpenCode does not have | Low | ✅ Real |
| **B. Separate user** — human commits as `alice`, OpenCode as `agent`; CI rejects `agent` commits that touch `derivations/theorems/` | Medium | ✅ Real |
| **C. `physctl seal` writes to a `SEALS/` directory that is `.gitignore`-d for agents and only committed by human via separate `git` invocation | Low | ⚠ Weak — agent can edit any file |

**Recommendation:** Adopt **A** (GPG-signed merges). Cost: one-time key setup. Benefit: every theorem has a cryptographic provenance chain. This is non-negotiable if the plan is serious about "scientific acceptance belongs to humans."

---

### **O-2 — Removing the falsifier BLOCKER is a regression** 🔴 BLOCKING

**What the previous plan had:**
> "A `⚠` with no `falsifier_id` is a **BLOCKER** — it cannot be promoted."

**What the current plan has:**
> (removed entirely; no mention of falsifier enforcement)

**Why it matters:** A prediction that cannot be falsified is not a prediction — it is prose. The single discipline that separates a research program from a publishing program is that every claim carries its own death condition. Without this gate:
- The Adversary has no checklist to apply to hypotheses.
- `physctl check` has no schema to validate.
- The "expected vs actual" table becomes a wishlist, not a test.

**Concrete fix:** Restore the gate, but scope it correctly:
| Field | Required? | Where |
|---|---|---|
| `prediction.falsifier` | ✅ For every prediction that claims to be testable | `research/evidence/` |
| `hypothesis.falsifiers[]` | ✅ For every hypothesis (may be multiple) | `research/hypotheses/` |
| `theorem.status=IDENTIFIED` | Requires all hypotheses in the chain to have a falsifier | `physctl seal` |

**Add to `physctl lint`:** "A hypothesis with `status ∈ {ACCEPTED, REFUTED}` and zero falsifiers → exit 1."

---

### **O-3 — The provenance chain is assumed, not built** 🔴 BLOCKING

**What the plan says:**
> "`physctl check` — Run available mathematical replay and derivation tests."

**The hidden assumption:** "available" means "the semantic graph of dependencies exists." But with flat files in Git, the only graph is *textual* (mentions, hashes, references). There is no semantic edge from "step 3 of derivation D uses rule R with inputs I and outputs O."

**Why it matters:** Without the semantic graph, `check` cannot replay. Without replay, there is no mathematical gate. Without a gate, `seal` is a rubber stamp. The plan's three-command chain (`lint` → `check` → `seal`) collapses to `lint` → `seal`.

**Concrete fix (three levels, pick by scale):**
| Level | Mechanism | Cost |
|---|---|---|
| **L1** (start here) | Pseudo files carry explicit `depends_on: [C-001, C-002]` headers. `physctl refs` parses these. `physctl check` topologically sorts and replays. | Low |
| **L2** | A `deps.tsv` adjacency file committed alongside each derivation (e.g., `derivations/C-001/deps.tsv`). | Low |
| **L3** | The SQLite index (deferred by the plan) built from L1/L2 headers. | Medium |

**Recommendation:** Start with **L1**. Add L2 if `refs` gets slow. Defer L3 until 1,000 hypotheses. But the *schema* for L1 must be declared in the first commit — otherwise retrofitting is a migration.

---

## Part 3: Serious Objections (should fix before Branch C derivation begins)

### **O-4 — `physctl refs --depth 3` is grep with a namespace** 🟠 SERIOUS

**The claim:** `physctl refs C-001 --depth 3 --json` traverses relationships.

**The problem:** With flat files, "traversal" is text matching. At depth 3 with 500 files, this is O(n²) reads. The plan does not declare how depth is computed:
- Regex over file contents? (fragile; misses multi-line refs)
- Parsed `depends_on` headers? (requires L1 from O-3)
- Git log? (wrong — that's temporal, not semantic)

**Fix:** Either (a) adopt L1 headers and make `refs` a parser, or (b) rename the command to `physctl grep` and admit it. Option (a) is correct; option (b) is honest. The current naming is misleading.

---

### **O-5 — Concurrency is unaddressed** 🟠 SERIOUS

**The scenario:** OpenCode runs three agents in three sessions. Researcher edits `research/issues/C-001.md`; Adversary edits `research/reviews/C-001-review.md`; Formalizer edits `derivations/C-001/step-03.pseudo`. No conflict — different files.

**But:** Adversary reads `research/hypotheses/C-001.pseudo` while Researcher writes it. Race condition. Adversary reviews a stale version. The review is recorded against a SHA that no longer matches.

**Why it matters:** The `replay_status=PASS` gate depends on the reviewed artifact matching the committed artifact. A race breaks this.

**Concrete fix:**
- Every write via `physctl` computes SHA *before* and *after* write, and aborts if the file changed between read and write.
- Every review file records the SHA of the hypothesis it reviewed.
- `physctl check` refuses to promote a hypothesis whose current SHA ≠ reviewed SHA.

This is a 30-line addition to the Go binary. It is not optional if agents run concurrently.

---

### **O-6 — The "pen and paper" handoff is unobservable** 🟠 SERIOUS

**What the plan says:**
> "6. You derive the result using pen and paper.
> 7. Formalizer prepares the executable Phys research artifact."

**The gap:** Between 6 and 7, the derivation exists only in the human's head and on paper. The Formalizer cannot transcribe what it cannot see. The plan does not specify the intermediate artifact.

**Why it matters:** The plan's own principle is "make the reasoning inspectable." A pen-and-paper derivation is the *least* inspectable form. If the Formalizer improvises the steps, the theorem's provenance is fake. If the human photographs the page and the Formalizer OCRs it, the OCR is a new source of error.

**Concrete fix:**
- **Step 6.5:** Human writes a `derivation-sketch.pseudo` (or `.tex`) with the step sequence. This is the inspectable artifact.
- **Step 7:** Formalizer maps each sketch step to a `phys-math` rule and emits `derivation.pseudo` with explicit `rule:` annotations.
- **Step 8:** `physctl check` replays the annotated version.

The sketch is ~20 lines per derivation. This is the minimum for the plan to be internally consistent.

---

### **O-7 — Branch tracking (Q vs C) is implicit** 🟠 SERIOUS

**The problem:** The entire research program hinges on the Q/C fork. The plan has no field, directory, or tag that says "this hypothesis is Branch C."

**Consequences:**
- `physctl status --open` cannot report "which branch is blocking."
- The Adversary cannot apply branch-specific criteria (Branch C has CPTP + constraint closure + torsion checks; Branch Q has quantization + anomaly checks).
- The MVP ("one complete research issue") is ambiguous about which branch it demonstrates.

**Concrete fix:** Add a `branch:` field to every hypothesis, issue, and theorem:
```pseudo
hypothesis cptp_hybrid:
  branch: C
  premises: ...
```
`physctl status --branch C` becomes possible. The Adversary's checklist becomes branch-dependent. This is a 5-character addition to the schema and a 10-line addition to the CLI.

---

## Part 4: Minor Objections (fix when convenient)

### **O-8 — `lint` is elevated above `check`** 🟡 MINOR

The plan lists `lint` first and describes it as "validate structure." But `lint` catches syntax; `check` catches mathematics. In a research program, syntax is cheap and mathematics is expensive. The gate hierarchy should be:

```
check (replay) > adversary review > lint (syntax)
```

**Fix:** Document the hierarchy. Or rename `lint` → `syntax` to signal it's cosmetic.

---

### **O-9 — "No JSON for every artifact" is too broad** 🟡 MINOR

The plan correctly rejects JSON-wrapping pseudo files. But agents need *some* machine-readable structure to process a hypothesis: premises, consequence, obligations, branch, falsifiers. Either:
- (a) Parse pseudo on every read (slow, fragile), or
- (b) Emit a YAML frontmatter block from `physctl lint --json` (fast, stable).

**Fix:** Adopt (b). The pseudo file is the source of truth; the YAML frontmatter is a derived index, regenerated by `lint`. This is not "JSON for every artifact" — it's "one derived index per artifact."

---

### **O-10 — The `?` mark needs two states** 🟡 MINOR

The plan says:
> "The linter should detect malformed syntax, not force every hypothesis to pretend it is complete."

Correct — but then `lint` must distinguish:
- `?` as **syntactic error** (e.g., `premise ?` with no name) → exit 1
- `?` as **obligation** (e.g., `derive X from Y ? whether Y implies X`) → exit 0, log to `obligations`

The plan doesn't specify the two states. Fix: add `OBLIGATION` as an explicit token.

---

### **O-11 — The 3-agent split is asserted, not argued** 🟡 MINOR

The plan says "three roles cover the initial work" — but doesn't say why three, not two or four. The previous proposal had four (SCOUT, HYPOTHESIZER, ADVERSARY, DERIVER). Merging SCOUT+HYPOTHESIZER into "Researcher" is defensible. Merging DERIVER into "Formalizer" is defensible *only if* the Formalizer is a transcriber, not a deriver. This should be stated.

**Fix:** One sentence: "The Formalizer transcribes human-authored derivation sketches into phys-math-annotated artifacts; it does not derive."

---

### **O-12 — "One complete research issue" is not scoped** 🟡 MINOR

The MVP is "from paper ingestion through adversarial review to a human-approved derivation." That is the entire loop. It is not scoped.

**Fix:** The MVP is: **one paper → one issue → one hypothesis with one falsifier → one adversary finding → one human-signed commit.** Derivation is the *second* MVP, not the first. This is achievable in a week.

---

### **O-13 — No migration path for the SQLite index** 🟡 MINOR

The plan says SQLite can be added later as a "disposable, rebuildable search index." But it does not declare the index schema upfront. When you add it, you will either (a) rebuild from scratch (fine) or (b) discover that the file format doesn't support the queries you need (painful).

**Fix:** Declare the future index schema in the first commit, as a comment in `cmd/physctl/schema.sql.txt`. It's free now; it's expensive later.

---

### **O-14 — The `AGENTS.md` vs `.opencode/agents/` distinction is unexplained** 🟡 MINOR

The plan says "use the existing `AGENTS.md` conventions." But also has `.opencode/agents/{researcher,adversary,formalizer}.md`. What is the relationship? Is `AGENTS.md` the constitution, and the agent files are roles? Or do they overlap?

**Fix:** One sentence: "`AGENTS.md` is the constitution (rules all agents share); `.opencode/agents/*.md` are role-specific instructions."

---

## Part 5: The Complete Ledger

| # | Objection | Severity | Fix cost |
|---|---|---|---|
| O-1 | Human authority is not enforced | 🔴 BLOCKING | GPG-signed merges (1 day) |
| O-2 | Falsifier BLOCKER removed | 🔴 BLOCKING | Restore gate (half day) |
| O-3 | Provenance chain assumed | 🔴 BLOCKING | L1 `depends_on` headers (1 day) |
| O-4 | `refs` is grep | 🟠 SERIOUS | Adopt L1 headers (see O-3) |
| O-5 | Concurrency unaddressed | 🟠 SERIOUS | SHA-check-before-write (1 day) |
| O-6 | Pen-and-paper handoff unobservable | 🟠 SERIOUS | Add sketch step (half day) |
| O-7 | Branch Q/C tracking implicit | 🟠 SERIOUS | Add `branch:` field (half day) |
| O-8 | `lint` elevated above `check` | 🟡 MINOR | Rename or document |
| O-9 | "No JSON" too broad | 🟡 MINOR | Derived YAML frontmatter |
| O-10 | `?` needs two states | 🟡 MINOR | Add `OBLIGATION` token |
| O-11 | 3-agent split asserted, not argued | 🟡 MINOR | One sentence |
| O-12 | MVP not scoped | 🟡 MINOR | Restrict to one-paper-one-issue |
| O-13 | SQLite migration undeclared | 🟡 MINOR | Comment the schema |
| O-14 | `AGENTS.md` vs roles unexplained | 🟡 MINOR | One sentence |

**Total fix time:** ~4 days. Total deferral time: unbounded (the plan will discover these problems one by one, in production, at 10× the cost).

---

## Part 6: The Verdict

**The plan is the best proposal so far — it correctly removes the database, the JSON-for-everything, the custom orchestrator, the automatic approval. It is 7/10.**

**The three blocking objections are not minor.** They are the difference between:
- A research program that *says* it enforces epistemic discipline, and
- A research program that *does*.

The plan itself admits the human-authority problem ("a human-only command in a Go executable is not genuinely human-only") and then removes the mechanism (PR/merge) that would fix it. That is the single most important thing to fix. Without GPG-signed merges (or a real permissions boundary), the entire "human acceptance" step is theater.

The second blocking issue is the falsifier gate. The previous plan had it; this plan dropped it. Dropping it is a regression, not a simplification. A prediction without a death condition is not a scientific claim.

The third blocking issue is the provenance chain. The plan assumes `physctl check` can replay, but does not specify how the semantic graph is built from flat files. Without it, `check` is a no-op and `seal` is a rubber stamp.

**The honest summary:** The plan is the right architecture with three holes punched in the hull. Patch them before the first commit, not after.

---

## Part 7: The One-Paragraph Summary

> **The plan is the best proposal so far — 7/10. It correctly removes the database, the JSON envelopes, the custom orchestrator, and the automatic approval, and it correctly preserves the seven Phys-Pseudo constructs, the `?` mark as legitimate content, and the three-agent split. But it has three blocking objections: (O-1) the human-authority boundary is a fiction — the plan admits it and then removes the only fix (signed merges); (O-2) the falsifier BLOCKER is removed, which is a regression, not a simplification; (O-3) the provenance chain is assumed but not built, so `check` and `seal` are no-ops. It has three serious objections: (O-4) `refs` is grep; (O-5) concurrency is unaddressed; (O-6) the pen-and-paper handoff is unobservable; (O-7) Branch Q/C tracking is implicit. And it has seven minor objections: `lint` is elevated above `check`, "no JSON" is too broad, `?` needs two states, the 3-agent split is asserted not argued, the MVP is not scoped, the SQLite migration is undeclared, and the `AGENTS.md`/roles distinction is unexplained. The total fix time is ~4 days. The total deferral time is unbounded. Fix the three blockers before the first commit; fix the four serious issues before the first derivation; fix the seven minors when convenient. The plan is the right architecture. It is not yet a disciplined one.**


## qwen

As the **Adversary Agent** defined in your plan, my role is to challenge assumptions, check for missing derivations, and search for structural counterexamples. 

I have reviewed `plan.md`. The architectural synthesis is excellent—it successfully strips away the bloat of databases and microservices. However, under adversarial stress-testing, the plan reveals **five critical structural weaknesses** regarding scaling, concurrency, and epistemic boundaries. 

Here is the adversarial review.

---

### Finding 1: The `physctl refs` $O(N)$ Trap
**The Claim:** Git and files are the single source of truth. `physctl refs` follows relationships between papers, problems, and hypotheses.
**The Adversarial Challenge:** Git is a version control system, not a graph database. To execute `physctl refs C-001 --depth 3`, the Go binary must recursively scan the `research/` and `derivations/` directories, parse every Markdown/JSON file, extract cross-references, build an in-memory adjacency list, and traverse it. 
- **The Failure Mode:** This works flawlessly for 100 files. It will cause unacceptable latency and memory spikes at 10,000 files. 
- **The Contradiction:** The plan states, *"SQLite can be added later... if repository scanning becomes measurably slow."* But if `refs` and `status` are core daily commands, the system will hit this wall much sooner than anticipated, forcing an unplanned architectural pivot.

### Finding 2: The Orchestration Vacuum
**The Claim:** "OpenCode coordinates." "No custom message bus."
**The Adversarial Challenge:** OpenCode is a coding agent (an IDE/CLI tool), not a persistent background daemon or event loop. It does not "coordinate" asynchronously. 
- **The Failure Mode:** If the Researcher agent creates an issue, the Adversary agent does not automatically wake up to challenge it. OpenCode requires explicit human prompting or a wrapper script to invoke the next agent. 
- **The Reality:** The plan inadvertently relies on the **human** as the message bus. If the human forgets to invoke the Adversary, the workflow stalls. The plan must explicitly define how state transitions are triggered without a daemon.

### Finding 3: The Markdown Merge Conflict Hazard
**The Claim:** Git handles the workflow. Agents create and update files.
**The Adversarial Challenge:** Git is notoriously hostile to concurrent writes on the same file. 
- **The Failure Mode:** If the Researcher updates `issues/C-001.md` with new claims, and simultaneously the Adversary updates the same file with objections, Git will produce a merge conflict. LLMs (and humans) are exceptionally bad at resolving semantic merge conflicts in structured Markdown.
- **The Missing Derivation:** The plan lacks a concurrency control mechanism. How do agents safely mutate shared state (the issue file) without corrupting the Git history?

### Finding 4: The Epistemic Limit of `physctl lint`
**The Claim:** `physctl lint` validates Phys-Pseudo structure. Natural language is allowed (rejecting Qwen's strict no-prose rule).
**The Adversarial Challenge:** If natural language is permitted for source attribution and explanation, `physctl lint` cannot verify scientific correctness; it can only verify structural integrity (e.g., "Does the referenced file `C-002.pseudo` actually exist on disk?").
- **The Failure Mode:** An agent could write a beautifully formatted, structurally valid Phys-Pseudo file that contains complete physical nonsense or hallucinated math. `physctl lint` will return `exit 0` (success), giving a false sense of validation.
- **The Missing Boundary:** The plan correctly notes that passing a Go validator does not establish scientific truth, but it fails to define what `lint` *actually* guarantees. 

### Finding 5: The "Actual" Data Void
**The Claim:** Every issue tracks "Expected vs Actual" for both Theory and Experiment.
**The Adversarial Challenge:** The plan thoroughly defines how to generate "Expected" (Theory) via Phys-Pseudo and derivations. It completely omits how "Actual" (Experiment) is ingested.
- **The Failure Mode:** Where does the empirical data come from? If an agent is supposed to compare a theory against an experiment, how does it ingest the experimental dataset? Is it a CSV? A DOI link? A hardcoded number in the Markdown? Without a defined schema for empirical evidence, the "Actual — Experiment" field will become a dumping ground for unstructured, unverified prose.

---

### Required Corrections (Adversary's Mandates)

To resolve these findings without reintroducing the bloat we just deleted, I mandate the following corrections to the plan:

#### Correction 1: Restrict `physctl refs` to Local Neighborhoods
Do not attempt to build a global graph in memory. `physctl refs` must only parse files that explicitly contain a backlink or forward-link to the target ID. 
- *Implementation:* Require a strict YAML frontmatter or a specific Markdown link syntax (e.g., `[C-001](../issues/C-001.md)`). `physctl` uses `grep` (via Go's `os.ReadDir` and string matching) to find only the files that reference the target, keeping it $O(K)$ where $K$ is the number of referencing files, not $O(N)$.

#### Correction 2: Explicit Human-in-the-Loop Orchestration
Acknowledge that OpenCode is a tool, not an orchestrator. Update the workflow to explicitly state:
- *The human invokes the Researcher.*
- *The human reviews the output, then invokes the Adversary.*
- *The human reviews the dispute, then invokes the Formalizer.*
Remove the phrase "OpenCode coordinates" and replace it with "Human drives OpenCode through sequential agent roles."

#### Correction 3: The "Append-Only" Concurrency Rule
To prevent Git merge conflicts, agents must **never overwrite** an existing issue file. 
- *Implementation:* If the Adversary has a finding, it does not edit `issues/C-001.md`. It creates a new file: `reviews/C-001-adv-001.md`. The issue file remains immutable once created. `physctl status` simply aggregates the issue and its linked reviews. This guarantees zero merge conflicts.

#### Correction 4: Redefine `physctl lint` as "Structural Integrity Only"
Update the plan to explicitly state: `physctl lint` verifies **reference integrity** (e.g., "All linked artifacts exist on disk") and **schema compliance** (e.g., "Required frontmatter is present"). It explicitly **does not** verify physical or mathematical truth. Truth is established by `physctl check` (math replay) and human acceptance.

#### Correction 5: Define the "Evidence" Schema
The `evidence/` directory must have a strict, minimal schema for experimental data. 
- *Implementation:* An evidence file must contain: `Source DOI/URL`, `Date`, `Measured Value`, `Uncertainty/Margin of Error`, and `Methodology Summary`. If an agent cannot fill out these fields, it cannot claim an "Actual — Experiment" result.

### Final Adversary Verdict
The plan is 90% there. By enforcing **append-only reviews**, **local-only graph traversal**, and **strict evidence schemas**, you eliminate the hidden scaling and concurrency traps while preserving the beautiful simplicity of Git + Go + OpenCode. 

Shall I proceed to draft the corrected `AGENTS.md` and the Go struct definitions for the append-only review model?





