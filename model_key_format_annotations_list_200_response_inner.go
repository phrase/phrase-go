package phrase

// KeyFormatAnnotationsList200ResponseInner struct for KeyFormatAnnotationsList200ResponseInner
type KeyFormatAnnotationsList200ResponseInner struct {
	// The file format that produced this annotation.
	FileFormat string `json:"file_format,omitempty"`
	// The original file-format snippet captured when the key was imported. Returned as-is, may contain multi-line content.
	OriginalRepresentation string `json:"original_representation,omitempty"`
}
