package phrase

// AutomationsCreateParameters struct for AutomationsCreateParameters
type AutomationsCreateParameters struct {
	// name of the automation
	Name    string `json:"name"`
	Trigger string `json:"trigger"`
	// List of project IDs to associate with the automation. Currently, only the first ID in the array is used. The array format leaves room for future support of multiple projects.
	ProjectIds []string `json:"project_ids"`
	// id of job template that the automation uses to create jobs from
	JobTemplateId string `json:"job_template_id,omitempty"`
	// Translation states used when selecting keys for a job.  States are derived from associated translations, not the keys themselves.  When review workflow is enabled, `ready_for_review` is internally treated as `translated`.
	StatusFilters []string `json:"status_filters"`
	// used to filter which keys are added to jobs
	Tags []string `json:"tags,omitempty"`
	// along with time_zone, specifies when the scheduled automation is supposed to run
	CronSchedule string `json:"cron_schedule,omitempty"`
	// along with cron_schedule, specifies when the scheduled automation is supposed to run
	TimeZone string `json:"time_zone,omitempty"`
}
