package practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person interface {
	ID() string
}

func Equals(s1, s2 Student) bool {
	return s1 == s2
}

type Student struct {
	IDCardNumber string
	Name         string
	Age          int
}

func TestGeneric(t *testing.T) {
	/*
		int1 := int8(1)
		int2 := int8(2)
		int3 := int16(10)
		int4 := int16(20)
		int5 := int32(4)
		int6 := int32(6)
		int7 := 10
		int8 := 20
	*/

	t.Run("int8, int16만 덧셈 가능하게 만들기", func(t *testing.T) {
		/*
			assert.EqualValues(t, Add1(int1, int2), 3)
			assert.EqualValues(t, Add1(int3, int4), 30)
		*/
	})

	t.Run("constraint 사용하여 integer 덧셈 가능하게 만들기", func(t *testing.T) {
		// import "golang.org/x/exp/constraints"
		/*
			assert.EqualValues(t, Add2(int1, int2), 3)
			assert.EqualValues(t, Add2(int3, int4), 30)
			assert.EqualValues(t, Add2(int5, int6), 10)
			assert.EqualValues(t, Add2(int7, int8), 30)
		*/
	})

	t.Run("parameter에 각각 다른 타입 덧셈 가능하게 만들기", func(t *testing.T) {
		/*
			assert.Equal(t, Add3(int1, int8), 21)
			assert.Equal(t, Add3(int2, int3), 12)
			assert.Equal(t, Add3(int4, int7), 30)
		*/
	})

	t.Run("ID가 같다면 같은 학생으로 취급한다.", func(t *testing.T) {
		s1 := Student{IDCardNumber: "1", Name: "tomas", Age: 24}
		s2 := Student{IDCardNumber: "1", Name: "tom", Age: 24}
		assert.True(t, Equals(s1, s2))
	})
}
