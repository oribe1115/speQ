/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { traPId } from '../models/traPId';
import type { UserInfo } from '../models/UserInfo';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class UserService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Get My Information
     * 現在アクセスしているユーザーの情報を取得する
     * @returns UserInfo OK
     * @throws ApiError
     */
    public getMe(): CancelablePromise<UserInfo> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users/me',
            errors: {
                401: `traPIDが確認できません`,
            },
        });
    }

    /**
     * Get Root Users
     * rootユーザーを取得する
     * @param requestBody
     * @returns traPId OK
     * @throws ApiError
     */
    public getRootUsers(
        requestBody?: any,
    ): CancelablePromise<traPId> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users/root',
            body: requestBody,
        });
    }

    /**
     * Get Admin Users
     * adminユーザーを取得する
     * @returns traPId OK
     * @throws ApiError
     */
    public getAdmins(): CancelablePromise<Array<traPId>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users/admin',
        });
    }

    /**
     * Get Contestant Users
     * contestantユーザーを取得する
     * @param requestBody
     * @returns traPId OK
     * @throws ApiError
     */
    public getContestants(
        requestBody?: Array<traPId>,
    ): CancelablePromise<Array<traPId>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users/contestant',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

}
