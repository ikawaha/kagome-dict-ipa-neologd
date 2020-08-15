package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGi, nil
}

func dictMod0IpaNeologdDictGi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
