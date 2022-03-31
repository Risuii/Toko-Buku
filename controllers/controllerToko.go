package controllers

import (
	"Toko/modules"
	"Toko/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ControllerToko struct {
	s services.Kontrak //manggil kontrak
}

func New(kontrak services.Kontrak) ControllerToko {
	return ControllerToko{
		s: kontrak, //s disini mewakili isi dari ControllerToko
	}
}

func (co *ControllerToko) CreateBook(c *gin.Context) {
	id := c.Param("kodeToko")
	var bukus modules.Buku
	if err := c.ShouldBindJSON(&bukus); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	err := co.s.CreateBook(id, &bukus)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Berhasil Membuat Data",
		"kode_buku": id,
	})
}

func (co *ControllerToko) CreateToko(c *gin.Context) {
	var tokos modules.Toko
	tokos.Buku = []modules.Buku{}
	tokos.KodeToko = uuid.New().String()
	if err := c.ShouldBindJSON(&tokos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := co.s.CreateToko(&tokos)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"kode_toko": tokos.KodeToko,
	})
}

func (co *ControllerToko) GetBook(c *gin.Context) {
	id := c.Param("kodeToko")
	idBuku := c.Param("kodeBuku")
	toko, err := co.s.GetToko(id)
	buku := toko.Buku
	fmt.Println(buku)
	var b modules.Buku
	for _, value := range buku {
		if value.KodeBuku == idBuku {
			b = value
			fmt.Println(value.KodeBuku)
			break
		}
		if len(value.KodeBuku) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Buku Tidak Ditemukan",
			})
			fmt.Println(value.KodeBuku)
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, b)
}

func (co *ControllerToko) GetToko(c *gin.Context) {
	id := c.Param("kodeToko")
	toko, err := co.s.GetToko(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": toko,
	})
}

func (co *ControllerToko) GetAll(c *gin.Context) {
	toko, err := co.s.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": toko,
	})
}

func (co *ControllerToko) GetAllBooks(c *gin.Context) {
	id := c.Param("kodeToko")
	toko, err := co.s.GetToko(id)
	buku := toko.Buku
	if len(buku) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Tidak ada Buku",
		})
	}
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": buku,
	})

}

func (co *ControllerToko) UpdateToko(c *gin.Context) {
	id := c.Param("kodeToko")
	var toko modules.Toko
	if err := c.ShouldBindJSON(&toko); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	err := co.s.UpdateToko(&toko, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"messsage": "Berhasil Update Data",
	})
}

func (co *ControllerToko) UpdateBook(c *gin.Context) {
	id := c.Param("kodeToko")
	idBuku := c.Param("kodeBuku")

	var buku modules.Buku

	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	err := co.s.UpdateBook(id, idBuku, &buku)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Update Buku",
	})
}

func (co *ControllerToko) DeleteToko(c *gin.Context) {
	id := c.Param("kodeToko")
	err := co.s.DeleteToko(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Hapus Data",
	})
}

func (co *ControllerToko) DeleteBook(c *gin.Context) {
	id := c.Param("kodeToko")
	idBuku := c.Param("kodeBuku")
	err := co.s.DeleteBook(id, idBuku)
	fmt.Println("ini ID Toko", id)
	fmt.Println("ini ID Buku", idBuku)
	fmt.Println(co.s.DeleteBook(id, idBuku))
	if err != nil {
		fmt.Println("ini error", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Buku Telah Dihapus",
	})
}

func (co *ControllerToko) Routing(rg *gin.RouterGroup) {
	r := rg.Group("/store")
	r.POST("/create", co.CreateToko)
	r.GET("/toko/:kodeToko", co.GetToko)
	r.GET("/getall", co.GetAll)
	r.PUT("/edit/:kodeToko", co.UpdateToko)
	r.DELETE("/delete/:kodeToko", co.DeleteToko)

	r.POST("/book/:kodeToko", co.CreateBook)
	r.GET("/book/:kodeToko/:kodeBuku", co.GetBook)
	r.GET("/book/all/:kodeToko", co.GetAllBooks)
	r.PUT("/book/:kodeToko/:kodeBuku", co.UpdateBook)
	r.DELETE("/book/:kodeToko/:kodeBuku", co.DeleteBook)
}
