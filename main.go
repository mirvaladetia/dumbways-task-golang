package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	e.Static("/public", "public")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog_list/:id", blogList)
	e.POST("/add-blog", addBlog)
	e.GET("/testimonials", testimonials)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":    id,
		"Title": "Putri Es Inazuma, Kamisato Ayaka",
		"Content": `Waifu cryo dari Inazuma ini merupakan salah satu karkater yang disukai oleh banyak player Genshin.
					Selain menjadi DPS yang kuat, dia juga memiliki kepribadian yang menarik.`,
	}

	var tmpl, err = template.ParseFiles("views/blog_list.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addBlog(c echo.Context) error {
	judul := c.FormValue("judul")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	konten := c.FormValue("konten")
	maindps := c.FormValue("maindps")
	subdps := c.FormValue("subdps")
	shielder := c.FormValue("shielder")
	healer := c.FormValue("healer")

	println("Judul : " + judul)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Konten : " + konten)
	println("Role : " + maindps)
	println("Role : " + subdps)
	println("Role : " + shielder)
	println("Role : " + healer)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
