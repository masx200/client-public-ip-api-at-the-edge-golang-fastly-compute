// https://developer.fastly.com/solutions/examples/client-public-ip-api-at-the-edge/
package main

import (
	"context"
	// "io"
	// "net/http"
	// "net/url"

	// "fmt"
	// "io"
	"encoding/json"
	"log"

	"github.com/fastly/compute-sdk-go/fsthttp"
	// "github.com/quic-go/quic-go/http3"
)

// BackendName is the name of our service backend.
// const BackendName = "origin_0"

func main() {
	fsthttp.ServeFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {

		w.Header().Add("Alt-Svc", "h3=\":443\"; ma=86400")
		w.Header().Add("Alt-Svc", "h2=\":443\"; ma=86400")
		w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// This requires your service to be configured with a backend
		// named "origin_0" and pointing to "https://http-me.glitch.me".
		// if r.URL.Path == "/api/http" {

		// 	o := &fsthttp.BackendOptions{}
		// 	o.UseSSL(true).HostOverride("fastly-compute-hello-world-javascript.edgecompute.app").SNIHostname("fastly-compute-hello-world-javascript.edgecompute.app")
		// 	b, err := fsthttp.RegisterDynamicBackend(
		// 		"fastly-compute-hello-world-javascript.edgecompute.app",
		// 		"fastly-compute-hello-world-javascript.edgecompute.app",
		// 		o,
		// 	)
		// 	if err != nil {
		// 		log.Printf("Error happened in http request. Err: %s", err)
		// 		w.WriteHeader(fsthttp.StatusBadGateway)
		// 		w.Write([]byte("Bad Gateway" + "\n" + err.Error()))
		// 		return
		// 	}
		// 	req, err := fsthttp.NewRequest(r.Method, "https://fastly-compute-hello-world-javascript.edgecompute.app/", nil)

		// 	if err != nil {
		// 		log.Printf("Error happened in http request. Err: %s", err)
		// 		w.WriteHeader(fsthttp.StatusBadGateway)
		// 		w.Write([]byte("Bad Gateway" + "\n" + err.Error()))
		// 		return
		// 	}
		// 	req.Header = r.Header
		// 	resp, err := req.Send(ctx, b.Name())
		// 	if err != nil {
		// 		log.Printf("Error happened in http request. Err: %s", err)
		// 		w.WriteHeader(fsthttp.StatusBadGateway)
		// 		w.Write([]byte("Bad Gateway" + "\n" + err.Error()))
		// 		return
		// 	}
		// 	defer io.Copy(w, resp.Body)
		// 	w.WriteHeader(resp.StatusCode)

		// 	for k, v := range resp.Header {
		// 		for _, h := range v {
		// 			w.Header().Add(k, h)
		// 		}

		// 	}
		// 	return
		// }
		// If the URL path matches the path below, then return the client IP as a JSON response.
		// if r.URL.Path == "/api/clientIP" {
		// Get client IP address
		ClientIP := r.RemoteAddr
		resp := make(map[string]any)
		resp["request"] = map[string]any{
			"method":  r.Method,
			"url":     r.URL.String(),
			"headers": r.Header,
		}

		resp["client"] = map[string]any{"address": ClientIP}

		// resp]["Address"] = ClientIP
		JsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Header().Set("Content-Type", "application/json")
		// Return 200 OK response
		w.WriteHeader(fsthttp.StatusOK)
		//return
		w.Write(JsonResp)
		// }
		// resp, err := r.Send(ctx, BackendName)
		// if err != nil {
		// 	w.WriteHeader(fsthttp.StatusBadGateway)
		// 	fmt.Fprintln(w, err.Error())
		// 	return
		// }

		// w.Header().Reset(resp.Header)
		// w.WriteHeader(resp.StatusCode)
		// io.Copy(w, resp.Body)
	})
}
