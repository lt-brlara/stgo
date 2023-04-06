package main

type stigData struct {
	Stig struct {
		Findings map[string]stigFinding `json:"findings"`
	} `json:"stig"`
}

type stigFinding struct {
	CheckId         string      `json:"checkid"`
	CheckText       string      `json:"checktext"`
	DescriptionInfo string      `json:"description"`
	FixId           string      `json:"fixid"`
	FixText         string      `json:"fixtext"`
	IaControls      interface{} `json:"iacontrols"`
	Id              string      `json:"id"`
	RuleID          string      `json:"ruleID"`
	Severity        string      `json:"severity"`
	TitleInfo       string      `json:"title"`
	Version         string      `json:"version"`
}

func (i stigFinding) Title() string       { return i.Id }
func (i stigFinding) Description() string { return i.TitleInfo }
func (i stigFinding) FilterValue() string { return i.Id }
