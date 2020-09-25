package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDg, nil
}

func dictMod0IpaNeologdDictDg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
