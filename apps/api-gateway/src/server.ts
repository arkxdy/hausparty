import express, { Application, Request, Response } from "express";
import cors from 'cors'
import compression from 'compression'
import cookieParser from "cookie-parser";
import apiVersionRouting from "./routes/version_routing";
import { login } from "./controller/auth_controller";

const server: Application = express();

server.use(cors({
    origin: true,
    credentials: true,
}));

server.get('/', (_req: Request, res: Response) => {
    res.status(200).json({"status": "Server is thick"})
})

server.use(cookieParser())
server.use(express.json())
server.use(express.urlencoded({ extended: true}))
server.use(compression())

server.use('/api/v1/', apiVersionRouting)

export default server;