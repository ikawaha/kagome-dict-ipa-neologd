package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBvBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBv, nil
}

func dictMod0IpaNeologdDictBv() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
