package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Task struct {
	bun.BaseModel `bun:"table:task,alias:t"`
	ID            int64         `bun:"id,pk,notnull,unique" json:"id,omitempty"`
	Name          string        `bun:"name,notnull" json:"name,omitempty"`
	Description   string        `bun:"description,notnull" json:"description,omitempty"`
	TaskCreator   User          `bun:"task_creator,notnull,rel:belongs-to" json:"task_creator,omitempty"`
	TaskAssignee  User          `bun:"task_assignee,notnull,rel:belongs-to" json:"task_assignee,omitempty"`
	CreateTime    time.Time     `bun:"create_time,notnull" json:"create_time,omitempty"`
	Deadline      time.Time     `bun:"deadline,notnull" json:"deadline,omitempty"`
	Priority      Priority      `bun:"priority,notnull" json:"priority,omitempty"`
	Tags          []string      `bun:"tags,array" json:"tags,omitempty"`
	Notes         string        `bun:"notes" json:"notes,omitempty"`
	Attachments   []*Attachment `bun:"attachments,array, " json:"attachments,omitempty"`
	SubTasks      []*Task       `bun:"tasks,rel:has-many" json:"tasks,omitempty"`
	Completed     bool          `bun:"completed,notnull" json:"completed,omitempty"`
	CompletedTime time.Time     `bun:"completed_time,notnull" json:"completed_time,omitempty"`
	Archived      bool          `bun:"archived" json:"archived,omitempty"`
}

type Priority int

const (
	Low    Priority = 0
	Medium Priority = 1
	High   Priority = 2
)

type Attachment struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}
