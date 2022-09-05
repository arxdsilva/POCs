package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")

	http.HandleFunc("/", middleware(handler))

	go http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second)

	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		log.Fatal(err)
	}

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("body: %s\n\n", string(respB))
	// Starting server at port 8080
	// timeout: 600, ok: true
	// Fields for context.valueCtx
	// context is empty (int)
	// field name: key
	// value: net/http context value http-server
	// field name: val
	// value: &{Addr::8080 Handler:<nil> TLSConfig:0x14000176000 ReadTimeout:0s ReadHeaderTimeout:0s WriteTimeout:0s IdleTimeout:0s MaxHeaderBytes:0 TLSNextProto:map[h2:0x104497340] ConnState:<nil> ErrorLog:<nil> BaseContext:<nil> ConnContext:<nil> inShutdown:0 disableKeepAlives:0 nextProtoOnce:{done:1 m:{state:0 sema:0}} nextProtoErr:<nil> mu:{state:0 sema:0} listeners:map[0x14000112cf0:{}] activeConn:map[0x14000130a00:{}] doneChan:<nil> onShutdown:[0x1044e1930] listenerGroup:{noCopy:{} state1:4294967296 state2:0}}
	// field name: key
	// value: net/http context value local-addr
	// field name: val
	// value: [::1]:8080
	// field name: mu
	// value: {state:0 sema:0}
	// field name: done
	// value: {v:0x1400007e060}
	// field name: children
	// value: map[context.Background.WithValue(type *http.contextKey, val <not Stringer>).WithValue(type *http.contextKey, val [::1]:8080).WithCancel.WithCancel:{}]
	// field name: err
	// value: <nil>
	// field name: mu
	// value: {state:0 sema:0}
	// field name: done
	// value: {v:<nil>}
	// field name: children
	// value: map[]
	// field name: err
	// value: <nil>
	// field name: key
	// value: http.timeout
	// field name: val
	// value: 600
	// field name: key
	// value: 1
	// field name: val
	// value: database
	// body:
}

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), "http.timeout", 600)) // nolint
		h.ServeHTTP(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), 1, "database")

	timeout, ok := ctx.Value("http.timeout").(int)
	fmt.Printf("timeout: %v, ok: %v\n\n", timeout, ok)
	if !ok {
		timeout = 60
	}
	printContextInternals(ctx, false)
}

func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}
