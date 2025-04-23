import { Router } from "express";
import { authProvider, authProviderCallback, login, register } from "../controller/auth_controller";

const authRoute: Router = Router();
authRoute.get('/:provider', authProvider);
authRoute.get('/:provider/callback', authProviderCallback);
authRoute.post('/login', login);
authRoute.post('/register', register);

export default authRoute;