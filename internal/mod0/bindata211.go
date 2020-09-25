package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictId, nil
}

func dictMod0IpaNeologdDictId() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.id", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
