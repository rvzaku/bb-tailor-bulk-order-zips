/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { LoginPayload } from '../models/LoginPayload';
import type { LoginSuccessResponse } from '../models/LoginSuccessResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class AuthService {
    /**
     * User Login
     * Logs in a user.
     * @param requestBody User login payload
     * @returns LoginSuccessResponse User logged in successfully, JWT token pair generated.
     * @throws ApiError
     */
    public static postAuthLogin(
        requestBody: LoginPayload,
    ): CancelablePromise<LoginSuccessResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/login',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `User performed a bad request, request body validation failure.`,
                401: `User not authorized, email and password mismatch.`,
            },
        });
    }
}
