package utils

type SomeInt int

const (
	ZeroInt SomeInt = iota
	OneInt
	SecondInt
)

// Пример структуры
type SomeStruct struct {
	A SomeInt
}

func (s *SomeStruct) Func() {
	s.A += 1
}

// st := &utils.SomeStruct{}
// 	var itf interface{} = st
// 	st.Func()
// 	fmt.Println(st.A)
// 	switch itf.(type) {
// 	case *utils.SomeInt:
// 	}

// st := utils.SomeInt(1)
// 	switch st {
// 	case utils.OneInt:
// 		fallthrough
// 	default:
// 	}
