package main

import (
	"bmstu/cc2024/lab2/grammar"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	input := []string{}

	input = append(input, "2")
	input = append(input, "A B")
	input = append(input, "4")
	input = append(input, "a b c d")
	input = append(input, "5")
	input = append(input, "A -> B a")
	input = append(input, "A -> c")
	input = append(input, "B -> A b")
	input = append(input, "B -> d")
	input = append(input, "B -> A")
	input = append(input, "A")

	g := grammar.ReadGrammar(input)
	// t.Fail()
	if g.G == nil {
		t.Fail()
	}

	fmt.Println(g.ToString())

	chainlessGrammar := g.Eliminate_Chain_Rules()
	fmt.Println("Грамматика после удаления цепных правил")
	fmt.Println(chainlessGrammar.ToString())

	recursionlessGrammar := grammar.Eliminate_left_recursion(chainlessGrammar)

	fmt.Println("Грамматика после удаления рекурсии")
	fmt.Println(recursionlessGrammar.ToString())

}

func Test2(t *testing.T) {

	input := []string{}

	input = append(input, "3")
	input = append(input, "E T F")
	input = append(input, "5")
	input = append(input, "* + ( ) a")
	input = append(input, "6")
	input = append(input, "E -> E + T")
	input = append(input, "E -> T")
	input = append(input, "T -> T * F")
	input = append(input, "T -> F")
	input = append(input, "F -> a")
	input = append(input, "F -> ( E )")
	input = append(input, "E")

	g := grammar.ReadGrammar(input)
	// t.Fail()
	if g.G == nil {
		t.Fail()
	}

	fmt.Println(g.ToString())

	chainlessGrammar := g.Eliminate_Chain_Rules()
	fmt.Println("Грамматика после удаления цепных правил")
	fmt.Println(chainlessGrammar.ToString())

	recursionlessGrammar := grammar.Eliminate_left_recursion(chainlessGrammar)

	fmt.Println("Грамматика после удаления рекурсии")
	fmt.Println(recursionlessGrammar.ToString())

}

func Test3(t *testing.T) {

	input := []string{}

	input = append(input, "3")
	input = append(input, "A B C")
	input = append(input, "2")
	input = append(input, "a b")
	input = append(input, "7")
	input = append(input, "A -> B C")
	input = append(input, "A -> a")
	input = append(input, "B -> C A")
	input = append(input, "B -> A b")
	input = append(input, "C -> A B")
	input = append(input, "C -> C C")
	input = append(input, "C -> a")
	input = append(input, "A")

	g := grammar.ReadGrammar(input)
	// t.Fail()
	if g.G == nil {
		t.Fail()
	}

	fmt.Println(g.ToString())

	chainlessGrammar := g.Eliminate_Chain_Rules()
	fmt.Println("Грамматика после удаления цепных правил")
	fmt.Println(chainlessGrammar.ToString())

	recursionlessGrammar := grammar.Eliminate_left_recursion(chainlessGrammar)

	fmt.Println("Грамматика после удаления рекурсии")
	fmt.Println(recursionlessGrammar.ToString())
}

func Test4(t *testing.T) {
	input := []string{}

	input = append(input, "4")
	input = append(input, "A B C D")
	input = append(input, "3")
	input = append(input, "a b c")
	input = append(input, "6")
	input = append(input, "A -> B")
	input = append(input, "A -> a")
	input = append(input, "B -> C")
	input = append(input, "B -> b")
	input = append(input, "C -> D D")
	input = append(input, "C -> c")
	input = append(input, "A")

	g := grammar.ReadGrammar(input)

	// t.Fail()
	if g.G == nil {
		t.Fail()
	}

	fmt.Println(g.ToString())

	chainlessGrammar := g.Eliminate_Chain_Rules()
	fmt.Println("Грамматика после удаления цепных правил")
	fmt.Println(chainlessGrammar.ToString())

}
