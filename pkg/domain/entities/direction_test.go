package entities

import "testing"

func TestNewDirection(t *testing.T) {
	tests := []struct {
		name      string
		direction string
		want      Direction
		wantErr   bool
	}{
		{
			name:      "north",
			direction: "N",
			want:      North,
			wantErr:   false,
		},
		{
			name:      "east",
			direction: "E",
			want:      East,
			wantErr:   false,
		},
		{
			name:      "south",
			direction: "S",
			want:      South,
			wantErr:   false,
		},
		{
			name:      "west",
			direction: "W",
			want:      West,
			wantErr:   false,
		},
		{
			name:      "invalid",
			direction: "Z",
			want:      Direction(0),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDirection(tt.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewDirection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
