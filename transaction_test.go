/*
MIT License

Copyright (c) 2016 Sascha Hanse
Copyright (c) 2017 Shinya Yagyu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package gcornode

import (
	"encoding/json"
	"testing"
	"time"
)

// XXX: come up with more test cases that could ideally be
// 			shared accross implementations.

func TestNewTransactionFromTrits(t *testing.T) {
	var trytes Trytes = "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999A9RGRKVGWMWMKOLVMDFWJUHNUNYWZTJADGGPZGXNLERLXYWJE9WQHWWBMCPZMVVMJUMWWBLZLNMLDCGDJ999999999999999999999999999999999999999999999999999999YGYQIVD99999999999999999999TXEFLKNPJRBYZPORHZU9CEMFIFVVQBUSTDGSJCZMBTZCDTTJVUFPTCCVHHORPMGCURKTH9VGJIXUQJVHK999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	//var hash Trytes = "IPQYUNLDGKCLJVEJGVVISSQYVDJJWOXCW9RZXIDFKMBXDVZDXFBZNZJKBSTIMBKAXHFTGETEIPTZGNTJK"
	out := Transaction{
		SignatureMessageFragment: "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999",
		Address:                  "A9RGRKVGWMWMKOLVMDFWJUHNUNYWZTJADGGPZGXNLERLXYWJE9WQHWWBMCPZMVVMJUMWWBLZLNMLDCGDJ",
		Value:                    0,
		Tag:                      "999999999999999999999999999",
		Timestamp:                time.Unix(1482522289, 0),
		CurrentIndex:             0,
		LastIndex:                0,
		Bundle:                   "TXEFLKNPJRBYZPORHZU9CEMFIFVVQBUSTDGSJCZMBTZCDTTJVUFPTCCVHHORPMGCURKTH9VGJIXUQJVHK",
		Nonce:                    "999999999999999999999999999999999999999999999999999999999999999999999999999999999",
		TrunkTransaction:         "999999999999999999999999999999999999999999999999999999999999999999999999999999999",
		BranchTransaction:        "999999999999999999999999999999999999999999999999999999999999999999999999999999999",
	}
	tt, err := NewTransaction(trytes)
	if err != nil {
		t.Fatalf("TransactionFromTrits() expected err to be %#v but got %#v", nil, err)
	}
	if out.Trytes() != tt.Trytes() {
		t.Errorf("TransactionFromTrits()\nwant: %#v\nhave: %#v", out, tt)
	}
	js, err := json.Marshal(tt)
	if err != nil {
		t.Error(err)
	}
	if trytes != Trytes(js[1:len(js)-1]) {
		t.Error("transaction marshaller is incorrect.")
	}
	var tt2 Transaction
	if err := json.Unmarshal([]byte(`"`+trytes+`"`), &tt2); err != nil {
		t.Error(err)
	}
	if tt.Trytes() != tt2.Trytes() {
		t.Error("transaction unmarshaller is incorrect.")
	}

	//t.Logf("tt: %#v\n", tt)
}

func TestNonce(t *testing.T) {
	var trytes Trytes = "QQ9VLGQTJICJH9UZNSBXRSKBWVXCTRNGLJPEYFSZBISDCGBHGAV9TEHIMJS9LMQNHFNWVAXALKGESPWZMUMZPQCKCUFSEJZVHWZHZOEZBPTERXPHUQLQLZMENKOKAWMJ9LCFSIBSBEYCJQVQQMTMRJMDEKRXLCIRZFWQBRJYLPSK9XLWFYFTU9FBJATWPQRJBHWYJRUEXSXMBJLVWNYRTZJTHLEKDTWCGJ9OXDEZNWTKLTXXTKVFXDMRJUDAMDACRHJKZIJFJBZRVDLSTIMOWPTLLIVHCFUBMSQIPVPCSQAPZGHHNNQLWEHDTIQZQTAXJMTTROYOTIZSZKKQFXBHXFKSNAGZWHWEGXCKHSJQGBJGC9IKUSSJEAOHTPGPMIYPEXJXJKRLX9IOCUDKK9ONEUONYVGHRHXSAUZJXGVQNWSDZIYXXZMMJKRTOXSWJHBIMXPEMJTKIBSQKPICTKAQLJVOLZSGNVVBIOFJJKKJ9UBYKJWCBRPQLGBNEKEFRYCHB9PORAKEKTJZYZBFGZLONUJCPNINGUTYWTMDKTEPDASCHNVHXTUIOZ9PDALXKQYLILMUJEJWTYYAGEQXNNPLOMFGHSTFCNRMDFUSQREHFASDXZZOYNWNVKCRFZIDF9X9YKME9O9NJH9LFVNNSOXHSQOMBULFHBBPSGRCCGKWENQOZHRSIIHKKXAQTFISNVGIVMBDKSJYDW9VTZBAVYWWJAIYSCHGBIXKMKHBRTHWQMIWQVPWFUJQVUDZRDX9MMCXOQTQZGKECGIHATHUC9TDGKUKOYXEEBGTQYEO9KADOSMYWGQARTIUG9IUHQEBWSFHWDZCRCH9WDRJEUSQDBBGGKTOYBKYZX9LUKUGBVTEPXDOVVNSKPUWNARSILQNOQKATCUHMRVMKVHF9B9TEJDOMBDXLNKDJI9IIYRXNOQPBOVOEEGQFSZJJOFPNFVXOYZNLSOOA9FWDFJKDOQUHWHBNDWQZZVRBZJSLDVBRCGFCWXNUFTMLCHNXODSQMUIBPN9NMAXZXKUYYRUEDSLW9UEQYGSEBEOHGI9W9WAUFDEQCXVYOTBAAHFXHRGJBWGFZKIUMSWEXAD9EDALOOYQZIXKUOWJDHBHSEUDPJJRXANPPHUZWJM9KDVGJUAQHFJTWNFSULGSWMGMAOCAHTIYYJONOLKGNKRXWWZYWKITSFPQJNHQWAIJULYWPXQENPZTWZISXJBWYLENOINLFBRWWGWSWJONHIA99VGCSAT9XNMZUCSVEJMZJASUSXVUWFSVFXUIFWGIFKLEFLANHTITFCOPXRXTDXRDHLWP9RJHVQJMLPBFVHTIOQMSZECNPITTBTPLUHDQQZP9BWTWIZSOTMTZQWHCETRTKQPOHPJMJMLSBWC9ZRQAZDLAPJHAXRQZ9RUHHBPASWVUHBYX9FH9PLEHQEKCOMIHUNDVKOPKXEFM9CNP9LOLMBVHMMWNDCKHSYZEXOJAQOHMRXEGWMWFW9YVOZ9YDPHWVTKWFYUECSROYJENFPSIKHBJRSNWO9KQUENGPVULVYAQFAIYFSNIYR9LLRMYNCNQQAOPFNMWFSWSPMWNNIRNVKDZWRLGMPCBOVIMMXEZUFYPFPIGWTGOLDAEBZQADSOGVDZXHEWZGHNAAHMFPSOZD9SEPNCW9GTN9WLFDDCKMMXPXAEUUJPJFKKWMGUKVMYBH9AIEFAIIDJOAWRDKECA99XOYRSFZQKTVRJMTUAQJZUTKGXROESUYYSHTLRIZRPSNDFEEZWXCQONZYCD9TOHCBP9ISXQ9YSRZJ999Z9TETYOINSDGRBQSDTVFABHQNLTWGFYLHBHPVKLIBUMLVSURAOS9QHXDTIPKOJDLYOKRCEKCBMKVYIAKVA9WTGDWHIRUAWOVRKOSYTNIZAZNTJRFJDMNLGHTDKPKZDLBPQXRIRIVREMOBCPHMBBAUKNXHU9XIZNG9GD9LDIBBFPSI9PJNRCHXHNWAZXIACE9LUBNUWOK9LGJ9MKZQRI9CBCJUNALQKKVGGSPRJFAGCXNFO99YMLMKI9NVUZCZ9BCUEBSGMAVNKGWYWWQPZISMKAROXLQWEHOJIJOIIYRUDBNHRD9DEDQWQONAXKKSYMYCFTITZFKIXKZCGAVAFQIYEMESOIMWUUDSXJRR9RVWTAAHCOA9SCQBF9LAGPPYDXPEBKLHZ9KHKTXFP9XOVMVWIXEWMOISJHMQEXMYMZCUGEQNKGUNVRPUDPRX9IR9LBASIARWNFXXESPITSLYAQMLCLVTLHW9999999999999999999999999999999999999999999999999999FBIEUWD99A99999999C99999999DEXRPLKGBROUQMKCLMRPG9HFKCACDZ9AB9HOJQWERTYWERJNOYLW9PKLOGDUPC9DLGSUH9UHSKJOASJRU9MMRRSLICRITOROFC9FBVWLFEDNN9KJKYHUMRCJEUDGCYCWTBP9HHBEEJRFAU9FALRJWTU99NZK999999UE9VSBDVSRNTBZWPXYZPGAUTSWFLARLPXMHYBSTEUWIDOFJQJMVIACGUPTOMBWQO9AEADCFCMFJ999999WQKHJEXIHMOKQETOUTEO9JUPCDNAJQYZVXQRCXGYGEBOTMHE9HSJXVYVQUS9FPDLQWWKSYVDPCXX9LLAT"
	tx, err := NewTransaction(trytes)
	if err != nil {
		t.Fatal(err)
	}
	if !tx.HasValidNonce(18) {
		t.Error("cannot validate nonce")
	}
}
