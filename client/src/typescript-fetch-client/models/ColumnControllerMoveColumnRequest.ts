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
 * @interface ColumnControllerMoveColumnRequest
 */
export interface ColumnControllerMoveColumnRequest {
    /**
     * 
     * @type {string}
     * @memberof ColumnControllerMoveColumnRequest
     */
    direction: ColumnControllerMoveColumnRequestDirectionEnum;
    /**
     * 
     * @type {number}
     * @memberof ColumnControllerMoveColumnRequest
     */
    id: number;
    /**
     * 
     * @type {number}
     * @memberof ColumnControllerMoveColumnRequest
     */
    relativeId: number;
}


/**
 * @export
 */
export const ColumnControllerMoveColumnRequestDirectionEnum = {
    Before: 'before',
    After: 'after'
} as const;
export type ColumnControllerMoveColumnRequestDirectionEnum = typeof ColumnControllerMoveColumnRequestDirectionEnum[keyof typeof ColumnControllerMoveColumnRequestDirectionEnum];


/**
 * Check if a given object implements the ColumnControllerMoveColumnRequest interface.
 */
export function instanceOfColumnControllerMoveColumnRequest(value: object): value is ColumnControllerMoveColumnRequest {
    if (!('direction' in value) || value['direction'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('relativeId' in value) || value['relativeId'] === undefined) return false;
    return true;
}

export function ColumnControllerMoveColumnRequestFromJSON(json: any): ColumnControllerMoveColumnRequest {
    return ColumnControllerMoveColumnRequestFromJSONTyped(json, false);
}

export function ColumnControllerMoveColumnRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ColumnControllerMoveColumnRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'direction': json['direction'],
        'id': json['id'],
        'relativeId': json['relative_id'],
    };
}

export function ColumnControllerMoveColumnRequestToJSON(json: any): ColumnControllerMoveColumnRequest {
    return ColumnControllerMoveColumnRequestToJSONTyped(json, false);
}

export function ColumnControllerMoveColumnRequestToJSONTyped(value?: ColumnControllerMoveColumnRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'direction': value['direction'],
        'id': value['id'],
        'relative_id': value['relativeId'],
    };
}

