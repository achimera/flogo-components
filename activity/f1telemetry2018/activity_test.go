package f1telemetry2018

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/lunixbochs/struc"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalQuality(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing -  MSGTYPE 3")

	var buf bytes.Buffer

	arr, err1 := hex.DecodeString("e2070103eb15f8232a247e7200000000000000001353535441")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}

	tc.SetInput("buffer", arr)
	err := struc.Pack(&buf, arr)

	o := &F1Header{}
	o2 := &F1Event{}

	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)
	if err != nil {
		fmt.Printf("error code: %v \n", err)
		return
	}

	err = struc.Unpack(&buf, o2)
	fmt.Printf("F1 Event : \n %+x \n", o2)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype := tc.GetOutput("msgtype")
	rdata := tc.GetOutput("data")
	rdata2 := tc.GetOutput("array")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)
	fmt.Printf("csv array: %v \n", rdata2)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println("#######   Testing -  MSGTYPE 1")

	arr, err1 = hex.DecodeString("e20701010dbc53c77cf22db20e745a3c0000000013001a13052a120a0400201c201c500000ff00126b46783f006fdb943d00957f1b3e00c7f75d3e008ea3763e00f28aa83e00cb2ec13e003edcd83e00b7df043f00014e0f3f00521b1b3f0089b6243f000ca6383f00dff0473f000bea4d3f00ef595a3f001101623f00f29e683f000000000000000000000000000000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")
	rdata2 = tc.GetOutput("array")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)
	fmt.Printf("csv array: %v \n", rdata2)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println("#######   Testing -  MSGTYPE 5")
	arr, err1 = hex.DecodeString("e20701050dbc53c77cf22db20e745a3c000000001306064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000204106064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac410600002041")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")
	rdata2 = tc.GetOutput("array")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)
	fmt.Printf("csv array: %v \n", rdata2)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println("#######   Testing -  MSGTYPE ?")
	arr, err1 = hex.DecodeString("e20701040dbc53c77cf22db20e745a3c000000001314013c061c36422e20484152544c455900000000000000000000000000000000000000000000000000000000000000000000000000000103080e4e462e20414c4f4e534f000000000000000000000000000000000000000000000000000000000000000000000000000000010d01051d532e2056455454454c000000000000000000000000000000000000000000000000000000000000000000000000000000010e040b34532e20504552455a00000000000000000000000000000000000000000000000000000000000000000000000000000000011303120d4c2e205354524f4c4c000000000000000000000000000000000000000000000000000000000000000000000000000000010601071b4b2e2052c384494b4bc3964e454e00000000000000000000000000000000000000000000000000000000000000000000013a091035432e204c45434c45524300000000000000000000000000000000000000000000000000000000000000000000000000000112080207532e2056414e444f4f524e4500000000000000000000000000000000000000000000000000000000000000000000000001090221164d2e205645525354415050454e0000000000000000000000000000000000000000000000000000000000000000000000010809094f4d2e204552494353534f4e00000000000000000000000000000000000000000000000000000000000000000000000000010f004d1b562e20424f54544153000000000000000000000000000000000000000000000000000000000000000000000000000000010a051b1d4e2e2048c39c4c4b454e42455247000000000000000000000000000000000000000000000000000000000000000000000102020303442e2052494343494152444f000000000000000000000000000000000000000000000000000000000000000000000000010b0714154b2e204d41474e555353454e000000000000000000000000000000000000000000000000000000000000000000000000010c07081c522e2047524f534a45414e000000000000000000000000000000000000000000000000000000000000000000000000000111041f1c452e204f434f4e0000000000000000000000000000000000000000000000000000000000000000000000000000000000013d032344532e205349524f544b494e00000000000000000000000000000000000000000000000000000000000000000000000000010005374e432e205341494e5a00000000000000000000000000000000000000000000000000000000000000000000000000000000013b060a1c502e204741534c59000000000000000000000000000000000000000000000000000000000000000000000000000000000007002c0a4c2e2048414d494c544f4e00000000000000000000000000000000000000000000000000000000000000000000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")
	rdata2 = tc.GetOutput("array")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)
	fmt.Printf("csv array: %v \n", rdata2)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}
}
