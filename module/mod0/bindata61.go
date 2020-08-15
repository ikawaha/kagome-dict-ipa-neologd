package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCj, nil
}

func dictMod0IpaNeologdDictCj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
