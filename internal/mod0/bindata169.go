package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGnBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGn, nil
}

func dictMod0IpaNeologdDictGn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
