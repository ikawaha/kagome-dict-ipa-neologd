package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEq, nil
}

func dictMod0IpaNeologdDictEq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.eq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
