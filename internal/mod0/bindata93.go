package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDpBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDp, nil
}

func dictMod0IpaNeologdDictDp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
