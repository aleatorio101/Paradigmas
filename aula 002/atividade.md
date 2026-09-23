# Evolução das Principais Linguagens de Programação

## 1. A genealogia das linguagens não é uma escada de progresso

A história das linguagens de programação não acontece de forma linear, como se cada linguagem nova fosse necessariamente melhor que a anterior e tivesse que substituí-la. Na prática, várias linguagens continuam sendo utilizadas ao mesmo tempo, porque possuem objetivos, paradigmas e contextos diferentes.

Dois fatores históricos ajudam a explicar por que uma linguagem pode influenciar outra sem substituí-la:

* **Necessidades e contextos diferentes:** uma linguagem pode ser criada para uma finalidade específica e continuar sendo adequada para ela mesmo depois do surgimento de linguagens mais novas. Por exemplo, JavaScript se tornou muito importante no desenvolvimento para navegadores, enquanto Java possui características que favoreceram seu uso em aplicações multiplataforma.

* **Legado e comunidade:** linguagens que já são utilizadas há muitos anos acumulam sistemas, bibliotecas, ferramentas e profissionais com conhecimento naquela tecnologia. Por isso, trocar completamente de linguagem pode ser caro e trabalhoso. Além disso, linguagens novas costumam aproveitar conceitos das anteriores, fazendo com que exista influência mesmo sem uma substituição direta.

---

## 2. Plankalkül

O **Plankalkül**, projeto desenvolvido por Konrad Zuse, foi um dos primeiros projetos de uma linguagem de programação de alto nível. Mesmo sem ter sido implementado em um computador naquela época, o projeto já mostrava que era possível representar algoritmos complexos de maneira estruturada.

Entre os recursos que já apareciam no projeto estavam:

* **Vetores e registros:** permitiam organizar informações em estruturas de dados mais complexas, inclusive possibilitando o aninhamento de estruturas, como um vetor dentro de um registro.

* **Sub-rotinas:** permitiam dividir um algoritmo em partes menores que poderiam ser reutilizadas. Isso ajudava a organizar programas grandes e evitava a repetição desnecessária de operações.

* **Laços de repetição:** o projeto também apresentava mecanismos para representar repetições, permitindo descrever iterações sem depender apenas de comandos de desvio como o `GOTO`.

Dessa forma, o Plankalkül antecipou vários conceitos que posteriormente se tornariam comuns em outras linguagens.

---

## 4. Fortran e a competição com o código de máquina

Antes do Fortran, existia a ideia de que, para conseguir o melhor desempenho possível, era necessário programar diretamente em código de máquina ou Assembly. Por isso, um dos principais desafios do projeto Fortran era convencer os programadores de que um código escrito em uma linguagem de alto nível e depois traduzido por um compilador poderia ter um desempenho competitivo.

O Fortran procurou resolver esse problema principalmente em aplicações científicas e numéricas. Seus compiladores eram capazes de gerar código relativamente eficiente para a época, principalmente em operações matemáticas.

* **Desempenho:** o código gerado pelo compilador conseguia chegar perto do desempenho de programas escritos manualmente em Assembly em determinadas situações.

* **Custo de programação:** escrever em Fortran era muito mais simples e rápido do que programar diretamente em código de máquina, principalmente para cálculos matemáticos e operações com números de ponto flutuante.

* **Adoção:** mesmo que em alguns casos fosse possível obter um desempenho maior com código escrito manualmente, a economia de tempo e esforço na programação tornava o Fortran uma alternativa muito interessante.

Assim, o Fortran ajudou a mostrar que uma linguagem de alto nível poderia oferecer um bom equilíbrio entre **desempenho e produtividade**, facilitando sua adoção.

---

## 6. Importância histórica do ALGOL 60

O **ALGOL 60** teve uma importância histórica muito grande, mesmo não tendo alcançado uma adoção comercial tão ampla. Sua principal contribuição foi influenciar o desenvolvimento de várias linguagens posteriores e ajudar a consolidar conceitos importantes da programação.

