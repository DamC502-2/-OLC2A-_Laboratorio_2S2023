// ocl2_proyecto1 project main.go

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var filepath string
var count int = 1
var comments bool = false
var Test string

func main() {
	a := app.New()
	w := a.NewWindow("Ejemplo")
	input := widget.NewMultiLineEntry()
	//console
	console := widget.NewMultiLineEntry()
	console.Disable()

	labelC := widget.NewLabel("Salida")
	//errors
	errors := widget.NewMultiLineEntry()
	errors.Disable()
	labelE := widget.NewLabel("Errores Hallados: ")
	//warnings
	warnings := widget.NewMultiLineEntry()
	warnings.Disable()
	labelW := widget.NewLabel("Advertencias: ")

	//config
	chbxComments := widget.NewCheck("Comentarios", func(b bool) {
		comments = b
	})
	pConf := container.NewGridWithColumns(1, chbxComments)

	labelConf := widget.NewLabel("Configuración: ")

	tabC1 := container.NewTabItem("Salida", container.NewBorder(labelC, nil, nil, nil, console))
	tabC2 := container.NewTabItem("Errores: ", container.NewBorder(labelE, nil, nil, nil, errors))
	tabC3 := container.NewTabItem("Advertencias", container.NewBorder(labelW, nil, nil, nil, warnings))
	tabC4 := container.NewTabItem("Configuración", container.NewBorder(labelConf, nil, nil, nil, pConf))

	tabsConsole := container.NewAppTabs(tabC1, tabC2, tabC3, tabC4)

	splitEditor := container.NewVSplit(input, tabsConsole)

	tabs := container.NewAppTabs(
		container.NewTabItem("Editor", splitEditor),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	input.SetPlaceHolder("Abra un archivo para procesar ó Ingrese el texto aquí")
	input.Move(fyne.NewPos(0, 0))
	input.Resize(fyne.NewSize(1000, 500))
	input.TextStyle = fyne.TextStyle{
		Bold:      false,
		Italic:    false,
		Monospace: true,
		TabWidth:  4,
	}
	console.TextStyle = fyne.TextStyle{
		Bold:      false,
		Italic:    false,
		Monospace: true,
		TabWidth:  4,
	}

	new1 := fyne.NewMenuItem("Nuevo", func() {
		filepath = ""
		w.SetTitle("Nuevo.tswift")
		input.Text = ""
		input.Refresh()
	})
	save1 := fyne.NewMenuItem("Guardar", func() {
		if filepath != "" {
			var f *os.File
			f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				//handle error

			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			//f.Write([]byte(input.Text))
			f.WriteString(input.Text)
		} else {
			saveFileDialog := dialog.NewFileSave(
				func(r fyne.URIWriteCloser, _ error) {
					if r != nil {
						textData := []byte(input.Text)
						r.Write(textData)
						filepath = r.URI().Path()
						w.SetTitle(filepath)
					}

				}, w)
			saveFileDialog.SetFileName("NuevoArchivo" + strconv.Itoa(count-1) + ".tswift")
			saveFileDialog.Show()
		}
	})

	open1 := fyne.NewMenuItem("Abrir", func() {
		openfileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				if r != nil {
					data, _ := ioutil.ReadAll(r)
					result := fyne.NewStaticResource("Nombre", data)
					input.SetText(string(result.StaticContent))
					fmt.Println(result.StaticName + r.URI().Path())
					filepath = r.URI().Path()
					w.SetTitle(filepath)
				}
			}, w)
		openfileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".tswift"}))
		openfileDialog.Show()
	})
	menuItem := fyne.NewMenu("Archivo", new1, save1, open1)

	parseMenu := fyne.NewMenuItem("Leer", func() {
		// funcion leer para ejecutar

		// leer lo del input y poner el console
		// el codigo que hara el boton
		console.Enable()
		console.Text = input.Text
		console.Refresh()

	})
	runMenu := fyne.NewMenuItem("Ejecutar", func() {

		// funcion ejecutar
	})

	menuInterprete := fyne.NewMenu("Ejecutar", parseMenu, runMenu)

	compileMenu := fyne.NewMenuItem("Compilar", func() {
		// funcion compilar
	})

	compileInFile := fyne.NewMenuItem("Compilar en archivo", func() {

		//funcion compilar archivo

	})

	//TODO  añadir las optimizaciones
	menuCompilar := fyne.NewMenu("Compilar", compileMenu, compileInFile)

	disableEty := fyne.NewMenuItem("Desactivar Salidas", func() {
		console.Disable()
		warnings.Disable()
		errors.Disable()
	})

	enableEty := fyne.NewMenuItem("Habilitar Salidas", func() {
		console.Enable()
		warnings.Enable()
		errors.Enable()
	})

	theme := fyne.NewMenuItem("Seleccionar Tema", func() {
		w := a.NewWindow("Fyne Theme")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})

	menuGUI := fyne.NewMenu("GUI", disableEty, enableEty, theme)
	menux1 := fyne.NewMainMenu(menuItem, menuInterprete, menuCompilar, menuGUI)

	w.SetMainMenu(menux1)
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1000, 510))
	w.ShowAndRun()

}
