package usecase

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"path"
	"strconv"
	"voting/model"
	"voting/repository"
)

func HandlerEventPage(w http.ResponseWriter, r *http.Request) {
	data_event, err := repository.GetEventPage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "event.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var n_data = 3
	var n_page = int(math.Ceil(float64(len(data_event)) / float64(n_data)))
	var page = make([]int, n_page)
	//fmt.Println(page)
	for i := 0; i < len(page); i++ {
		page[i] = i + 1
	}
	//fmt.Println(page)

	var idx_awal = 0
	var idx_akhir = len(data_event)

	//fmt.Println(r.Host, r.URL)
	param_page := r.URL.Query().Get("page")
	if param_page == "" {
		param_page = "1"
	}
	if param_page != "" {
		int_page, err := strconv.Atoi(param_page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idx_awal = (int_page - 1) * n_data
		idx_akhir = int_page * n_data
		if idx_akhir > len(data_event) {
			idx_akhir = len(data_event)
		}
		// fmt.Println("TAMPILKAN HAL :", int_page, "INDEX AWAL", idx_awal)
		// fmt.Println(data_event[idx_awal:idx_akhir])
	} else {
		// fmt.Println("TAMPILKAN SEMUA")
	}
	// TAMBAHAN CODE

	var data_view = struct {
		Page       []int
		Data_event []model.Event
	}{page, data_event[idx_awal:idx_akhir]}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func HandlerUserEditEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	namaevent := r.FormValue("nama_event")
	tujuan := r.FormValue("tujuan")
	status := r.FormValue("status")
	err := repository.GetUserEditEvent(namaevent, status, tujuan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/eventpage", http.StatusSeeOther)

}

func HandlerUserDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	idevent := r.FormValue("id_event")
	i, err := strconv.Atoi(idevent)
	err = repository.GetUserDelete(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/eventpage", http.StatusSeeOther)

}

func HandlerTambahEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	namaevent := r.FormValue("nama_event")
	idevent := r.FormValue("id_event")
	i, err := strconv.Atoi(idevent)
	tujuanevent := r.FormValue("tujuan")
	statusevent := r.FormValue("status")
	err = repository.GetUserTambahEvent(namaevent, i, tujuanevent, statusevent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/eventpage", http.StatusSeeOther)

}
func HandlerCalonPage(w http.ResponseWriter, r *http.Request) {
	data_calon, err := repository.GetCalonPage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "calon.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var n_data = 3
	var n_page = int(math.Ceil(float64(len(data_calon)) / float64(n_data)))
	var page = make([]int, n_page)
	//fmt.Println(page)
	for i := 0; i < len(page); i++ {
		page[i] = i + 1
	}
	//fmt.Println(page)

	var idx_awal = 0
	var idx_akhir = len(data_calon)

	//fmt.Println(r.Host, r.URL)
	param_page := r.URL.Query().Get("page")
	if param_page == "" {
		param_page = "1"
	}
	if param_page != "" {
		int_page, err := strconv.Atoi(param_page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idx_awal = (int_page - 1) * n_data
		idx_akhir = int_page * n_data
		if idx_akhir > len(data_calon) {
			idx_akhir = len(data_calon)
		}
		// fmt.Println("TAMPILKAN HAL :", int_page, "INDEX AWAL", idx_awal)
		// fmt.Println(data_event[idx_awal:idx_akhir])
	} else {
		// fmt.Println("TAMPILKAN SEMUA")
	}
	// TAMBAHAN CODE

	var data_view = struct {
		Page       []int
		Data_calon []model.Calon
	}{page, data_calon[idx_awal:idx_akhir]}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func HandlerUserEditCalon(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	namacalon := r.FormValue("nama_calon")
	visi := r.FormValue("visi")
	misi := r.FormValue("misi")
	err := repository.GetUserEditCalon(namacalon, visi, misi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/calonpage", http.StatusSeeOther)

}

func HandlerUserDeleteCalon(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	idcalon := r.FormValue("id_calon")
	x, err := strconv.Atoi(idcalon)
	err = repository.GetUserDeleteCalon(x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/calonpage", http.StatusSeeOther)

}
func HandlerTambahCalon(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	namacalon := r.FormValue("nama_calon")
	idevent := r.FormValue("id_event")
	i, err := strconv.Atoi(idevent)
	visi := r.FormValue("visi")
	misi := r.FormValue("misi")
	ttl := r.FormValue("ttl")
	hobi := r.FormValue("hobi")
	err = repository.GetUserTambahCalon(namacalon, i, visi, misi, ttl, hobi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/calonpage", http.StatusSeeOther)

}

func HandlerUserProfilCalon(w http.ResponseWriter, r *http.Request) {
	idcalon := r.FormValue("id_calon")
	x, err := strconv.Atoi(idcalon)
	data_calon, err := repository.GetProfilCalonPage(x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "profilcalon.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_calon []model.Calon
	}{data_calon}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerUserVotePage(w http.ResponseWriter, r *http.Request) {
	idevent := r.FormValue("id_event")
	id, err := strconv.Atoi(idevent)
	data_calon, err := repository.GetUserVote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "vote.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_calon []model.Calon
		Idevent    string
	}{data_calon, idevent}
	fmt.Println(data_view)
	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerUserVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	idcalon := r.FormValue("id_calon")
	idevent := r.FormValue("id_event")
	i, err := strconv.Atoi(idcalon)
	j, err := strconv.Atoi(idevent)
	data_event, err := repository.GetUserCekStatus(j)
	fmt.Println(data_event)
	if (len(data_event)) > 0 {
		if data_event[0].Status != "tutup" {
			err = repository.GetUserVoteLogic(i)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Redirect(w, r, "/eventpage", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/eventpage", http.StatusSeeOther)
		}
	}
	// if data_event[0].Status != "tutup" {
	// 	err = repository.GetUserVoteLogic(i)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// 	http.Redirect(w, r, "/eventpage", http.StatusSeeOther)
	// } else {
	// 	http.Redirect(w, r, "/eventpage", http.StatusSeeOther)
	// }

}

func HandlerUserLiveSuara(w http.ResponseWriter, r *http.Request) {
	idevent := r.FormValue("id_event")
	dd, err := strconv.Atoi(idevent)
	data_calon, err := repository.GetUserLiveSuara(dd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "livesuara.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_calon []model.Calon
	}{data_calon}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerKodePage(w http.ResponseWriter, r *http.Request) {
	data_event, err := repository.GetEventPage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "kode.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_event []model.Event
	}{data_event}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerKodeEventPage(w http.ResponseWriter, r *http.Request) {
	data_event, err := repository.GetEventPage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filepath = path.Join("html", "kodeevent.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_event []model.Event
	}{data_event}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
