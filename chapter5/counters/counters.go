package counters

// our package-private type
type alertCounter int

// explicit constructor for the private type
func New(value int) alertCounter {
	return alertCounter(value)
}