// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Betting {
    struct Aposta {
        address user;
        uint valor;
        string escolha;
    }

    Aposta[] public apostas;
    uint public pote;
    string public resultado;
    bool public isFinished;

    // Registrar uma aposta
    function registrarAposta(string memory _escolha) public payable {
        require(msg.value > 0, "Aposta deve ter valor positivo");
        require(!isFinished, "O evento ja foi finalizado");

        // Adicionar aposta ao array
        apostas.push(Aposta(msg.sender, msg.value, _escolha));
        pote += msg.value;
    }

    // Finalizar o evento e definir o resultado
    function finalizar(string memory _resultado) public {
        require(!isFinished, "Evento ja finalizado");
        require(
            keccak256(abi.encodePacked(_resultado)) == keccak256(abi.encodePacked("cara")) ||
            keccak256(abi.encodePacked(_resultado)) == keccak256(abi.encodePacked("coroa")),
            "Resultado invalido"
        );

        resultado = _resultado;
        isFinished = true;

        // Calcular o total apostado pelos vencedores
        uint totalVencedores = 0;
        for (uint i = 0; i < apostas.length; i++) {
            if (keccak256(abi.encodePacked(apostas[i].escolha)) == keccak256(abi.encodePacked(resultado))) {
                totalVencedores += apostas[i].valor;
            }
        }

        // Distribuir prêmios
        for (uint i = 0; i < apostas.length; i++) {
            if (keccak256(abi.encodePacked(apostas[i].escolha)) == keccak256(abi.encodePacked(resultado))) {
                uint recompensa = (apostas[i].valor * pote) / totalVencedores;
                (bool sucesso, ) = apostas[i].user.call{value: recompensa}("");
                require(sucesso, "Falha ao transferir recompensa");
            }
        }
    }

    // Consultar o total de apostas no pote
    function getPote() public view returns (uint) {
        return pote;
    }
}
