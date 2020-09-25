package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNmBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNm, nil
}

func dictMod1IpaNeologdDictNm() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
