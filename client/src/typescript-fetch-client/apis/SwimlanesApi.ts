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
  ModelsErrorResponse,
  SwimlaneControllerCreateSwimlaneRequest,
  SwimlaneControllerCreateSwimlaneResponse,
  SwimlaneControllerDeleteSwimlaneRequest,
  SwimlaneControllerDeleteSwimlaneResponse,
  SwimlaneControllerEditSwimlaneRequest,
  SwimlaneControllerEditSwimlaneResponse,
} from '../models/index';
import {
    ModelsErrorResponseFromJSON,
    ModelsErrorResponseToJSON,
    SwimlaneControllerCreateSwimlaneRequestFromJSON,
    SwimlaneControllerCreateSwimlaneRequestToJSON,
    SwimlaneControllerCreateSwimlaneResponseFromJSON,
    SwimlaneControllerCreateSwimlaneResponseToJSON,
    SwimlaneControllerDeleteSwimlaneRequestFromJSON,
    SwimlaneControllerDeleteSwimlaneRequestToJSON,
    SwimlaneControllerDeleteSwimlaneResponseFromJSON,
    SwimlaneControllerDeleteSwimlaneResponseToJSON,
    SwimlaneControllerEditSwimlaneRequestFromJSON,
    SwimlaneControllerEditSwimlaneRequestToJSON,
    SwimlaneControllerEditSwimlaneResponseFromJSON,
    SwimlaneControllerEditSwimlaneResponseToJSON,
} from '../models/index';

export interface ApiV1SwimlanesCreatePostRequest {
    request: SwimlaneControllerCreateSwimlaneRequest;
}

export interface ApiV1SwimlanesDeletePostRequest {
    request: SwimlaneControllerDeleteSwimlaneRequest;
}

export interface ApiV1SwimlanesEditPostRequest {
    request: SwimlaneControllerEditSwimlaneRequest;
}

/**
 * 
 */
export class SwimlanesApi extends runtime.BaseAPI {

    /**
     * Create a swimlane for a board
     * Create a swimlane
     */
    async apiV1SwimlanesCreatePostRaw(requestParameters: ApiV1SwimlanesCreatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SwimlaneControllerCreateSwimlaneResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiV1SwimlanesCreatePost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["X-CSRF-Token"] = await this.configuration.apiKey("X-CSRF-Token"); // csrf authentication
        }

        const response = await this.request({
            path: `/api/v1/swimlanes/create`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SwimlaneControllerCreateSwimlaneRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SwimlaneControllerCreateSwimlaneResponseFromJSON(jsonValue));
    }

    /**
     * Create a swimlane for a board
     * Create a swimlane
     */
    async apiV1SwimlanesCreatePost(requestParameters: ApiV1SwimlanesCreatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SwimlaneControllerCreateSwimlaneResponse> {
        const response = await this.apiV1SwimlanesCreatePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete a swimlane by ID
     * Delete a swimlane
     */
    async apiV1SwimlanesDeletePostRaw(requestParameters: ApiV1SwimlanesDeletePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SwimlaneControllerDeleteSwimlaneResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiV1SwimlanesDeletePost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["X-CSRF-Token"] = await this.configuration.apiKey("X-CSRF-Token"); // csrf authentication
        }

        const response = await this.request({
            path: `/api/v1/swimlanes/delete`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SwimlaneControllerDeleteSwimlaneRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SwimlaneControllerDeleteSwimlaneResponseFromJSON(jsonValue));
    }

    /**
     * Delete a swimlane by ID
     * Delete a swimlane
     */
    async apiV1SwimlanesDeletePost(requestParameters: ApiV1SwimlanesDeletePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SwimlaneControllerDeleteSwimlaneResponse> {
        const response = await this.apiV1SwimlanesDeletePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit a swimlane by ID
     * Edit a swimlane
     */
    async apiV1SwimlanesEditPostRaw(requestParameters: ApiV1SwimlanesEditPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SwimlaneControllerEditSwimlaneResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiV1SwimlanesEditPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["X-CSRF-Token"] = await this.configuration.apiKey("X-CSRF-Token"); // csrf authentication
        }

        const response = await this.request({
            path: `/api/v1/swimlanes/edit`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SwimlaneControllerEditSwimlaneRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SwimlaneControllerEditSwimlaneResponseFromJSON(jsonValue));
    }

    /**
     * Edit a swimlane by ID
     * Edit a swimlane
     */
    async apiV1SwimlanesEditPost(requestParameters: ApiV1SwimlanesEditPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SwimlaneControllerEditSwimlaneResponse> {
        const response = await this.apiV1SwimlanesEditPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

}