package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIxBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIx, nil
}

func dictMod0IpaNeologdDictIx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ix", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
