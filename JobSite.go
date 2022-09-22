package main

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (j *JobSite) addVacancy(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}

func (j *JobSite) removeVacancy(vacancy string) {
	j.vacancies = removeFromVacancies(j.vacancies, vacancy)
	j.sendAll()
}

func removeFromVacancies(vacancies []string, vacancyToRemove string) []string {
	for i, curVacancy := range vacancies {
		if vacancyToRemove == curVacancy {
			return append(vacancies[:i], vacancies[i+1:]...)
		}
	}
	return vacancies
}
