package repository

import (
	"database/sql"
	"fmt"

	"github.com/viniciusarambul/go-flight/app/entity"
)

type PlaneMySQLRepository struct {
	Db *sql.DB
}

func (p PlaneMySQLRepository) Insert(plane entity.Plane) error {
	stmt, err := p.Db.Prepare(`Insert into planes(id, model, desription, status) values (?,?,?,?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		plane.ID,
		plane.Model,
		plane.Description,
		plane.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PlaneMySQLRepository) Select(plane entity.Plane) error {
	fmt.Sprintln("dentro do select: ")
	err := p.Db.QueryRow("select id, model, description, status from planes").Scan(&plane.ID, &plane.Model, &plane.Description, &plane.Status)

	if err != nil {
		return err
	}

	return nil
	/*switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("id is %d, modelo: %v, descricao: %v, status:%v\n", &plane.ID, &plane.Model, &plane.Description, &plane.Status)
	}
	/*
		for result.Next() {
			var plane entity.Plane
			err = result.Scan(&plane.ID, &plane.Model, &plane.Description, &plane.Status)
		}
		log.Printf(plane.ID)
	*/

}
