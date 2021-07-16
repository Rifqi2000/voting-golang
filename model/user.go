package model

type Event struct {
	Idevent   int    `json:"id_event"`
	Namaevent string `json:"nama_event"`
	Status    string `json:"status"`
	Tujuan    string `json:"tujuan"`
}

type Calon struct {
	Idcalon     int    `json:"id_calon"`
	Idevent     int    `json:"id_event"`
	Namacalon   string `json:"nama_calon"`
	Jumlahsuara int    `json:"jumlah_suara"`
	Visi        string `json:"visi"`
	Misi        string `json:"misi"`
	TTL         string `json:"ttl"`
	Hobi        string `json:"hobi"`
}
