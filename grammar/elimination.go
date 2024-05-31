package grammar

import (
	"fmt"
	"strings"
)

func Eliminate_left_recursion(g FullGrammar) FullGrammar {
	new_grammar := FullGrammar{}

	new_grammar.Terms = g.Terms

	new_grammar.G = eliminate_all_left_recursion(g.G, g.Nonterms)
	new_grammar.Nonterms = new_grammar.G.GetNonTerms()

	return new_grammar

}

func eliminate_direct_left_recursion(g grammar) grammar {
	new_grammar := NewGrammar()

	non_terms := g.GetNonTerms()

	// visited_nonterms := map[string]interface{}{}

	for _, a := range non_terms {
		prods := g[a]

		a_prime := fmt.Sprintf("%s'", a)
		A_prods := []production{}
		A_primed_prods := []production{}

		for _, elem := range prods.ToSlice() {
			prod := elem.(production)
			if prod.ToSlice()[0] == a {
				A_primed_prods = append(A_primed_prods, NewProd(append(prod.ToSlice()[1:], a_prime)...))
			} else {
				A_prods = append(A_prods, NewProd(append(prod.ToSlice(), a_prime)...))
			}
		}

		if len(A_primed_prods) > 0 {
			for _, prod := range A_prods {
				new_grammar.AddRule(a, prod)
			}
			for _, prod := range A_primed_prods {
				new_grammar.AddRule(a_prime, prod)
			}
			new_grammar.AddRule(a_prime, NewProd("ε"))
		} else {
			for _, prod := range prods.ToSlice() {
				new_grammar.AddRule(a, prod.(production))
			}

		}

	}

	return new_grammar
}

func eliminate_all_left_recursion(g grammar, nonTerms []string) grammar {

	newGrammar := NewGrammar()

	// nonTerms := g.GetNonTerms()

	// fmt.Println(nonTerms)

	for i, Ai := range nonTerms {
		for j := 0; j < i; j++ {
			Aj := nonTerms[j]
			newProds := make([]production, 0)
			for _, elem := range g[Ai].ToSlice() {
				prod := elem.(production)
				if prod.ToSlice()[0] == Aj {
					for _, beta := range g[Aj].ToSlice() {
						newProds = append(newProds, NewProd(append(beta.(production).ToSlice(), prod.ToSlice()[1:]...)...))
					}

				} else {
					newProds = append(newProds, prod)
				}
			}
			g[Ai] = NewProds()
			for _, val := range newProds {
				g.AddRule(Ai, val)
			}
		}

		directProductions, recursiveProductions := make([]production, 0), make([]production, 0)

		for _, elem := range g[Ai].ToSlice() {
			prod := elem.(production)
			if prod.ToSlice()[0] == Ai {
				recursiveProductions = append(recursiveProductions, NewProd(prod.ToSlice()[1:]...))
			} else {
				directProductions = append(directProductions, prod)
			}
		}

		if len(recursiveProductions) > 0 {
			AiPrime := Ai + "'"
			for _, prod := range directProductions {
				newGrammar.AddRule(Ai, NewProd(append(prod.ToSlice(), AiPrime)...))
			}

			for _, prod := range recursiveProductions {
				newGrammar.AddRule(AiPrime, NewProd(append(prod.ToSlice(), AiPrime)...))
			}
			newGrammar.AddRule(AiPrime, NewProd("ε"))
		} else {
			newGrammar[Ai] = g[Ai]
		}

	}
	return newGrammar
}

func (fg FullGrammar) Eliminate_Chain_Rules() FullGrammar {
	newGrammar := FullGrammar{}

	newGrammar.Nonterms, newGrammar.Terms, newGrammar.Init = fg.Nonterms, fg.Terms, fg.Init
	newGrammar.G = fg.G.eliminate_all_chain_rules()
	return newGrammar
}

func (g grammar) eliminate_all_chain_rules() grammar {
	newGrammar := NewGrammar()

	for A := range g {
		reachable := map[string]bool{}
		reachable[A] = true
		visitQueue := []string{A}

		for len(visitQueue) > 0 {
			current := visitQueue[0]
			visitQueue = visitQueue[1:]

			for _, elem := range g[current].ToSlice() {
				prod := elem.(production)
				if len(prod) == 1 && isIn(prod.ToSlice()[0], g.GetNonTerms()) {
					if reachable[prod.ToSlice()[0]] == false {
						reachable[prod.ToSlice()[0]] = true
						visitQueue = append(visitQueue, prod.ToSlice()[0])
					}
				}
			}
		}

		newProds := make([]production, 0)
		for B := range reachable {
			for _, elem := range g[B].ToSlice() {
				prod := elem.(production)
				if len(prod) != 1 || isIn(prod.ToSlice()[0], g.GetNonTerms()) == false {
					newProds = append(newProds, prod)
				}
			}
		}

		for _, prod := range newProds {
			newGrammar.AddRule(A, prod)
		}
	}

	return newGrammar

}

func isIn(symbol string, nonterms []string) (ret bool) {
	ret = false

	for _, val := range nonterms {
		if strings.Compare(val, symbol) == 0 {
			ret = true
		}
	}

	return
}
