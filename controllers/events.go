package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/evanlib/evan_monitor/models"

	"github.com/julienschmidt/httprouter"
)

type EventsController struct {
	Controller
}

var Events EventsController

func (self EventsController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Fprint(w, "Error retrieving events")
		return err
	}

	tpl, err := template.ParseFiles("views/header_footer.html", "views/events.html")
	if err != nil {
		return err
	}
	tpl.Execute(w, events)

	//fmt.Print(events)
	return nil
}
func (self EventsController) CreateAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	bs, err := ioutil.ReadAll(r.Body)
	sbs := string(bs)
	fmt.Println(sbs)
	if err != nil {
		return err
	}
	return nil
}
func (self EventsController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	var evType int8
	if r.FormValue("eventType") == "pos" {
		evType = 1
	} else {
		evType = 2
	}

	ev := models.Event{
		Type:        evType,
		Name:        r.FormValue("eventName"),
		Description: r.FormValue("eventDescription"),
		Created:     time.Now(),
	}
	fmt.Println("Event Name: ", ev.Name)
	err := ev.Save()
	return err
}

func EventTest() {
	ev := new(models.Event)
	ev.Type = 1
	ev.Name = "Test Event"
	ev.Description = "Event Description"
	ev.Created = time.Now()

	//Save Event
	ev.Save()
}