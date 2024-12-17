const { Web3 } = require("web3");

// Conexão com o nó Ethereum (ajuste o URL conforme necessário)
const web3 = new Web3("http://127.0.0.1:8545");

async function checkBalance() {
    const accounts = await web3.eth.getAccounts();
    const balanceWei = await web3.eth.getBalance(accounts[0]);
    const balanceEth = web3.utils.fromWei(balanceWei, "ether");
    console.log(`Saldo da conta (${accounts[0]}): ${balanceEth} ETH`);
}

checkBalance();


// ABI do contrato
const abi = [
    {
        "inputs": [],
        "name": "criarAposta",
        "outputs": [{ "name": "", "type": "uint256" }],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "depositar",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [{ "name": "_id", "type": "uint256" }],
        "name": "finalizarAposta",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [{ "name": "_id", "type": "uint256" }, { "name": "_escolha", "type": "string" }],
        "name": "registrarAposta",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [{ "name": "_valor", "type": "uint256" }],
        "name": "sacar",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "apostaCount",
        "outputs": [{ "name": "", "type": "uint256" }],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [{ "name": "", "type": "uint256" }],
        "name": "apostas",
        "outputs": [
            { "name": "criador", "type": "address" },
            { "name": "valorTotal", "type": "uint256" },
            { "name": "resultado", "type": "string" },
            { "name": "finalizado", "type": "bool" }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [{ "name": "_id", "type": "uint256" }],
        "name": "getPote",
        "outputs": [{ "name": "", "type": "uint256" }],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getSaldo",
        "outputs": [{ "name": "", "type": "uint256" }],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [{ "name": "", "type": "address" }],
        "name": "saldo",
        "outputs": [{ "name": "", "type": "uint256" }],
        "stateMutability": "view",
        "type": "function"
    }
];

// Endereço do contrato
const contractAddress = "0xAeD0A1B21A0baB9f9EfAeF3E50B749ad985B1409";

// Instanciar o contrato
const contract = new web3.eth.Contract(abi, contractAddress);

// Exemplo: chamar uma função
async function main() {
    const accounts = await web3.eth.getAccounts();
    const result = await contract.methods.criarAposta().send({ from: accounts[0] });
    console.log("Aposta criada! Transação:", result.transactionHash);
}

main().catch(console.error);
