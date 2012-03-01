package basedir

import "testing"

func TestMakeMap(t *testing.T) {
	kvForms := []string{"name=peter", "age=42"}
	mapping := makeMap(kvForms)
	expectedValue := make(environ)
	expectedValue["name"] = "peter"
	expectedValue["age"] = "42"
	if !mapping.Eq(expectedValue) {
		t.Errorf("expected %v to be equal to %v", mapping, expectedValue)
	}
}

func TestGetEnvAccessible(t *testing.T) {
	env := make(environ)
	env["PAGER"] = "LESS"
	pager := getEnv(env, "PAGER", "XY")
	if pager != "LESS" {
		t.Errorf("%v ≠ \"LESS\"", pager)
	}
}

func TestGetEnvWithDefault(t *testing.T) {
	env := make(environ)
	pager := getEnv(env, "PAGER", "XY")
	if pager != "XY" {
		t.Errorf("%v ≠ \"XY\"", pager)
	}
}
