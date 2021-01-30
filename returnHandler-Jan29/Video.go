package main

type Video struct {
	Id          string `jason:"Id"`
	User        string `jason:"User"`
	Date        string `jason:"Date"`
	Subject     string `jason:"Subject"`
	Description string `jason:"Description"`
	Fulefilled  string `jason:"Fulefilled"`
}
