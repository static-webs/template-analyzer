package main

import (
	"fmt"
	"log"
	"template-analyzer/internal/clients"
	"template-analyzer/internal/models"
	"template-analyzer/internal/services"
)

func main() {
	llmClient := clients.NewLLMClient()
	analyzer := services.NewAnalyzer(llmClient)
	
	templates, err := analyzer.ReadTemplates()
	mustNot(err)

	fmt.Println(fmt.Sprintf("Obtained {%d} templates", len(templates)))
	analyzer.ExtractFeatures(models.Template{})
}

func mustNot(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
