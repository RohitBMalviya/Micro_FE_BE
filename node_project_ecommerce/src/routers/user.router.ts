import { Router } from 'express';

const userRouter: Router = Router();

userRouter.post('/sign-in');
userRouter.post('/sign-up');
userRouter.post('/verify-user');
userRouter.patch('/update-password');
userRouter.post('/refresh-token');
userRouter.patch('/update-email');
userRouter.post('/reset-password');
userRouter.post('/forgot-password');
userRouter.delete('/delete-account');
userRouter.route('/user-details').get().patch().put();

export default userRouter;
