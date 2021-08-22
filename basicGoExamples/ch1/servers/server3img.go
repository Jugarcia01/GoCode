// Server1echo es un miniServer web que retorna el path/ruta actual
// El prog. es solo un puñado de lineas, puesto que la mayoria del trabajo la realizan las funciones de las librerias empleadas. El prog. conecta un manejador para aquellas URLs entrantes hacia el puerto 8000 que comienzan con /
// Una peticion se representa como una estructura de tipo http.Requestla cual contiene un num. de campos relacionados, uno de los cuales es la URL de las peticiones entrantes. Cuando una petic. llega, se pasa a la func manejadora, la cual extrae la ruta del componente (/hello) de la URL de la petic. y la retorna a la respuesta, empleando fmt.Fprintf

// NOTA IMPORTANTE: Para iniciar este server en segundo plano en Linux/macOS ejecuta desde la terminal: 
//$ go run Server1echo.go &
// Posteriormente realiza peticiones desde ka linea de comandos así por ejem:
//$ go build gopl.io/ch1/fetch
//$ ./fetch http://localhost:8080
//URL.Path = "/"
//$ ./fetch http://localhost:8080/help
//URL.Path = "/help"

package main

import (
    //"fmt"
    "log"
    "net/http" 
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
 
  )

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
    
  // El manejador retorna la imagen del componente de la URL en la peticion
/*  handler := func(w http.ResponseWriter, r *http.Request) {
//gifLissa(w)
	  lissajous(w)
  }
    // Cada solicitud llama al manejador/handler
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
*/

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  lissajous(w)
})

//func lissajous(out io.Writer) {
func lissajous(out *http.ResponseWriter) anim gif {
	const (
		cycles = 1 // number of complete x oscillator
		revolutions
		res     = 0.001 // angular resolution
		size    = 50   // image canvas covers [-size..+size]
		nframes = 10    // number of animation frames
		delay   = 7     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
