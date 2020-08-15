package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFk, nil
}

func dictMod0IpaNeologdDictFk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
