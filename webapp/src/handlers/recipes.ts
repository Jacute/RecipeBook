import { Request, Response } from "express";
import * as db from "../database";


async function index(req: Request, res: Response) {
    let recipes: db.Recipe[] | null;
    try {
        recipes = await db.getRecipes();
    } catch (err) {
        recipes = null;
    }
    res.render("index", {username: req.session.username, recipes: recipes, nonce: res.locals.nonce});
}

function recipe(req: Request, res: Response) {
    res.render("recipe", {username: req.session.username, nonce: res.locals.nonce});
}

function create(req: Request, res: Response) {
    res.render("create", {username: req.session.username});
}

export { index, create, recipe };