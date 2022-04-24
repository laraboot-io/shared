package shared

// LarabootStruct : Laraboot file schema.
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
	Build struct {
		Tasks []struct {
			Name   string `json:"name"`
			Path   string `json:"path"`
			Local  bool   `json:"local"`
			Format string `json:"format"`
			Env    struct {
			} `json:"env"`
		} `json:"tasks"`
	} `json:"Build"`
}
