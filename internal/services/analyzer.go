package services

import (
	"fmt"
	"os"
	"path/filepath"
	"template-analyzer/internal/models"
)

type LLMClient interface {
	Send(content string) (string, error)
}

type Analyzer struct {
	llmClient LLMClient
}

func NewAnalyzer(llmClient LLMClient) *Analyzer {
	return &Analyzer{
		llmClient: llmClient,
	}
}

func (analyzer Analyzer) ReadTemplates() ([]models.Template, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	templatesPath := filepath.Join(home, "static-webs", "local", "templates")

	providers, err := os.ReadDir(templatesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read templates directory: %w", err)
	}

	var templates []models.Template

	for _, providerEntry := range providers {
		if !providerEntry.IsDir() {
			continue
		}

		providerName := providerEntry.Name()
		providerPath := filepath.Join(templatesPath, providerName)

		templateDirs, err := os.ReadDir(providerPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read provider %s: %w", providerName, err)
		}

		for _, tplEntry := range templateDirs {
			if !tplEntry.IsDir() {
				continue
			}

			templateName := tplEntry.Name()
			templatePath := filepath.Join(providerPath, templateName)

			size, fileCount, err := analyzer.computeDirectoryStats(templatePath)
			if err != nil {
				return nil, fmt.Errorf("failed to compute stats for %s: %w", templateName, err)
			}

			t := models.Template{
				Name:     templateName,
				Provider: providerName,
				Path:     templatePath,
			}

			t.Features.SizeBytes = size
			t.Features.FileCount = fileCount

			templates = append(templates, t)
		}
	}

	return templates, nil
}

func (analyzer Analyzer) computeDirectoryStats(root string) (int64, int, error) {
	var totalSize int64
	var fileCount int

	if err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk dir %s: %w", path, err)
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %w", d.Name(), err)
		}

		totalSize += info.Size()
		fileCount++

		return nil
	}); err != nil {
		return 0, 0, fmt.Errorf("failed to compute directory size: %w", err)
	}

	return totalSize, fileCount, nil
}

func (analyzer Analyzer) ExtractFeatures(template models.Template) models.Features {
	return models.Features{
		SizeBytes:       0,
		FileCount:       0,
		HTMLCount:       0,
		ImageCount:      0,
		ImageBytes:      0,
		UsesBootstrap:   false,
		UsesTailwind:    false,
		UsesJQuery:      false,
		UsesFrameworkJS: false,
		UsesSCSS:        false,
		HasContactForm:  false,
		HasBackend:      false,
		IsResponsive:    false,
		UsesModernCSS:   false,
		HasAnimations:   false,
		HasIcons:        false,
		HasReadme:       false,
		Style:           "",
		UseCases:        nil,
		TargetClients:   nil,
		Complexity:      "",
		Confidence:      0,
	}
}
