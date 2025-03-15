import { Router } from 'express';
import userRouter from '@/routers/user.router';

const initialRoute: Router = Router();

initialRoute.use('/auth', userRouter);

export default initialRoute;
