import { Request, Response, NextFunction } from "express";

function AuthMiddleware(req: Request, res: Response, next: NextFunction) {
    if (req.session.username == null) {
        res.redirect('/');
        return;
    }
    next();
};

export { AuthMiddleware };