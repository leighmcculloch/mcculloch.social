package main

import (
	"net/http"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/.well-known/webfinger", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		resource := q.Get("resource")
		if resource == "acct:leigh@mcculloch.social" {
			_, err := w.Write([]byte(`{
	"subject":"acct:leigh@mcculloch.social",
	"aliases":[
		"https://inuh.net/@leigh",
		"https://inuh.net/users/leigh"
	],
	"links":[
		{"rel":"http://webfinger.net/rel/profile-page","type":"text/html","href":"https://inuh.net/@leigh"},
		{"rel":"self","type":"application/activity+json","href":"https://inuh.net/users/leigh"},
		{"rel":"http://ostatus.org/schema/1.0/subscribe","template":"https://inuh.net/authorize_interaction?uri={uri}"}
	]
}`))
			if err != nil {
				panic(err)
			}
		} else {
			w.WriteHeader(404)
		}
	})

	m.HandleFunc("/.well-known/host-meta", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
	<Link rel="lrdd" template="https://mcculloch.social/.well-known/webfinger?resource={uri}"/>
</XRD>`))
		if err != nil {
			panic(err)
		}
	})

	m.HandleFunc("/.well-known/nodeinfo", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{
	"links":[
		{"rel":"http://nodeinfo.diaspora.software/ns/schema/2.0", "href":"https://inuh.net/nodeinfo/2.0"}
	]
}`))
		if err != nil {
			panic(err)
		}
	})

	m.HandleFunc("/nodeinfo/2.0", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{
	"version":"2.0",
	"software":{"name":"mastodon"},
	"protocols":["activitypub"],
	"usage":{"users":{"total":1,"activeMonth":1,"activeHalfyear":1}},
	"openRegistrations":false
}`))
		if err != nil {
			panic(err)
		}
	})

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://inuh.net/@leigh", 301)
	})

	s := http.Server{Addr: ":8000", Handler: m}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
