package gutilz

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"log"
)

// CreateFolder - create or force re-create a folder
func CreateFolder(folder string, forceCreate int) error {

	folderInfo, err := os.Stat(folder)
	if os.IsNotExist(err) {
		log.Printf("%s\n", err)
	} else {
		log.Printf("%s already exists\n", folderInfo.Name())
		if forceCreate == 1 {
			log.Printf("Deleting %s\n", folder)
			os.RemoveAll(folder)
		} else {
			return nil
		}
	}

	log.Printf("Creating = %+v\n", folder)
	errDir := os.MkdirAll(folder, 0755)

	if errDir != nil {
		log.Println(errDir)
	} else {
		log.Printf("created = %+v\n", folder)
	}
	return errDir
}

// Dumper - similar to Data::Dumper module in Perl.
func Dumper(data ...interface{}) (string, error) {
	prefix := ""
	indent := "  "
	if len(data) == 2 {
		if fmt.Sprintf("%s", reflect.TypeOf(data[1]).Kind()) == "map" {
			params := data[1].(map[string]string)
			prefix = params["prefix"]
			indent = params["indent"]
		}
	}
	b, err := json.MarshalIndent(data[0], prefix, indent)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// CreateSymLink creates a sym link
func CreateSymLink(src string, p string, forceCreate int) error {
	if _, err := os.Lstat(p); err == nil {
		log.Printf("Symlink already exists. = %+v\n", p)

		if forceCreate == 1 {
			log.Printf("Removing symlink...")
			os.Remove(p)
		} else {
			s, _ := os.Readlink(p)

			log.Printf("Symlink %s points to %s", p, s)
			if s != src {
				log.Printf("WARN: Symlink destination is NOT pointing to %s", p)
			}
			return nil
		}
	}

	err := os.Symlink(src, p)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Symlink created %s --> %s", p, src)
	return nil
}
