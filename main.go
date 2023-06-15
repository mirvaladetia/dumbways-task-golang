package main

import (
	"belajargolang/connection"
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id             int
	Image          string
	Judul          string
	StartDate      time.Time
	EndDate        time.Time
	FormattedStart string
	FormattedEnd   string
	Durasi         string
	Konten         string
	Maindps        bool
	Subdps         bool
	Shielder       bool
	Healer         bool
}
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	e.Static("/public", "public")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog_list/:id", blogList) //membuka halaman detail blog dari klik judul blog di myproject
	e.GET("/testimonial", testimonials)
	e.GET("/add-blog/:id", addBlog) //membuka halaman blog form dari klik edit di my project
	// Register
	e.GET("/form-register", formRegister)
	e.POST("/register", register)
	// Login
	e.GET("/form-login", formLogin)
	e.POST("/login", login)
	//Logout
	e.POST("/logout", logout)
	//CRUD
	e.POST("/add-new-blog", addNewBlog)
	e.POST("/delete-blog/:id", deleteBlog)
	e.POST("/update-blog/:id", updateBlog)

	e.Logger.Fatal(e.Start("localhost:8000"))

}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM tb_blogs")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Judul, &each.StartDate, &each.EndDate, &each.Konten, &each.Image, &each.Id, &each.Maindps, &each.Subdps, &each.Shielder, &each.Healer, &each.Durasi)
		if err != nil {
			println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.FormattedStart = each.StartDate.Format("02 January 2006")
		each.FormattedEnd = each.EndDate.Format("02 January 2006")

		result = append(result, each)
	}
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"Blogs":        result,
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), projects)
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

	blogDetail := Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_blogs WHERE id=$1", id).Scan(&blogDetail.Id, &blogDetail.Judul, &blogDetail.StartDate, &blogDetail.EndDate,
		&blogDetail.Konten, &blogDetail.Image, &blogDetail.Id, &blogDetail.Maindps, &blogDetail.Subdps,
		&blogDetail.Shielder, &blogDetail.Healer, &blogDetail.Durasi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogDetail.FormattedStart = blogDetail.StartDate.Format("2 January 2006")
	blogDetail.FormattedEnd = blogDetail.EndDate.Format("2 January 2006")

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}
	dt := map[string]interface{}{
		"Blogs":       blogDetail,
		"DataSession": userData,
	}
	var tmpl, errTmpl = template.ParseFiles("views/blog_list.html")

	if errTmpl != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), dt)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	blogDetail := Project{}
	err := connection.Conn.QueryRow(context.Background(), "SELECT judul, start_date, end_date, konten, image, id, main_dps, sub_dps, shielder, healer, durasi FROM tb_blogs WHERE id=$1", id).
		Scan(&blogDetail.Judul, &blogDetail.StartDate, &blogDetail.EndDate,
			&blogDetail.Konten, &blogDetail.Image, &blogDetail.Id, &blogDetail.Maindps, &blogDetail.Subdps,
			&blogDetail.Shielder, &blogDetail.Healer, &blogDetail.Durasi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogDetail.FormattedStart = blogDetail.StartDate.Format("2 January 2006")
	blogDetail.FormattedEnd = blogDetail.EndDate.Format("2 January 2006")

	dt := map[string]interface{}{
		"Blog": blogDetail,
	}
	var tmpl, tmplErr = template.ParseFiles("views/blog_update.html")

	if tmplErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": tmplErr.Error()})
	}

	return tmpl.Execute(c.Response(), dt)
}

func addNewBlog(c echo.Context) error {

	judul := c.FormValue("judul")
	image := c.FormValue("image")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	durasi := sumDurasi(startDate, endDate)
	konten := c.FormValue("konten")
	maindps := c.FormValue("maindps")
	subdps := c.FormValue("subdps")
	shielder := c.FormValue("shielder")
	healer := c.FormValue("healer")

	println("Judul : " + judul)
	println("Image : " + image)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Konten : " + konten)
	println("Role : " + maindps)
	println("Role : " + subdps)
	println("Role : " + shielder)
	println("Role : " + healer)

	parsedStartDate, _ := time.Parse("2006-01-02", startDate)
	parsedEndDate, _ := time.Parse("2006-01-02", endDate)
	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blogs (judul, start_date, end_date, konten, image, main_dps, sub_dps, shielder, healer, durasi) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		judul, parsedStartDate.Local(), parsedEndDate.Local(), konten, "/public/image/"+image, maindps != "", subdps != "", shielder != "", healer != "", durasi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	println("Index : ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blogs WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	println("Index :", id)

	judul := c.FormValue("judul")
	image := c.FormValue("image")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	durasi := sumDurasi(startDate, endDate)
	konten := c.FormValue("konten")
	maindps := c.FormValue("maindps")
	subdps := c.FormValue("subdps")
	shielder := c.FormValue("shielder")
	healer := c.FormValue("healer")

	parsedStartDate, _ := time.Parse("2006-01-02", startDate)
	parsedEndDate, _ := time.Parse("2006-01-02", endDate)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_blogs SET judul=$1, start_date=$2, end_date=$3, durasi=$4, konten=$5, main_dps=$6, sub_dps=$7, shielder=$8, healer=$9, image=$10 WHERE id=$11",
		judul, parsedStartDate.Local(), parsedEndDate.Local(), konten, maindps != "", subdps != "", durasi, shielder != "", healer != "", "/public/image/"+image, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func sumDurasi(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durasiJam := int(endTime.Sub(startTime).Hours())
	durasiHari := durasiJam / 24
	durasiMinggu := durasiHari / 7
	durasiBulan := durasiMinggu / 4
	durasiTahun := durasiBulan / 12

	var durasi string

	if durasiTahun > 1 {
		durasi = strconv.Itoa(durasiTahun) + " Tahun"
	} else if durasiTahun > 0 {
		durasi = strconv.Itoa(durasiTahun) + "Tahun"
	} else {
		if durasiBulan > 1 {
			durasi = strconv.Itoa(durasiBulan) + " Bulan"
		} else if durasiBulan > 0 {
			durasi = strconv.Itoa(durasiBulan) + " Bulan"
		} else {
			if durasiMinggu > 1 {
				durasi = strconv.Itoa(durasiMinggu) + " Minggu"
			} else if durasiMinggu > 0 {
				durasi = strconv.Itoa(durasiMinggu) + " Minggu"
			} else {
				if durasiHari > 1 {
					durasi = strconv.Itoa(durasiHari) + " Hari"
				} else {
					durasi = strconv.Itoa(durasiHari) + " Hari"
				}
			}
		}
	}

	return durasi
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)
	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/blog_login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}
func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/blog_register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
func register(c echo.Context) error {
	// to make sure request body is form data format, not JSON, XML, etc.
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_users(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success!", true, "/form-login")
}
func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
