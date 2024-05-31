package grammar

import "fmt"

func (this grammar) ToString() (ret string) {
	ret = fmt.Sprintf("Грамматика:\n")
	for key, prods := range this {
		ret = fmt.Sprintf("%s%s --> ", ret, key)
		// fmt.Println(prods.ToSlice())

		for i, prod := range prods.ToSlice() {
			if i == 0 {
				ret = fmt.Sprintf("%s%s", ret, prod.(production).ToString())
				continue
			}
			ret = fmt.Sprintf("%s | %s", ret, prod.(production).ToString())
		}
		ret = fmt.Sprintf("%s\n", ret)
	}
	return ret
}


func (this FullGrammar) ToString() (ret string){
  ret = fmt.Sprintf("Терминалы: %v\n", this.Terms)
  ret = fmt.Sprintf("%sНетерминалы: %v\n", ret, this.Nonterms)
  ret = fmt.Sprintf("%s%s", ret, this.G.ToString())
  return ret
}
