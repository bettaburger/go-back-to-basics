package main

import (
	"reflect"
	"testing"
)

func TestTwoSumsHash(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSumsHash(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSumsHash(%v, %d) = %v; want %v",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func TestTwoSumsBrute(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSumsBrute(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSumsBrute(%v, %d) = %v; want %v",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

// to run all tests 
// go test -v

// to run individual test
// go test -v -run TestTwoSumsHash

