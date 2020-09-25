package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEk, nil
}

func dictMod0IpaNeologdDictEk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ek", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
