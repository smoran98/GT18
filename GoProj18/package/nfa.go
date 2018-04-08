package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

// Helper state
type nfa struct {
	initial *state
	accept  *state
}

// Returns pointer
func ReturnNFA(pofix string) *nfa {

	nfaStack := []*nfa{}
	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}

			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '+':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &initial

			nfaStack = append(nfaStack, &nfa{initial: frag.initial, accept: &accept})

		case '?':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag.initial, edge2: frag.accept}

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: frag.accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}
	// Handle errors
	if len(nfaStack) != 1 {
		fmt.Println("Sorry more than 1 nfa found", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}

// MATCH
func Pomatch(po string, s string) bool {

	ismatch := false

	ponfa := ReturnNFA(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}

		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}
	return ismatch
}

func addState(l []*state, s *state, a *state) []*state {

	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

// infix to post
func IntoPost(infix string) string {

	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '?': 7, '+': 6}
	postfix := []rune{}
	stack := []rune{}

	// Range converts each character of string into rune
	for _, r := range infix {
		switch {
		case r == '(':
			//if there is a ( then this gets appended to the stack
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// Discard open bracket
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			// If the length of stack is greater than zero
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default:
			postfix = append(postfix, r)
		}
	}
	// Handle errors
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	// Return the string
	return string(postfix)
}

func ReadUserInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}
