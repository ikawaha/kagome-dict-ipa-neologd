package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMoBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMo, nil
}

func dictMod1IpaNeologdDictMo() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
