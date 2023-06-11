package main

import (
	"html/template"
	"net/http"
	"strconv"

	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id        int
	Image     string
	Judul     string
	StartDate string
	EndDate   string
	Durasi    string
	Konten    string
	Maindps   bool
	Subdps    bool
	Shielder  bool
	Healer    bool
}

var dataBlog = []Project{
	{
		Judul:     "Kamisato Ayaka",
		Image:     "/public/image/ayaka.jpeg",
		StartDate: "2023-05-09",
		EndDate:   "2023-06-10",
		Durasi:    "1 Bulan",
		Konten:    "Putri Es Inazuma",
		Maindps:   true,
		Subdps:    false,
		Shielder:  false,
		Healer:    false,
	},
	{
		Judul:     "Keqing",
		Image:     "/public/image/keqing.jpeg",
		StartDate: "09-05-2023",
		EndDate:   "10-06-2023",
		Durasi:    "1 Bulan",
		Konten:    "Liyue Qixing",
		Maindps:   true,
		Subdps:    false,
		Shielder:  false,
		Healer:    false,
	},
}

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	e.Static("/public", "public")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog_list/:id", blogList) //membuka halaman detail blog dari klik judul blog di myproject
	e.GET("/testimonial", testimonials)
	e.GET("/add-blog/:id", addBlog) //membuka halaman blog form dari klik edit di my project

	e.POST("/add-new-blog", addNewBlog)
	e.POST("/delete-blog/:id", deleteBlog)
	e.POST("/update-blog/:id", updateBlog)

	e.Logger.Fatal(e.Start("localhost:8000"))

}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	projects := map[string]interface{}{
		"Blogs": dataBlog,
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

	// 	"Id":    id,
	// 	"Title": "Putri Es Inazuma, Kamisato Ayaka",
	// 	"Content": `Waifu cryo dari Inazuma ini merupakan salah satu karkater yang disukai oleh banyak player Genshin.
	// 				Selain menjadi DPS yang kuat, dia juga memiliki kepribadian yang menarik.`,
	// }
	blogDetail := Project{}
	blogDetail = dataBlog[id]
	blogDetail.Id = id
	// for i, data := range dataBlog {
	// 	if id == i {
	// 		blogDetail = Project{
	// 			Judul:     data.Judul,
	// 			Image:     data.Image,
	// 			StartDate: data.StartDate,
	// 			EndDate:   data.EndDate,
	// 			Durasi:    data.Durasi,
	// 			Konten:    data.Konten,
	// 			Maindps:   data.Maindps,
	// 			Subdps:    data.Subdps,
	// 			Shielder:  data.Shielder,
	// 			Healer:    data.Healer,
	// 		}
	// 	}
	// }
	// data := map[string]interface{}{
	// 	"Blog": blogDetail,
	// }
	var tmpl, err = template.ParseFiles("views/blog_list.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogDetail)
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
	for i, data := range dataBlog {
		if id == i {
			blogDetail = Project{
				Id:        id,
				Judul:     data.Judul,
				Image:     data.Image,
				StartDate: data.StartDate,
				EndDate:   data.EndDate,
				Durasi:    data.Durasi,
				Konten:    data.Konten,
				Maindps:   data.Maindps,
				Subdps:    data.Subdps,
				Shielder:  data.Shielder,
				Healer:    data.Healer,
			}
		}
	}
	data := map[string]interface{}{
		"Blog":  blogDetail,
		"Index": id,
	}
	var tmpl, err = template.ParseFiles("views/blog_update.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
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

	// newData := Project{
	// 	judul: judul,
	// 	StartDate: startDate,
	// 	EndDate: endDate,
	// 	konten: konten,
	// 	maindps: maindps,
	dataNewBlog := Project{
		Judul:     judul,
		Image:     "/public/image/" + image,
		StartDate: startDate,
		EndDate:   endDate,
		Durasi:    durasi,
		Konten:    konten,
		Maindps:   (maindps == "maindps"),
		Subdps:    (subdps == "subdps"),
		Shielder:  (shielder == "subdps"),
		Healer:    (healer == "healer"),
	}
	println(image)
	dataBlog = append(dataBlog, dataNewBlog)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	println("Index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	println("Index :", id)

	judul := c.FormValue("judul")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	durasi := sumDurasi(startDate, endDate)
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

	dataUpdateBlog := Project{
		Judul:     judul,
		StartDate: startDate,
		EndDate:   endDate,
		Durasi:    durasi,
		Konten:    konten,
		Maindps:   (maindps == "maindps"),
		Subdps:    (subdps == "subdps"),
		Shielder:  (shielder == "subdps"),
		Healer:    (healer == "healer"),
	}
	dataBlog[id] = dataUpdateBlog
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
