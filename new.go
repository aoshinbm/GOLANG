
func dynamicPort() Settings {
	// Open and read the settings file
	file, err := os.Open("setting.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the JSON settings
	var setting Settings
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&setting)
	if err != nil {
		log.Fatal(err)
	}
	return setting
}

+++++++++++++++++++++++++++++
file, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Unmarshal using the `encoding/json` package
	var person1 Person
	err = json.Unmarshal(data, &person1)
	if err != nil {
		panic(err)
	}