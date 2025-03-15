import express, { Express } from 'express';
import cookieParser from 'cookie-parser';
import cors from 'cors';
import { config } from '@/utils';
import initialRoute from '@/routers';

const app: Express = express();

app.use(cors({ origin: config.CORS_ORIGIN, credentials: true }));
app.use(express.json());
app.use(cookieParser());
app.use(express.urlencoded({ extended: true, limit: '16kb' }));

app.use('/api/v1', initialRoute);

export default app;
