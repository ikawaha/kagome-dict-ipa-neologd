package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIa, nil
}

func dictMod0IpaNeologdDictIa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ia", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
