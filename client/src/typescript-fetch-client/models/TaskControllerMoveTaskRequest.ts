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
 * @interface TaskControllerMoveTaskRequest
 */
export interface TaskControllerMoveTaskRequest {
    /**
     * 
     * @type {number}
     * @memberof TaskControllerMoveTaskRequest
     */
    columnId: number;
    /**
     * 
     * @type {number}
     * @memberof TaskControllerMoveTaskRequest
     */
    position?: number;
    /**
     * 
     * @type {number}
     * @memberof TaskControllerMoveTaskRequest
     */
    swimlaneId: number;
    /**
     * 
     * @type {number}
     * @memberof TaskControllerMoveTaskRequest
     */
    taskId: number;
}

/**
 * Check if a given object implements the TaskControllerMoveTaskRequest interface.
 */
export function instanceOfTaskControllerMoveTaskRequest(value: object): value is TaskControllerMoveTaskRequest {
    if (!('columnId' in value) || value['columnId'] === undefined) return false;
    if (!('swimlaneId' in value) || value['swimlaneId'] === undefined) return false;
    if (!('taskId' in value) || value['taskId'] === undefined) return false;
    return true;
}

export function TaskControllerMoveTaskRequestFromJSON(json: any): TaskControllerMoveTaskRequest {
    return TaskControllerMoveTaskRequestFromJSONTyped(json, false);
}

export function TaskControllerMoveTaskRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): TaskControllerMoveTaskRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'columnId': json['column_id'],
        'position': json['position'] == null ? undefined : json['position'],
        'swimlaneId': json['swimlane_id'],
        'taskId': json['task_id'],
    };
}

export function TaskControllerMoveTaskRequestToJSON(json: any): TaskControllerMoveTaskRequest {
    return TaskControllerMoveTaskRequestToJSONTyped(json, false);
}

export function TaskControllerMoveTaskRequestToJSONTyped(value?: TaskControllerMoveTaskRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'column_id': value['columnId'],
        'position': value['position'],
        'swimlane_id': value['swimlaneId'],
        'task_id': value['taskId'],
    };
}
