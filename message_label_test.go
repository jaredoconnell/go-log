package log_test

import (
	"testing"

	"go.arcalot.io/log/v2"
)

func TestMessageLabels(t *testing.T) {
	var l log.Labels = map[string]string{}
	if l.String() != "" {
		t.Fatalf("Incorrect label string: %s", l.String())
	}
	l = map[string]string{"foo": "bar"}
	if l.String() != "foo=bar" {
		t.Fatalf("Incorrect label string: %s", l.String())
	}
	l = map[string]string{"foo": "bar", "baz": "Hello world!"}
	if l.String() != "foo=bar;baz=Hello world!" && l.String() != "baz=Hello world!;foo=bar" {
		t.Fatalf("Incorrect label string: %s", l.String())
	}
}
