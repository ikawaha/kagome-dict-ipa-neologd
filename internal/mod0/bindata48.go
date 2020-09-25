package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBwBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBw, nil
}

func dictMod0IpaNeologdDictBw() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
