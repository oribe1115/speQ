/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ContestInfo } from '../models/ContestInfo';
import type { ProblemInfo } from '../models/ProblemInfo';
import type { ProblemSolvedInfo } from '../models/ProblemSolvedInfo';
import type { ScoreInfo } from '../models/ScoreInfo';
import type { traPId } from '../models/traPId';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class AdminOnlyService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Put Admin Users
     * adminユーザー全体を更新する
     * @param requestBody
     * @returns traPId OK
     * @throws ApiError
     */
    public putAdminUsers(
        requestBody?: Array<traPId>,
    ): CancelablePromise<Array<traPId>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/admin/users/admin',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Put Contestant Users
     * contestantユーザ全体を更新する
     * @param requestBody
     * @returns traPId OK
     * @throws ApiError
     */
    public putContestants(
        requestBody?: Array<traPId>,
    ): CancelablePromise<Array<traPId>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/admin/users/contestant',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Put Contest Info
     * コンテスト情報全体を更新する
     * @param requestBody
     * @returns ContestInfo OK
     * @throws ApiError
     */
    public putContestInfo(
        requestBody?: ContestInfo,
    ): CancelablePromise<ContestInfo> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/admin/info',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Put Problems
     * 問題情報全体を更新
     * @param requestBody
     * @returns ProblemInfo OK
     * @throws ApiError
     */
    public putProblems(
        requestBody?: Array<ProblemInfo>,
    ): CancelablePromise<Array<ProblemInfo>> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/admin/problems',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Mark Problems as Solved
     * 問題が解かれたことをマークする
     * @param requestBody
     * @returns ProblemSolvedInfo OK
     * @throws ApiError
     */
    public markProblemAsSolved(
        requestBody?: ProblemSolvedInfo,
    ): CancelablePromise<ProblemSolvedInfo> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/admin/problems/solved',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Unmark Problem as Solved
     * 問題が解かれたというマークを解除する
     * @param requestBody
     * @returns any OK
     * @throws ApiError
     */
    public unmarkProblemAsSolved(
        requestBody?: ProblemSolvedInfo,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/admin/problems/solved',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Post New Score
     * 最新スコアを登録する
     * @param requestBody
     * @returns ScoreInfo OK
     * @throws ApiError
     */
    public postScore(
        requestBody?: ScoreInfo,
    ): CancelablePromise<ScoreInfo> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/admin/scores',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

}
