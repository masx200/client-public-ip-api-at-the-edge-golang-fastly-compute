// https://developer.fastly.com/solutions/examples/client-public-ip-api-at-the-edge/
package main

import (
	"context"
	// "fmt"
	// "io"
	"encoding/json"
	"log"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

// BackendName is the name of our service backend.
// const BackendName = "origin_0"

func main() {
	fsthttp.ServeFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		// This requires your service to be configured with a backend
		// named "origin_0" and pointing to "https://http-me.glitch.me".

		// If the URL path matches the path below, then return the client IP as a JSON response.
		// if r.URL.Path == "/api/clientIP" {
		// Get client IP address
		ClientIP := r.RemoteAddr
		resp := make(map[string]any)

		resp["method"] = r.Method
		resp["url"] = r.URL.String()

		resp["client"] = map[string]any{"address": ClientIP}

		resp["headers"] = r.Header
		// resp]["Address"] = ClientIP
		JsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(JsonResp)
		w.Header().Set("Content-Type", "application/json")
		// Return 200 OK response
		w.WriteHeader(fsthttp.StatusOK)
		//return

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
