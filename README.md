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

A escalabilidade do sistema é limitada principalmente pela capacidade de processamento de transações da blockchain privada, que depende do muito do tamanho dos blocos. Por outro lado, a estabilidade é garantida pelo uso de Go-Ethereum, que lida automaticamente com mineração, consenso e persistência dos dados, reduzindo o risco de falhas no processamento das transações.

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

<div id="conclusao">
<h2> Conclusão </h2>
<div align="justify">

O projeto proporciona uma solução descentralizada para um sistema de apostas, e a utilização da tecnologia blockchain garantiu a imutabilidade e a segurança dos dados. Os contratos inteligentes viabilizaram a gestão automatizada das apostas e pagamentos, e a estrutura é escalável para redes privadas de pequeno a médio porte podendo ser facilmente adaptada para incluir novos recursos ou operações.

Porém, a solução pode enfrentar limitações em cenários de alta demanda devido à capacidade de processamento da blockchain, que é sempre bom lembrar que acaba por criar problemas de grande complexidade que podem se provar desafiadores para hardwares mais limitados. No geral, a solução demonstrou ser funcional, robusta e alinhada aos princípios fundamentais da tecnologia blockchain.


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