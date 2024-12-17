const { ethers } = require("hardhat");  // Certifique-se de que essa linha está presente

async function main() {
  const [deployer] = await ethers.getSigners();  // Obtendo a conta que será usada para deploy

  console.log("Deploying contracts with the account:", deployer.address);

  const Contract = await ethers.getContractFactory("Betting");  // Substitua "NomeDoContrato" pelo nome do seu contrato
  const contract = await Contract.deploy();

  console.log("Contract deployed to:", contract.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
