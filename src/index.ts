import * as dotenv from "dotenv";
import startBot from "./bot.js";
import server from "./server.js";

dotenv.config();

console.log("Start bot");

startBot().catch((err) => {
  console.error("Failed to start the bot", err);
});

server();
