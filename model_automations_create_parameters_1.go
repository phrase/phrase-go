package phrase

// AutomationsCreateParameters1 struct for AutomationsCreateParameters1
type AutomationsCreateParameters1 struct {
	// name of the automation
	Name    string `json:"name"`
	Trigger string `json:"trigger"`
	// List of project IDs to associate with the automation. Providing more than one project ID requires the `advanced_job_automation` plan feature; accounts without this feature receive a 422 response with error field `project_ids`.
	ProjectIds []string `json:"project_ids"`
	// id of job template that the automation uses to create jobs from
	JobTemplateId string `json:"job_template_id,omitempty"`
	// translation key statuses used to filter keys that are added to jobs
	StatusFilters []string `json:"status_filters"`
	// used to filter which keys are added to jobs
	Tags []string `json:"tags,omitempty"`
	// along with time_zone, specifies when the scheduled automation is supposed to run
	CronSchedule string `json:"cron_schedule,omitempty"`
	// along with cron_schedule, specifies when the scheduled automation is supposed to run
	TimeZone string `json:"time_zone,omitempty"`
	// User ID of the job owner that newly created jobs are assigned to.
	JobOwnerId string `json:"job_owner_id,omitempty"`
	// When `true`, the automation only acts on locales that changed since its last run.
	IncludeOnlyUpdatedLocales *bool `json:"include_only_updated_locales,omitempty"`
}
