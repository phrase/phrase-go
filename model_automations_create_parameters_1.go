package phrase

// AutomationsCreateParameters1 struct for AutomationsCreateParameters1
type AutomationsCreateParameters1 struct {
	// name of the automation
	Name    string `json:"name"`
	Trigger string `json:"trigger"`
	// List of project IDs to associate with the automation. Currently, only the first ID in the array is used. The array format leaves room for future support of multiple projects.
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
}
