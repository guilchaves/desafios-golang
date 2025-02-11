#  DESAFIO: Componentes e injeção de dependência | Go


### Desafio
Você deve criar um sistema para calcular o **valor total** de um pedido, considerando uma porcentagem
de desconto e o frete. O cálculo do **valor total** do pedido consiste em aplicar o desconto ao **valor
básico** do pedido e adicionar o valor do frete. A regra para cálculo do frete é a seguinte: </br>


Valor básico do pedido (sem desconto) | Frete 
---|---
Abaixo de R$ 100.00 | R$ 20.00
De R$ 100.00 até R$ 200.00 exclusive | R$ 12.00
R$ 200.00 ou mais | Grátis

_**Exemplo 1**_</br>

**ENTRADA** (dados do pedido: código, valor básico, porcentagem de desconto)</br>
```
1034
150.0
20.0
```
**SAÍDA**
```
Pedido codigo 1034
Valor total: R$ 132.00
```

_**Exemplo 2**_</br>

**ENTRADA** (dados do pedido: código, valor básico, porcentagem de desconto)</br>
```
2282
800.00
10.0
```
**SAÍDA**
```
Pedido codigo 2282
Valor total: R$ 720.00
```
_**Exemplo 3**_</br>

**ENTRADA** (dados do pedido: código, valor básico, porcentagem de desconto)</br>
```
1309
95.90
0.0
```
**SAÍDA**
```
Pedido codigo 1309 
Valor total: R$ 115.90
```

