import express from 'express';
import { router } from './routes';
import path from 'path';
import cookieParser from 'cookie-parser';
import session from 'express-session';
import bodyParser from 'body-parser';
import { createTableUsers } from './database';

const app: express.Application = express();

app.use(cookieParser());
app.use(session({
    secret: 'mysecret',
    resave: false,
    saveUninitialized: true,
    cookie: {
        secure: false,
        httpOnly: true,
        sameSite: 'lax'
    }
}));

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.set('view engine', 'ejs');
app.use("/static", express.static(path.join(__dirname, '../static')));

app.use("/", router);

app.listen(3223, '0.0.0.0', () => {
    createTableUsers();
    console.log('listening on http://0.0.0.0:3223');
});