package theme

import "context"

type Mode struct {
	Name                string `json:"name"`
	PrimaryBackground   string `json:"primary_background"`
	SecondaryBackground string `json:"secondary_background"`
	PrimaryText         string `json:"primary_text"`
	SecondaryText       string `json:"secondary_text"`
	PrimaryBorder       string `json:"primary_border"`
	SecondaryBorder     string `json:"secondary_border"`
	Accent              string `json:"accent"`
	IconFilter          string `json:"icon_filter"`
}

type Theme interface {
	LoadThemes() error
	T(ctx context.Context, key string) Mode
}
