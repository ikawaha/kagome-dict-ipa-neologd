package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIl, nil
}

func dictMod0IpaNeologdDictIl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.il", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
