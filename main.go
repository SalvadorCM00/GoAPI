package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type contacts struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
	Email  string `json:"Email"`
	Notas  string `json:"Notas"`
}

func createContact(c echo.Context) error {
	cont := new(contacts)
	if err := c.Bind(cont); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Create(&cont)

	return c.JSON(http.StatusOK, cont)
}

func getContact(c echo.Context) error {
	var cont contacts
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Find(&cont)
	return c.JSON(http.StatusOK, cont)
}

func getContactID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	var res contacts
	db.First(&res, id)
	return c.JSON(http.StatusOK, res)
}

func updateContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	cont := new(contacts)
	db.First(&cont, id)
	if errs := c.Bind(cont); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Save(&cont)
	return c.JSON(http.StatusOK, cont)

}

func deleteContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var res contacts
	db.Where("ID = ? ", id).Delete(&res)
	return c.NoContent(http.StatusNoContent)
}

func getEmails(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var res contacts
	db.First(&res, id)
	return c.JSON(http.StatusOK, res)
}

func getOnlyEmails(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var res contacts
	db.First(&res, id)
	return c.JSON(http.StatusOK, res.Email)
}

func getOnlyNotes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	var res contacts
	db.First(&res, id)
	return c.JSON(http.StatusOK, res.Notas)
}

func updateNotes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	cont := new(contacts)
	db.First(&cont, id)
	if errs := c.Bind(cont); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Model(&cont).Update("Notas", cont.Notas)
	return c.JSON(http.StatusOK, cont)

}

func deleteNotes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	cont := new(contacts)
	db.First(&cont, id)
	if errs := c.Bind(cont); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Model(&cont).Update("Notas", "")
	return c.JSON(http.StatusOK, cont)
}

func createNotes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLO CONEXION A BDD")
	}
	cont := new(contacts)
	db.First(&cont, id)
	if errs := c.Bind(cont); errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Model(&cont).Update("Notas", cont.Notas)
	return c.JSON(http.StatusOK, cont)
}
func main() {
	e := echo.New()

	// CREAR
	e.POST("/contacts", createContact)
	e.POST("/contacts/:id/notes", createNotes)

	// VER
	e.GET("/contacts", getContact)
	e.GET("/contacts/:id", getContactID)
	e.GET("/contacts/:id/emails/:id", getEmails)
	e.GET("/emails/:id", getOnlyEmails)
	e.GET("/contacts/:id/notes", getOnlyNotes)

	// ACTUALIZAR
	e.PUT("/contacts/:id", updateContact)
	e.PUT("/contacts/:id/notes/:id", updateNotes)

	// ELIMINAR
	e.DELETE("/contacts/:id", deleteContact)
	e.DELETE("/contacts/:id/notes/:id", deleteNotes)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
