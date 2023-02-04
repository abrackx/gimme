package template

import (
	"gimme/database"
	"os"
	"text/template"
)

func getSpringTemplate() []byte {
	return []byte(`
spring:
  datasource:
    driver-class-name: org.postgresql.Driver
    url: jdbc:postgresql://localhost:{{ .Port }}/postgres
    username: {{ .Username }}
    password: {{ .Password }}
`)
}

func PrintSpringTemplate(db database.Database) {
	spring := template.Must(template.New("spring").Parse(string(getSpringTemplate())))
	err := spring.Execute(os.Stdout, db)
	if err != nil {
		panic(err)
	}
}
