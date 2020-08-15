package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIrBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIr, nil
}

func dictMod0IpaNeologdDictIr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ir", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
