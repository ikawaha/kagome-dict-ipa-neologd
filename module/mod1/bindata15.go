package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKeBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKe, nil
}

func dictMod1IpaNeologdDictKe() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ke", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
