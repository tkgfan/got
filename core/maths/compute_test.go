// author gmfan
// date 2024/5/8

package maths

import (
	"testing"
)

func TestRatioToAB(t *testing.T) {
	tests := []struct {
		name         string
		ra, rb, a, b int
		wantA, wantB int
	}{
		{
			name:  "test1",
			ra:    1,
			rb:    2,
			a:     10,
			b:     10,
			wantA: 1,
			wantB: 2,
		},
		{
			name:  "test2",
			ra:    73,
			rb:    103,
			a:     10,
			b:     10,
			wantA: 5,
			wantB: 7,
		},
		{
			name:  "test3",
			ra:    130,
			rb:    88,
			a:     10,
			b:     10,
			wantA: 3,
			wantB: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA, gotB := RatioToAB(tt.ra, tt.rb, tt.a, tt.b); gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("RatioToAB() = %v, %v, want %v, %v", gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}
