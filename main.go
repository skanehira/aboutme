package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type SNS struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	ID   string `json:"id,omitempty" yaml:"id,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Language struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Experience string `json:"experience,omitempty" yaml:"experience,omitempty"`
}

type Job struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Experience string `json:"experience,omitempty" yaml:"experience,omitempty"`
}

type Tool struct {
	Editor         []string `json:"editor,omitempty" yaml:"editor,omitempty"`
	Terminal       []string `json:"terminal,omitempty" yaml:"terminal,omitempty"`
	VersionControl []string `json:"version_control,omitempty" yaml:"version_control,omitempty"`
	Browser        []string `json:"browser,omitempty" yaml:"browser,omitempty"`
	CD             []string `json:"cd,omitempty" yaml:"cd,omitempty"`
	Other          []string `json:"other,omitempty" yaml:"other,omitempty"`
}

type Dev struct {
	Language  []Language `json:"language,omitempty" yaml:"language,omitempty"`
	OS        []string   `json:"os,omitempty" yaml:"os,omitempty"`
	DB        []string   `json:"db,omitempty" yaml:"db,omitempty"`
	Framework []string   `json:"framework,omitempty" yaml:"framework,omitempty"`
	Tool      Tool       `json:"tool,omitempty" yaml:"tool,omitempty"`
}

type Status struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Age     int    `json:"age,omitempty" yaml:"age,omitempty"`
	Species string `json:"species,omitempty" yaml:"species,omitempty"`
	Sex     string `json:"sex,omitempty" yaml:"sex,omitempty"`
}

type Profile struct {
	Status     Status `json:"status,omitempty" yaml:"status,omitempty"`
	SNS        []SNS  `json:"sns,omitempty" yaml:"sns,omitempty"`
	JobHistory []Job  `json:"job_history,omitempty" yaml:"job_history,omitempty"`
	Dev        Dev    `json:"dev,omitempty" yaml:"dev,omitempty"`
}

var (
	all  = flag.Bool("all", false, "all of profile")
	stat = flag.Bool("stat", false, "about status")
	sns  = flag.Bool("sns", false, "about sns")
	job  = flag.Bool("job", false, "about job history")
	lang = flag.Bool("lang", false, "about programming language")
	os   = flag.Bool("os", false, "about os")
	db   = flag.Bool("db", false, "about db")
	fw   = flag.Bool("fw", false, "about framework")
	tool = flag.Bool("tool", false, "about tool")
)

func toJson(i interface{}) string {
	j, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	json.Indent(buf, j, "", "    ")
	return buf.String() + "\n"
}

func main() {
	// get user home dir
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	// get yaml
	buf, err := ioutil.ReadFile(filepath.Join(usr.HomeDir, ".aboutme.yaml"))
	if err != nil {
		panic(err)
	}

	p := new(Profile)
	if err := yaml.Unmarshal(buf, p); err != nil {
		panic(err)
	}

	// parse args
	flag.Parse()
	var out string

	args := []*bool{
		stat, sns, job, lang, os, db, fw, tool,
	}

	if *all {
		fmt.Println(toJson(p))
		return
	}

	for _, arg := range args {
		// output json
		switch {
		case arg == stat && *stat:
			out += toJson(p.Status)
		case arg == sns && *sns:
			out += toJson(p.SNS)
		case arg == job && *job:
			out += toJson(p.JobHistory)
		case arg == lang && *lang:
			out += toJson(p.Dev.Language)
		case arg == os && *os:
			out += toJson(p.Dev.OS)
		case arg == db && *db:
			out += toJson(p.Dev.DB)
		case arg == fw && *fw:
			out += toJson(p.Dev.Framework)
		case arg == tool && *tool:
			out += toJson(p.Dev.Tool)
		}
	}

	fmt.Println(out)
}
