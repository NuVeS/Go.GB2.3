package gb23

import (
	"fmt"

	"github.com/andybalholm/brotli"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println(fasthttp.StateActive)
	fmt.Println(brotli.BestCompression)
}
