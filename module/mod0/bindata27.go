package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBb, nil
}

func dictMod0IpaNeologdDictBb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
