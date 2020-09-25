package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIb, nil
}

func dictMod0IpaNeologdDictIb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ib", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
