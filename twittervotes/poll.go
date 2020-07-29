package main

type poll struct {
	Options []string
}

func loadOptions() (options []string, err error) {
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	defer iter.Close()

	var p poll

	for iter.Next(&p) {
		options = append(options, p.Options...)
	}

	return
}
