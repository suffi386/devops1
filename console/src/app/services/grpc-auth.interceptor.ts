import { Injectable } from '@angular/core';
import { Metadata } from 'grpc-web';

import { GrpcHandler } from './grpc-handler';
import { GrpcInterceptor } from './grpc-interceptor';
import { StorageService } from './storage.service';

const authorizationKey = 'Authorization';
const bearerPrefix = 'Bearer ';
const accessTokenStorageField = 'access_token';

@Injectable({ providedIn: 'root' })
export class GrpcAuthInterceptor implements GrpcInterceptor {
    constructor(private readonly authStorage: StorageService) { }

    public async intercept(
        req: unknown,
        metadata: Metadata,
        next: GrpcHandler,
    ): Promise<any> {
        if (!metadata[authorizationKey]) {
            const accessToken = this.authStorage.getItem(accessTokenStorageField);
            if (accessToken) {
                metadata[authorizationKey] = bearerPrefix + accessToken;
            }
        }

        return await next.handle(req, metadata);
        // .catch(error => {
        //     if (error.code === 7) {
        //         this.authService.authenticate();
        //     }
        // });
    }
}
