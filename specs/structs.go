package specs

// Computer structs similar to java pojos
type Computer struct {
	Mfg string
	Model string
	Price int
	Year int
	Configurations // its anonmyous struct, its sub fields gets promoted. other syntx Configuration Configurations.
}

// Configurations structs similar to java pojos
type Configurations struct {
	Hdd string
	Memory string
	Processor string
}