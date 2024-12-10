package algorithm

import "push/utils"

//push the top first element of stack b to stack a
func pa(s []*utils.Stack) { 
	if len(s[1].Nums) > 0 {
		s[0].PushBack(s[1].PopBack())
	}
}

//push the top first element of stack a to stack b
func pb(s []*utils.Stack) {
	if len(s[0].Nums) > 0 {
		s[1].PushBack(s[0].PopBack())
	}
}

//swap first 2 elements of stack a
func sa(s []*utils.Stack) {
	if len(s[0].Nums) > 1 {
		f1 := s[0].PopBack()
		f2 := s[0].PopBack()
		s[0].PushBack(f1)
		s[0].PushBack(f2)
	}
}

//swap first 2 elements of stack b
func sb(s []*utils.Stack) {
	if len(s[1].Nums) > 1 {
		f1 := s[1].PopBack()
		f2 := s[1].PopBack()
		s[1].PushBack(f1)
		s[1].PushBack(f2)
	}
}

//execute sa and sb
func ss(s []*utils.Stack) {
	sa(s)
	sb(s)
}


//rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func ra(s []*utils.Stack) {
	if len(s[0].Nums) == 0 {
		return
	}
	s[0].PushFront(s[0].PopBack())
}

//rotate stack b
func rb(s []*utils.Stack) {
	if len(s[1].Nums) == 0 {
		return
	}
	s[1].PushFront(s[1].PopBack())
}

//execute ra and rb
func rr(s []*utils.Stack) {
	ra(s)
	rb(s)
}

//reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func rra(s []*utils.Stack){
	if len(s[0].Nums) == 0 {
		return
	}
	s[0].PushBack(s[0].PopFront())
}

//reverse rotate b
func rrb(s []*utils.Stack) {
	if len(s[1].Nums) == 0 {
		return
	}
	s[1].PushBack(s[1].PopFront())
}

//execute rra and rrb
func rrr(s []*utils.Stack) {
	rra(s)
	rrb(s)
}