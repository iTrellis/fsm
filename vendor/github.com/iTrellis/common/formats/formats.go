/*
Copyright © 2020 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package formats

import (
	"math/big"
	"regexp"
	"time"
)

const (
	timeReg = `^(?P<value>([0-9]+(\.[0-9]+)?))\s*(?P<unit>(nanoseconds|nanosecond|nanos|nano|ns|microseconds|microsecond|micros|micro|us|milliseconds|millisecond|millis|milli|ms|seconds|second|s|minutes|minute|m|hours|hour|h|days|day|d))$`

	bitReg = `^(?P<value>([0-9]+(\.[0-9]+)?))\s*(?P<unit>(b|byte|bytes|kb|kilobyte|kilobytes|mb|megabyte|megabytes|gb|gigabyte|gigabytes|tb|terabyte|terabytes|pb|petabyte|petabytes|eb|exabyte|exabytes|zb|zettabyte|zettabytes|yb|yottabyte|yottabytes|k|ki|kib|kibibyte|kibibytes|m|mi|mib|mebibyte|mebibytes|g|gi|gib|gibibyte|gibibytes|t|ti|tib|tebibyte|tebibytes|p|pi|pib|pebibyte|pebibytes|e|ei|eib|exbibyte|exbibytes|z|zi|zib|zebibyte|zebibytes|y|yi|yib|yobibyte|yobibytes))$`
)

// FindStringSubmatchMap infomation:
// returns a map of strings holding the text of the
// leftmost match of the regular expression in s and the matches, if any, of
// its subexpressions, as defined by the 'Submatch' description in the
// package comment.
// A return value of nil indicates no match.
func FindStringSubmatchMap(s, exp string) (map[string]string, bool) {
	reg := regexp.MustCompile(exp)
	captures := make(map[string]string)

	match := reg.FindStringSubmatch(s)
	if match == nil {
		return captures, false
	}

	for i, name := range reg.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}
		captures[name] = match[i]
	}
	return captures, true
}

// ByteSizes
var (
	_Num1000 = big.NewInt(1000)
	_Num1024 = big.NewInt(1024)

	_Byte   = big.NewInt(1)
	_KiByte = (&big.Int{}).Mul(_Byte, _Num1024)
	_MiByte = (&big.Int{}).Mul(_KiByte, _Num1024)
	_GiByte = (&big.Int{}).Mul(_MiByte, _Num1024)
	_TiByte = (&big.Int{}).Mul(_GiByte, _Num1024)
	_PiByte = (&big.Int{}).Mul(_TiByte, _Num1024)
	_EiByte = (&big.Int{}).Mul(_PiByte, _Num1024)
	_ZiByte = (&big.Int{}).Mul(_EiByte, _Num1024)
	_YiByte = (&big.Int{}).Mul(_ZiByte, _Num1024)

	_KByte = (&big.Int{}).Mul(_Byte, _Num1000)
	_MByte = (&big.Int{}).Mul(_KByte, _Num1000)
	_GByte = (&big.Int{}).Mul(_MByte, _Num1000)
	_TByte = (&big.Int{}).Mul(_GByte, _Num1000)
	_PByte = (&big.Int{}).Mul(_TByte, _Num1000)
	_EByte = (&big.Int{}).Mul(_PByte, _Num1000)
	_ZByte = (&big.Int{}).Mul(_EByte, _Num1000)
	_YByte = (&big.Int{}).Mul(_ZByte, _Num1000)
)

// ParseStringByteSize return big size
func ParseStringByteSize(key string) *big.Int {
	groups, matched := FindStringSubmatchMap(key, bitReg)
	if !matched {
		return nil
	}
	i, _ := ToInt64(groups["value"])

	switch groups["unit"] {
	case "b", "byte", "bytes":
		return (&big.Int{}).Mul(big.NewInt(i), _Byte)
	case "kb", "kilobyte", "kilobytes":
		return (&big.Int{}).Mul(big.NewInt(i), _KByte)
	case "mb", "megabyte", "megabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _MByte)
	case "gb", "gigabyte", "gigabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _GByte)
	case "tb", "terabyte", "terabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _TByte)
	case "pb", "petabyte", "petabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _PByte)
	case "eb", "exabyte", "exabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _EByte)
	case "zb", "zettabyte", "zettabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _ZByte)
	case "yb", "yottabyte", "yottabytes":
		return (&big.Int{}).Mul(big.NewInt(i), _YByte)
	case "k", "ki", "kib", "kibibyte", "kibibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _KiByte)
	case "m", "mi", "mib", "mebibyte", "mebibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _MiByte)
	case "g", "gi", "gib", "gibibyte", "gibibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _GiByte)
	case "t", "ti", "tib", "tebibyte", "tebibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _TiByte)
	case "p", "pi", "pib", "pebibyte", "pebibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _PiByte)
	case "e", "ei", "eib", "exbibyte", "exbibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _EiByte)
	case "z", "zi", "zib", "zebibyte", "zebibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _ZiByte)
	case "y", "yi", "yib", "yobibyte", "yobibytes":
		return (&big.Int{}).Mul(big.NewInt(i), _YiByte)
	default:
		return nil
	}
}

// ParseStringTime return time.Duration
func ParseStringTime(s string, defValue ...time.Duration) time.Duration {
	groups, matched := FindStringSubmatchMap(s, timeReg)

	if !matched {
		if len(defValue) == 0 {
			return 0
		}
		return defValue[0]
	}

	i, _ := ToInt64(groups["value"])

	switch groups["unit"] {
	case "nanoseconds", "nanosecond", "nanos", "nano", "ns":
		return time.Nanosecond * time.Duration(i)
	case "microseconds", "microsecond", "micros", "micro", "us":
		return time.Microsecond * time.Duration(i)
	case "milliseconds", "millisecond", "millis", "milli", "ms":
		return time.Millisecond * time.Duration(i)
	case "seconds", "second", "s":
		return time.Second * time.Duration(i)
	case "minutes", "minute", "m":
		return time.Minute * time.Duration(i)
	case "hours", "hour", "h":
		return time.Hour * time.Duration(i)
	case "days", "day", "d":
		return time.Hour * 24 * time.Duration(i)
	default:
		return 0
	}
}
