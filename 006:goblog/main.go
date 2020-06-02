package main 

import (
	"time"
	"net/http"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Blog struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Content string `json:"content"`
	CreatedOn time.Time `json:"createdon"`
}

var blogStore = make(map[string]Blog)

var id int = 0 

// HTTP Post - /api/blogs

func PostBlogHandler(w http.ResponseWriter, r *http.Request) {
	var blog Blog 
	// Decode incoming Json
	err := json.NewDecoder(r.Body).Decode(&Blog)
	if err != nil {
		panic(err)
	}
	blog.CreatedOn = time.Now()
	id++ 
	k := strconv.Itoa(id)
	blogStore[k] = blog 
	j, err := json.Marshal(blog)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// HTTP Get - /api/blogs 

func GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	var blogs []Blog
	for _, v := range blogStore{
		blogs = append(blogs, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(blogs)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Put - /api/blog/{id}

func PutBlogHandler(w http.ResponseWriter, r *http.Request) {
	var err error 
	vars := mux.Vars(r)
	k := vars["id"]
	var blogUpd Blog 
	// Decode incoming blog JSON 
	err = json.NewDecoder(r.Body).Decode(&blogUpd)
	if err != nil {
		panic(err)
	}
	if blog, ok := blogStore[k]; ok {
		blogUpd.CreatedOn = blog.CreatedOn
		// delete existing item and add the updated item 
		delete(blogStore, k)
		blogStore[k] = blogUpd
	} else {
		log.Printf("COuld not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	// Remove from store 
	if _, ok := blogStore[k]; ok {
		// delete existing item 
		delete(blogStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/blogs", GetBlogHandler).Methods("GET")
	r.HandleFunc("/api/blogs", PostBlogHandler).Methods("POST")
	r.HandleFunc("/api/blogs/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/blogs/{id}", DeleteBlogHandler).Methods("DELETE")

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}