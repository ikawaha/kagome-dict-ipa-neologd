package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDb, nil
}

func dictMod0IpaNeologdDictDb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.db", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
