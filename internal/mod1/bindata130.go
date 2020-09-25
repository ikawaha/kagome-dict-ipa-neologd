package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOpBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOp, nil
}

func dictMod1IpaNeologdDictOp() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.op", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
