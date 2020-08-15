package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJe, nil
}

func dictMod0IpaNeologdDictJe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.je", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
