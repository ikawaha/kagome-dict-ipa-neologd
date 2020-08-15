package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAk, nil
}

func dictMod0IpaNeologdDictAk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ak", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
