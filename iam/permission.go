package iam

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	//Actions 每個RESTful APIs的動作
	Actions *actionEntries
)

const (
	awsAction = "Action"
)

func init() {
	Actions = &actionEntries{
		entries: make(map[string]*ActionEntry),
	}
}

//GinIRouter ...
type GinIRouter func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

type actionEntries struct {
	entries map[string]*ActionEntry
}

//ActionEntry ...
type ActionEntry struct {
	Name          string
	Administrator bool
	IsAWSAction   bool
	Action        string
}

func (a *actionEntries) AddAction(name, action string, isAWSAction, administorator bool) {
	if v, ok := a.entries[name]; ok {
		panic(fmt.Sprintf("Duplicate Action Entry: %+v", v))
	}
	a.entries[name] = &ActionEntry{
		Name:          name,
		Administrator: administorator,
		IsAWSAction:   isAWSAction,
		Action:        action,
	}
}

func (a *actionEntries) GET(rg *gin.RouterGroup, relativePath string,
	handler gin.HandlerFunc, action string, administorator bool) {
	a.AddActionWithRouterGroup(rg, http.MethodGet, relativePath, handler, action, administorator)
}

func (a *actionEntries) POST(rg *gin.RouterGroup, relativePath string,
	handler gin.HandlerFunc, action string, administorator bool) {
	a.AddActionWithRouterGroup(rg, http.MethodPost, relativePath, handler, action, administorator)
}

func (a *actionEntries) DELETE(rg *gin.RouterGroup, relativePath string,
	handler gin.HandlerFunc, action string, administorator bool) {
	a.AddActionWithRouterGroup(rg, http.MethodDelete, relativePath, handler, action, administorator)
}

func (a *actionEntries) PUT(rg *gin.RouterGroup, relativePath string,
	handler gin.HandlerFunc, action string, administorator bool) {
	a.AddActionWithRouterGroup(rg, http.MethodPut, relativePath, handler, action, administorator)
}

func (a *actionEntries) AddActionWithRouterGroup(rg *gin.RouterGroup, httpMethod, relativePath string,
	handler gin.HandlerFunc, action string, administorator bool) {

	name := fmt.Sprintf("%s:%s/%s", httpMethod, rg.BasePath(), strings.TrimPrefix(relativePath, "/"))
	a.AddAction(name, action, true, administorator)

	rg.Handle(httpMethod, relativePath, handler)
}

func (a *actionEntries) AddAWSAction(name, action string, administorator bool) {
	a.AddAction(name, action, true, administorator)
}

func (a *actionEntries) Len() int {
	return len(a.entries)
}

func (a *actionEntries) Get(key string) *ActionEntry {
	return a.entries[key]
}

func (a *actionEntries) GetMap() map[string]*ActionEntry {
	return a.entries
}

func (a *actionEntries) HasActionPrefix(prefix string) (actions []string) {
	for _, v := range a.GetMap() {
		if !strings.HasPrefix(v.Action, prefix) {
			actions = append(actions, v.Action)
		}
	}
	return
}

func (a *actionEntries) GetActions() (actions []string) {
	for _, v := range a.GetMap() {
		actions = append(actions, v.Action)

	}
	return
}

func (a *actionEntries) GetNames() (names []string) {
	for k := range a.GetMap() {
		names = append(names, k)
	}
	return
}

func (a *actionEntries) HasName(name string) (ok bool) {
	_, ok = a.entries[name]
	return
}

func (a *actionEntries) GenName(c *gin.Context) string {
	return fmt.Sprintf("%s:%s", c.Request.Method, c.FullPath())
}

func (a *actionEntries) Checker(c *gin.Context) (isNoRouter bool, actionEntry *ActionEntry) {
	if !a.HasName(a.GenName(c)) {
		name := ""
		switch c.Request.Method {
		case http.MethodPost:
			c.Request.ParseForm()
			name = c.Request.PostForm.Get(awsAction)
		case http.MethodGet:
			name = c.Request.URL.Query().Get(awsAction)
		}

		if name != "" {
			return false, a.Get(name)
		}
		return true, nil
	}
	return false, a.Get(a.GenName(c))
}
