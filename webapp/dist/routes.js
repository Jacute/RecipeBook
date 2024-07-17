"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.router = void 0;
const express_1 = require("express");
const auth_1 = require("./middlewares/auth");
const other_1 = require("./handlers/other");
const auth_2 = require("./handlers/auth");
const router = (0, express_1.Router)();
exports.router = router;
router.get('*', auth_1.UserMiddleware);
router.get('/', other_1.index);
router.get('/recipe', other_1.recipe);
router.use('/create', auth_1.AuthMiddleware);
router.get('/create', other_1.create);
router.post('/register', auth_2.register);
router.post('/login', auth_2.login);
router.get('/logout', auth_2.logout);
