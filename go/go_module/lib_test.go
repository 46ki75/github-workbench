package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"addition of positive integers", 1, 2, 3},
		{"addition of negative integers", -1, -2, -3},
		{"addition of zeros", 0, 0, 0},
		{"addition of positive and negative integers", -1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestGenerateID(t *testing.T) {
	tests := []struct {
		name        string
		length      int
		expectError bool
		expectedLen int
	}{
		{"Valid length 10", 10, false, 10},
		{"Valid length 21 (default NanoID)", 21, false, 21},
		{"Zero length", 0, true, 0},
		{"Negative length", -1, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := GenerateID(tt.length)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("did not expect error but got: %v", err)
				return
			}

			if len(id) != tt.expectedLen {
				t.Errorf("expected ID length %d, but got %d", tt.expectedLen, len(id))
			}
		})
	}
}
