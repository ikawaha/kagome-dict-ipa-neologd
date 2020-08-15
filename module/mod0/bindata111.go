package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictEhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEh, nil
}

func dictMod0IpaNeologdDictEh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictEhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.eh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
