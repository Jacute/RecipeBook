import express from 'express';
import { router } from './routes';
import path from 'path';
import cookieParser from 'cookie-parser';
import session from 'express-session';
import bodyParser from 'body-parser';
import { createTableUsers } from './database';
import helmet, { contentSecurityPolicy } from 'helmet';

const app: express.Application = express();

app.use(cookieParser());
app.use(session({
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
})

app.use(helmet.contentSecurityPolicy({
    useDefaults: true,
    directives: {
        scriptSrc: [
            "'self'",
            (req, res) => `'nonce-${res.locals.nonce}'`,
        ],
        styleSrc: [
            "'self'"
        ]
    }
));

app.use((req, res, next) => {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
  });

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.set('view engine', 'ejs');
app.use("/static", express.static(path.join(__dirname, '../static')));

app.use("/", router);

app.listen(3223, '0.0.0.0', () => {
    createTableUsers();
    console.log('listening on http://0.0.0.0:3223');
});