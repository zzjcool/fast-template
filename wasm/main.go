package main

import (
	"syscall/js"

	"github.com/zzjcool/fast-template/render"
)

func renderf(this js.Value,args []js.Value)any{
	return render.Render("","")
}
func main(){
	c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
	js.Global().Set("renderf", js.FuncOf(renderf))
    <-c
}