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
 * @interface SwimlaneControllerEditSwimlaneRequest
 */
export interface SwimlaneControllerEditSwimlaneRequest {
    /**
     * 
     * @type {number}
     * @memberof SwimlaneControllerEditSwimlaneRequest
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof SwimlaneControllerEditSwimlaneRequest
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof SwimlaneControllerEditSwimlaneRequest
     */
    order?: number;
}

/**
 * Check if a given object implements the SwimlaneControllerEditSwimlaneRequest interface.
 */
export function instanceOfSwimlaneControllerEditSwimlaneRequest(value: object): value is SwimlaneControllerEditSwimlaneRequest {
    return true;
}

export function SwimlaneControllerEditSwimlaneRequestFromJSON(json: any): SwimlaneControllerEditSwimlaneRequest {
    return SwimlaneControllerEditSwimlaneRequestFromJSONTyped(json, false);
}

export function SwimlaneControllerEditSwimlaneRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SwimlaneControllerEditSwimlaneRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'order': json['order'] == null ? undefined : json['order'],
    };
}

export function SwimlaneControllerEditSwimlaneRequestToJSON(json: any): SwimlaneControllerEditSwimlaneRequest {
    return SwimlaneControllerEditSwimlaneRequestToJSONTyped(json, false);
}

export function SwimlaneControllerEditSwimlaneRequestToJSONTyped(value?: SwimlaneControllerEditSwimlaneRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'name': value['name'],
        'order': value['order'],
    };
}
