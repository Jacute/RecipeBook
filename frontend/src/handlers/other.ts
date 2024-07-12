import { Request, Response } from "express";
import * as db from "../database";


async function index(req: Request, res: Response) {
    let recipes: db.Recipe[] | null;
    try {
        recipes = await db.getRecipes();
    } catch (err) {
        recipes = null;
    }
    res.render("index", {username: req.session.username, recipes: recipes});
}

function receipe(req: Request, res: Response) {
    res.render("receipe", {username: req.session.username});
}

function create(req: Request, res: Response) {
    res.render("create", {username: req.session.username});
}

export { index, create, receipe };