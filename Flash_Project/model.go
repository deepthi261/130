package mem

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"strings"
)

type model struct {
	Name     string
	State    bool
	Pictures []string
}


