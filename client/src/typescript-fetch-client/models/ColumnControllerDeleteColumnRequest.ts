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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ColumnControllerDeleteColumnRequest
 */
export interface ColumnControllerDeleteColumnRequest {
    /**
     * 
     * @type {number}
     * @memberof ColumnControllerDeleteColumnRequest
     */
    id?: number;
}

/**
 * Check if a given object implements the ColumnControllerDeleteColumnRequest interface.
 */
export function instanceOfColumnControllerDeleteColumnRequest(value: object): value is ColumnControllerDeleteColumnRequest {
    return true;
}

export function ColumnControllerDeleteColumnRequestFromJSON(json: any): ColumnControllerDeleteColumnRequest {
    return ColumnControllerDeleteColumnRequestFromJSONTyped(json, false);
}

export function ColumnControllerDeleteColumnRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ColumnControllerDeleteColumnRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
    };
}

export function ColumnControllerDeleteColumnRequestToJSON(json: any): ColumnControllerDeleteColumnRequest {
    return ColumnControllerDeleteColumnRequestToJSONTyped(json, false);
}

export function ColumnControllerDeleteColumnRequestToJSONTyped(value?: ColumnControllerDeleteColumnRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
    };
}
