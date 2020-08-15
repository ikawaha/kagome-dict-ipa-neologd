package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictItBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIt, nil
}

func dictMod0IpaNeologdDictIt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictItBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.it", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