Alguns exemplos são:

* **Recursividade:** permitia que um procedimento chamasse a si mesmo, possibilitando soluções mais naturais para determinados problemas.

* **Blocos e escopo:** ajudaram a organizar os programas e estabelecer regras mais claras sobre onde as variáveis poderiam ser utilizadas.

* **BNF (Backus-Naur Form):** foi utilizada para descrever formalmente a sintaxe da linguagem, contribuindo para uma especificação mais precisa.

Essas ideias influenciaram linguagens como **C, C++, Pascal e Java**. Portanto, mesmo sem ter sido uma linguagem comercialmente dominante, o ALGOL 60 teve uma influência muito grande na evolução das linguagens de programação.

---

## 7. COBOL, processamento comercial e FLOW-MATIC

O COBOL foi criado pensando principalmente no **processamento de dados comerciais**. Por isso, tanto o domínio quanto o público para o qual a linguagem era destinada influenciaram bastante suas características.

### Legibilidade

A linguagem buscava utilizar uma sintaxe relativamente próxima do inglês, com comandos como `IF`, `MOVE`, `ADD` e `PERFORM`. A ideia era facilitar a leitura e manutenção dos programas, inclusive por profissionais que não fossem especialistas em programação.

### Registros

Como o objetivo era trabalhar com informações empresariais, o COBOL dava bastante importância à representação de **registros e arquivos estruturados**. Era possível definir campos e organizar os dados de maneira bastante explícita.

### Relação com a FLOW-MATIC

O COBOL recebeu forte influência da **FLOW-MATIC**, linguagem desenvolvida por Grace Hopper. A FLOW-MATIC já tinha como objetivo facilitar a programação de aplicações comerciais e utilizar uma linguagem mais próxima do inglês.

O COBOL aproveitou várias dessas ideias e as levou para uma linguagem mais abrangente e padronizada, voltada para o processamento de dados empresariais.

### Orientação a dados x orientação a objetos

| Orientado a dados                                                                        | Orientado a objetos                                                            |
| ---------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| O foco está nos dados e no seu processamento.                                            | O foco está nos objetos e na combinação de dados e comportamentos.             |
| Os dados geralmente são manipulados por procedimentos.                                   | Os objetos possuem atributos e métodos.                                        |
| A organização gira principalmente em torno das informações que precisam ser processadas. | A organização gira em torno das entidades e das ações que elas podem realizar. |

---

## 11. ALGOL 60 → Pascal → C

O **ALGOL 60** consolidou diversos conceitos importantes da programação, como estruturas de controle, blocos, escopo de variáveis, tipos e procedimentos. Por isso, acabou influenciando várias linguagens que surgiram depois.

A **Pascal** recebeu forte influência do ALGOL e manteve sua preocupação com a programação estruturada. Além disso, trouxe uma preocupação maior com tipagem, organização dos dados e estruturas de programação. Por isso, acabou sendo bastante utilizada no ensino de programação.

A linguagem **C** também recebeu influência dessa tradição de programação estruturada e imperativa. Porém, C deu mais destaque à eficiência, ao controle da execução e à manipulação de memória. Isso fez com que se tornasse uma linguagem muito importante para o desenvolvimento de sistemas e também influenciasse várias linguagens posteriores.

O **Prolog** segue uma direção diferente. Enquanto ALGOL, Pascal e C trabalham principalmente com a ideia de indicar como o computador deve executar determinada tarefa, o Prolog utiliza **programação lógica**.

Em vez de descrever passo a passo o que deve ser feito, o programador informa fatos, regras e consultas. A partir dessas informações, o sistema pode fazer inferências para encontrar uma resposta.

---

## 12. Exemplo de programação lógica em Prolog

Uma pequena base de conhecimento pode ser representada da seguinte maneira:

### Fatos

* O gato tem pelos.
* O gato caça ratos.

### Regra

* Se um animal tem pelos e caça ratos, então ele é um bom caçador.

Em Prolog:

