package units

import (
	"reflect"
	"testing"

	"gitlab.ozon.ru/re/teams/itc-quotas/internal/quotas/measure/prefix"
	"gitlab.ozon.ru/re/teams/itc-quotas/internal/quotas/measure/prefix/binary"
)

func Test_value_Prettier(t *testing.T) {
	type testCase[P prefix.Prefix[P]] struct {
		name string
		m    value[P]
		want value[P]
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Prettier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prettier() = %v, want %v", got, tt.want)
			}
		})
	}
}
