package usecase

import "github.com/stobita/cleanarchitecture-api-sample/domain"

type WorkerRepository interface {
	Store(domain.Worker) (int, error)
	FindByID(int) (domain.Worker, error)
	FindAll() (domain.Workers, error)
}

type WorkerInteractor struct {
	WorkerRepository
}

func (interactor *WorkerInteractor) Add(worker domain.Worker) (err error) {
	_, err = interactor.WorkerRepository.Store(worker)
	return
}

func (interactor *WorkerInteractor) Workers() (workers domain.Workers, err error) {
	workers, err = interactor.WorkerRepository.FindAll()
	return
}

func (interactor *WorkerInteractor) WorkerByID(identifier int) (worker domain.Worker, err error) {
	worker, err = interactor.WorkerRepository.FindByID(identifier)
	return
}
