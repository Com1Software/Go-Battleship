package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"image/color"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	xip := fmt.Sprintf("%s", GetOutboundIP())
	port := "8080"
	a := app.New()
	w := a.NewWindow("Listening on " + xip + ":" + port)
	ctla := true
	memo := widget.NewEntry()
	memo.SetPlaceHolder("Enter an IP address to sync with...")
	memo.MultiLine = true // Enable multiline for larger text fields
	memoa := container.NewGridWrap(fyne.NewSize(380, 40), memo)

	memo1 := widget.NewEntry()
	memo1.SetPlaceHolder("...")
	memo1.MultiLine = true // Enable multiline for larger text fields
	memo1a := container.NewGridWrap(fyne.NewSize(480, 40), memo1)

	helloButton := widget.NewButton("Connect", func() {
		url := memo.Text
		go func() {
			for {
				data := ReadURL(url)

				// Use fyne.Do to safely update the UI from a goroutine
				fyne.Do(func() {
					memo1.SetText(data)
				})

				time.Sleep(1 * time.Second) // Adjust the refresh rate as needed
			}
		}()

	})
	hellobuttonContainer := container.NewGridWrap(fyne.NewSize(100, 40), helloButton)

	hellolayoutContainer := container.NewHBox(memoa, hellobuttonContainer)

	button := NewCustomButton("", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button1 := NewCustomButton("1", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button2 := NewCustomButton("2", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button3 := NewCustomButton("3", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button4 := NewCustomButton("4", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button5 := NewCustomButton("5", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button6 := NewCustomButton("6", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button7 := NewCustomButton("7", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button8 := NewCustomButton("8", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button9 := NewCustomButton("9", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button10 := NewCustomButton("10", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0a := NewCustomButton("A", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0b := NewCustomButton("B", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0c := NewCustomButton("C", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0d := NewCustomButton("D", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0e := NewCustomButton("E", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0f := NewCustomButton("F", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0g := NewCustomButton("G", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0h := NewCustomButton("H", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0i := NewCustomButton("I", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button0j := NewCustomButton("J", color.RGBA{0, 0, 0, 0}, func() {
	}, ctla)

	button1a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10a := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10b := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10c := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10d := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10e := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10f := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10g := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10h := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10i := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button1j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button2j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button3j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button4j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button5j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button6j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button7j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button8j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	button9j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)
	button10j := NewCustomButton("", color.RGBA{0, 0, 255, 255}, func() {
		println("Button clicked!")
	}, ctla)

	buttonContainer := container.NewGridWrap(fyne.NewSize(40, 40), button)
	buttonContainer1 := container.NewGridWrap(fyne.NewSize(40, 40), button1)
	buttonContainer2 := container.NewGridWrap(fyne.NewSize(40, 40), button2)
	buttonContainer3 := container.NewGridWrap(fyne.NewSize(40, 40), button3)
	buttonContainer4 := container.NewGridWrap(fyne.NewSize(40, 40), button4)
	buttonContainer5 := container.NewGridWrap(fyne.NewSize(40, 40), button5)
	buttonContainer6 := container.NewGridWrap(fyne.NewSize(40, 40), button6)
	buttonContainer7 := container.NewGridWrap(fyne.NewSize(40, 40), button7)
	buttonContainer8 := container.NewGridWrap(fyne.NewSize(40, 40), button8)
	buttonContainer9 := container.NewGridWrap(fyne.NewSize(40, 40), button9)
	buttonContainer10 := container.NewGridWrap(fyne.NewSize(40, 40), button10)

	buttonContainer0a := container.NewGridWrap(fyne.NewSize(40, 40), button0a)
	buttonContainer0b := container.NewGridWrap(fyne.NewSize(40, 40), button0b)
	buttonContainer0c := container.NewGridWrap(fyne.NewSize(40, 40), button0c)
	buttonContainer0d := container.NewGridWrap(fyne.NewSize(40, 40), button0d)
	buttonContainer0e := container.NewGridWrap(fyne.NewSize(40, 40), button0e)
	buttonContainer0f := container.NewGridWrap(fyne.NewSize(40, 40), button0f)
	buttonContainer0g := container.NewGridWrap(fyne.NewSize(40, 40), button0g)
	buttonContainer0h := container.NewGridWrap(fyne.NewSize(40, 40), button0h)
	buttonContainer0i := container.NewGridWrap(fyne.NewSize(40, 40), button0i)
	buttonContainer0j := container.NewGridWrap(fyne.NewSize(40, 40), button0j)

	buttonContainer1a := container.NewGridWrap(fyne.NewSize(40, 40), button1a)
	buttonContainer2a := container.NewGridWrap(fyne.NewSize(40, 40), button2a)
	buttonContainer3a := container.NewGridWrap(fyne.NewSize(40, 40), button3a)
	buttonContainer4a := container.NewGridWrap(fyne.NewSize(40, 40), button4a)
	buttonContainer5a := container.NewGridWrap(fyne.NewSize(40, 40), button5a)
	buttonContainer6a := container.NewGridWrap(fyne.NewSize(40, 40), button6a)
	buttonContainer7a := container.NewGridWrap(fyne.NewSize(40, 40), button7a)
	buttonContainer8a := container.NewGridWrap(fyne.NewSize(40, 40), button8a)
	buttonContainer9a := container.NewGridWrap(fyne.NewSize(40, 40), button9a)
	buttonContainer10a := container.NewGridWrap(fyne.NewSize(40, 40), button10a)

	buttonContainer1b := container.NewGridWrap(fyne.NewSize(40, 40), button1b)
	buttonContainer2b := container.NewGridWrap(fyne.NewSize(40, 40), button2b)
	buttonContainer3b := container.NewGridWrap(fyne.NewSize(40, 40), button3b)
	buttonContainer4b := container.NewGridWrap(fyne.NewSize(40, 40), button4b)
	buttonContainer5b := container.NewGridWrap(fyne.NewSize(40, 40), button5b)
	buttonContainer6b := container.NewGridWrap(fyne.NewSize(40, 40), button6b)
	buttonContainer7b := container.NewGridWrap(fyne.NewSize(40, 40), button7b)
	buttonContainer8b := container.NewGridWrap(fyne.NewSize(40, 40), button8b)
	buttonContainer9b := container.NewGridWrap(fyne.NewSize(40, 40), button9b)
	buttonContainer10b := container.NewGridWrap(fyne.NewSize(40, 40), button10b)

	buttonContainer1c := container.NewGridWrap(fyne.NewSize(40, 40), button1c)
	buttonContainer2c := container.NewGridWrap(fyne.NewSize(40, 40), button2c)
	buttonContainer3c := container.NewGridWrap(fyne.NewSize(40, 40), button3c)
	buttonContainer4c := container.NewGridWrap(fyne.NewSize(40, 40), button4c)
	buttonContainer5c := container.NewGridWrap(fyne.NewSize(40, 40), button5c)
	buttonContainer6c := container.NewGridWrap(fyne.NewSize(40, 40), button6c)
	buttonContainer7c := container.NewGridWrap(fyne.NewSize(40, 40), button7c)
	buttonContainer8c := container.NewGridWrap(fyne.NewSize(40, 40), button8c)
	buttonContainer9c := container.NewGridWrap(fyne.NewSize(40, 40), button9c)
	buttonContainer10c := container.NewGridWrap(fyne.NewSize(40, 40), button10c)

	buttonContainer1d := container.NewGridWrap(fyne.NewSize(40, 40), button1d)
	buttonContainer2d := container.NewGridWrap(fyne.NewSize(40, 40), button2d)
	buttonContainer3d := container.NewGridWrap(fyne.NewSize(40, 40), button3d)
	buttonContainer4d := container.NewGridWrap(fyne.NewSize(40, 40), button4d)
	buttonContainer5d := container.NewGridWrap(fyne.NewSize(40, 40), button5d)
	buttonContainer6d := container.NewGridWrap(fyne.NewSize(40, 40), button6d)
	buttonContainer7d := container.NewGridWrap(fyne.NewSize(40, 40), button7d)
	buttonContainer8d := container.NewGridWrap(fyne.NewSize(40, 40), button8d)
	buttonContainer9d := container.NewGridWrap(fyne.NewSize(40, 40), button9d)
	buttonContainer10d := container.NewGridWrap(fyne.NewSize(40, 40), button10d)

	buttonContainer1e := container.NewGridWrap(fyne.NewSize(40, 40), button1e)
	buttonContainer2e := container.NewGridWrap(fyne.NewSize(40, 40), button2e)
	buttonContainer3e := container.NewGridWrap(fyne.NewSize(40, 40), button3e)
	buttonContainer4e := container.NewGridWrap(fyne.NewSize(40, 40), button4e)
	buttonContainer5e := container.NewGridWrap(fyne.NewSize(40, 40), button5e)
	buttonContainer6e := container.NewGridWrap(fyne.NewSize(40, 40), button6e)
	buttonContainer7e := container.NewGridWrap(fyne.NewSize(40, 40), button7e)
	buttonContainer8e := container.NewGridWrap(fyne.NewSize(40, 40), button8e)
	buttonContainer9e := container.NewGridWrap(fyne.NewSize(40, 40), button9e)
	buttonContainer10e := container.NewGridWrap(fyne.NewSize(40, 40), button10e)

	buttonContainer1f := container.NewGridWrap(fyne.NewSize(40, 40), button1f)
	buttonContainer2f := container.NewGridWrap(fyne.NewSize(40, 40), button2f)
	buttonContainer3f := container.NewGridWrap(fyne.NewSize(40, 40), button3f)
	buttonContainer4f := container.NewGridWrap(fyne.NewSize(40, 40), button4f)
	buttonContainer5f := container.NewGridWrap(fyne.NewSize(40, 40), button5f)
	buttonContainer6f := container.NewGridWrap(fyne.NewSize(40, 40), button6f)
	buttonContainer7f := container.NewGridWrap(fyne.NewSize(40, 40), button7f)
	buttonContainer8f := container.NewGridWrap(fyne.NewSize(40, 40), button8f)
	buttonContainer9f := container.NewGridWrap(fyne.NewSize(40, 40), button9f)
	buttonContainer10f := container.NewGridWrap(fyne.NewSize(40, 40), button10f)

	buttonContainer1g := container.NewGridWrap(fyne.NewSize(40, 40), button1g)
	buttonContainer2g := container.NewGridWrap(fyne.NewSize(40, 40), button2g)
	buttonContainer3g := container.NewGridWrap(fyne.NewSize(40, 40), button3g)
	buttonContainer4g := container.NewGridWrap(fyne.NewSize(40, 40), button4g)
	buttonContainer5g := container.NewGridWrap(fyne.NewSize(40, 40), button5g)
	buttonContainer6g := container.NewGridWrap(fyne.NewSize(40, 40), button6g)
	buttonContainer7g := container.NewGridWrap(fyne.NewSize(40, 40), button7g)
	buttonContainer8g := container.NewGridWrap(fyne.NewSize(40, 40), button8g)
	buttonContainer9g := container.NewGridWrap(fyne.NewSize(40, 40), button9g)
	buttonContainer10g := container.NewGridWrap(fyne.NewSize(40, 40), button10g)

	buttonContainer1h := container.NewGridWrap(fyne.NewSize(40, 40), button1h)
	buttonContainer2h := container.NewGridWrap(fyne.NewSize(40, 40), button2h)
	buttonContainer3h := container.NewGridWrap(fyne.NewSize(40, 40), button3h)
	buttonContainer4h := container.NewGridWrap(fyne.NewSize(40, 40), button4h)
	buttonContainer5h := container.NewGridWrap(fyne.NewSize(40, 40), button5h)
	buttonContainer6h := container.NewGridWrap(fyne.NewSize(40, 40), button6h)
	buttonContainer7h := container.NewGridWrap(fyne.NewSize(40, 40), button7h)
	buttonContainer8h := container.NewGridWrap(fyne.NewSize(40, 40), button8h)
	buttonContainer9h := container.NewGridWrap(fyne.NewSize(40, 40), button9h)
	buttonContainer10h := container.NewGridWrap(fyne.NewSize(40, 40), button10h)

	buttonContainer1i := container.NewGridWrap(fyne.NewSize(40, 40), button1i)
	buttonContainer2i := container.NewGridWrap(fyne.NewSize(40, 40), button2i)
	buttonContainer3i := container.NewGridWrap(fyne.NewSize(40, 40), button3i)
	buttonContainer4i := container.NewGridWrap(fyne.NewSize(40, 40), button4i)
	buttonContainer5i := container.NewGridWrap(fyne.NewSize(40, 40), button5i)
	buttonContainer6i := container.NewGridWrap(fyne.NewSize(40, 40), button6i)
	buttonContainer7i := container.NewGridWrap(fyne.NewSize(40, 40), button7i)
	buttonContainer8i := container.NewGridWrap(fyne.NewSize(40, 40), button8i)
	buttonContainer9i := container.NewGridWrap(fyne.NewSize(40, 40), button9i)
	buttonContainer10i := container.NewGridWrap(fyne.NewSize(40, 40), button10i)

	buttonContainer1j := container.NewGridWrap(fyne.NewSize(40, 40), button1j)
	buttonContainer2j := container.NewGridWrap(fyne.NewSize(40, 40), button2j)
	buttonContainer3j := container.NewGridWrap(fyne.NewSize(40, 40), button3j)
	buttonContainer4j := container.NewGridWrap(fyne.NewSize(40, 40), button4j)
	buttonContainer5j := container.NewGridWrap(fyne.NewSize(40, 40), button5j)
	buttonContainer6j := container.NewGridWrap(fyne.NewSize(40, 40), button6j)
	buttonContainer7j := container.NewGridWrap(fyne.NewSize(40, 40), button7j)
	buttonContainer8j := container.NewGridWrap(fyne.NewSize(40, 40), button8j)
	buttonContainer9j := container.NewGridWrap(fyne.NewSize(40, 40), button9j)
	buttonContainer10j := container.NewGridWrap(fyne.NewSize(40, 40), button10j)

	layoutContainer := container.NewHBox(buttonContainer, buttonContainer0a, buttonContainer0b, buttonContainer0c,
		buttonContainer0d, buttonContainer0e, buttonContainer0f, buttonContainer0g, buttonContainer0h,
		buttonContainer0i, buttonContainer0j)

	layoutContainera := container.NewHBox(buttonContainer1, buttonContainer1a, buttonContainer2a, buttonContainer3a,
		buttonContainer4a, buttonContainer5a, buttonContainer6a, buttonContainer7a, buttonContainer8a,
		buttonContainer9a, buttonContainer10a)

	layoutContainerb := container.NewHBox(buttonContainer2, buttonContainer1b, buttonContainer2b, buttonContainer3b,
		buttonContainer4b, buttonContainer5b, buttonContainer6b, buttonContainer7b, buttonContainer8b,
		buttonContainer9b, buttonContainer10b)

	layoutContainerc := container.NewHBox(buttonContainer3, buttonContainer1c, buttonContainer2c, buttonContainer3c,
		buttonContainer4c, buttonContainer5c, buttonContainer6c, buttonContainer7c, buttonContainer8c,
		buttonContainer9c, buttonContainer10c)

	layoutContainerd := container.NewHBox(buttonContainer4, buttonContainer1d, buttonContainer2d, buttonContainer3d,
		buttonContainer4d, buttonContainer5d, buttonContainer6d, buttonContainer7d, buttonContainer8d,
		buttonContainer9d, buttonContainer10d)

	layoutContainere := container.NewHBox(buttonContainer5, buttonContainer1e, buttonContainer2e, buttonContainer3e,
		buttonContainer4e, buttonContainer5e, buttonContainer6e, buttonContainer7e, buttonContainer8e,
		buttonContainer9e, buttonContainer10e)

	layoutContainerf := container.NewHBox(buttonContainer6, buttonContainer1f, buttonContainer2f, buttonContainer3f,
		buttonContainer4f, buttonContainer5f, buttonContainer6f, buttonContainer7f, buttonContainer8f,
		buttonContainer9f, buttonContainer10f)

	layoutContainerg := container.NewHBox(buttonContainer7, buttonContainer1g, buttonContainer2g, buttonContainer3g,
		buttonContainer4g, buttonContainer5g, buttonContainer6g, buttonContainer7g, buttonContainer8g,
		buttonContainer9g, buttonContainer10g)

	layoutContainerh := container.NewHBox(buttonContainer8, buttonContainer1h, buttonContainer2h, buttonContainer3h,
		buttonContainer4h, buttonContainer5h, buttonContainer6h, buttonContainer7h, buttonContainer8h,
		buttonContainer9h, buttonContainer10h)

	layoutContaineri := container.NewHBox(buttonContainer9, buttonContainer1i, buttonContainer2i, buttonContainer3i,
		buttonContainer4i, buttonContainer5i, buttonContainer6i, buttonContainer7i, buttonContainer8i,
		buttonContainer9i, buttonContainer10i)

	layoutContainerj := container.NewHBox(buttonContainer10, buttonContainer1j, buttonContainer2j, buttonContainer3j,
		buttonContainer4j, buttonContainer5j, buttonContainer6j, buttonContainer7j, buttonContainer8j,
		buttonContainer9j, buttonContainer10j)

	exitButton := widget.NewButton("Exit", func() {
		os.Exit(0)
	})
	exitButtona := container.NewGridWrap(fyne.NewSize(480, 40), exitButton)
	w.Resize(fyne.NewSize(1000, 300))
	w.SetContent(container.NewVBox(
		hellolayoutContainer,
		memo1a,
		layoutContainer,
		layoutContainera,
		layoutContainerb,
		layoutContainerc,
		layoutContainerd,
		layoutContainere,
		layoutContainerf,
		layoutContainerg,
		layoutContainerh,
		layoutContaineri,
		layoutContainerj,
		exitButtona,
	))
	w.ShowAndRun()
}

// CustomButton creates a button with a colored background
type CustomButton struct {
	widget.BaseWidget
	text     string
	color    color.Color
	onTapped func()
	bgRect   *canvas.Rectangle
}

func NewCustomButton(text string, bgColor color.Color, onTapped func(), ctl bool) *CustomButton {
	btn := &CustomButton{text: text, color: bgColor, onTapped: onTapped}
	btn.ExtendBaseWidget(btn)
	return btn
}

// CreateRenderer defines the custom rendering for the button
func (b *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(b.text)
	b.bgRect = canvas.NewRectangle(b.color) // Use a modifiable background

	content := container.NewMax(b.bgRect, label)
	return widget.NewSimpleRenderer(content)
}

func (b *CustomButton) Tapped(_ *fyne.PointEvent) {
	if b.onTapped != nil {
		b.onTapped()
	}
	hit := false
	if hit {
		b.bgRect.FillColor = color.RGBA{255, 0, 0, 255} // Change to Red
	} else {
		b.bgRect.FillColor = color.RGBA{255, 255, 255, 255} // Change to Green
	}
	b.bgRect.Refresh() // Update UI
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

type message struct {
	Controller string `xml:"controller"`
	DateTime   string `xml:"date_time"`
	Command    string `xml:"command"`
	Value      string `xml:"value"`
}

func ReadURL(url string) string {
	msg := &message{}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)

	for {
		line, erra := reader.ReadBytes('\n')
		if erra != nil {
			log.Fatal(erra) // Ensure correct error logging
		}
		xml.Unmarshal(line, &msg)
		break // Exit after the first read
	}

	// Construct the string with extracted values
	return msg.Controller + " " + msg.DateTime + " " + msg.Command + " " + msg.Value
}

type Agent struct {
	Notifier    chan []byte
	newuser     chan chan []byte
	closinguser chan chan []byte
	user        map[chan []byte]bool
}

func SSE() (agent *Agent) {
	agent = &Agent{
		Notifier:    make(chan []byte, 1),
		newuser:     make(chan chan []byte),
		closinguser: make(chan chan []byte),
		user:        make(map[chan []byte]bool),
	}
	go agent.listen()
	return
}

func (agent *Agent) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Error ", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	mChan := make(chan []byte)
	agent.newuser <- mChan
	defer func() {
		agent.closinguser <- mChan
	}()
	notify := req.Context().Done()
	go func() {
		<-notify
		agent.closinguser <- mChan
	}()
	for {
		fmt.Fprintf(rw, "%s", <-mChan)
		flusher.Flush()
	}

}

func (agent *Agent) listen() {
	for {
		select {
		case s := <-agent.newuser:
			agent.user[s] = true
		case s := <-agent.closinguser:
			delete(agent.user, s)
		case event := <-agent.Notifier:
			for userMChan, _ := range agent.user {
				userMChan <- event
			}
		}
	}

}
