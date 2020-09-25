package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAl, nil
}

func dictMod0IpaNeologdDictAl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.al", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
