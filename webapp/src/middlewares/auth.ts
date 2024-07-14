import { Request, Response, NextFunction } from "express";
import { verifyToken } from "../utils/jwt";

function AuthMiddleware(req: Request, res: Response, next: NextFunction) {
    if (req.session.username == null) {
        res.redirect('/');
        return;
    }
    next();
};

function UserMiddleware(req: Request, res: Response, next: NextFunction) {
    if (req.session.username == null) {
        const token = req.cookies["auth-token"];
        if (req.cookies["auth-token"] != null) {
            const { tokenExp, error, decoded } = verifyToken(token);
            if (!tokenExp && error == null) {
                req.session.username = decoded.username;
            }
        }
    }
    next();
}

export { UserMiddleware, AuthMiddleware };