package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEs, nil
}

func dictMod0IpaNeologdDictEs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.es", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
