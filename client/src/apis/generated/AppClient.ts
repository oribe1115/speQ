/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseHttpRequest } from './core/BaseHttpRequest';
import type { OpenAPIConfig } from './core/OpenAPI';
import { FetchHttpRequest } from './core/FetchHttpRequest';

import { AdminOnlyService } from './services/AdminOnlyService';
import { ContestService } from './services/ContestService';
import { UserService } from './services/UserService';
import { VoteService } from './services/VoteService';

type HttpRequestConstructor = new (config: OpenAPIConfig) => BaseHttpRequest;

export class AppClient {

    public readonly adminOnly: AdminOnlyService;
    public readonly contest: ContestService;
    public readonly user: UserService;
    public readonly vote: VoteService;

    public readonly request: BaseHttpRequest;

    constructor(config?: Partial<OpenAPIConfig>, HttpRequest: HttpRequestConstructor = FetchHttpRequest) {
        this.request = new HttpRequest({
            BASE: config?.BASE ?? '/api',
            VERSION: config?.VERSION ?? '1.0',
            WITH_CREDENTIALS: config?.WITH_CREDENTIALS ?? false,
            CREDENTIALS: config?.CREDENTIALS ?? 'include',
            TOKEN: config?.TOKEN,
            USERNAME: config?.USERNAME,
            PASSWORD: config?.PASSWORD,
            HEADERS: config?.HEADERS,
            ENCODE_PATH: config?.ENCODE_PATH,
        });

        this.adminOnly = new AdminOnlyService(this.request);
        this.contest = new ContestService(this.request);
        this.user = new UserService(this.request);
        this.vote = new VoteService(this.request);
    }
}

