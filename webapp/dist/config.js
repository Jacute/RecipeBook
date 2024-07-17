"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.RATELIMIT_PERIOD = exports.RATELIMIT_COUNT = exports.COOKIE_LIFETIME = exports.TOKEN_LIFETIME = exports.JWT_PUBLIC_KEY = exports.JWT_PRIVATE_KEY = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const JWT_PRIVATE_KEY = fs_1.default.readFileSync(path_1.default.join(__dirname, '../jwt_keys/private.pem'));
exports.JWT_PRIVATE_KEY = JWT_PRIVATE_KEY;
const JWT_PUBLIC_KEY = fs_1.default.readFileSync(path_1.default.join(__dirname, '../jwt_keys/public.pem'));
exports.JWT_PUBLIC_KEY = JWT_PUBLIC_KEY;
const TOKEN_LIFETIME = "60 days";
exports.TOKEN_LIFETIME = TOKEN_LIFETIME;
const COOKIE_LIFETIME = 1000 * 60 * 60 * 24 * 60;
exports.COOKIE_LIFETIME = COOKIE_LIFETIME;
const RATELIMIT_COUNT = 500;
exports.RATELIMIT_COUNT = RATELIMIT_COUNT;
const RATELIMIT_PERIOD = 10 * 60 * 1000; // 10 minutes
exports.RATELIMIT_PERIOD = RATELIMIT_PERIOD;
