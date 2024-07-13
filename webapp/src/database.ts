import mysql, { RowDataPacket } from 'mysql2/promise';


interface User extends RowDataPacket {
    id: number;
    username: string;
    email: string;
    password: string;
}

interface Recipe extends RowDataPacket {
    id: number;
    name: string;
    description: string;
    ingredients: string;
}

// dotenv.config({path: "../.env_db"});
const pool = mysql.createPool({
    connectionLimit: 10,
    host: process.env.DATABASE_HOSTNAME,
    user: process.env.MYSQL_USER,
    password: process.env.MYSQL_PASSWORD,
    database: process.env.MYSQL_DATABASE,
});

async function createTableUsers() {
    try {
        await pool.query(`CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(64) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`);
    } catch (error) {
        throw error;
    }
}

async function registerUser(username: string, email: string, password: string): Promise<void> {
    try {
        await pool.query("INSERT INTO users (username, email, password) VALUES (?,?,?)", [username, email, password]);
    } catch (error) {
        throw error;
    }
}

async function getUserByEmail(email: string): Promise<User | null> {
    try {
        const [rows]: [User[], any] = await pool.query("SELECT * FROM users WHERE email = ?", [email]);
        return rows.length > 0 ? rows[0] : null;
    } catch (error) {
        throw error;
    }
}

async function getUserByUsername(username: string): Promise<User | null> {
    try {
        const [rows]: [User[], any] = await pool.query("SELECT * FROM users WHERE username = ?", [username]);
        return rows.length > 0 ? rows[0] : null;
    } catch (error) {
        throw error;
    }
}

async function getRecipes() {
    try {
        const [rows]: [Recipe[], any] = await pool.query("SELECT name, description, ingredients FROM recipes WHERE is_private = false");
        return rows.length > 0 ? rows : null;
    } catch (error) {
        throw error;
    }
}

export { Recipe, User, registerUser, getUserByEmail, getUserByUsername, createTableUsers, getRecipes };