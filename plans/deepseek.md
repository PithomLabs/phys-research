# OpenCode as Orchestrator
## The agents are prompts. The Go code is the tool. OpenCode drives the loop.

---

## 1. The Reframe

Your previous design had four Go binaries that each ran a `for { select }` loop. **That was wrong for your setup.** You have OpenCode — a coding agent that can read files, run commands, edit code, and use tools. OpenCode **is** the orchestrator. The Go code is the **deterministic tool layer** it calls.

| Layer | Who does it | What it is |
|---|---|---|
| **Judgment** | OpenCode (LLM) | Read paper, decide it's CONTESTED, draft hypothesis, run MRC review |
| **Determinism** | Go binary (`physctl`) | Parse pseudo, validate schema, write DB, replay derivations, compute SHA |
| **State** | SQLite | Single file, six tables |
| **Coordination** | OpenCode's tool loop | No custom orchestrator, no message bus |

**The Go binary is not an agent.** It is a **CLI tool** with verbs. OpenCode calls `physctl scout`, `physctl hypothesize`, `physctl review`, `physctl derive` — and reads the JSON output.

---

## 2. The Tool Surface (one binary, six verbs)

```
physctl
├── ingest   --arxiv 2510.12345 --status CONTESTED --objection-to 42
├── problem  --add "Runaway solutions in semiclassical backreaction" \
│            --mrc Calculus --severity CRITICAL
├── hypoth   --from-problem 7 --pseudo ./branch-c/cptp-001.pseudo
├── review   --hypothesis 3 --verdict PASS --gaps "IVP not proven; torsion algebraic"
├── derive   --hypothesis 3 --step 2 --rule DIFFERENTIATE-001 \
│            --input sha:abc --output sha:def
├── seal     --theorem 5 --session S-2026-11-04-001   # human-only
└── status   --branch C --format json                 # agent-readable snapshot
```

**Every verb:**
- Reads/writes SQLite (single file, `modernc.org/sqlite`)
- Emits JSON to stdout (OpenCode parses)
- Exits non-zero on gate failure (e.g., `seal` without a valid session)

**The `seal` verb is the only one gated by a human session ID.** OpenCode cannot call it with a synthetic session. That is the epistemic debt discharge, enforced in Go, not in a prompt.

---

## 3. The Prompt Surface (OpenCode skills)

OpenCode reads skill files. Four skills, one per role. Each skill is a **thin prompt** that tells OpenCode:

1. What verb to call
2. What judgment to make
3. What to write back

```
.opencode/
  skills/
    scout.md          # "Ingest new arXiv papers. Call physctl ingest. 
                       #  If contested, set objection_to. Never invent status."
    hypothesizer.md   # "Read problems with status=OPEN. Draft phys-pseudo.
                       #  Every premise needs provenance. Use ? for gaps."
    adversary.md      # "Read hypotheses with mrc_verdict IS NULL. 
                       #  Run MRC_REVIEW_PROTOCOL.md. Call physctl review.
                       #  FAIL → file new problems from the ? items."
    deriver.md        # "Read hypotheses with mrc_verdict=PASS. 
                       #  Map each step to one of nine phys-math rules. 
                       #  Call physctl derive. Never call seal."
```

**The skill is the role.** It is ~200 words, not a 2000-word system prompt. OpenCode loads it when invoked.

---

## 4. The Loop (OpenCode drives it, Go enforces gates)

```
OpenCode session
  │
  ├─ reads scout.md → runs `physctl ingest ...` → reads JSON
  ├─ decides paper is CONTESTED → runs `physctl ingest --objection-to 42`
  ├─ reads hypothesizer.md → drafts PSEUDO.md → runs `physctl hypoth`
  ├─ reads adversary.md → runs MRC_REVIEW_PROTOCOL.md (prompt) 
  │                     → runs `physctl review --verdict PASS`
  ├─ reads deriver.md → maps steps → runs `physctl derive` per step
  └─ STOPS. Cannot call `seal`. Human runs it in a separate session.
```

**The gates are in Go, not in the prompt:**

| Gate | Enforced by |
|---|---|
| `theorem.IDENTIFIED` requires session ID | SQLite `CHECK` + `physctl seal` validates session |
| Hypothesis with `?` gaps cannot be reviewed PASS | `physctl review` rejects if `pseudo` has unresolved `?` |
| Derivation with `replay_status=FAIL` cannot promote | `physctl derive` refuses to write `theorem` |
| Prediction without `falsifier_id` is `BLOCKER` | `physctl status` flags it |

**OpenCode cannot bypass these.** The prompt can say "seal the theorem" — the tool will exit non-zero. This is the discipline. The LLM proposes; the Go code disposes.

---

## 5. Why This Is Simpler Than the Previous Design

| Previous | Now |
|---|---|
| 4 Go binaries, each a `for { select }` loop | 1 Go binary, 6 verbs |
| Custom orchestrator (ARGUS MCP or equivalent) | OpenCode's native tool loop |
| Agents coordinate via DB state | OpenCode coordinates via tool calls + DB state |
| Prompt = 2000-word system prompt | Prompt = 200-word skill file |
| MCP server needed | Just a CLI — OpenCode calls it via bash |

**The Go binary is ~800 lines.** The skills are ~800 words total. The schema is unchanged. The SQLite file is the state. OpenCode is the loop. That is the whole system.

---

## 6. What OpenCode Actually Does (concrete example)

**Session start:**
```bash
$ physctl status --branch C --format json
{"problems": [{"id": 7, "title": "Runaway solutions", "status": "OPEN"}], 
 "hypotheses": [{"id": 3, "mrc_verdict": "PASS", "derivations": 2}],
 "blockers": [{"id": 12, "reason": "prediction without falsifier"}]}
```

**OpenCode reads `deriver.md`, sees hypothesis 3 is PASS, opens the pseudo file:**

```pseudo
hypothesis cptp_hybrid:
  premise ...
  consequence cptp: hybrid_dynamics is CPTP
  derive cptp from [linearity, trace_preservation, complete_positivity]
    ? whether CPTP is sufficient for IVP  # ← this is a gap
```

**OpenCode calls:**
```bash
$ physctl derive --hypothesis 3 --step 1 \
    --rule DIFFERENTIATE-001 --input sha:abc --output sha:def
{"step": 1, "replay_status": "PASS", "next_obligation": "IVP"}
```

**OpenCode sees `next_obligation: IVP`, and the `?` gap. It does NOT call `seal`. It files a problem:**
```bash
$ physctl problem --add "IVP for CPTP hybrid dynamics" \
    --mrc Calculus --severity BLOCKER --hypothesis 3
{"problem": 8, "blocker": true}
```

**The loop continues.** OpenCode never promotes. It only proposes, derives, and files gaps. The human seals.

---

## 7. The One-Sentence Answer

**Use OpenCode as the orchestrator: the Go binary (`physctl`) is a deterministic CLI with six verbs that enforce gates via SQLite `CHECK` constraints, the four agent roles are 200-word skill files in `.opencode/skills/`, and the loop is OpenCode's native tool-call loop — no custom orchestrator, no MCP server, no message bus, just a tool the LLM calls and a database it cannot lie to.**
