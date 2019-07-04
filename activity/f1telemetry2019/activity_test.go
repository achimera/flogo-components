package f1telemetry2019

import (
	"bytes"
	"encoding/hex"
	"fmt"

	//"io/ioutil"
	"testing"

	"github.com/lunixbochs/struc"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

/* func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	aInput := &Input{XmlData: `<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`}
	tc.SetInputObject(aInput)
	done, _ := act.Eval(tc)
	assert.True(t, done)
	aOutput := &Output{}
	err := tc.GetOutputObject(aOutput)
	assert.Nil(t, err)
	assert.Equal(t, "world", aOutput.JsonObject["hello"])
} */

func TestEvalQuality(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 3 - Event Data ###############################################################")

	var buf bytes.Buffer

	arr, err1 := hex.DecodeString("e30701040103c655e025b6e12cb300000000000000001353535441a009d74101")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}

	tc.SetInput("buffer", arr)
	err := struc.Pack(&buf, arr)

	o := &F1Header{}

	err = struc.Unpack(&buf, o)
	if err != nil {
		fmt.Printf("error code: %v \n", err)
		return
	}

	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype := tc.GetOutput("msgtype")
	rdata := tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 0 - Motion Data ###############################################################")

	arr, err1 = hex.DecodeString("e30701040100c655e025b6e12cb3595f643e03000000137410c543daa995c241e49dc3830118c2099a90bf43596a425ebafafd606bac94d7fc57bac74279bd22ff4d3f732f89bf485413bf0065f0bc7e68ca3cc920c5437da795c284129ec3e37617c234e175bf30fb674213ba13fe306bdf947efc0cba10baca3dddd6333fe21dd7bec30b14bf809effbc17bee03ca402c543d9b095c2f5219ec37cdf18c233a77ebfe70c67426db924fec36a4b9594fc67b95199723d2b6e2f3f6db887bdb49615bf007ff8bc7a0ddb3c4cfdc443ccae95c242dd9dc3701419c2d98a7abff6826a4214ba37fe316bdd9488fc0ebafe59d33da10f383f0c885dbe7e0814bfc07cf8bc5a0cde3cf446c543c59d95c2f5789ec30cc515c200fb78bf1dd16342d5b922fe086b07957afccfb9f5818e3e20b5403f3f3a33be549e14bf4044febc2c8ee13cbb0fc543e49e95c268049ec34c9e18c23a9a82bf77f86842eeb92bfe186bf6949efce8b983063e3e8e9b3e3f39a285bedf6114bf4096f5bc0ea9d83c55fdc443f8aa95c27bf09dc3585019c2788a82bf4cb16942e4b90ffe116bfd9485fcdeb9a8c1263ee9423c3fec1203bf527a14bf40aefebca808df3cd424c543659c95c2e8529ec3bfe017c2b2bb7abfd08165426bb924fec26a4c9593fc65b96c633d3d8a63433f67bcd1bd249c15bf80b9f8bc0963db3ca013c543aba395c29a369ec38c4218c256907abff479664282b921fed16a3d9591fc7cb9173f933d7e81363f0f97e4bd186515bf4094f9bc32f1db3c1ff7c44395b195c266ef9dc3322c19c286bd80bf34306942d7b925fe096b069586fcd1b977e54b3e38ed293f72042cbe279a14bf8068fbbc549cde3c05edc44324aa95c2e6e89dc327d219c2556982bfb8fc6942d3b924fe066b089580fccdb9f2c2803e8c42373f6a3755bed4a314bfc0b8fcbcbc11e03c36fec44362a595c23def9dc3217a19c2780a87bf7af16942e0b910fe0e6b019568fcd9b9a11b2d3e4ee3403f5b15c2be748614bfa03a02bd6111e63cb841c543f1a395c203639ec38d3a16c2c21b6dbffd906442d1b928fe056b0a956bfccbb91200673eae6f3f3fdd79c4bd5da814bf805300bdb67ae53cee01c543d5a795c29c069ec38df518c2bbb581bf21bb6842d4b924fe076b07957efcceb96218753eedc93b3f3c035abef79f14bf8038fdbce4ace03ce207c54333ad95c294279ec399aa18c2aeec7bbf35b066426ab920fec16a4d9594fc64b9f2215a3d203a2a3fae4dc0bd449d15bf401cf9bcd31edb3c9624c54319a195c271119ec3f8c217c2dce37fbfa1ce684224ba1efe3b6bd294affc1eba1565a43d3f83483f33e902bfecdf13bf80b7f3bc5051d43c0bfec443a6a095c2b7179ec33ba519c289517ebf262a68426bb927fec26a4c9594fc65b97014373d9c60413fc99fbcbdba9c15bf802cf8bceb3adb3cd2fec44396ae95c2572f9ec3f1d017c294447abff9326742fab926fe206bef9472fcf4b93cd2cc3e9a2b2a3f66b637bcf34514bf408fffbc31b6e33c542bc543039c95c29dff9dc3649916c25a8785bf526a6942a1ba0ffe8c6b8194b4fc9aba01ea18bef134393ff0f671bff6b612bfc0e1f4bc9c2bd33ce31dcf436f9c94c2863fadc31b0a13c237f33ebfdfeb60422fb9f2fd996a7495b5fc28b9d618f5be7b4cf23dc3fe0bbfa12b16bfc097f8bc3d04d33c7ec2854014b590418b735d41f8ee6c415bc485c46ec6c04280c5fd427f1883c39ec1d5c65e0c0fc69aea0f47851b99c693808942936b884286698642be418642c8f1b63ccd687b3cd767e7b7dc54b6b81729013f1400a43edf5c86426b6be3bd5e1444bd1e4bb3be7494fa400ba3cd3fd57b00c2ef71b8bb")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 1 - Session Data ###############################################################")

	arr, err1 = hex.DecodeString("e30701040101c655e025b6e12cb307e8223f0b0000001300211a632a120904000000f000500000ff00126b46783f006fdb943d00957f1b3e00c7f75d3e008ea3763e00f28aa83e00cb2ec13e003edcd83e00b7df043f00014e0f3f00521b1b3f0089b6243f000ca6383f00dff0473f000bea4d3f00ef595a3f001101623f00f29e683f000000000000000000000000000000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 2 - Lap Data ###############################################################")

	arr, err1 = hex.DecodeString("e30701040102c655e025b6e12cb397aa183f0a0000001300000000db263b4000000000000000000000000080be4bc380be4bc30000008006010002000000010200000000db263b4000000000000000000000000080c04bc380c04bc3000000800b010002000001010200000000db263b40000000000000000000000000a0c04bc3a0c04bc3000000800c010002000002010200000000db263b4000000000000000000000000000bd4bc300bd4bc30000008003010002000003010200000000db263b40000000000000000000000000e0ca4bc3e0ca4bc30000008014010002000004010200000000db263b40000000000000000000000000e0c24bc3e0c24bc3000000800e010002000005010200000000db263b4000000000000000000000000060bb4bc360bb4bc30000008002010002000006010200000000db263b40000000000000000000000000a0c64bc3a0c64bc30000008012010002000007010200000000db263b40000000000000000000000000c0c34bc3c0c34bc30000008010010002000008010200000000db263b40000000000000000000000000c0bd4bc3c0bd4bc30000008005010002000009010200000000db263b4000000000000000000000000060bd4bc360bd4bc3000000800401000200000a010200000000db263b40000000000000000000000000c0be4bc3c0be4bc3000000800701000200000b010200000000db263b40000000000000000000000000a0c74bc3a0c74bc3000000801301000200000c010200000000db263b4000000000000000000000000020c04bc320c04bc3000000800801000200000d010200000000db263b4000000000000000000000000040c04bc340c04bc3000000800901000200000e010200000000db263b40000000000000000000000000c0c44bc3c0c44bc3000000801101000200000f010200000000db263b4000000000000000000000000080c34bc380c34bc3000000800f010002000010010200000000db263b4000000000000000000000000060c04bc360c04bc3000000800a010002000011010200000000db263b4000000000000000000000000020c24bc320c24bc3000000800d010002000012010200000000db263b40000000000000000000000000801428c3801428c300000080010100020000130102")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 4 - Paricipants ###############################################################")
	arr, err1 = hex.DecodeString("e30701040104c655e025b6e12cb3dec2323d010000001314010d01051d532e2056455454454c000000000000000000000000000000000000000000000000000000000000000000000000000000000102050303442e2052494343494152444f00000000000000000000000000000000000000000000000000000000000000000000000000013b020a1c502e204741534c590000000000000000000000000000000000000000000000000000000000000000000000000000000000013a011035432e204c45434c4552430000000000000000000000000000000000000000000000000000000000000000000000000000000101061a44442e204b565941540000000000000000000000000000000000000000000000000000000000000000000000000000000000014b035840522e204b554249434100000000000000000000000000000000000000000000000000000000000000000000000000000000010f004d1b562e20424f5454415300000000000000000000000000000000000000000000000000000000000000000000000000000000010609071b4b2e2052c384494b4bc3964e454e0000000000000000000000000000000000000000000000000000000000000000000000013e061751412e20414c424f4e000000000000000000000000000000000000000000000000000000000000000000000000000000000001090221164d2e205645525354415050454e000000000000000000000000000000000000000000000000000000000000000000000000010b0714154b2e204d41474e555353454e00000000000000000000000000000000000000000000000000000000000000000000000000010e040b34532e20504552455a0000000000000000000000000000000000000000000000000000000000000000000000000000000000010a051b1d4e2e2048c39c4c4b454e424552470000000000000000000000000000000000000000000000000000000000000000000000010c07081c522e2047524f534a45414e0000000000000000000000000000000000000000000000000000000000000000000000000000013608040a4c2e204e4f5252495300000000000000000000000000000000000000000000000000000000000000000000000000000000011304120d4c2e205354524f4c4c000000000000000000000000000000000000000000000000000000000000000000000000000000000132033f0a472e2052555353454c4c000000000000000000000000000000000000000000000000000000000000000000000000000000010008374e432e205341494e5a0000000000000000000000000000000000000000000000000000000000000000000000000000000000014a096329412e2047494f56494e415a5a490000000000000000000000000000000000000000000000000000000000000000000000000007002c0a4c2e2048414d494c544f4e0000000000000000000000000000000000000000000000000000000000000000000000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 5 - Setup Data ###############################################################")
	arr, err1 = hex.DecodeString("e30701040105c655e025b6e12cb3dec2323d010000001306064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c04006064b4b000040c00000c0bfcecccc3d3433b33e0606060606064b3c0000b8410000ac41060000c040")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 6 - Car Telemtery ######################################################")
	arr, err1 = hex.DecodeString("e30701040106c655e025b6e12cb3dec2323d0100000013f8000000803f1c07803b000000000006142c00074500480050005100590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f6000000803f34bdc5bb000000000006b02b0019450047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f6000000803f743275bc000000000006aa2b00254500470050005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84100000000f8000000803fba9986bb000000000006212c000a4500480050005100590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f1000000803f70d580bc000000000006ec2a0003440047004f004f00590059005a005900590059005a005a005a000000ac410000ac410000b8410000b84101000100f7000000803f6d0439bc000000000006772b000d450047004f0050005a0059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f8000000803f929018bc000000000006be2a00004500480050005100590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f4000000803fc39853bc000000000006522b0000440047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84100000000f5000000803f1db349bc000000000006842b001d450047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84100000000f8000000803ffc503abc000000000006002c00324500480050005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f8000000803f37a54ebc000000000006222c000a4500480050005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f8000000803f30c008bc000000000006d42a00004500480050005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f2000000803f36fe4bbc0000000000060f2b000b440047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f7000000803f2ba165bc000000000006e32b00054500480050005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f5000000803f667245bc0000000000060d2c0027450047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84100000000f6000000803fb822a6bb000000000006762a0000440047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f7000000803f822f4ebc000000000006682b000e450047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84100000000f5000000803f2592b3bc000000000006092c0026450047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84100000000f6000000803fb609ecbb000000000006de2b0005450047004f005000590059005a0059005a0059005a005a005a000000ac410000ac410000b8410000b84101000100f4000000000000000080000000000006fa280000400043004b004b005a005a005a005a005a005a005a005a0059000000ac410000ac410000b8410000b8410000010000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

	fmt.Println(" ")
	fmt.Println("#######   Testing -  MSGTYPE 7 - Car Status  ###########################################################")
	arr, err1 = hex.DecodeString("e30701040107c655e025b6e12cb3dec2323d01000000130001033c009526c7400000dc428a177640bc34cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c009429c7400000dc424e1b7640bc34d70e0900000000001210000000000000000000000024744a000000000000000000000000000001033c003c2ac7400000dc42201c7640bc34d70e0900000000001210000000000000000000000024744a000000000000000000000000000001033c003f26c7400000dc421e177640bc34cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c005b2ac7400000dc42471c7640bc34ab0d0900000000001210000000000000000000000024744a000000000000000000000000000001033c008e24c7400000dc42fe1476402035cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c00d326c7400000dc42d81776402035cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c00e627c7400000dc4231197640bc34cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c003d28c7400000dc429f197640bc34ab0d0900000000001210000000000000000000000024744a000000000000000000000000000001033c00a828c7400000dc42251a7640bc34d70e0900000000001210000000000000000000000024744a000000000000000000000000000001033c00b025c7400000dc426b167640bc34cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c002125c7400000dc42b71576402035cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c00a82bc7400000dc42ea1d7640bc34d70e0900000000001210000000000000000000000024744a000000000000000000000000000001033c007f26c7400000dc426f177640bc34cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c00c228c7400000dc42461a7640bc34ab0d0900000000001210000000000000000000000024744a000000000000000000000000000001033c007923c7400000dc42a31376402035cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c008f24c7400000dc42001576402035cc100900000000001210000000000000000000000024744a000000000000000000000000000001033c00e728c7400000dc42741a7640bc34ab0d0900000000001210000000000000000000000024744a000000000000000000000000000001033c00b725c7400000dc4273167640bc34cc100900000000001210000000000000000000000024744a000000000000000000000000000201003c009a99c1400000dc42f31e6f402035cc100901000000001210000000000000000000000024744a00000000000000000000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}
	tc.SetInput("buffer", arr)
	buf.Reset()
	err = struc.Pack(&buf, arr)
	err = struc.Unpack(&buf, o)
	fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype = tc.GetOutput("msgtype")
	rdata = tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}
}
