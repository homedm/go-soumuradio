package soumuradio

import "testing"

func TestParseRadioSpec1(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []RadioSpec
	}{
		{
			name: "jj1hgp",
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
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := parseRadioSpec1(tt.str); err == nil {
				t.Errorf("parseRadioSpec = \nget %v,\n want %v", got, tt.want)
			} else {
				t.Fatal(err)
			}
		})
	}
}
