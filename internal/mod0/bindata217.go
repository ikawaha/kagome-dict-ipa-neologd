package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIj, nil
}

func dictMod0IpaNeologdDictIj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ij", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
