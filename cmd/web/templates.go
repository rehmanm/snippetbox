package main

import (
	"github.com/rehmanm/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
