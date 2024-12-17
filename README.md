<h4 align="center">VultureBet</h4>

<div align="justify"> 
<div id="sobre-o-projeto"> 
<h2> Descrição do Projeto</h2>

O projeto desenvolvido consiste em um sistema de apostas descentralizado baseado em uma blockchain privada. A ideia principal é permitir que usuários realizem apostas sobre o resultado do lançamento de moedas ("cara" ou "coroa"), onde a soma dos valores apostados é distribuído proporcionalmente entre os vencedores. Este sistema utiliza contratos inteligentes escritos em solidity para gerenciar toda a lógica das apostas, desde o registro até a distribuição dos prêmios.

O objetivo principal é explorar conceitos fundamentais de blockchain, como descentralização, imutabilidade, privacidade e transparência, aplicados a uma situação prática. Além disso, fornece uma solução escalável que pode ser utilizada por múltiplos participantes, sem depender de um servidor central, garantindo a segurança no processamento das apostas.

</div>
</div>

<h2> Autores <br></h2>
<uL>
  <li><a href="https://github.com/felipe-py">Felipe Silva</a></li>
  <li><a href="https://github.com/Lucas-L-Rodrigues">Lucas Lima</a></li>
</ul>


<h1 align="center"> Sumário </h1>
<div id="sumario">
	<ul>
        <li><a href="#arquitetura"> Arquitetura da Solução </a></li>
        <li><a href="#funcionamento"> Funcionamento do Sistema </a></li>
        <li><a href="#confiabilidade"> Confiabilidade da Solução </a></li>
        <li><a href="#blockchain"> Blockchain </a></li>
        <li><a href="contrato"> Contrato </a></li>
        <li><a href="#conclusao"> Conclusão </a></li>
        <li><a href="#referencias"> Referências </a></li>
	</ul>	
</div>

</div id="arquitetura">
<h2> Arquitetura da Solução </h2>
<div align="justify">

O sistema é baseado em uma blockchain privada com o uso de contratos inteligentes escritos em Solidity, e uma interface de interação desenvolvida em GO utilizando a biblioteca go-ethereum (tambem conhecida como Geth). A solução funciona em três camadas principais:

<h4> Blockchain Privada </h4>

A blockchain é criada utilizando o cliente Go-Ethereum (Geth). Ela é configurada com um arquivo “genesis.json”, que define os parâmetros iniciais. A blockchain é mantida
localmente por cada nó participante, que minera blocos para processar as transações relacionadas às apostas.

<h4> Contratos Inteligentes </h4>

Os contratos inteligentes encapsulam a lógica das apostas. Eles são escritos em Solidity e implementam funcionalidades como: registro de apostas com os devidos valores 
apostados; validação das condições para finalizar o evento e o cálculo (e distribuição) do pote aos vencedores.

<h4> Interação com a Blockchain </h4>

Os contratos são implantados na blockchain privada utilizando bytecode e ABI gerados pelo compilador Solidity. A interação com o sistema é feita via API JSON-RPC do Geth,
utilizando a biblioteca go-ethereum. Essas APIs permitem que os usuários enviem transações, consultem o estado da blockchain e interajam com os contratos inteligentes.

</div>

<div id="funcionamento">
<h2> Funcionamento do Sistema </h2>
<div align="justify">

<h4> Registro de apostas </h4>

* O usuário envia uma transação ao contrato chamando a função placeBet, indicando sua escolha (“cara” ou “coroa”) e o valor da aposta.
* A transação é registrada na blockchain como parte do bloco seguinte.

* O contrato atualiza o pote total das apostas e registra os dados do usuário.

<h4> Finalização de aposta </h4>

* O administrador chama a função finalizeEvent, especificando o resultado do evento.

* O contrato calcula o total apostado pelos vencedores e distribui o pote proporcionalmente aos seus valores apostados.

* Os pagamentos são processados automaticamente utilizando a função transfer do Ethereum.

* Validação, segurança e imutabilidade

Toda a validação do sistema é feita pelos nós do blockchain, antes da inclusão no bloco. É verificado se o emissor da transação possui saldo suficiente para iniciar buma operação, se a transação assinada possui autenticidade e a validade do contrato inteligente.

A segurança do sistema de maneira geral é garntida pelas várias camadas fornecidas pelo modelo blockchain. Hashes são utilizados para forma uma cadeia inquebrável e descentralizada, onde cada nó mantem uma cópia completa do blockchain.

A imutabilidade do sistema é fornecidade de maneira geral pelo uso do contrato inteligente, assegurando que as regras para as apostas sejam executadas sema  possibilidade de manipulação.

</div>
</div>

<div id="confiabilidade">
<h2> Confiabilidade da Solução </h2>
<div align="justify">

A solução foi projetada para suportar múltiplos nós em simultâneo, permitindo que diferentes máquinas participem da rede distribuída. Cada nó mantém uma cópia local da blockchain, processando transações independentemente e sincronizando periodicamente com os demais nós da rede.

O projeto entretanto, não foi integrado em sua totalidade, ou seja, não é possível atestar sua confiabilidade de maneira completa.
</div>
</div>

<div id="blockchain">
<h2> Blockchain </h2>
<div align="justify">

<h4> Definição </h4>

Blockchain é uma estrutura de dados distribuída e imutável que organiza informações em blocos interligados por endereços codificados por hashes. Cada bloco armazena um conjunto de 
informações (como transações, por exemplo) e um hash que referencia o bloco anterior, formando uma cadeia, daí o nome blockchain. Sua natureza descentralizada e transparente a 
torna ideal para aplicações onde é necessário garantir integridade, privacidade e segurança.
	
