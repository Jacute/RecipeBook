import { Request, Response } from "express";
import { comparePassword, hashPassword } from '../utils/utils';
import { encodeToken } from '../utils/jwt';
import * as db from '../database';
import * as config from '../config';


async function register(req: Request, res: Response) {
    const { username, email, password } = req.body;
    const hashedPassword = await hashPassword(password);

    try {
        let user: db.User | null = await db.getUserByUsername(username);
        if (user != null) {
            res.status(409).json({message: "Username already exists."});
            return;
        }
        user = await db.getUserByEmail(email);
        if (user != null) {
            res.status(409).json({message: "Email already exists."});
            return;
        }
    } catch (error) {
        console.error(error);
        res.status(500).json({message: "Internal Server Error."});
        return;
    }

    try {
        await db.registerUser(username, email, hashedPassword);
    } catch (error) {
        console.error(error);
        res.status(500).json({message: "Internal Server Error."});
        return;
    }
    req.session.username = username;
    const token = encodeToken(username);

    res
    .status(201)
    .cookie("auth-token", token, {maxAge: config.COOKIE_LIFETIME})
    .json({message: "User registered successfully."});
}

async function login(req: Request, res: Response) {
    const { username, password } = req.body;
    let user: db.User | null;
    try {
        user = await db.getUserByUsername(username);
    } catch (error) {
        console.error(error);
        res.status(500).json({message: "Internal Server Error."});
        return;
    }
    if (user == null || await comparePassword(password, user.password) === false) {
        res.status(401).json({message: "Invalid username or password."});
        return;
    }
    req.session.username = username;
    const token = encodeToken(username);

    res.status(200)
    .cookie("auth-token", token, {maxAge: config.COOKIE_LIFETIME})
    .json({message: "User logged in successfully."});;
}

async function logout(req: Request, res: Response) {
    res.clearCookie("auth-token")
    .redirect("/");
}

export { register, login, logout };