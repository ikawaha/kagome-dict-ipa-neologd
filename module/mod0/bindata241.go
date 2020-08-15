package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJh, nil
}

func dictMod0IpaNeologdDictJh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
