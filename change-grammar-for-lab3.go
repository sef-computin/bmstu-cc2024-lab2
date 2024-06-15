package main

import (
	"bmstu/cc2024/lab2/grammar"
	"fmt"
)


func main(){
  input := []string{}

	input = append(input, "11")
	input = append(input, "<программа> <блок> <список_операторов> <оператор> <выражение> <логическое_выражение> <логический_одночлен> <вторичное_лог_выражение> <первичное_лог_выражение> <логическое_значение> <знак_операции>")
	input = append(input, "9")
	input = append(input, "begin end ; = ~ ! & true false")
	input = append(input, "20")
	input = append(input, "<программа> -> <блок>")
	input = append(input, "<блок> -> begin <список_операторов> end")
	input = append(input, "<список_операторов> -> <оператор>")
	input = append(input, "<список_операторов> -> <список_операторов> ; <оператор>")
	input = append(input, "<оператор> -> <идентификатор> = <выражение>")
	input = append(input, "<оператор> -> <блок>")
  input = append(input, "<выражение> -> <логическое_выражение>")
	input = append(input, "<логическое_выражение> -> <логический_одночлен>")
	input = append(input, "<логическое_выражение> -> <логическое_выражение> ! <логический_одночлен>")
	input = append(input, "<логический_одночлен> -> <вторичное_лог_выражение>") 
  input = append(input, "<логический_одночлен> -> <логический_одночлен> & <вторичное_лог_выражение>")
	input = append(input, "<вторичное_лог_выражение> -> <первичное_лог_выражение>")
  input = append(input, "<логический_одночлен> -> ~ <первичное_лог_выражение>")
	input = append(input, "<первичное_лог_выражение> -> <логическое_значение>")
	input = append(input, "<первичное_лог_выражение> -> <идентификатор>")
	input = append(input, "<логическое_значение> -> true")
	input = append(input, "<логическое_значение> -> false")

	input = append(input, "<знак_операции> -> ~")
	input = append(input, "<знак_операции> -> &")
	input = append(input, "<знак_операции> -> !")
	
  input = append(input, "<программа>")

  gr := grammar.ReadGrammar(input)

  fmt.Println(gr.ToString())


  new_gr := grammar.Eliminate_left_recursion(gr)

  fmt.Println(new_gr.ToString())
}
