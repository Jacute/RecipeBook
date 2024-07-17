"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AuthMiddleware = AuthMiddleware;
function AuthMiddleware(req, res, next) {
    if (req.session.username == null) {
        res.redirect('/');
        return;
    }
    next();
}
;
