/* tslint:disable */
/* eslint-disable */
/**
 * Server API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ColumnControllerCreateColumnRequest,
  ColumnControllerCreateColumnResponse,
  ColumnControllerDeleteColumnRequest,
  ColumnControllerDeleteColumnResponse,
  ColumnControllerEditColumnRequest,
  ColumnControllerEditColumnResponse,
  ModelsErrorResponse,
} from '../models/index';
import {
    ColumnControllerCreateColumnRequestFromJSON,
    ColumnControllerCreateColumnRequestToJSON,
    ColumnControllerCreateColumnResponseFromJSON,
    ColumnControllerCreateColumnResponseToJSON,
    ColumnControllerDeleteColumnRequestFromJSON,
    ColumnControllerDeleteColumnRequestToJSON,
    ColumnControllerDeleteColumnResponseFromJSON,
    ColumnControllerDeleteColumnResponseToJSON,
    ColumnControllerEditColumnRequestFromJSON,
    ColumnControllerEditColumnRequestToJSON,
    ColumnControllerEditColumnResponseFromJSON,
    ColumnControllerEditColumnResponseToJSON,
    ModelsErrorResponseFromJSON,
    ModelsErrorResponseToJSON,
} from '../models/index';

export interface ApiV1ColumnsCreatePostRequest {
    request: ColumnControllerCreateColumnRequest;
}

export interface ApiV1ColumnsDeletePostRequest {
    request: ColumnControllerDeleteColumnRequest;
}

export interface ApiV1ColumnsEditPostRequest {
    request: ColumnControllerEditColumnRequest;
}

/**
 * 
 */
export class ColumnsApi extends runtime.BaseAPI {

    /**
     * Create a column for a board
     * Create a column
     */
    async apiV1ColumnsCreatePostRaw(requestParameters: ApiV1ColumnsCreatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ColumnControllerCreateColumnResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiV1ColumnsCreatePost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["X-CSRF-Token"] = await this.configuration.apiKey("X-CSRF-Token"); // csrf authentication
        }

        const response = await this.request({
            path: `/api/v1/columns/create`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ColumnControllerCreateColumnRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ColumnControllerCreateColumnResponseFromJSON(jsonValue));
    }

    /**
     * Create a column for a board
     * Create a column
     */
    async apiV1ColumnsCreatePost(requestParameters: ApiV1ColumnsCreatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ColumnControllerCreateColumnResponse> {
        const response = await this.apiV1ColumnsCreatePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete a column by ID
     * Delete a column
     */
    async apiV1ColumnsDeletePostRaw(requestParameters: ApiV1ColumnsDeletePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ColumnControllerDeleteColumnResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiV1ColumnsDeletePost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["X-CSRF-Token"] = await this.configuration.apiKey("X-CSRF-Token"); // csrf authentication
        }

        const response = await this.request({
            path: `/api/v1/columns/delete`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ColumnControllerDeleteColumnRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ColumnControllerDeleteColumnResponseFromJSON(jsonValue));
    }

    /**
     * Delete a column by ID
     * Delete a column
     */
    async apiV1ColumnsDeletePost(requestParameters: ApiV1ColumnsDeletePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ColumnControllerDeleteColumnResponse> {
        const response = await this.apiV1ColumnsDeletePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit a column by ID
     * Edit a column
     */
    async apiV1ColumnsEditPostRaw(requestParameters: ApiV1ColumnsEditPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ColumnControllerEditColumnResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiV1ColumnsEditPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["X-CSRF-Token"] = await this.configuration.apiKey("X-CSRF-Token"); // csrf authentication
        }

        const response = await this.request({
            path: `/api/v1/columns/edit`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ColumnControllerEditColumnRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ColumnControllerEditColumnResponseFromJSON(jsonValue));
    }

    /**
     * Edit a column by ID
     * Edit a column
     */
    async apiV1ColumnsEditPost(requestParameters: ApiV1ColumnsEditPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ColumnControllerEditColumnResponse> {
        const response = await this.apiV1ColumnsEditPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

}