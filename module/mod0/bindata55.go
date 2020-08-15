package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCd, nil
}

func dictMod0IpaNeologdDictCd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
