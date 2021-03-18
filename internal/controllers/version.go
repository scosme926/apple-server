package controllers

import (
    "net/http"
)

func (c *Controller) getVersion(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Apple Sever v1.0"))
}
