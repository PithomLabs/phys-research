// Package frontmatter parses the deliberately restricted research-record
// metadata header defined in AGENTS.md §6.1.
//
// It is a line-based parser for a fixed, tiny subset: scalar fields from a
// closed key set and a `refs` list of {relation, id} pairs. It is not a
// general YAML parser, it does not guess, and it rejects everything outside
// the documented subset with a line-numbered error. It never touches
// Phys-Pseudo content: the body after the closing delimiter is returned
// as-is (apart from CRLF→LF normalization).
package frontmatter

import (
	"fmt"
	"strings"
)

// SyntaxError reports a frontmatter parse failure with the 1-based line
// number of the offending line in the file.
type SyntaxError struct {
	Line int
	Msg  string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("frontmatter line %d: %s", e.Line, e.Msg)
}

// Ref is one declared relationship in a record header.
type Ref struct {
	Relation string
	ID       string
}

// Header is the parsed metadata header of a research record. Fields that
// were absent remain the empty string; semantic validation of required
// values lives in the validate package, not here.
type Header struct {
	ID       string
	Type     string
	State    string
	Branch   string
	Priority string
	Class    string
	Refs     []Ref
}

// AllowedKeys is the closed set of header fields (AGENTS.md §6.1).
var AllowedKeys = []string{"id", "type", "state", "branch", "priority", "class", "refs"}

var allowedKeySet = func() map[string]bool {
	m := make(map[string]bool, len(AllowedKeys))
	for _, k := range AllowedKeys {
		m[k] = true
	}
	return m
}()

// Parse splits content into header and body and parses the header subset.
// The body is returned untouched apart from line-ending normalization, and
// bodyLine is the 1-based line number in content where the body begins
// (the line after the closing --- delimiter), so callers can report
// file-absolute line numbers.
func Parse(content []byte) (Header, []byte, int, error) {
	var h Header

	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(text, "\n")

	if len(lines) == 0 || lines[0] != "---" {
		return h, nil, 0, &SyntaxError{Line: 1, Msg: "file must start with a --- frontmatter delimiter"}
	}

	closeIdx := -1
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			closeIdx = i
			break
		}
	}
	if closeIdx < 0 {
		return h, nil, 0, &SyntaxError{Line: 1, Msg: "unterminated frontmatter: closing --- delimiter not found"}
	}

	body := strings.Join(lines[closeIdx+1:], "\n")

	seen := map[string]bool{}
	inRefs := false
	var cur *Ref

	finalize := func(line int) error {
		if cur == nil {
			return nil
		}
		if cur.Relation == "" {
			return &SyntaxError{Line: line, Msg: "refs item missing 'relation'"}
		}
		if cur.ID == "" {
			return &SyntaxError{Line: line, Msg: "refs item missing 'id'"}
		}
		h.Refs = append(h.Refs, *cur)
		cur = nil
		return nil
	}

	headerLines := lines[1:closeIdx]
	for i, raw := range headerLines {
		abs := i + 2 // line 1 is the opening ---

		if strings.Contains(raw, "\t") {
			return h, nil, 0, &SyntaxError{Line: abs, Msg: "tab characters are not allowed"}
		}
		line := strings.TrimRight(raw, " ")
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue // full-line comment
		}

		indent := len(line) - len(strings.TrimLeft(line, " "))

		if indent == 0 {
			if err := finalize(abs); err != nil {
				return h, nil, 0, err
			}
			inRefs = false

			key, value, err := splitKV(line, abs)
			if err != nil {
				return h, nil, 0, err
			}
			if !allowedKeySet[key] {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: fmt.Sprintf("unknown field %q (allowed: %s)", key, strings.Join(AllowedKeys, ", "))}
			}
			if seen[key] {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: fmt.Sprintf("duplicate field %q", key)}
			}
			seen[key] = true

			if key == "refs" {
				if value != "" {
					return h, nil, 0, &SyntaxError{Line: abs, Msg: "refs must use the list form ('- relation:' / '    id:' items)"}
				}
				inRefs = true
				continue
			}

			v, err := cleanValue(value, abs)
			if err != nil {
				return h, nil, 0, err
			}
			switch key {
			case "id":
				h.ID = v
			case "type":
				h.Type = v
			case "state":
				h.State = v
			case "branch":
				h.Branch = v
			case "priority":
				h.Priority = v
			case "class":
				h.Class = v
			}
			continue
		}

		// Indented lines are only valid as refs list items.
		trimmed := strings.TrimLeft(line, " ")
		if !inRefs {
			return h, nil, 0, &SyntaxError{Line: abs, Msg: "malformed indentation: indented line outside a 'refs:' list"}
		}

		if strings.HasPrefix(trimmed, "- ") {
			if indent != 2 {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: "malformed indentation: list items must be indented exactly 2 spaces"}
			}
			if err := finalize(abs); err != nil {
				return h, nil, 0, err
			}
			key, value, err := splitKV(strings.TrimLeft(trimmed[2:], " "), abs)
			if err != nil {
				return h, nil, 0, err
			}
			if key != "relation" && key != "id" {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: fmt.Sprintf("unknown refs field %q (allowed: relation, id)", key)}
			}
			v, err := cleanValue(value, abs)
			if err != nil {
				return h, nil, 0, err
			}
			if v == "" {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: fmt.Sprintf("empty %q value", key)}
			}
			cur = &Ref{}
			if key == "relation" {
				cur.Relation = v
			} else {
				cur.ID = v
			}
			continue
		}

		if indent != 4 {
			return h, nil, 0, &SyntaxError{Line: abs, Msg: "malformed indentation: refs item keys must be indented exactly 4 spaces"}
		}
		if cur == nil {
			return h, nil, 0, &SyntaxError{Line: abs, Msg: "refs item key outside of a '- ' list item"}
		}
		key, value, err := splitKV(trimmed, abs)
		if err != nil {
			return h, nil, 0, err
		}
		if key != "relation" && key != "id" {
			return h, nil, 0, &SyntaxError{Line: abs, Msg: fmt.Sprintf("unknown refs field %q (allowed: relation, id)", key)}
		}
		v, err := cleanValue(value, abs)
		if err != nil {
			return h, nil, 0, err
		}
		if v == "" {
			return h, nil, 0, &SyntaxError{Line: abs, Msg: fmt.Sprintf("empty %q value", key)}
		}
		if key == "relation" {
			if cur.Relation != "" {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: "duplicate 'relation' in refs item"}
			}
			cur.Relation = v
		} else {
			if cur.ID != "" {
				return h, nil, 0, &SyntaxError{Line: abs, Msg: "duplicate 'id' in refs item"}
			}
			cur.ID = v
		}
	}

	if err := finalize(len(headerLines) + 1); err != nil {
		return h, nil, 0, err
	}
	// The closing --- sits at 1-based line closeIdx+1, so the body starts
	// on the following line.
	return h, []byte(body), closeIdx + 2, nil
}

