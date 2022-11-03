package render

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"gopkg.in/yaml.v2"
)

// Render 使用字符串进行输入输出
func Render(config, tpl string) string {

	t := template.New("test").Funcs(sprig.TxtFuncMap())
	t = template.Must(t.Parse(
		`hello {{.UserName}}!`))
	p := map[string]any{
		"UserName": "ddd",
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, p)
	return buf.String()
}

func BuildTemplate[T any](configFile, templateFile, outputFile string) {

	// 读取配置

	var conf T //定义一个结构体变量

	//读取yaml文件到缓存中
	config, err := os.ReadFile(configFile)
	errPanic(err)

	//yaml文件内容影射到结构体中
	err = yaml.Unmarshal(config, &conf)
	errPanic(err)

	tpl := template.Must(
		template.New(filepath.Base(templateFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(templateFile),
	)

	var filename = outputFile
	var f *os.File
	if checkFileIsExist(filename) { //如果文件存在
		err := os.Remove(filename)
		errPanic(err)
	}
	f, err = os.Create(filename) //创建文件
	errPanic(err)
	defer f.Close()
	err = tpl.Execute(f, conf)
	errPanic(err)
}

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

type Config map[string]any
