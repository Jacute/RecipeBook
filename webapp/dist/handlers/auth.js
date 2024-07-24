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
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.register = register;
exports.login = login;
exports.logout = logout;
const utils_1 = require("../utils/utils");
const jwt_1 = require("../utils/jwt");
const db = __importStar(require("../database"));
const config = __importStar(require("../config"));
function register(req, res) {
    return __awaiter(this, void 0, void 0, function* () {
        const { username, email, password } = req.body;
        const hashedPassword = yield (0, utils_1.hashPassword)(password);
        try {
            let user = yield db.getUserByUsername(username);
            if (user != null) {
                res.status(409).json({ message: "Username already exists." });
                return;
            }
            user = yield db.getUserByEmail(email);
            if (user != null) {
                res.status(409).json({ message: "Email already exists." });
                return;
            }
        }
        catch (error) {
            console.error(error);
            res.status(500).json({ message: "Internal Server Error." });
            return;
        }
        try {
            yield db.registerUser(username, email, hashedPassword);
        }
        catch (error) {
            console.error(error);
            res.status(500).json({ message: "Internal Server Error." });
            return;
        }
        req.session.username = username;
        const token = (0, jwt_1.encodeToken)(username);
        res
            .status(201)
            .cookie("auth-token", token, { maxAge: config.COOKIE_LIFETIME })
            .json({ message: "User registered successfully." });
    });
}
function login(req, res) {
    return __awaiter(this, void 0, void 0, function* () {
        const { username, password } = req.body;
        let user;
        try {
            user = yield db.getUserByUsername(username);
        }
        catch (error) {
            console.error(error);
            res.status(500).json({ message: "Internal Server Error." });
            return;
        }
        if (user == null || (yield (0, utils_1.comparePassword)(password, user.password)) === false) {
            res.status(401).json({ message: "Invalid username or password." });
            return;
        }
        req.session.username = username;
        const token = (0, jwt_1.encodeToken)(username);
        res.status(200)
            .cookie("auth-token", token, { maxAge: config.COOKIE_LIFETIME })
            .json({ message: "User logged in successfully." });
        ;
    });
}
function logout(req, res) {
    return __awaiter(this, void 0, void 0, function* () {
        res.clearCookie("auth-token")
            .redirect("/");
    });
}
