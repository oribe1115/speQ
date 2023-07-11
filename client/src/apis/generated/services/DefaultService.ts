/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { TripleVote } from '../models/TripleVote';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class DefaultService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * 現在の3連単予想を取得する
     * @returns TripleVote OK
     * @throws ApiError
     */
    public getVoteTriple(): CancelablePromise<TripleVote> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/vote/triple',
        });
    }

}
