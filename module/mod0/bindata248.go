package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJo, nil
}

func dictMod0IpaNeologdDictJo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
