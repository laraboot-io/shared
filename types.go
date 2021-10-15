package main

type LarabootStruct struct {
	Version   string `json:"version"`
	ProjectID string `json:"project_id"`
	Php       struct {
		Version string `json:"version"`
	} `json:"php"`
	Framework struct {
		Config struct {
			Overrides []struct {
				Key     string `json:"key"`
				Envs    string `json:"envs"`
				Default string `json:"default"`
			} `json:"overrides"`
		} `json:"config"`
		Auth struct {
			Stack string `json:"stack"`
		} `json:"auth"`
		Models []struct {
			Name    string `json:"name"`
			Columns []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"columns"`
		} `json:"models"`
	} `json:"Framework"`
}
