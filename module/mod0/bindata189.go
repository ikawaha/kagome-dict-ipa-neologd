package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHh, nil
}

func dictMod0IpaNeologdDictHh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
