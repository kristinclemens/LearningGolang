package interfaces

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestTabbyCat_Meow(t *testing.T) {
	tabbyCat := &TabbyCat{}
	assert.Equal(t, "*purr*", tabbyCat.Purr())
}

func TestTabbyCat_Purr(t *testing.T) {
	tabbyCat := &TabbyCat{}
	assert.Equal(t, "Meow!", tabbyCat.Meow())
}

// Because a TabbyCat can Purr, it must implement the Cat interface
func TestTabbyCat_ImplementsCat(t *testing.T) {
	typeOfTabbyCat := reflect.TypeOf((*TabbyCat)(nil))
	typeOfCat := reflect.TypeOf((*Cat)(nil)).Elem()

	assert.True(t, typeOfTabbyCat.Implements(typeOfCat))
}

// Because a TabbyCat can Meow, it must implement the HouseCat interface
func TestTabbyCat_ImplementsHouseCat(t *testing.T) {
	typeOfTabbyCat := reflect.TypeOf((*TabbyCat)(nil))
	typeOfHouseCat := reflect.TypeOf((*HouseCat)(nil)).Elem()
	assert.True(t, typeOfTabbyCat.Implements(typeOfHouseCat))
}

func TestLion_Purr(t *testing.T) {
	lion := &Lion{}
	assert.Equal(t, "*rumble*", lion.Purr())
}

// Because a Lion can Purr, but not Meow, it must implement the Cat interface, but not the HouseCat interface
func TestLion_ImplementsCat_NotHouseCat(t *testing.T) {
	typeOfLion := reflect.TypeOf((*Lion)(nil))
	typeOfCat := reflect.TypeOf((*Cat)(nil)).Elem()
	assert.True(t, typeOfLion.Implements(typeOfCat))

	typeOfHouseCat := reflect.TypeOf((*HouseCat)(nil)).Elem()
	assert.False(t, typeOfLion.Implements(typeOfHouseCat))
}

// Both TabbyCat and Lion implement the Animal interface because it has no methods
func TestTabbyCat_ImplementsAnimal(t *testing.T) {
	typeOfTabbyCat := reflect.TypeOf((*TabbyCat)(nil))
	typeOfAnimal := reflect.TypeOf((*Animal)(nil)).Elem()
	assert.True(t, typeOfTabbyCat.Implements(typeOfAnimal))
}

func TestLion_ImplementsAnimal(t *testing.T) {
	typeOfLion := reflect.TypeOf((*Lion)(nil))
	typeOfAnimal := reflect.TypeOf((*Animal)(nil)).Elem()
	assert.True(t, typeOfLion.Implements(typeOfAnimal))
}
