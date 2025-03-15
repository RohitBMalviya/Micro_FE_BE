import { AsyncHandler } from '@/utils';

export const AuthMiddleware = AsyncHandler(async (request, response, next) => {
  console.log('Auth Middleware');
  next!();
});

export default AuthMiddleware;
