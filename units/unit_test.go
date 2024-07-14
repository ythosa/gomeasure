package units

import (
	"reflect"
	"testing"

	"github.com/ythosa/gomeasure/prefix"
	"github.com/ythosa/gomeasure/prefix/binary"
)

func Test_value_Prettier(t *testing.T) {
	t.Parallel()

	type testCase[P prefix.Prefix[P]] struct {
		name string
		m    unit[P]
		want unit[P]
	}
	tests := []testCase[binary.Prefix]{
		{
			name: "already a lot of prettier <3",
			m:    newValue(123, binary.Tebi),
			want: newValue(123, binary.Tebi),
		},
		{
			name: "need prettier",
			m:    newValue(4096, binary.Kibi),
			want: newValue(4, binary.Mebi),
		},
		{
			name: "not prettier but max in scale",
			m:    newValue(2048, binary.Pebi),
			want: newValue(2048, binary.Pebi),
		},
		{
			name: "already prettier in minimal scale",
			m:    newValue(1023, binary.None),
			want: newValue(1023, binary.None),
		},
		{
			name: "negative",
			m:    newValue(-4096, binary.Kibi),
			want: newValue(-4, binary.Mebi),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.prettier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prettier() = %v, want %v", got, tt.want)
			}
		})
	}
}