```prolog
tem_pelos(gato).
caca_ratos(gato).

bom_cacador(X) :- tem_pelos(X), caca_ratos(X).
```

### Consulta

```prolog
?- bom_cacador(gato).
```

Resultado:

```text
true
```

Isso representa programação lógica porque o programa não está apenas armazenando informações. Existe uma **regra** que permite chegar a uma nova conclusão a partir dos fatos existentes.

Nesse caso, a partir de `tem_pelos(gato)` e `caca_ratos(gato)`, o Prolog consegue concluir que `bom_cacador(gato)` é verdadeiro.

---

## 13. Ada e os sistemas críticos

A linguagem **Ada** foi desenvolvida para atender principalmente às necessidades do Departamento de Defesa dos Estados Unidos, especialmente em sistemas embarcados e projetos de grande escala.

Como esses sistemas poderiam envolver consequências graves em caso de falhas, a linguagem foi projetada com foco em **confiabilidade, segurança e organização**.

* **Confiabilidade:** Ada possui mecanismos que ajudam a detectar erros durante a compilação, diminuindo a possibilidade de determinados problemas aparecerem somente durante a execução.

* **Tipos:** sua tipagem forte permite definir os dados de maneira mais precisa e restringir operações inadequadas. Isso ajuda o compilador a encontrar erros antes da execução.

* **Pacotes:** permitem organizar o código em módulos, separando interfaces e implementações. Isso facilita a manutenção e ajuda na organização de projetos grandes.

* **Concorrência:** Ada possui recursos próprios para trabalhar com tarefas concorrentes, permitindo coordenar diferentes atividades que precisam ocorrer simultaneamente. Isso é especialmente útil em sistemas embarcados que precisam responder a vários eventos.

Essas características estão diretamente relacionadas às necessidades de sistemas críticos, nos quais previsibilidade, segurança e confiabilidade são aspectos importantes.

---

## 15. Java e a mudança de contexto para a Web

A primeira aplicação do Java não foi a Web. A linguagem surgiu inicialmente dentro de um projeto voltado para dispositivos e sistemas embarcados. Porém, esse contexto inicial não apresentou o sucesso comercial esperado.

Com o crescimento da **Web**, surgiu uma nova oportunidade para a linguagem. O Java possuía algumas características que combinavam bem com esse novo cenário, principalmente a possibilidade de executar o mesmo programa em diferentes plataformas através da **JVM (Java Virtual Machine)**.

Os **applets** também tiveram um papel importante na popularização inicial do Java na Web.

Isso mostra como uma mudança no contexto tecnológico pode mudar completamente a posição de uma linguagem. O Java foi criado pensando em um determinado tipo de aplicação, mas acabou encontrando na Web um ambiente onde suas características se tornaram bastante interessantes.

Posteriormente, a linguagem passou a ser utilizada em diversos outros tipos de sistemas, incluindo aplicações empresariais e servidores.

---

## 17. C# e suas decisões em relação a Java e C++

O C# surgiu dentro do ambiente **.NET** e combinou características que já existiam em linguagens como C++ e Java, mas fazendo algumas escolhas diferentes.

Uma diferença entre **C# e Java** está no uso do comando `goto`. O Java não permite o uso de `goto` como comando de controle de fluxo, enquanto o C# mantém essa possibilidade, assim como C e C++.

Apesar de o uso excessivo de `goto` poder dificultar a leitura do código, sua existência oferece uma alternativa para alguns casos específicos de controle de fluxo.

Outra diferença aparece nos **enums**. O C# possui uma implementação mais controlada do que a tradicionalmente encontrada em C++, principalmente porque não permite simplesmente tratar qualquer valor de enum como um `int` sem uma conversão adequada. Nesse aspecto, existe uma preocupação maior com a segurança de tipos, semelhante à encontrada em Java.

Essas decisões mostram que o C# buscou aproveitar ideias de linguagens anteriores, mas fazendo ajustes para equilibrar **flexibilidade, segurança de tipos e organização do código**.
