package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJa, nil
}

func dictMod0IpaNeologdDictJa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ja", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
