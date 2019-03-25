package printer

import (
	"fmt"

	info "github.com/Helcaraxan/moddep-direct-dep"
)

func Print(i *info.Info) {
	if i == nil {
		i = info.DefaultInfo()
	}
	fmt.Println(i)
}
