package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHvBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHv, nil
}

func dictMod0IpaNeologdDictHv() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
