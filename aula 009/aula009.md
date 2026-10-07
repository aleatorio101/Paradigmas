1 · Java
(a) Saída: au

(b) Decisão: Execução. Ocorre o polimorfismo em tempo de execução (dynamic method dispatch), onde o Java identifica o tipo real do objeto (Cachorro) e executa o método sobrescrito correspondente.

2 · C++
(a) Saída: A

(b) Decisão: Compilação. Em C++, por padrão, métodos não são virtuais (virtual). Como o ponteiro p é do tipo A*, a resolução do método f() é feita estaticamente com base no tipo do ponteiro em tempo de compilação.

3 · Java
(a) Saída: A B

(b) Decisão:

Para o atributo (x.nome): Compilação (campos em Java sofrem hiding, ou seja, são escondidos e são resolvidos pelo tipo estático da referência, que é A).

Para o método (x.getNome()): Execução (métodos sofrem sobrescrita e usam despacho dinâmico com base no objeto real B).

4 · Python
(a) Saída: 1 2 2

(b) Decisão: Execução. Python é uma linguagem dinamicamente tipada e interpretada, portanto todas as atribuições, chamadas de construtores e incrementos de variáveis de classe ocorrem em tempo de execução.

5 · Java
(a) Saída: A

(b) Decisão: Compilação. Métodos estáticos (static) não participam do polimorfismo de sobrescrita; eles pertencem à classe. A escolha de qual método chamar é feita estaticamente com base no tipo declarado da variável (A). 

6 · Go
(a) Saída: faz ... au

(b) Decisão: Compilação. O mecanismo de embedding em Go realiza composição estática. O método Falar() pertence a Animal e espera um receptor do tipo Animal, logo o a.Som() interno chama o método de Animal (...), enquanto Cao{}.Som() chama diretamente o método sobrescrito de Cao (au).