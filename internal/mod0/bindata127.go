package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictExBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEx, nil
}

func dictMod0IpaNeologdDictEx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictExBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ex", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
