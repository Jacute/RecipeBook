"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.encodeToken = encodeToken;
exports.verifyToken = verifyToken;
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
const config = __importStar(require("../config"));
function encodeToken(username) {
    const payload = {
        username: username
    };
    const options = {
        algorithm: "RS256",
        expiresIn: config.TOKEN_LIFETIME
    };
    const token = jsonwebtoken_1.default.sign(payload, config.JWT_PRIVATE_KEY, options);
    return token;
}
function verifyToken(token) {
    const options = {
        algorithms: ["RS256"],
        maxAge: config.TOKEN_LIFETIME
    };
    try {
        const decoded = jsonwebtoken_1.default.verify(token, config.JWT_PUBLIC_KEY, options);
        return { tokenExp: false, error: null, decoded: decoded };
    }
    catch (err) {
        return { tokenExp: true, error: err, decoded: null };
    }
}
