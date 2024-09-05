import { AuthProvider, AuthActionResponse, OnErrorResponse } from '@refinedev/core'
import { AuthService } from '@services/openapi/services/AuthService'
import { LoginPayload } from '@services/openapi/models/LoginPayload'
import { LoginSuccessResponse } from '@services/openapi/models/LoginSuccessResponse'
import { HttpError } from '@refinedev/core'

export const authProvider: AuthProvider = {
    login: async (payload: LoginPayload): Promise<AuthActionResponse> => {
        try {
            const response: LoginSuccessResponse = await AuthService.postAuthLogin(payload)
            localStorage.setItem('user.accessToken', response.accessToken as string)
            localStorage.setItem('user.refreshToken', response.refreshToken as string)

            return {
                success: true,
                redirectTo: '/home',
            }
        } catch (error) {
            const httpError: HttpError = error as HttpError
            return { success: false, error: httpError }
        }
    },
    logout: async () => {
        try {
            localStorage.removeItem('user.accessToken')
            localStorage.removeItem('user.refreshToken')
            return { success: true }
        } catch (error) {
            const httpError: HttpError = error as HttpError
            return { success: false, error: httpError }
        }
    },
    check: async () => {
        const userAccessToken = localStorage.getItem('user.accessToken')
        const userRefreshToken = localStorage.getItem('user.refreshToken')
        if (!userAccessToken && !userRefreshToken) {
            return {
                authenticated: false,
            }
        }

        return {
            authenticated: true,
        }
    },
    onError: async (error: any): Promise<OnErrorResponse> => {
        console.error(error)
        return {
            error,
        }
    },
}
