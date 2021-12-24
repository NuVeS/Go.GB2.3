package gb23

import (
	"fmt"

	"github.com/andybalholm/brotli"
)

func main() {
	fmt.Println(brotli.BestCompression)
}
