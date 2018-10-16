package interfaces

import "github.com/stobita/cleanarchitecture-api-sample/domain"

type CompanyRepository struct {
	SqlHandler
}

func (repo *CompanyRepository) Store(company domain.Company) (id int, err error) {
	result, err := repo.Execute("INSERT INTO company (name) VALUE (?)", company.Name)
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

func (repo *CompanyRepository) FindByID(identifier int) (company domain.Company, err error) {
	row, err := repo.Query("SELECT id, name FROM company WHERE id =?", identifier)
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
	company = domain.Company{
		ID:   id,
		Name: name,
	}
	return
}

func (repo *CompanyRepository) FindAll() (companies domain.Companies, err error) {
	rows, err := repo.Query("SELECT id, name FROM company")
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
		company := domain.Company{
			ID:   id,
			Name: name,
		}
		companies = append(companies, company)
	}
	return
}
