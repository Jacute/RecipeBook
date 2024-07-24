"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const routes_1 = require("./routes");
const path_1 = __importDefault(require("path"));
const cookie_parser_1 = __importDefault(require("cookie-parser"));
const express_session_1 = __importDefault(require("express-session"));
const body_parser_1 = __importDefault(require("body-parser"));
const database_1 = require("./database");
const helmet_1 = __importDefault(require("helmet"));
const app = (0, express_1.default)();
app.use((0, cookie_parser_1.default)());
app.use((0, express_session_1.default)({
    secret: crypto.randomUUID(),
    resave: false,
    saveUninitialized: true,
    cookie: {
        secure: true,
        httpOnly: true,
        sameSite: 'strict'
    }
}));
app.use((req, res, next) => {
    res.locals.nonce = crypto.randomUUID();
    next();
});
app.use((0, helmet_1.default)({
    contentSecurityPolicy: {
        useDefaults: true,
        directives: {
            scriptSrc: [
                "'self'",
                (req, res) => `'nonce-${res.locals.nonce}'`,
            ],
            styleSec: [
                "'self'"
            ]
        },
    },
    crossOriginResourcePolicy: {
        policy: 'cross-origin'
    }
}));
app.use((req, res, next) => {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
});
app.use(body_parser_1.default.json());
app.use(body_parser_1.default.urlencoded({ extended: true }));
app.set('view engine', 'ejs');
app.use("/static", express_1.default.static(path_1.default.join(__dirname, '../static')));
app.use("/", routes_1.router);
app.listen(3223, '0.0.0.0', () => {
    (0, database_1.createTableUsers)();
    console.log('listening on http://0.0.0.0:3223');
});
