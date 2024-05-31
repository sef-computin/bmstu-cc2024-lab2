package grammar

type FullGrammar struct{
  G grammar
  Terms []string
  Nonterms []string
  Init  string
}


type grammar map[string]productions

func NewGrammar() grammar {

	gr := make(grammar)
	return gr
}

func (this *grammar) AddRule(from string, prod production) {
	if _, ok := (*this)[from]; !ok {
		(*this)[from] = NewProds()
	}
  // fmt.Println("Adding prod: \"", prod.ToSlice(), "\"")
	(*this)[from].Add(prod)
}

func (this *grammar) RemoveRule(from string, prod production){
  if _, ok := (*this)[from]; ok{
    (*this)[from].Remove(prod)
  }
}

func (this grammar) GetNonTerms() []string {
	ret := []string{}
	for key := range this {
		ret = append(ret, key)
	}
	return ret
}

// func (this grammar) GetTerms() []string{
// }


