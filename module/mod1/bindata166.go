package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPz, nil
}

func dictMod1IpaNeologdDictPz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
