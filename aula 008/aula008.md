# Conceitos de Linguagens de Programação

---

### 1. Python (Argumentos Padrão Mutáveis)

```python
def adicionar(item, lista=[]):
    lista.append(item)
    return lista

print(adicionar(1))
print(adicionar(2))
```

* **(a) Saída:**
  ```text
  [1]
  [1, 2]
  ```
* **(b) Conceito da aula:** Argumentos padrão mutáveis (*Mutable Default Arguments*). Em Python, os valores padrão dos argumentos de funções são avaliados uma única vez, no momento em que a função é definida (e não a cada chamada). Portanto, a mesma lista é reutilizada em chamadas subsequentes se nenhum novo argumento de lista for fornecido.

---

### 2. Java (Passagem de Parâmetros)

```java
static void zera(int[] v, int n) {
    v[0] = 0;
    n = 0;
}

int[] v = {5, 5};
int n = 5;
zera(v, n);
System.out.println(v[0] + " " + n);
```

* **(a) Saída:**
  ```text
  0 5
  ```
* **(b) Conceito da aula:** Passagem por valor e referência de objetos (*Pass-by-value / References*). Em Java, variáveis primitivas (`int n`) são passadas por valor (uma cópia é enviada para o método, logo o original não muda). Já os arrays (`int[] v`) passam a cópia da referência ao objeto na memória; portanto, modificar `v[0]` afeta diretamente o array original.

---

### 3. Python (Late Binding em Closures)

```python
fs = [lambda: i for i in range(3)]
print([f() for f in fs])
```

* **(a) Saída:**
  ```text
  [2, 2, 2]
  ```
* **(b) Conceito da aula:** Ligação tardia (*Late Binding*) em *Closures*. As funções lambdas criadas no *comprehension* avaliam a variável $i$ apenas no momento em que são executadas (`f()`), e não no momento da criação. Quando as funções são chamadas, o loop já terminou e a variável $i$ reteve seu último valor ($2$).

---

### 4. C (Variáveis Estáticas Locais)

```c
int contador(void) {
    static int n = 0;
    return ++n;
}

// em main:
contador(); 
contador();
printf("%d\n", contador());
```

* **(a) Saída:**
  ```text
  3
  ```
* **(b) Conceito da aula:** Variáveis Locais Estáticas (`static`) em C. Uma variável local declarada como `static` mantém o seu valor entre as chamadas da função e é inicializada apenas uma vez. As duas primeiras chamadas incrementam $n$ para $1$ e $2$, e a terceira chamada incrementa para $3$ antes de imprimir.

---

### 5. Rust (Ownership e Borrowing)

```rust
fn dobra(v: Vec<i32>) -> Vec<i32> {
    v.iter().map(|x| x * 2).collect()
}

let v = vec![1, 2, 3];
let d = dobra(v);
println!("{:?} {:?}", v, d);
```

* **(a) Saída:**
  ```text
  Erro de Compilação (Compile Error: use of moved value: v)
  ```
* **(b) Conceito da aula:** Sistema de Propriedade (*Ownership*). Em Rust, quando a variável `v` (um `Vec`, que não implementa o trait `Copy`) é passada como argumento para a função `dobra(v)`, a propriedade (*ownership*) de `v` é transferida (movida) para dentro da função. Tentar usar `v` no `println!` após a chamada gera um erro de compilação por uso de valor movido.

---

### 6. Python (Escopo de Variáveis e a Palavra-chave global)

```python
total = 0

def adiciona(x):
    total = total + x
    return total

print(adiciona(5))
```

* **(a) Saída:**
  ```text
  Erro em Tempo de Execução: UnboundLocalError: local variable 'total' referenced before assignment
  ```
* **(b) Conceito da aula:** Escopo de Variáveis (*Global vs Local*). Quando o interpretador Python encontra uma atribuição a uma variável dentro de uma função (`total = ...`), ele classifica automaticamente essa variável como local para todo o escopo da função. Como `total` está sendo lida (`total + x`) antes de ser atribuída localmente, ocorre o erro. Para corrigir, seria necessário utilizar a palavra-chave `global total` ou utilizar o total dentro da definição.