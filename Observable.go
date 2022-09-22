package main

type Observable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	sendAll()
}

func (j *JobSite) subscribe(subscriber Observer) {
	j.subscribers = append(j.subscribers, subscriber)
}

func (j *JobSite) unsubscribe(subscriber Observer) {
	j.subscribers = removeSubscriber(j.subscribers, subscriber)
}

func removeSubscriber(subscribers []Observer, subscriberToRemove Observer) []Observer {
	for i, curSubscriber := range subscribers {
		if curSubscriber == subscriberToRemove {
			return append(subscribers[:i], subscribers[i+1:]...)
		}
	}
	return subscribers
}

func (j *JobSite) sendAll() {
	for _, subscriber := range j.subscribers {
		subscriber.handleEvent(j.vacancies)
	}
}
