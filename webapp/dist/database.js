"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.registerUser = registerUser;
exports.getUserByEmail = getUserByEmail;
exports.getUserByUsername = getUserByUsername;
exports.createTableUsers = createTableUsers;
exports.getRecipes = getRecipes;
const promise_1 = __importDefault(require("mysql2/promise"));
// dotenv.config({path: "../.env_db"});
const pool = promise_1.default.createPool({
    connectionLimit: 10,
    host: process.env.DATABASE_HOSTNAME,
    user: process.env.MYSQL_USER,
    password: process.env.MYSQL_PASSWORD,
    database: process.env.MYSQL_DATABASE,
});
function createTableUsers() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            yield pool.query(`CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(64) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`);
        }
        catch (error) {
            throw error;
        }
    });
}
function registerUser(username, email, password) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            yield pool.query("INSERT INTO users (username, email, password) VALUES (?,?,?)", [username, email, password]);
        }
        catch (error) {
            throw error;
        }
    });
}
function getUserByEmail(email) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const [rows] = yield pool.query("SELECT * FROM users WHERE email = ?", [email]);
            return rows.length > 0 ? rows[0] : null;
        }
        catch (error) {
            throw error;
        }
    });
}
function getUserByUsername(username) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const [rows] = yield pool.query("SELECT * FROM users WHERE username = ?", [username]);
            return rows.length > 0 ? rows[0] : null;
        }
        catch (error) {
            throw error;
        }
    });
}
function getRecipes() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const [rows] = yield pool.query("SELECT name, description, ingredients FROM recipes WHERE is_private = false");
            return rows.length > 0 ? rows : null;
        }
        catch (error) {
            throw error;
        }
    });
}
