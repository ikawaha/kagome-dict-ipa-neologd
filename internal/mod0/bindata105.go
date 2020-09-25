package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEb, nil
}

func dictMod0IpaNeologdDictEb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.eb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
