# A linguagem escolhida foi C 

- A notacao utilizada pela linguagem: BNF e versoes proximas como: ISO/IEC C

- https://learn.microsoft.com/pt-br/cpp/c-language/lexical-grammar?view=msvc-170 - link documentacao

## Regras utilizadas 

- Os simbolos entre < > sao nao terminais; os demais representam terminais da linguagem.

```C
<if_statement> ::= if ( <condition> ) <statement> else <statement>

<condition> ::= <expression> > <expression>

<statement> ::= <assignment_statement> | <compound_statement>

<assignment_statement> ::= <identifier> = <expression> ;

<compound_statement> ::= { <statement_list> }

<statement_list> ::= <statement> | <statement> <statement_list>

<expression> ::= <identifier> | <division_expression> | <integer_literal>

<division_expression> ::= <expression> / <expression>

<identifier> ::= i | x | y

<integer_literal> ::= 0
```

## Derivacao 

<if_statement>

⇒ if ( <condition> ) <statement> else <statement>

⇒ if ( <expression> > <expression> )
   <statement>
   else
   <statement>

⇒ if ( i > <expression> )
   <statement>
   else
   <statement>

⇒ if ( i > 0 )
   <statement>
   else
   <statement>

⇒ if ( i > 0 )
   <assignment_statement>
   else
   <statement>

⇒ if ( i > 0 )
   y = <expression> ;
   else
   <statement>

⇒ if ( i > 0 )
   y = <division_expression> ;
   else
   <statement>

⇒ if ( i > 0 )
   y = x / i ;
   else
   <compound_statement>

⇒ if ( i > 0 )
   y = x / i ;
   else
   { <statement_list> }

⇒ if ( i > 0 )
   y = x / i ;
   else
   { <statement> }

⇒ if ( i > 0 )
   y = x / i ;
   else
   { <assignment_statement> }

⇒ if ( i > 0 )
   y = x / i ;
   else
   { x = <expression> ; }

⇒ if ( i > 0 )
   y = x / i ;
   else
   { x = i ; }


## Codigo concreto

```C
 if ( i > 0 )
    y = x / i;
else
{
    x = i;
}
```

