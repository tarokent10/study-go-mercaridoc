package config

// Configs is config
type Configs struct {
	env   Env
	texts Texts
}

// Load load config and env
func Load() (*Configs, error) {

	env := loadEnv()
	texts, err := loadWords(env)
	if err != nil {
		return nil, err
	}
	c := &Configs{
		env:   *env,
		texts: *texts,
	}
	// fmt.Printf("config loaded: %v\n", c)
	return c, nil
}
