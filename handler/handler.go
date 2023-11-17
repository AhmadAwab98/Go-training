package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go-training/model"
	"encoding/json"
	// "github.com/go-chi/chi"
	// "strconv"
	"sort"
)

type OwnerSorted []model.Owners

func (a OwnerSorted) Len() int {
    return len(a)
}

func (a OwnerSorted) Less(i, j int) bool {
    return a[i].ID < a[j].ID
}

func (a OwnerSorted) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func GetOwners(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Owner []model.Owners

		// using select * from owners
		result := db.Find(&Owner)

		// error then panic
		if result.Error != nil {
			panic(result.Error)
		}

		// sort the rows by id
		sort.Sort(OwnerSorted(Owner))

		// convert to json
		jsonData, _ := json.MarshalIndent(Owner, "", "	")

		// write on website
		w.Write([]byte(jsonData))
	}
}