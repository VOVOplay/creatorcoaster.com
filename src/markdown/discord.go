package markdown

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
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
//
// The only errors returned are errors occurred while looking up a variable
// in the database.
func GenerateDiscordMarkdownFromString(s string) (string, error) {
	var b strings.Builder
	b.Grow(len(s))

	var variableCache map[variableCacheKey]string

	for line := range strings.Lines(s) {
		if isHorizontalRule(line) {
			continue
		}

		last := 0

		for i := 0; i < len(line); i++ {
			if line[i] != '{' {
				continue
			}

			nameStart := i + 1
			var kind variableKind

			if nameStart < len(line) {
				switch line[nameStart] {
				case '#':
					kind = channelVar
					nameStart++
				case '@':
					kind = roleUserVar
					nameStart++
				}
			}

			if nameStart >= len(line) || !isVariableNameStart(line[nameStart]) {
				continue
			}

			end := nameStart
			for end < len(line) && isVariableNameChar(line[end]) {
				end++
			}

			if end >= len(line) || line[end] != '}' {
				continue
			}

			b.WriteString(line[last:i])

			name := line[nameStart:end]
			key := variableCacheKey{kind: kind, name: name}

			value, found := variableCache[key]

			if !found {
				var err error
				value, err = expandVariable(kind, name)
				if err != nil {
					return "", err
				}

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

	return b.String(), nil
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

func expandVariable(kind variableKind, name string) (string, error) {
	switch kind {
	case standardVar:
		value, err := database.GetVariableByName(name)
		if err != nil {
			return "", fmt.Errorf("get variable %q: %w", name, err)
		}
		return value, nil

	case channelVar:
		channel, err := database.GetChannelVariableByName(name)
		if err != nil {
			return "", fmt.Errorf("get channel variable %q: %w", name, err)
		}
		s := strconv.FormatUint(channel.ID, 10)

		var b strings.Builder
		b.Grow(len(s) + 3)
		b.WriteString("<#")
		b.WriteString(s)
		b.WriteByte('>')
		return b.String(), nil

	case roleUserVar:
		// TODO: Update to use <@UserID> and <@&RoleID> syntax rather
		// than the "pretty name" once the database returns a
		// myTypes.RoleUserVariable instead of string.

		prettyName, err := database.GetRoleUserVariableByName(name)
		if err != nil {
			return "", fmt.Errorf("get role/user variable %q: %w", name, err)
		}
		return prettyName, nil

	default:
		return "", fmt.Errorf("unknown variable kind %d", kind)
	}
}

func isVariableNameStart(c byte) bool {
	return c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isVariableNameChar(c byte) bool {
	return isVariableNameStart(c) || (c >= '0' && c <= '9')
}
