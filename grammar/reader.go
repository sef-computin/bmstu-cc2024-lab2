package grammar

import (
	"strconv"
	"strings"
  "bufio"
  "os" 
  "fmt"
)

func ReadGrammarFromStdin() (fg FullGrammar){
  var input []string

	var line string = ""

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите грамматику: (q прекращает выполнение) ")

	for {
		line, _ = reader.ReadString('\n')
    line = strings.Trim(line, "\n")
		if strings.Compare(line, "q") == 0 {
			break
		}
		// fmt.Println(line, err)
		input = append(input, line)
	}
  g := ReadGrammar(input)

  return g
}

func ReadGrammar(text []string) (fg FullGrammar) {
	
  g := NewGrammar()

	nontermNum, err := strconv.Atoi(text[0])
	if err != nil {
		panic("argument mismatch")
	}
  // fmt.Println(termNum)

	fg.Nonterms = strings.Split(text[1], " ")

  // fmt.Println(termSlice)
	if len(fg.Nonterms) != nontermNum {
		panic("wrong number of arguments")
	}

	termNum, err := strconv.Atoi(text[2])
	if err != nil {
		panic("argument mismatch")
	}
  // fmt.Println(nontermNum)
	fg.Terms = strings.Split(text[3], " ")
  // fmt.Println(nontermSlice)

	if len(fg.Terms) != termNum {
		panic("wrong number of arguments")
	}

	prodsNum, err := strconv.Atoi(text[4])
	if err != nil {
		panic("argument mismatch")
	}

  // fmt.Println(prodsNum)

	for i := 0; i < prodsNum; i++ {
		var from string
    var prodstring string
    var args []string
    // fmt.Println(text[5+i])
    args = strings.Split(text[5+i], " -> ")
    from, prodstring = args[0], args[1]
		g.AddRule(from, NewProd(strings.Split(prodstring, " ")...))
	}

  fg.Init = text[5+prodsNum]

  fg.G = g
	return
}
