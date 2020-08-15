package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEj, nil
}

func dictMod0IpaNeologdDictEj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ej", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
