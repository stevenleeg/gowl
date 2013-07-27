package gowl

import (
    "net/http"
    "net/url"
    "strconv"
    "errors"
)

const (
    API_PREFIX = "https://api.prowlapp.com/publicapi/"
)

type Gowl struct {
    ApiKey string
}

type Notification struct {
    Priority int   
    Url, Application, Event, Description string
}

func (this *Notification) Validate() error {
    if this.Application == "" {
        return errors.New("Notification application field must be specified")
    }

    if this.Event == "" {
        return errors.New("Notification event field must be specified")
    }

    if this.Description == "" {
        return errors.New("Notification description field must be specified")
    }

    return nil
}

func New(api_key string) *Gowl {
    gowl := new(Gowl)
    gowl.ApiKey = api_key
    return gowl
}

func (this *Gowl) Add(notification *Notification) error {
    // Check to see if we need to use default values on the notification struct
    if err := notification.Validate(); err != nil {
        return errors.New("Notification validation failed: " + err.Error())
    }

    post_data := url.Values{
        "apikey": { this.ApiKey },
        "priority": { strconv.FormatInt(int64(notification.Priority), 10) },
        "url": { notification.Url },
        "application": { notification.Application },
        "event": { notification.Event },
        "description": { notification.Description },
    }

    _, err := http.PostForm(API_PREFIX + "add", post_data)

    if err != nil {
        return errors.New("Notification add failed: " + err.Error())
    }

    return nil
}
