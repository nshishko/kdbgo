package kdb

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xff, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if d.(bool) {
		t.Fail()
	}
}

// 1i
func TestInt(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x01, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if d != int32(1) {
		t.Fail()
	}
}

// `GOOG
func TestSymbol(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x00, 0x00, 0xf5, 0x47, 0x4f, 0x4f, 0x47, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	fmt.Println(d)
	if d != "GOOG" {
		t.Fail()
	}
}

// "GOOG"
func TestCharArray(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x04, 0x00, 0x00, 0x00, 0x47, 0x4f, 0x4f, 0x47}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	fmt.Println(d)
	if d != "GOOG" {
		t.Fail()
	}
}
func ExampleInt(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0xfa, 0x01, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	Decode(r)
}

// enlist 1i
func TestIntVector(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if vec, ok := d.([]int32); ok {
		fmt.Println(vec)

		if len(vec) != 1 || vec[0] != int32(1) {
			t.Fail()
		}
	}

}

// `byte$til 5
func TestByteVector(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x04, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if vec, ok := d.([]byte); ok {
		fmt.Println(vec)

		if len(vec) != 5 || vec[4] != 0x04 {
			t.Fail()
		}
	}

}

// 1?0Ng - enlist ddb87915-b672-2c32-a6cf-296061671e9d
func TestGUIDVector(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01,
		0x00, 0x00, 0x00, 0xdd, 0xb8, 0x79, 0x15, 0xb6, 0x72, 0x2c, 0x32, 0xa6, 0xcf, 0x29, 0x60, 0x61, 0x67, 0x1e, 0x9d}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if vec, ok := d.([]uuid.UUID); ok {
		fmt.Println(vec[0].String())

		if len(vec) != 1 || vec[0].String() != "ddb87915-b672-2c32-a6cf-296061671e9d" {
			t.Fail()
		}
	}

}
func TestGUID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00,
		0xfe, 0xdd, 0xb8, 0x79, 0x15, 0xb6, 0x72, 0x2c, 0x32, 0xa6, 0xcf, 0x29, 0x60, 0x61, 0x67, 0x1e, 0x9d}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	var d1 uuid.UUID
	d1 = d.(uuid.UUID)
	fmt.Println(d1.String())
	if d1.String() != "ddb87915-b672-2c32-a6cf-296061671e9d" {
		t.Fail()
	}
}

//q)-8!0N!0D01:22:33.444555666*1+til 2
// 0D01:22:33.444555666 0D02:45:06.889111332
func TestTimespanVector(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x10, 0x00, 0x02, 0x00, 0x00, 0x00,
		0x92, 0x9b, 0x4d, 0x50, 0x81, 0x04, 0x00, 0x00, 0x24, 0x37, 0x9b, 0xa0, 0x02, 0x09, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if vec, ok := d.([]time.Duration); ok {
		fmt.Println(vec[0].String())

		if len(vec) != 2 || vec[0].String() != "1h22m33.444555666s" {
			t.Fail()
		}
	}

}

