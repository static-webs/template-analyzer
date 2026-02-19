package models

type Features struct {
	// non-semantic
	SizeBytes       int64
	FileCount       int
	HTMLCount       int
	ImageCount      int
	ImageBytes      int64
	UsesBootstrap   bool
	UsesTailwind    bool
	UsesJQuery      bool
	UsesFrameworkJS bool
	UsesSCSS        bool
	HasContactForm  bool
	HasBackend      bool
	IsResponsive    bool
	UsesModernCSS   bool
	HasAnimations   bool
	HasIcons        bool
	HasReadme       bool

	// semantic
	Style         string   // minimalist, colored, informative, etc
	UseCases      []string // landing page, portfolio, blog, etc
	TargetClients []string // independent professional, small business, etc
	Complexity    string   // low, medium, hight
	Confidence    float64  // llm confidence
}
