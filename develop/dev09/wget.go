package main

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type MSet struct {
	mx sync.RWMutex
	mp map[string]bool
}

func NewSet() *MSet {
	return &MSet{
		mp: make(map[string]bool),
	}
}

func (m *MSet) Set(key string) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.mp[key] = true
}

func (m *MSet) Check(key string) bool {
	m.mx.RLock()
	defer m.mx.RUnlock()
	_, ok := m.mp[key]
	return ok
}

func getSite(urlStr string) {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalln(err)
	}

	visitedMap := NewSet()

	urlHost := u.Hostname()

	//err = os.Mkdir(urlHost, os.ModePerm)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	c := colly.NewCollector(
		colly.AllowedDomains(urlHost),
		// Parallelism
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if ok := visitedMap.Check(link); !ok {
			c.Visit(e.Request.AbsoluteURL(link))
			visitedMap.Set(link)
		}
	})
	//var mx sync.RWMutex
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			fmt.Printf("Visiting: %s\n", r.Request.URL)
			respUrl := r.Request.URL
			dirname := respUrl.Hostname() + respUrl.Path
			s := strings.Split(respUrl.Path, "/")
			filename := s[len(s)-1]
			if filepath.Ext(filename) == "" {
				filename = "index.html"
			}
			fileNameFull := filepath.Join(dirname, filename)
			if _, err := os.Stat(dirname); errors.Is(err, os.ErrNotExist) {
				os.MkdirAll(dirname, os.ModePerm)
			}
			if _, err := os.Stat(fileNameFull); errors.Is(err, os.ErrNotExist) {
				err = r.Save(fileNameFull)
				if err != nil {
					log.Println(err)
				} else {
					log.Println("Created file:", fileNameFull)
				}
			}
		}
	})

	if err := c.Visit(u.String()); err != nil {
		log.Fatalln(err)
	}
	c.Wait()
}
