package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

type Person struct {
	Nombre        string
	Edad          int
	Correo        string
	Activo        bool
	Pais          string
	FechaRegistro string
	Puntos        int
}

func main() {
	// Leer el JSON
	jsonData, err := os.ReadFile("Data/data.json")
	if err != nil {
		fmt.Println("Error reading the JSON:", err)
		return
	}
	// Decodificar el JSON
	var people []Person
	err = json.Unmarshal(jsonData, &people)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Escribir el Excel con Excelize
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheet := "Sheet1"
	// Setear nombre de la pestaña
	f.SetSheetName("Sheet1", sheet)
	// Setear el ancho de la columna
	f.SetColWidth("Sheet1", "A", "A", 20)
	f.SetColWidth("Sheet1", "C", "C", 30)
	f.SetColWidth("Sheet1", "E", "F", 20)
	// Setear estilos
	style, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"90B4D6"},
			Pattern: 1},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	styleColumnBeginingFromLeft, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	cellStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"FFF8E1"},
			Pattern: 1},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	// Agregar estilos a celdas
	f.SetCellStyle("Sheet1", "B1", "B100", styleColumnBeginingFromLeft)
	f.SetCellStyle("Sheet1", "D1", "D100", styleColumnBeginingFromLeft)
	f.SetCellStyle("Sheet1", "G1", "G100", styleColumnBeginingFromLeft)
	f.SetCellStyle("Sheet1", "A1", "G1", style)

	// Escribir encabezados
	f.SetCellValue(sheet, "A1", "Nombre")
	f.SetCellValue(sheet, "B1", "Edad")
	f.SetCellValue(sheet, "C1", "Correo")
	f.SetCellValue(sheet, "D1", "Activo")
	f.SetCellValue(sheet, "E1", "País")
	f.SetCellValue(sheet, "F1", "FechaRegistro")
	f.SetCellValue(sheet, "G1", "Puntos")

	// Ingresar filas
	/* for i, p := range people {
		row := i + 2 // Iniciar desde la fila 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.Nombre)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.Edad)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.Correo)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), p.Activo)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), p.Pais)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), p.FechaRegistro)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), p.Puntos)
	} */

	for i, p := range people {
		row := i + 2

		celdas := map[string]interface{}{
			fmt.Sprintf("A%d", row): p.Nombre,
			fmt.Sprintf("B%d", row): p.Edad,
			fmt.Sprintf("C%d", row): p.Correo,
			fmt.Sprintf("D%d", row): p.Activo,
			fmt.Sprintf("E%d", row): p.Pais,
			fmt.Sprintf("F%d", row): p.FechaRegistro,
			fmt.Sprintf("G%d", row): p.Puntos,
		}

		for celda, valor := range celdas {
			f.SetCellValue(sheet, celda, valor)
			f.SetCellStyle(sheet, celda, celda, cellStyle)
		}
	}

	// Crear nombre
	timestamp := time.Now().Format("2006-01-02")
	fileName := "document_" + timestamp + ".xlsx"
	// Guardar
	err = f.SaveAs("Excels/" + fileName)
	if err != nil {
		fmt.Println("Error durante el guardado:", err)
		return
	}
}
