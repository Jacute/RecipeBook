import { Router } from "express";
import { UserMiddleware, AuthMiddleware } from "./middlewares/auth";
import { index, create, recipe } from "./handlers/other";
import { register, login, logout } from "./handlers/auth";

const router: Router = Router();
router.get('*', UserMiddleware);

router.get('/', index);
router.get('/recipe', recipe);

router.use('/create', AuthMiddleware);
router.get('/create', create);

router.post('/register', register);
router.post('/login', login);
router.get('/logout', logout);

export { router };