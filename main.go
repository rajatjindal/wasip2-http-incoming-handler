package main

import (
	incominghandler "github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/incoming-handler"
	"github.com/rajatjindal/wasip2-http-incoming-handler/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

func init() {
	incominghandler.Handle = func(request types.IncomingRequest, responseOut types.ResponseOutparam) {
		err := cm.Err[cm.ErrResult[types.OutgoingResponse, types.ErrorCode]](types.ErrorCodeConfigurationError())
		types.ResponseOutparamSet(cm.ResourceNone, err)
	}
}

func main() {}
