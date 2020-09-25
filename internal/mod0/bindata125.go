package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEvBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEv, nil
}

func dictMod0IpaNeologdDictEv() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ev", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
