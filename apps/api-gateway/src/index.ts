import server from "./server";
import dotenv from "dotenv";

dotenv.config();

const app = server;

const PORT = process.env.API_GATEWAY_PORT || 3001;

app.listen(PORT, () => {
  console.log(`API Gateway listening on port ${PORT}`);
});
