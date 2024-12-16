require("dotenv").config();

module.exports = {
  solidity: "0.8.18",
  networks: {
    sepolia: {
      url: process.env.RPC_URL,
      accounts: {
        mnemonic: process.env.MNEMONIC, // Usando a mnemonic para gerar as contas
      },
    },
  },
};
