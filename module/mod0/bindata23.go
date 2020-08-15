package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAxBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAx, nil
}

func dictMod0IpaNeologdDictAx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ax", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
