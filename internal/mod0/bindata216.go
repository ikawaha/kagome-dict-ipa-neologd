package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIi, nil
}

func dictMod0IpaNeologdDictIi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ii", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
