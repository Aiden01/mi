package mi

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	hitToken       = regexp.MustCompile("^(\\d+(?:,\\d+)*)(\\.*)(\\+*|-*)$")
	directiveToken = regexp.MustCompile("^([^:]+):(.*)$")
	tokenizer      = regexp.MustCompile("(?m)\\s+")
	comment        = regexp.MustCompile("#[^\n]*")

	drumNotes  = map[string]byte{}
	velocities = map[string]Velocity{
		"-----": PPP,
		"----":  PP,
		"---":   P,
		"--":    MP,
		"-":     MF,
		"":      F,
		"+":     FF,
		"++":    FFF,
	}
	durations = map[string]uint{
		"":       96 * 4,
		".":      96 * 2,
		"..":     96,
		"...":    96 / 2,
		"....":   96 / 4,
		".....":  96 / 8,
		"......": 96 / 16,
	}

	directives = map[string]directive{
		"bpm": bpmDirective,
	}
)

func init() {
	for i := byte(35); i <= 81; i++ {
		drumNotes[fmt.Sprint(i)] = i
	}
}

func ParseTrack(s string) (*Track, error) {
	t := &Track{}
	for i, token := range tokenize(s) {
		switch {
		case hitToken.MatchString(token):
			h, err := parseHit(token)
			if err != nil {
				return nil, fmt.Errorf("token #%v: %v", i, err)
			}
			t.Hits = append(t.Hits, h)
		case directiveToken.MatchString(token):
			if err := t.parseDirective(token); err != nil {
				return nil, fmt.Errorf("token #%v: %v", i, err)
			}
		default:
			return nil, fmt.Errorf("token #%v: unrecognized token: %q", i, token)
		}
	}
	return t, nil
}

func tokenize(s string) []string {
	s = comment.ReplaceAllString(s, "")
	var result []string
	for _, t := range tokenizer.Split(s, -1) {
		if t == "" {
			continue
		}
		result = append(result, t)
	}
	return result
}

func parseHit(s string) (*Hit, error) {
	m := hitToken.FindStringSubmatch(s)
	if m == nil {
		return nil, fmt.Errorf("bad hit: %q", s)
	}

	notes := []byte{}
	for _, n := range strings.Split(m[1], ",") {
		if drumNotes[n] == 0 {
			return nil, fmt.Errorf("bad drum number: %q", n)
		}
		notes = append(notes, drumNotes[n])
	}

	d := durations[m[2]]
	if d == 0 {
		return nil, fmt.Errorf("bad duration: %q", m[2])
	}

	v := velocities[m[3]]
	if v == 0 {
		return nil, fmt.Errorf("bad velocity: %q", m[3])
	}

	return NewHit(d, v, notes...), nil
}

type directive func(*Track, string) error

func (t *Track) parseDirective(s string) error {
	m := directiveToken.FindStringSubmatch(s)
	if m == nil {
		return fmt.Errorf("bad directive: %q", s)
	}
	d := directives[m[1]]
	if d == nil {
		return fmt.Errorf("unknown directive: %q", m[1])
	}
	return d(t, m[2])
}

func bpmDirective(t *Track, s string) error {
	bpm, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("bad input to BPM: %v", err)
	}
	if bpm < 1 || bpm > 500 {
		return fmt.Errorf("bad BPM: %v, must be between 1 and 500")
	}
	t.BPM = uint(bpm)
	return nil
}