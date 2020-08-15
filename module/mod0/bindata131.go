package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFb, nil
}

func dictMod0IpaNeologdDictFb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
