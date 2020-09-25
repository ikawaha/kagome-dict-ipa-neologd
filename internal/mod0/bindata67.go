package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCpBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCp, nil
}

func dictMod0IpaNeologdDictCp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
