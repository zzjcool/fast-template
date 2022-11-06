package render

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"gopkg.in/yaml.v2"
)

// Str 使用字符串进行输入输出
func Str(config, tpl string) (string,error) {

	t := template.New("template").Funcs(sprig.TxtFuncMap())
	t,err := t.Parse(tpl)
	if err!=nil{
		return "",err
	}
	conf := new(map[string]any)
	//yaml文件内容影射到结构体中
	err = yaml.Unmarshal([]byte(config), conf)
	if err!=nil{
		return "",err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, conf);err!=nil{
		return "",err
	}
	return buf.String(),nil
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
