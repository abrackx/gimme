package util

import "testing"

func TestGenerateName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Should work",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateName(); got != tt.want {
				t.Errorf("GenerateName() = %v, want %v", got, tt.want)
			}
		})
	}
}
