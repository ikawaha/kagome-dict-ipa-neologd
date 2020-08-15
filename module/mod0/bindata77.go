package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCzBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCz, nil
}

func dictMod0IpaNeologdDictCz() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
