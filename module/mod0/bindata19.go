package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAtBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAt, nil
}

func dictMod0IpaNeologdDictAt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.at", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
