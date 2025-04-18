import express from "express";
import dotenv from "dotenv";

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3001;
console.log(process.env.API_GATEWAY)
app.get("/", (_, res) => {
  res.send("API Gateway is working!");
});

app.listen(PORT, () => {
  console.log(`API Gateway listening on port ${PORT}`);
});
