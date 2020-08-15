package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBxBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBx, nil
}

func dictMod0IpaNeologdDictBx() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
