"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserMiddleware = UserMiddleware;
exports.AuthMiddleware = AuthMiddleware;
const jwt_1 = require("../utils/jwt");
function AuthMiddleware(req, res, next) {
    if (req.session.username == null) {
        res.redirect('/');
        return;
    }
    next();
}
;
function UserMiddleware(req, res, next) {
    if (req.session.username == null) {
        const token = req.cookies["auth-token"];
        if (req.cookies["auth-token"] != null) {
            const { tokenExp, error, decoded } = (0, jwt_1.verifyToken)(token);
            if (!tokenExp && error == null) {
                req.session.username = decoded.username;
            }
        }
    }
    next();
}
