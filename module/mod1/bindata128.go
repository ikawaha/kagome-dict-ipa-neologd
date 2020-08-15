package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOnBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOn, nil
}

func dictMod1IpaNeologdDictOn() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.on", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
