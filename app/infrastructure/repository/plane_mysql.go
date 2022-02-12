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

}
