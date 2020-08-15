package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNeBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNe, nil
}

func dictMod1IpaNeologdDictNe() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ne", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
