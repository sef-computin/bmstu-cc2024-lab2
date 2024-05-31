package main

import (
	"bmstu/cc2024/lab2/grammar"
	"fmt"
)

func main() {

	// input = append(input, "3")
	// input = append(input, "A B C")
	// input = append(input, "2")
	// input = append(input, "a b")
	// input = append(input, "3")
	// input = append(input, "A -> A a")
	// input = append(input, "A -> B b")
	// input = append(input, "B -> a C")

	g := grammar.ReadGrammarFromStdin()

	fmt.Println("Начальная грамматика")
	fmt.Print(g.ToString())

	new_g := grammar.Eliminate_left_recursion(g)

	fmt.Println("Грамматика после удаления левой рекурсии")
	fmt.Print(new_g.ToString())
}
