// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VultureBet {
    struct Aposta {
        address criador;    // Criador da aposta
        uint valorTotal;    // Pote total da aposta
        string resultado;   // Resultado da aposta ("cara" ou "coroa")
        bool finalizado;    // Status de finalização
        mapping(address => uint) valores; // Valores apostados por cada usuário
        mapping(address => string) escolhas; // Escolhas dos usuários
    }

    mapping(uint => Aposta) public apostas; // Lista de apostas por ID
    uint public apostaCount;                // Contador de IDs das apostas
    mapping(address => uint) public saldo;  // Saldo dos usuários

    // Evento para registrar o resultado da aposta
    event ApostaFinalizada(
        uint indexed idAposta,
        string resultado,
        uint valorTotal,
        uint premioDistribuido
    );

    // Criar uma nova aposta
    function criarAposta() public returns (uint) {
        apostaCount++;
        Aposta storage novaAposta = apostas[apostaCount];
        novaAposta.criador = msg.sender;
        return apostaCount;
    }

    // Registrar uma aposta em um evento
    function registrarAposta(uint _id, uint _valor, string memory _escolha) public {
        require(_valor > 0, "Valor da aposta deve ser positivo.");
        require(!apostas[_id].finalizado, "Aposta ja foi finalizada.");
        require(saldo[msg.sender] >= _valor, "Saldo insuficiente para apostar.");

        // Descontar o valor do saldo do usuário
        saldo[msg.sender] -= _valor;

        // Atualizar o registro da aposta
        Aposta storage aposta = apostas[_id];
        aposta.valores[msg.sender] += _valor;
        aposta.escolhas[msg.sender] = _escolha;
        aposta.valorTotal += _valor;
    }

    // Finalizar a aposta (somente pelo criador)
    function finalizarAposta(uint _id) public {
        Aposta storage aposta = apostas[_id];
        require(msg.sender == aposta.criador, "Somente o criador pode finalizar.");
        require(!aposta.finalizado, "Aposta ja foi finalizada.");
        require(aposta.valorTotal > 0, "Nenhuma aposta registrada.");

        // Decidir o resultado aleatoriamente
        uint random = uint(keccak256(abi.encodePacked(block.timestamp, block.difficulty))) % 2;
        aposta.resultado = random == 0 ? "cara" : "coroa";
        aposta.finalizado = true;

        // Calcular e distribuir o prêmio
        uint premioDistribuido = 0;
        for (uint i = 0; i < apostaCount; i++) {
            address apostador = msg.sender; // Endereço atual
            string memory escolha = aposta.escolhas[apostador];

            // Verificar se o apostador venceu
            if (keccak256(abi.encodePacked(escolha)) == keccak256(abi.encodePacked(aposta.resultado))) {
                uint premio = aposta.valores[apostador];
                premioDistribuido += premio;
                payable(apostador).transfer(premio);
            }
        }

        // Registrar o resultado no evento
        emit ApostaFinalizada(_id, aposta.resultado, aposta.valorTotal, premioDistribuido);
    }

    // Consultar o pote de uma aposta
    function getPote(uint _id) public view returns (uint) {
        return apostas[_id].valorTotal;
    }

    // Consultar saldo
    function getSaldo() public view returns (uint) {
        return saldo[msg.sender];
    }

    // Depositar saldo
    function depositar(uint _valor) public payable {
        require(_valor > 0, "Valor deve ser positivo.");
        saldo[msg.sender] += _valor;
    }

    // Sacar saldo
    function sacar(uint _valor) public {
        require(saldo[msg.sender] >= _valor, "Saldo insuficiente.");
        require(_valor > 0, "Valor deve ser positivo.");
        saldo[msg.sender] -= _valor;
        //payable(msg.sender).transfer(_valor);
    }
}