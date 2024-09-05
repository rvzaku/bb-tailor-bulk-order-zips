/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { GetUsersSuccessResponse } from '../models/GetUsersSuccessResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class UsersService {
    /**
     * Get all users with their metadata
     * Returns a list of users with their metadata.
     * @returns GetUsersSuccessResponse Successful response
     * @throws ApiError
     */
    public static getUsers(): CancelablePromise<GetUsersSuccessResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/users',
            errors: {
                401: `Unauthorized response`,
                403: `Forbidden response`,
            },
        });
    }
}
