package compute

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected *Command
		err      bool
	}{

		{"SET key value", &Command{"SET", []string{"key", "value"}}, false},
		{"SET key", nil, true},
		{"SET", nil, true},
		{"SET key value extra", nil, true},

		{"GET key", &Command{"GET", []string{"key"}}, false},
		{"GET", nil, true},
		{"GET key extra", nil, true},

		{"DEL key", &Command{"DEL", []string{"key"}}, false},
		{"DEL", nil, true},
		{"DEL key extra", nil, true},

		{"INVALID key", nil, true},
		{"", nil, true},

		{"SET key$", nil, true},
		{"GET key$", nil, true},
		{"DEL key$", nil, true},

		{"SET key_1 value_1", &Command{"SET", []string{"key_1", "value_1"}}, false},
		{"GET key_1", &Command{"GET", []string{"key_1"}}, false},
		{"DEL key_1", &Command{"DEL", []string{"key_1"}}, false},

		{"  SET   key   value  ", &Command{"SET", []string{"key", "value"}}, false},
		{"GET    key", &Command{"GET", []string{"key"}}, false},

		{"set key value", nil, true},
		{"Get key", nil, true},
		{"Del key", nil, true},
	}

	for _, test := range tests {
		result, err := Parse(test.input)
		if test.err {
			if err == nil {
				t.Errorf("Expected error for input '%s', got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", test.input, err)
			} else if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%s', expected %v, got %v", test.input, test.expected, result)
			}
		}
	}
}
