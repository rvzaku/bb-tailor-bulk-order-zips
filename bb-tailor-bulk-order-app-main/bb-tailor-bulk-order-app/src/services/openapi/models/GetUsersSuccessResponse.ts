/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * The list of users with their metadata
 */
export type GetUsersSuccessResponse = Array<{
    /**
     * The ID of the user
     */
    ID?: string;
    /**
     * The email of the user
     */
    Email?: string;
    /**
     * The list of roles of the user
     */
    Roles?: Array<string>;
    Profile?: {
        /**
         * The profile ID of the user
         */
        ID?: string;
        /**
         * The first name of the user
         */
        FirstName?: string;
        /**
         * The last name of the user
         */
        LastName?: string;
        /**
         * The phone number of the user
         */
        Phone?: string;
        /**
         * The age of the user
         */
        Age?: number;
        /**
         * The gender of the user
         */
        Gender?: string;
    };
    /**
     * The created at date-time of the user
     */
    CreatedAt?: string;
    /**
     * The updated at date-time of the user
     */
    UpdatedAt?: string;
}>;
