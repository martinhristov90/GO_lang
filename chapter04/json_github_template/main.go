package main


import (
	"github" // using local /Users/marti/go/src/github
	"log"
	"os"
	"text/template"
	"time"
)

const templ = ` Total count of the issues : {{ .TotalCount }}
{{ range .Items }}-------------------------------------------------------
Number: {{ .Number }}
User: {{ .User.Login }}
Title: {{ .Title | printf "%.64s" }}
Age: {{ .CreatedAt | daysAgo }} days
{{ end }}`

const templ_html = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {

	//templ_txt()

	templ_html_report()
}

func templ_txt() {

	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	// top to bottom left to right : creating new template named "issuelist"
	// argumenting it with Funcs, which uses FuncMap to relate names inside the template with functions inside the code
	// finally parsing with Parse(templ)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err = report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func templ_html_report() {
	var template_html = template.Must(template.New("issuelist_html").Parse(templ_html))

	result, err := github.SearchIssues(os.Args[1:]) // 1st of os.Args is the name of the program, taking care of the rest
	if err != nil {
		log.Fatal(err)
	}

	if err = template_html.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
