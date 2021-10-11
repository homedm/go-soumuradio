package soumuradio

import (
	"reflect"
	"testing"
)

func TestParseRadioSpec1(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []RadioSpec
	}{
		{
			name: "case1",
			str:  "3MA\\t1910 kHz\\t50   W\\n3HA\\t3537.5 kHz\\t50   W\\n3HD\\t3798 kHz\\t50   W\\n3HA\\t7100 kHz\\t50   W\\n2HC\\t10125 kHz\\t50   W\\n2HA\\t14175 kHz\\t50   W\\n3HA\\t18118 kHz\\t50   W\\n3HA\\t21225 kHz\\t50   W\\n3HA\\t24940 kHz\\t50   W\\n3VA\\t28.85 MHz\\t50   W\\n3VA\\t52 MHz\\t50   W\\n3VA\\t145 MHz\\t50   W\\n3VA\\t435 MHz\\t50   W\\nA1A\\t4630 kHz\\t50   W",
			want: []RadioSpec{
				{RadioFormat: []string{"3MA"}, Freq: "1910kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "3537.5kHz", Power: 50},
				{RadioFormat: []string{"3HD"}, Freq: "3798kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "7100kHz", Power: 50},
				{RadioFormat: []string{"2HC"}, Freq: "10125kHz", Power: 50},
				{RadioFormat: []string{"2HA"}, Freq: "14175kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "18118kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "21225kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "24940kHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "28.85MHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "52MHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "145MHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "435MHz", Power: 50},
				{RadioFormat: []string{"A1A"}, Freq: "4630kHz", Power: 50},
			},
		},
		{
			name: "case2",
			str:  "3MA\\t1910 kHz\\t50   W\\n3HA\\t3537.5 kHz\\t50   W\\n3HD\\t3798 kHz\\t50   W\\n3HA\\t7100 kHz\\t50   W\\n2HC\\t10125 kHz\\t50   W\\n2HA\\t14175 kHz\\t50   W\\n3HA\\t18118 kHz\\t50   W\\n3HA\\t21225 kHz\\t50   W\\n3HA\\t24940 kHz\\t50   W\\n3VA\\t28.85 MHz\\t50   W\\n3VA\\t52 MHz\\t50   W\\n3VA\\t145 MHz\\t50   W\\n3VA\\t435 MHz\\t50   W\\n3SA\\t1280 MHz\\t10   W\\n3SA\\t2425 MHz\\t1   W\\n3SA\\t5750 MHz\\t2   W\\n3SA\\t10.125 GHz\\t2   W\\n3SA\\t10.475 GHz\\t2   W",
			want: []RadioSpec{
				{RadioFormat: []string{"3MA"}, Freq: "1910kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "3537.5kHz", Power: 50},
				{RadioFormat: []string{"3HD"}, Freq: "3798kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "7100kHz", Power: 50},
				{RadioFormat: []string{"2HC"}, Freq: "10125kHz", Power: 50},
				{RadioFormat: []string{"2HA"}, Freq: "14175kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "18118kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "21225kHz", Power: 50},
				{RadioFormat: []string{"3HA"}, Freq: "24940kHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "28.85MHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "52MHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "145MHz", Power: 50},
				{RadioFormat: []string{"3VA"}, Freq: "435MHz", Power: 50},
				{RadioFormat: []string{"3SA"}, Freq: "1280MHz", Power: 10},
				{RadioFormat: []string{"3SA"}, Freq: "2425MHz", Power: 1},
				{RadioFormat: []string{"3SA"}, Freq: "5750MHz", Power: 2},
				{RadioFormat: []string{"3SA"}, Freq: "10.125GHz", Power: 2},
				{RadioFormat: []string{"3SA"}, Freq: "10.475GHz", Power: 2},
			},
		},
		{
			name: "case3",
			str:  "500HA1A\\t437.075 MHz\\t100  mW\\n16K0F2D\\t437.075 MHz\\t800  mW\\n16K0F3E\\t \\t \\n16K0F3F\\t \\t \\n26K0F1D\\t437.075 MHz\\t800  mW\\n30K0G1D\\t435.9 MHz\\t400  mW\\n30K0F1D\\t \\t \\n500HA1A\\t435.88 MHz　から　      435.91 MHz　まで\\t500  mW\\n3K00J3E\\t435.88 MHz　から　      435.91 MHz　まで\\t500  mW",
			want: []RadioSpec{
				{RadioFormat: []string{"500HA1A"}, Freq: "437.075MHz", Power: 0.1},
				{RadioFormat: []string{"16K0F2D"}, Freq: "437.075MHz", Power: 0.8},
				{RadioFormat: []string{"16K0F3E"}, Freq: "", Power: 0},
				{RadioFormat: []string{"16K0F3F"}, Freq: "", Power: 0},
				{RadioFormat: []string{"26K0F1D"}, Freq: "437.075MHz", Power: 0.8},
				{RadioFormat: []string{"30K0G1D"}, Freq: "435.9MHz", Power: 0.4},
				{RadioFormat: []string{"30K0F1D"}, Freq: "", Power: 0},
				{RadioFormat: []string{"500HA1A"}, Freq: "435.88MHzから435.91MHzまで", Power: 0.5},
				{RadioFormat: []string{"3K00J3E"}, Freq: "435.88MHzから435.91MHzまで", Power: 0.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := parseRadioSpec1(tt.str); err == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("parseRadioSpec = \nget %v,\n want %v", got, tt.want)
				}
			} else {
				t.Fatal(err)
			}
		})
	}
}
