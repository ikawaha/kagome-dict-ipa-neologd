package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPu, nil
}

func dictMod1IpaNeologdDictPu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
