// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Betting {
    struct Aposta {
        address user;  // Endere\u00e7o do apostador
        uint valor;    // Valor apostado
        string escolha; // Escolha do resultado: "cara" ou "coroa"
    }

    Aposta[] public apostas; // Lista de apostas
    uint public pote;        // Valor total no pote
    string public resultado; // Resultado do evento
    bool public isFinished;  // Indica se o evento foi finalizado

    // Registrar uma aposta
    function registrarAposta(string memory _escolha) public payable {
        require(msg.value > 0, "Aposta deve ter valor positivo");
        require(!isFinished, "O evento ja foi finalizado");

        // Registra a nova aposta
        apostas.push(Aposta(msg.sender, msg.value, _escolha));
        pote += msg.value;
    }

    // Finalizar o evento e definir o resultado
    function finalizar(string memory _resultado) public {
        require(!isFinished, "Evento ja foi finalizado");
        require(
            keccak256(abi.encodePacked(_resultado)) == keccak256(abi.encodePacked("cara")) ||
            keccak256(abi.encodePacked(_resultado)) == keccak256(abi.encodePacked("coroa")),
            "Resultado invalido"
        );

        resultado = _resultado;
        isFinished = true;

        // Distribuir os pr\u00eamios
        uint totalVencedores = 0;

        // Calcula o total apostado pelos vencedores
        for (uint i = 0; i < apostas.length; i++) {
            if (keccak256(abi.encodePacked(apostas[i].escolha)) == keccak256(abi.encodePacked(resultado))) {
                totalVencedores += apostas[i].valor;
            }
        }

        // Distribui o pote proporcionalmente aos vencedores
        for (uint i = 0; i < apostas.length; i++) {
            if (keccak256(abi.encodePacked(apostas[i].escolha)) == keccak256(abi.encodePacked(resultado))) {
                uint recompensa = (apostas[i].valor * pote) / totalVencedores;
                payable(apostas[i].user).transfer(recompensa);
            }
        }
    }

    // Consultar o total de apostas no pote
    function getPote() public view returns (uint) {
        return pote;
    }
}

