package documentmanager

import (
	"encoding/json"
	"fmt"
	"github.com/adaptive-scale/dockyard/asset"
	"github.com/adaptive-scale/dockyard/internal/configuration"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

const OutputDir = "public"

type DocumentManager struct {
	config *configuration.Configuration
}

func New(config *configuration.Configuration) *DocumentManager {
	return &DocumentManager{config:config}
}

func (d *DocumentManager) ListAllFiles() map[string]map[string]string {
	return render(d.config.Location, OutputDir)
}

func (d *DocumentManager) Reset() {
	err := os.RemoveAll(OutputDir)
	if err != nil {
		fmt.Println("error while removing outdir - "+err.Error())
		os.Exit(1)
	}
}

func (d *DocumentManager) Generate() {
	d.Reset()

	for r, f := range d.ListAllFiles() {
		js, menu := d.GenerateJS(f)
		tmpl := d.GenerateIndexHTML(menu, js)
		mkdir(r)
		write(r, tmpl)
	}
}

func write(r string, tmpl string) {
	err := ioutil.WriteFile(path.Join(r, "index.html"), []byte(tmpl), os.FileMode(0777))
	if err != nil {
		fmt.Println("Error while writing -"+ err.Error())
		os.Exit(1)
	}
}

func mkdir( r string) {
	err := os.MkdirAll(r, os.FileMode(0700))
	if err != nil {
		fmt.Println("Error - "+err.Error())
		os.Exit(1)
	}
}

type FileLocation struct {
	FileLocation string
	Render map[string]string
}

func render(location, root string) map[string]map[string]string {

	fileLoc := map[string]map[string]string{}
	renderedJS := map[string]string{}

	var currentPath string
	filepath.Walk(location,
		func(filePath string, info os.FileInfo, err error) error {

			if currentPath != path.Join(root, filepath.Dir(filePath)) {
				renderedJS = map[string]string{}
			}
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if filepath.Ext(filePath) == ".md"{
					md, _ := ioutil.ReadFile(filePath)
					output := markdown.ToHTML(md, nil, nil)
					renderedJS[strings.ReplaceAll(info.Name(), filepath.Ext(info.Name()), "")] = string(output)
					join := path.Join(root, filepath.Dir(filePath))
					currentPath = join
					fileLoc[join] = renderedJS
				}
			}

			return nil
		})

	return fileLoc
}

type Content struct {
	index string
	submenu []*Content
	id string
	content string
}

func (d *DocumentManager) GetMenu(menuAndContent map[string]string) (string,string, string) {
	fmt.Println("==> generating menu")

	start := "{"
	end := "}"


	var contents []*Content
	var menu string
	var activated string
	for key, value := range menuAndContent {
		if string(key[0]) == "_" {
			continue
		} else {
			key := strings.Split(key, "_")

			if len(key) >= 2 {

				var content Content

				content.index = key[0]
				content.content = value
				content.id = strings.Join(key[1:], "_")

				contents = append(contents, &content)
			}
		}
	}

	for key, _ := range menuAndContent {
		if string(key[0]) == "_" {
			key := strings.Split(key[1:], "_")
			if len(key) >= 2 {
				for _, c := range contents {
					if c.index == key[0] {
						c.submenu = append(c.submenu, &Content{id: strings.Join(key[2:], "_"), index:key[1]})
					}
				}
			}
		}
	}

	sort.Slice(contents, func(i, j int) bool { return contents[i].index < contents[j].index })


	for _, content := range contents {
		log.Println(content.submenu)

		if activated == "" {
			activated = content.id
		}

		sort.Slice(content.submenu, func(i, j int) bool { return content.submenu[i].index < content.submenu[j].index })

		d, _ := json.Marshal(content.submenu)

		var data string

		if len(content.submenu) > 0 {
			for _, s := range content.submenu  {
				data = data+`<li onclick="onsection(this)" id="` + content.id + `-`+s.id+`" style="border: 0px; padding: 10px; color: #888; margin-left: 23px; font-size: 14px; border-radius:6px;cursor: pointer;"  class="list list-group-item">`+strings.ReplaceAll(s.id, "_", " ")+ `</li>`
			}
		}
		start = start + "'"+content.id+"':{'id':'" + content.id + "', 'content':" + "`" + content.content + "`, 'submenu':"+fmt.Sprintf("%v", string(d))+"},"
		menu = menu + `<li id="` + content.id + `" style="border: 0px; border-radius:6px;cursor: pointer;" class="list list-group-item" onclick="activate(this);">` + beautify(content.id) + data+ `
</li>
`
	}

	 start = start[:len(start)-1] + end

	 return start, menu, activated
}

func (d *DocumentManager) GenerateJS(menuAndContent map[string]string) (string, string) {

	start, menu, activated := d.GetMenu(menuAndContent)

	fmt.Println("==> generating content")

	return `var activated = '`+activated+`';
   let url = window.location.href;
    url = url.split('#')[1];
    if(url != null){
        activated = url;
    }
var advent = {id: activated};
var content=` + start, menu

}

func (d *DocumentManager) GenerateIndexHTML( menu, js string) string {

	fmt.Println("==> generating template")

	css, err := asset.Asset("templates/style.css")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ""
	}

	newTemplate, err := asset.Asset("templates/index_template.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ""
	}

	theme := "templates/theme/" + d.config.Theme + ".css"
	themeInfo, err := asset.Asset(theme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ""
	}

	tmpl := strings.ReplaceAll(string(newTemplate), "<MENU />", menu)
	tmpl = strings.ReplaceAll(tmpl, "<CSS />", string(css))
	tmpl = strings.ReplaceAll(tmpl, "<JS />", js)
	tmpl = strings.ReplaceAll(tmpl, "<Brand />", d.config.Branding)
	tmpl = strings.ReplaceAll(tmpl, "<Theme />", string(themeInfo))

	fmt.Println("==> generating index.html length=", len(tmpl))

	return tmpl
}

func beautify(key string) string {
	return strings.Title(strings.ReplaceAll(key, "_", " "))
}