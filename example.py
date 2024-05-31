
def eliminate_left_recursion(grammar):
    non_terminals = grammar.keys()
    new_grammar = {}

    for A in non_terminals:
        productions = grammar[A]
        A_productions = []
        A_primed_productions = []
        
        for production in productions:
            if production[0] == A:  # проверка на левостороннюю рекурсию
                print("рекурсия в ", production)
                A_primed_productions.append(production[1:] + (A + "'",))
                print("primed: ", A_primed_productions)
            else:
                A_productions.append(production + (A + "'",))
                print("nPrime: ", A_productions)

        if A_primed_productions:
            new_grammar[A] = A_productions
            new_grammar[A + "'"] = A_primed_productions + [("ε",)]  # добавляем пустую продукцию
        else:
            new_grammar[A] = productions

    return new_grammar

def print_grammar(grammar):
    for non_terminal, productions in grammar.items():
        print(non_terminal + " -> ", end="")
        for production in productions:
            print(" ".join(production), end=" | ")
        print("\b\b  ")  # Удаление последних двух символов "| " и переход на новую строку

# Пример грамматики с левосторонней рекурсией
original_grammar = {
    'A': [('A', 'a'), ('A', 'b'), ('c',)],
    'B': [('A', 'd'), ('B', 'e')]
}

print("Исходная грамматика:")
print_grammar(original_grammar)

new_grammar = eliminate_left_recursion(original_grammar)

print("\nГрамматика без левосторонней рекурсии:")
print_grammar(new_grammar)
print(new_grammar)
