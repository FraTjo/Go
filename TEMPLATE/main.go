package main

import (
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	tasks := map[string][]Task{
		"tasks": {
			{"Learn Go", false},
		},
	}

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "todo", tasks) // using echo renderer
	})

	e.POST("/add-todo/", func(c echo.Context) error {
		todo := c.Request().PostFormValue("name")
		newTask := Task{todo, false}
		tasks["tasks"] = append(tasks["tasks"], newTask)

		return c.Render(200, "todoElement", newTask)
	})

	e.Start(":8080")
}
