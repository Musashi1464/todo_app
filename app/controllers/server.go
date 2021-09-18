package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"todo_app_heroku/app/models"
	"todo_app_heroku/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// このfuncを用いてアクセス制限を実現する
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie") // route_auth.goのauthenticate()で指定したcookieの名前(Name)
	if err == nil {
		sess = models.Session{UUID: cookie.Value} // sessionのUUIDがcookieのValueに登録されている
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session") // 無効なセッション
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	// 返り値のHandlerFuncはハンドル関数である
	return func(w http.ResponseWriter, r *http.Request) {
		// /todo/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2]) // q[2]はtodoのidが入っている
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

// 第二引数はデフォルトのマルチプレクサを使うため、nilにする
// 登録されていないアドレスにアクセスする際に、"404 page not found"を表示する
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
}
