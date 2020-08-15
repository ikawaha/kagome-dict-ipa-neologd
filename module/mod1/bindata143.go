package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPcBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPc, nil
}

func dictMod1IpaNeologdDictPc() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
