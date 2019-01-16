package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PostCommentPost ...
func PostCommentPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var commentsInput CommentInput
	err := decoder.Decode(&commentsInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	} else {
		user, err := GetValueFromKey(commentsInput.Apikey)
		if user == "" {
			errdata, _ := json.Marshal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errdata)
		} else {
			err = AddComment(commentsInput)
			fmt.Println(err)
			if err != nil {
				errdata, _ := json.Marshal(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(errdata)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Comment successfully added to the post"))
			}
		}
	}
}

// PostDislikePost ...
func PostDislikePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postInput PostInput
	err := decoder.Decode(&postInput)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
	} else {
		user, err := GetValueFromKey(postInput.Apikey)
		if err != nil || user == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error"))
		} else {
			err = DislikePost(postInput)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post disliked successfully"))
			}
		}

	}
}

// PostLikePost ...
func PostLikePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postInput PostInput
	err := decoder.Decode(&postInput)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
	} else {
		user, err := GetValueFromKey(postInput.Apikey)
		if err != nil || user == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error"))
		} else {
			err = LikePost(postInput)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post liked successfully"))
			}
		}

	}
}

// PostCreatePost ...
func PostCreatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var createPost CreatePost
	err := decoder.Decode(&createPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed creating post"))
	} else {
		// w.WriteHeader(http.StatusOK)
		user, err := GetValueFromKey(createPost.Apikey)
		if err != nil || user == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Probablty you are not logged in"))
		} else {
			var postid = randStringRunes(32)
			var post = Posts{Id: postid, User: user, ImageURL: createPost.ImageURL, Likes: 0, Dislikes: 0, Content: createPost.Content}
			err = AddPost(post)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal error"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post success Post ID: " + postid))
			}
		}
	}

}
