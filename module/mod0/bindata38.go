package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBmBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBm, nil
}

func dictMod0IpaNeologdDictBm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
