package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEg, nil
}

func dictMod0IpaNeologdDictEg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.eg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
