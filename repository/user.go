package repository

import (
	"fmt"
	"voting/model"
)

var tname_event string = prefixDb("event")
var tname_calon string = prefixDb("calon")

func GetEventPage() ([]model.Event, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT nama_event, tujuan,id_event,status  FROM " + tname_event
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	//rows, err := db.Query("SELECT id, name, age, grade FROM " + tname_student + " WHERE age=?", age)
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data_event := []model.Event{}
	for rows.Next() {
		each := model.Event{}
		err := rows.Scan(
			&each.Namaevent, &each.Tujuan, &each.Idevent, &each.Status,
		)
		if err != nil {
			return nil, err
		}

		data_event = append(data_event, each)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data_event, nil
}

// func GetEditEventPage() ([]model.Event, error) {
// 	db, err := connectDb()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	query := "SELECT nama_event FROM " + tname_event
// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//rows, err := db.Query("SELECT id, name, age, grade FROM " + tname_student + " WHERE age=?", age)
// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	data_event := []model.Event{}
// 	for rows.Next() {
// 		each := model.Event{}
// 		err := rows.Scan(
// 			&each.Namaevent, &each.Tujuan,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		data_event = append(data_event, each)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return data_event, nil
// }

func GetUserEditEvent(in_namaevent string, in_statusevent string, in_tujuanevent string) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "UPDATE " + tname_event + " SET status=? WHERE nama_event=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_statusevent, in_namaevent,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserDelete(in_idevent int) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "DELETE FROM `tb_event` WHERE id_event=? "
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_idevent,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserTambahEvent(in_namaevent string, in_idevent int, in_tujuanevent string, in_statusevent string) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "INSERT INTO tb_event (nama_event,id_event,tujuan,status) VALUES (?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_namaevent, in_idevent, in_tujuanevent, in_statusevent,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetCalonPage() ([]model.Calon, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT nama_calon, visi,id_calon  FROM " + tname_calon
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	//rows, err := db.Query("SELECT id, name, age, grade FROM " + tname_student + " WHERE age=?", age)
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data_calon := []model.Calon{}
	for rows.Next() {
		each := model.Calon{}
		err := rows.Scan(
			&each.Namacalon, &each.Visi, &each.Idcalon,
		)
		if err != nil {
			return nil, err
		}

		data_calon = append(data_calon, each)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data_calon, nil
}

func GetUserEditCalon(in_namacalon string, in_visi string, in_misi string) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "UPDATE " + tname_calon + " SET visi=?, misi=? WHERE nama_calon=? "
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_visi, in_misi, in_namacalon,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserDeleteCalon(in_idcalon int) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "DELETE FROM `tb_calon` WHERE id_calon=? "
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_idcalon,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserTambahCalon(in_namacalon string, in_idevent int, in_visi string, in_misi string, in_ttl string, in_hobi string) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "INSERT INTO tb_calon (nama_calon,id_event,visi,misi,ttl,hobi) VALUES (?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_namacalon, in_idevent, in_visi, in_misi, in_ttl, in_hobi,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetProfilCalonPage(in_idcalon int) ([]model.Calon, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT nama_calon,visi,misi,ttl,hobi  FROM " + tname_calon + " WHERE id_calon=?"

	stmt, err := db.Prepare(query)
	if err != nil {

		return nil, err

	}
	defer stmt.Close()

	each := model.Calon{}
	err = stmt.QueryRow(in_idcalon).Scan(&each.Namacalon, &each.Visi, &each.Misi, &each.TTL, &each.Hobi)
	if err != nil {
		// fmt.Println(in_idcalon)
		return nil, err

	}

	data_calon := []model.Calon{}
	data_calon = append(data_calon, each)
	return data_calon, nil
}

// SELECT tb_calon.nama_calon, tb_calon.id_calon, tb_calon.visi FROM tb_event INNER JOIN tb_calon ON tb_event.id_event = tb_calon.id_event WHERE tb_event.id_event=2
func GetUserVote(in_idevent int) ([]model.Calon, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT tb_calon.nama_calon, tb_calon.id_calon, tb_calon.visi FROM tb_event INNER JOIN tb_calon ON tb_event.id_event = tb_calon.id_event WHERE tb_event.id_event=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	//rows, err := db.Query("SELECT id, name, age, grade FROM " + tname_student + " WHERE age=?", age)
	rows, err := stmt.Query(in_idevent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data_calon := []model.Calon{}
	for rows.Next() {
		each := model.Calon{}
		err := rows.Scan(
			&each.Namacalon, &each.Idcalon, &each.Visi,
		)
		if err != nil {
			return nil, err
		}

		data_calon = append(data_calon, each)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data_calon, nil
}

func GetUserCekStatus(in_idevent int) ([]model.Event, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT status  FROM " + tname_event + " WHERE id_event=?"
	stmt, err := db.Prepare(query)
	if err != nil {

		return nil, err

	}
	defer stmt.Close()
	fmt.Println("TEST", in_idevent)
	each := model.Event{}
	err = stmt.QueryRow(in_idevent).Scan(&each.Status)
	if err != nil {
		// fmt.Println(in_idcalon)
		return nil, err

	}
	fmt.Println(each, in_idevent)
	data_event := []model.Event{}
	data_event = append(data_event, each)
	return data_event, nil

}

func GetUserVoteLogic(in_idcalon int) error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	// query := "SELECT username, password, name FROM " + tname_user
	query := "UPDATE " + tname_calon + " SET jumlah_suara=jumlah_suara+1 WHERE id_calon=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		in_idcalon,
	)
	if err != nil {
		return err
	}
	return nil
}

// SELECT tb_calon.nama_calon, tb_calon.id_calon, tb_calon.jumlah_suara FROM tb_calon INNER JOIN tb_event ON tb_calon.id_event = tb_event.id_event
func GetUserLiveSuara(in_idevent int) ([]model.Calon, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT tb_calon.nama_calon, tb_calon.id_calon, tb_calon.jumlah_suara FROM tb_event INNER JOIN tb_calon ON tb_event.id_event = tb_calon.id_event WHERE tb_event.id_event=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	//rows, err := db.Query("SELECT id, name, age, grade FROM " + tname_student + " WHERE age=?", age)
	rows, err := stmt.Query(in_idevent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data_calon := []model.Calon{}
	for rows.Next() {
		each := model.Calon{}
		err := rows.Scan(
			&each.Namacalon, &each.Idcalon, &each.Jumlahsuara,
		)
		if err != nil {
			return nil, err
		}

		data_calon = append(data_calon, each)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data_calon, nil
}

func GetKodePage() ([]model.Event, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT id_event, nama_event  FROM " + tname_event
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	//rows, err := db.Query("SELECT id, name, age, grade FROM " + tname_student + " WHERE age=?", age)
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data_event := []model.Event{}
	for rows.Next() {
		each := model.Event{}
		err := rows.Scan(
			&each.Idevent, &each.Namaevent,
		)
		if err != nil {
			return nil, err
		}

		data_event = append(data_event, each)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data_event, nil
}
