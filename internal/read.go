package internal

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadPath(path string) (bool, string, error) {
	path = strings.TrimLeft(path, "/")
	if !strings.HasPrefix(path, "./") {
		path = "./" + path
	}
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true, BuildTemplate(Html404Template, map[string]interface{}{"Path": path}), nil
		}
		return false, "", err
	}

	if fileinfo.IsDir() {
		content, err := readDir(path)
		if err != nil {
			return false, "", err
		}
		return true, content, nil
	}

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return false, "", err
	}

	return false, string(bs), err
}

func readDir(path string) (string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	res := []string{}
	if path != "./" {
		res = append(res, "../")
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			res = append(res, dir.Name())
		} else {
			res = append(res, dir.Name()+"/")
		}
	}

	return BuildTemplate(htmlDirTemplate, map[string]interface{}{
		"Path": path,
		"Dirs": res,
	}), nil
}