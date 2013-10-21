package mosquito


import (
	"io"
	"regexp"
	"io/ioutil"
	"html/template"
	"path/filepath"
)


type FileSystemRenderer struct {
	Root	string
}

func (renderer *FileSystemRenderer) Render(w io.Writer, file string, data interface{}) error {
	var tmpl *template.Template

	if err := renderer.parse(file, tmpl); err != nil {
		return err
	}

	tmpl.ExecuteTemplate(w, file, data)

	return nil
}

func (renderer *FileSystemRenderer) parse(file string, output *template.Template) error {
	path := filepath.Join(renderer.Root, file)
    content, err := ioutil.ReadFile(path)

	if err != nil {
	    return err
	}

	if _, err := output.New(file).Parse(string(content)); err != nil {
	    return err
	}

	reg := regexp.MustCompile("{{[ ]*template[ ]+\"([^\"]+)\"")
	matches := reg.FindAllStringSubmatch(string(content), -1)

	for _, match := range matches {
		if t := output.Lookup(match[1]); t != nil {
            continue
		}
		if err := renderer.parse(match[1], output); err != nil {
		    return err
		}
	}
	return nil
}
