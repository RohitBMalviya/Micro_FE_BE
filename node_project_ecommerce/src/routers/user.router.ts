import {
  SignIn,
  SignUp,
  VerifyUser,
  GetUserDetailById,
  UpdateDetailById,
  UpdatePassword,
  UpdateEmail,
  ResetPassword,
  ForgotPassword,
  RefreshToken,
  DeleteAccountById,
} from '@/controllers';
import { AuthMiddleware } from '@/middleware';
import { Router } from 'express';

const userRouter: Router = Router();

userRouter.post('/sign-in', SignIn);
userRouter.post('/sign-up', SignUp);
userRouter.post('/verify-user', VerifyUser);
userRouter.patch('/update-password', AuthMiddleware, UpdatePassword);
userRouter.post('/refresh-token', RefreshToken);
userRouter.patch('/update-email', AuthMiddleware, UpdateEmail);
userRouter.post('/reset-password', ResetPassword);
userRouter.post('/forgot-password', ForgotPassword);
userRouter.delete('/delete-account', AuthMiddleware, DeleteAccountById);
userRouter
  .use(AuthMiddleware)
  .route('/user-details')
  .get(GetUserDetailById)
  .patch(UpdateDetailById);

export default userRouter;
