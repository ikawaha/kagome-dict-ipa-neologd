package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictElBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictEl, nil
}

func dictMod0IpaNeologdDictEl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictElBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.el", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
