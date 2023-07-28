package entities

import "testing"

func Test_position_UpdateDirection(t *testing.T) {
	type fields struct {
		direction Direction
		X         int
		Y         int
	}

	tests := []struct {
		name      string
		fields    fields
		increment int
		want      Direction
	}{
		{
			name: "increase one from north position",
			fields: fields{
				direction: North,
			},
			increment: 1,
			want:      East,
		},
		{
			name: "increase one from east position",
			fields: fields{
				direction: East,
			},
			increment: 1,
			want:      South,
		},
		{
			name: "increase one from south position",
			fields: fields{
				direction: South,
			},
			increment: 1,
			want:      West,
		},
		{
			name: "increase one from west position",
			fields: fields{
				direction: West,
			},
			increment: 1,
			want:      North,
		},
		{
			name: "decrease one from north position",
			fields: fields{
				direction: North,
			},
			increment: -1,
			want:      West,
		},
		{
			name: "decrease one from west position",
			fields: fields{
				direction: West,
			},
			increment: -1,
			want:      South,
		},
		{
			name: "decrease one from south position",
			fields: fields{
				direction: South,
			},
			increment: -1,
			want:      East,
		},
		{
			name: "decrease one from north position",
			fields: fields{
				direction: East,
			},
			increment: -1,
			want:      North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := position{
				direction: tt.fields.direction,
			}
			p.UpdateDirection(tt.increment)
			if tt.want != p.direction {
				t.Errorf("expected %v, got %v", tt.want, p.direction)
			}
		})
	}
}
