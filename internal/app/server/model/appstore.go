package model

type Application struct {
	Name   string `json:"name"`
	Intro  string `json:"intro"`
	Image  string `json:"image"`
	Port   []int  `json:"port"`
	Config string `json:"config"`
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Render() {
	// tpl := template.New()
	// tpl.Execute(os.Stdout, app)
}
