/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ContestInfo } from '../models/ContestInfo';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class ContestService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Get Contest Info
     * コンテスト情報を取得する
     * @returns ContestInfo OK
     * @throws ApiError
     */
    public getContestInfo(): CancelablePromise<ContestInfo> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/info',
        });
    }

}
