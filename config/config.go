package config

func Load() {
	app()
	hashing()
	redis()
	database()
}
