package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input   int
		want    string
		wantErr error
	}{
		{-1, "", fmt.Errorf("unable to convert numbers less than 1")},
		{0, "", fmt.Errorf("unable to convert numbers less than 1")},
		{1, "I", nil},
		{2, "II", nil},
		{3, "III", nil},
		{4, "IV", nil},
		{5, "V", nil},
		{6, "VI", nil},
		{7, "VII", nil},
		{8, "VIII", nil},
		{9, "IX", nil},
		{10, "X", nil},
		{11, "XI", nil},
		{12, "XII", nil},
		{13, "XIII", nil},
		{14, "XIV", nil},
		{15, "XV", nil},
		{16, "XVI", nil},
		{17, "XVII", nil},
		{18, "XVIII", nil},
		{19, "XIX", nil},
		{20, "XX", nil},
		{21, "XXI", nil},
		{22, "XXII", nil},
		{23, "XXIII", nil},
		{24, "XXIV", nil},
		{25, "XXV", nil},
		{26, "XXVI", nil},
		{27, "XXVII", nil},
		{28, "XXVIII", nil},
		{29, "XXIX", nil},
		{30, "XXX", nil},
		{31, "XXXI", nil},
		{32, "XXXII", nil},
		{33, "XXXIII", nil},
		{34, "XXXIV", nil},
		{35, "XXXV", nil},
		{36, "XXXVI", nil},
		{37, "XXXVII", nil},
		{38, "XXXVIII", nil},
		{39, "XXXIX", nil},
		{40, "XL", nil},
		{41, "XLI", nil},
		{42, "XLII", nil},
		{43, "XLIII", nil},
		{44, "XLIV", nil},
		{45, "XLV", nil},
		{46, "XLVI", nil},
		{47, "XLVII", nil},
		{48, "XLVIII", nil},
		{49, "XLIX", nil},
		{50, "L", nil},
		{51, "LI", nil},
		{52, "LII", nil},
		{53, "LIII", nil},
		{54, "LIV", nil},
		{55, "LV", nil},
		{56, "LVI", nil},
		{57, "LVII", nil},
		{58, "LVIII", nil},
		{59, "LIX", nil},
		{60, "LX", nil},
		{61, "LXI", nil},
		{62, "LXII", nil},
		{63, "LXIII", nil},
		{64, "LXIV", nil},
		{65, "LXV", nil},
		{66, "LXVI", nil},
		{67, "LXVII", nil},
		{68, "LXVIII", nil},
		{69, "LXIX", nil},
		{70, "LXX", nil},
		{71, "LXXI", nil},
		{72, "LXXII", nil},
		{73, "LXXIII", nil},
		{74, "LXXIV", nil},
		{75, "LXXV", nil},
		{76, "LXXVI", nil},
		{77, "LXXVII", nil},
		{78, "LXXVIII", nil},
		{79, "LXXIX", nil},
		{80, "LXXX", nil},
		{81, "LXXXI", nil},
		{82, "LXXXII", nil},
		{83, "LXXXIII", nil},
		{84, "LXXXIV", nil},
		{85, "LXXXV", nil},
		{86, "LXXXVI", nil},
		{87, "LXXXVII", nil},
		{88, "LXXXVIII", nil},
		{89, "LXXXIX", nil},
		{90, "XC", nil},
		{91, "XCI", nil},
		{92, "XCII", nil},
		{93, "XCIII", nil},
		{94, "XCIV", nil},
		{95, "XCV", nil},
		{96, "XCVI", nil},
		{97, "XCVII", nil},
		{98, "XCVIII", nil},
		{99, "XCIX", nil},
		{100, "C", nil},
		{225, "CCXXV", nil},
		{845, "DCCCXLV", nil},
		{1984, "MCMLXXXIV", nil},
		{2024, "MMXXIV", nil},
		{2550, "MMDL", nil},
		{3999, "MMMCMXCIX", nil},
		{4000, "", fmt.Errorf("unable to convert numbers greater than 3999")},
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			got, err := Convert(test.input)
			if err != nil {
				if test.wantErr == nil {
					t.Errorf("Convert(%d) gave error %q, want nil error", test.input, err)
				} else if err.Error() != test.wantErr.Error() {
					t.Errorf("Convert(%d) gave error %q, want error %q", test.input, err, test.wantErr)
				}
			} else if test.wantErr != nil {
				t.Errorf("Convert(%d) = %q, want error %q", test.input, got, test.wantErr)
			}
			if got != test.want {
				t.Errorf("Convert(%d) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
