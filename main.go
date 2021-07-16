package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"voting/usecase"
)

type sEvent struct {
	Id_event   int
	Nama_event string
	Status     string
	Tujuan     string
}

type nCalon struct {
	Idcalon     int
	Idevent     int
	Namacalon   string
	Jumlahsuara string
	Visi        string
	Misi        string
	TTL         string
	Hobi        string
}

var data_event = []sEvent{
	{1, "Pemilihan Osis SMA ", "Buka", "Pemilihan Osis SMA 1 Semarang"}, {2, "Pemilihan Osis SMP", "Buka", "Pemilihan Osis SMP 2 Semarang"},
	{3, "Pemilihan MPK", "Buka", "Pemilihan MPK SMA 1 Bandung"},
}

var data_calon = []nCalon{
	{2, 1, "Feb", "10", "Jalan Terus", " ", "Bandung, 20 Januari 2001", "Belajar"},
}

func handlerBerandaPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("html", "beranda.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Title string
	}{"Hi VOTERS !!! "}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerEventPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("html", "event.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_user []sEvent
	}{data_event}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerEditEventPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	namaevent := r.FormValue("nama_event")
	var filepath = path.Join("html", "editevent.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Namaevent string
	}{namaevent}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerTambahEventPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("html", "tambahevent.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_user []sEvent
	}{data_event}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerEditCalonPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	namacalon := r.FormValue("nama_calon")
	var filepath = path.Join("html", "editcalon.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Namacalon string
	}{namacalon}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerTambahCalonPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("html", "tambahcalon.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data_view = struct {
		Data_user []nCalon
	}{data_calon}

	err = tmpl.Execute(w, data_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("asset"))))

	http.HandleFunc("/berandapage", handlerBerandaPage)
	// http.HandleFunc("/eventpage", handlerEventPage)
	// http.HandleFunc("/delete", handlerDelete)
	http.HandleFunc("/eventpage", usecase.HandlerEventPage)
	//http.HandleFunc("/editeventpage", usecase.HandlerEditEventPage)
	http.HandleFunc("/editeventpage", handlerEditEventPage)
	http.HandleFunc("/editevent", usecase.HandlerUserEditEvent)
	http.HandleFunc("/delete", usecase.HandlerUserDelete)
	http.HandleFunc("/tambahevent", usecase.HandlerTambahEvent)
	http.HandleFunc("/tambaheventpage", handlerTambahEventPage)
	http.HandleFunc("/calonpage", usecase.HandlerCalonPage)
	http.HandleFunc("/editcalonpage", handlerEditCalonPage)
	http.HandleFunc("/editcalon", usecase.HandlerUserEditCalon)
	http.HandleFunc("/deletecalon", usecase.HandlerUserDeleteCalon)
	http.HandleFunc("/tambahcalon", usecase.HandlerTambahCalon)
	http.HandleFunc("/tambahcalonpage", handlerTambahCalonPage)
	// http.HandleFunc("/profilcalonpage", handlerUserProfilCalonPage)
	http.HandleFunc("/profilcalon", usecase.HandlerUserProfilCalon)
	http.HandleFunc("/votepage", usecase.HandlerUserVotePage)
	http.HandleFunc("/vote", usecase.HandlerUserVote)
	http.HandleFunc("/livesuarapage", usecase.HandlerUserLiveSuara)
	http.HandleFunc("/kodepage", usecase.HandlerKodePage)
	http.HandleFunc("/kodeeventpage", usecase.HandlerKodeEventPage)

	fmt.Println("starting web server at http://localhost:7008/")
	http.ListenAndServe("localhost:7008", nil)

}
