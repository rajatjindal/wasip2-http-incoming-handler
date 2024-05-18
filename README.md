## wasip2 http incoming handler

When trying out wasip2 changes, I am running into following issue:

Environment:

- Macbook pro M2 processor
- Tinygo compiled from https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2 (3ef711afab91ddedc66dc06d91eaa16cdb1f8c6a)
- wit-bindgen-go compiled from https://github.com/ydnar/wasm-tools-go (aa782301db370ce82be09d373efc94a281ccee9a)

steps to reproduce function mismatch error

```bash
git clone https://github.com/rajatjindal/wasip2-http-incoming-handler.git
git checkout 058507c5540fa69524f717f32f31b53328329a0f
tinygo build -target=wasip2 -wit-package ./wit --wit-world http-trigger -o main.wasm -x -work main.go
```

Error:

```
error: failed to encode a component from module

Caused by:
    0: failed to validate import interface `wasi:http/types@0.2.0`
    1: type mismatch for function `[static]response-outparam.set`: expected `[I32, I32, I32, I32, I64, I32, I32, I32, I32] -> []` but found `[I32, I32, I32, I32, I32, I32, I32, I32, I32] -> []`
error: wasm-tools failed: exit status 1
```

Full logs:

```
WORK=/var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888
wasm-ld --stack-first --no-demangle -L /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/ -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main.o /Users/rajatjindal/Library/Caches/tinygo/compiler-rt-wasm32-unknown-wasi-generic/lib.a /Users/rajatjindal/Library/Caches/tinygo/obj-363c24a529264e57aec91eacb3cbc247bacb915b67966242e30f7fb9.bc /Users/rajatjindal/Library/Caches/tinygo/obj-870ef2ac5a778e9704c61e87bde8b24392ff596a3b0126fe5821e7b1.bc /Users/rajatjindal/Library/Caches/tinygo/wasmbuiltins-wasm32-unknown-wasi-generic/lib.a -mllvm -mcpu=generic --lto-O2 --thinlto-cache-dir=/Users/rajatjindal/Library/Caches/tinygo/thinlto -mllvm --rotation-max-header-size=0
/opt/homebrew/bin/wasm-opt --asyncify -Oz -g /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main --output /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main
wasm-tools component embed -w http-trigger ./wit /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main
wasm-tools component new /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo4056239888/main
error: failed to encode a component from module

Caused by:
    0: failed to validate import interface `wasi:http/types@0.2.0`
    1: type mismatch for function `[static]response-outparam.set`: expected `[I32, I32, I32, I32, I64, I32, I32, I32, I32] -> []` but found `[I32, I32, I32, I32, I32, I32, I32, I32, I32] -> []`
error: wasm-tools failed: exit status 1
```

I am able to fix that error by making following change to [wasi/http/types/types.wit.go](types.wit.go):

```
diff --git a/main.go b/main.go
index 9edb840..efb9e22 100644
--- a/main.go
+++ b/main.go
@@ -8,7 +8,7 @@ import (
 
 func init() {
 	incominghandler.Handle = func(request types.IncomingRequest, responseOut types.ResponseOutparam) {
-		err := cm.Err[cm.ErrResult[types.OutgoingResponse, types.ErrorCode]](types.ErrorCodeConfigurationError())
+		err := cm.Err[cm.OKResult[types.OutgoingResponse, types.ErrorCode]](types.ErrorCodeConfigurationError())
 		types.ResponseOutparamSet(cm.ResourceNone, err)
 	}
 }
diff --git a/wasi/http/types/types.wit.go b/wasi/http/types/types.wit.go
index f7c3e6c..b52d16e 100644
--- a/wasi/http/types/types.wit.go
+++ b/wasi/http/types/types.wit.go
@@ -1487,13 +1487,13 @@ func (self ResponseOutparam) wasmimport_ResourceDrop()
 //	error-code>)
 //
 //go:nosplit
-func ResponseOutparamSet(param ResponseOutparam, response cm.ErrResult[OutgoingResponse, ErrorCode]) {
+func ResponseOutparamSet(param ResponseOutparam, response cm.OKResult[OutgoingResponse, ErrorCode]) {
 	wasmimport_ResponseOutparamSet(param, response)
 }
 
 //go:wasmimport wasi:http/types@0.2.0 [static]response-outparam.set
 //go:noescape
-func wasmimport_ResponseOutparamSet(param ResponseOutparam, response cm.ErrResult[OutgoingResponse, ErrorCode])
+func wasmimport_ResponseOutparamSet(param ResponseOutparam, response cm.OKResult[OutgoingResponse, ErrorCode])
 
 // StatusCode represents the imported type "wasi:http/types@0.2.0#status-code".
 //
```

After which I am able to compile the code successfully:

