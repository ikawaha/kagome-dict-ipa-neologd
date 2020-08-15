package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictClBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCl, nil
}

func dictMod0IpaNeologdDictCl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictClBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
