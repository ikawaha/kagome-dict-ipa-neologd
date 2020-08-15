package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOmBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOm, nil
}

func dictMod1IpaNeologdDictOm() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.om", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
