// Code generated by templ@(devel) DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func show(message string) templ.ComponentScript {
	return templ.ComponentScript{
		Name:     `__templ_show_3e0b`,
		Function: `function __templ_show_3e0b(message){alert(message)}`,
		Call:     templ.SafeScript(`__templ_show_3e0b`, message),
	}
}

func page(message string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html><head><title>")
		if err != nil {
			return err
		}
		var_2 := `Test`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title></head><body><h1>")
		if err != nil {
			return err
		}
		var_3 := `Test`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1>")
		if err != nil {
			return err
		}
		err = templ.RenderScriptItems(ctx, templBuffer, show(message))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<input type=\"button\" value=\"Click me\" onclick=\"")
		if err != nil {
			return err
		}
		var var_4 templ.ComponentScript = show(message)
		_, err = templBuffer.WriteString(var_4.Call)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}