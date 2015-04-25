package config

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	cfg, err := New("test.ini")
	if err != nil {
		t.Error("Got error on parse " + err.Error())
	} else {
		log.Printf("Defaults: %+v\n", cfg.Defaults)
		log.Println("Pools:")
		for k, v := range cfg.Pool {
			log.Printf("%s ==> %+v\n", k, *v)
		}
		log.Println("Jobs: ")
		for k, v := range cfg.Job {
			log.Printf("%s ==> %+v\n", k, *v)
		}
	}
}
