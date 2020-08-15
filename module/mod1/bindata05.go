package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJu, nil
}

func dictMod1IpaNeologdDictJu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ju", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