// splitKV splits "key: value" (or "key:" at end of line) and rejects YAML
// shapes outside the documented subset.
func splitKV(line string, abs int) (key, value string, err error) {
	idx := strings.Index(line, ":")
	if idx < 0 {
		return "", "", &SyntaxError{Line: abs, Msg: fmt.Sprintf("expected 'key: value', got %q", line)}
	}
	if idx+1 < len(line) && line[idx+1] != ' ' {
		return "", "", &SyntaxError{Line: abs, Msg: "expected a space after ':'"}
	}
	key = line[:idx]
	if key == "" || strings.ContainsAny(key, " \t") {
		return "", "", &SyntaxError{Line: abs, Msg: fmt.Sprintf("invalid field name %q", key)}
	}
	return key, strings.TrimSpace(line[idx+1:]), nil
}

// cleanValue rejects unsupported YAML value shapes and strips one layer
// of matching quotes. It never silently reinterprets a value.
func cleanValue(v string, abs int) (string, error) {
	if v == "" {
		return "", nil
	}
	switch v[0] {
	case '&', '*':
		return "", &SyntaxError{Line: abs, Msg: "anchors and aliases are not supported"}
	case '|', '>':
		return "", &SyntaxError{Line: abs, Msg: "multiline strings are not supported"}
	case '[', '{':
		return "", &SyntaxError{Line: abs, Msg: "flow collections are not supported"}
	case '"', '\'':
		q := v[0]
		if len(v) < 2 || v[len(v)-1] != q {
			return "", &SyntaxError{Line: abs, Msg: "malformed quoted value"}
		}
		return v[1 : len(v)-1], nil
	}
	return v, nil
}
