//modificar o programa lissajous para gerar imagens em varias cores
//modificando pallete e SetColorIndex
//para executar usar os comandos: go build
//e : ./nome_do_arquivo >out.gif

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0xAA, 0xAA, 0x77, 0xAA}}

const (
	whiteIndex = 0 //primeira cor da paleta
	blackIndex = 1 //segunda cor da paleta
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //numero de voltas completas do oscilador x
		res     = 0.001 //resolucao angular
		size    = 100   //canvas da imagem cobre [-size..+size]
		nframes = 64    //numero de quadros da animacao
		delay   = 8     //tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 //frequencia relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //diferenca de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 3)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTA: ignorando erros de codificacao
}
