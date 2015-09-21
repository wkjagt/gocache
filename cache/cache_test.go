package cache
import (
  "testing"
)

func TestExtractSetCommand(t *testing.T) {
  payload := "SET a b"
  expectedResult := Command{"SET", []string{"a", "b"}}
  result, err := ExtractCommand(payload)

  if err != nil {
    t.Fatalf("Unexpected error returned")
  }

  if expectedResult.name != result.name {
    t.Fatalf("expected command name: %s, got %s", expectedResult.name, result.name)
  }

  if len(expectedResult.params) != len(result.params) {
    t.Fatalf("expected command param length: %q, got %q", expectedResult.params, result.params)
  }

  for index, param := range(result.params) {
    if expectedResult.params[index] != param {
      t.Fatalf("expected param %d to be %s, but got %s", index, expectedResult.params[index], param)
    }
  }
}

func TestExtractGetCommand(t *testing.T) {
  payload := "GET a b"
  expectedResult := Command{"GET", []string{"a", "b"}}
  result, err := ExtractCommand(payload)

  if err != nil {
    t.Fatalf("Unexpected error returned")
  }

  if expectedResult.name != result.name {
    t.Fatalf("expected command name: %s, got %s", expectedResult.name, result.name)
  }

  if len(expectedResult.params) != len(result.params) {
    t.Fatalf("expected command param length: %q, got %q", expectedResult.params, result.params)
  }

  for index, param := range(result.params) {
    if expectedResult.params[index] != param {
      t.Fatalf("expected param %d to be %s, but got %s", index, expectedResult.params[index], param)
    }
  }
}

func TestExtractDeleteCommand(t *testing.T) {
  payload := "DELETE a"
  expectedResult := Command{"DELETE", []string{"a"}}
  result, err := ExtractCommand(payload)

  if err != nil {
    t.Fatalf("Unexpected error returned")
  }

  if expectedResult.name != result.name {
    t.Fatalf("expected command name: %s, got %s", expectedResult.name, result.name)
  }

  if len(expectedResult.params) != len(result.params) {
    t.Fatalf("expected command param length: %q, got %q", expectedResult.params, result.params)
  }

  for index, param := range(result.params) {
    if expectedResult.params[index] != param {
      t.Fatalf("expected param %d to be %s, but got %s", index, expectedResult.params[index], param)
    }
  }
}
