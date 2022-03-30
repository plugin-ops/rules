package main

import (
	"action"
	"fmt"
	"path"
	"rule"
	"strings"
	"value"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic:", r)
		}
	}()
	path := value.Path.String()
	if path == "" {
		return
	}
	if err := AutoRemove(path); err != nil {
		rule.Error(err)
		return
	}
}

func AutoRemove(p string) error {
	fmt.Println("p:", p)
	if p == "" {
		return nil
	}
	stats, err := action.Info(p)
	if err != nil {
		return err
	}
	if !value.NewValue(stats["4"]).Bool() {
		if value.NewValue(stats["5"]).Float64() > value.Size.Float64() {
			fmt.Println("remove:", p)
			_, err = action.File("remove", p)
			return err
		}
		return nil
	}
	files := strings.Split(stats["7"].String(), "\n")
	for _, file := range files {
		err = AutoRemove(path.Join(p, file))
		if err != nil {
			return err
		}
	}
	return nil
}

//--config--

var rule_name = "remove big file"
var rule_version = 1.0
var rule_description = "脚本将自动删除指定目录中所有空间占用大于指定大小的文件, 不会删除文件夹"
var rule_rely_on = map[string]string{
	"File": "file",
	"Info": "path-info@0.1",
}
var rule_params = map[string]string{
	"Path": "",
	"Size": "0",
}
