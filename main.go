package main

func main() {
	var HeadHunterKZ JobSite
	HeadHunterKZ.addVacancy("Senior Data Science Developer")
	Naruto := &Person{"Naruto"}
	HeadHunterKZ.subscribe(Naruto)
	HeadHunterKZ.addVacancy("Junior UI/UX Developer")
	HeadHunterKZ.addVacancy("Middle Go Back-End Developer")

	Eren := &Person{"Eren"}
	HeadHunterKZ.subscribe(Eren)
	HeadHunterKZ.removeVacancy("Junior UI/UX Developer")
	HeadHunterKZ.unsubscribe(Naruto)
	HeadHunterKZ.addVacancy("Junior Business Analytic")

	HeadHunterKZ.sendAll()
}
