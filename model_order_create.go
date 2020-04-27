package phrase
// OrderCreate struct for OrderCreate
type OrderCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Name of the LSP that should process this order. Can be one of gengo, textmaster.
	Lsp string `json:"lsp,omitempty"`
	// Source locale for the order. Can be the name or public id of the source locale. Preferred is the public id.
	SourceLocaleId string `json:"source_locale_id,omitempty"`
	// List of target locales you want the source content translate to. Can be the name or public id of the target locales. Preferred is the public id.
	TargetLocaleIds []string `json:"target_locale_ids,omitempty"`
	// Name of the quality level, availability depends on the LSP. Can be one of:  standard, pro (for orders processed by Gengo) and one of regular, premium, enterprise (for orders processed by TextMaster)
	TranslationType string `json:"translation_type,omitempty"`
	// Tag you want to order translations for.
	Tag string `json:"tag,omitempty"`
	// Message that is displayed to the translators for description.
	Message string `json:"message,omitempty"`
	// Style guide for translators to be sent with the order.
	StyleguideId string `json:"styleguide_id,omitempty"`
	// Unverify translations upon delivery.
	UnverifyTranslationsUponDelivery bool `json:"unverify_translations_upon_delivery,omitempty"`
	// Order translations for keys with untranslated content in the selected target locales.
	IncludeUntranslatedKeys bool `json:"include_untranslated_keys,omitempty"`
	// Order translations for keys with unverified content in the selected target locales.
	IncludeUnverifiedTranslations bool `json:"include_unverified_translations,omitempty"`
	// Category to use (required for orders processed by TextMaster).
	Category string `json:"category,omitempty"`
	// Extra proofreading option to ensure consistency in vocabulary and style. Only available for orders processed by TextMaster.
	Quality bool `json:"quality,omitempty"`
	// Indicates whether the priority option should be ordered which decreases turnaround time by 30%. Available only for orders processed by TextMaster.
	Priority bool `json:"priority,omitempty"`
}
