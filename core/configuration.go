package core

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path"
)

const (
	HOME_DIRECTORY_CONFIG = "my home dir"
	BLOCKCHAIN_DIRECTORY  = ".blockchain/"

	BLOCKHAIN_KEYS_FILENAME = "keys.json"
)

func getDirectoryWithBaseDir(dir string) string {

	if dir == HOME_DIRECTORY_CONFIG {

		usr, err := user.Current()
		logOnError(err)
		dir = usr.HomeDir
	}

	return path.Join(dir, BLOCKCHAIN_DIRECTORY)
}

func OpenConfiguration(dir string) (*Keypair, error) {

	dir = getDirectoryWithBaseDir(dir)

	//Create ~/.blockchain directory
	err := os.MkdirAll(dir, 0777)
	logOnError(err)

	//Search for keys file inside directory
	f, err := os.OpenFile(path.Join(dir, BLOCKHAIN_KEYS_FILENAME), os.O_RDWR|os.O_CREATE, 0660)
	defer f.Close()

	if err != nil {

		return nil, err
	}

	k := &Keypair{}
	json.NewDecoder(f).Decode(k)

	if k == nil || k.Public == nil || k.Private == nil {
		return nil, nil
	}

	return k, f.Close()
}

func WriteConfiguration(dir string, keypair *Keypair) error {

	if keypair != nil {

		dir = getDirectoryWithBaseDir(dir)

		err := os.MkdirAll(dir, 0777)
		logOnError(err)

		f, err := os.OpenFile(path.Join(dir, BLOCKHAIN_KEYS_FILENAME), os.O_RDWR|os.O_CREATE, 0660)

		err = json.NewEncoder(f).Encode(keypair)
		logOnError(err)

		return f.Close()
	}

	return errors.New("No keypair provided to save")
}
