package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGtBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGt, nil
}

func dictMod0IpaNeologdDictGt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
