package markdown

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
)

type variableKind int

const (
	standardVar variableKind = iota
	channelVar
	roleUserVar
)

type variableCacheKey struct {
	kind variableKind
	name string
}

// GenerateDiscordMarkdownFromString removes horizontal rules and expands all
// variables in the input Markdown to produce Discord-compatible Markdown.
//
// Textual transformations (not Markdown parsing or validation) are applied to
// remove any lines that only consist of three or more contiguous hyphens,
// asterisks, or underscores, ignoring surrounding whitespace, and to replace
// variables with their values (formatting is not escaped). It does not
// distinguish between ordinary text, code, escaped characters, etc.
//
// Input Markdown must already be Discord-compatible besides horizontal rules
// (no Setext-style headings, raw HTML, images, indented code blocks, etc).
// Variable names must not contain any whitespace, and are terminated by
// closing braces (and begun with opening braces).
func GenerateDiscordMarkdownFromString(source string) string {
	var b strings.Builder
	b.Grow(len(source))

	var variableCache map[variableCacheKey]string

	for line := range strings.Lines(source) {
		if isHorizontalRule(line) {
			continue
		}

		last := 0

		for i := 0; i < len(line); i++ {
			if line[i] != '{' {
				continue
			}

			kind, name, end, ok := parseVariable(line, i)
			if !ok {
				continue
			}

			b.WriteString(line[last:i])

			key := variableCacheKey{kind: kind, name: name}

			value, found := variableCache[key]
			if !found {
				value = expandVariable(kind, name)

				if variableCache == nil {
					variableCache = make(map[variableCacheKey]string)
				}

				variableCache[key] = value
			}

			b.WriteString(value)

			i = end
			last = end + 1
		}

		b.WriteString(line[last:])
	}

	return b.String()
}

func isHorizontalRule(line string) bool {
	line = strings.TrimSpace(line)

	if len(line) < 3 {
		return false
	}

	marker := line[0]
	if marker != '-' && marker != '*' && marker != '_' {
		return false
	}

	for i := 1; i < len(line); i++ {
		if line[i] != marker {
			return false
		}
	}

	return true
}

func parseVariable(line string, start int) (kind variableKind, name string, end int, ok bool) {
	nameStart := start + 1

	if nameStart < len(line) {
		switch line[nameStart] {
		case '#':
			kind = channelVar
			nameStart = start + 2
		case '@':
			kind = roleUserVar
			nameStart = start + 2
		}
	}

	for offset, rune := range line[nameStart:] {
		if unicode.IsSpace(rune) {
			return 0, "", 0, false
		}

		if rune == '}' {
			end = nameStart + offset
			if end == nameStart {
				return 0, "", 0, false
			}

			return kind, line[nameStart:end], end, true
		}
	}

	return 0, "", 0, false
}

func expandVariable(kind variableKind, name string) string {
	switch kind {
	case standardVar:
		standardVariable, err := database.GetVariableByName(name)
		if err != nil {
			return "{" + name + "}"
		}
		return standardVariable.VariableContents

	case channelVar:
		channelVariable, err := database.GetChannelVariableByName(name)
		if err != nil {
			return "{#" + name + "}"
		}
		return "<#" + strconv.FormatUint(channelVariable.ID, 10) + ">"

	case roleUserVar:
		roleUserVariable, err := database.GetRoleUserVariableByName(name)
		if err != nil {
			return "{@" + name + "}"
		}
		switch roleUserVariable.Type {
		case myTypes.Role:
			return "<@&" + strconv.FormatUint(roleUserVariable.ID, 10) + ">"
		case myTypes.User:
			return "<@" + strconv.FormatUint(roleUserVariable.ID, 10) + ">"
		default:
			// TODO: log this
			return "[ERROR]"
		}

	default:
		// TODO: log this
		return "[ERROR]"
	}
}
