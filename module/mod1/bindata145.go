package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPeBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPe, nil
}

func dictMod1IpaNeologdDictPe() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pe", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