// 	q)-8!`abc`bc`c
func TestSymbolVec(t *testing.T) {

	b := []byte{0x01, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x0b,
		0x00, 0x03, 0x00, 0x00, 0x00, 0x61, 0x62, 0x63, 0x00, 0x62, 0x63, 0x00, 0x63, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	if vec, ok := d.([]string); ok {
		fmt.Println(vec)

		if len(vec) != 3 || vec[0] != "abc" || vec[1] != "bc" || vec[2] != "c" {
			t.Fail()
		}
	}

}

// -8!'type
func TestError(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x80, 0x74, 0x79, 0x70, 0x65, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	if d != nil {
		t.Fail()
	}

	if err.Error() != "type" {
		t.Fail()
	}

	fmt.Println(err)
}

//
//q)-8!`a`b!2 3
func TestDictWithAtoms(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00,
		0x63, 0x0b, 0x00, 0x02, 0x00, 0x00, 0x00, 0x61, 0x00, 0x62, 0x00, 0x06, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, _ := Decode(r)
	fmt.Println(d)

}

//-8!`s#`a`b!2 3
func TestSortedDict(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x7f, 0x0b, 0x01, 0x02,
		0x00, 0x00, 0x00, 0x61, 0x00, 0x62, 0x00, 0x06,
		0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Sorted dict", d, err)
	if err != nil {
		t.Fail()
	}

}

//-8!`a`b!enlist each 2 3
func TestDictWithVectors(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x63, 0x0b, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x61, 0x00, 0x62, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
		0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Dict with vectors", d, err)

}

// ([]a:enlist 2;b:enlist 3)
func TestTable(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x2f, 0x00, 0x00, 0x00, 0x62, 0x00, 0x63, 0x0b, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x61, 0x00, 0x62, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Table:", d, err)

}

//-8!`s#([]a:enlist 2;b:enlist 3)
func TestSortedTable(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x2f, 0x00, 0x00, 0x00, 0x62, 0x01, 0x63, 0x0b, 0x00, 0x02, 0x00,
		0x00, 0x00, 0x61, 0x00, 0x62, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x03, 0x01, 0x00,
		0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Sorted Table:", d, err)

}

// -8!([a:enlist 2]b:enlist 3)
func TestKeyedTable(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x63, 0x62, 0x00, 0x63, 0x0b, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x61, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x62, 0x00, 0x63, 0x0b, 0x00, 0x01, 0x00, 0x00, 0x00, 0x62, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Keyed Table:", d, err)

}

// -8!`s#([a:enlist 2]b:enlist 3)
func TestSortedKeyedTable(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x7f, 0x62, 0x00, 0x63, 0x0b, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x61, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x62, 0x00, 0x63, 0x0b, 0x00, 0x01, 0x00, 0x00, 0x00, 0x62, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Sorted Keyed Table:", d, err)

}

// -8!{x+y}
func TestFunc(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x64, 0x00, 0x0a,
		0x00, 0x05, 0x00, 0x00, 0x00, 0x7b, 0x78, 0x2b, 0x79, 0x7d}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Function:", d, err)
}

//q)\d .d
//q.d)test:{x+y}
//q.d)-8!test
func TestFuncNonRoot(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x64, 0x64, 0x00, 0x0a,
		0x00, 0x05, 0x00, 0x00, 0x00, 0x7b, 0x78, 0x2b, 0x79, 0x7d}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println("Function:", d, err)
}

// `byte$enlist til 5
func TestGeneralList(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04}
	r := bufio.NewReader(bytes.NewReader(b))
	Decode(r)
}

//q)-8!1986.07.23D03:10:45.000639000 2013.06.10D20:49:14.999361000
func TestTimestampVec(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x02, 0x00, 0x00, 0x00, 0x18, 0x92,
		0x00, 0xc6, 0xed, 0xe4, 0x1c, 0xfa, 0xe8, 0x6d, 0xff, 0x39, 0x12, 0x1b, 0xe3, 0x05}
	r := bufio.NewReader(bytes.NewReader(b))
	d, err := Decode(r)
	fmt.Println(d)
	if err != nil {
		t.Error("Decoding failed.", err)
	}
	if vec, ok := d.([]time.Time); ok {
		if len(vec) != 2 || vec[0] != time.Date(1986, time.July, 23, 03, 10, 45, 639000, time.UTC) || vec[1] != time.Date(2013, time.June, 10, 20, 49, 14, 999361000, time.UTC) {
			t.Error("Decoding is incorrect. Result was ", vec)
		}
	} else {
		t.Error("Result is not time array")
	}

}

// -8!2013.06m +til 3
func TestMonthList(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x03, 0x00, 0x00, 0x00, 0xa1, 0x00,
		0x00, 0x00, 0xa2, 0x00, 0x00, 0x00, 0xa3, 0x00, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	fmt.Println(Decode(r))
}

// -8!21:22*til 2
func TestMinuteList(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x11, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x05, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	fmt.Println(Decode(r))

}

// -8!21:22:01 + 1 2
func TestSecondList(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x12, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7a, 0x2c, 0x01, 0x00, 0x7b, 0x2c, 0x01, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	fmt.Println(Decode(r))

}

//-8!1#21:53:37.963
func TestTimeVec(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x13, 0x00, 0x01, 0x00, 0x00, 0x00, 0xab, 0xaa, 0xb2, 0x04}
	r := bufio.NewReader(bytes.NewReader(b))
	fmt.Println(Decode(r))

}

//-8!1#2013.06.10
func TestDateVec(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x01, 0x00, 0x00, 0x00, 0x2d, 0x13, 0x00, 0x00}
	r := bufio.NewReader(bytes.NewReader(b))
	fmt.Println(Decode(r))

}

//-8!1#2013.06.10T22:03:49.713
func TestDateTimeVec(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd6, 0x81, 0xe8, 0x58, 0xeb, 0x2d, 0xb3, 0x40}
	r := bufio.NewReader(bytes.NewReader(b))
	fmt.Println(Decode(r))

}
