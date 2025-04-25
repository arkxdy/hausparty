import { Request, Response } from "express";
import crypto from 'crypto'

const login = async (req: Request, res: Response) => {
    try{
        res.status(200).json({"success": true, data: {}})
    } catch (err) {
        res.status(501).json({"success": true, data: {}})
    }
}
    
const register = async (req: Request, res: Response) => {
    try{

        res.status(201).json({"success": true, data: {}})
    } catch (err) {
        res.status(501).json({"success": true, data: {}})
    }
}

const authProvider = async (req: Request, res: Response) => {
    try{
        // 1. Validate provider parameter
        const provider = req.params["provider"]

        // 2. Validate redirect_uri format
        const redirectUri = req.query.redirect_uri as string;
        console.log(req.query)
        if (!redirectUri) {
            res.status(400).json({ error: 'Missing redirect_uri' });
            return
        }
    
        // 3. Generate and store state
        const state = crypto.randomBytes(16).toString('hex');
        //stateStore.set(state, redirectUri);
    
        // 4. Redirect to Auth Service
        const authServiceUrl = new URL(
            `${process.env.AUTH_SERVICE_URL}/auth/${provider}`
        );
        authServiceUrl.searchParams.append('redirect_uri', redirectUri);
        authServiceUrl.searchParams.append('state', state);
        console.log("req", authServiceUrl.toString())
        res.redirect(authServiceUrl.toString());
    } catch (err) {
        console.error('Auth initiation error:', err);
        res.status(500).json({ error: 'Authentication failed' });
    }
}

const authProviderCallback = async (req: Request, res: Response) => {
    try {
        console.log("in callback")
        // 1. Validate state parameter
        const provider = req.params["provider"]
        console.log(provider)
        // const state = req.query.state as string;
        // if (!state /*|| !stateStore.has(state)*/) {
        //   res.status(400).json({ error: 'Invalid state token' });
        //   return
        // }
    
        // 2. Get original redirect_uri
        const redirectUri = new URL(
            `${process.env.AUTH_SERVICE_URL}/auth/${provider}/callback`
        );
        //stateStore.delete(state);
        console.log("error", req.query.error)
        // 3. Validate error parameters
        if (req.query.error) {
          const error = new Error(req.query.error_description as string);
          res.redirect(`${redirectUri}?error=${error.message}`);
          return
        }
    
        // 4. Forward to frontend
        redirectUri.searchParams.append('redirect_uri', `${process.env.API_GATEWAY_URL}/auth/${provider}/callback`);
        redirectUri.searchParams.append('state', req.query.state as string);
        redirectUri.searchParams.append('code', req.query.code as string);
        console.log(redirectUri.toString())
        res.redirect(
          redirectUri.toString()
        );
      } catch (error) {
        console.error('Callback validation error:', error);
        res.status(500).json({ error: 'Callback processing failed' });
      }
}

export {
    login,
    register,
    authProvider,
    authProviderCallback
}