import { NextFunction, Request, Response } from 'express';

export default function AsyncHandler(
  requestHandler: (
    request: Request,
    response: Response,
    next?: NextFunction
  ) => Promise<void> | void
) {
  return (request: Request, response: Response, next: NextFunction) => {
    Promise.resolve(requestHandler(request, response, next)).catch((err) =>
      next(err)
    );
  };
}
