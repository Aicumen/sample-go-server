package api

import (
	"encoding/json"
	"net/http"

	"github.com/akshay111meher/sample-go-server/go/db"
	keyMgmt "github.com/akshay111meher/sample-go-server/go/key_management"
)

var apikeyQueryParam = "apikey"

// LatestGet ...
func LatestGet(w http.ResponseWriter, r *http.Request) {
	apikey := r.URL.Query().Get(apikeyQueryParam)
	user, err := keyMgmt.GetValueFromKey(apikey)
	if err != nil || user == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Probablty you are not logged in"))
	} else {
		posts, err := db.LatestPosts()
		if err != nil || len(posts) < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error or No Posts"))
		} else {
			w.WriteHeader(http.StatusOK)
			postsMarshall, err := json.Marshal(posts)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error. Marshal Error"))
			}
			w.Write(postsMarshall)
		}

	}
}
