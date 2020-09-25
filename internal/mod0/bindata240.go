package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJg, nil
}

func dictMod0IpaNeologdDictJg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
