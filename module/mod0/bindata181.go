package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGzBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGz, nil
}

func dictMod0IpaNeologdDictGz() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
