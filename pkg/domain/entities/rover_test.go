package entities

import "testing"

func TestRover_Explore(t *testing.T) {
	type args struct {
		initialX     int
		initialY     int
		direction    string
		instructions []rune
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid case",
			args: args{
				initialX:     1,
				initialY:     2,
				direction:    "N",
				instructions: []rune{'L', 'M', 'L', 'M', 'L', 'M', 'L', 'M', 'M'},
			},
			want:    "1 3 N",
			wantErr: false,
		},
		{
			name: "valid case",
			args: args{
				initialX:     3,
				initialY:     3,
				direction:    "E",
				instructions: []rune{'M', 'M', 'R', 'M', 'M', 'R', 'M', 'R', 'R', 'M'},
			},
			want:    "5 1 E",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := NewRover(tt.args.initialX, tt.args.initialY, tt.args.direction, tt.args.instructions)
			got, err := r.Explore(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Explore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Explore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
