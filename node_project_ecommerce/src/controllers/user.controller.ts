import { AsyncHandler } from '@/utils';

export const SignIn = AsyncHandler(async (request, response) => {
  response.send('SignIn Controller');
});

export const SignUp = AsyncHandler(async (request, response) => {
  response.send('SignUp Controller');
});

export const VerifyUser = AsyncHandler(async (request, response) => {
  response.send('VerifyUser Controller');
});

export const GetUserDetailById = AsyncHandler(async (request, response) => {
  response.send('GetUserDetailById Controller');
});

export const UpdateDetailById = AsyncHandler(async (request, response) => {
  response.send('UpdateDetailById Controller');
});

export const UpdatePassword = AsyncHandler(async (request, response) => {
  response.send('UpdatePassword Controller');
});

export const UpdateEmail = AsyncHandler(async (request, response) => {
  response.send('UpdateEmail Controller');
});

export const ResetPassword = AsyncHandler(async (request, response) => {
  response.send('ResetPassword Controller');
});

export const ForgotPassword = AsyncHandler(async (request, response) => {
  response.send('ForgotPassword Controller');
});

export const RefreshToken = AsyncHandler(async (request, response) => {
  response.send('RefreshToken Controller');
});

export const DeleteAccountById = AsyncHandler(async (request, response) => {
  response.send('DeleteAccountById Controller');
});
