package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGo, nil
}

func dictMod0IpaNeologdDictGo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.go", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
