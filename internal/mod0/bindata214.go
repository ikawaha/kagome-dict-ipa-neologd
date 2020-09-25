package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIg, nil
}

func dictMod0IpaNeologdDictIg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ig", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
