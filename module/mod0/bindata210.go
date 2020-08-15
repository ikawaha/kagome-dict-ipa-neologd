package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIc, nil
}

func dictMod0IpaNeologdDictIc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ic", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
