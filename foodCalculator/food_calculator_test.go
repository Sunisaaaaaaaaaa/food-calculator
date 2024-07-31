package main_test

import (
	cal "food_calculator/foodCalculator"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name     string
		order    []cal.OrderDetail
		member   bool
		expected float32
	}{
		{
			name: "discount, no membership",
			order: []cal.OrderDetail{
				{"Orange", 2},
				{"Pink", 3},
				{"Green", 5},
			},
			member:   false,
			expected: 652,
		}, {
			name: "no discount, no membership",
			order: []cal.OrderDetail{
				{"Red", 2},
				{"Blue", 3},
				{"Yellow", 1},
				{"Purple", 4},
			},
			member:   false,
			expected: 600,
		}, {
			name: "discount, membership",
			order: []cal.OrderDetail{
				{"Orange", 2},
				{"Pink", 3},
				{"Red", 1},
			},
			member:   true,
			expected: 459,
		},
		{
			name: "no discount, membership",
			order: []cal.OrderDetail{
				{"Blue", 3},
				{"Yellow", 1},
				{"Purple", 4},
			},
			member:   true,
			expected: 450,
		}, {
			name: "food not found",
			order: []cal.OrderDetail{
				{"Unknown", 2},
			},
			member:   false,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc := cal.FoodCalculator{
				Order:  tc.order,
				Member: tc.member,
			}

			total := fc.Calculator()
			if total != tc.expected {
				t.Errorf("expected: %f, total: %f", tc.expected, total)
			}
		})
	}
}
