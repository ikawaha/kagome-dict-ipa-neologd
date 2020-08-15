package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKgBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKg, nil
}

func dictMod1IpaNeologdDictKg() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
