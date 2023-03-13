package constants

const (
	AnimalGenderMale   = "MALE"
	AnimalGenderFemale = "FEMALE"
	AnimalGenderOther  = "OTHER"
)

var animalGenders = []string{
	AnimalGenderMale,
	AnimalGenderFemale,
	AnimalGenderOther,
}

func IsAnimalGender(value string) bool {
	for _, gender := range animalGenders {
		if gender == value {
			return true
		}
	}
	return false
}
