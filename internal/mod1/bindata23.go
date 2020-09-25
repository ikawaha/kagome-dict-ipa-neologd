package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKmBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKm, nil
}

func dictMod1IpaNeologdDictKm() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.km", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
