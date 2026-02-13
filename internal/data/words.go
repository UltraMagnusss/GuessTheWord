package data

type Word struct {
	Hint   string
	Answer string
}

func GetWrods() []Word {
	return []Word{
		{
			Hint:   "Capital of Tajikistan",
			Answer: "Dushanbe",
		},
		{
			Hint:   "Capital of Russia",
			Answer: "Moscow",
		},
		{
			Hint:   "Capital of Canada",
			Answer: "Toronto",
		},
		{
			Hint:   "Capital of USA",
			Answer: "Washington",
		},
		{
			Hint:   "Capital of France",
			Answer: "Paris",
		},
		{
			Hint:   "Capital of Germany",
			Answer: "Berlin",
		},
		{
			Hint:   "Capital of Ierland",
			Answer: "Dublin",
		},
		{
			Hint:   "Capital of China",
			Answer: "Beijing",
		},
		{
			Hint:   "Capital of Japan",
			Answer: "Tokyo",
		},
		{
			Hint:   "Capital of Belorus",
			Answer: "Minsk",
		},
	}
}
