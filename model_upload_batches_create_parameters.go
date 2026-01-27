package phrase

// UploadBatchesCreateParameters struct for UploadBatchesCreateParameters
type UploadBatchesCreateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// If set to true, after all uploads in the batch are completed, translation keys that were not mentioned in any of the uploaded files will be deleted.
	DeleteUnmentionedKeys *bool `json:"delete_unmentioned_keys,omitempty"`
	// Array of upload IDs to include in the batch
	UploadIds []string `json:"upload_ids"`
}
