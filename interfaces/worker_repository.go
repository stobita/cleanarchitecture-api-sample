package interfaces

import "github.com/stobita/cleanarchitecture-api-sample/domain"

type WorkerRepository struct {
	SqlHandler
}

func (repo *WorkerRepository) Store(worker domain.Worker) (id int, err error) {
	result, err := repo.Execute("INSERT INTO worker (name) VALUE (?)", worker.Name)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *WorkerRepository) FindByID(identifier int) (worker domain.Worker, err error) {
	row, err := repo.Query("SELECT id, name FROM worker WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	worker = domain.Worker{
		ID:   id,
		Name: name,
	}
	return
}

func (repo *WorkerRepository) FindAll() (workers domain.Workers, err error) {
	rows, err := repo.Query("SELECT id, name FROM worker")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			continue
		}
		worker := domain.Worker{
			ID:   id,
			Name: name,
		}
		workers = append(workers, worker)
	}
	return
}
