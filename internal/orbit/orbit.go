package orbit

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// global state and http helper methods
type Orbit struct{}

func Launch(r http.Handler) {
	fmt.Printf("🔥 Launching at http://%s:%s 🔥\n", os.Getenv("HOST"), os.Getenv("PORT"))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
	if err != nil {
		log.Fatal(err)
	}
}

func (o *Orbit) Text(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

func (o *Orbit) HTML(w http.ResponseWriter, code int, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(html))
}

// func (o *Orbit) Render(w http.ResponseWriter, name string, code int, data pongo2.Context) {
// 	template := pongo2.Must(pongo2.FromCache(path.Join("web/view", name) + ".ehtml"))

// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(code)
// 	template.ExecuteWriter(data, w)
// }

// func (o *Orbit) StaticRender(w http.ResponseWriter, hs *static.HTMLStore, code int, slug string) {
// 	post, ok := hs.StaticPostMap[slug]
// 	if !ok {
// 		o.Error(w, 404, fmt.Sprintf("%s is not found", slug))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(code)
// 	w.Write([]byte(post))
// }

func (o *Orbit) Error(w http.ResponseWriter, code int, errorMessage string) {
	http.Error(w, errorMessage, code)
}
