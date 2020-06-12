package main_test

import "testing"

func TestAddition(t *testing.T){
    got := 2 + 2
    expected := 4
    if got != expected{
        t.Errorf("Not expected result. Got: '%v', wanted: '%v'", got, expected)
    }
}

func TestLoop(t *testing.T){
    data := [{"id": 1, "first_name": "Maurise", "last_name": "Shieldon", "email": "mshieldon0@squidoo.com", "ip_address": "192.57.232.111", "latitude": 34.003135, "longitude": -117.7228641}, {"id": 2, "first_name": "Bendix", "last_name": "Halgarth", "email": "bhalgarth1@timesonline.co.uk", "ip_address": "4.185.73.82", "latitude": -2.9623869, "longitude": 104.7399789}, {"id": 3, "first_name": "Meghan", "last_name": "Southall", "email": "msouthall2@ihg.com", "ip_address": "21.243.184.215", "latitude": "15.45033", "longitude": "44.12768"}]

    got := 4
    expected := 4
    if got == expected{
        t.Errorf("Not expected result. Got: '%v', wanted: '%v'", got, expected)
    }
}