<h4> Como foi usada </h4>

Neste projeto, a blockchain foi empregada para:
    - Registrar transações de apostas e pagamentos.
    - Garantir que todos os dados sejam imutáveis e auditáveis.
    - Manter uma cópia descentralizada e sincronizada em todos os nós.
		
<h4> Ferramentas empregados </h4>

Para o emprego da blockchain nesse projeto foram utilizadas as seguintes ferramentas de terceiros:
    - Go-Ethereum (Geth): Para criação e gerenciamento da blockchain privada.
    - Solidity: Para desenvolver os contratos inteligentes.
    - go-ethereum: Biblioteca GO para interagir com a blockchain.
    - solc: Compilador de Solidity.


</div>
</div>

<div id="contrato">
<h2> Contrato </h2>
<div align="justify">

A solução utiliza um smart contract que é um programa autoexecutável baseado em blockchain, contendo as regras e acordos necessários entre as partes envolvidas codificados em 'Solidity'. Ele garante segurança e transparência, eliminando intermediários, e no caso de contratos em 'Solidity', eles são implementados na blockchain Ethereum, garantindo que as transações sejam imutáveis e verificáveis.
O VultureBet gerencia apostas de forma descentralizada, permitindo a criação de apostas, registro de participantes e distribuição de prêmios de acordo com o resultado. A blockchain registra todas as operações, protegendo a integridade do sistema.

O contrato está estruturado para criar um sistema de apostas descentralizado, nele estão contidas variáveis como 'apostas' (um mapeamento de estruturas de apostas) e saldo de cada usuário. Há ainda funções para criar apostas, registrar participantes e finalizar o evento, distribuindo os valores de acordo com os resultados. Alguns eventos, como ApostaFinalizada, são emitidos para notificar as interações importantes.

Para poder usar o contrato em uma blockchain são necessários 2 arquivos: o 'ABI' e o binário. Ambos podem ser gerados ao compilar o contrato se utilizando do compilador 'solc' com o seguinte comando:
								solc --optimize --bin --abi contrato_bet.sol -o ./build
	
	1.1.Variáveis e estruturas do contrato:
		1.1.1.Aposta:
		estrutura que define uma aposta contendo as seguintes informações: criador da aposta; valor total acumulado no pote; resultado do evento; se foi finalizado; valores apostados pelos usuários e quantias apostadas pelos usuários

		1.1.2.apostas:
		uma lista das apostas feitas naquele evento

		1.1.3.apostaCount:
		um contador para o número da aposta, usado para garantir que nenhuma aposta tenha o ID de outra

		1.1.3.saldo:
		uma lista contendo o saldo dos usuários
	

	1.2. Funções do contrato:
	As funções descrevem uma série de procedimentos que são executados de acordo com a política do contrato, cada uma pode receber alguns parâmetros e retornar alguns resultados, as funções do contrato seguem:
		1.2.1.criarAposta:
		Não precisa de nenhuma entrada e retorna apenas o ID único da nova aposta que foi gerada permitindo ao usuário criar um novo evento em que podem ser realizadas as apostas
		
		1.2.2.registrarAposta:
		Recebe de entrada o ID do evento em que se deseja apostar, o valor a ser apostado e o resultado em que se quer apostar. Não retorna qualquer valor, contudo cria uma nova aposta para aquele evento no nome do usuário que esteja apostando

		1.2.3.finalizarAposta:
		Recebe o ID da aposta em que se deseja finalizar. Caso o usuário seja o criador da aposta essa função determina o resultado da aposta e distribui o prêmio entre os vencedores. Ela retorna os dados: o ID; o resultado da aposta; o valor do pote e o valor distribuído

		1.2.4.getPote:
		Recebe o ID da aposta e retorna o valor acumulado no pote da mesma.

		1.2.5.getSaldo:
		Não recebe parâmetros e retorna o saldo que um dado usuário tem em sua conta

		1.2.6.depositar:
		Recebe o valor a ser depositado e o adiciona no saldo do usuário

		1.2.7.sacar:
		Recebe o valor a ser sacado e o deduz do saldo

	1.3.Eventos do contrato:
	Eventos são estruturas cujo setido é autoexplicativo e descrevem acontecimentos durante a execução da blockchain
		
		1.3.1.ApostaFinalizada:
		finaliza a aposta estabelecendo um resultado e a distribuição do pote estre os participantes. Ocorre quando o dono da aposta determina

</div>
</div>

<div id="conclusao">
<h2> Conclusão </h2>
<div align="justify">

O projeto cumpriu de maneira limitada o que foi proposto, devido a não integração de todos os elementos presentes no desenvolvimento.
Cada tópico desenvolvido foi testado e executado de forma independente, entretanto, devido a problemas de dependência sua integração não foi realizada.

É interessante que a solução seja totalmente integrada, atestando assim o cumprimento de cada funcionalidade em sua forma individual e conjunta.

</div>
</div>

<div id="referencias">  
<h2> Referências</h2>
<div align="justify">

Ethereum Foundation. (2023). Ethereum Documentation. Disponível em: https://ethereum.org/

Go-Ethereum (Geth). (2023). Official Documentation. Disponível em: https://geth.ethereum.org/docs/

Solidity. (2023). Solidity Documentation. Disponível em: https://docs.soliditylang.org/

</div>
</div>