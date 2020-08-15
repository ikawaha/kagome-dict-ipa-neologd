package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGb, nil
}

func dictMod0IpaNeologdDictGb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
