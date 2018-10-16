package usecase

import "github.com/stobita/cleanarchitecture-api-sample/domain"

type CompanyRepository interface {
	Store(domain.Company) (int, error)
	FindByID(int) (domain.Company, error)
	FindAll() (domain.Companies, error)
}

type CompanyInteractor struct {
	CompanyRepository
}

func (interactor *CompanyInteractor) Add(company domain.Company, err error) {
	_, err = interactor.CompanyRepository.Store(company)
	return
}

func (interactor *CompanyInteractor) Companies() (companies domain.Companies, err error) {
	companies, err = interactor.CompanyRepository.FindAll()
	return
}

func (interactor *CompanyInteractor) CompanyByID(identifier int) (company domain.Company, err error) {
	company, err = interactor.CompanyRepository.FindByID(identifier)
	return
}
