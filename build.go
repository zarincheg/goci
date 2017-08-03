package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// New An object containing information about the state of the reference after the push.
type New struct {
	Type string
	Name string
}

// BitbucketPush The object with information about a git push.
type BitbucketPush struct {
	Actor      string
	Repository string
	Push       New
}

// GitPushHandler Handler for the Bitbucket push webhook.
func GitPushHandler(w http.ResponseWriter, r *http.Request) {
	var payload BitbucketPush
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &payload)

	var buf bytes.Buffer
	logger := log.New(&buf, "info: ", log.Lshortfile)
	logger.Print(payload.Repository)

	fmt.Print(&buf)
}

func main() {
	http.HandleFunc("/build", GitPushHandler)
	http.ListenAndServe(":8031", nil)
}
