package config

// Configs is config
type Configs struct {
	Env   Env
	Texts Texts
}

// Env os env info
type Env struct {
	TimeLimit     int
	WordsFilePath string
}

// Texts is input texts
type Texts struct {
	Words []string
}

// Load load config and env
func Load() (Configs, error) {

	env := loadEnv()
	texts, err := loadWords(env)
	if err != nil {
		return *new(Configs), err
	}
	c := &Configs{
		Env:   *env,
		Texts: *texts,
	}
	// fmt.Printf("config loaded: %v\n", c)
	return *c, nil
}