```
$ tinygo build -target=wasip2 -wit-package ./wit --wit-world http-trigger -o main.wasm -x -work main.go

WORK=/var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658
wasm-ld --stack-first --no-demangle -L /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/ -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main.o /Users/rajatjindal/Library/Caches/tinygo/compiler-rt-wasm32-unknown-wasi-generic/lib.a /Users/rajatjindal/Library/Caches/tinygo/obj-363c24a529264e57aec91eacb3cbc247bacb915b67966242e30f7fb9.bc /Users/rajatjindal/Library/Caches/tinygo/obj-870ef2ac5a778e9704c61e87bde8b24392ff596a3b0126fe5821e7b1.bc /Users/rajatjindal/Library/Caches/tinygo/wasmbuiltins-wasm32-unknown-wasi-generic/lib.a -mllvm -mcpu=generic --lto-O2 --thinlto-cache-dir=/Users/rajatjindal/Library/Caches/tinygo/thinlto -mllvm --rotation-max-header-size=0
/opt/homebrew/bin/wasm-opt --asyncify -Oz -g /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main --output /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main
wasm-tools component embed -w http-trigger ./wit /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main
wasm-tools component new /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main -o /var/folders/5r/6y6fr12x1dq6qnt74zn4_1wm0000gn/T/tinygo2565817658/main
```

and able to start the `wasmtime serve` with it:

```
RUST_LOG=trace WASMTIME_BACKTRACE_DETAILS=1 wasmtime serve main.wasm --wasi common
Serving HTTP on http://0.0.0.0:8080/
```

However, when I curl this endpoint, I get following stacktrace:


```
    context: "error while executing at wasm backtrace:\n    0: 0x1332a - wit-component:shim!indirect-wasi:io/streams@0.2.0-[method]output-stream.blocking-write-and-flush\n    1: 0x25f5 - (internal/wasi/io/v0.2.0/streams.OutputStream).BlockingWriteAndFlush\n                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/internal/wasi/io/v0.2.0/streams/streams.wit.go:266:39              - runtime.putchar\n                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/runtime/runtime_tinygowasmp2.go:29:38\n    2: 0x24e1 - runtime.printnl\n                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/runtime/print.go:290:9\n    3: 0x14b0 - runtime._panic\n                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/runtime/panic.go:64:9\n    4: 0x36b6 - (*github.com/ydnar/wasm-tools-go/cm.result[github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse, github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse, github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode]).validate[github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode]\n                    at /Users/rajatjindal/go/pkg/mod/github.com/ydnar/wasm-tools-go@v0.0.0-20240511034548-aa782301db37/cm/result.go:90:8\n    5:  0x4d8 - github.com/ydnar/wasm-tools-go/cm.Err[github.com/ydnar/wasm-tools-go/cm.OKResult[github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse, github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode] github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode]\n                    at /Users/rajatjindal/go/pkg/mod/github.com/ydnar/wasm-tools-go@v0.0.0-20240511034548-aa782301db37/cm/result.go:144:12\n    6:  0x47b - main.init#1$1\n                    at /Users/rajatjindal/go/src/github.com/rajatjindal/wasip2-http-incoming-handler/main.go:11:70              - wasi:http/incoming-handler@0.2.0#handle\n                    at /Users/rajatjindal/go/src/github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/incoming-handler/incoming-handler.wit.go:36:8",
    source: "unknown handle index 0",
}    
error: hyper::Error(User(Service), guest never invoked `response-outparam::set` method: error while executing at wasm backtrace:
    0: 0x1332a - wit-component:shim!indirect-wasi:io/streams@0.2.0-[method]output-stream.blocking-write-and-flush
    1: 0x25f5 - (internal/wasi/io/v0.2.0/streams.OutputStream).BlockingWriteAndFlush
                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/internal/wasi/io/v0.2.0/streams/streams.wit.go:266:39              - runtime.putchar
                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/runtime/runtime_tinygowasmp2.go:29:38
    2: 0x24e1 - runtime.printnl
                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/runtime/print.go:290:9
    3: 0x14b0 - runtime._panic
                    at /Users/rajatjindal/go/src/github.com/tinygo-org/tinygo/build/release/tinygo/src/runtime/panic.go:64:9
    4: 0x36b6 - (*github.com/ydnar/wasm-tools-go/cm.result[github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse, github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse, github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode]).validate[github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode]
                    at /Users/rajatjindal/go/pkg/mod/github.com/ydnar/wasm-tools-go@v0.0.0-20240511034548-aa782301db37/cm/result.go:90:8
    5:  0x4d8 - github.com/ydnar/wasm-tools-go/cm.Err[github.com/ydnar/wasm-tools-go/cm.OKResult[github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse, github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode] github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.OutgoingResponse github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types.ErrorCode]
                    at /Users/rajatjindal/go/pkg/mod/github.com/ydnar/wasm-tools-go@v0.0.0-20240511034548-aa782301db37/cm/result.go:144:12
    6:  0x47b - main.init#1$1
                    at /Users/rajatjindal/go/src/github.com/rajatjindal/wasip2-http-incoming-handler/main.go:11:70              - wasi:http/incoming-handler@0.2.0#handle
                    at /Users/rajatjindal/go/src/github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/incoming-handler/incoming-handler.wit.go:36:8

Caused by:
    unknown handle index 0)

```