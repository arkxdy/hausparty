import { Request, Response, NextFunction } from 'express';
import crypto from 'crypto'
import { URL } from 'url';

const validateGoogleAuth = (req: Request, res: Response, next: NextFunction) => {
    try {
        const redirectUri = req.query.redirect_uri as string;
        if (!redirectUri) {
            return res.status(400).json({ error: 'redirect_uri is required' });
        }

        // 3. Prevent open redirect attacks
        // const allowedDomains = process.env.ALLOWED_REDIRECT_DOMAINS?.split(',') || [];
        // const isValidRedirect = allowedDomains.some(domain => {
        //     try {
        //         return new URL(redirectUri).hostname === domain;
        //     } catch {
        //         return false;
        //     }
        // });
        
        // if (!isValidRedirect) {
        //     return res.status(400).json({ error: 'Invalid redirect URI' });
        // }

        // // 4. Generate and store CSRF token (state)
        const state = crypto.randomBytes(16).toString('hex');
        // storeStateToken(state, req.sessionID);

        // Attach validated parameters to request
        res.locals.oauthParams = {
            redirect_uri: redirectUri,
            scopes: req.query.scope || 'profile email'
        };

        next();
    } catch (err) {
        console.error('Validation error:', err);
        res.status(500).json({ error: 'Authentication failed' });
    }
}


export {
    validateGoogleAuth
}