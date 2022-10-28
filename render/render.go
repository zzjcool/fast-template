package render

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func main2() {
	
	valueFilename:="example/value.yaml"
	templateFilename:="example/README.template"
	outputFilename:="example/README.md"

	// flag.StringVar(&valueFilename,"value","example/value.yaml","value file in yaml format")
	rootCmd:=&cobra.Command{
		Use:   "ft",
		Short: "Fast rendering with parameters and templates",
		Long: `Fast template is a tool for template rendering`,
	}
	rootCmd.PersistentFlags().StringVarP(&valueFilename, "value", "v", "example/value.yaml", "value file")

	// rootCmd.SetHelpCommand(&cobra.Command{
	// 	Use:   "ft",
	// 	Short: "Fast rendering with parameters and templates",
	// 	Long: `Fast template is a tool for template rendering`,
	// })
	rootCmd.Execute()
	



	BuildTemplate[Config](valueFilename, templateFilename, outputFilename)
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