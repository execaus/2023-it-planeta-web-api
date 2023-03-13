package constants

const (
	AnimalLifeStatusDead  = "DEAD"
	AnimalLifeStatusAlive = "ALIVE"
)

var animalLifeStatuses = []string{
	AnimalLifeStatusAlive,
	AnimalLifeStatusDead,
}

func IsAnimalLifeStatus(value string) bool {
	for _, status := range animalLifeStatuses {
		if status == value {
			return true
		}
	}
	return false
}
