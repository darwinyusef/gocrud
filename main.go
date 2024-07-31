package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Modelo de datos
type Persona struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

// Slice para almacenar los datos
var personas []Persona

func main() {
	// Inicializar Gin
	app := gin.Default()

	// Rutas
	app.GET("/personas", obtenerPersonas)
	app.POST("/personas", crearPersona)
	app.GET("/personas/:id", obtenerPersona)
	app.PUT("/personas/:id", actualizarPersona)
	app.DELETE("/personas/:id", eliminarPersona)

	// Iniciar servidor
	app.Run(":8000")
}

// Obtener todas las personas
func obtenerPersonas(c *gin.Context) {
	c.JSON(http.StatusOK, personas)
}

// Crear una nueva persona
func crearPersona(c *gin.Context) {
	var nuevaPersona Persona
	if err := c.BindJSON(&nuevaPersona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personas = append(personas, nuevaPersona)
	c.JSON(http.StatusCreated, nuevaPersona)
}

// Obtener una persona por ID
func obtenerPersona(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	for _, persona := range personas {
		if persona.ID == idInt {
			c.JSON(http.StatusOK, persona)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrada"})
}

// Actualizar una persona
func actualizarPersona(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	for i, persona := range personas {
		if persona.ID == idInt {
			var nuevaPersona Persona
			if err := c.BindJSON(&nuevaPersona); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			personas[i] = nuevaPersona
			c.JSON(http.StatusOK, nuevaPersona)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrada"})
}

// Eliminar una persona
func eliminarPersona(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	for i, persona := range personas {
		if persona.ID == idInt {
			personas = append(personas[:i], personas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensaje": "Persona eliminada"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrada"})
}
