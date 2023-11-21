package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go-training/handler"
	"go-training/model"
	"net/http"
	"github.com/go-chi/chi"
	//"github.com/casbin/chi-authz"
	"github.com/casbin/casbin"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "user=postgres password=123456 dbname=postgres")
	if err != nil {
		panic(err.Error())
	
	}
	//var Own
	db.AutoMigrate(&model.Owners{})
}


func main() {
	r := chi.NewRouter()

	// load the casbin model and policy from files, database is also supported.
	e := casbin.NewEnforcer("model.conf", "policy.csv")
	r.Use(Authorizer(e))

	// define endpoints and associate with handler
	r.Get("/object", handler.GetOwners(db))
	r.Get("/object/{id}", handler.GetOwnersbyId(db))
	r.Post("/object", handler.CreateOwner(db))
	r.Patch("/object/{id}", handler.UpdateOwner(db))
	r.Delete("/object/{id}", handler.DeleteOwner(db))

	// setup http server on port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, password, _ := r.BasicAuth()
			method := r.Method
			path := r.URL.Path
			if e.Enforce(user, password, path, method) {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, http.StatusText(403), 403)
			}
		}

		return http.HandlerFunc(fn)
	}
}