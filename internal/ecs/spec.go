// Package ecs source: https://www.elastic.co/docs/reference/ecs/ecs-field-reference
package ecs

import (
	"crypto/sha1"
	"encoding/hex"
	"path/filepath"
	"strings"
	"time"
)

const (
	Version = "9.0.0"
)

type Log struct {
	Base

	Ecs      *Ecs      `json:"ecs"`
	Agent    *Agent    `json:"agent"`
	Event    *Evt      `json:"event"`
	File     *File     `json:"file"`
	Url      *Url      `json:"url,omitempty"`
	Host     *Host     `json:"host,omitempty"`
	User     *User     `json:"user,omitempty"`
	Process  *Process  `json:"process,omitempty"`
	Registry *Registry `json:"registry,omitempty"`
}

type Base struct {
	Timestamp time.Time              `json:"@timestamp"`
	Message   string                 `json:"message"`
	Tags      string                 `json:"tags,omitempty"`
	Labels    map[string]interface{} `json:"labels,omitempty"`
}

type Ecs struct {
	Version string `json:"version"`
}

type Agent struct {
	Type string `json:"type"`
}

type Evt struct {
	Kind     string    `json:"kind,omitempty"`
	Module   string    `json:"module,omitempty"`
	Dataset  string    `json:"dataset,omitempty"`
	Severity int64     `json:"severity,omitempty"`
	ID       string    `json:"id,omitempty"`
	Code     string    `json:"code,omitempty"`
	Provider string    `json:"provider,omitempty"`
	Ingested time.Time `json:"ingested,omitempty"`
	Original string    `json:"original,omitempty"`
	Hash     string    `json:"hash,omitempty"`
}

type File struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Extension   string `json:"extension,omitempty"`
	Directory   string `json:"directory,omitempty"`
	DriveLetter string `json:"drive_letter,omitempty"`
	Path        string `json:"path,omitempty"`
}

type Url struct {
	Original string `json:"original,omitempty"`
	Full     string `json:"full,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Port     int64  `json:"port,omitempty"`
	Path     string `json:"path,omitempty"`
	Query    string `json:"query,omitempty"`
	Fragment string `json:"fragment,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Host struct {
	Hostname string `json:"hostname,omitempty"`
	MAC      string `json:"mac,omitempty"`
}

type User struct {
	ID string `json:"id,omitempty"`
}

type Process struct {
	PID              int64    `json:"pid,omitempty"`
	Thread           *Thread  `json:"thread,omitempty"`
	EntityID         string   `json:"entity_id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Title            string   `json:"title,omitempty"`
	Args             []string `json:"args,omitempty"`
	ArgsCount        int64    `json:"args_count,omitempty"`
	Executable       string   `json:"executable,omitempty"`
	CommandLine      string   `json:"command_line,omitempty"`
	WorkingDirectory string   `json:"working_directory,omitempty"`
}

type Thread struct {
	ID int64 `json:"id,omitempty"`
}

type Registry struct {
	Path  string `json:"path,omitempty"`
	Hive  string `json:"hive,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func NewLog(s, src string, base *Base) *Log {
	return &Log{
		Base: *base,
		Ecs: &Ecs{
			Version: Version,
		},
		Agent: &Agent{
			Type: "flog",
		},
		Event: &Evt{
			Ingested: time.Now().UTC(),
			Original: s,
			Hash:     Hash(s),
		},
		File: file(src),
	}
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

func file(f string) *File {
	dir, _ := filepath.Abs(filepath.Dir(f))
	abs, _ := filepath.Abs(f)

	return &File{
		Type:        "file",
		Name:        filepath.Base(f),
		Extension:   strings.Replace(filepath.Ext(f), ".", "", 1),
		DriveLetter: strings.Replace(filepath.VolumeName(f), ":", "", 1),
		Directory:   dir,
		Path:        abs,
	}
}
