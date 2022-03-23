rule_name := "remove big file"
rule_version := 1.0
rule_description := "脚本将自动删除指定目录中所有空间占用大于指定大小的文件, 不会删除文件夹"
rule_rely_on := map[string]string{
"Info":"path-info",
"File":"file",
}
rule_params := map[string]string{
"Path":"",
"Size":0,
}

//--body--

import (
	"action"
	"fmt"
	"strings"
	"value"
)

func main() {
	path := value.Path.String()
	if path == "" {
		return
	}

	if err := AutoRemove(path); err != nil {
		fmt.Println("err:", err)
		return
	}
}

func AutoRemove(path string) error {
	if path == "" {
		return nil
	}
	stats, err := action.Info(value.Path)
	if err != nil {
		return err
	}
	if !value.NewValue(stats[4]).Bool() {
		if value.NewValue(stats[5]).Float64 > value.Size.Float64 {
			_, err = action.File("remove")
			return err
		}
		return nil
	}
	files := strings.Split(stats[7], "\n")
	for _, file := range files {
		err = AutoRemove(file)
		if err != nil {
			return err
		}
	}
	return nil
}
