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
 * @interface AuthControllerConfirmTOTPRequest
 */
export interface AuthControllerConfirmTOTPRequest {
    /**
     * 
     * @type {string}
     * @memberof AuthControllerConfirmTOTPRequest
     */
    code: string;
}

/**
 * Check if a given object implements the AuthControllerConfirmTOTPRequest interface.
 */
export function instanceOfAuthControllerConfirmTOTPRequest(value: object): value is AuthControllerConfirmTOTPRequest {
    if (!('code' in value) || value['code'] === undefined) return false;
    return true;
}

export function AuthControllerConfirmTOTPRequestFromJSON(json: any): AuthControllerConfirmTOTPRequest {
    return AuthControllerConfirmTOTPRequestFromJSONTyped(json, false);
}

export function AuthControllerConfirmTOTPRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthControllerConfirmTOTPRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'],
    };
}

export function AuthControllerConfirmTOTPRequestToJSON(json: any): AuthControllerConfirmTOTPRequest {
    return AuthControllerConfirmTOTPRequestToJSONTyped(json, false);
}

export function AuthControllerConfirmTOTPRequestToJSONTyped(value?: AuthControllerConfirmTOTPRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'code': value['code'],
    };
}

