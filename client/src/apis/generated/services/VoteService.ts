/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { traPId } from '../models/traPId';
import type { TripleVote } from '../models/TripleVote';
import type { VoteStatsItem } from '../models/VoteStatsItem';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class VoteService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Get Current Vote
     * 現在のvote対象を取得する
     * @returns traPId OK
     * @throws ApiError
     */
    public getVote(): CancelablePromise<traPId> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/vote',
        });
    }

    /**
     * Post New Vote
     * 新しいvote対象を登録する
     * @param requestBody
     * @returns traPId OK
     * @throws ApiError
     */
    public postVote(
        requestBody?: traPId,
    ): CancelablePromise<traPId> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/vote',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * 三連単を投票する
     * @param requestBody
     * @returns any OK
     * @throws ApiError
     */
    public postVoteTriple(
        requestBody?: TripleVote,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/vote/triple',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Your GET endpoint
     * 1位予想の統計情報
     * @returns VoteStatsItem OK
     * @throws ApiError
     */
    public getVoteStats(): CancelablePromise<Array<VoteStatsItem>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/vote/stats',
        });
    }

}
